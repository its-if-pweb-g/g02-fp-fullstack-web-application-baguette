package store

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	ID       string
	UserID   string
	Date     time.Time
	Status   string
	Products []Cart
	Address  string
}

type Cart struct {
	ProductID string `bson:"_id,omitempty"`
	Price     int    `bson:"price,omitempty"`
	Quantity  int    `bson:"quantity,omitempy"`
}

type TransactionStore struct {
	db *mongo.Client
}
