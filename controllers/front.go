package controllers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware"
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
	"github.com/webAPi/controllers/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/webAPi/database"
)

func RegisterController() {
	/*	uc := newUserController()

		http.Handle("/users", *uc)
		http.Handle("/users/", *uc)
		http.Handle("/biometricVerification", *uc) */
	database.TestingDb()
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func test() {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	r := gin.Default()

	uc := newUserController()

	v1 := r.Group("/users")
	{
		accounts := v1.Group("")
		{
			accounts.GET("", uc.ListAccounts)
			/*	accounts.GET("", c.ListAccounts)
				accounts.POST("", c.AddAccount)
				accounts.DELETE(":id", c.DeleteAccount)
				accounts.PATCH(":id", c.UpdateAccount)
				accounts.POST(":id/images", c.UploadAccountImage)  */
		}
		//...
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
