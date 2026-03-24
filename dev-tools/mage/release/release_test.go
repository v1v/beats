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
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdateVersion(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()

	// Create libbeat/version directory
	versionDir := filepath.Join(tmpDir, "libbeat", "version")
	err := os.MkdirAll(versionDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create version dir: %v", err)
	}

	// Create version.go file
	versionFile := filepath.Join(versionDir, "version.go")
	initialContent := `package version

const defaultBeatVersion = "9.3.0"
`
	err = os.WriteFile(versionFile, []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create version file: %v", err)
	}

	// Change to temp directory
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Test updating version
	err = UpdateVersion("9.4.0")
	if err != nil {
		t.Fatalf("UpdateVersion failed: %v", err)
	}

	// Verify the file was updated
	content, err := os.ReadFile(versionFile)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	if !strings.Contains(string(content), `const defaultBeatVersion = "9.4.0"`) {
		t.Errorf("Version not updated correctly. Got:\n%s", string(content))
	}
}

func TestUpdateDocs(t *testing.T) {
	tmpDir := t.TempDir()

	// Create necessary directories
	dirs := []string{
		"libbeat/docs",
		"deploy/kubernetes",
	}
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(tmpDir, dir), 0755)
		if err != nil {
			t.Fatalf("Failed to create dir %s: %v", dir, err)
		}
	}

	// Create version.asciidoc
	versionAsciidoc := filepath.Join(tmpDir, "libbeat/docs/version.asciidoc")
	versionContent := `:stack-version: 9.3.0
:doc-branch: 9.3
`
	err := os.WriteFile(versionAsciidoc, []byte(versionContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create version.asciidoc: %v", err)
	}

	// Create K8s manifest with wolfi suffix
	k8sFile := filepath.Join(tmpDir, "deploy/kubernetes/metricbeat-kubernetes.yaml")
	k8sContent := `image: docker.elastic.co/beats/metricbeat-wolfi:9.3.0
`
	err = os.WriteFile(k8sFile, []byte(k8sContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create K8s file: %v", err)
	}

	// Create filebeat K8s manifest without suffix
	filebeat8sFile := filepath.Join(tmpDir, "deploy/kubernetes/filebeat-kubernetes.yaml")
	filebeat8sContent := `image: docker.elastic.co/beats/filebeat:9.3.0
`
	err = os.WriteFile(filebeat8sFile, []byte(filebeat8sContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create filebeat K8s file: %v", err)
	}

	// Change to temp directory
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Test updating docs
	err = UpdateDocs("9.4.0")
	if err != nil {
		t.Fatalf("UpdateDocs failed: %v", err)
	}

	// Verify version.asciidoc was updated
	content, _ := os.ReadFile(versionAsciidoc)
	if !strings.Contains(string(content), ":stack-version: 9.4.0") {
		t.Errorf("version.asciidoc stack-version not updated. Got:\n%s", string(content))
	}
	if !strings.Contains(string(content), ":doc-branch: 9.4") {
		t.Errorf("version.asciidoc doc-branch not updated. Got:\n%s", string(content))
	}

	// Verify K8s file with -wolfi suffix was updated correctly
	content, _ = os.ReadFile(k8sFile)
	if !strings.Contains(string(content), "docker.elastic.co/beats/metricbeat-wolfi:9.4.0") {
		t.Errorf("K8s file with -wolfi not updated. Got:\n%s", string(content))
	}

	// Verify K8s file without suffix was updated correctly
	content, _ = os.ReadFile(filebeat8sFile)
	if !strings.Contains(string(content), "docker.elastic.co/beats/filebeat:9.4.0") {
		t.Errorf("K8s file without suffix not updated. Got:\n%s", string(content))
	}
}

func TestUpdateTestEnv(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test env directory
	testEnvDir := filepath.Join(tmpDir, "testing/environments")
	err := os.MkdirAll(testEnvDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test env dir: %v", err)
	}

	// Create latest.yml with multiple versions in the same minor line
	latestYml := filepath.Join(testEnvDir, "latest.yml")
	latestContent := `services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:9.3.1
  kibana:
    image: docker.elastic.co/kibana/kibana:9.3.2
  logstash:
    image: docker.elastic.co/logstash/logstash:9.3.0
`
	err = os.WriteFile(latestYml, []byte(latestContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create latest.yml: %v", err)
	}

	// Create metricbeat docker-compose with variable defaults
	metricbeatDir := filepath.Join(tmpDir, "metricbeat")
	err = os.MkdirAll(metricbeatDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create metricbeat dir: %v", err)
	}
	metricbeatCompose := filepath.Join(metricbeatDir, "docker-compose.yml")
	metricbeatContent := `services:
  elasticsearch:
    image: docker.elastic.co/integrations-ci/beats-elasticsearch:${ELASTICSEARCH_VERSION:-9.3.1}-1
    build:
      args:
        ELASTICSEARCH_VERSION: ${ELASTICSEARCH_VERSION:-9.3.1}
  kibana:
    image: docker.elastic.co/integrations-ci/beats-kibana:${KIBANA_VERSION:-9.3.0}-1
`
	err = os.WriteFile(metricbeatCompose, []byte(metricbeatContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create metricbeat docker-compose: %v", err)
	}

	// Change to temp directory
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Test updating test env
	// When releasing 9.3.4, we want test envs to use 9.3.3 (latest stable)
	// UpdateTestEnv(latestVersion, currentVersion)
	// - Matches all 9.x.x versions (based on major version from currentVersion)
	// - Replaces with latestVersion (9.3.3)
	err = UpdateTestEnv("9.3.3", "9.3.4")
	if err != nil {
		t.Fatalf("UpdateTestEnv failed: %v", err)
	}

	// Verify all 9.3.x versions in latest.yml were updated to 9.3.3 (not 9.3.4)
	content, _ := os.ReadFile(latestYml)
	contentStr := string(content)
	if !strings.Contains(contentStr, "docker.elastic.co/elasticsearch/elasticsearch:9.3.3") {
		t.Errorf("elasticsearch not updated. Got:\n%s", contentStr)
	}
	if !strings.Contains(contentStr, "docker.elastic.co/kibana/kibana:9.3.3") {
		t.Errorf("kibana not updated. Got:\n%s", contentStr)
	}
	if !strings.Contains(contentStr, "docker.elastic.co/logstash/logstash:9.3.3") {
		t.Errorf("logstash not updated. Got:\n%s", contentStr)
	}
	// Verify old versions are gone
	if strings.Contains(contentStr, ":9.3.0") || strings.Contains(contentStr, ":9.3.1") || strings.Contains(contentStr, ":9.3.2") {
		t.Errorf("Old versions still present in latest.yml. Got:\n%s", contentStr)
	}

	// Verify metricbeat docker-compose variable defaults were updated
	content, _ = os.ReadFile(metricbeatCompose)
	contentStr = string(content)
	if !strings.Contains(contentStr, "${ELASTICSEARCH_VERSION:-9.3.3}") {
		t.Errorf("ELASTICSEARCH_VERSION default not updated. Got:\n%s", contentStr)
	}
	if !strings.Contains(contentStr, "${KIBANA_VERSION:-9.3.3}") {
		t.Errorf("KIBANA_VERSION default not updated. Got:\n%s", contentStr)
	}
	// Verify old variable defaults are gone
	if strings.Contains(contentStr, ":-9.3.0}") || strings.Contains(contentStr, ":-9.3.1}") {
		t.Errorf("Old variable defaults still present in metricbeat compose. Got:\n%s", contentStr)
	}
}

func TestCheckRequirements(t *testing.T) {
	tests := []struct {
		name        string
		version     string
		shouldError bool
		errorMsg    string
	}{
		{
			name:        "6.x minor release blocked",
			version:     "6.5.0",
			shouldError: true,
			errorMsg:    "deprecated and blocked",
		},
		{
			name:        "7.x minor release blocked",
			version:     "7.5.0",
			shouldError: true,
			errorMsg:    "deprecated and blocked",
		},
		{
			name:        "8.x minor release blocked",
			version:     "8.5.0",
			shouldError: true,
			errorMsg:    "deprecated and blocked",
		},
		{
			name:        "9.x minor release allowed",
			version:     "9.3.0",
			shouldError: false,
		},
		{
			name:        "6.x patch release allowed",
			version:     "6.5.1",
			shouldError: false,
		},
		{
			name:        "7.x patch release allowed",
			version:     "7.5.1",
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary git repository for testing
			tmpDir := t.TempDir()

			cfg := &ReleaseConfig{
				CurrentRelease: tt.version,
				BaseBranch:     "main",
			}

			// Change to temp directory
			origDir, _ := os.Getwd()
			defer os.Chdir(origDir)
			os.Chdir(tmpDir)

			// Initialize git repo (this will fail if git is not available, but that's ok for this test)
			// We're mainly testing the version validation logic
			err := checkRequirements(cfg)

			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error for version %s, got nil", tt.version)
				} else if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', got: %v", tt.errorMsg, err)
				}
			} else if err != nil && !strings.Contains(err.Error(), "repository") {
				// Allow errors about repository not existing, but not version validation errors
				if strings.Contains(err.Error(), "deprecated") {
					t.Errorf("Unexpected error for version %s: %v", tt.version, err)
				}
			}
		})
	}
}

