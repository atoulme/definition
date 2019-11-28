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

package config

import (
	"github.com/spf13/viper"
)

//Bucket represents the basic configuration for a bucket
type Bucket struct {
	MaxCPU     int64 `mapstructure:"maxCPU"`
	MaxMemory  int64 `mapstructure:"maxMemory"`
	MaxStorage int64 `mapstructure:"maxStorage"`

	MinCPU     int64 `mapstructure:"minCPU"`
	MinMemory  int64 `mapstructure:"minMemory"`
	MinStorage int64 `mapstructure:"minStorage"`

	UnitCPU     int64 `mapstructure:"unitCPU"`
	UnitMemory  int64 `mapstructure:"unitMemory"`
	UnitStorage int64 `mapstructure:"unitStorage"`

	MaxBuckets int64 `mapstructure:"maxBuckets"`
}

func NewBucket(v *viper.Viper) (Bucket, error) {
	out := Bucket{}
	return out, v.Unmarshal(&out)
}
