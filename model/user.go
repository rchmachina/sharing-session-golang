package model

type Login struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type LoginResponse struct {
	UserId   string `json:"userId"`
	Roles    string `json:"roles"`
	UserName string `json:"userName"`
	Password string `json:"hashedPassword"`
	Expired  int64  `json:"expired"`
}

type ReturnLogin struct {
	UserId   string `json:"userId"`
	Roles    string `json:"roles"`
	UserName string `json:"userName"`
}

type CreateUser struct {
	UserName       string `json:"userName"`
	HashedPassword string `json:"password"`
	Roles          string `json:"roles"`
}

type DeleteUser struct {
	UserId string `json:"userId"`
}

type UpdateUser struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type GetAllUser struct {
	Data      interface{} `json:"data"`
	TotalPage int         `json:"totalPage"`
	Page      int         `json:"page"`
}

type SearchUser struct {
	Search string `json:"searchName"`
	Page   int    `json:"page"`
}
