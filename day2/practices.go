package day2

import (
	"fmt"
	"sort"
	"strconv"

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

type Employee struct {
	Name      string
	Salary    int
	Subsidize int
}

// Giải bài tập số 1
func practice01() {
	numbers := []int{2, 1, 3, 4, 7, 5, 10}
	fmt.Println("Cho trước một mảng các số: ", numbers)
	fmt.Println("Số lớn thứ nhì trong dãy số trên: ", max2Numbers(numbers))
}

// Giải bài tập số 2
func practice02() {
	strs := []string{"aa", "aab", "bcd", "a", "cdf", "bb"}
	fmt.Println("Cho trước một mảng bao gồm các chuỗi: ", strs)
	fmt.Println("Các chuỗi có độ dài lớn nhất trong mảng vừa rồi là: ", findMaxLengthElement(strs))
}

// Giải bài tập số 3
func practice03() {
	n := []int{1, 2, 5, 2, 6, 2, 5, 2, 6, 1, 5}
	fmt.Println("Cho trước một mảng bao gồm các số: ", n)
	fmt.Println("Mảng mới sau khi đã remove các phần tử bị duplicate: ", removeDuplicates(n))
}

// Giải bài tập số 4
func practice04() {
	e := []Employee{
		{"Walle", 3, 4500000},
		{"Agent", 1, 2500000},
		{"Adam", 2, 2000000},
		{"Alice", 2, 1000000},
		{"Gopher", 4, 3000000},
	}

	fmt.Println("Danh sách nhân viên của một công ty với các thông tin như sau: ")
	printEmployee(e)

	fmt.Println("")
	fmt.Println("Sắp xếp tên nhân viên tăng dần theo bảng chữ cái:")
	eOrderByName := orderEmployeeByName(e)
	printEmployee(eOrderByName)

	fmt.Println("")
	fmt.Println("Sắp xếp nhân viên theo mức lương giảm dần: ")
	eOrderBySalary := orderEmployeeBySalary(e)
	printEmployee(eOrderBySalary)

	fmt.Println("")
	fmt.Println("Nhân viên có mức lương cao thứ 2:")
	eMax2 := []Employee{top2EmployeeBySalary(e)}
	printEmployee(eMax2)
}

// in danh sách nhân viên ra dưới dạng bảng
func printEmployee(e []Employee) {
	data := [][]string{}
	for _, v := range e {
		data = append(data, []string{
			v.Name,
			strconv.Itoa(v.Salary),
			strconv.Itoa(v.Subsidize),
		})
	}
	u.TablePrint(data, []string{"Tên", "Hệ số lương", "Tiền trợ cấp"}, nil)
}

// Tìm số lớn thứ 2 trong 1 mảng các số bất kỳ
func max2Numbers(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-2]
}

// findMaxLengthElement sẽ trả ra các chuỗi có độ dài lớn nhất
func findMaxLengthElement(e []string) (output []string) {
	maxLength := len(e[0])

	for _, v := range e {
		switch {
		case len(v) < maxLength:
			continue
		case len(v) == maxLength:
			output = append(output, v)
		case len(v) > maxLength:
			maxLength = len(v)
			output = output[:0]
			output = append(output, v)
		}
	}

	return
}

// removeDuplicates sẽ remove hết những phần tử bị trùng trong 1 chuỗi các số
func removeDuplicates(n []int) (result []int) {
	keys := map[int]bool{}
	for _, entry := range n {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return
}

// Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
func orderEmployeeByName(e []Employee) []Employee {
	sort.SliceStable(e, func(i, j int) bool { return e[i].Name < e[j].Name })
	return e
}

// Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
func orderEmployeeBySalary(e []Employee) []Employee {
	sort.Slice(e, func(i, j int) bool {
		return (e[i].Salary*1500000 + e[i].Subsidize) > (e[j].Salary*1500000 + e[j].Subsidize)
	})
	return e
}

// Lấy ra nhân viên có mức lương lớn thứ 2 trong danh sách nhân viên
func top2EmployeeBySalary(e []Employee) Employee {
	top2 := orderEmployeeBySalary(e)[1]
	return top2
}
