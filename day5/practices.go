package day5

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/wayarmy/learning-go/helper"
)

const (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// Commands là hàm khai báo các subcommands cho câu lệnh được khai báo ở main.go bên ngoài.
func Commands() []*cli.Command {
	return []*cli.Command{
		commandDocker(),
		{
			Name:    "employee",
			Aliases: []string{"e"},
			Usage:   "giải bài tập thứ 2 ngày 5",
			Action: func(c *cli.Context) error {
				practice2()
				return nil
			},
		},
	}
}

// Định nghĩa command docker
func commandDocker() *cli.Command {
	return &cli.Command{
		// Usage:
		// wayarmy day3 docker
		Name:  "docker",
		Usage: "Câu lệnh dùng cho Docker",
		Subcommands: []*cli.Command{
			commandListContainer(),
			commandStopContainer(),
			commandStartContainer(),
		},
	}
}

// Định nghĩa câu lệnh start container
func commandStartContainer() *cli.Command {
	return &cli.Command{
		Name:    "start-container",
		Aliases: []string{"start"},
		Usage:   "Khởi động lại container đã tạm dừng",
		Flags: []cli.Flag{
			// Usage: wayarmy day3 docker start --id xxxx
			&cli.StringFlag{
				Name:    "container-id",
				Aliases: []string{"id"},
				Usage:   "Cung cấp container_id",
			},
			// Usage: wayarmy day3 docker start --n container_name
			&cli.StringFlag{
				Name:    "container-name",
				Aliases: []string{"n"},
				Usage:   "Cung cấp container_name",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("container-id") == "" && c.String("container-name") == "" {
				return errors.New("vui lòng cung cấp container-id hoặc container-name bạn đang cần start!.")
			}
			d, err := NewDockerClient()
			id := c.String("id")
			if c.String("container-id") == "" {
				id, err = d.DockerClient.convertContainerNameToId(c.String("container-name"))
				if err != nil {
					return nil
				}
			}

			err = d.DockerClient.startContainerWithId(id)
			if err != nil {
				return err
			}
			fmt.Printf("Container %s%s khởi động thành công!.", c.String("container-id"), c.String("container-name"))
			return nil
		},
	}
}

// Định nghĩa câu lệnh lấy danh sách containers
func commandListContainer() *cli.Command {
	return &cli.Command{
		Name:  "ls",
		Usage: "List các containers trên local",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "state",
				Aliases: []string{"s"},
				Usage:   "Filter theo status của các container: running/exit/all",
			},
		},
		Action: func(c *cli.Context) error {
			var state string
			if c.String("state") == "" {
				state = "all"
			} else {
				state = c.String("state")
			}

			err := PrintListContainers(state)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

// Định nghĩa câu lệnh stop container
func commandStopContainer() *cli.Command {
	return &cli.Command{
		Name:    "stop-container",
		Aliases: []string{"stop"},
		Usage:   "Tạm dừng container đang hoạt động",
		Flags: []cli.Flag{
			// Usage: wayarmy day3 docker stop --id xxxx
			&cli.StringFlag{
				Name:    "container-id",
				Aliases: []string{"id"},
				Usage:   "Cung cấp container_id",
			},
			// Usage: wayarmy day3 docker stop --n container_name
			&cli.StringFlag{
				Name:    "container-name",
				Aliases: []string{"n"},
				Usage:   "Cung cấp container_name",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("container-id") == "" && c.String("container-name") == "" {
				return errors.New("vui lòng cung cấp container-id hoặc container-name bạn đang cần stop!.")
			}

			d, err := NewDockerClient()
			id := c.String("id")
			if c.String("container-id") == "" {
				id, err = d.DockerClient.convertContainerNameToId(c.String("container-name"))
				if err != nil {
					return nil
				}
			}

			err = d.DockerClient.stopContainerWithId(id)
			if err != nil {
				return err
			}
			fmt.Printf("Dừng hoạt động của container %s%s thành công!.", c.String("container-id"), c.String("container-name"))
			return nil
		},
	}
}

// Giải bài 2 của day 5
func practice2() {
	fmt.Println(colorRed, "Giả sử đề bài cho 1 công ty có 2 danh sách nhân viên:", colorReset)
	fmt.Println(colorGreen, "Danh sách nhân viên chính thức: Salary = BasicPay + PF ", colorReset)
	perm := [][]string{}
	for _, e := range InitListPermanent() {
		perm = append(perm, []string{
			e.Name(), strconv.Itoa(e.EmpId()), strconv.Itoa(e.Salary()),
		})
	}
	permF := []string{"Name", "EmpID", "Salary"}
	helper.TablePrint(perm, nil, permF)

	fmt.Println(colorGreen, "Danh sách nhân viên hợp đồng: ", colorReset)
	cont := [][]string{}
	for _, e := range InitListContract() {
		cont = append(cont, []string{
			e.Name(), strconv.Itoa(e.EmpId()), strconv.Itoa(e.Salary()),
		})
	}
	contF := []string{"Name", "EmpID", "Salary"}
	helper.TablePrint(cont, nil, contF)

	fmt.Println("")
	t := TotalSalary(InitListPermanent()) + TotalSalary(InitListContract())
	fmt.Println(colorRed, "Tổng lương mà công ty trả cho nhân viên hàng tháng:", colorReset, colorGreen, t, "$")
}
