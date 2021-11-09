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
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"google.golang.org/api/option"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func TestPgCreateDatabase(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	out := runSampleWithContext(ctx, t, pgCreateDatabase, dbName, "failed to create a Spanner PG database")
	assertContains(t, out, fmt.Sprintf("Created Spanner PostgreSQL database [%s]", dbName))
}

func TestPgQueryParameter(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(
		dbName,
		`CREATE TABLE Singers (
		   SingerId  bigint NOT NULL PRIMARY KEY,
		   FirstName varchar(1024),
		   LastName  varchar(1024)
		 )`)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}
	// TODO: Remove endpoint
	client, err := spanner.NewClient(context.Background(), dbName, option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443"))
	if err != nil {
		t.Fatalf("failed to create Spanner client: %v", err)
	}
	if _, err := client.Apply(context.Background(), []*spanner.Mutation{
		spanner.InsertOrUpdateMap("Singers", map[string]interface{}{
			"SingerId":  1,
			"FirstName": "Bruce",
			"LastName":  "Allison",
		}),
		spanner.InsertOrUpdateMap("Singers", map[string]interface{}{
			"SingerId":  2,
			"FirstName": "Alice",
			"LastName":  "Bruxelles",
		}),
	}); err != nil {
		t.Fatalf("failed to insert test records: %v", err)
	}

	out := runSample(t, pgQueryParameter, dbName, "failed to execute PG query with parameter")
	assertContains(t, out, "1 Bruce Allison")
	assertNotContains(t, out, "2 Alice Bruxelles")
}

func TestPgDmlWithParameters(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(
		dbName,
		`CREATE TABLE Singers (
		   SingerId  bigint NOT NULL PRIMARY KEY,
		   FirstName varchar(1024),
		   LastName  varchar(1024)
		 )`)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgDmlWithParameters, dbName, "failed to execute PG DML with parameter")
	assertContains(t, out, "Inserted 2 singers")
}

func TestPgNumericDataType(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgNumericDataType, dbName, "failed to execute PG Numeric sample")
	assertContains(t, out, "Inserted 1 venue(s)")
	assertContains(t, out, "Revenues of Venue 1: 3150.25")
	assertContains(t, out, "Revenues of Venue 2: <null>")
	assertContains(t, out, "Revenues of Venue 3: NaN")
	assertContains(t, out, "Inserted 2 Venues using mutations")
}

func TestPgFunctions(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgFunctions, dbName, "failed to execute PG functions sample")
	assertContains(t, out, "1284352323 seconds after epoch is 2010-09-13 04:32:03 +0000 UTC")
}

func TestPgInformationSchema(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(dbName)
	if matches == nil || len(matches) != 3 {
		t.Fatalf("Invalid database id %s", dbName)
	}
	out := runSample(t, pgInformationSchema, dbName, "failed to execute PG INFORMATION_SCHEMA sample")
	assertContains(t, out, fmt.Sprintf("Table: %s.public.venues (User defined type: null)", matches[2]))
}

func TestPgCastDataType(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgCastDataType, dbName, "failed to execute PG cast data type sample")
	assertContains(t, out, "String: 1")
	assertContains(t, out, "Int: 2")
	assertContains(t, out, "Decimal: 3")
	assertContains(t, out, "Bytes: 4")
	assertContains(t, out, "Bool: true")
	assertContains(t, out, "Timestamp: 2021-11-03 09:35:01 +0000 UTC")
}

func TestPgInterleavedTable(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgInterleavedTable, dbName, "failed to execute PG interleaved table sample")
	assertContains(t, out, "Created interleaved table hierarchy using PostgreSQL dialect")
}

func TestPgBatchDml(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(
		dbName,
		`CREATE TABLE Singers (
		   SingerId  bigint NOT NULL PRIMARY KEY,
		   FirstName varchar(1024),
		   LastName  varchar(1024)
		 )`)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgBatchDml, dbName, "failed to execute PG Batch DML sample")
	assertContains(t, out, "Inserted [1 1] singers")
}

