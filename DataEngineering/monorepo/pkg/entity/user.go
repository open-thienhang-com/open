/*
Same as Entities or models, will used in all layer.
This layer, will store any Object’s Struct and its method.
Example : Article, Student, Book.
*/
package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Education struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	From        string             `bson:"from" json:"from"`
	To          string             `bson:"to" json:"to"` // becase can be "present"
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Major       string             `bson:"major" json:"major"`
	Grade       float32            `bson:"grade" json:"grade,string,omitempty"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}

type Experience struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	From           string             `json:"from"`
	To             string             `json:"to"` // becase can be "present"
	Title          string             `json:"title"`
	Company        string             `json:"company"`
	Responsibility string             `json:"responsibility"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
}

type Skill struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name"`
	Level      int16              `json:"level"` // from 1 to 5
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type Paper struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `json:"name"`
	Authors       []string           `json:"authors"`
	PublishedDate string             `json:"publishedDate"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}

type Reference struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name"`
	Email      string             `json:"email"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type Award struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name"`
	Time       string             `json:"time"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type Qualification struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name"`
	IssueDate  string             `bson:"issueDate" json:"issueDate"`
	ExpireAt   string             `bson:"expireAt" json:"expireAt"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type DBRef struct {
	Ref        interface{} `bson:"ref" json:"ref"`
	ID         interface{} `bson:"id" json:"id"`
	DB         interface{} `bson:"db" json:"db"`
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`
}

type User struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Email          string               `default:"" idx:"{email},unique" json:"email" binding:"required"`
	PhoneNumber    string               `bson:"phone" json:"phone"`
	FirstName      string               `default:"" json:"firstname"`
	LastName       string               `default:"" json:"lastname"`
	About          string               `default:"" json:"about"`
	LivesIn        []primitive.ObjectID `default:"" json:"lives_in"`
	Occupation     string               `default:"" json:"occupation"`
	Status         string               `default:"" json:"status"`
	Gender         string               `default:"" json:"gender"`
	Website        string               `default:"" json:"website"`
	DOB            string               `bson:"dob" json:"dob"`
	Addresses      string               `bson:"address" json:"address"`
	Educations     []Education          `bson:"educations" json:"educations"`
	Experiences    []Experience         `bson:"experiences" json:"experiences"`
	Skills         []Skill              `bson:"skills" json:"skills"`
	References     []Reference          `bson:"references" json:"references"`
	Awards         []Award              `bson:"awards" json:"awards"`
	Qualifications []Qualification      `bson:"qualifications" json:"qualifications"`
	Pages          []primitive.ObjectID `bson:"pages" json:"pages"`
	Courses        []primitive.ObjectID `bson:"courses" json:"courses"`
	DisplayName    string               `bson:"usernames" json:"usernames"`
	PhotoURL       string               `bson:"photo_url" json:"photo_url"`
	Created_at     time.Time            `json:"created_at"`
	Updated_at     time.Time            `json:"updated_at"`
	Score          []primitive.ObjectID `json:"score"`
	Penalty        []primitive.ObjectID `json:"penalty"`
}

type Account struct {
	UID           string `default:"" json:"uid" example:"Empty for create"`
	Email         string `default:"" idx:"{email},unique" json:"email" binding:"required" example:"me@thienhang.com"`
	EmailVerified bool   `bson:"verified" json:"verified" example:"false"`
	PhoneNumber   string `bson:"phone" json:"phone" example:"+840924202404"`
	Password      string `default:"" json:"password" binding:"required" example:"Abc@123456"`
}

type DeletingObject struct {
	IDs      []string `json:"ids" bson:"ids"`
	Property string   `json:"property" bson:"property"`
}
