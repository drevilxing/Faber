package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func RunContainer(imageName string, containerName string, srcPort string, dstPort string, volume map[string]string) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	m := make([]mount.Mount, 0, len(volume))
	for k, v := range volume {
		m = append(m, mount.Mount{Type: mount.TypeBind, Source: k, Target: v})
	}
	// m = append(m, mount.Mount{Type: mount.TypeBind, Source: "/etc/prometheus/prometheus.yml", Target: "/etc/prometheus/prometheus.yml"})

	exports := make(nat.PortSet)
	netport := make(nat.PortMap)
	port, err := nat.NewPort("tcp", srcPort)
	if err != nil {
		panic(err)
	}
	exports[port] = struct{}{}
	portlist := make([]nat.PortBinding, 0, 1)
	portlist = append(portlist, nat.PortBinding{HostIP: "0.0.0.0", HostPort: dstPort})
	netport[port] = portlist

	// imageName := "prom/prometheus"

	containerID := GetDockerContainerID(containerName)
	if containerID == "" {
		out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}
		defer out.Close()
		io.Copy(os.Stdout, out)

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image:        imageName,
			ExposedPorts: exports,
		}, &container.HostConfig{
			PortBindings: netport,
			Mounts:       m,
		}, nil, nil, containerName)
		if err != nil {
			panic(err)
		}
		containerID = resp.ID
	}
	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(containerID)
}

func StopContainer(name string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerID := GetDockerContainerID(name)
	timeout := time.Second * 10

	err = cli.ContainerStop(ctx, containerID, &timeout)
	if err != nil {
		panic(err)
	}
}

func RemoveContainer(name string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerID := GetDockerContainerID(name)

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
	})
	if err != nil {
		panic(err)
	}
}

func GetDockerContainerID(name string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Names[0] == name {
			fmt.Println(container.ID)
			return container.ID
		}
	}
	return ""
}

func ConnectToNecwork(containerName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerID := GetDockerContainerID(containerName)
	networkID := GetNetworkOfFabric()

	err = cli.NetworkConnect(ctx, networkID, containerID, &network.EndpointSettings{})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetNetworkOfFabric() string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	ID := GetDockerContainerID("/peer0.org1.example.com")
	CJSON, err := cli.ContainerInspect(ctx, ID)
	fmt.Println(CJSON.NetworkSettings.Networks)
	for key, value := range CJSON.NetworkSettings.Networks {
		fmt.Println(key, value)
		fmt.Println(value.NetworkID)
		return value.NetworkID
	}
	return ""
}
