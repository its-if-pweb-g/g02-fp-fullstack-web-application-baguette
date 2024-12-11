package store

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID                string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string    `bson:"name,omitempty" json:"name"`
	HeaderDescription string    `bson:"header_description,omitempty" json:"header_description"`
	Description       string    `bson:"description,omitempty" json:"description,omitempty"`
	Price             int       `bson:"price,omitempty" json:"price,omitempty"`
	Stock             int       `bson:"stock,omitempty" json:"stock,omitempty"`
	Sold              int       `bson:"sold,omitempty" json:"sold,omitempty"`
	Image             []byte    `bson:"image,omitempty" json:"image,omitempty"`
	CreatedAt         time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Type              []string  `bson:"type,omitempty json:type,omitempty"`
	Flavor            []string  `bson:"flavor,omitempty" json:"flavor,omitempty"`
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

	filter := bson.M{"_id": objectID}
	projection := bson.M{"image":0}
	opts := options.FindOne().SetProjection(projection)

	var result Product
	if err := p.db.Database(Database).Collection(ProductCollection).FindOne(ctx, filter, opts).Decode(&result); err != nil {
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

	filter := bson.M{"_id": objectID}
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

	filter := bson.M{"_id": objectID}
	if _, err := p.db.Database(Database).Collection(ProductCollection).DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (p *ProductStore) GetProductByFilter(ctx context.Context, q string, tipe []string, flavor []string, start int, end int) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{}
	opts := options.Find().SetProjection(bson.M{"image": 0})

	if q != "" {
		filter["$or"] = []bson.M{
			{"name": bson.M{"$regex": q, "$options": "i"}},
			{"description": bson.M{"$regex": q, "$options": "i"}},
			{"header_description": bson.M{"$regex": q, "$options": "i"}},
		}
	}

	if len(tipe) > 0 {
		filter["type"] = bson.M{"$in": tipe}
	}

	if len(flavor) > 0 {
		filter["flavor"] = bson.M{"$in": flavor}
	}

	if start > 0 && end > start {
		filter["price"] = bson.M{"$gte": start, "$lte": end}
	}

	cursor, err := p.db.Database(Database).Collection(ProductCollection).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []Product
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, err
}

func (p *ProductStore) GetProductImage(ctx context.Context, id string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	opts := options.FindOne().SetProjection(bson.M{"image": 1})

	var result bson.M
	err = p.db.Database(Database).Collection(ProductCollection).FindOne(ctx, bson.M{"_id": objectID}, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	imageData, ok := result["image"].(primitive.Binary)
	if !ok {
		return nil, errors.New("image field not found or is of incorrect type")
	}

	return imageData.Data, nil
}
