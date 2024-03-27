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

func CreatePhoto(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	Photo := models.Photo{}
	userID := int(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         &Photo.ID,
		"title":      &Photo.Title,
		"caption":    &Photo.Caption,
		"photo_url":  &Photo.PhotoURL,
		"user_id":    &Photo.UserID,
	})

}

func ShowPhoto(c *gin.Context) {
	db := config.ConnectDatabase()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	Photo := []models.Photo{}
	RespPhoto := []models.Photo{}

	err := db.Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "photos cannot accessed",
		})
	}

	User := models.User{}

	for _, v := range Photo {
		db.First(&User, "id = ?", v.UserID)
		input := models.Photo{}
		input.ID = v.ID
		input.Title = v.Title
		input.UserID = v.UserID
		input.PhotoURL = v.PhotoURL
		input.Caption = v.Caption
		input.User.Username = User.Username
		input.User.Email = User.Email
		RespPhoto = append(RespPhoto, input)
	}
	c.JSON(http.StatusOK, RespPhoto)

}

func UpdatePhoto(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	Photo := models.Photo{}
	userId := int(userData["id"].(float64))

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	photoResult := models.Photo{}
	err = db.First(&photoResult, "id = ?", photoId).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}

	if photoResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this photo",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&photoResult, "id = ?", photoId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         &photoResult.ID,
		"title":      &photoResult.Title,
		"caption":    &photoResult.Caption,
		"photo_url":  &photoResult.PhotoURL,
		"user_id":    &photoResult.UserID,
	})
}

func DeletePhoto(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	// contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	Photo := models.Photo{}
	userId := int(userData["id"].(float64))

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	photoResult := models.Photo{}
	err = db.First(&photoResult, "id = ?", photoId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}
	if photoResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this photo",
		})
		return
	}

	Comment := models.Comment{}
	err = db.Model(&Comment).Where("photo_id = ?", photoId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "failed deleting comment from this photo",
		})
	}

	err = db.Model(&Photo).Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Your Photo has been succesfully deleted",
	})
}
