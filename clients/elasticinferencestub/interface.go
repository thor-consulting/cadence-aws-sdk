// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package elasticinferencestub

import (
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error)
	DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *DescribeAcceleratorOfferingsFuture

	DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error)
	DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *DescribeAcceleratorTypesFuture

	DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error)
	DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *DescribeAcceleratorsFuture

	ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *UntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
