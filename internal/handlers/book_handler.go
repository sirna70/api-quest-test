package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"api-quest-test/internal/models"
	"api-quest-test/internal/storage"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and author required"})
		return
	}

	newBook.ID = strconv.Itoa(len(storage.Books) + 1)
	storage.Books = append(storage.Books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func GetBooks(c *gin.Context) {
	author := c.Query("author")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	filtered := storage.Books

	if author != "" {
		temp := []models.Book{}
		for _, b := range storage.Books {
			if strings.EqualFold(b.Author, author) {
				temp = append(temp, b)
			}
		}
		filtered = temp
	}

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit

		if start >= len(filtered) {
			filtered = []models.Book{}
		} else {
			if end > len(filtered) {
				end = len(filtered)
			}
			filtered = filtered[start:end]
		}
	}

	c.JSON(http.StatusOK, filtered)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, b := range storage.Books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var updated models.Book
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	for i, b := range storage.Books {
		if b.ID == id {
			storage.Books[i].Title = updated.Title
			storage.Books[i].Author = updated.Author
			c.JSON(http.StatusOK, storage.Books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, b := range storage.Books {
		if b.ID == id {
			storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}
