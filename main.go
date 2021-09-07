package main

import (
	"fmt"

	"github.com/wayarmy/learning-go/day1"
	"github.com/wayarmy/learning-go/day2"
	u "github.com/wayarmy/learning-go/helper"
)

func main() {
	day, err := u.ReadNumberFromKeyboard("Muốn tìm kiếm bài tập về nhà của ngày thứ : ")
	u.CheckErr(err)

	switch {
	case int(day) == 1:
		day1.DoHomeWork()
	case int(day) == 2:
		day2.DoHomeWork()
	default:
		fmt.Println("Chỉ chấp nhận 1/2/3, mời chọn lại.!")
		main()
	}
}
