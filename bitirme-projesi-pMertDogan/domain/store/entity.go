package store

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Stores []Store

type Store struct {
	gorm.Model
	Name string
	Description string
	Phone string
	Email string
	Address string
}

// fromJson Store
func UnmarshalStore(data []byte) (Stores, error) {
	var r Stores
	err := json.Unmarshal(data, &r)
	return r, err
}


//Stores toJson
func (r *Stores) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//Store toJson
func (r *Store) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
