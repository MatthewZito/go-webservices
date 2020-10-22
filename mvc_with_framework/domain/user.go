package domain

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}
