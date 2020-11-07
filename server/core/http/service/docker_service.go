// Copyright 2020 liuxiaodong Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func NewDockerService() *DockerService {
	return &DockerService{}
}

type DockerService struct {
}

func (this *DockerService) ListContainers() ([]types.Container, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	resources, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	return resources, nil
}

func (this *DockerService) ListImages() ([]types.ImageSummary, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	resources, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	return resources, nil
}

func (this *DockerService) ListNets() ([]types.NetworkResource, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	resources, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	return resources, nil
}

func (this *DockerService) ListVolumes() (volume.VolumesListOKBody, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return volume.VolumesListOKBody{}, err
	}

	resources, err := cli.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		panic(err)
	}

	return resources, nil
}
