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

func CreateSocialMedia(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userId

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               &SocialMedia.ID,
		"name":             &SocialMedia.Name,
		"social_media_url": &SocialMedia.SocialMediaURL,
		"user_id":          &SocialMedia.UserID,
	})
}

func ShowSocialMedia(c *gin.Context) {
	db := config.ConnectDatabase()
	SocialMedia := []models.SocialMedia{}
	err := db.Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "Data cannot accessed",
		})
	}
	User := models.User{}

	for _, v := range SocialMedia {
		db.First(&User, "id = ?", v.UserID)
		input := models.SocialMedia{}
		input.ID = v.ID
		input.Name = v.Name
		input.SocialMediaURL = v.SocialMediaURL
		input.UserID = v.UserID
		input.User.ID = User.ID
		input.User.Username = User.Username
		SocialMedia = append(SocialMedia, input)
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	socialMediaResult := models.SocialMedia{}
	err = db.First(&socialMediaResult, "id = ?", socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}
	if socialMediaResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this Social Media",
		})
		return
	}
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err = db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	err = db.First(&socialMediaResult, "id = ?", socialMediaId).Error
	c.JSON(http.StatusOK, gin.H{
		"id":               &socialMediaResult.ID,
		"name":             &socialMediaResult.Name,
		"social_media_url": &socialMediaResult.SocialMediaURL,
		"user_id":          &socialMediaResult.UserID,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := config.ConnectDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	SocialMedia := models.SocialMedia{}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	socialMediaResult := models.SocialMedia{}
	err = db.First(&socialMediaResult, "id = ?", socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}
	if socialMediaResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this Social Media",
		})
		return
	}

	err = db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(201, gin.H{
		"message": "Your Social Media has been succesfully deleted",
	})
}
