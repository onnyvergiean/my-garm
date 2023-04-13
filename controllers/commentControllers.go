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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Comment := models.Comment{}
	userID := userData["id"].(float64)
	

	photoID, err := strconv.Atoi(c.Param("photoId"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	User := models.User{}
	err = db.Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	Photo := models.Photo{}
	err = db.Where("id = ?", photoID).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	if contentType == appJSON {
		c.BindJSON(&Comment)
	} else {
		c.Bind(&Comment)
	}

	Comment.UserID = uint(userID)
	Comment.User = &User

	Comment.PhotoID = uint(photoID)
	Comment.Photo = &Photo

	err = db.Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}


func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	commentId,err  := strconv.Atoi(c.Param("commentId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	photoId, err  := strconv.Atoi(c.Param("photoId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Preload("User").Preload("Photo").Where("photo_id = ?", photoId).First(&Comment, "id= ?",commentId).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Comment)
}


func GetComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}
	photoId, err  := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}
	err = db.Preload("User").Preload("Photo").Where("photo_id = ?", photoId).Find(&Comments).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
	}

	c.JSON(http.StatusOK, Comments)
}


func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	commentId,_ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}


	if contentType == appJSON {
		c.BindJSON(&Comment)
	} else {
		c.Bind(&Comment)
	}

	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		"comment": Comment.Message,
	})
}


func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId,_ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}