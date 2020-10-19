package utils

import (
	"github.com/harjyotbagga/meeting-scheduler-api/models"
)

func db_setup() {
	c := GetClient()
	participant := models.Participant{
		Name:  "Participant-A",
		Email: "participant-a@gmail.com",
		RSVP:  "yes",
	}
	CreateParticipant(c, participant)
}
