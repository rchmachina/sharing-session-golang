package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/rchmachina/sharing-session-golang/model"
	"github.com/rchmachina/sharing-session-golang/utils/helper"

	_ "github.com/lib/pq"

	"gorm.io/gorm"
)

// kontrak
type UserRepository interface {
	CreateUserJson(user model.CreateUser) error
	CreateUserDb(user model.CreateUser) error
	LoginUserDB(string) (model.LoginResponse, error)
	DeleteUserDb(string) (bool, error)
	UpdateUserDb(user model.UpdateUser) (bool, error)
	GetAllUserDb(string, string) model.GetAllUser
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUserDb(user model.CreateUser) error {
	var result bool

	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = r.db.Raw("SELECT * FROM pklbram.createuser($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return err
	}

	if !result {
		return errors.New("not found")
	}


	return nil
}
func (r *repository) CreateUserJson(user model.CreateUser) error {
	var result map[string]interface{}
	fmt.Println("masuk pak cik",user)

	
	err := r.db.Raw("select create_user(?)", helper.ToJSON(user)).Scan(&result).Error
	if err != nil {
		return err
	}

	log.Println("isi dari json biasa ae" , helper.ToJSON(result))


	return nil
}

func (r *repository) LoginUserDB(userName string) (model.LoginResponse, error) {

	var responseLogin model.LoginResponse
	var result string

	paramsJSON, err := json.Marshal(map[string]interface{}{"userName": userName})
	if err != nil {
		return responseLogin, err
	}


	err = r.db.Raw("SELECT * FROM pklbram.login_user($1::jsonb)", string(paramsJSON)).Scan(&result).Error
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

func (r *repository) DeleteUserDb(userId string) (bool, error) {
	var isDeleted bool = false
	var result string
	fmt.Println("masuk pak cik", userId)
	paramsJSON, err := json.Marshal(map[string]interface{}{
		"userId": userId,
	})
	if err != nil {
		return isDeleted, err
	}

	err = r.db.Raw("SELECT * from pklbram.delete_user($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return isDeleted, err
	}

	err = json.Unmarshal([]byte(result), &isDeleted)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return isDeleted, err
	}

	return isDeleted, err
}

func (r *repository) UpdateUserDb(user model.UpdateUser) (bool, error) {
	var isUpdated bool
	var result string
	fmt.Println("masuk pak cik")
	paramsJSON, err := json.Marshal(user)
	if err != nil {
		return isUpdated, err
	}

	err = r.db.Raw("SELECT * from pklbram.update_user($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return isUpdated, err
	}
	err = json.Unmarshal([]byte(result), &isUpdated)
	if err != nil {
		log.Println(err)
		return isUpdated, err
	}

	return isUpdated, nil
}
func (r *repository) GetAllUserDb(searchName string, page string) model.GetAllUser {

	var GetAllUser model.GetAllUser

	var searchQuery model.SearchUser

	pageData, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
	}

	searchQuery.Page = pageData
	searchQuery.Search = searchName

	paramsJSON, err := json.Marshal(searchQuery)

	if err != nil {
		fmt.Println("ko ade error ke?")
	}

	var result string
	err = r.db.Raw("select * from pklbram.get_all_user($1::jsonb)", string(paramsJSON)).Scan(&result).Error
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
