// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package networkmanagerstub

import (
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AssociateCustomerGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateCustomerGatewayFuture) Get(ctx workflow.Context) (*networkmanager.AssociateCustomerGatewayOutput, error) {
	var output networkmanager.AssociateCustomerGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateLinkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateLinkFuture) Get(ctx workflow.Context) (*networkmanager.AssociateLinkOutput, error) {
	var output networkmanager.AssociateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDeviceFuture) Get(ctx workflow.Context) (*networkmanager.CreateDeviceOutput, error) {
	var output networkmanager.CreateDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateGlobalNetworkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateGlobalNetworkFuture) Get(ctx workflow.Context) (*networkmanager.CreateGlobalNetworkOutput, error) {
	var output networkmanager.CreateGlobalNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateLinkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLinkFuture) Get(ctx workflow.Context) (*networkmanager.CreateLinkOutput, error) {
	var output networkmanager.CreateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSiteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSiteFuture) Get(ctx workflow.Context) (*networkmanager.CreateSiteOutput, error) {
	var output networkmanager.CreateSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDeviceFuture) Get(ctx workflow.Context) (*networkmanager.DeleteDeviceOutput, error) {
	var output networkmanager.DeleteDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteGlobalNetworkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteGlobalNetworkFuture) Get(ctx workflow.Context) (*networkmanager.DeleteGlobalNetworkOutput, error) {
	var output networkmanager.DeleteGlobalNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteLinkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteLinkFuture) Get(ctx workflow.Context) (*networkmanager.DeleteLinkOutput, error) {
	var output networkmanager.DeleteLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSiteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSiteFuture) Get(ctx workflow.Context) (*networkmanager.DeleteSiteOutput, error) {
	var output networkmanager.DeleteSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeregisterTransitGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeregisterTransitGatewayFuture) Get(ctx workflow.Context) (*networkmanager.DeregisterTransitGatewayOutput, error) {
	var output networkmanager.DeregisterTransitGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeGlobalNetworksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeGlobalNetworksFuture) Get(ctx workflow.Context) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	var output networkmanager.DescribeGlobalNetworksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateCustomerGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateCustomerGatewayFuture) Get(ctx workflow.Context) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
	var output networkmanager.DisassociateCustomerGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateLinkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateLinkFuture) Get(ctx workflow.Context) (*networkmanager.DisassociateLinkOutput, error) {
	var output networkmanager.DisassociateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetCustomerGatewayAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetCustomerGatewayAssociationsFuture) Get(ctx workflow.Context) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	var output networkmanager.GetCustomerGatewayAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDevicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDevicesFuture) Get(ctx workflow.Context) (*networkmanager.GetDevicesOutput, error) {
	var output networkmanager.GetDevicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLinkAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLinkAssociationsFuture) Get(ctx workflow.Context) (*networkmanager.GetLinkAssociationsOutput, error) {
	var output networkmanager.GetLinkAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetLinksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetLinksFuture) Get(ctx workflow.Context) (*networkmanager.GetLinksOutput, error) {
	var output networkmanager.GetLinksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSitesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSitesFuture) Get(ctx workflow.Context) (*networkmanager.GetSitesOutput, error) {
	var output networkmanager.GetSitesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetTransitGatewayRegistrationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetTransitGatewayRegistrationsFuture) Get(ctx workflow.Context) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	var output networkmanager.GetTransitGatewayRegistrationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*networkmanager.ListTagsForResourceOutput, error) {
	var output networkmanager.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RegisterTransitGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RegisterTransitGatewayFuture) Get(ctx workflow.Context) (*networkmanager.RegisterTransitGatewayOutput, error) {
	var output networkmanager.RegisterTransitGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*networkmanager.TagResourceOutput, error) {
	var output networkmanager.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*networkmanager.UntagResourceOutput, error) {
	var output networkmanager.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDeviceFuture) Get(ctx workflow.Context) (*networkmanager.UpdateDeviceOutput, error) {
	var output networkmanager.UpdateDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateGlobalNetworkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateGlobalNetworkFuture) Get(ctx workflow.Context) (*networkmanager.UpdateGlobalNetworkOutput, error) {
	var output networkmanager.UpdateGlobalNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateLinkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateLinkFuture) Get(ctx workflow.Context) (*networkmanager.UpdateLinkOutput, error) {
	var output networkmanager.UpdateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateSiteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateSiteFuture) Get(ctx workflow.Context) (*networkmanager.UpdateSiteOutput, error) {
	var output networkmanager.UpdateSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateCustomerGateway(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error) {
	var output networkmanager.AssociateCustomerGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-AssociateCustomerGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) *AssociateCustomerGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-AssociateCustomerGateway", input)
	return &AssociateCustomerGatewayFuture{Future: future}
}

func (a *stub) AssociateLink(ctx workflow.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error) {
	var output networkmanager.AssociateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-AssociateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateLinkAsync(ctx workflow.Context, input *networkmanager.AssociateLinkInput) *AssociateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-AssociateLink", input)
	return &AssociateLinkFuture{Future: future}
}

func (a *stub) CreateDevice(ctx workflow.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error) {
	var output networkmanager.CreateDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDeviceAsync(ctx workflow.Context, input *networkmanager.CreateDeviceInput) *CreateDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateDevice", input)
	return &CreateDeviceFuture{Future: future}
}

func (a *stub) CreateGlobalNetwork(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error) {
	var output networkmanager.CreateGlobalNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateGlobalNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) *CreateGlobalNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateGlobalNetwork", input)
	return &CreateGlobalNetworkFuture{Future: future}
}

func (a *stub) CreateLink(ctx workflow.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error) {
	var output networkmanager.CreateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLinkAsync(ctx workflow.Context, input *networkmanager.CreateLinkInput) *CreateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateLink", input)
	return &CreateLinkFuture{Future: future}
}

