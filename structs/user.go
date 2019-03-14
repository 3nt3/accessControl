package structs

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Permission int `json:"permission"`
	UID string `json:"uid"`
} 
