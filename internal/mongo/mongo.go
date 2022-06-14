package mongo

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/konstellation/swap/internal/config"
	"github.com/konstellation/swap/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	*mongo.Client
	DB  *mongo.Database
	Ctx context.Context
}

func InitConnection(ctx context.Context, c *config.DBConfig) (*Connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", c.UserName, c.Password, c.Host, c.Port)))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database(c.DBName)

	return &Connection{
		client,
		db,
		ctx,
	}, nil
}

func (c *Connection) InsertTx(tx *model.Tx) (interface{}, error) {
	txs := c.DB.Collection("txs")
	result, err := txs.InsertOne(c.Ctx, tx)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *Connection) GetTx(id primitive.ObjectID) (interface{}, error) {
	var tx model.Tx
	txs := c.DB.Collection("txs")
	err := txs.FindOne(c.Ctx, bson.M{"_id": id}).Decode(&tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Connection) FindTx(where map[string]string) (interface{}, error) {
	var tx model.Tx
	var filter bson.D
	txs := c.DB.Collection("txs")
	for condition, value := range where {
		if condition == "amount" {
			val, _ := strconv.ParseFloat(value, 64)
			if val == math.Trunc(val) {
				filter = append(filter, bson.E{Key: condition, Value: uint64(val)})
			} else {
				filter = append(filter, bson.E{Key: condition, Value: val})
			}

			continue
		}
		if condition == "source_network_completed" {
			val, _ := strconv.ParseBool(value)
			filter = append(filter, bson.E{Key: condition, Value: val})
			continue
		}
		if condition == "destination_network_completed" {
			val, _ := strconv.ParseBool(value)
			filter = append(filter, bson.E{Key: condition, Value: val})
			continue
		}
		if condition == "completed" {
			val, _ := strconv.ParseBool(value)
			filter = append(filter, bson.E{Key: condition, Value: val})
			continue
		}
		filter = append(filter, bson.E{Key: condition, Value: value})
	}
	err := txs.FindOne(c.Ctx, filter).Decode(&tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Connection) UpdateTx(tx *model.Tx) (interface{}, error) {
	// Reference: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.UpdateOne
	txs := c.DB.Collection("txs")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "_id", Value: tx.ID}}
	result, err := txs.UpdateOne(c.Ctx, filter, bson.D{primitive.E{Key: "$set", Value: tx}}, opts)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

func (c *Connection) InsertBlacklist(address *model.Blacklist) (interface{}, error) {
	blacklist := c.DB.Collection("blacklist")
	result, err := blacklist.InsertOne(c.Ctx, address)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *Connection) FindBlacklist(where map[string]string) (interface{}, error) {
	var address model.Blacklist
	var filter bson.D
	blacklist := c.DB.Collection("blacklist")
	for condition, value := range where {
		filter = append(filter, bson.E{Key: condition, Value: value})
	}
	err := blacklist.FindOne(c.Ctx, filter).Decode(&address)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (c *Connection) InsertBscTx(bscTx *types.Log) (interface{}, error) {
	blacklist := c.DB.Collection("bsctxs")
	result, err := blacklist.InsertOne(c.Ctx, bscTx)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *Connection) FindBscTx(where map[string]string) (*mongo.Cursor, error) {
	var filter bson.D
	blacklist := c.DB.Collection("bsctxs")
	for condition, value := range where {
		if condition == "removed" {
			val, _ := strconv.ParseBool(value)
			filter = append(filter, bson.E{Key: condition, Value: val})
			continue
		}
		filter = append(filter, bson.E{Key: condition, Value: value})
	}
	cur, err := blacklist.Find(c.Ctx, filter)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

func (c *Connection) UpdateBscTx(bscTx *types.Log) (interface{}, error) {
	// Reference: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.UpdateOne
	bsctxs := c.DB.Collection("bsctxs")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "blocknumber", Value: bscTx.BlockNumber}}
	result, err := bsctxs.UpdateOne(c.Ctx, filter, bson.D{primitive.E{Key: "$set", Value: bscTx}}, opts)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}
