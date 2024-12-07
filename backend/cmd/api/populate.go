package main

import (
	"api/internal/store"
	"context"

)


func addAdmin(currentStore store.Storage) {
	ctx := context.Background()

	admin := &store.User{
		Name: "deva",
		Email: "deva@gmail.com",
		Role: "admin",
		Phone: "",
		Addrres: "",
	}

	if err := admin.SetPassword("password"); err != nil {
		return 
	}

	_, err := currentStore.Users.Create(ctx, admin)
	if err != nil {
		return 
	}
}