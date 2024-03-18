package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"greatcomcatengineering.com/backend/models"
)

// AddUser adds a new user to the database
func AddUser(ctx context.Context, user models.User) error {
	client, err := GetClient()
	if err != nil {
		return err
	}

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_USERS)

	var existingUser models.User
	err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("a user with this email already exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = collection.InsertOne(ctx, user)
	return err
}

// GetUserByEmail retrieves a user from the database by id
func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	client, err := GetClient()
	if err != nil {
		return models.User{}, err
	}

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_USERS)
	var user models.User
	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

// UpdateUser updates a user in the database
func UpdateUser(ctx context.Context, user models.User) error {
	client, err := GetClient()
	if err != nil {
		return err
	}

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_USERS)
	_, err = collection.ReplaceOne(ctx, models.User{Email: user.Email}, user)
	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(ctx context.Context, email string) error {
	client, err := GetClient()
	if err != nil {
		return err
	}

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_USERS)
	_, err = collection.DeleteOne(ctx, models.User{Email: email})
	return err
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(ctx context.Context) ([]models.User, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(DATABASE_NAME).Collection(COLLECTION_USERS)
	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
