package model

//Person object for REST(CRUD)
type ContactService struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string  `json:"email"`
}
