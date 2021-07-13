package apis

import (
	"errors"
	"net/http"

	"github.com/game-backend/common"
	"github.com/game-backend/logger"
	"github.com/game-backend/models"
	userservice "github.com/game-backend/services/user_service"
	"github.com/gin-gonic/gin"
)

const (
	USER    = "/user"
	FRIENDS = "/user/:userID/friends"
)

type UserController struct {
	service *userservice.UserService
}

// NewUserController initialises all the routes for user functions
func NewUserController(router *gin.RouterGroup, service *userservice.UserService) {
	controller := UserController{
		service: service,
	}

	router.POST(USER, controller.createUser)
	router.PUT(FRIENDS, controller.updateFriends)
	router.GET(FRIENDS, controller.getFriends)
	router.GET(USER, controller.getAllUsers)

}

func (c *UserController) createUser(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", USER)

	var request models.CreateUser
	if errors.Is(ginCtx.ShouldBindJSON(&request), nil) {
		resp, err := c.service.CreateUser(context, request)
		if !errors.Is(err, nil) {
			ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
				"data": "",
			})
			context.Logger.Error(err)
			return
		}

		ginCtx.SecureJSON(http.StatusOK, gin.H{
			"data": resp,
		})
		return
	}
	ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
		"data": "",
	})
}

func (c *UserController) updateFriends(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", FRIENDS)

	var request models.UpdateFriendsRequest
	if errors.Is(ginCtx.ShouldBindJSON(&request), nil) {
		request.UserID = ginCtx.Param("userID")
		err := c.service.UpdateFriends(context, request)
		if !errors.Is(err, nil) {
			ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{})
			context.Logger.Error(err)
			return
		}

		ginCtx.SecureJSON(http.StatusNoContent, gin.H{})
		return
	}
	ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{})
}

func (c *UserController) getFriends(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", FRIENDS)

	userID := ginCtx.Param("userID")
	resp, err := c.service.GetFriends(context, userID)

	if !errors.Is(err, nil) {
		ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
			"date": "",
		})
		context.Logger.Error(err)
		return
	}

	ginCtx.SecureJSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (c *UserController) getAllUsers(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", USER)

	resp, err := c.service.GetAllUsers(context)

	if !errors.Is(err, nil) {
		ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{
			"date": "",
		})
		context.Logger.Error(err)
		return
	}

	ginCtx.SecureJSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
