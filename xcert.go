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
var logDebug func(string, ...interface{}) = func(string, ...any) {}

// Set to any function with a Printf-like signature to enable WARNING logging
var logWarn func(string, ...interface{}) = func(string, ...any) {}

func SetLogger(d func(string, ...any), w func(string, ...any)) {
	if d != nil {
		logDebug = d
	}
	if w != nil {
		logWarn = w
	}
}

func LogWarn(s string, a ...any) {
	ps := fmt.Sprintf("[WARN][%s]%s", strPackageName, s)
	logWarn(ps, a...)
}

func LogDebug(s string, a ...any) {
	ps := fmt.Sprintf("[DEBUG][%s]%s", strPackageName, s)
	logDebug(ps, a...)
}

// GetFormattedVersionString gets a friendly printable string to represent the version
func GetFormattedVersionString() string {
	return fmt.Sprintf("%v [%v] (%v)", VersionString, Commit, VersionBuildTimeStamp)
}
