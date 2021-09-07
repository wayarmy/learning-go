package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Đọc dữ liệu từ bàn phím, và trả về dữ liệu float64
func ReadNumberFromKeyboard(msg string) (float64, error) {
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
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("[Error]: %s", err)
		os.Exit(1)
	}
}
