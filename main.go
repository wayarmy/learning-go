package main

import (
	"fmt"

	"github.com/wayarmy/learning-go/day1"
	"github.com/wayarmy/learning-go/day2"
	u "github.com/wayarmy/learning-go/helper"
)

const current = 2

func main() {
	day, err := u.ReadNumberFromKeyboard("Muốn tìm kiếm bài tập về nhà của ngày thứ : ")
	u.CheckErr(err)

	switch {
	case int(day) == 1:
		day1.DoHomeWork()
	case int(day) == 2:
		day2.DoHomeWork()
	default:
		fmt.Printf("Chỉ mới học đến ngày thứ %d, làm gì đã có bài tập ngày %g, mời chọn lại.! \n", current, day)
		fmt.Println("")
		main()
	}
}
