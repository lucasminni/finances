package routes

import (
	financialV1 "finances/pkg/routes/v1"
	"log"

	"finances/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Grouper(r *gin.Engine) {
	group := r.Group("/api")
	v1 := group.Group("/v1")
	financialV1.Register(v1.Group("/finances"))
}

func Docs(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Financial control API"
	docs.SwaggerInfo.Description = "A simple service for financial control using http requests."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1/finances/"
	docs.SwaggerInfo.Schemes = []string{"http, https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Visit swagger at http://localhost:8080/swagger/index.html")
}
