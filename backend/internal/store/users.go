package store

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password []byte `bson:"password" json:"password"`
	Phone    string `bson:"phone" json:"phone"`
	Role     string `bson:"role" json:"role"`
	Addrres  string `bson:"address" json:"address"`
}

type UserStore struct {
	db *mongo.Client
}

func (u *User) SetPassword(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = hash
	return nil
}

func (u *User) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(text))
}

func (s *UserStore) Create(ctx context.Context, user *User) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if _, err := s.GetByEmail(ctx, user.Email); err == nil {
		return "", ErrDuplicateEmail
	}

	result, err := s.db.Database(Database).Collection(UserCollection).InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)

	return objectID.Hex(), nil
}

func (s *UserStore) Update(ctx context.Context, user *User, user_id string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	updateData := bson.M{"$set": user}

	if _, err := s.db.Database(Database).Collection(UserCollection).UpdateOne(ctx, filter, updateData); err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	filter := bson.M{"email": email}

	var result User

	err := s.db.Database(Database).Collection(UserCollection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *UserStore) GetByID(ctx context.Context, id string) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var result User

	if err := s.db.Database(Database).Collection(UserCollection).FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *UserStore) GetUserAddress(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": objectID}
	projection := options.FindOne().SetProjection(bson.M{"address": 1, "_id": 0})

	var result struct {
		Address string `bson:"address"`
	}

	err = s.db.Database(Database).Collection(UserCollection).FindOne(ctx, filter, projection).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Address, nil
}
