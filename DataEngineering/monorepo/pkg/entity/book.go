package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Author    string             `bson:"author"`
	PageCount int                `bson:"page_count"`
}

type Author struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName string             `bson:"full_name"`
}
