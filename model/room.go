package model

// /// output model stuct ///////
// struct for model get all room

type GetAllRoom struct {
	Data         interface{} `json:"data"`
	TotalPages   int         `json:"totalPages"`
	TotalRecords int         `json:"totalRecords"`
}

type RoomData struct {
	RoomID         string `json:"roomId"`
	NameRoom       string `json:"nameRoom"`
	CreatedAt      string `json:"created_at"`
	CreatedBy      string `json:"created_by"`
	Description    string `json:"description"`
	TotalPhonebook int    `json:"totalPhonebook"`
}

// struct for model get specific room
type GetRoom []struct {
	RoomId      string   `json:"roomId"`
	NameRoom    string   `json:"nameRoom"`
	CreatedAt   string   `json:"created_at"`
	CreatedBy   string   `json:"created_by"`
	Description string   `json:"description"`
	InTheRoom   []InRoom `json:"inTheRoom"`
}

// struct for childmodel of get room
type InRoom struct {
	Ext  int    `json:"ext"`
	Name string `json:"name"`
}

// ////// input model struct room /////////
// struct for create new room
type CreateRoom struct {
	RoomName    string `json:"roomName"`
	CreatedBy   string `json:"createdBy"`
	Description string `json:"description"`
}

type DeleteRoom struct {
	RoomId string `json:"roomId"`
}

type UpdateRoom struct {
	RoomName    string `json:"roomName"`
	Description string `json:"description"`
	RoomId      string `json:"roomId"`
}
type SearchRoom struct {
	Search string `json:"searchRoom"`
	Page   string `json:"pageNumber"`
}

type ChoiceRoom [] struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
