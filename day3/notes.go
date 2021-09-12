package day3

import "fmt"

/*
File này để notes lại code cần phải nhớ trong buổi thứ 3
*/

func modifyString(s *string) {
	t := *s
	*s = t[len(t)-1:] + t[1:len(t)-1] + t[:1]
	fmt.Println("Inside func s = ", s)
}

func DemoPointer() {
	s := "hello"
	modifyString(&s)
	fmt.Println("Outside func s = ", s)
}
