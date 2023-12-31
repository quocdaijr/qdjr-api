package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"qdjr-api/requests"
	"qdjr-api/services"
	"strconv"
)

type ArticleController struct{}

var articleService = new(services.ArticleService)

func (articleController ArticleController) List(c *gin.Context) {
	var input requests.SearchUserRequest
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	articles, totalItems := articleService.Search(input, input.Page, input.PerPage, input.Sort)
	c.JSON(
		http.StatusOK,
		gin.H{
			"articles":    articles,
			"total_items": totalItems,
			"total_pages": math.Ceil(float64(totalItems) / float64(input.PerPage)),
		},
	)
}

func (articleController ArticleController) Create(c *gin.Context) {

	var input requests.CreateArticleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	article, err := articleService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    article,
		})
	}
	return
}

func (articleController ArticleController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input requests.UpdateArticleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
		return
	}
	article, err := articleService.Update(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    article,
		})
	}
	return
}
