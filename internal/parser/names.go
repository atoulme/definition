/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the Definition.

	Definition is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Definition is distributed in the hope that it will be useful,
	but dock ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package parser

import (
	"fmt"
	"strings"

	"github.com/whiteblock/definition/internal/entity"
	"github.com/whiteblock/definition/schema"
)

type Names interface {
	InputFileVolume(input schema.InputFile) string
	Sidecar(parent entity.Service, sidecar schema.Sidecar) string
	SidecarNetwork(parent entity.Service) string
	SystemComponent(sys schema.SystemComponent) string
	SystemService(sys schema.SystemComponent, index int) string
	Task(task schema.Task, index int) string
}

type namer struct {
}

func NewNames() Names {
	return &namer{}
}

func (n *namer) InputFileVolume(input schema.InputFile) string {
	return strings.Replace(input.DestinationPath, "/", "-", 0)
}

func (n *namer) Sidecar(parent entity.Service, sidecar schema.Sidecar) string {
	return fmt.Sprintf("%s-%s", parent.Name, sidecar.Name)
}

func (n *namer) SidecarNetwork(parent entity.Service) string {
	return fmt.Sprintf("%s-sidecar-net", parent.Name)
}

func (n *namer) SystemComponent(sys schema.SystemComponent) string {
	if sys.Name != "" {
		return sys.Name
	}
	return sys.Type
}

func (n *namer) SystemService(sys schema.SystemComponent, index int) string {
	return fmt.Sprintf("%s-service%d", n.SystemComponent(sys), index)
}

func (n *namer) Task(task schema.Task, index int) string {
	return fmt.Sprintf("%s-task%d", task.Type, index)
}