package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/rchmachina/sharing-session-golang/model"
	"github.com/rchmachina/sharing-session-golang/repositories"
	. "github.com/rchmachina/sharing-session-golang/utils/bcrypt"
	jwtToken "github.com/rchmachina/sharing-session-golang/utils/jwt"
)

type userHandler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *userHandler {
	return &userHandler{UserRepository}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	userData := new(model.CreateUser)

	if err := c.Bind(userData); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: err.Error(),
		})
		return
	}

	if userData.Roles != "admin" {
		userData.Roles = "nonadmin"
	}

	hashedPassword, err := HashingPassword(userData.HashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to hash password",
		})
		return
	}
	userData.HashedPassword = hashedPassword

	err = u.UserRepository.CreateUserDb(*userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": fmt.Sprintf("registration successful for username %s", userData.UserName),
	})
}

func (u *userHandler) LoginUser(c *gin.Context) {
	checkAuth := new(model.Login)
	if err := c.Bind(checkAuth); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: err.Error(),
		})
		return
	}

	user, err := u.UserRepository.LoginUserDB(checkAuth.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "credential is wrong",
		})
		return
	}

	isValid := CheckPasswordHash(checkAuth.Password, user.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "credential is wrong",
		})
		return
	}

	expiredTime := time.Now().Add(4 * time.Hour).Unix()
	user.Expired = expiredTime

	claims := jwt.MapClaims{
		"id":      user.UserId,
		"userName": user.UserName,
		"roles":   user.Roles,
		"expiry":  user.Expired,
	}

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userName": user.UserName,
		"roles":    user.Roles,
		"userId":   user.UserId,
		"token":    token,
		"expiry":   user.Expired,
	})
}

func (u *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	isDeleted, err := u.UserRepository.DeleteUserDb(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "there is something wrong",
		})
		return
	}

	if !isDeleted {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "there is something wrong",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Message: "success deleting user!",
	})
}

func (u *userHandler) UpdateUser(c *gin.Context) {
	updateUser := new(model.UpdateUser)
	if err := c.Bind(updateUser); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: err.Error(),
		})
		return
	}

	if updateUser.Password != "" {
		hashedPassword, err := HashingPassword(updateUser.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Response{
				Message: "Failed to hash password",
			})
			return
		}
		updateUser.Password = hashedPassword
	}

	isUpdated, err := u.UserRepository.UpdateUserDb(*updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "there is something wrong",
		})
		return
	}

	if !isUpdated {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "there is something wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "success updating user!" + updateUser.UserName,
	})
}

func (u *userHandler) GetAllUser(c *gin.Context) {
	page := c.Query("page")
	search := c.Query("search")

	getAllUser := u.UserRepository.GetAllUserDb(search, page)

	c.JSON(http.StatusOK, getAllUser)
}
