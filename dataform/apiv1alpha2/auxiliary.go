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

package dataform

import (
	dataformpb "cloud.google.com/go/dataform/apiv1alpha2/dataformpb"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

// CompilationResultActionIterator manages a stream of *dataformpb.CompilationResultAction.
type CompilationResultActionIterator struct {
	items    []*dataformpb.CompilationResultAction
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.CompilationResultAction, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *CompilationResultActionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CompilationResultActionIterator) Next() (*dataformpb.CompilationResultAction, error) {
	var item *dataformpb.CompilationResultAction
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CompilationResultActionIterator) bufLen() int {
	return len(it.items)
}

func (it *CompilationResultActionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// CompilationResultIterator manages a stream of *dataformpb.CompilationResult.
type CompilationResultIterator struct {
	items    []*dataformpb.CompilationResult
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.CompilationResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *CompilationResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CompilationResultIterator) Next() (*dataformpb.CompilationResult, error) {
	var item *dataformpb.CompilationResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CompilationResultIterator) bufLen() int {
	return len(it.items)
}

func (it *CompilationResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// LocationIterator manages a stream of *locationpb.Location.
type LocationIterator struct {
	items    []*locationpb.Location
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
	InternalFetch func(pageSize int, pageToken string) (results []*locationpb.Location, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *LocationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *LocationIterator) Next() (*locationpb.Location, error) {
	var item *locationpb.Location
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *LocationIterator) bufLen() int {
	return len(it.items)
}

func (it *LocationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// QueryDirectoryContentsResponse_DirectoryEntryIterator manages a stream of *dataformpb.QueryDirectoryContentsResponse_DirectoryEntry.
type QueryDirectoryContentsResponse_DirectoryEntryIterator struct {
	items    []*dataformpb.QueryDirectoryContentsResponse_DirectoryEntry
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.QueryDirectoryContentsResponse_DirectoryEntry, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *QueryDirectoryContentsResponse_DirectoryEntryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *QueryDirectoryContentsResponse_DirectoryEntryIterator) Next() (*dataformpb.QueryDirectoryContentsResponse_DirectoryEntry, error) {
	var item *dataformpb.QueryDirectoryContentsResponse_DirectoryEntry
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *QueryDirectoryContentsResponse_DirectoryEntryIterator) bufLen() int {
	return len(it.items)
}

func (it *QueryDirectoryContentsResponse_DirectoryEntryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RepositoryIterator manages a stream of *dataformpb.Repository.
type RepositoryIterator struct {
	items    []*dataformpb.Repository
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.Repository, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *RepositoryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RepositoryIterator) Next() (*dataformpb.Repository, error) {
	var item *dataformpb.Repository
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RepositoryIterator) bufLen() int {
	return len(it.items)
}

func (it *RepositoryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkflowInvocationActionIterator manages a stream of *dataformpb.WorkflowInvocationAction.
type WorkflowInvocationActionIterator struct {
	items    []*dataformpb.WorkflowInvocationAction
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.WorkflowInvocationAction, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *WorkflowInvocationActionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkflowInvocationActionIterator) Next() (*dataformpb.WorkflowInvocationAction, error) {
	var item *dataformpb.WorkflowInvocationAction
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkflowInvocationActionIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkflowInvocationActionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkflowInvocationIterator manages a stream of *dataformpb.WorkflowInvocation.
type WorkflowInvocationIterator struct {
	items    []*dataformpb.WorkflowInvocation
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.WorkflowInvocation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *WorkflowInvocationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkflowInvocationIterator) Next() (*dataformpb.WorkflowInvocation, error) {
	var item *dataformpb.WorkflowInvocation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkflowInvocationIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkflowInvocationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkspaceIterator manages a stream of *dataformpb.Workspace.
type WorkspaceIterator struct {
	items    []*dataformpb.Workspace
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataformpb.Workspace, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *WorkspaceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkspaceIterator) Next() (*dataformpb.Workspace, error) {
	var item *dataformpb.Workspace
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkspaceIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkspaceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
