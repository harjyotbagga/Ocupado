package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meeting struct {
	ID               primitive.ObjectID `bson:"_id"`
	Title            string             `json:"title,omitempty" bson:"title"`
	ParticipantsList []*Participant     `json:"participantsList" bson:"participantsList"`
	StartTime        string             `json:"startTime,omitempty" bson:"startTime"`
	EndTime          string             `json:"endTime,omitempty" bson:"endTime"`
	CreationTime     string             `json:"creationTime,omitempty" bson:"creationTime"`
}

type MeetingScheduled struct {
	Meeting_ID string `json:"meetingID"`
}
