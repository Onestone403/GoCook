package model

type Role string

const (
	COOK string = "cook"
	USER string = "user"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Role      Role
}
