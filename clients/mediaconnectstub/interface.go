// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mediaconnectstub

import (
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddFlowOutputs(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error)
	AddFlowOutputsAsync(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) *AddFlowOutputsFuture

	AddFlowSources(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) (*mediaconnect.AddFlowSourcesOutput, error)
	AddFlowSourcesAsync(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) *AddFlowSourcesFuture

	AddFlowVpcInterfaces(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) (*mediaconnect.AddFlowVpcInterfacesOutput, error)
	AddFlowVpcInterfacesAsync(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) *AddFlowVpcInterfacesFuture

	CreateFlow(ctx workflow.Context, input *mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error)
	CreateFlowAsync(ctx workflow.Context, input *mediaconnect.CreateFlowInput) *CreateFlowFuture

	DeleteFlow(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error)
	DeleteFlowAsync(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) *DeleteFlowFuture

	DescribeFlow(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error)
	DescribeFlowAsync(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) *DescribeFlowFuture

	DescribeOffering(ctx workflow.Context, input *mediaconnect.DescribeOfferingInput) (*mediaconnect.DescribeOfferingOutput, error)
	DescribeOfferingAsync(ctx workflow.Context, input *mediaconnect.DescribeOfferingInput) *DescribeOfferingFuture

	DescribeReservation(ctx workflow.Context, input *mediaconnect.DescribeReservationInput) (*mediaconnect.DescribeReservationOutput, error)
	DescribeReservationAsync(ctx workflow.Context, input *mediaconnect.DescribeReservationInput) *DescribeReservationFuture

	GrantFlowEntitlements(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error)
	GrantFlowEntitlementsAsync(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) *GrantFlowEntitlementsFuture

	ListEntitlements(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error)
	ListEntitlementsAsync(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) *ListEntitlementsFuture

	ListFlows(ctx workflow.Context, input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error)
	ListFlowsAsync(ctx workflow.Context, input *mediaconnect.ListFlowsInput) *ListFlowsFuture

	ListOfferings(ctx workflow.Context, input *mediaconnect.ListOfferingsInput) (*mediaconnect.ListOfferingsOutput, error)
	ListOfferingsAsync(ctx workflow.Context, input *mediaconnect.ListOfferingsInput) *ListOfferingsFuture

	ListReservations(ctx workflow.Context, input *mediaconnect.ListReservationsInput) (*mediaconnect.ListReservationsOutput, error)
	ListReservationsAsync(ctx workflow.Context, input *mediaconnect.ListReservationsInput) *ListReservationsFuture

	ListTagsForResource(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) *ListTagsForResourceFuture

	PurchaseOffering(ctx workflow.Context, input *mediaconnect.PurchaseOfferingInput) (*mediaconnect.PurchaseOfferingOutput, error)
	PurchaseOfferingAsync(ctx workflow.Context, input *mediaconnect.PurchaseOfferingInput) *PurchaseOfferingFuture

	RemoveFlowOutput(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error)
	RemoveFlowOutputAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) *RemoveFlowOutputFuture

	RemoveFlowSource(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) (*mediaconnect.RemoveFlowSourceOutput, error)
	RemoveFlowSourceAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) *RemoveFlowSourceFuture

	RemoveFlowVpcInterface(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error)
	RemoveFlowVpcInterfaceAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) *RemoveFlowVpcInterfaceFuture

	RevokeFlowEntitlement(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error)
	RevokeFlowEntitlementAsync(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) *RevokeFlowEntitlementFuture

	StartFlow(ctx workflow.Context, input *mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error)
	StartFlowAsync(ctx workflow.Context, input *mediaconnect.StartFlowInput) *StartFlowFuture

	StopFlow(ctx workflow.Context, input *mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error)
	StopFlowAsync(ctx workflow.Context, input *mediaconnect.StopFlowInput) *StopFlowFuture

	TagResource(ctx workflow.Context, input *mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediaconnect.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediaconnect.UntagResourceInput) *UntagResourceFuture

	UpdateFlow(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) (*mediaconnect.UpdateFlowOutput, error)
	UpdateFlowAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) *UpdateFlowFuture

	UpdateFlowEntitlement(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error)
	UpdateFlowEntitlementAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) *UpdateFlowEntitlementFuture

	UpdateFlowOutput(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error)
	UpdateFlowOutputAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) *UpdateFlowOutputFuture

	UpdateFlowSource(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error)
	UpdateFlowSourceAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) *UpdateFlowSourceFuture
}

func NewClient() Client {
	return &stub{}
}
