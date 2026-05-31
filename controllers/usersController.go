package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/NaufalParamaRafif/CICO-Learn-Backend/initializers"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Username string `json:"username" form:"username" binding:"required,min=5"`
		Name     string `json:"name" form:"name" binding:"required,min=3"`
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required,min=8"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		fmt.Println("Error: " + err.Error())

		return
	}

	var user models.User
	users := initializers.DB.Where("username = ?", body.Username).Or("email = ?", body.Email).Find(&user)

	if users.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to signup, user has been registered",
		})

		return
	}
	if users.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to login, internal server error",
		})
		fmt.Println("Error: " + users.Error.Error())

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user = models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup success",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Failed to login, User not found",
		})

		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to login, internal server error",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) // Todo: settings this later (both true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}

func Validate(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
