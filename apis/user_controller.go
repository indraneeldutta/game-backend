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
	CREATEUSER = "/user"
)

type UserController struct {
	service *userservice.UserService
}

func NewUserController(router *gin.RouterGroup, service *userservice.UserService) {
	controller := UserController{
		service: service,
	}

	router.POST(CREATEUSER, controller.createUser)
}

func (c *UserController) createUser(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, logger.Log)
	context.Logger.Infof("request received for %v", CREATEUSER)

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
