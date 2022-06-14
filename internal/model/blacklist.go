package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blacklist struct {
	ID        primitive.ObjectID `bson:"_id"`
	Address   string             `json:"address" bson:"address"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewBlacklist() *Blacklist {
	return &Blacklist{
		ID: primitive.NewObjectID(),
	}
}
