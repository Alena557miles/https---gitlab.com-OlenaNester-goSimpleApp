package todo

import (
	"fmt"
)

type TODO struct {
	owner *User
	tasks []Task
}

type User struct {
	name  string
	email string
	td    *TODO
	on    *Course
}

type Course struct {
	title string
	users []*User
}

func (t TODO) ShowTasks() {
	for _, task := range t.tasks {
		fmt.Println(task)
	}
}

func (t *TODO) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}

type Task struct {
	name        string
	description string
}

func (t Task) String() string {
	return fmt.Sprintf(`name: %s, description: %s`, t.name, t.description)
}

// NewTODO expects name, description, name, description and else...
func NewTODO(sl Getter, text ...string) *TODO {
	if len(text)%2 != 0 {
		panic(`arg must be even`)
	}

	var us NewUserer
	sl.Get(&us)
	user := us.NewUser()

	res := &TODO{owner: user}

	var i int
	for {
		res.AddTask(Task{text[i], text[i+1]})
		i += 2
		if i >= len(text) {
			break
		}
	}

	return res
}

type NewUserer interface {
	NewUser() *User
}

type Getter interface {
	Get(some any) bool
}
