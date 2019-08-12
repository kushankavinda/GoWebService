package servises

import (
	"fmt"

	"github.com/webAPi/models"
	"github.com/webAPi/repository"
)

type User interface {
	UserService() string
}

func UserService() string {
	fmt.Print("ranji")
	repository.Create(struct_name{})
	return "user serviss"
}

func AddUser(u models.User) (models.User, error) {
	fmt.Print("user added : " + u.FirstName)

	fmt.Print(u.ID)
	repository.Create(u)

	return u, nil
}

/* define an interface */
type interface_name interface {
	method_name1() string
}

/* define a struct */
type struct_name struct {
	/* variables */
	name  string
	email string
}

/* implement interface methods*/
func (struct_name_variable struct_name) method_name1() string {
	/* method implementation */
	return "yes"
}
