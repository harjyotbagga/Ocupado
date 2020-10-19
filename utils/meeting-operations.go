package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/harjyotbagga/meeting-scheduler-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ScheduleNewMeeting(client *mongo.Client, meeting models.Meeting) (interface{}, bool) {
	err_status := false
	collection := client.Database("meeting-scheduler").Collection("meetings")
	insertResult, err := collection.InsertOne(context.TODO(), meeting)
	if err != nil {
		log.Fatalln("Error on creating new meeting", err)
		err_status = true
	}
	return insertResult.InsertedID, err_status
}

func ReturnMeetings(client *mongo.Client, filter bson.M) ([]*models.Meeting, bool) {
	err_status := false
	var meetings []*models.Meeting
	collection := client.Database("meeting-scheduler").Collection("meetings")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the participants", err)
		err_status = true
	}
	for cur.Next(context.TODO()) {
		var meeting models.Meeting
		err := cur.Decode(&meeting)
		if err != nil {
			log.Fatal("Error on Decoding the meeting", err)
			err_status = true
		}

		meetings = append(meetings, &meeting)
	}
	return meetings, err_status
}

func ReturnMeeting(client *mongo.Client, filter bson.M) (models.Meeting, bool) {
	err_status := false
	var meeting models.Meeting
	collection := client.Database("meeting-scheduler").Collection("meetings")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&meeting)
	return meeting, err_status
}

func GetMeetingsInDuration(client *mongo.Client, start time.Time, end time.Time) ([]*models.Meeting, bool) {
	err_status := false
	filter := bson.D{{
		"startTime",
		bson.D{{
			"$gte", start,
		}}}, {
		"endTime",
		bson.D{{
			"$lte", end,
		}},
	}}
	var meetings []*models.Meeting
	collection := client.Database("meeting-scheduler").Collection("meetings")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error while querying DB", err)
		err_status = true
	}
	err = cur.All(context.TODO(), &meetings)
	if err != nil {
		log.Fatal("Error while querying DB", err)
		err_status = true
	}
	return meetings, err_status
}

func GetMeetingsOfParticipant(client *mongo.Client, participant_id string) ([]*models.Meeting, bool) {
	err_status := false
	var meetings []*models.Meeting
	collection := client.Database("meeting-scheduler").Collection("participants")
	unwindStage := bson.D{{
		"$unwind", "$participants",
	}}

	matchStage := bson.D{{
		"$match", bson.M{
			"participants.email": participant_id,
		},
	}}
	meetingsCursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{unwindStage, matchStage})

	if err = meetingsCursor.All(context.TODO(), &meetings); err != nil {
		log.Fatal(err.Error())
		err_status = true
	}
	fmt.Println(meetings)
	return meetings, err_status
}
