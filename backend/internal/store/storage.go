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
		GetByID(context.Context, string) (*User, error)
		GetUserAddress(ctx context.Context, id string) (string, error)
		Update(ctx context.Context, user *User, user_id string) error
	}
	Products interface {
		GetAllProducts(ctx context.Context) ([]Product, error)
		GetBySort(ctx context.Context, token string, limit int) ([]Product, error)
		GetDetailProduct(ctx context.Context, id string) (*Product, error)
		CreateProduct(ctx context.Context, product *Product) (string, error)
		UpdateProduct(ctx context.Context, product *Product, id string) error
		DeleteProduct(ctx context.Context, id string) error
		GetProductByFilter(ctx context.Context, q string, tipe []string, flavor []string, start int, end int) ([]Product, error)
		GetProductImage(ctx context.Context, id string) ([]byte, error)
	}
	Transaction interface {
		GetUserCart(ctx context.Context, user_id string) ([]Cartproduct, error)
		AddProduct(ctx context.Context, newProduct *Cartproduct, user_id string) error
	}
}


func NewStorage(db *mongo.Client) Storage {
	return Storage{
	}
}
