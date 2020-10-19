package models

type Participant struct {
	Name  string `json:"name,omitempty" bson:"name"`
	Email string `json:"email,omitempty" bson:"email"`
	RSVP  string `json:"rsvp,omitempty" bson:"rsvp"`
}
