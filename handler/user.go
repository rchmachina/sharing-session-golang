package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/rchmachina/sharing-session-golang/model"
	"github.com/rchmachina/sharing-session-golang/repositories"
	. "github.com/rchmachina/sharing-session-golang/utils/bcrypt"
	"github.com/rchmachina/sharing-session-golang/utils/helper"
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
		helper.JSONResponse(c, 400, err.Error())
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
		helper.JSONResponse(c, 200, err.Error())
		return
	}

	resp := model.Response{
		Message: fmt.Sprintf("registration successful for username %s", userData.UserName),
		Status: "success",
	}
	
	helper.JSONResponse(c, 200, resp)
}

func (u *userHandler)LoginUser(c *gin.Context) {
	checkAuth := new(model.Login)
	if err := c.Bind(checkAuth); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: err.Error(),
		})
		return
	}

	user, err := u.UserRepository.LoginUserDB(checkAuth.UserName)
	if err != nil {
		helper.JSONResponse(c, 401, "no user found")
		return
	}

	isValid := CheckPasswordHash(checkAuth.Password, user.Password)
	if !isValid {
		helper.JSONResponse(c, 401, "credential is wrong")
		return
	}

	expiredTime := time.Now().Add(4 * time.Hour).Unix()
	user.Expired = expiredTime

	claims := jwt.MapClaims{
		"id":       user.UserId,
		"userName": user.UserName,
		"roles":    user.Roles,
		"expiry":   user.Expired,
	}

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		helper.JSONResponse(c, 401, errGenerateToken)
		return
	}

	returnData := map[string]interface{}{
		"userName": user.UserName,
		"roles":    user.Roles,
		"userId":   user.UserId,
		"token":    token,
		"expiry":   user.Expired,
	}

	resp := model.Response{
		Message: returnData,
		Status: "success",
	}

	helper.JSONResponse(c, 200, resp)
}

func (u *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	message, err := u.UserRepository.DeleteUserDb(id)
	if err != nil {
		helper.JSONResponse(c, 400, message)
		return
	}

	helper.JSONResponse(c, 200, message)
}

func (u *userHandler) UpdateUser(c *gin.Context) {
	updateUser := new(model.UpdateUser)
	if err := c.Bind(updateUser); err != nil {
		helper.JSONResponse(c, 400, err.Error())
		return
	}

	if updateUser.Password != "" {
		hashedPassword, err := HashingPassword(updateUser.Password)
		if err != nil {
			log.Println("update user err:", err)
			helper.JSONResponse(c, 500, "Failed to hash password")
			return
		}
		updateUser.Password = hashedPassword
	}

	err := u.UserRepository.UpdateUserDb(*updateUser)

	if err != nil {
		log.Println("update user err:", err)
		helper.JSONResponse(c, 500, "there is something wrong")
		return
	}

	resp := model.Response{
		Message: fmt.Sprintf("success updating user %s", updateUser.UserName),
		Status: "success",
	}

	helper.JSONResponse(c, 200, resp)
}

func (u *userHandler) GetAllUser(c *gin.Context) {
	search := c.Query("search")
	page := c.Query("page")
	pageInt, _ := strconv.Atoi(page)
	pageSize := c.Query("pageSize")
	pageSizeInt, _ := strconv.Atoi(pageSize)



	searchGetAllUser := model.SearchUser{
		Page:     pageInt,
		PageSize: pageSizeInt,
		Search:   search,
	}

	getAllUser := u.UserRepository.GetAllUserDb(searchGetAllUser)

	c.JSON(http.StatusOK, getAllUser)
}
