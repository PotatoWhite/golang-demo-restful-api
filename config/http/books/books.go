package books

import (
	"demo-restful-api/config/database"
	"demo-restful-api/logics"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	api := &logics.APIEnv{
		DB: database.GetDB(),
	}

	books := r.Group("/books")

	books.GET("", api.GetBooks)
	books.GET("/:id", api.GetBook)
	books.POST("", api.CreateBook)
	books.PUT("/:id", api.UpdateBook)
	books.DELETE("/:id", api.DeleteBook)
}
