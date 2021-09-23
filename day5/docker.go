package day5

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/wayarmy/learning-go/helper"
)

// Container trả về dữ liệu của container
type Container struct {
	Id        string
	Names     []string
	ImageName string
	Ports     []types.Port
	Status    string
}

type Docker struct {
	DockerClient DockerClient
}

type DockerClient interface {
	stopContainerWithId(id string) error
	listContainersWithState(state string) ([]Container, error)
	startContainerWithId(id string) error
	convertContainerNameToId(name string) (string, error)
}

type DockerRepository struct {
	Client *client.Client
}

// Khởi tạo Docker client
func NewDockerClient() (*Docker, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	d := DockerRepository{
		Client: c,
	}
	return &Docker{
		DockerClient: d,
	}, nil
}

// listContainers trả về danh sách các container đang chạy trên local
// Filter danh sách containers theo status, nếu muốn trả về tất cả các containers thì truyền vào state=all
func (d DockerRepository) listContainersWithState(state string) ([]Container, error) {
	ctx := context.Background()

	t := types.ContainerListOptions{}
	if state == "all" {
		t.All = true
	} else {
		args := filters.NewArgs(filters.KeyValuePair{
			Key:   "label",
			Value: "state=" + state,
		})
		t.Filters = args
	}

	result, err := d.Client.ContainerList(ctx, t)
	if err != nil {
		return nil, err
	}
	var containers []Container
	for _, c := range result {
		containers = append(containers, Container{
			Id:        c.ID[:12],
			Names:     c.Names,
			ImageName: c.Image,
			Ports:     c.Ports,
			Status:    c.Status,
		})
	}
	return containers, nil
}

// Hiển thị danh sách Containers ra ngoài terminal
func PrintListContainers(state string) error {
	d, err := NewDockerClient()
	if err != nil {
		return err
	}

	containers, err := d.DockerClient.listContainersWithState(state)
	if err != nil {
		return err
	}
	data := [][]string{}
	// fmt.Printf("%v", containers)

	for _, c := range containers {
		data = append(data, []string{
			c.Id,
			strings.Trim(strings.Join(c.Names, ","), "/"),
			c.ImageName,
			c.Status,
			portsToS(c.Ports),
		})
	}

	header := []string{"ID", "Name", "Image Name", "Status", "Ports"}
	helper.TablePrint(data, header, nil)
	return nil
}

// convert list port sang dạng string
func portsToS(ports []types.Port) string {
	ps := []string{}
	for _, p := range ports {
		ps = append(ps, strconv.Itoa(int(p.PrivatePort))+":"+strconv.Itoa(int(p.PublicPort)))
	}
	return strings.Join(ps, ", ")
}

// khởi động lại container đã stop
func (d DockerRepository) startContainerWithId(id string) error {
	ctx := context.Background()

	opts := types.ContainerStartOptions{}

	err := d.Client.ContainerStart(ctx, id, opts)
	if err != nil {
		return err
	}
	return nil
}

// convert từ container name sang container id
func (d DockerRepository) convertContainerNameToId(name string) (string, error) {
	containers, err := d.listContainersWithState("all")
	if err != nil {
		return "", err
	}

	for _, c := range containers {
		for _, n := range c.Names {
			n = strings.Split(n, "/")[1]
			if n == name {
				return c.Id, nil
			}
		}
	}

	return "", errors.New("cannot find container with that name")
}

// tạm dừng hoạt động của 1 container với id đã có
func (d DockerRepository) stopContainerWithId(id string) error {
	ctx := context.Background()

	err := d.Client.ContainerStop(ctx, id, nil)
	if err != nil {
		return err
	}

	return nil
}
