package http

import (
	"demo-restful-api/config/http/books"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	books.Routes(r)

	return r
}
