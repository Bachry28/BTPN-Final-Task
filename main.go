package main

import (
	"github.com/Bachry28/BTPN-Final-Task/controllers/photocontroller"

	"github.com/Bachry28/BTPN-Final-Task/controllers/usercontroller"

	"github.com/Bachry28/BTPN-Final-Task/model"

	"github.com/Bachry28/BTPN-Final-Task/tools"

	"github.com/gin-gonic/gin"
)

func init() {
	tools.LoadEnv()
	model.ConnectDB()

}
func main() {
	r := gin.Default()

	r.GET("/api/users", usercontroller.GetAllUser)
	r.GET("/api/users/:id", usercontroller.GetUserById)
	r.POST("/api/users/register", usercontroller.Register)
	r.POST("/api/users/login", usercontroller.Login)
	r.PUT("/api/users/:id", usercontroller.UpdateUser)
	r.DELETE("/api/users/:id", usercontroller.DeleteUser)

	r.GET("/api/photo", tools.Auth, photocontroller.GetAllPhoto)
	r.GET("/api/photo/:id", tools.Auth, photocontroller.GetPhotoById)
	r.POST("/api/photo", tools.Auth, photocontroller.CreatePhoto)
	r.PUT("/api/photo/:id", tools.Auth, photocontroller.UpdatePhoto)
	r.DELETE("/api/photo/:id", tools.Auth, photocontroller.DeletePhoto)

	r.Run()
}
