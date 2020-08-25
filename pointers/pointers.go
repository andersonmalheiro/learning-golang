package pointers

import (
	"time"
)

// A pointer is a memory address reference
// The zero value of a pointer is nil
var pointer *int // An integer pointer declaration

// User is an user struct
// Structs are a collection of fields
type User struct {
	Name      string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Structs can be created passing all parameters like this
var fullUser = User{"Anderson", "Major", time.Now(), time.Now()}

// Or passing only specific params with their names
var partialUser = User{Name: "Test"}

// CreateUser creates an user and returns it
func CreateUser(name string, nickname string) (user User) {
	user = User{
		Name:      name,
		Nickname:  nickname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return
}

// UpdateUser is a function to change an user data
// When passing a pointer as parameter in a function we can change the value of the original variable
// When we pass just the variable, the function receives only a copy of the value
func UpdateUser(u *User, name string, nickname string) {
	u.Name = name
	u.Nickname = nickname
	u.UpdatedAt = time.Now()
}
