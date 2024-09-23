// Copyright (c) BeardedPrincess 2024
// SPDX-License-Identifier: MPL-2.0

package xcert

type Config struct {
	CertIssuer string         `json:"certIssuer"`
	Config     map[string]any `json:"config"`
}
