package domain

// domain is basically your models

type User struct {
	ID       int
	Email    string
	Password string
}

// acts as constructor in OOP
func NewPerson(id int, email string, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

// methods
func (u *User) GetEmail() string {
	return u.Email
}
