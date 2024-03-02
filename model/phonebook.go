package model

// //// input model struct for phone book ///////
type CreatePhonebook struct {
	Ext                  string    `json:"ext"`
	RoomId               string `json:"roomId"`
	UserNamePhonebook    string `json:"namePhonebook"`
	PhonebookDescription string `json:"description"`
	CreatedBy            string `json:"createdBy"`
}

type UpdatePhonebook struct {
	Ext                  string    `json:"ext"`
	RoomId               string `json:"roomId"`
	UserNamePhonebook    string `json:"namePhonebook"`
	PhonebookDescription string `json:"description"`
	PhonebookId          string `json:"phonebookId"`
}

/// output model struct for phone book ///

type PhonebookEntry struct {
	Ext                  int    `json:"ext"`
	RoomID               string `json:"roomId"`
	NameRoom             string `json:"nameRoom"`
	UserNamePhonebook    string `json:"userNamePhonebook"`
	PhonebookDescription string `json:"phonebookDescription"`
}

type GetPhonebook struct {
	Data      []interface{} `json:"dataPhonebook"`
	TotalData int           `json:"totalData"`
}
