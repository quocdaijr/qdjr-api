package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"qdjr-api/requests"
	"qdjr-api/services"
	"strconv"
)

type UserController struct{}

var userService = new(services.UserService)

func (userController UserController) Search(c *gin.Context) {
	var input requests.SearchUserRequest
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, totalItems := userService.Search(input, input.Page, input.PerPage, input.Sort)
	c.JSON(
		http.StatusOK,
		gin.H{
			"users":       users,
			"total_items": totalItems,
			"total_pages": math.Ceil(float64(totalItems) / float64(input.PerPage)),
		},
	)
}

func (userController UserController) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := userService.Detail(id)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}
