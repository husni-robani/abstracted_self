package models


type Technology struct {
	Id int `json:"id,omitempty"`	
	TypeId int `json:"type_id,omitempty"`
	Name string `json:"name,omitempty"`
}