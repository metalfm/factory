package entity

type User struct {
	ID       int
	Name     string
	Location string
	Friend   *User
}
