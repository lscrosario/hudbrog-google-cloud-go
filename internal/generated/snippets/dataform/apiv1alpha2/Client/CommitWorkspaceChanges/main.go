// Copyright 2025 Google LLC
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

// [START dataform_v1alpha2_generated_Dataform_CommitWorkspaceChanges_sync]

package main

import (
	"context"

	dataform "cloud.google.com/go/dataform/apiv1alpha2"
	dataformpb "cloud.google.com/go/dataform/apiv1alpha2/dataformpb"
)

func main() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := dataform.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataformpb.CommitWorkspaceChangesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dataform/apiv1alpha2/dataformpb#CommitWorkspaceChangesRequest.
	}
	err = c.CommitWorkspaceChanges(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END dataform_v1alpha2_generated_Dataform_CommitWorkspaceChanges_sync]
