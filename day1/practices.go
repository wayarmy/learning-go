package day1

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"

	u "github.com/wayarmy/learning-go/helper"
)

func DoHomeWork() {
	str, err := u.ReadNumberFromKeyboard("Chọn bài tập muốn giải (1, 2, 3): ")
	u.CheckErr(err)
	fmt.Println("")

	switch {
	case str == 1:
		practice01()
	case str == 2:
		practice02()
	case str == 3:
		practice03()
	default:
		fmt.Println("Chỉ chấp nhận 1/2/3, mời chọn lại.!")
	}
}

// Giải bài tập 01
func practice01() {
	a, err := u.ReadNumberFromKeyboard("Giá trị của a: ")
	u.CheckErr(err)

	b, err := u.ReadNumberFromKeyboard("Giá trị của b: ")
	u.CheckErr(err)

	c, err := u.ReadNumberFromKeyboard("Giá trị của c: ")
	u.CheckErr(err)

	_, _ = quadraticEquation2(a, b, c)
}

// Giải bài tập 02
func practice02() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	i := 1
	loop := true
	for loop {
		n, err := u.ReadNumberFromKeyboard("Nhập số bạn đang đoán vào đây: ")
		u.CheckErr(err)
		switch {
		case int(n) < x:
			fmt.Println("Số bạn chọn nhỏ hơn số X mà hệ thống sinh ra. mời chọn lại!.")
			i++
			continue
		case int(n) > x:
			fmt.Println("Số bạn chọn lớn hơn số X mà hệ thống sinh ra. mời chọn lại!.")
			i++
			continue
		case int(n) == x:
			fmt.Printf("Congras, bạn đã chọn đúng số mà hệ thống đã chọn sau %v lần chọn", i)
			loop = false
		}
	}
}

// Giải bài tập 03
func practice03() {
	n, err := u.ReadNumberFromKeyboard("Nhập vào số N bất kỳ mà bạn chọn < 100.000: ")
	u.CheckErr(err)

	if int(n) > 100000 {
		fmt.Println("Số bạn nhập lớn hơn 100.000, không thoả mãn điều kiện, mời chạy lại chương trình!.")
		os.Exit(1)
	}
	fmt.Println("Tập hợp các số nguyên tố nhỏ hơn số bạn chọn:")
	for i := 2; i <= int(n); i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Printf("%d ", i)
		}
	}
}

/*
quadraticEquation2 giải phương trình bậc 2 với 3 tham số: a, b, c được truyền vào với kiểu dữ liệu float64
hàm này sẽ trả về kết quả
*/
func quadraticEquation2(a, b, c float64) (interface{}, interface{}) {
	delta := (b * b) - (4 * a * c)
	switch {
	case delta < 0:
		x1 := complex(-b/(2*a), math.Sqrt(-delta)/(2*a))
		x2 := complex(-b/(2*a), -math.Sqrt(-delta)/(2*a))
		fmt.Printf("Phương trình có 2 nghiệm ở dạng số phức: x1 = %v, x2 = %v", x1, x2)
		return x1, x2
	case delta == 0:
		fmt.Printf("Phương trình có 1 nghiệm duy nhất: %g", (-b / 2 * a))
		return (-b / 2 * a), (-b / 2 * a)
	case delta > 0:
		fmt.Printf("Phương trình có 2 nghiệm: %g, %g", (-b+math.Sqrt(delta))/(2*a), (-b-math.Sqrt(delta))/(2*a))
		return (-b + math.Sqrt(delta)/(2*a)), (-b - math.Sqrt(delta)/(2*a))
	}

	return nil, nil
}
