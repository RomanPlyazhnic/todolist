package models

type Checkbox struct {
	Id         int64
	TodoListId int64
	Checked    bool
	Text       string
}
