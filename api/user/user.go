package api

type User struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Type    string `json:"type"`
}

//Use for signin and signup actions
type UserData struct {
	User     User   `json:"user"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
