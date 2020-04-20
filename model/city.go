package model

type City struct {
	Name      string      `json:"name,omitempty" schema:"name" structs:"name,omitempty" bson:"name,omitempty"`
	SortName  string      `json:"sort_name,omitempty" schema:"sort_name" structs:"sort_name,omitempty" bson:"sort_name,omitempty"`
	CreatedAt string      `json:"created_at,omitempty" schema:"created_at" structs:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty" schema:"updated_at" structs:"updated_at,omitempty" bson:"updated_at,omitempty"`
	ID        interface{} `json:"id,omitempty" structs:",omitempty" bson:"_id,omitempty"`
}
