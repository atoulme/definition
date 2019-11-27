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

package process

import (
	"fmt"

	"github.com/whiteblock/definition/internal/entity"
	"github.com/whiteblock/definition/internal/parser"
	"github.com/whiteblock/definition/schema"
)

type State struct {
	systemState map[string]schema.SystemComponent
}

func NewState() State {
	return State{systemState: map[string]schema.SystemComponent{}}
}

//System is for diff calculations
type System interface {
	//Add modifies State
	Add(state *State, systems []schema.SystemComponent) ([]entity.Service, error)
	//Remove modifies state
	Remove(state *State, systems []string) ([]entity.Service, error)
}

type system struct {
	namer  parser.Names
	parser parser.Service
}

func NewSystem(namer parser.Names, parser parser.Service) System {
	return &system{namer: namer, parser: parser}
}

//Add modifies State
func (sys system) Add(state *State, systems []schema.SystemComponent) ([]entity.Service, error) {
	out := []entity.Service{}

	for _, system := range systems {
		name := sys.namer.SystemComponent(system)
		_, exists := state.systemState[name]
		if exists {
			return nil, fmt.Errorf("already have a system with the name \"%s\"", name)
		}
		services, err := sys.parser.FromSystem(system)
		if err != nil {
			return nil, err
		}
		out = append(out, services...)
	}

	for _, system := range systems {
		name := sys.namer.SystemComponent(system)
		state.systemState[name] = system
	}

	return out, nil
}

//Remove modifies state
func (sys system) Remove(state *State, systems []string) ([]entity.Service, error) {
	out := []entity.Service{}
	for _, toRemove := range systems {
		system, exists := state.systemState[toRemove]
		if !exists {
			return nil, fmt.Errorf("system not found")
		}
		services, err := sys.parser.FromSystem(system)
		if err != nil {
			return nil, err
		}
		out = append(out, services...)
	}
	for _, toRemove := range systems {
		delete(state.systemState, toRemove)
	}
	return out, nil
}
