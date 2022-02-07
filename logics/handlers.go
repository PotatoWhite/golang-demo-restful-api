package logics

import (
	"demo-restful-api/models"
	"demo-restful-api/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBooks(c *gin.Context) {
	books, err := repositories.GetBooks(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func (a *APIEnv) GetBook(c *gin.Context) {
	id := c.Params.ByName("id")
	book, exists, err := repositories.GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no book in db")
		return
	}

	c.JSON(http.StatusOK, book)
}

func (a *APIEnv) CreateBook(c *gin.Context) {
	book := models.Book{}
	if err := c.BindJSON(&book); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := a.DB.Create(&book).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (a *APIEnv) DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := repositories.DeleteBook(id, a.DB); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
	return
}

func (a *APIEnv) UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")
	if _, exists, err := repositories.GetBookByID(id, a.DB); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	} else if !exists {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	updatedBook := models.Book{}
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := repositories.UpdateBook(a.DB, &updatedBook).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	a.GetBook(c)
}
