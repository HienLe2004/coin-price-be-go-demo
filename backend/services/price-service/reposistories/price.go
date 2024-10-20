package reposistories

import (
	"context"
	"fmt"

	"github.com/HienLe2004/coin-price-be-go-demo/services/price-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PriceRepo struct {
	MongoCollection *mongo.Collection
}

func (r *PriceRepo) InsertPrice(price *models.Price) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), price)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PriceRepo) FindPriceBySymbol(symbol string) (*models.Price, error) {
	var price models.Price
	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "symbol", Value: symbol}}).Decode(&price)
	if err != nil {
		return nil, err
	}
	return &price, nil
}

func (r *PriceRepo) FindAllPrice() ([]models.Price, error) {
	result, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var prices []models.Price
	err = result.All(context.Background(), &prices)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}
	return prices, nil
}

func (r *PriceRepo) UpdatePriceBySymbol(symbol string, updatedPrice *models.Price) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "symbol", Value: symbol}},
		bson.D{{Key: "$set", Value: updatedPrice}})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (r *PriceRepo) DeletePriceBySymbol(symbol string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "symbol", Value: symbol}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *PriceRepo) DeleteAllPrice() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
