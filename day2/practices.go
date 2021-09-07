package day2

import (
	"fmt"

	u "github.com/wayarmy/learning-go/helper"
)

func DoHomeWork() {
	str, err := u.ReadNumberFromKeyboard("Chọn bài tập muốn giải (1, 2, 3, 4): ")
	u.CheckErr(err)
	fmt.Println("")

	switch {
	case str == 1:
		practice01()
	case str == 2:
		practice02()
	case str == 3:
		practice03()
	case str == 4:
		practice04()
	default:
		fmt.Println("Chỉ chấp nhận 1/2/3, mời chọn lại.!")
		DoHomeWork()
	}
}

// Giải bài tập số 1
func practice01() {

}

// Giải bài tập số 2
func practice02() {

}

// Giải bài tập số 3
func practice03() {

}

// Giải bài tập số 4
func practice04() {

}
