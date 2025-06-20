package models

// User represents user's object model
// User has one-to-many relation to todos table
type User struct {
	Id       int64
	Username string
	Password string
	//Todos *[]Todo
}
