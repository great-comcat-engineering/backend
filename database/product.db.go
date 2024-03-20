package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"greatcomcatengineering.com/backend/models"
)

func GetAllProducts(ctx context.Context) ([]models.Product, error) {
	collection := Client.Database(DATABASE_NAME).Collection(COLLECTION_PRODUCTS)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	collection := Client.Database(DATABASE_NAME).Collection(COLLECTION_PRODUCTS)
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return models.Product{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return models.Product{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	hexID := insertedID.Hex()
	product.ID = hexID

	return product, nil
}

func CreateManyProducts(ctx context.Context, products []models.Product) ([]models.Product, error) {
	collection := Client.Database(DATABASE_NAME).Collection(COLLECTION_PRODUCTS)
	documents := make([]interface{}, len(products))
	for i, product := range products {
		documents[i] = product
	}

	_, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}

	return products, nil
}
