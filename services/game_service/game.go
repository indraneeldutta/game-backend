package gameservice

import (
	"errors"
	"time"

	"github.com/game-backend/common"
	"github.com/game-backend/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameService struct {
	collection *mongo.Collection
}

func NewGameService(client *mongo.Client) GameService {
	return GameService{
		collection: client.Database(viper.GetString("MONGO.DB_NAME")).Collection("game_state"),
	}
}

func (s *GameService) SaveGameState(ctx *common.Context, request models.GameState) error {
	request.CreatedAt = time.Now()
	_, insertErr := s.collection.InsertOne(ctx.Ctx, request)
	if !errors.Is(insertErr, nil) {
		ctx.Logger.Error(insertErr)
		return insertErr
	}

	return nil
}

func (s *GameService) GetGameState(ctx *common.Context, userID string) (models.GameState, error) {
	var gameState models.GameState
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"created_at", -1}})
	err := s.collection.FindOne(ctx.Ctx, bson.M{"user_id": userID}, findOptions).Decode(&gameState)

	if !errors.Is(err, nil) {
		ctx.Logger.Error(err)
		return gameState, err
	}

	return gameState, nil
}
