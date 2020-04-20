package model

import "time"

type Place struct {
	Name      string      `json:"name,omitempty" schema:"name" structs:"name,omitempty" bson:"name,omitempty"`
	CityID    string      `json:"city_id,omitempty" schema:"city_id" structs:"city_id,omitempty" bson:"city_id,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty" schema:"created_at" structs:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" schema:"updated_at" structs:"updated_at,omitempty" bson:"updated_at,omitempty"`
	ID        interface{} `json:"id,omitempty" structs:",omitempty" bson:"_id,omitempty"`
}
