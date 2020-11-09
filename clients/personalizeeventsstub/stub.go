// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package personalizeeventsstub

import (
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type PutEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutEventsFuture) Get(ctx workflow.Context) (*personalizeevents.PutEventsOutput, error) {
	var output personalizeevents.PutEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutItemsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutItemsFuture) Get(ctx workflow.Context) (*personalizeevents.PutItemsOutput, error) {
	var output personalizeevents.PutItemsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutUsersFuture) Get(ctx workflow.Context) (*personalizeevents.PutUsersOutput, error) {
	var output personalizeevents.PutUsersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) PutEvents(ctx workflow.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
	var output personalizeevents.PutEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws-personalizeevents-PutEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutEventsAsync(ctx workflow.Context, input *personalizeevents.PutEventsInput) *PutEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-personalizeevents-PutEvents", input)
	return &PutEventsFuture{Future: future}
}

func (a *stub) PutItems(ctx workflow.Context, input *personalizeevents.PutItemsInput) (*personalizeevents.PutItemsOutput, error) {
	var output personalizeevents.PutItemsOutput
	err := workflow.ExecuteActivity(ctx, "aws-personalizeevents-PutItems", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutItemsAsync(ctx workflow.Context, input *personalizeevents.PutItemsInput) *PutItemsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-personalizeevents-PutItems", input)
	return &PutItemsFuture{Future: future}
}

func (a *stub) PutUsers(ctx workflow.Context, input *personalizeevents.PutUsersInput) (*personalizeevents.PutUsersOutput, error) {
	var output personalizeevents.PutUsersOutput
	err := workflow.ExecuteActivity(ctx, "aws-personalizeevents-PutUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutUsersAsync(ctx workflow.Context, input *personalizeevents.PutUsersInput) *PutUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-personalizeevents-PutUsers", input)
	return &PutUsersFuture{Future: future}
}
