// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mqstub

import (
	"github.com/aws/aws-sdk-go/service/mq"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateBroker(ctx workflow.Context, input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error)
	CreateBrokerAsync(ctx workflow.Context, input *mq.CreateBrokerRequest) *CreateBrokerFuture

	CreateConfiguration(ctx workflow.Context, input *mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error)
	CreateConfigurationAsync(ctx workflow.Context, input *mq.CreateConfigurationRequest) *CreateConfigurationFuture

	CreateTags(ctx workflow.Context, input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *mq.CreateTagsInput) *CreateTagsFuture

	CreateUser(ctx workflow.Context, input *mq.CreateUserRequest) (*mq.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *mq.CreateUserRequest) *CreateUserFuture

	DeleteBroker(ctx workflow.Context, input *mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error)
	DeleteBrokerAsync(ctx workflow.Context, input *mq.DeleteBrokerInput) *DeleteBrokerFuture

	DeleteTags(ctx workflow.Context, input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *mq.DeleteTagsInput) *DeleteTagsFuture

	DeleteUser(ctx workflow.Context, input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *mq.DeleteUserInput) *DeleteUserFuture

	DescribeBroker(ctx workflow.Context, input *mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error)
	DescribeBrokerAsync(ctx workflow.Context, input *mq.DescribeBrokerInput) *DescribeBrokerFuture

	DescribeBrokerEngineTypes(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error)
	DescribeBrokerEngineTypesAsync(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) *DescribeBrokerEngineTypesFuture

	DescribeBrokerInstanceOptions(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error)
	DescribeBrokerInstanceOptionsAsync(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) *DescribeBrokerInstanceOptionsFuture

	DescribeConfiguration(ctx workflow.Context, input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error)
	DescribeConfigurationAsync(ctx workflow.Context, input *mq.DescribeConfigurationInput) *DescribeConfigurationFuture

	DescribeConfigurationRevision(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error)
	DescribeConfigurationRevisionAsync(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) *DescribeConfigurationRevisionFuture

	DescribeUser(ctx workflow.Context, input *mq.DescribeUserInput) (*mq.DescribeUserResponse, error)
	DescribeUserAsync(ctx workflow.Context, input *mq.DescribeUserInput) *DescribeUserFuture

	ListBrokers(ctx workflow.Context, input *mq.ListBrokersInput) (*mq.ListBrokersResponse, error)
	ListBrokersAsync(ctx workflow.Context, input *mq.ListBrokersInput) *ListBrokersFuture

	ListConfigurationRevisions(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error)
	ListConfigurationRevisionsAsync(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) *ListConfigurationRevisionsFuture

	ListConfigurations(ctx workflow.Context, input *mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error)
	ListConfigurationsAsync(ctx workflow.Context, input *mq.ListConfigurationsInput) *ListConfigurationsFuture

	ListTags(ctx workflow.Context, input *mq.ListTagsInput) (*mq.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *mq.ListTagsInput) *ListTagsFuture

	ListUsers(ctx workflow.Context, input *mq.ListUsersInput) (*mq.ListUsersResponse, error)
	ListUsersAsync(ctx workflow.Context, input *mq.ListUsersInput) *ListUsersFuture

	RebootBroker(ctx workflow.Context, input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error)
	RebootBrokerAsync(ctx workflow.Context, input *mq.RebootBrokerInput) *RebootBrokerFuture

	UpdateBroker(ctx workflow.Context, input *mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error)
	UpdateBrokerAsync(ctx workflow.Context, input *mq.UpdateBrokerRequest) *UpdateBrokerFuture

	UpdateConfiguration(ctx workflow.Context, input *mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error)
	UpdateConfigurationAsync(ctx workflow.Context, input *mq.UpdateConfigurationRequest) *UpdateConfigurationFuture

	UpdateUser(ctx workflow.Context, input *mq.UpdateUserRequest) (*mq.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *mq.UpdateUserRequest) *UpdateUserFuture
}

func NewClient() Client {
	return &stub{}
}
