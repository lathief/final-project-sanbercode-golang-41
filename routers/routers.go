package routers

import (
	"bioskop/controllers"
	"bioskop/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/films", controllers.GetAllFilm)
	router.GET("/films/:id", controllers.GetFilmById)
	router.GET("/films/search", controllers.GetFilmByNameStatus)

	router.GET("/schedulelist", controllers.GetAllSchedule)
	router.GET("/schedule/:id", controllers.GetScheduleById)

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
	}))
	authorized.GET("/user", controllers.GetAllUserCustomer)

	authorized.POST("/films", controllers.InsertFilm)
	authorized.PUT("/films/:id", controllers.UpdateFilm)
	authorized.DELETE("/films/:id", controllers.DeleteFilm)

	authorized.POST("/schedule", controllers.InsertSchedule)
	authorized.PUT("/schedule/:id", controllers.UpdateSchedule)
	authorized.DELETE("/schedule/:id", controllers.DeleteSchedule)

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/user/order", controllers.OrderTicket)
	protectedRoutes.POST("/user/update", controllers.UpdateProfile)
	protectedRoutes.GET("/user/showticket", controllers.ShowTicket)

	// protectedRoutes.POST("/films", controllers.InsertFilm)
	// protectedRoutes.PUT("/films/:id", controllers.UpdateFilm)
	// protectedRoutes.DELETE("/films/:id", controllers.DeleteFilm)

	// protectedRoutes.POST("/schedule", controllers.InsertSchedule)
	// protectedRoutes.PUT("/schedule/:id", controllers.UpdateSchedule)
	// protectedRoutes.DELETE("/schedule/:id", controllers.DeleteSchedule)

	return router
}
