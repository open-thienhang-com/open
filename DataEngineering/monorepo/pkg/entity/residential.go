package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	Building primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	Block    string             `bson:"block" json:"block"`
	Floor    string             `bson:"floor" json:"floor"`
	NoHouse  string             `bson:"no_house" json:"website"`
}

type Building struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	Name    string             `bson:"name" json:"name"`
	Icon    string             `bson:"icon" json:"icon"`
	Address string             `bson:"address" json:"address"` // becase can be "present"
	// Ward        primitive.ObjectID `bson:"ward" json:"ward"`
	// Province    primitive.ObjectID `bson:"province" json:"province"`
	// District    primitive.ObjectID `bson:"district" json:"district"`
	Description string    `bson:"description" json:"desctiption"`
	Website     string    `bson:"website" json:"website"`
	Size        int       `bson:"size" json:"size"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Addresses   []Address `bson:"addresses" json:"addresses"`
}

type Street struct {
}

type Ward struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	Name     string             `bson:"name" json:"name"`
	Pre      string             `bson:"pre" json:"pre"`
	Building []Building         `bson:"building" json:"building"`
}

type District struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	Name  string             `bson:"name" json:"name"`
	Pre   string             `bson:"pre" json:"pre"`
	Wards []Ward             `bson:"ward" json:"ward"`
}

type Province struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" binding:"required"`
	Code      string             `bson:"code" json:"code"`
	Name      string             `bson:"name" json:"name"`
	Districts []District         `bson:"district" json:"district"`
}
