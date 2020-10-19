package utils

import (
	"context"
	"log"

	"github.com/harjyotbagga/meeting-scheduler-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReturnParticipants(client *mongo.Client, filter bson.M) []*models.Participant {
	var participants []*models.Participant
	collection := client.Database("meeting-scheduler").Collection("participants")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the participants", err)
	}
	for cur.Next(context.TODO()) {
		var participant models.Participant
		err := cur.Decode(&participant)
		if err != nil {
			log.Fatal("Error on Decoding the participant", err)
		}
		participants = append(participants, &participant)
	}
	return participants
}

func GetParticipant(client *mongo.Client, filter bson.M) models.Participant {
	var participant models.Participant
	collection := client.Database("meeting-scheduler").Collection("participants")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&participant)
	return participant
}

func CreateParticipant(client *mongo.Client, participant models.Participant) interface{} {
	collection := client.Database("meeting-scheduler").Collection("participants")
	insertResult, err := collection.InsertOne(context.TODO(), participant)
	if err != nil {
		log.Fatalln("Error on inserting new Participant", err)
	}
	return insertResult.InsertedID
}
