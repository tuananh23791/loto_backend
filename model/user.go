package model

type User struct {
	PhoneNumber string      `json:"phone_number,omitempty" schema:"phone_number" structs:"phone_number,omitempty" bson:"phone_number,omitempty"`
	FullName    string      `json:"full_name,omitempty" schema:"full_name" structs:"full_name,omitempty" bson:"full_name,omitempty"`
	Password    string      `json:"-" schema:"password" structs:"password,omitempty"`
	Role        string      `json:"role,omitempty" schema:"role" structs:"role,omitempty" bson:"role,omitempty"`
	Token       string      `json:"token,omitempty" schema:"token" structs:"token,omitempty" bson:"token,omitempty"`
	ID          interface{} `json:"id,omitempty" structs:",omitempty" bson:"_id,omitempty"`
}
