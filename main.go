package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/harjyotbagga/meeting-scheduler-api/models"
	"github.com/harjyotbagga/meeting-scheduler-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	c := utils.GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected to the DB!")
	}

	// db_setup()

	http.HandleFunc("/", apiStatus)
	http.HandleFunc("/meetings", meetings_handler)
	http.HandleFunc("/meeting", meeting_handler)
	fmt.Println("Server Starting....")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func meeting_handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		meeting_id := utils.ExtractParam(r)
		getMeeting(w, r, meeting_id)
	default:
		utils.InvalidRequest(w, 403, "Not a Valid Endpoint!!")
	}
}

func meetings_handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		params := r.URL.Query()
		if len(params) != 0 {
			utils.InvalidRequest(w, 400, "Invalid Endpoint! Please check and retry again!")
		} else {
			req_type := r.Header.Get("Content-Type")
			if req_type != "application/json" {
				utils.InvalidRequest(w, 400, "Please ONLY send JSON queries at this endpoint.")
			} else {
				scheduleMeeting(w, r)
			}
		}
	case "GET":
		params := r.URL.Query()
		participantArr, participant_query := params["participant"]
		if !participant_query || len(participantArr) < 1 {
			startArr, noStart := params["start"]
			endArr, noEnd := params["end"]
			if !noStart || !noEnd || len(startArr) < 1 || len(endArr) < 1 {
				utils.InvalidRequest(w, 400, "Invalid Request! Please check and retry again!")
				return
			}
			start, _ := time.Parse("RFC", startArr[0])
			end, _ := time.Parse("RFC", endArr[0])
			getMeetingsinDuration(w, r, start, end)
		} else if participant_query && len(participantArr) == 1 {
			participant_id := participantArr[0]
			getMeetingsOfParticipant(w, r, participant_id)
		}

	default:
		utils.InvalidRequest(w, 400, "Invalid Request! Please check and retry again!")
	}

}

func apiStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World! SERVER IS UP & RUNNING"}`))
}

func scheduleMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c := utils.GetClient()
	var meeting models.Meeting
	err := json.NewDecoder(r.Body).Decode(&meeting)
	if err != nil {
		panic(err)
	}
	meeting.CreationTime = time.Now().Format("RFC")
	meeting.ID = primitive.NewObjectID()
	meeting.ParticipantsList = utils.ReturnParticipants(c, bson.M{})
	fmt.Println(meeting.ParticipantsList)
	insertedID, err_status := utils.ScheduleNewMeeting(c, meeting)
	if err_status {
		utils.InternalError(w)
	}
	if err != nil {
		panic(err)
		return
	}
	scheduled_meeting := models.MeetingScheduled{Meeting_ID: insertedID.(primitive.ObjectID).Hex()}
	json.NewEncoder(w).Encode(scheduled_meeting)
}

func getMeeting(w http.ResponseWriter, r *http.Request, meeting_id string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c := utils.GetClient()
	id := utils.ExtractParam(r)
	meeting, err_status := utils.ReturnMeeting(c, bson.M{"_id": id})
	if err_status {
		utils.InternalError(w)
		return
	}
	json.NewEncoder(w).Encode(meeting)
}

func getMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c := utils.GetClient()
	meetings, err_status := utils.ReturnMeetings(c, bson.M{})
	if err_status {
		utils.InternalError(w)
		return
	}
	json.NewEncoder(w).Encode(meetings)
}

func getParticipants(w http.ResponseWriter, r *http.Request) {
	c := utils.GetClient()
	participants := utils.ReturnParticipants(c, bson.M{})
	for _, participant := range participants {
		log.Println(participant.Name, participant.Email)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(participants)
}

func getMeetingsinDuration(w http.ResponseWriter, r *http.Request, start time.Time, end time.Time) {
	c := utils.GetClient()
	meetings, err_status := utils.GetMeetingsInDuration(c, start, end)
	if err_status {
		utils.InternalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meetings)
}

func getMeetingsOfParticipant(w http.ResponseWriter, r *http.Request, participant_id string) {
	c := utils.GetClient()
	meetings, err_status := utils.GetMeetingsOfParticipant(c, participant_id)
	if err_status {
		utils.InternalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meetings)
}
