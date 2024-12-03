package store

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password []byte             `bson:"password" json:"password"`
	Phone    string             `bson:"phone" json:"phone"`
	Role     string             `bson:"role" json:"role"`
	Addrres  UserAddress        `bson:"address" json:"address"`
}

type UserAddress struct {
	Label     string `bson:"label" json:"label"`
	Recipient string `bson:"recipient" json:"recipient"`
	Phone     string `bson:"phone" json:"phone"`
	Message   string `bson:"message" json:"message"`
	Street    string `bson:"street" json:"street"`
	City      string `bson:"city" json:"city"`
	Province  string `bson:"province" json:"province"`
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
