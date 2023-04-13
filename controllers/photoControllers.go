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


func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Photo := models.Photo{}
	userID := userData["id"].(float64)
	User := models.User{}

	err := db.Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	if contentType == appJSON {
		c.BindJSON(&Photo)
	} else {
		c.Bind(&Photo)
	}

	Photo.UserID = uint(userID)
	Photo.User = &User

	err = db.Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}


func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}
	photoId,_ := strconv.Atoi(c.Param("photoId"))
	
	err := db.Preload("User").First(&Photo, "id= ?",photoId).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
	}
	c.JSON(http.StatusOK, Photo)
}


func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}
	
	err := db.Preload("User").Find(&Photos).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
	}
	c.JSON(http.StatusOK, Photos)
}

func UpdatePhoto(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Photo := models.Photo{}
	photoId,_ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.BindJSON(&Photo)
	} else {
		c.Bind(&Photo)
	}

	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title: Photo.Title,
		Caption: Photo.Caption,
		PhotoURL: Photo.PhotoURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated", "title": Photo.Title, "caption": Photo.Caption, "photoURL": Photo.PhotoURL})
}

func DeletePhoto(c *gin.Context){
	db := database.GetDB()
	Photo := models.Photo{}
	photoId,_ := strconv.Atoi(c.Param("photoId"))

	err := db.Delete(&Photo, "id = ?", photoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}

