package userservice

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(client *mongo.Client) UserService {
	return UserService{
		collection: client.Database(viper.GetString("MONGO.DB_NAME")).Collection("users"),
	}
}
