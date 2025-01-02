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

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
)

func ExampleNewSecurityPoliciesRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleSecurityPoliciesClient_AddRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AddRuleSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddRuleSecurityPolicyRequest.
	}
	op, err := c.AddRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecurityPoliciesClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AggregatedListSecurityPoliciesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListSecurityPoliciesRequest.
	}
	it := c.AggregatedList(ctx, req)
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

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.SecurityPoliciesAggregatedList)
	}
}

func ExampleSecurityPoliciesClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteSecurityPolicyRequest.
	}
	op, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecurityPoliciesClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetSecurityPolicyRequest.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_GetRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetRuleSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRuleSecurityPolicyRequest.
	}
	resp, err := c.GetRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.InsertSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertSecurityPolicyRequest.
	}
	op, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecurityPoliciesClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListSecurityPoliciesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListSecurityPoliciesRequest.
	}
	it := c.List(ctx, req)
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

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.SecurityPolicyList)
	}
}

func ExampleSecurityPoliciesClient_ListPreconfiguredExpressionSets() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListPreconfiguredExpressionSetsSecurityPoliciesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListPreconfiguredExpressionSetsSecurityPoliciesRequest.
	}
	resp, err := c.ListPreconfiguredExpressionSets(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.PatchSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchSecurityPolicyRequest.
	}
	op, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecurityPoliciesClient_PatchRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.PatchRuleSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRuleSecurityPolicyRequest.
	}
	op, err := c.PatchRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecurityPoliciesClient_RemoveRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.RemoveRuleSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveRuleSecurityPolicyRequest.
	}
	op, err := c.RemoveRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecurityPoliciesClient_SetLabels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewSecurityPoliciesRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetLabelsSecurityPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetLabelsSecurityPolicyRequest.
	}
	op, err := c.SetLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
