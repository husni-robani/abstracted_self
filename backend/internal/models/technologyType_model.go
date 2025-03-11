package models

type TechnologyType struct {
	Id int `json:"id"`
	TypeName string `json:"type_name"`
	Technologies []Technology `json:"technologies,omitempty"`
}