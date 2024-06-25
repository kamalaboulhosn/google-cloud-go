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

package binaryauthorization_test

import (
	"context"

	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
	binaryauthorizationpb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	"google.golang.org/api/iterator"
)

func ExampleNewBinauthzManagementClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewBinauthzManagementRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleBinauthzManagementClient_CreateAttestor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.CreateAttestorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#CreateAttestorRequest.
	}
	resp, err := c.CreateAttestor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBinauthzManagementClient_DeleteAttestor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.DeleteAttestorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#DeleteAttestorRequest.
	}
	err = c.DeleteAttestor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBinauthzManagementClient_GetAttestor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.GetAttestorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#GetAttestorRequest.
	}
	resp, err := c.GetAttestor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBinauthzManagementClient_GetPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.GetPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#GetPolicyRequest.
	}
	resp, err := c.GetPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBinauthzManagementClient_ListAttestors() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.ListAttestorsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#ListAttestorsRequest.
	}
	it := c.ListAttestors(ctx, req)
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
		_ = it.Response.(*binaryauthorizationpb.ListAttestorsResponse)
	}
}

func ExampleBinauthzManagementClient_UpdateAttestor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.UpdateAttestorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#UpdateAttestorRequest.
	}
	resp, err := c.UpdateAttestor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBinauthzManagementClient_UpdatePolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := binaryauthorization.NewBinauthzManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &binaryauthorizationpb.UpdatePolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb#UpdatePolicyRequest.
	}
	resp, err := c.UpdatePolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
