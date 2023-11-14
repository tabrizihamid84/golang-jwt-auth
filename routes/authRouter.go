package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tabrizihamid84/jwt_auth-golang/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())

}