func (a *stub) CreateSite(ctx workflow.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error) {
	var output networkmanager.CreateSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateSite", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSiteAsync(ctx workflow.Context, input *networkmanager.CreateSiteInput) *CreateSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-CreateSite", input)
	return &CreateSiteFuture{Future: future}
}

func (a *stub) DeleteDevice(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error) {
	var output networkmanager.DeleteDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDeviceAsync(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) *DeleteDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteDevice", input)
	return &DeleteDeviceFuture{Future: future}
}

func (a *stub) DeleteGlobalNetwork(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error) {
	var output networkmanager.DeleteGlobalNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteGlobalNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) *DeleteGlobalNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteGlobalNetwork", input)
	return &DeleteGlobalNetworkFuture{Future: future}
}

func (a *stub) DeleteLink(ctx workflow.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error) {
	var output networkmanager.DeleteLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteLink", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLinkAsync(ctx workflow.Context, input *networkmanager.DeleteLinkInput) *DeleteLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteLink", input)
	return &DeleteLinkFuture{Future: future}
}

func (a *stub) DeleteSite(ctx workflow.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error) {
	var output networkmanager.DeleteSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteSite", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSiteAsync(ctx workflow.Context, input *networkmanager.DeleteSiteInput) *DeleteSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeleteSite", input)
	return &DeleteSiteFuture{Future: future}
}

func (a *stub) DeregisterTransitGateway(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error) {
	var output networkmanager.DeregisterTransitGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeregisterTransitGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) *DeregisterTransitGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DeregisterTransitGateway", input)
	return &DeregisterTransitGatewayFuture{Future: future}
}

func (a *stub) DescribeGlobalNetworks(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	var output networkmanager.DescribeGlobalNetworksOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DescribeGlobalNetworks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeGlobalNetworksAsync(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) *DescribeGlobalNetworksFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DescribeGlobalNetworks", input)
	return &DescribeGlobalNetworksFuture{Future: future}
}

func (a *stub) DisassociateCustomerGateway(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
	var output networkmanager.DisassociateCustomerGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DisassociateCustomerGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) *DisassociateCustomerGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DisassociateCustomerGateway", input)
	return &DisassociateCustomerGatewayFuture{Future: future}
}

func (a *stub) DisassociateLink(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error) {
	var output networkmanager.DisassociateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-DisassociateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateLinkAsync(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) *DisassociateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-DisassociateLink", input)
	return &DisassociateLinkFuture{Future: future}
}

func (a *stub) GetCustomerGatewayAssociations(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	var output networkmanager.GetCustomerGatewayAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetCustomerGatewayAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCustomerGatewayAssociationsAsync(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) *GetCustomerGatewayAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetCustomerGatewayAssociations", input)
	return &GetCustomerGatewayAssociationsFuture{Future: future}
}

func (a *stub) GetDevices(ctx workflow.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
	var output networkmanager.GetDevicesOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetDevices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDevicesAsync(ctx workflow.Context, input *networkmanager.GetDevicesInput) *GetDevicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetDevices", input)
	return &GetDevicesFuture{Future: future}
}

func (a *stub) GetLinkAssociations(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
	var output networkmanager.GetLinkAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetLinkAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLinkAssociationsAsync(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) *GetLinkAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetLinkAssociations", input)
	return &GetLinkAssociationsFuture{Future: future}
}

func (a *stub) GetLinks(ctx workflow.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
	var output networkmanager.GetLinksOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetLinks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLinksAsync(ctx workflow.Context, input *networkmanager.GetLinksInput) *GetLinksFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetLinks", input)
	return &GetLinksFuture{Future: future}
}

func (a *stub) GetSites(ctx workflow.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
	var output networkmanager.GetSitesOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetSites", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSitesAsync(ctx workflow.Context, input *networkmanager.GetSitesInput) *GetSitesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetSites", input)
	return &GetSitesFuture{Future: future}
}

func (a *stub) GetTransitGatewayRegistrations(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	var output networkmanager.GetTransitGatewayRegistrationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetTransitGatewayRegistrations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTransitGatewayRegistrationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) *GetTransitGatewayRegistrationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-GetTransitGatewayRegistrations", input)
	return &GetTransitGatewayRegistrationsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
	var output networkmanager.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) RegisterTransitGateway(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error) {
	var output networkmanager.RegisterTransitGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-RegisterTransitGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) *RegisterTransitGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-RegisterTransitGateway", input)
	return &RegisterTransitGatewayFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error) {
	var output networkmanager.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *networkmanager.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error) {
	var output networkmanager.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *networkmanager.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateDevice(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error) {
	var output networkmanager.UpdateDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDeviceAsync(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) *UpdateDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateDevice", input)
	return &UpdateDeviceFuture{Future: future}
}

func (a *stub) UpdateGlobalNetwork(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error) {
	var output networkmanager.UpdateGlobalNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateGlobalNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) *UpdateGlobalNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateGlobalNetwork", input)
	return &UpdateGlobalNetworkFuture{Future: future}
}

func (a *stub) UpdateLink(ctx workflow.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error) {
	var output networkmanager.UpdateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateLinkAsync(ctx workflow.Context, input *networkmanager.UpdateLinkInput) *UpdateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateLink", input)
	return &UpdateLinkFuture{Future: future}
}

func (a *stub) UpdateSite(ctx workflow.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error) {
	var output networkmanager.UpdateSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateSite", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSiteAsync(ctx workflow.Context, input *networkmanager.UpdateSiteInput) *UpdateSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws-networkmanager-UpdateSite", input)
	return &UpdateSiteFuture{Future: future}
}
