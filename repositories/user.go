package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"


	"github.com/rchmachina/sharing-session-golang/model"
	"github.com/rchmachina/sharing-session-golang/utils/helper"

	_ "github.com/lib/pq"

	"gorm.io/gorm"
)

// kontrak
type UserRepository interface {
	CreateUserDb(user model.CreateUser) error
	LoginUserDB(string) (model.LoginResponse, error)
	DeleteUserDb(userId string) error
	UpdateUserDb(user model.UpdateUser) error
	GetAllUserDb(model.SearchUser) model.GetAllUser
	
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUserDb(user model.CreateUser) error {
	var result map[string]interface{}
	fmt.Println("masuk pak cik", user)

	err := r.db.Raw("select create_user(?)", helper.ToJSON(user)).Scan(&result).Error
	if err != nil {
		return err
	}

	log.Println("result from repository ", helper.ToJSON(result))

	return nil
}

func (r *repository) LoginUserDB(userName string) (model.LoginResponse, error) {

	var responseLogin model.LoginResponse
	var result string

	paramsJSON, err := json.Marshal(map[string]interface{}{"userName": userName})
	if err != nil {
		return responseLogin, err
	}

	err = r.db.Raw("SELECT * blm jadi njeng", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return responseLogin, err
	}
	err = json.Unmarshal([]byte(result), &responseLogin)
	if err != nil {
		return responseLogin, errors.New("not found")
	}
	fmt.Print("isi response", responseLogin)

	return responseLogin, err
}

func (r *repository) DeleteUserDb(userId string) error {

	var result map[string]interface{}
	fmt.Println("masuk pak cik", userId)

	err := r.db.Raw("select users_delete(?)", helper.ToJSON(userId)).Scan(&result).Error
	if err != nil {
		return err
	}

	log.Println("result from repository ", helper.ToJSON(result))

	return nil

}

func (r *repository) UpdateUserDb(user model.UpdateUser) error {
	var result map[string]interface{}
	fmt.Println("masuk pak cik", user)

	err := r.db.Raw("select users_update(?)", helper.ToJSON(user)).Scan(&result).Error
	if err != nil {
		return err
	}

	log.Println("result from repository ", helper.ToJSON(result))

	return nil
}
func (r *repository) GetAllUserDb(query model.SearchUser) model.GetAllUser {

	var GetAllUser model.GetAllUser

	var result string
	err := r.db.Raw("SELECT users_get_all(?)", helper.ToJSON(query)).Scan(&result).Error
	if err != nil {
		log.Println(err)
		return GetAllUser
	}
	err = json.Unmarshal([]byte(result), &GetAllUser)
	if err != nil {
		log.Println(err)
		return GetAllUser
	}

	return GetAllUser
}
