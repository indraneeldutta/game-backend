package apis

import (
	"errors"
	"net/http"

	"github.com/game-backend/common"
	"github.com/game-backend/logger"
	"github.com/game-backend/models"
	gameservice "github.com/game-backend/services/game_service"
	"github.com/gin-gonic/gin"
)

const (
	GAMESTATE = "/user/:userID/state"
)

type GameStateController struct {
	service *gameservice.GameService
}

func NewGameStateController(router *gin.RouterGroup, service *gameservice.GameService) {
	controller := GameStateController{
		service: service,
	}

	router.PUT(GAMESTATE, controller.saveGameState)
	router.GET(GAMESTATE, controller.getGameState)
}

func (c *GameStateController) saveGameState(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", GAMESTATE)

	var request models.GameState
	if errors.Is(ginCtx.ShouldBindJSON(&request), nil) {
		request.UserID = ginCtx.Param("userID")
		err := c.service.SaveGameState(context, request)
		if !errors.Is(err, nil) {
			ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
				"date": "",
			})
			context.Logger.Error(err)
			return
		}

		ginCtx.SecureJSON(http.StatusNoContent, gin.H{})
		return
	}
	ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
		"date": "",
	})
	return
}

func (c *GameStateController) getGameState(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", GAMESTATE)

	userID := ginCtx.Param("userID")
	response, err := c.service.GetGameState(context, userID)
	if !errors.Is(err, nil) {
		ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
			"date": "",
		})
		context.Logger.Error(err)
		return
	}

	ginCtx.SecureJSON(http.StatusOK, gin.H{
		"data": response,
	})
}
