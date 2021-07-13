package mongoservice

import (
	"context"
	"fmt"
	"time"

	"github.com/game-backend/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoConnection connects to a mongo instance
func NewMongoConnection() *mongo.Client {
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=admin&readPreference=primary&ssl=%v", viper.GetString("MONGO.USER"),
		viper.GetString("MONGO.PASS"),
		viper.GetString("MONGO.HOST"),
		viper.GetString("MONGO.PORT"),
		viper.GetString("MONGO.SSL"))
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		logger.Log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("Connected to database")

	// defer client.Disconnect(ctx)

	return client
}
