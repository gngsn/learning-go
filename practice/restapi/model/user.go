package model

import "time"

// user struct
type User struct {
	ID		  int		`json:"id"`
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Email	  string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}