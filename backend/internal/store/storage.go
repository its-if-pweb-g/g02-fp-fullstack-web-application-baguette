package store

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrorNotFound	= errors.New("resource not found")
	ErrorConflict	= errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
	Database = "pweb-g"
	UserCollection = "users"
	ProductCollection = "products"
	TransactionCollection = "transaction"
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) (string, error)
		GetByEmail(context.Context, string) (*User, error)
	}
	Products interface {

	}
}


func NewStorage(db *mongo.Client) Storage {
	return Storage{
	}
}