func TestUpdateTestEnvNoChanges(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test env directory
	testEnvDir := filepath.Join(tmpDir, "testing/environments")
	err := os.MkdirAll(testEnvDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test env dir: %v", err)
	}

	// Create latest.yml with a version from a different major line
	latestYml := filepath.Join(testEnvDir, "latest.yml")
	latestContent := `services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.15.5
`
	err = os.WriteFile(latestYml, []byte(latestContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create latest.yml: %v", err)
	}

	// Change to temp directory
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Test updating test env when releasing 9.3.4
	// UpdateTestEnv("9.3.3", "9.3.4") looks for 9.x.x versions to update to 9.3.3
	// File has 8.15.5 (different major version), so pattern 9.\d+.\d+ won't match
	// This should complete without error even though no changes are made
	err = UpdateTestEnv("9.3.3", "9.3.4")
	if err != nil {
		t.Fatalf("UpdateTestEnv should not fail when no matches found: %v", err)
	}

	// Verify file was NOT updated (version should still be 8.15.5)
	content, _ := os.ReadFile(latestYml)
	if !strings.Contains(string(content), "docker.elastic.co/elasticsearch/elasticsearch:8.15.5") {
		t.Errorf("Test env file should not be updated when version doesn't match major. Got:\n%s", string(content))
	}
}
