package main

import (
	"github.com/abiosoft/ishell"
)

func initShell() *ishell.Shell {
	shell := ishell.New()
	shell.Println("Gatekeeper Server Console")
	shell.Println("--------------------------------------")
	shell.Println(" © Isaac True <isaac@is.having.coffee>")
	shell.Println()
	/*
		shell.AddCmd(&ishell.Cmd{
			Name: "addUser",
			Help: "Add a new user",
			Func: shellAddUser,
		})*/

	return shell
}

/*
func shellAddUser(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	var email string
	if len(c.Args) > 0 {
		email = c.Args[0]
	} else {
		c.Print("Email: ")
		email = c.ReadLine()
	}

	c.Print("First Name: ")
	firstName := c.ReadLine()
	c.Print("Last Name: ")
	lastName := c.ReadLine()

	c.Print("Password: ")
	password := c.ReadPassword()

	if err := users.Users.AddUser(email, password, firstName, lastName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK!")
	}
}
*/
