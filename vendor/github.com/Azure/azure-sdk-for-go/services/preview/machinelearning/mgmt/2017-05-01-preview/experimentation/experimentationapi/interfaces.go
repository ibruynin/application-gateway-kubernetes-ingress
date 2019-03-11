package experimentationapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-05-01-preview/experimentation"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result experimentation.OperationListResult, err error)
}

var _ OperationsClientAPI = (*experimentation.OperationsClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, parameters experimentation.Account) (result experimentation.Account, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result experimentation.Account, err error)
	List(ctx context.Context) (result experimentation.AccountListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result experimentation.AccountListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters experimentation.AccountUpdateParameters) (result experimentation.Account, err error)
}

var _ AccountsClientAPI = (*experimentation.AccountsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, parameters experimentation.Workspace) (result experimentation.Workspace, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, workspaceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, workspaceName string) (result experimentation.Workspace, err error)
	ListByAccounts(ctx context.Context, accountName string, resourceGroupName string) (result experimentation.WorkspaceListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, parameters experimentation.WorkspaceUpdateParameters) (result experimentation.Workspace, err error)
}

var _ WorkspacesClientAPI = (*experimentation.WorkspacesClient)(nil)

// ProjectsClientAPI contains the set of methods on the ProjectsClient type.
type ProjectsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, projectName string, parameters experimentation.Project) (result experimentation.Project, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, projectName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, projectName string) (result experimentation.Project, err error)
	ListByWorkspace(ctx context.Context, accountName string, workspaceName string, resourceGroupName string) (result experimentation.ProjectListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, workspaceName string, projectName string, parameters experimentation.ProjectUpdateParameters) (result experimentation.Project, err error)
}

var _ ProjectsClientAPI = (*experimentation.ProjectsClient)(nil)
