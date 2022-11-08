package main

import (
	"ci_hw/course"
	"ci_hw/sl"
	"ci_hw/todo"
	"ci_hw/user"
	"fmt"
	"reflect"
)

func main() {

	sl := &sl.ServiceLocator{}
	sl.Register(course.NewCourse)
	sl.Register(todo.NewTODO)

	u1 := user.NewUser(`Maks`, `Morozov`)
	sl.Register(u1)
	td1 := todo.NewTODO(sl,
		`mon`, `Trim the cat`,
		`tue`, `Wash the car`,
		`wed`, `By cell phone`,
		`thu`, `Play guitar`,
		`fri`, `Go sleep`,
	)
	sl.Register(td1)
	u1.SetTODO(sl)

	u2 := user.NewUser(`Alla`, `Pugachova`)
	sl.Register(u2)
	td2 := todo.NewTODO(sl,
		`mon`, `Bim the cat`,
		`tue`, `Wash the car`,
		`wed`, `By cell phone`,
		`thu`, `Fly guitar`,
		`fri`, `Go sleep`,
	)
	sl.Register(td2)
	u2.SetTODO(sl)

	c := course.NewCourse(`Golang for beginners`, sl, sl)
	// fmt.Println(`Course info:`, c)
	// u2.ShowMyTasks(sl)
	// u2.SetCourse(sl)
	// fmt.Println(u2.On)

	for _, u := range c.Users() {
		fmt.Println(u)
		fmt.Println(reflect.TypeOf(u))
		fmt.Println(reflect.TypeOf(u1))

		// u.ShowMyTasks()
	}
	// fmt.Println("Show my tasks(user)", u2.S)
}
