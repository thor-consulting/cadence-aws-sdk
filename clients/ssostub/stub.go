// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package ssostub

import (
	"github.com/aws/aws-sdk-go/service/sso"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type GetRoleCredentialsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetRoleCredentialsFuture) Get(ctx workflow.Context) (*sso.GetRoleCredentialsOutput, error) {
	var output sso.GetRoleCredentialsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAccountRolesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAccountRolesFuture) Get(ctx workflow.Context) (*sso.ListAccountRolesOutput, error) {
	var output sso.ListAccountRolesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAccountsFuture) Get(ctx workflow.Context) (*sso.ListAccountsOutput, error) {
	var output sso.ListAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LogoutFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LogoutFuture) Get(ctx workflow.Context) (*sso.LogoutOutput, error) {
	var output sso.LogoutOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRoleCredentials(ctx workflow.Context, input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error) {
	var output sso.GetRoleCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-GetRoleCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRoleCredentialsAsync(ctx workflow.Context, input *sso.GetRoleCredentialsInput) *GetRoleCredentialsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-GetRoleCredentials", input)
	return &GetRoleCredentialsFuture{Future: future}
}

func (a *stub) ListAccountRoles(ctx workflow.Context, input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error) {
	var output sso.ListAccountRolesOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-ListAccountRoles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountRolesAsync(ctx workflow.Context, input *sso.ListAccountRolesInput) *ListAccountRolesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-ListAccountRoles", input)
	return &ListAccountRolesFuture{Future: future}
}

func (a *stub) ListAccounts(ctx workflow.Context, input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error) {
	var output sso.ListAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-ListAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountsAsync(ctx workflow.Context, input *sso.ListAccountsInput) *ListAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-ListAccounts", input)
	return &ListAccountsFuture{Future: future}
}

func (a *stub) Logout(ctx workflow.Context, input *sso.LogoutInput) (*sso.LogoutOutput, error) {
	var output sso.LogoutOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-Logout", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) LogoutAsync(ctx workflow.Context, input *sso.LogoutInput) *LogoutFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-Logout", input)
	return &LogoutFuture{Future: future}
}
