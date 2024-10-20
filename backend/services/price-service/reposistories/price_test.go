package reposistories

import (
	"context"
	"log"
	"testing"

	"github.com/HienLe2004/coin-price-be-go-demo/services/price-service/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://hienlengoc2004:hienlemgoc2004@cluster0.ods9u.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal("error while connecting mongodb ", err)
	}
	log.Println("mongodb successfully connected")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("ping success")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	collection := mongoTestClient.Database("coin-price-system").Collection("price-test")

	priceRepo := PriceRepo{MongoCollection: collection}

	t.Run("Insert price number one", func(t *testing.T) {
		price := models.Price{
			Price:  3.23,
			Name:   "bitcoin",
			Symbol: "btc",
		}
		result, err := priceRepo.InsertPrice(&price)
		if err != nil {
			t.Fatal("insert 1 failed", err)
		}
		t.Log("Insert 1 successfull", result)
	})

	t.Run("Get price number one", func(t *testing.T) {
		result, err := priceRepo.FindPriceBySymbol("btc")
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("Get price number one", result.Name)
	})

	t.Run("Get all prices", func(t *testing.T) {
		result, err := priceRepo.FindAllPrice()
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("Get all price", result)
	})

	t.Run("Update price number one name", func(t *testing.T) {
		price := models.Price{
			Name:   "bigcoin",
			Price:  3.33,
			Symbol: "btc",
		}
		result, err := priceRepo.UpdatePriceBySymbol("btc", &price)
		if err != nil {
			t.Fatal("Update price number one failed", err)
		}
		t.Log("Update count", result)
	})

	t.Run("Get price number one again", func(t *testing.T) {
		result, err := priceRepo.FindPriceBySymbol("btc")
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("Get price number one", result.Name)
	})

	t.Run("Delete price number one", func(t *testing.T) {
		result, err := priceRepo.DeletePriceBySymbol("btc")
		if err != nil {
			t.Fatal("delete operation failed", err)
		}
		t.Log("Get price number one", result)
	})
}
