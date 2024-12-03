package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (

)

type Product struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Slug        string    `bson:"slug"`
	Description string    `bson:"description"`
	Price       int       `bson:"price"`
	Stock       int       `bson:"stock"`
	ExpiredIn   time.Time `bson:"expired_in"`
	Sold        int       `bson:"sold"`
	MinPurchase int       `bson:"min_purchase"`
	Image       []byte    `bson:"image"`
	Category    []string  `bson:"category"`
}

type ProductStore struct {
	db *mongo.Client
}

func (p *ProductStore) GetData(ctx context.Context) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{}
	opts := options.Find().SetSort(bson.D{{Key : "sold", Value : -1}})

	cursor, err := p.db.Database(Database).Collection(ProductCollection).Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	
}

