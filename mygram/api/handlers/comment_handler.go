package handlers

import (
	"mygram/config"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	Photos := models.Photo{}
	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId

	err := db.First(&Photos, "id = ?", Comment.PhotoID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "Data not found",
		})
		return
	}

	err = db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         &Comment.ID,
		"message":    &Comment.Message,
		"photo_id":   &Comment.PhotoID,
		"user_id":    &Comment.UserID,
	})

}

func ShowComment(c *gin.Context) {
	db := config.ConnectDatabase()
	Comment := []models.Comment{}
	respComment := []models.Comment{}

	err := db.Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "Can't accessed data",
		})
	}

	User := models.User{}
	Photo := models.Photo{}

	for _, v := range Comment {
		db.First(&Photo, "id = ?", v.PhotoID)
		db.First(&User, "id = ?", v.UserID)
		input := models.Comment{}
		input.ID = v.ID
		input.UserID = v.UserID
		input.PhotoID = v.PhotoID
		input.Message = v.Message
		input.User.ID = User.ID
		input.User.Email = User.Email
		input.User.Username = User.Username
		input.Photo.ID = Photo.ID
		input.Photo.Title = Photo.Title
		input.Photo.Caption = Photo.Caption
		input.Photo.PhotoURL = Photo.PhotoURL
		input.Photo.UserID = Photo.UserID
		respComment = append(respComment, input)
	}
	c.JSON(http.StatusOK, respComment)
}

func UpdateComment(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	commentResult := models.Comment{}
	err = db.First(&commentResult, "id = ?", commentId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}
	if commentResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this comment",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err = db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	err = db.First(&commentResult, "id = ?", commentId).Error

	c.JSON(http.StatusOK, gin.H{
		"id":         &commentResult.ID,
		"message":    &commentResult.Message,
		"photo_id":   &commentResult.PhotoID,
		"user_id":    &commentResult.UserID,
	})

}

func DeleteComment(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	Comment := models.Comment{}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	commentResult := models.Comment{}

	err = db.First(&commentResult, "id = ?", commentId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}

	if commentResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you can't edit this comment",
		})
		return
	}

	err = db.Model(&Comment).Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	c.JSON(201, gin.H{
		"message": "Your Comment has been succesfully deleted",
	})

}
