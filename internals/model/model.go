package model

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DbEnvs struct {
	DbName    string
	DbUserCol string
}
