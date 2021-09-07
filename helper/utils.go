package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
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

// In ra màn hình dữ liệu dứoi dạng table
// Dữ liệu truyền vào phải đảm bảo đúng kiểu [][]string
func TablePrint(data [][]string, header, footer []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("╪")
	table.SetRowSeparator("-")

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
