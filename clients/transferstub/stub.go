// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package transferstub

import (
	"github.com/aws/aws-sdk-go/service/transfer"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type TransferCreateServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferCreateServerFuture) Get(ctx workflow.Context) (*transfer.CreateServerOutput, error) {
	var output transfer.CreateServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferCreateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferCreateUserFuture) Get(ctx workflow.Context) (*transfer.CreateUserOutput, error) {
	var output transfer.CreateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferDeleteServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferDeleteServerFuture) Get(ctx workflow.Context) (*transfer.DeleteServerOutput, error) {
	var output transfer.DeleteServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferDeleteSshPublicKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferDeleteSshPublicKeyFuture) Get(ctx workflow.Context) (*transfer.DeleteSshPublicKeyOutput, error) {
	var output transfer.DeleteSshPublicKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferDeleteUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferDeleteUserFuture) Get(ctx workflow.Context) (*transfer.DeleteUserOutput, error) {
	var output transfer.DeleteUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferDescribeSecurityPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferDescribeSecurityPolicyFuture) Get(ctx workflow.Context) (*transfer.DescribeSecurityPolicyOutput, error) {
	var output transfer.DescribeSecurityPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferDescribeServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferDescribeServerFuture) Get(ctx workflow.Context) (*transfer.DescribeServerOutput, error) {
	var output transfer.DescribeServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferDescribeUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferDescribeUserFuture) Get(ctx workflow.Context) (*transfer.DescribeUserOutput, error) {
	var output transfer.DescribeUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferImportSshPublicKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferImportSshPublicKeyFuture) Get(ctx workflow.Context) (*transfer.ImportSshPublicKeyOutput, error) {
	var output transfer.ImportSshPublicKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferListSecurityPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferListSecurityPoliciesFuture) Get(ctx workflow.Context) (*transfer.ListSecurityPoliciesOutput, error) {
	var output transfer.ListSecurityPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferListServersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferListServersFuture) Get(ctx workflow.Context) (*transfer.ListServersOutput, error) {
	var output transfer.ListServersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferListTagsForResourceFuture) Get(ctx workflow.Context) (*transfer.ListTagsForResourceOutput, error) {
	var output transfer.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferListUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferListUsersFuture) Get(ctx workflow.Context) (*transfer.ListUsersOutput, error) {
	var output transfer.ListUsersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferStartServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferStartServerFuture) Get(ctx workflow.Context) (*transfer.StartServerOutput, error) {
	var output transfer.StartServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferStopServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferStopServerFuture) Get(ctx workflow.Context) (*transfer.StopServerOutput, error) {
	var output transfer.StopServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferTagResourceFuture) Get(ctx workflow.Context) (*transfer.TagResourceOutput, error) {
	var output transfer.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferTestIdentityProviderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferTestIdentityProviderFuture) Get(ctx workflow.Context) (*transfer.TestIdentityProviderOutput, error) {
	var output transfer.TestIdentityProviderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferUntagResourceFuture) Get(ctx workflow.Context) (*transfer.UntagResourceOutput, error) {
	var output transfer.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferUpdateServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferUpdateServerFuture) Get(ctx workflow.Context) (*transfer.UpdateServerOutput, error) {
	var output transfer.UpdateServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TransferUpdateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TransferUpdateUserFuture) Get(ctx workflow.Context) (*transfer.UpdateUserOutput, error) {
	var output transfer.UpdateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateServer(ctx workflow.Context, input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error) {
	var output transfer.CreateServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-CreateServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateServerAsync(ctx workflow.Context, input *transfer.CreateServerInput) *TransferCreateServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-CreateServer", input)
	return &TransferCreateServerFuture{Future: future}
}

func (a *stub) CreateUser(ctx workflow.Context, input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error) {
	var output transfer.CreateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-CreateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateUserAsync(ctx workflow.Context, input *transfer.CreateUserInput) *TransferCreateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-CreateUser", input)
	return &TransferCreateUserFuture{Future: future}
}

func (a *stub) DeleteServer(ctx workflow.Context, input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error) {
	var output transfer.DeleteServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-DeleteServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteServerAsync(ctx workflow.Context, input *transfer.DeleteServerInput) *TransferDeleteServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-DeleteServer", input)
	return &TransferDeleteServerFuture{Future: future}
}

func (a *stub) DeleteSshPublicKey(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error) {
	var output transfer.DeleteSshPublicKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-DeleteSshPublicKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSshPublicKeyAsync(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) *TransferDeleteSshPublicKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-DeleteSshPublicKey", input)
	return &TransferDeleteSshPublicKeyFuture{Future: future}
}

func (a *stub) DeleteUser(ctx workflow.Context, input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error) {
	var output transfer.DeleteUserOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-DeleteUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteUserAsync(ctx workflow.Context, input *transfer.DeleteUserInput) *TransferDeleteUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-DeleteUser", input)
	return &TransferDeleteUserFuture{Future: future}
}

func (a *stub) DescribeSecurityPolicy(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error) {
	var output transfer.DescribeSecurityPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-DescribeSecurityPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSecurityPolicyAsync(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) *TransferDescribeSecurityPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-DescribeSecurityPolicy", input)
	return &TransferDescribeSecurityPolicyFuture{Future: future}
}

func (a *stub) DescribeServer(ctx workflow.Context, input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error) {
	var output transfer.DescribeServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-DescribeServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeServerAsync(ctx workflow.Context, input *transfer.DescribeServerInput) *TransferDescribeServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-DescribeServer", input)
	return &TransferDescribeServerFuture{Future: future}
}

func (a *stub) DescribeUser(ctx workflow.Context, input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error) {
	var output transfer.DescribeUserOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-DescribeUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeUserAsync(ctx workflow.Context, input *transfer.DescribeUserInput) *TransferDescribeUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-DescribeUser", input)
	return &TransferDescribeUserFuture{Future: future}
}

func (a *stub) ImportSshPublicKey(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error) {
	var output transfer.ImportSshPublicKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-ImportSshPublicKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportSshPublicKeyAsync(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) *TransferImportSshPublicKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-ImportSshPublicKey", input)
	return &TransferImportSshPublicKeyFuture{Future: future}
}

func (a *stub) ListSecurityPolicies(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error) {
	var output transfer.ListSecurityPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-ListSecurityPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSecurityPoliciesAsync(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) *TransferListSecurityPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-ListSecurityPolicies", input)
	return &TransferListSecurityPoliciesFuture{Future: future}
}

func (a *stub) ListServers(ctx workflow.Context, input *transfer.ListServersInput) (*transfer.ListServersOutput, error) {
	var output transfer.ListServersOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-ListServers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListServersAsync(ctx workflow.Context, input *transfer.ListServersInput) *TransferListServersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-ListServers", input)
	return &TransferListServersFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error) {
	var output transfer.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *transfer.ListTagsForResourceInput) *TransferListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-ListTagsForResource", input)
	return &TransferListTagsForResourceFuture{Future: future}
}

