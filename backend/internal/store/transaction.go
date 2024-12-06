package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cartproduct struct {
	ProductID string `bson:"product_id,omitempty" json:"product_id"`
	Name      string `bson:"name" json:"name"`
	Price     int    `bson:"price,omitempty" json:"price"`
	Quantity  int    `bson:"quantity,omitempty" json:"quantity"`
}

type Order struct {
	ID       string        `bson:"_id,omitempty"`
	UserID   string        `bson:"user_id"`
	Date     string        `bson:"date"`
	Status   string        `bson:"status"`
	Products []Cartproduct `bson:"products"`
	Address  string        `bson:"address"`
}

type TransactionStore struct {
	db *mongo.Client
}

func (t *TransactionStore) GetUserCart(ctx context.Context, user_id string) ([]Cartproduct, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{"user_id": user_id, "status": "cart"}
	opts := options.FindOne().SetProjection(bson.M{"_id": 1, "products": 1})

	var result Order
	if err := t.db.Database(Database).Collection(TransactionCollection).FindOne(ctx, filter, opts).Decode(&result); err != nil {
		return nil, err
	}
	
	return result.Products, nil
}

func (t *TransactionStore) AddProduct(ctx context.Context, newProduct *Cartproduct, user_id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	userCart, err := t.GetUserCart(ctx, user_id)
	if err != nil {
		return err
	}

	for _, product := range userCart {
		if product.ProductID == newProduct.ProductID {
			return fmt.Errorf("Product already exist")
		}
	}

	filter := bson.M{"user_id": user_id, "status": "cart"}
    update := bson.M{"$push": bson.M{"products": newProduct}}
    opts := options.Update().SetUpsert(true)

	_, err = t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
