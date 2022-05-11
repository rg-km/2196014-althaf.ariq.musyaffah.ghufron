package handlers

import (
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	shorturl := c.Param("shorturl")
	url, err := h.repo.Get(shorturl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//should use gin redirect
	c.Redirect(http.StatusFound, url.LongURL)

	// c.Writer.Header().Set("Location", url.LongURL)
	// c.JSON(http.StatusFound, gin.H{
	// 	"message": "Redirect",
	// })

}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	req := c.ShouldBindJSON(&url)
	if req != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadRequest)
	}

	urlResponse, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"Data": urlResponse,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	req := c.ShouldBindJSON(&url)
	if req != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadRequest)
	}

	urlResponse, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": urlResponse,
	})

}
