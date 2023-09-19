package controllers

import (
	"fmt"
	"jwtEx/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.Userlogindetails
	if err := context.ShouldBindJSON(&user); err != nil {
		context.XML(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.XML(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := Instance.Create(&user)
	if record.Error != nil {
		context.XML(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.XML(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Name})
}

type TokenRequest struct {
	Email    string `xml:"email"`
	Password string `xml:"password"`
}
type TokenResponse struct {
	Email string
	Token string
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.Userlogindetails
	var tokenres TokenResponse
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "please Sign up"})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := GenerateJWT(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	tokenres.Email = user.Email
	tokenres.Token = tokenString
	context.JSON(http.StatusOK, gin.H{"token": tokenres})
}

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
func GetUser(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
		c.Abort()
		return
	}
	claims, err := ValidateToken1(token)
	email := claims.Email

	var user models.Userlogindetails
	if err := Instance.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	fmt.Println(err)

	//c.JSON(200, gin.H{"token": claims})
}
func GetUserInXML(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.XML(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
		c.Abort()
		return
	}
	claims, err := ValidateToken1(token)
	email := claims.Email

	var user models.Userlogindetails
	if err := Instance.Where("email = ?", email).First(&user).Error; err != nil {
		c.XML(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.XML(http.StatusOK, gin.H{"data": user})
	fmt.Println(err)

	//c.JSON(200, gin.H{"token": claims})
}

func GenerateTokenByXML(context *gin.Context) {
	var request TokenRequest
	var user models.Userlogindetails
	var tokenres TokenResponse
	if err := context.ShouldBindXML(&request); err != nil {
		context.XML(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.XML(http.StatusInternalServerError, gin.H{"error": "please Sign up"})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.XML(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := GenerateJWT(user.Email)
	if err != nil {
		context.XML(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	tokenres.Email = user.Email
	tokenres.Token = tokenString
	context.XML(http.StatusOK, gin.H{"token": tokenres})
}
