package day3

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Commands là hàm khai báo các subcommands cho câu lệnh được khai báo ở main.go bên ngoài.
func Commands() []*cli.Command {
	return []*cli.Command{
		commandCal(),
		commandDocker(),
		commandTree(),
	}
}

func commandCal() *cli.Command {
	return &cli.Command{
		// Usage:
		// - wayarmy day3 calendar
		// - wayarmy day3 cal
		Name:    "calendar",
		Aliases: []string{"cal"},
		Usage:   "Hiển thị lịch trên terminal",
		Action: func(c *cli.Context) error {
			err := showCalendar()
			if err != nil {
				return err
			}
			return nil
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

// Định nghĩa command tree
func commandTree() *cli.Command {
	return &cli.Command{
		// Usage:
		// wayarmy tree -L
		Name:  "tree",
		Usage: "Câu lệnh hiển thị file và thư mục ở dạng tree",
		Flags: []cli.Flag{
			// Usage: wayarmy day3 tree -L 2
			&cli.StringFlag{
				Name:    "level",
				Aliases: []string{"L"},
				Usage:   "Cấp độ folder cần list",
			},
		},
		Action: func(c *cli.Context) error {
			depth := 1
			if c.Int("level") != 0 {
				depth = c.Int("level")
			}
			currentPath, err := os.Getwd()
			if err != nil {
				return err
			}
			err = printDirectory(currentPath, depth)
			if err != nil {
				return err
			}
			return nil
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
				id, err = d.convertContainerNameToId(c.String("container-name"))
				if err != nil {
					return nil
				}
			}

			err = d.startContainerWithId(id)
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
				id, err = d.convertContainerNameToId(c.String("container-name"))
				if err != nil {
					return nil
				}
			}

			err = d.stopContainerWithId(id)
			if err != nil {
				return err
			}
			fmt.Printf("Dừng hoạt động của container %s%s thành công!.", c.String("container-id"), c.String("container-name"))
			return nil
		},
	}
}
