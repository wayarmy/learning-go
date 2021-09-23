package day5

import (
	"math/rand"
	"time"
)

type Permanent struct {
	name     string
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	name     string
	empId    int
	basicpay int
}

type Employee interface {
	Name() string
	EmpId() int
	Salary() int
}

// Tạo danh sách nhân viên permanent random
func InitListPermanent() (perm []Employee) {
	emp := []string{"Mark", "Tim", "Satya", "Sundar", "Jeff"}
	for _, p := range emp {
		e := Permanent{
			name:     p,
			empId:    getRandomInt(10000, 30000),
			basicpay: getRandomInt(1, 10) * 1000,
			pf:       getRandomInt(1, 5) * 500,
		}
		perm = append(perm, e)
	}
	return perm
}

// Trả về số random trong khoảng min->max
func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Tạo danh sách nhân viên permanent random
func InitListContract() (cont []Employee) {
	emp := []string{"Tom", "Wayne", "Shane", "Sheeran"}
	for _, p := range emp {
		e := Contract{
			name:     p,
			empId:    getRandomInt(10000, 30000),
			basicpay: getRandomInt(1, 10) * 1000,
		}
		cont = append(cont, e)
	}
	return cont
}

// Trả về employee ID permanent
func (p Permanent) EmpId() int {
	return p.empId
}

// Trả về employee name permanent
func (p Permanent) Name() string {
	return p.name
}

// Trả về salary của employee permanent
func (p Permanent) Salary() int {
	return p.basicpay + p.pf
}

// Trả về employee ID contract
func (c Contract) EmpId() int {
	return c.empId
}

// Trả về employee name permanent
func (c Contract) Name() string {
	return c.name
}

// Trả về salary của employee contract
func (c Contract) Salary() int {
	return c.basicpay
}

// Tính tổng tiền lương công ty phải trả cho nhân viên
func TotalSalary(emp []Employee) (t int) {
	for _, e := range emp {
		t = t + e.Salary()
	}
	return
}
