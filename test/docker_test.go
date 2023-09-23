package test

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/containers"
	"testing"
)

func TestRun(t *testing.T) {
	docker := containers.Docker{}

	containerList, err := docker.ContainerList()
	if err != nil {
		panic(err)
	}

	fmt.Println(containerList)
}