package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	eMock = []Employee{
		{"Walle", 3, 4500000},
		{"Agent", 1, 2500000},
		{"Adam", 2, 2000000},
		{"Alice", 2, 1000000},
		{"Gopher", 4, 3000000},
	}
)

func Test_Max2Numbers(t *testing.T) {
	arrayMock := []int{2, 1, 3, 4, 7, 5, 10}
	n := max2Numbers(arrayMock)

	assert.Equal(t, 7, n)
}

func Test_FindMaxLengthElement(t *testing.T) {
	strsMock := []string{"aa", "aab", "bcd", "a", "cdf", "bb"}
	expected := []string{"aab", "bcd", "cdf"}
	acctual := findMaxLengthElement(strsMock)

	assert.Equal(t, expected, acctual)
}

func Test_RemoveDuplicates(t *testing.T) {
	numbersMock := []int{1, 2, 5, 2, 6, 2, 5, 2, 6, 1, 5}
	expected := []int{1, 2, 5, 6}

	actual := removeDuplicates(numbersMock)
	assert.Equal(t, expected, actual)
}

func Test_OrderEmployeeByName(t *testing.T) {
	expected := []Employee{
		{"Adam", 2, 2000000},
		{"Agent", 1, 2500000},
		{"Alice", 2, 1000000},
		{"Gopher", 4, 3000000},
		{"Walle", 3, 4500000},
	}

	actual := orderEmployeeByName(eMock)
	assert.Equal(t, expected, actual)
}

func Test_OrderEmployeeBySalary(t *testing.T) {
	expected := []Employee{
		{"Gopher", 4, 3000000},
		{"Walle", 3, 4500000},
		{"Adam", 2, 2000000},
		{"Agent", 1, 2500000},
		{"Alice", 2, 1000000},
	}

	actual := orderEmployeeBySalary(eMock)
	assert.Equal(t, expected, actual)
}

func Test_Top2EmployeeBySalary(t *testing.T) {
	expected := Employee{"Walle", 3, 4500000}

	actual := top2EmployeeBySalary(eMock)
	assert.Equal(t, expected, actual)
}
