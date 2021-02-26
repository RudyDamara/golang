package structs

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Islogin  bool   `json:"islogin"`
}

type ResultLogin struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"access_token"`
}

type AccessToken struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
