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

// CreateComment godoc
// @Summary Create Comment
// @Description Create New Comment Based on Photo ID
// @Security Bearer
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param photoId path int true "photo id"
// @Param CommentCreate body models.CommentCreate true "create comment"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string "Bad Request"
// @Router /comments/{photoId} [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

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


// GetComment godoc
// @Summary Get Comment
// @Description Get Details of Comment by corresponding ID
// @Security Bearer
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param photoId path int true "photo id"
// @Param commentId path int true "comment id"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string "Bad Request"
// @Router /comments/{photoId}/{commentId} [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	commentId,_  := strconv.Atoi(c.Param("commentId"))
	photoId, err  := strconv.Atoi(c.Param("photoId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Preload("User").Preload("Photo").Preload("Photo.User").Where("photo_id = ?", photoId).First(&Comment, "id= ?",commentId).Error
	

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

// GetComments godoc
// @Summary Get Comments
// @Description Get Details of Comments by corresponding Photo ID
// @Security Bearer
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param photoId path int true "photo id"
// @Success 200 {object} []models.Comment
// @Failure 400 {object} string "Bad Request"
// @Router /comments/{photoId} [get]
func GetComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}
	photoId, err  := strconv.Atoi(c.Param("photoId"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}
	err = db.Preload("User").Preload("Photo").Preload("Photo.User").Where("photo_id = ?", photoId).Find(&Comments).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
	}

	c.JSON(http.StatusOK, Comments)
}

// UpdateComment godoc
// @Summary Update Comment
// @Description Update Comment by corresponding ID
// @Security Bearer
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param photoId path int true "photo id"
// @Param commentId path int true "comment id"
// @Param CommentUpdate body models.CommentCreate true "update comment"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string "Bad Request"
// @Router /comments/{photoId}/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	commentId,_ := strconv.Atoi(c.Param("commentId"))
	photoId,err := strconv.Atoi(c.Param("photoId"))
	Comment := models.Comment{}
	Photo := models.Photo{}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Where("id = ?", photoId).Take(&Photo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return		 
	}

	if contentType == appJSON {
		c.BindJSON(&Comment)
	} else {
		c.Bind(&Comment)
	}

	Comment.ID = uint(commentId)

	err = db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return		 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		"comment": Comment.Message,
	})
}


// DeleteComment godoc
// @Summary Delete Comment
// @Description Delete Comment by corresponding ID
// @Security Bearer
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param photoId path int true "photo id"
// @Param commentId path int true "comment id"
// @Success 200 {message} string "Comment deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Router /comments/{photoId}/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId,_ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	photoId,err := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Where("id = ?", photoId).Take(&Photo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	err = db.Where("id = ? ", commentId).Delete(&Comment).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return		 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}