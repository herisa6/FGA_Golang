package routes

import (
	"mygram/api/handlers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handlers.UserRegister)
		userRouter.POST("/login", handlers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), handlers.UpdateUsers)
		userRouter.DELETE("/:userId", middlewares.Authentication(), handlers.DeleteUsers)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", handlers.CreatePhoto)
		photoRouter.GET("/", handlers.ShowPhoto)
		photoRouter.PUT("/:photoId", handlers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", handlers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", handlers.CreateComment)
		commentRouter.GET("/", handlers.ShowComment)
		commentRouter.PUT("/:commentId", handlers.UpdateComment)
		commentRouter.DELETE("/:commentId", handlers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", handlers.CreateSocialMedia)
		socialMediaRouter.GET("/", handlers.ShowSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", handlers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", handlers.DeleteSocialMedia)
	}
	return r
}
