package models

import(
	"fmt"
)
type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

func GetUsers() []*User  {
	return users	
}

func AddUser (u User) (User , error){
	if u.ID != 0 {
		return User {} , nil
	}
	u.ID = nextID
	nextID++
	users = append(users,&u)
	return u , nil
}

func GetUserByID (id int) (User , error){
	for _,u := range users {
		if u.ID == id {
			return *u , nil
		}
	}
	return User {}, fmt.Errorf("No such User");
}