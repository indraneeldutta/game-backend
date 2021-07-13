package userservice

import (
	"errors"
	"time"

	"github.com/game-backend/common"
	"github.com/game-backend/models"
	uuid "github.com/satori/go.uuid"
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

func (s *UserService) CreateUser(ctx *common.Context, request models.CreateUser) (*models.CreateUserResponse, error) {
	request.ID = uuid.Must(uuid.NewV1(), nil).String()
	request.CreatedAt, request.UpdatedAt = time.Now(), time.Now()
	_, insertErr := s.collection.InsertOne(ctx.Ctx, request)

	if !errors.Is(insertErr, nil) {
		ctx.Logger.Error(insertErr)
		return nil, insertErr
	}

	resp := models.CreateUserResponse{
		ID:   request.ID,
		Name: request.Name,
	}

	return &resp, nil
}
