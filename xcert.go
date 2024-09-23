// Copyright (c) BeardedPrincess 2024
// SPDX-License-Identifier: MPL-2.0

package xcert

import (
	"fmt"
)

var (
	VersionBuildTimeStamp string
	VersionString         string
	Commit                string
)

// Set to any function with a Printf-like signature to enable DEBUG logging
var LogDebug func(string, ...interface{}) = func(string, ...any) {}

// Set to any function with a Printf-like signature to enable WARNING logging
var LogWarn func(string, ...interface{}) = func(string, ...any) {}

// GetFormattedVersionString gets a friendly printable string to represent the version
func GetFormattedVersionString() string {
	return fmt.Sprintf("%v [%v] (%v)", VersionString, Commit, VersionBuildTimeStamp)
}
