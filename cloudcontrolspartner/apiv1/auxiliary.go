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

package cloudcontrolspartner

import (
	cloudcontrolspartnerpb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	"google.golang.org/api/iterator"
)

// AccessApprovalRequestIterator manages a stream of *cloudcontrolspartnerpb.AccessApprovalRequest.
type AccessApprovalRequestIterator struct {
	items    []*cloudcontrolspartnerpb.AccessApprovalRequest
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cloudcontrolspartnerpb.AccessApprovalRequest, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *AccessApprovalRequestIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AccessApprovalRequestIterator) Next() (*cloudcontrolspartnerpb.AccessApprovalRequest, error) {
	var item *cloudcontrolspartnerpb.AccessApprovalRequest
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AccessApprovalRequestIterator) bufLen() int {
	return len(it.items)
}

func (it *AccessApprovalRequestIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// CustomerIterator manages a stream of *cloudcontrolspartnerpb.Customer.
type CustomerIterator struct {
	items    []*cloudcontrolspartnerpb.Customer
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cloudcontrolspartnerpb.Customer, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *CustomerIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CustomerIterator) Next() (*cloudcontrolspartnerpb.Customer, error) {
	var item *cloudcontrolspartnerpb.Customer
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CustomerIterator) bufLen() int {
	return len(it.items)
}

func (it *CustomerIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ViolationIterator manages a stream of *cloudcontrolspartnerpb.Violation.
type ViolationIterator struct {
	items    []*cloudcontrolspartnerpb.Violation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cloudcontrolspartnerpb.Violation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ViolationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ViolationIterator) Next() (*cloudcontrolspartnerpb.Violation, error) {
	var item *cloudcontrolspartnerpb.Violation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ViolationIterator) bufLen() int {
	return len(it.items)
}

func (it *ViolationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkloadIterator manages a stream of *cloudcontrolspartnerpb.Workload.
type WorkloadIterator struct {
	items    []*cloudcontrolspartnerpb.Workload
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cloudcontrolspartnerpb.Workload, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *WorkloadIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkloadIterator) Next() (*cloudcontrolspartnerpb.Workload, error) {
	var item *cloudcontrolspartnerpb.Workload
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkloadIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkloadIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
