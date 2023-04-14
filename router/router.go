package router

import (
	"my-garm/controllers"
	"my-garm/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My-Garm API Documentation
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE0ODA2MDUsImlzcyI6Ik1ZLUdBUk0iLCJpZCI6NCwiRW1haWwiOiJva0Bvay5jb20ifQ.bPcX66MlkvCWypEXpLEadFToWYDoe8w3iq9YXp_JlHs

// @description This is a sample server for a photo sharing app.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @version 1.0.0
// @host my-garm-production-b7b2.up.railway.app
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
		socialMediaRouter.POST("/",controllers.CreateSocialMedia)
		socialMediaRouter.GET("/",controllers.GetSocialMedias)
		socialMediaRouter.GET("/:socialMediaId",controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId",middlewares.SocialMediaAuthorization(),controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId",middlewares.SocialMediaAuthorization(),controllers.DeleteSocialMedia)
		
	}

	
	
	return r
}