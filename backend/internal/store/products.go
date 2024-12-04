package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

var ()

type Product struct {
	ID                string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string    `bson:"name,omitempty" json:"name"`
	HeaderDescription string    `bson:"header_description,omitempty" json:"header_description"`
	Description       string    `bson:"description,omitempty" json:"description"`
	Price             int       `bson:"price,omitempty" json:"price"`
	Stock             int       `bson:"stock,omitempty" json:"stock"`
	Sold              int       `bson:"sold,omitempty" json:"sold"`
	Image             []byte    `bson:"image,omitempty" json:"image"`
	CreatedAt         time.Time `bson:"created_at,omitempty" json:"created_at"` 
	Category          []string  `bson:"category,omitempty" json:"category"`
}

type ProductStore struct {
	db *mongo.Client
}

func (p *ProductStore) GetAllProducts(ctx context.Context) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	projection := bson.M{"image": 0}
	opts := options.Find().SetProjection(projection)

	cursor, err := p.db.Database(Database).Collection(ProductCollection).Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []Product
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *ProductStore) GetBySort(ctx context.Context, token string, limit int) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	projection := bson.M{"image": 0}
	opts := options.Find().
		SetProjection(projection).
		SetSort(bson.D{{Key: token, Value: -1}}).
		SetLimit(int64(limit))

	cursor, err := p.db.Database(Database).Collection(ProductCollection).Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var result []Product
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}	

	return result, nil
}

func (p *ProductStore) GetDetailProduct(ctx context.Context, id string) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id" : objectID}

	var result Product
	if err := p.db.Database(Database).Collection(ProductCollection).FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *ProductStore) CreateProduct(ctx context.Context, product *Product) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	result, err := p.db.Database(Database).Collection(ProductCollection).InsertOne(ctx, product)
	if err != nil {
		return "", err 
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.Hex(), nil
}

func (p *ProductStore) UpdateProduct(ctx context.Context, product *Product, id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id" : objectID}
	update := bson.M{"$set": product}

	if _, err := p.db.Database(Database).Collection(ProductCollection).UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (p *ProductStore) DeleteProduct(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id" : objectID}
	if _, err := p.db.Database(Database).Collection(ProductCollection).DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}


