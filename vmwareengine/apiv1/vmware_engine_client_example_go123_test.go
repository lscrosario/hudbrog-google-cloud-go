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

//go:build go1.23

package vmwareengine_test

import (
	"context"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	vmwareengine "cloud.google.com/go/vmwareengine/apiv1"
	vmwareenginepb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func ExampleClient_FetchNetworkPolicyExternalAddresses_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.FetchNetworkPolicyExternalAddressesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#FetchNetworkPolicyExternalAddressesRequest.
	}
	for resp, err := range c.FetchNetworkPolicyExternalAddresses(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListClusters_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListClustersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListClustersRequest.
	}
	for resp, err := range c.ListClusters(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListExternalAccessRules_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListExternalAccessRulesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListExternalAccessRulesRequest.
	}
	for resp, err := range c.ListExternalAccessRules(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListExternalAddresses_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListExternalAddressesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListExternalAddressesRequest.
	}
	for resp, err := range c.ListExternalAddresses(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListHcxActivationKeys_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListHcxActivationKeysRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListHcxActivationKeysRequest.
	}
	for resp, err := range c.ListHcxActivationKeys(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListLoggingServers_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListLoggingServersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListLoggingServersRequest.
	}
	for resp, err := range c.ListLoggingServers(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListManagementDnsZoneBindings_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListManagementDnsZoneBindingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListManagementDnsZoneBindingsRequest.
	}
	for resp, err := range c.ListManagementDnsZoneBindings(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListNetworkPeerings_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListNetworkPeeringsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListNetworkPeeringsRequest.
	}
	for resp, err := range c.ListNetworkPeerings(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListNetworkPolicies_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListNetworkPoliciesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListNetworkPoliciesRequest.
	}
	for resp, err := range c.ListNetworkPolicies(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListNodeTypes_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListNodeTypesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListNodeTypesRequest.
	}
	for resp, err := range c.ListNodeTypes(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListNodes_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListNodesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListNodesRequest.
	}
	for resp, err := range c.ListNodes(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListPeeringRoutes_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListPeeringRoutesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListPeeringRoutesRequest.
	}
	for resp, err := range c.ListPeeringRoutes(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListPrivateClouds_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListPrivateCloudsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListPrivateCloudsRequest.
	}
	for resp, err := range c.ListPrivateClouds(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListPrivateConnectionPeeringRoutes_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListPrivateConnectionPeeringRoutesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListPrivateConnectionPeeringRoutesRequest.
	}
	for resp, err := range c.ListPrivateConnectionPeeringRoutes(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListPrivateConnections_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListPrivateConnectionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListPrivateConnectionsRequest.
	}
	for resp, err := range c.ListPrivateConnections(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListSubnets_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListSubnetsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListSubnetsRequest.
	}
	for resp, err := range c.ListSubnets(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListVmwareEngineNetworks_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vmwareenginepb.ListVmwareEngineNetworksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb#ListVmwareEngineNetworksRequest.
	}
	for resp, err := range c.ListVmwareEngineNetworks(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListLocations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	for resp, err := range c.ListLocations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListOperations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := vmwareengine.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	for resp, err := range c.ListOperations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
