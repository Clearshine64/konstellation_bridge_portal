package model

import (
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BscTx struct {
	ID primitive.ObjectID `bson:"_id"`
	*types.Log
}

func NewBscTx() *BscTx {
	return &BscTx{
		primitive.NewObjectID(),
		&types.Log{},
	}
}
