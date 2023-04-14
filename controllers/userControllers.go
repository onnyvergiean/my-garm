package controllers

import (
	"my-garm/database"
	"my-garm/helpers"
	"my-garm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"


// UserRegister godoc
// @Summary Register new user
// @Description Register new user
// @ID register-user
// @Tags User
// @Accept  json
// @Produce  json
// @Param UserRegister body models.UserRegister true "UserRegister"
// @Success 200 {object} models.User
// @Router /users/register [post]
func UserRegister(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	

	if contentType == appJSON {
		c.BindJSON(&User)
	} else {
		c.Bind(&User)
	}

	err := db.Create(&User).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": User.ID, "username": User.Username, "email": User.Email})
}


// UserLogin godoc
// @Summary Login user
// @Description Login user
// @Tags User
// @ID login-user
// @Accept json
// @Produce json
// @Param UserLogin body models.UserLogin true "email"
// @Success 200 {object} string "token"
// @Router /users/login [post]
func UserLogin(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.BindJSON(&User)
	} else {
		c.Bind(&User)
	}

	password = User.Password
	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	comparePassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Password is not match"})
		return
	}

	token, err := helpers.GenerateToken(User.ID, User.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}