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

// [START spanner_postgresql_query_parameter]

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func pgQueryParameter(w io.Writer, db string) error {
	ctx := context.Background()
	// TODO: Remove endpoint
	client, err := spanner.NewClient(ctx, db, option.WithEndpoint("staging-wrenchworks.sandbox.googleapis.com:443"))
	if err != nil {
		return err
	}
	defer client.Close()

	stmt := spanner.Statement{
		SQL: `SELECT SingerId, FirstName, LastName
              FROM Singers
              WHERE LastName LIKE $1`,
		// Use 'p1' to bind to the parameter with index 1.
		Params: map[string]interface{}{"p1": "A%"},
	}
	fmt.Fprint(w, "Listing all singers with a last name that starts with 'A'\n")
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		var singerID int64
		var firstName, lastName string
		if err := row.Columns(&singerID, &firstName, &lastName); err != nil {
			return err
		}
		fmt.Fprintf(w, "%d %s %s\n", singerID, firstName, lastName)
	}
}

// [END spanner_postgresql_query_parameter]