func TestPgPartitionedDml(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(
		dbName,
		`CREATE TABLE users (
			user_id   bigint NOT NULL PRIMARY KEY,
			user_name varchar(1024),
			active    boolean
		)`)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}
	// TODO: Remove endpoint
	client, err := spanner.NewClient(context.Background(), dbName, option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443"))
	if err != nil {
		t.Fatalf("failed to create Spanner client: %v", err)
	}
	if _, err := client.Apply(context.Background(), []*spanner.Mutation{
		spanner.InsertOrUpdateMap("users", map[string]interface{}{
			"user_id":   1,
			"user_name": "User 1",
			"active":    false,
		}),
		spanner.InsertOrUpdateMap("users", map[string]interface{}{
			"user_id":   2,
			"user_name": "User 2",
			"active":    false,
		}),
		spanner.InsertOrUpdateMap("users", map[string]interface{}{
			"user_id":   3,
			"user_name": "User 3",
			"active":    true,
		}),
	}); err != nil {
		t.Fatalf("failed to insert test records: %v", err)
	}

	out := runSample(t, pgPartitionedDml, dbName, "failed to execute PG Partitioned DML sample")
	assertContains(t, out, "Deleted at least 2 inactive users")
}

func TestPgCaseSensitivity(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgCaseSensitivity, dbName, "failed to execute PG case sensitivity sample")
	assertContains(t, out, "SingerId: 1, FirstName: Bruce, LastName: Allison")
	assertContains(t, out, "SingerId: 1, FullName: Bruce Allison")
}

func TestPgOrderNulls(t *testing.T) {
	_ = SystemTest(t)
	t.Parallel()

	_, dbName, cleanup := initTest(t, randomID())
	defer cleanup()
	dbCleanup, err := createTestPgDatabase(dbName)
	if dbCleanup != nil {
		defer dbCleanup()
	}
	if err != nil {
		t.Fatalf("failed to create test database: %v", err)
	}

	out := runSample(t, pgOrderNulls, dbName, "failed to execute PG order nulls sample")
	assertContains(t, out, "Singers ORDER BY Name\n\tAlice\n\tBruce\n\t<null>")
	assertContains(t, out, "Singers ORDER BY Name DESC\n\t<null>\n\tBruce\n\tAlice")
	assertContains(t, out, "Singers ORDER BY Name NULLS FIRST\n\t<null>\n\tAlice\n\tBruce")
	assertContains(t, out, "Singers ORDER BY Name DESC NULLS LAST\n\tBruce\n\tAlice\n\t<null>")
}

func createTestPgDatabase(db string, extraStatements ...string) (func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(db)
	if matches == nil || len(matches) != 3 {
		return func() {}, fmt.Errorf("invalid database id %s", db)
	}

	// TODO: Remove endpoint
	adminClient, err := database.NewDatabaseAdminClient(ctx, option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443"))
	if err != nil {
		return func() {}, err
	}
	defer adminClient.Close()

	opCreate, err := adminClient.CreateDatabase(ctx, &adminpb.CreateDatabaseRequest{
		Parent:          matches[1],
		DatabaseDialect: adminpb.DatabaseDialect_POSTGRESQL,
		CreateStatement: `CREATE DATABASE "` + matches[2] + `"`,
	})
	if err != nil {
		return func() {}, err
	}
	if _, err := opCreate.Wait(ctx); err != nil {
		return func() {}, err
	}
	dropDb := func() {
		adminClient, err := database.NewDatabaseAdminClient(ctx, option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443"))
		if err != nil {
			return
		}
		adminClient.DropDatabase(context.Background(), &adminpb.DropDatabaseRequest{
			Database: db,
		})
		defer adminClient.Close()
	}
	if len(extraStatements) > 0 {
		opUpdate, err := adminClient.UpdateDatabaseDdl(ctx, &adminpb.UpdateDatabaseDdlRequest{
			Database:   db,
			Statements: extraStatements,
		})
		if err != nil {
			return dropDb, err
		}
		if err := opUpdate.Wait(ctx); err != nil {
			return dropDb, err
		}
	}
	return dropDb, nil
}
