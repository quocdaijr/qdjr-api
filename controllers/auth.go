package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qdjr-api/helpers"
	"qdjr-api/requests"
	"qdjr-api/services"
	"strconv"
)

type AuthController struct{}

var authService = new(services.AuthService)

var authHelper = new(helpers.AuthHelper)
var baseHelper = new(helpers.BaseHelper)

func (authController AuthController) Register(c *gin.Context) {
	var input requests.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := authService.Register(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (authController AuthController) Login(c *gin.Context) {
	var input requests.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := authService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login Failed"})
		return
	}
	sub := strconv.Itoa(int(user.Id))
	ttl, _ := strconv.Atoi(baseHelper.GetEnv("SECRET_TTL", "60"))
	refreshTtl, _ := strconv.Atoi(baseHelper.GetEnv("REFRESH_SECRET_TTL", "84000"))
	token, err := authHelper.GenerateToken(sub, baseHelper.GetEnv("SECRET_KEY", ""), ttl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refreshToken, _ := authHelper.GenerateToken(sub, baseHelper.GetEnv("REFRESH_SECRET_KEY", ""), refreshTtl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":       token,
		"refresh_token":      refreshToken,
		"expires_in":         ttl,
		"refresh_expires_in": refreshTtl,
	})
}

func (authController AuthController) RefreshToken(c *gin.Context) {
	var input requests.RefreshTokenRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refreshToken := input.RefreshToken
	verifiedToken, err := authHelper.VerifyToken(refreshToken, baseHelper.GetEnv("REFRESH_SECRET_KEY", ""))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, _ := verifiedToken.Claims.GetSubject()
	ttl, _ := strconv.Atoi(baseHelper.GetEnv("SECRET_TTL", "60"))
	token, err := authHelper.GenerateToken(sub, baseHelper.GetEnv("SECRET_KEY", ""), ttl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refreshTtl, _ := strconv.Atoi(baseHelper.GetEnv("REFRESH_SECRET_TTL", "84000"))
	refreshToken, _ = authHelper.GenerateToken(sub, baseHelper.GetEnv("REFRESH_SECRET_KEY", ""), refreshTtl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":       token,
		"refresh_token":      refreshToken,
		"expires_in":         ttl,
		"refresh_expires_in": refreshTtl,
	})
}
