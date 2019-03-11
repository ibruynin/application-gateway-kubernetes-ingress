// Package consumption implements the Azure ARM Consumption service API version 2018-03-31.
//
// Consumption management client provides access to consumption resources for Azure Enterprise Subscriptions.
package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Consumption
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Consumption.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// GetBalancesByBillingAccount gets the balances for a scope by billingAccountId. Balances are available via this API
// only for May 1, 2014 or later.
// Parameters:
// billingAccountID - billingAccount ID
func (client BaseClient) GetBalancesByBillingAccount(ctx context.Context, billingAccountID string) (result Balance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetBalancesByBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetBalancesByBillingAccountPreparer(ctx, billingAccountID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BaseClient", "GetBalancesByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBalancesByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BaseClient", "GetBalancesByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.GetBalancesByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BaseClient", "GetBalancesByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// GetBalancesByBillingAccountPreparer prepares the GetBalancesByBillingAccount request.
func (client BaseClient) GetBalancesByBillingAccountPreparer(ctx context.Context, billingAccountID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/balances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBalancesByBillingAccountSender sends the GetBalancesByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetBalancesByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBalancesByBillingAccountResponder handles the response to the GetBalancesByBillingAccount request. The method always
// closes the http.Response Body.
func (client BaseClient) GetBalancesByBillingAccountResponder(resp *http.Response) (result Balance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
