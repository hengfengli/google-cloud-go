// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	instancepb "google.golang.org/genproto/googleapis/spanner/admin/instance/v1"
	"google.golang.org/grpc/codes"
)

type sampleFunc func(w io.Writer, dbName string) error
type sampleFuncWithContext func(ctx context.Context, w io.Writer, dbName string) error

var (
	validInstancePattern = regexp.MustCompile("^projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)$")
)

func initTest(t *testing.T, id string) (instName, dbName string, cleanup func()) {
	//projectID := getSampleProjectId(t)
	instName = os.Getenv("GOLANG_SAMPLES_SPANNER")
	cleanup = func() {}
	// instName, cleanup = createTestInstance(t, projectID, "regional-us-central1")
	dbID := validLength(fmt.Sprintf("smpl-%s", id), t)
	dbName = fmt.Sprintf("%s/databases/%s", instName, dbID)

	return
}

func testContext() (Context, error) {
	tc := Context{}

	tc.ProjectID = os.Getenv("GOLANG_SAMPLES_PROJECT_ID")
	if tc.ProjectID == "" {
		return tc, errNoProjectID
	}

	dir, err := os.Getwd()
	if err != nil {
		return tc, fmt.Errorf("could not find current directory")
	}
	if !strings.Contains(dir, "golang-samples") {
		return tc, fmt.Errorf("could not find golang-samples directory")
	}
	tc.Dir = dir[:strings.Index(dir, "golang-samples")+len("golang-samples")]

	if tc.Dir == "" {
		return tc, fmt.Errorf("could not find golang-samples directory")
	}

	return tc, nil
}

var errNoProjectID = errors.New("GOLANG_SAMPLES_PROJECT_ID not set")

// Context holds information useful for tests.
type Context struct {
	ProjectID string
	Dir       string
}

func SystemTest(t *testing.T) Context {
	tc, err := testContext()
	if err == errNoProjectID {
		t.Skip(err)
	} else if err != nil {
		t.Fatal(err)
	}

	return tc
}

func runSample(t *testing.T, f sampleFunc, dbName, errMsg string) string {
	var b bytes.Buffer
	if err := f(&b, dbName); err != nil {
		t.Errorf("%s: %v", errMsg, err)
	}
	return b.String()
}

func runSampleWithContext(ctx context.Context, t *testing.T, f sampleFuncWithContext, dbName, errMsg string) string {
	var b bytes.Buffer
	if err := f(ctx, &b, dbName); err != nil {
		t.Errorf("%s: %v", errMsg, err)
	}
	return b.String()
}

func mustRunSample(t *testing.T, f sampleFuncWithContext, dbName, errMsg string) string {
	var b bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	if err := f(ctx, &b, dbName); err != nil {
		t.Fatalf("%s: %v", errMsg, err)
	}
	return b.String()
}

func createTestInstance(t *testing.T, projectID string, instanceConfigName string) (instanceName string, cleanup func()) {
	ctx := context.Background()
	instanceID := fmt.Sprintf("go-sample-%s", uuid.New().String()[:16])
	instanceName = fmt.Sprintf("projects/%s/instances/%s", projectID, instanceID)
	instanceAdmin, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		t.Fatalf("failed to create InstanceAdminClient: %v", err)
	}
	databaseAdmin, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		t.Fatalf("failed to create DatabaseAdminClient: %v", err)
	}

	// Cleanup old test instances that might not have been deleted.
	iter := instanceAdmin.ListInstances(ctx, &instancepb.ListInstancesRequest{
		Parent: fmt.Sprintf("projects/%v", projectID),
		Filter: "labels.cloud_spanner_samples_test:true",
	})
	for {
		instance, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatalf("failed to list existing instances: %v", err)
		}
		if createTimeString, ok := instance.Labels["create_time"]; ok {
			seconds, err := strconv.ParseInt(createTimeString, 10, 64)
			if err != nil {
				t.Logf("could not parse create time %v: %v", createTimeString, err)
				continue
			}
			createTime := time.Unix(seconds, 0)
			diff := time.Now().Sub(createTime)
			if diff > time.Hour*24 {
				t.Logf("deleting stale test instance %v", instance.Name)
				deleteInstanceAndBackups(t, instance.Name, instanceAdmin, databaseAdmin)
			}
		}
	}

	instanceConfigName = fmt.Sprintf("projects/%s/instanceConfigs/%s", projectID, instanceConfigName)

	Retry(t, 20, time.Minute, func(r *R) {
		op, err := instanceAdmin.CreateInstance(ctx, &instancepb.CreateInstanceRequest{
			Parent:     fmt.Sprintf("projects/%s", projectID),
			InstanceId: instanceID,
			Instance: &instancepb.Instance{
				Config:      instanceConfigName,
				DisplayName: instanceID,
				NodeCount:   1,
				Labels: map[string]string{
					"cloud_spanner_samples_test": "true",
					"create_time":                fmt.Sprintf("%v", time.Now().Unix()),
				},
			},
		})
		if err != nil {
			// Retry if the instance could not be created because there have
			// been too many create requests in the past minute.
			if spanner.ErrCode(err) == codes.ResourceExhausted && strings.Contains(err.Error(), "Quota exceeded for quota metric 'Instance create requests'") {
				r.Errorf("could not create instance %s: %v", fmt.Sprintf("projects/%s/instances/%s", projectID, instanceID), err)
				return
			} else {
				t.Fatalf("could not create instance %s: %v", fmt.Sprintf("projects/%s/instances/%s", projectID, instanceID), err)
			}
		}
		_, err = op.Wait(ctx)
		if err != nil {
			t.Fatalf("waiting for instance creation to finish failed: %v", err)
		}
	})

	return instanceName, func() {
		deleteInstanceAndBackups(t, instanceName, instanceAdmin, databaseAdmin)
		instanceAdmin.Close()
		databaseAdmin.Close()
	}
}

