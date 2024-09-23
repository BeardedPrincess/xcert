// Copyright (c) BeardedPrincess
// SPDX-License-Identifier: MPL-2.0

package xcert

var (
	versionBuildTimeStamp string
	versionString         string
)

// GetFormattedVersionString gets a friendly printable string to represent the version
func GetFormattedVersionString() string {
	if versionString == "" {
		versionString = "Unknown"
	}
	return versionString
}

func GetFormatedBuildTimeStamp() string {
	if versionBuildTimeStamp == "" {
		versionBuildTimeStamp = "Unknown"
	}
	return versionBuildTimeStamp
}