func (a *stub) ListUsers(ctx workflow.Context, input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error) {
	var output transfer.ListUsersOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-ListUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListUsersAsync(ctx workflow.Context, input *transfer.ListUsersInput) *TransferListUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-ListUsers", input)
	return &TransferListUsersFuture{Future: future}
}

func (a *stub) StartServer(ctx workflow.Context, input *transfer.StartServerInput) (*transfer.StartServerOutput, error) {
	var output transfer.StartServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-StartServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartServerAsync(ctx workflow.Context, input *transfer.StartServerInput) *TransferStartServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-StartServer", input)
	return &TransferStartServerFuture{Future: future}
}

func (a *stub) StopServer(ctx workflow.Context, input *transfer.StopServerInput) (*transfer.StopServerOutput, error) {
	var output transfer.StopServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-StopServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopServerAsync(ctx workflow.Context, input *transfer.StopServerInput) *TransferStopServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-StopServer", input)
	return &TransferStopServerFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error) {
	var output transfer.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *transfer.TagResourceInput) *TransferTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-TagResource", input)
	return &TransferTagResourceFuture{Future: future}
}

func (a *stub) TestIdentityProvider(ctx workflow.Context, input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error) {
	var output transfer.TestIdentityProviderOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-TestIdentityProvider", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TestIdentityProviderAsync(ctx workflow.Context, input *transfer.TestIdentityProviderInput) *TransferTestIdentityProviderFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-TestIdentityProvider", input)
	return &TransferTestIdentityProviderFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error) {
	var output transfer.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *transfer.UntagResourceInput) *TransferUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-UntagResource", input)
	return &TransferUntagResourceFuture{Future: future}
}

func (a *stub) UpdateServer(ctx workflow.Context, input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error) {
	var output transfer.UpdateServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-UpdateServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateServerAsync(ctx workflow.Context, input *transfer.UpdateServerInput) *TransferUpdateServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-UpdateServer", input)
	return &TransferUpdateServerFuture{Future: future}
}

func (a *stub) UpdateUser(ctx workflow.Context, input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error) {
	var output transfer.UpdateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws-transfer-UpdateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateUserAsync(ctx workflow.Context, input *transfer.UpdateUserInput) *TransferUpdateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws-transfer-UpdateUser", input)
	return &TransferUpdateUserFuture{Future: future}
}
