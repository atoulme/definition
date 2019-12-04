/*
	Copyright 2019 Whiteblock Inc.
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

package command

import (
	"time"
)

//StartContainer is the command for starting a container
type StartContainer struct {
	Name   string `json:"name"`
	Attach bool   `json:"attach"`
	// Timeout is the maximum amount of time to wait for the task before terminating it.
	// This is ignored if attach is false
	Timeout time.Duration `json:"timeout"`
}
