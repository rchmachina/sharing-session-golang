package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	//"time"

	"github.com/rchmachina/sharing-session-golang/utils/helper"
	jwtToken "github.com/rchmachina/sharing-session-golang/utils/jwt"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized"})
			c.Abort()
			return
		}

		token = strings.Split(token, "Bearer ")[1]

		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized"})
			c.Abort()
			return
		}

		expiry, ok := claims["expiry"].(float64)
		if !ok {
			log.Println("expiry value is not of type float64")
			return
			
		}
		log.Println(expiry)
		roles, ok := claims["roles"].(string)
		if !ok {
			log.Println("expiry value is not of type float64")
			return
		}
		log.Println(roles)
		if float64(time.Now().Unix()) > expiry {
			helper.JSONResponse(c,401,"you are logged out")
			return
		}
		c.Set("roles",roles)
		handlerFunc(c)
	}
}
func IsAdmin(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles,isRolesExist := c.Get("roles")
		if !isRolesExist{c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized"})}

		role, ok := roles.(string) 
		if !ok{
			log.Println("Role:", role)
		}
		if role == "admin"{
			handlerFunc(c)
			return
		}
		helper.JSONResponse(c,401,"you are not admin")
		
	}
}


func UnmarshalToken(c *gin.Context) (User, error) {
	userMap := c.MustGet("userLogin")

	// Marshal the map back to JSON bytes
	jsonData, err := json.Marshal(userMap)
	if err != nil {
		return User{}, err
	}

	// Unmarshal the JSON bytes into the User struct
	var user User
	if err := json.Unmarshal(jsonData, &user); err != nil {
		return User{}, err
	}

	return user, nil
}

type User struct {
	Expiry   float64 `json:"expiry"`
	ID       string  `json:"id"`
	Roles    string  `json:"roles"`
	UserName string  `json:"userName"`
}
