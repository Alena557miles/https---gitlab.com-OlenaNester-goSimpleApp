package user

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
type TODO struct {
	owner *User
	tasks []Task
}
type Task struct {
	name        string
	description string
}

func (u *User) ShowMyTasks(sl Getter) {
	var t TODOER
	sl.Get(&t)
	t.ShowTasks()
}

func (u *User) SetTODO(sl Getter) {
	var t TODOER
	sl.Get(&t)
	u.td = t.NewTODO()
}

func (u *User) SetCourse(c *Course) {
	u.on = c
}

func NewUser(n, e string) *User {
	return &User{name: n, email: e}
}

type TODOER interface {
	ShowTasks()
	NewTODO() *TODO
}
type Getter interface {
	Get(some any) bool
}
