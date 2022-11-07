package main

import (
	"ci_hw/sl"
	"ci_hw/todo"
	"ci_hw/user"
)

func main() {
	sl := &sl.ServiceLocator{}
	u1 := user.NewUser(`Maks`, `Morozov`)
	sl.Register(user.NewUser(`Maks`, `Morozov`))
	td1 := todo.NewTODO(sl,
		`mon`, `Trim the cat`,
		`tue`, `Wash the car`,
		`wed`, `By cell phone`,
		`thu`, `Play guitar`,
		`fri`, `Go sleep`,
	)
	sl.Register(td1)
	u1.SetTODO(sl)

	// u2 := user.NewUser(`Alla`, `Pugachova`)
	// sl.Register(u2)
	// td2 := todo.NewTODO(sl,
	// 	`mon`, `Trim the cat`,
	// 	`tue`, `Wash the car`,
	// 	`wed`, `By cell phone`,
	// 	`thu`, `Play guitar`,
	// 	`fri`, `Go sleep`,
	// )
	// sl.Register(td2)
	// u2.SetTODO(sl)

	// c := course.NewCourse(`Golang for beginners`, u1, u2)
	// for _, u := range c.Users() {
	// 	u.ShowMyTasks(sl)
	// }
}
