package userservice

import (
	"errors"
	"time"

	"github.com/game-backend/common"
	"github.com/game-backend/models"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	db *mongo.Database
}

// NewUserService initialises user service
func NewUserService(client *mongo.Client) UserService {
	return UserService{
		db: client.Database(viper.GetString("MONGO.DB_NAME")),
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx *common.Context, request models.CreateUser) (*models.CreateUserResponse, error) {
	request.ID = uuid.Must(uuid.NewV1(), nil).String()
	request.CreatedAt, request.UpdatedAt = time.Now(), time.Now()
	_, insertErr := s.db.Collection("users").InsertOne(ctx.Ctx, request)

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

// UpdateFriends updates the friends list for the user
func (s *UserService) UpdateFriends(ctx *common.Context, request models.UpdateFriendsRequest) error {
	filter := bson.M{"_id": bson.M{"$eq": request.UserID}}
	update := bson.M{
		"$set": bson.M{
			"friends":    request.Friends,
			"updated_at": time.Now(),
		},
	}
	_, err := s.db.Collection("users").UpdateOne(
		ctx.Ctx,
		filter,
		update,
	)

	if !errors.Is(err, nil) {
		ctx.Logger.Error(err)
		return err
	}

	return nil
}

// GetFriends retrieves all the friends and their details of the user
func (s *UserService) GetFriends(ctx *common.Context, userID string) ([]models.FriendsData, error) {
	var friendsData []models.FriendsData
	var friends models.UpdateFriendsRequest
	err := s.db.Collection("users").FindOne(ctx.Ctx, bson.M{"_id": userID}).Decode(&friends)

	if !errors.Is(err, nil) {
		ctx.Logger.Error(err)
		return friendsData, err
	}

	data, err := s.db.Collection("users").Find(ctx.Ctx, bson.M{"_id": bson.M{"$in": friends.Friends}})
	if !errors.Is(err, nil) {
		ctx.Logger.Error(err)
		return friendsData, err
	}

	if err = data.All(ctx.Ctx, &friendsData); err != nil {
		ctx.Logger.Error(err)
		return friendsData, err
	}

	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"created_at", -1}})
	for i, value := range friendsData {
		var gameState models.GameState
		err := s.db.Collection("game_state").FindOne(ctx.Ctx, bson.M{"user_id": value.ID}, findOptions).Decode(&gameState)
		if errors.Is(err, nil) {
			friendsData[i].HighScore = gameState.Score
		}
	}
	return friendsData, nil
}

// GetAllUsers retrieves all the users
func (s *UserService) GetAllUsers(ctx *common.Context) ([]models.CreateUser, error) {
	var users []models.CreateUser
	data, err := s.db.Collection("users").Find(ctx.Ctx, bson.D{})

	if !errors.Is(err, nil) {
		ctx.Logger.Error(err)
		return users, err
	}

	if err = data.All(ctx.Ctx, &users); err != nil {
		ctx.Logger.Error(err)
		return users, err
	}

	return users, nil
}
