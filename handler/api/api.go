package api

import (
	"fmt"
	"handler/database"
	"handler/types"
)





type ApiHandler struct {
	dbstore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbstore: dbStore,
	}
}

// This is where we will implement the API handlers
func (api ApiHandler) RegisterUserHandler(user types.RegisterUser) error {
	// Check if the user exists

	if user.Username == "" || user.Password == "" {
		// return an error
		return fmt.Errorf("request is missing required fields")
	}

	// Does a user with this username already exist?
	userExists, err := api.dbstore.UserExists(user.Username)

	if err != nil {
		return fmt.Errorf("error checking if user exists: %v", err)
	}

	if userExists {
		return fmt.Errorf("username %s already exists", user.Username)
	}

	// A user with this username does not exist, insert the user
	err = api.dbstore.InsertUser(user)

	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}

	return nil
}