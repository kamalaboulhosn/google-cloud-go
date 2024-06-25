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

// [START merchantapi_v1beta_generated_EmailPreferencesService_GetEmailPreferences_sync]

package main

import (
	"context"

	accounts "cloud.google.com/go/shopping/merchant/accounts/apiv1beta"
	accountspb "cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb"
)

func main() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := accounts.NewEmailPreferencesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accountspb.GetEmailPreferencesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb#GetEmailPreferencesRequest.
	}
	resp, err := c.GetEmailPreferences(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END merchantapi_v1beta_generated_EmailPreferencesService_GetEmailPreferences_sync]
