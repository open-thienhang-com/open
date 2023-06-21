package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title,omitempty" bson:"title,omitempty"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	Image       string               `json:"image,omitempty" bson:"image,omitempty"`
	Price       int                  `json:"price,omitempty" bson:"price,omitempty"`
	Lessions    []Lession            `json:"lessions" bson:"lessions,omitempty"`
	Lecturer    []primitive.ObjectID `json:"lecturer,omitempty"`
	Mentors     []primitive.ObjectID `json:"mentors,omitempty"`
	Judges      []primitive.ObjectID `json:"judges,omitempty"`
	Members     []primitive.ObjectID `json:"members,omitempty"`
	Created_at  time.Time            `json:"created_at"`
	Updated_at  time.Time            `json:"updated_at"`
}

type Lession struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
