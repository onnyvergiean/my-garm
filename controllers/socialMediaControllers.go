package controllers

import (
	"errors"
	"my-garm/database"
	"my-garm/helpers"
	"my-garm/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CreateSocialmedia(c *gin.Context){
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := userData["id"].(float64)
	user := models.User{}


	err := db.First(&user, "id = ?" , userID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found", "message": err.Error()})
		return
	}

	if contentType == appJSON {
		c.BindJSON(&SocialMedia)
	} else {
		c.Bind(&SocialMedia)
	}

	SocialMedia.UserID = uint(userID)
	SocialMedia.User = &user

	err = db.Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetSocialMedia(c *gin.Context){
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	socialMediaId,_ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Preload("User").First(&SocialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func GetSocialMedias(c *gin.Context){
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}

	err := db.Preload("User").Order("id").Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, SocialMedias)
}

func UpdateSocialMedia(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	socialMediaId,_ := strconv.Atoi(c.Param("socialMediaId"))

	if contentType == appJSON {
		c.BindJSON(&SocialMedia)
	} else {
		c.Bind(&SocialMedia)
	}

	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Updates(models.SocialMedia{
		Name: SocialMedia.Name,
		SocialMediaURL: SocialMedia.SocialMediaURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK,  gin.H{"message": "Social Media Updated", "Name": SocialMedia.Name, "URL": SocialMedia.SocialMediaURL})
}

func DeleteSocialMedia(c *gin.Context){
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	socialMediaId,_ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Delete(&SocialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social Media Deleted"})
}

