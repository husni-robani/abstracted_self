package models

type Profile struct {
	Name string `json:"name"`
	Summary string `json:"summary"`
	Bio string `json:"bio"`
	Taglines []string `json:"taglines"`
	ResumeFileName string `json:"resume_file_name"`
	Skills []string `json:"skills"`
}

