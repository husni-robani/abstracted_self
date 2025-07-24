package models

type Profile struct {
	Name string `json:"name"`
	Summary string `json:"summary"`
	Bio string `json:"bio"`
	Taglines []string `json:"taglines"`
}