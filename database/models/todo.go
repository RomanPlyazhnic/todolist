package models

// TodoList represents todolist's object model
// TodoList has one-to-many relation to checkboxes
type TodoList struct {
	Id     int64
	UserId int64
	Text   string
	//Checkboxes *[]Checkbox
}
