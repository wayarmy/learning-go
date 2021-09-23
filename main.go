package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
	"github.com/wayarmy/learning-go/day3"
	"github.com/wayarmy/learning-go/day5"
)

const current = 2

// // This function would work with day1 and day2. From day 3, i will use cli for real cli
// func main() {
// 	day, err := u.ReadNumberFromKeyboard("Muốn tìm kiếm bài tập về nhà của ngày thứ : ")
// 	u.CheckErr(err)

// 	switch {
// 	case int(day) == 1:
// 		day1.DoHomeWork()
// 	case int(day) == 2:
// 		day2.DoHomeWork()
// 	default:
// 		fmt.Printf("Chỉ mới học đến ngày thứ %d, làm gì đã có bài tập ngày %g, mời chọn lại.! \n", current, day)
// 		fmt.Println("")
// 		main()
// 	}
// }

func main() {
	app := &cli.App{
		Name:  "wayarmy",
		Usage: "Golang rules the world!.",
		Commands: []*cli.Command{
			{
				Name:        "day3",
				Aliases:     []string{"3"},
				Usage:       "các hàm tiện ích để tương tác với Elastic Search",
				Subcommands: day3.Commands(),
			},
			{
				Name:        "day5",
				Aliases:     []string{"3"},
				Usage:       "các hàm tiện ích để tương tác với Elastic Search",
				Subcommands: day5.Commands(),
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
