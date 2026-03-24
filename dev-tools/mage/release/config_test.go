// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package release

import (
	"os"
	"testing"
)

func TestInferLatestRelease(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
		wantErr bool
	}{
		{
			name:    "valid version",
			version: "9.2.1",
			want:    "9.2.0",
			wantErr: false,
		},
		{
			name:    "patch version greater than 1",
			version: "8.15.10",
			want:    "8.15.9",
			wantErr: false,
		},
		{
			name:    "patch version is 0",
			version: "9.2.0",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid format - missing patch",
			version: "9.2",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid format - non-numeric patch",
			version: "9.2.x",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := inferLatestRelease(tt.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("inferLatestRelease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inferLatestRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInferNextRelease(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
		wantErr bool
	}{
		{
			name:    "valid version",
			version: "9.2.0",
			want:    "9.2.1",
			wantErr: false,
		},
		{
			name:    "patch version greater than 0",
			version: "8.15.9",
			want:    "8.15.10",
			wantErr: false,
		},
		{
			name:    "invalid format - missing patch",
			version: "9.2",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid format - non-numeric patch",
			version: "9.2.x",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := inferNextRelease(tt.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("inferNextRelease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inferNextRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInferReleaseBranch(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
	}{
		{
			name:    "valid version",
			version: "9.2.1",
			want:    "9.2",
		},
		{
			name:    "double digit minor",
			version: "8.15.10",
			want:    "8.15",
		},
		{
			name:    "missing patch",
			version: "9.2",
			want:    "9.2",
		},
		{
			name:    "single part version",
			version: "9",
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferReleaseBranch(tt.version)
			if got != tt.want {
				t.Errorf("inferReleaseBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadConfigForMajorMinorRelease(t *testing.T) {
	// Save and restore env
	oldCurrent := os.Getenv("CURRENT_RELEASE")
	defer func() {
		if oldCurrent != "" {
			os.Setenv("CURRENT_RELEASE", oldCurrent)
		} else {
			os.Unsetenv("CURRENT_RELEASE")
		}
	}()

	// Test with 9.3.0 (major/minor release)
	os.Setenv("CURRENT_RELEASE", "9.3.0")

	cfg, err := LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("LoadConfigFromEnv failed: %v", err)
	}

	if cfg.CurrentRelease != "9.3.0" {
		t.Errorf("CurrentRelease = %s, want 9.3.0", cfg.CurrentRelease)
	}

	// For .0 releases, LatestRelease should be empty (can't be inferred)
	if cfg.LatestRelease != "" {
		t.Errorf("LatestRelease should be empty for .0 release, got %s", cfg.LatestRelease)
	}

	if cfg.NextRelease != "9.3.1" {
		t.Errorf("NextRelease = %s, want 9.3.1", cfg.NextRelease)
	}

	if cfg.ReleaseBranch != "9.3" {
		t.Errorf("ReleaseBranch = %s, want 9.3", cfg.ReleaseBranch)
	}
}
