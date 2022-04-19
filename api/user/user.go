package user

type User struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Type    string `json:"type"`
}
