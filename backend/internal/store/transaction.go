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
	ID       string        `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   string        `bson:"user_id" json:"user_id,omitempty"`
	Date     string        `bson:"date" json:"date,omitempty"`
	Products []Cartproduct `bson:"products" json:"products,omitempty"`
	Address  string        `bson:"address" json:"address,omitempty"`
}

type TransactionStore struct {
	db *mongo.Client
}

func (t *TransactionStore) GetUserCart(ctx context.Context, user_id string) ([]Cartproduct, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{"user_id": user_id}
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
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	for _, product := range userCart {
		if product.ProductID == newProduct.ProductID {
			return fmt.Errorf("Product already exist")
		}
	}

	filter := bson.M{"user_id": user_id}
    update := bson.M{"$push": bson.M{"products": newProduct}}
    opts := options.Update().SetUpsert(true)

	_, err = t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionStore) DeleteProduct(ctx context.Context,user_id string, product_id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{"user_id": user_id, "products.product_id": product_id}
	update := bson.M{
        "$pull": bson.M{"products": bson.M{"product_id": product_id}, },
    }

	_, err := t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    return nil
}

func (t *TransactionStore) IncrementQuantity(ctx context.Context, user_id string, product_id string) error {
    ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
    defer cancel()

    filter := bson.M{
        "user_id": user_id,
        "products.product_id": product_id,
    }

    update := bson.M{
        "$inc": bson.M{
            "products.$.quantity": 1, 
        },
    }

    _, err := t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    return nil
}

func (t *TransactionStore) DecrementQuantity(ctx context.Context, user_id string, product_id string) error {
    ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
    defer cancel()

    filter := bson.M{
        "user_id": user_id,
        "products.product_id": product_id,
    }

    update := bson.M{
        "$inc": bson.M{
            "products.$.quantity": -1, 
        },
    }

    _, err := t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    return nil
}


