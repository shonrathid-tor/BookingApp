package controllers

import (
	"booking_app/config"
	"booking_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})

		return

	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message ": " User refisterd succssfull !!!"})

}

func Login(c *gin.Context) {
	var user models.User

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	//check  user 's email
	config.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful"})

}
