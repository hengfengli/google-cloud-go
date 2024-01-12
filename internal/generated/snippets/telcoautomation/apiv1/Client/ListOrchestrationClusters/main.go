// Copyright 2024 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

// [START telcoautomation_v1_generated_TelcoAutomation_ListOrchestrationClusters_sync]

package main

import (
	"context"

	telcoautomation "cloud.google.com/go/telcoautomation/apiv1"
	telcoautomationpb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := telcoautomation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &telcoautomationpb.ListOrchestrationClustersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb#ListOrchestrationClustersRequest.
	}
	it := c.ListOrchestrationClusters(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

// [END telcoautomation_v1_generated_TelcoAutomation_ListOrchestrationClusters_sync]
