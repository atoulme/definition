/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package schema

import "strconv"

const DefaultMode int64 = 0777

type InputFile struct {
	SourcePath      string `yaml:"source-path,omitempty" json:"path,omitempty"`
	DestinationPath string `yaml:"destination-path,omitempty" json:"destination-path,omitempty"`
	FileMode        string `yaml:"mode,omitempty" json:"mode,omitempty"`
}

// Mode gets the Unix mode bits
func (in InputFile) Mode() int64 {
	val, err := strconv.ParseInt(in.FileMode, 0, 64)
	if err != nil {
		return DefaultMode
	}
	return val
}

// GetSource makes it easy to change the name of the source member, as it is
// expected to change in the near future. (Also, there might be some logic behind it as well in
// the future)
func (in InputFile) Source() string {
	return in.SourcePath
}

// GetDestination makes it easy to change the name of the source member, as it is
// expected to change in the near future. (Also, there might be some logic behind it as well in
// the future)
func (in InputFile) Destination() string {
	return in.DestinationPath
}
