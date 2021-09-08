// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START orgpolicy_v2_generated_OrgPolicy_ListConstraints_sync]

package main

import (
	"context"

	orgpolicy "cloud.google.com/go/orgpolicy/apiv2"
	"google.golang.org/api/iterator"
	orgpolicypb "google.golang.org/genproto/googleapis/cloud/orgpolicy/v2"
)

func main() {
	ctx := context.Background()
	c, err := orgpolicy.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &orgpolicypb.ListConstraintsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/orgpolicy/v2#ListConstraintsRequest.
	}
	it := c.ListConstraints(ctx, req)
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

// [END orgpolicy_v2_generated_OrgPolicy_ListConstraints_sync]
