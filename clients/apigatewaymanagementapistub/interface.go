// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package apigatewaymanagementapistub

import (
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteConnection(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) *DeleteConnectionFuture

	GetConnection(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error)
	GetConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) *GetConnectionFuture

	PostToConnection(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error)
	PostToConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) *PostToConnectionFuture
}

func NewClient() Client {
	return &stub{}
}
