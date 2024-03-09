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
	Address        string `json:"address"`
}

type DeleteUser struct {
	UserId string `json:"userId"`
}

type SetAndUpdateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UpdateUser struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Password   string `json:"newPassword"`
	NewAddress string `json:"newAddress"`
}

type GetAllUser struct {
	GetUserData  []interface{} `json:"getUserData"`
	TotalRecords int           `json:"TotalRecords"`
	Page         int           `json:"page"`
}

type SearchUser struct {
	Search   string `json:"searchQuery"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
