package day3

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

var (
	cal = [][]string{
		{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
	}
)

// showCalendar sẽ hiển thị ra calendar tương tự như lệnh cal trên máy tính
func showCalendar() error {
	t := time.Now()
	m := time.Now().Month()
	y := time.Now().Year()

	cal := buildCalendar(theFirstDayOfThisMonth(t), theLastDayOfThisMonth(t))
	header := [][]string{{m.String(), strconv.Itoa(y)}}
	tablePrint(header, nil, nil)
	tablePrint(cal, nil, nil)

	return nil
}

// buildCalendar sẽ xây dựng lên calendar dưới dạng table theo từng tuần [][]string
// dựa trên ngày đầu tiên của tháng và ngày cuối cùng của tháng đó
func buildCalendar(dStart, dEnd time.Time) [][]string {
	cal = [][]string{
		{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
	}
	w := []string{"", "", "", "", "", "", ""}
	for i := 0; i < 32; i++ {
		switch dStart.Weekday().String() {
		case "Sunday":
			w[0] = strconv.Itoa(dStart.Day())
		case "Monday":
			w[1] = strconv.Itoa(dStart.Day())
		case "Tuesday":
			w[2] = strconv.Itoa(dStart.Day())
		case "Wednesday":
			w[3] = strconv.Itoa(dStart.Day())
		case "Thursday":
			w[4] = strconv.Itoa(dStart.Day())
		case "Friday":
			w[5] = strconv.Itoa(dStart.Day())
		case "Saturday":
			w[6] = strconv.Itoa(dStart.Day())
			cal = append(cal, w)
			w = []string{"", "", "", "", "", "", ""}
		}

		if dStart.Day() == dEnd.Day() {
			cal = append(cal, w)
			break
		}

		dStart = dStart.Add(time.Hour * 24)
	}
	return cal
}

// theFirstDayOfThisMonth tìm và trả về ngày đầu tiên của tháng
func theFirstDayOfThisMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	loc := t.Location()
	return time.Date(y, m, 1, 0, 0, 0, 0, loc)
}

// theLastDayOfThisMonth tìm và trả về ngày cuối cùng của tháng
func theLastDayOfThisMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	loc := t.Location()
	return time.Date(y, m+1, 1, 0, 0, 0, -1, loc)
}

// tablePrint sử dụng tablewriter để hiển thị data [][]string dưới dạng bảng trên môi trường terminal
func tablePrint(data [][]string, header, footer []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	// Change table lines
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	if footer != nil {
		table.SetFooter(footer)
	}
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
