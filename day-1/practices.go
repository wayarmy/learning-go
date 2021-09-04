package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Chọn bài tập muốn giải (1, 2, 3): ")
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")

	switch {
	case str == "1":
		practice01()
	case str == "2":
		practice02()
	case str == "3":
		_ = practice03()
	default:
		fmt.Println("Chỉ chấp nhận 1/2/3, mời chọn lại.!")
	}
}

// Đọc dữ liệu từ bàn phím, và trả về dữ liệu float64
func readNumberFromKeyboard(msg string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	number, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// Kiểm tra xem dữ liệu truyền vào có phải error hay không ?
func checkErr(err error) {
	if err != nil {
		fmt.Printf("[Error]: %s", err)
	}
}

// Giải bài tập 01
func practice01() {
	a, err := readNumberFromKeyboard("Giá trị của a: ")
	checkErr(err)

	b, err := readNumberFromKeyboard("Giá trị của b: ")
	checkErr(err)

	c, err := readNumberFromKeyboard("Giá trị của c: ")
	checkErr(err)

	quadraticEquation2(a, b, c)
}

// Giải bài tập 02
func practice02() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	i := 1
	loop := true
	for loop {
		n, err := readNumberFromKeyboard("Nhập số bạn đang đoán vào đây: ")
		checkErr(err)
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
func practice03() []int {
	n, err := readNumberFromKeyboard("Nhập vào số N bất kỳ mà bạn chọn < 100.000: ")
	checkErr(err)

	if int(n) > 100000 {
		fmt.Println("Số bạn nhập lớn hơn 100.000, không thoả mãn điều kiện, mời chạy lại chương trình!.")
		os.Exit(1)
	}
	var primes []int
	fmt.Println("Tập hợp các số nguyên tố nhỏ hơn số bạn chọn:")
	for i := 2; i <= int(n); i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Printf("%d ", i)
			primes = append(primes, i)
		}
	}

	return primes
}

/*
quadraticEquation2 giải phương trình bậc 2 với 3 tham số: a, b, c được truyền vào với kiểu dữ liệu float64
hàm này sẽ trả về kết quả
*/
func quadraticEquation2(a, b, c float64) {
	delta := (b * b) - (4 * a * c)
	switch {
	case delta < 0:
		sqrt := math.Sqrt(-delta)
		// x1 := complex(-b/(2*a), math.Sqrt(-delta)/(2*a))
		// x2 := complex(-b/(2*a), -math.Sqrt(-delta)/(2*a))
		fmt.Printf("Phương trình có 2 nghiệm ở dạng số phức: (%g + %fi)/%g, (%g - %fi)/%g", -b, sqrt, (2 * a), -b, sqrt, (2 * a))
	case delta == 0:
		fmt.Printf("Phương trình có 1 nghiệm duy nhất: %g", (-b / 2 * a))
	case delta > 0:
		fmt.Printf("Phương trình có 2 nghiệm: %g, %g", (-b + math.Sqrt(delta)/(2*a)), (-b - math.Sqrt(delta)/(2*a)))
	}
}
