package structs

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Islogin  bool   `json:"islogin"`
}
