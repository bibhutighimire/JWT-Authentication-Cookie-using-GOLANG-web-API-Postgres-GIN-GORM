package controllers

import (
	"gocrudapibackend/initializers"
	"gocrudapibackend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "this email aready exists",
		})
		return
	}

	//hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "problem hashing password",
		})
		return
	}

	// create user record

	newuser := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&newuser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "uSER COULD NOT BE CREATED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "USER CREATED",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or Password incorrect",
		})
		return
	}

	//check if both the password matches or not

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or Password incorrect",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token string could not be created",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
