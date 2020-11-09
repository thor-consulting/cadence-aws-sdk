// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudtrailstub

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTags(ctx workflow.Context, input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *cloudtrail.AddTagsInput) *AddTagsFuture

	CreateTrail(ctx workflow.Context, input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)
	CreateTrailAsync(ctx workflow.Context, input *cloudtrail.CreateTrailInput) *CreateTrailFuture

	DeleteTrail(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)
	DeleteTrailAsync(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) *DeleteTrailFuture

	DescribeTrails(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)
	DescribeTrailsAsync(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) *DescribeTrailsFuture

	GetEventSelectors(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error)
	GetEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) *GetEventSelectorsFuture

	GetInsightSelectors(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error)
	GetInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) *GetInsightSelectorsFuture

	GetTrail(ctx workflow.Context, input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error)
	GetTrailAsync(ctx workflow.Context, input *cloudtrail.GetTrailInput) *GetTrailFuture

	GetTrailStatus(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)
	GetTrailStatusAsync(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) *GetTrailStatusFuture

	ListPublicKeys(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error)
	ListPublicKeysAsync(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) *ListPublicKeysFuture

	ListTags(ctx workflow.Context, input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *cloudtrail.ListTagsInput) *ListTagsFuture

	ListTrails(ctx workflow.Context, input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error)
	ListTrailsAsync(ctx workflow.Context, input *cloudtrail.ListTrailsInput) *ListTrailsFuture

	LookupEvents(ctx workflow.Context, input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)
	LookupEventsAsync(ctx workflow.Context, input *cloudtrail.LookupEventsInput) *LookupEventsFuture

	PutEventSelectors(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error)
	PutEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) *PutEventSelectorsFuture

	PutInsightSelectors(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error)
	PutInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) *PutInsightSelectorsFuture

	RemoveTags(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) *RemoveTagsFuture

	StartLogging(ctx workflow.Context, input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)
	StartLoggingAsync(ctx workflow.Context, input *cloudtrail.StartLoggingInput) *StartLoggingFuture

	StopLogging(ctx workflow.Context, input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)
	StopLoggingAsync(ctx workflow.Context, input *cloudtrail.StopLoggingInput) *StopLoggingFuture

	UpdateTrail(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
	UpdateTrailAsync(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) *UpdateTrailFuture
}

func NewClient() Client {
	return &stub{}
}