func deleteInstanceAndBackups(
	t *testing.T,
	instanceName string,
	instanceAdmin *instance.InstanceAdminClient,
	databaseAdmin *database.DatabaseAdminClient) {
	ctx := context.Background()
	// Delete all backups before deleting the instance.
	iter := databaseAdmin.ListBackups(ctx, &adminpb.ListBackupsRequest{
		Parent: instanceName,
	})
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatalf("Failed to list backups for instance %s: %v", instanceName, err)
		}
		databaseAdmin.DeleteBackup(ctx, &adminpb.DeleteBackupRequest{Name: resp.Name})
	}
	instanceAdmin.DeleteInstance(ctx, &instancepb.DeleteInstanceRequest{Name: instanceName})
}

func getSampleProjectId(t *testing.T) string {
	// These tests get the project id from the environment variable
	// GOLANG_SAMPLES_SPANNER that is also used by other integration tests for
	// Spanner samples. The tests in this file create a separate instance for
	// each test, so only the project id is used, and the rest of the instance
	// name is ignored.
	instance := os.Getenv("GOLANG_SAMPLES_SPANNER")
	if instance == "" {
		t.Skip("Skipping spanner integration test. Set GOLANG_SAMPLES_SPANNER.")
	}
	if !strings.HasPrefix(instance, "projects/") {
		t.Fatal("Spanner instance ref must be in the form of 'projects/PROJECT_ID/instances/INSTANCE_ID'")
	}
	projectId, _, err := parseInstanceName(instance)
	if err != nil {
		t.Fatalf("Could not parse project id from instance name %q: %v", instance, err)
	}
	return projectId
}

func assertContains(t *testing.T, out string, sub string) {
	t.Helper()
	if !strings.Contains(out, sub) {
		t.Errorf("got output %q; want it to contain %q", out, sub)
	}
}

func assertNotContains(t *testing.T, out string, sub string) {
	t.Helper()
	if strings.Contains(out, sub) {
		t.Errorf("got output %q; want it to not contain %q", out, sub)
	}
}

// Maximum length of database name is 30 characters, so trim if the generated name is too long
func validLength(databaseName string, t *testing.T) (trimmedName string) {
	if len(databaseName) > 30 {
		trimmedName := databaseName[:30]
		t.Logf("Name too long, '%s' trimmed to '%s'", databaseName, trimmedName)
		return trimmedName
	}

	return databaseName
}

func randomID() string {
	now := time.Now().UTC()
	return fmt.Sprintf("%s-%s", strconv.FormatInt(now.Unix(), 10), uuid.New().String()[:8])
}

func parseInstanceName(instanceName string) (project, instance string, err error) {
	matches := validInstancePattern.FindStringSubmatch(instanceName)
	if len(matches) == 0 {
		return "", "", fmt.Errorf("failed to parse database name from %q according to pattern %q",
			instanceName, validInstancePattern.String())
	}
	return matches[1], matches[2], nil
}

// Retry runs function f for up to maxAttempts times until f returns successfully, and reports whether f was run successfully.
// It will sleep for the given period between invocations of f.
// Use the provided *testutil.R instead of a *testing.T from the function.
func Retry(t *testing.T, maxAttempts int, sleep time.Duration, f func(r *R)) bool {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		r := &R{Attempt: attempt, log: &bytes.Buffer{}}

		f(r)

		if !r.failed {
			if r.log.Len() != 0 {
				t.Logf("Success after %d attempts:%s", attempt, r.log.String())
			}
			return true
		}

		if attempt == maxAttempts {
			t.Logf("FAILED after %d attempts:%s", attempt, r.log.String())
			t.Fail()
		}

		time.Sleep(sleep)
	}
	return false
}

// R is passed to each run of a flaky test run, manages state and accumulates log statements.
type R struct {
	// The number of current attempt.
	Attempt int

	failed bool
	log    *bytes.Buffer
}

// Fail marks the run as failed, and will retry once the function returns.
func (r *R) Fail() {
	r.failed = true
}

// Errorf is equivalent to Logf followed by Fail.
func (r *R) Errorf(s string, v ...interface{}) {
	r.logf(s, v...)
	r.Fail()
}

// Logf formats its arguments and records it in the error log.
// The text is only printed for the final unsuccessful run or the first successful run.
func (r *R) Logf(s string, v ...interface{}) {
	r.logf(s, v...)
}

func (r *R) logf(s string, v ...interface{}) {
	fmt.Fprint(r.log, "\n")
	fmt.Fprint(r.log, lineNumber())
	fmt.Fprintf(r.log, s, v...)
}

func lineNumber() string {
	_, file, line, ok := runtime.Caller(3) // logf, public func, user function
	if !ok {
		return ""
	}
	return filepath.Base(file) + ":" + strconv.Itoa(line) + ": "
}
