package vo

type User struct {
	IdUser string `json:"userid"`
	Email  string `json:"email"`
}

type Schedule struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Users       string `json:"users"`
}

type Output struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Emails      string `json:"emails"`
}
