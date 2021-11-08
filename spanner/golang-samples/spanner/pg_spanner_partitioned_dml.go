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

// [START spanner_postgresql_partitioned_dml]

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/option"
)

func pgPartitionedDml(w io.Writer, db string) error {
	ctx := context.Background()
	// TODO: Remove endpoint
	client, err := spanner.NewClient(ctx, db, option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443"))
	if err != nil {
		return err
	}
	defer client.Close()

	// Spanner PostgreSQL has the same transaction limits as normal Spanner. This includes a
	// maximum of 20,000 mutations in a single read/write transaction. Large update operations can
	// be executed using Partitioned DML. This is also supported on Spanner PostgreSQL.
	// See https://cloud.google.com/spanner/docs/dml-partitioned for more information.
	deletedCount, err := client.PartitionedUpdate(ctx, spanner.Statement{
		SQL: "DELETE FROM users WHERE active=false",
	})
	if err != nil {
		return err
	}
	// The returned update count is the lower bound of the number of records that was deleted.
	fmt.Fprintf(w, "Deleted at least %d inactive users\n", deletedCount)

	return nil
}

// [END spanner_postgresql_partitioned_dml]
