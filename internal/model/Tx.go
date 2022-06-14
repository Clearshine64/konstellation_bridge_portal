package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tx struct {
	ID                          primitive.ObjectID `bson:"_id"`
	FromAddress                 string             `json:"from_address" bson:"from_address"`
	ToAddress                   string             `json:"to_address" bson:"to_address"`
	SourceNetwork               string             `json:"source_network" bson:"source_network"`
	SourceNetworkHash           string             `json:"source_network_hash" bson:"source_network_hash"`
	DestinationNetwork          string             `json:"destination_network" bson:"destination_network"`
	DestinationNetworkHash      string             `json:"destination_network_hash" bson:"destination_network_hash"`
	Amount                      float64            `json:"amount" bson:"amount"`
	Timestamp                   uint64             `json:"timestamp" bson:"timestamp"`
	SourceNetworkCompleted      bool               `json:"source_network_completed" bson:"source_network_completed"`
	DestinationNetworkCompleted bool               `json:"destination_network_completed" bson:"destination_network_completed"`
	Completed                   bool               `json:"completed" bson:"completed"`
	CreatedAt                   string             `json:"created_at" bson:"created_at"`
	UpdatedAt                   string             `json:"updated_at" bson:"updated_at"`
	TxTryCount                  int                `json:"tx_try_count" bson:"tx_try_count"`
}

func NewTx() *Tx {
	return &Tx{
		ID: primitive.NewObjectID(),
	}
}
