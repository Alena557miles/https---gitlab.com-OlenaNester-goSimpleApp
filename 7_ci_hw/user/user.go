package user

type User struct {
	name  string
	email string
	td    TODOer
	on    Courser
}

func (u *User) ShowMyTasks(sl Getter) {
	var sh ShowTasker
	sl.Get(&sh)
	sh.ShowTasks()
}

func (u *User) SetTODO(sl Getter) {
	var t TODOer
	sl.Get(&t)
	u.td = t
}

func (u *User) SetCourse(sl Getter) {
	var c Courser
	sl.Get(&c)
	u.on = c
}

func NewUser(n, e string) *User {
	return &User{name: n, email: e}
}

type TODOer interface {
}

type Courser interface {
}
type ShowTasker interface {
	ShowTasks()
}

type Getter interface {
	Get(some any) bool
}
