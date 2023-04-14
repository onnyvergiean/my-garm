package router

import (
	"my-garm/controllers"
	"my-garm/middlewares"

	"github.com/gin-gonic/gin"
)


func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.GET("/:photoId", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/:photoId", controllers.CreateComment)
		commentRouter.GET("/:photoId", controllers.GetComments)
		commentRouter.GET("/:photoId/:commentId", controllers.GetComment)
		commentRouter.PUT("/:photoId/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:photoId/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}


	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/",controllers.CreateSocialmedia)
		socialMediaRouter.GET("/",controllers.GetSocialMedias)
		socialMediaRouter.GET("/:socialMediaId",controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId",middlewares.SocialMediaAuthorization(),controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId",middlewares.SocialMediaAuthorization(),controllers.DeleteSocialMedia)
		
	}

	
	return r
}