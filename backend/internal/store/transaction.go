package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cartproduct struct {
	ProductID string `bson:"product_id,omitempty" json:"id"`
	Name      string `bson:"name" json:"name"`
	Price     int    `bson:"price,omitempty" json:"price"`
	Quantity  int    `bson:"quantity,omitempty" json:"quantity"`
}

type Order struct {
	ID       string        `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   string        `bson:"user_id" json:"user_id,omitempty"`
	Date     time.Time     `bson:"date" json:"date,omitempty"`
	Status   string        `bson:"status" json:"status"`
	Products []Cartproduct `bson:"products" json:"products,omitempty"`
	Address  string        `bson:"address" json:"address,omitempty"`
}

type TransactionStore struct {
	db *mongo.Client
}

type UserCart struct {
	ID       primitive.ObjectID `bson:"_id"`
	Products []Cartproduct      `bson:"products"`
}

func (t *TransactionStore) GetUserCart(ctx context.Context, user_id string) (string, []Cartproduct, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{"user_id": user_id, "status": "cart"}
	opts := options.FindOne().SetProjection(bson.M{"_id": 1, "products": 1})

	var result UserCart
	if err := t.db.Database(Database).Collection(TransactionCollection).FindOne(ctx, filter, opts).Decode(&result); err != nil {
		return "", nil, err
	}

	return result.ID.Hex(), result.Products, nil
}

func (t *TransactionStore) CreateUserOrder(ctx context.Context, order Order) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	result, err := t.db.Database(Database).Collection(TransactionCollection).InsertOne(ctx, order)
	if err != nil {
		return "", err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)

	return objectID.Hex(), nil
}

func (t *TransactionStore) AddProduct(ctx context.Context, newProduct Cartproduct, user_id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, userCart, err := t.GetUserCart(ctx, user_id)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if err == mongo.ErrNoDocuments {
		newOrder := Order{
			UserID:   user_id,
			Date:     time.Now(),
			Status:   "cart",
			Products: []Cartproduct{newProduct},
			Address:  "",
		}

		if _, err := t.CreateUserOrder(ctx, newOrder); err != nil {
			return err
		}

		return nil
	} else {
		for _, product := range userCart {
			if product.ProductID == newProduct.ProductID {
				err := t.IncrementQuantity(ctx, user_id, product.ProductID, newProduct.Quantity)

				if err != nil {
					return err
				}

				return nil
			}
		}
	}

	filter := bson.M{"user_id": user_id, "status": "cart"}
	update := bson.M{"$push": bson.M{"products": newProduct}}

	_, err = t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionStore) MarkTransaction(ctx context.Context, transaction_id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(transaction_id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"status": "deleted"}}

	if _, err := t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (t *TransactionStore) DeleteProduct(ctx context.Context, user_id string, product_id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{
		"user_id":             user_id,
		"status":              "cart",
		"products.product_id": product_id,
	}

	update := bson.M{
		"$pull": bson.M{"products": bson.M{"product_id": product_id}},
	}

	_, err := t.db.Database(Database).Collection(TransactionCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionStore) IncrementQuantity(ctx context.Context, user_id string, product_id string, q int) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{
		"user_id":             user_id,
		"status":              "cart",
		"products.product_id": product_id,
	}

	update := bson.M{
		"$inc": bson.M{
			"products.$.quantity": q,
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
		"user_id":             user_id,
		"status":              "cart",
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
