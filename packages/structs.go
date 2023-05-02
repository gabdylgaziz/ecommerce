package packages

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string 
}