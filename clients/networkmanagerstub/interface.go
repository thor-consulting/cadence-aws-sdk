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

type Client interface {
	AssociateCustomerGateway(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error)
	AssociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) *AssociateCustomerGatewayFuture

	AssociateLink(ctx workflow.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error)
	AssociateLinkAsync(ctx workflow.Context, input *networkmanager.AssociateLinkInput) *AssociateLinkFuture

	AssociateTransitGatewayConnectPeer(ctx workflow.Context, input *networkmanager.AssociateTransitGatewayConnectPeerInput) (*networkmanager.AssociateTransitGatewayConnectPeerOutput, error)
	AssociateTransitGatewayConnectPeerAsync(ctx workflow.Context, input *networkmanager.AssociateTransitGatewayConnectPeerInput) *AssociateTransitGatewayConnectPeerFuture

	CreateConnection(ctx workflow.Context, input *networkmanager.CreateConnectionInput) (*networkmanager.CreateConnectionOutput, error)
	CreateConnectionAsync(ctx workflow.Context, input *networkmanager.CreateConnectionInput) *CreateConnectionFuture

	CreateDevice(ctx workflow.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error)
	CreateDeviceAsync(ctx workflow.Context, input *networkmanager.CreateDeviceInput) *CreateDeviceFuture

	CreateGlobalNetwork(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error)
	CreateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) *CreateGlobalNetworkFuture

	CreateLink(ctx workflow.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error)
	CreateLinkAsync(ctx workflow.Context, input *networkmanager.CreateLinkInput) *CreateLinkFuture

	CreateSite(ctx workflow.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error)
	CreateSiteAsync(ctx workflow.Context, input *networkmanager.CreateSiteInput) *CreateSiteFuture

	DeleteConnection(ctx workflow.Context, input *networkmanager.DeleteConnectionInput) (*networkmanager.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *networkmanager.DeleteConnectionInput) *DeleteConnectionFuture

	DeleteDevice(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error)
	DeleteDeviceAsync(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) *DeleteDeviceFuture

	DeleteGlobalNetwork(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error)
	DeleteGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) *DeleteGlobalNetworkFuture

	DeleteLink(ctx workflow.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error)
	DeleteLinkAsync(ctx workflow.Context, input *networkmanager.DeleteLinkInput) *DeleteLinkFuture

	DeleteSite(ctx workflow.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error)
	DeleteSiteAsync(ctx workflow.Context, input *networkmanager.DeleteSiteInput) *DeleteSiteFuture

	DeregisterTransitGateway(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error)
	DeregisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) *DeregisterTransitGatewayFuture

	DescribeGlobalNetworks(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error)
	DescribeGlobalNetworksAsync(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) *DescribeGlobalNetworksFuture

	DisassociateCustomerGateway(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error)
	DisassociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) *DisassociateCustomerGatewayFuture

	DisassociateLink(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error)
	DisassociateLinkAsync(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) *DisassociateLinkFuture

	DisassociateTransitGatewayConnectPeer(ctx workflow.Context, input *networkmanager.DisassociateTransitGatewayConnectPeerInput) (*networkmanager.DisassociateTransitGatewayConnectPeerOutput, error)
	DisassociateTransitGatewayConnectPeerAsync(ctx workflow.Context, input *networkmanager.DisassociateTransitGatewayConnectPeerInput) *DisassociateTransitGatewayConnectPeerFuture

	GetConnections(ctx workflow.Context, input *networkmanager.GetConnectionsInput) (*networkmanager.GetConnectionsOutput, error)
	GetConnectionsAsync(ctx workflow.Context, input *networkmanager.GetConnectionsInput) *GetConnectionsFuture

	GetCustomerGatewayAssociations(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error)
	GetCustomerGatewayAssociationsAsync(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) *GetCustomerGatewayAssociationsFuture

	GetDevices(ctx workflow.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error)
	GetDevicesAsync(ctx workflow.Context, input *networkmanager.GetDevicesInput) *GetDevicesFuture

	GetLinkAssociations(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error)
	GetLinkAssociationsAsync(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) *GetLinkAssociationsFuture

	GetLinks(ctx workflow.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error)
	GetLinksAsync(ctx workflow.Context, input *networkmanager.GetLinksInput) *GetLinksFuture

	GetSites(ctx workflow.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error)
	GetSitesAsync(ctx workflow.Context, input *networkmanager.GetSitesInput) *GetSitesFuture

	GetTransitGatewayConnectPeerAssociations(ctx workflow.Context, input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput) (*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, error)
	GetTransitGatewayConnectPeerAssociationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput) *GetTransitGatewayConnectPeerAssociationsFuture

	GetTransitGatewayRegistrations(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error)
	GetTransitGatewayRegistrationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) *GetTransitGatewayRegistrationsFuture

	ListTagsForResource(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) *ListTagsForResourceFuture

	RegisterTransitGateway(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error)
	RegisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) *RegisterTransitGatewayFuture

	TagResource(ctx workflow.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *networkmanager.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *networkmanager.UntagResourceInput) *UntagResourceFuture

	UpdateConnection(ctx workflow.Context, input *networkmanager.UpdateConnectionInput) (*networkmanager.UpdateConnectionOutput, error)
	UpdateConnectionAsync(ctx workflow.Context, input *networkmanager.UpdateConnectionInput) *UpdateConnectionFuture

	UpdateDevice(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error)
	UpdateDeviceAsync(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) *UpdateDeviceFuture

	UpdateGlobalNetwork(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error)
	UpdateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) *UpdateGlobalNetworkFuture

	UpdateLink(ctx workflow.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error)
	UpdateLinkAsync(ctx workflow.Context, input *networkmanager.UpdateLinkInput) *UpdateLinkFuture

	UpdateSite(ctx workflow.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error)
	UpdateSiteAsync(ctx workflow.Context, input *networkmanager.UpdateSiteInput) *UpdateSiteFuture
}

func NewClient() Client {
	return &stub{}
}
