// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package iotwirelessstub

import (
	"github.com/aws/aws-sdk-go/service/iotwireless"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateAwsAccountWithPartnerAccount(ctx workflow.Context, input *iotwireless.AssociateAwsAccountWithPartnerAccountInput) (*iotwireless.AssociateAwsAccountWithPartnerAccountOutput, error)
	AssociateAwsAccountWithPartnerAccountAsync(ctx workflow.Context, input *iotwireless.AssociateAwsAccountWithPartnerAccountInput) *AssociateAwsAccountWithPartnerAccountFuture

	AssociateWirelessDeviceWithThing(ctx workflow.Context, input *iotwireless.AssociateWirelessDeviceWithThingInput) (*iotwireless.AssociateWirelessDeviceWithThingOutput, error)
	AssociateWirelessDeviceWithThingAsync(ctx workflow.Context, input *iotwireless.AssociateWirelessDeviceWithThingInput) *AssociateWirelessDeviceWithThingFuture

	AssociateWirelessGatewayWithCertificate(ctx workflow.Context, input *iotwireless.AssociateWirelessGatewayWithCertificateInput) (*iotwireless.AssociateWirelessGatewayWithCertificateOutput, error)
	AssociateWirelessGatewayWithCertificateAsync(ctx workflow.Context, input *iotwireless.AssociateWirelessGatewayWithCertificateInput) *AssociateWirelessGatewayWithCertificateFuture

	AssociateWirelessGatewayWithThing(ctx workflow.Context, input *iotwireless.AssociateWirelessGatewayWithThingInput) (*iotwireless.AssociateWirelessGatewayWithThingOutput, error)
	AssociateWirelessGatewayWithThingAsync(ctx workflow.Context, input *iotwireless.AssociateWirelessGatewayWithThingInput) *AssociateWirelessGatewayWithThingFuture

	CreateDestination(ctx workflow.Context, input *iotwireless.CreateDestinationInput) (*iotwireless.CreateDestinationOutput, error)
	CreateDestinationAsync(ctx workflow.Context, input *iotwireless.CreateDestinationInput) *CreateDestinationFuture

	CreateDeviceProfile(ctx workflow.Context, input *iotwireless.CreateDeviceProfileInput) (*iotwireless.CreateDeviceProfileOutput, error)
	CreateDeviceProfileAsync(ctx workflow.Context, input *iotwireless.CreateDeviceProfileInput) *CreateDeviceProfileFuture

	CreateServiceProfile(ctx workflow.Context, input *iotwireless.CreateServiceProfileInput) (*iotwireless.CreateServiceProfileOutput, error)
	CreateServiceProfileAsync(ctx workflow.Context, input *iotwireless.CreateServiceProfileInput) *CreateServiceProfileFuture

	CreateWirelessDevice(ctx workflow.Context, input *iotwireless.CreateWirelessDeviceInput) (*iotwireless.CreateWirelessDeviceOutput, error)
	CreateWirelessDeviceAsync(ctx workflow.Context, input *iotwireless.CreateWirelessDeviceInput) *CreateWirelessDeviceFuture

	CreateWirelessGateway(ctx workflow.Context, input *iotwireless.CreateWirelessGatewayInput) (*iotwireless.CreateWirelessGatewayOutput, error)
	CreateWirelessGatewayAsync(ctx workflow.Context, input *iotwireless.CreateWirelessGatewayInput) *CreateWirelessGatewayFuture

	CreateWirelessGatewayTask(ctx workflow.Context, input *iotwireless.CreateWirelessGatewayTaskInput) (*iotwireless.CreateWirelessGatewayTaskOutput, error)
	CreateWirelessGatewayTaskAsync(ctx workflow.Context, input *iotwireless.CreateWirelessGatewayTaskInput) *CreateWirelessGatewayTaskFuture

	CreateWirelessGatewayTaskDefinition(ctx workflow.Context, input *iotwireless.CreateWirelessGatewayTaskDefinitionInput) (*iotwireless.CreateWirelessGatewayTaskDefinitionOutput, error)
	CreateWirelessGatewayTaskDefinitionAsync(ctx workflow.Context, input *iotwireless.CreateWirelessGatewayTaskDefinitionInput) *CreateWirelessGatewayTaskDefinitionFuture

	DeleteDestination(ctx workflow.Context, input *iotwireless.DeleteDestinationInput) (*iotwireless.DeleteDestinationOutput, error)
	DeleteDestinationAsync(ctx workflow.Context, input *iotwireless.DeleteDestinationInput) *DeleteDestinationFuture

	DeleteDeviceProfile(ctx workflow.Context, input *iotwireless.DeleteDeviceProfileInput) (*iotwireless.DeleteDeviceProfileOutput, error)
	DeleteDeviceProfileAsync(ctx workflow.Context, input *iotwireless.DeleteDeviceProfileInput) *DeleteDeviceProfileFuture

	DeleteServiceProfile(ctx workflow.Context, input *iotwireless.DeleteServiceProfileInput) (*iotwireless.DeleteServiceProfileOutput, error)
	DeleteServiceProfileAsync(ctx workflow.Context, input *iotwireless.DeleteServiceProfileInput) *DeleteServiceProfileFuture

	DeleteWirelessDevice(ctx workflow.Context, input *iotwireless.DeleteWirelessDeviceInput) (*iotwireless.DeleteWirelessDeviceOutput, error)
	DeleteWirelessDeviceAsync(ctx workflow.Context, input *iotwireless.DeleteWirelessDeviceInput) *DeleteWirelessDeviceFuture

	DeleteWirelessGateway(ctx workflow.Context, input *iotwireless.DeleteWirelessGatewayInput) (*iotwireless.DeleteWirelessGatewayOutput, error)
	DeleteWirelessGatewayAsync(ctx workflow.Context, input *iotwireless.DeleteWirelessGatewayInput) *DeleteWirelessGatewayFuture

	DeleteWirelessGatewayTask(ctx workflow.Context, input *iotwireless.DeleteWirelessGatewayTaskInput) (*iotwireless.DeleteWirelessGatewayTaskOutput, error)
	DeleteWirelessGatewayTaskAsync(ctx workflow.Context, input *iotwireless.DeleteWirelessGatewayTaskInput) *DeleteWirelessGatewayTaskFuture

	DeleteWirelessGatewayTaskDefinition(ctx workflow.Context, input *iotwireless.DeleteWirelessGatewayTaskDefinitionInput) (*iotwireless.DeleteWirelessGatewayTaskDefinitionOutput, error)
	DeleteWirelessGatewayTaskDefinitionAsync(ctx workflow.Context, input *iotwireless.DeleteWirelessGatewayTaskDefinitionInput) *DeleteWirelessGatewayTaskDefinitionFuture

	DisassociateAwsAccountFromPartnerAccount(ctx workflow.Context, input *iotwireless.DisassociateAwsAccountFromPartnerAccountInput) (*iotwireless.DisassociateAwsAccountFromPartnerAccountOutput, error)
	DisassociateAwsAccountFromPartnerAccountAsync(ctx workflow.Context, input *iotwireless.DisassociateAwsAccountFromPartnerAccountInput) *DisassociateAwsAccountFromPartnerAccountFuture

	DisassociateWirelessDeviceFromThing(ctx workflow.Context, input *iotwireless.DisassociateWirelessDeviceFromThingInput) (*iotwireless.DisassociateWirelessDeviceFromThingOutput, error)
	DisassociateWirelessDeviceFromThingAsync(ctx workflow.Context, input *iotwireless.DisassociateWirelessDeviceFromThingInput) *DisassociateWirelessDeviceFromThingFuture

	DisassociateWirelessGatewayFromCertificate(ctx workflow.Context, input *iotwireless.DisassociateWirelessGatewayFromCertificateInput) (*iotwireless.DisassociateWirelessGatewayFromCertificateOutput, error)
	DisassociateWirelessGatewayFromCertificateAsync(ctx workflow.Context, input *iotwireless.DisassociateWirelessGatewayFromCertificateInput) *DisassociateWirelessGatewayFromCertificateFuture

	DisassociateWirelessGatewayFromThing(ctx workflow.Context, input *iotwireless.DisassociateWirelessGatewayFromThingInput) (*iotwireless.DisassociateWirelessGatewayFromThingOutput, error)
	DisassociateWirelessGatewayFromThingAsync(ctx workflow.Context, input *iotwireless.DisassociateWirelessGatewayFromThingInput) *DisassociateWirelessGatewayFromThingFuture

	GetDestination(ctx workflow.Context, input *iotwireless.GetDestinationInput) (*iotwireless.GetDestinationOutput, error)
	GetDestinationAsync(ctx workflow.Context, input *iotwireless.GetDestinationInput) *GetDestinationFuture

	GetDeviceProfile(ctx workflow.Context, input *iotwireless.GetDeviceProfileInput) (*iotwireless.GetDeviceProfileOutput, error)
	GetDeviceProfileAsync(ctx workflow.Context, input *iotwireless.GetDeviceProfileInput) *GetDeviceProfileFuture

	GetPartnerAccount(ctx workflow.Context, input *iotwireless.GetPartnerAccountInput) (*iotwireless.GetPartnerAccountOutput, error)
	GetPartnerAccountAsync(ctx workflow.Context, input *iotwireless.GetPartnerAccountInput) *GetPartnerAccountFuture

	GetServiceEndpoint(ctx workflow.Context, input *iotwireless.GetServiceEndpointInput) (*iotwireless.GetServiceEndpointOutput, error)
	GetServiceEndpointAsync(ctx workflow.Context, input *iotwireless.GetServiceEndpointInput) *GetServiceEndpointFuture

	GetServiceProfile(ctx workflow.Context, input *iotwireless.GetServiceProfileInput) (*iotwireless.GetServiceProfileOutput, error)
	GetServiceProfileAsync(ctx workflow.Context, input *iotwireless.GetServiceProfileInput) *GetServiceProfileFuture

	GetWirelessDevice(ctx workflow.Context, input *iotwireless.GetWirelessDeviceInput) (*iotwireless.GetWirelessDeviceOutput, error)
	GetWirelessDeviceAsync(ctx workflow.Context, input *iotwireless.GetWirelessDeviceInput) *GetWirelessDeviceFuture

	GetWirelessDeviceStatistics(ctx workflow.Context, input *iotwireless.GetWirelessDeviceStatisticsInput) (*iotwireless.GetWirelessDeviceStatisticsOutput, error)
	GetWirelessDeviceStatisticsAsync(ctx workflow.Context, input *iotwireless.GetWirelessDeviceStatisticsInput) *GetWirelessDeviceStatisticsFuture

	GetWirelessGateway(ctx workflow.Context, input *iotwireless.GetWirelessGatewayInput) (*iotwireless.GetWirelessGatewayOutput, error)
	GetWirelessGatewayAsync(ctx workflow.Context, input *iotwireless.GetWirelessGatewayInput) *GetWirelessGatewayFuture

	GetWirelessGatewayCertificate(ctx workflow.Context, input *iotwireless.GetWirelessGatewayCertificateInput) (*iotwireless.GetWirelessGatewayCertificateOutput, error)
	GetWirelessGatewayCertificateAsync(ctx workflow.Context, input *iotwireless.GetWirelessGatewayCertificateInput) *GetWirelessGatewayCertificateFuture

	GetWirelessGatewayFirmwareInformation(ctx workflow.Context, input *iotwireless.GetWirelessGatewayFirmwareInformationInput) (*iotwireless.GetWirelessGatewayFirmwareInformationOutput, error)
	GetWirelessGatewayFirmwareInformationAsync(ctx workflow.Context, input *iotwireless.GetWirelessGatewayFirmwareInformationInput) *GetWirelessGatewayFirmwareInformationFuture

	GetWirelessGatewayStatistics(ctx workflow.Context, input *iotwireless.GetWirelessGatewayStatisticsInput) (*iotwireless.GetWirelessGatewayStatisticsOutput, error)
	GetWirelessGatewayStatisticsAsync(ctx workflow.Context, input *iotwireless.GetWirelessGatewayStatisticsInput) *GetWirelessGatewayStatisticsFuture

	GetWirelessGatewayTask(ctx workflow.Context, input *iotwireless.GetWirelessGatewayTaskInput) (*iotwireless.GetWirelessGatewayTaskOutput, error)
	GetWirelessGatewayTaskAsync(ctx workflow.Context, input *iotwireless.GetWirelessGatewayTaskInput) *GetWirelessGatewayTaskFuture

	GetWirelessGatewayTaskDefinition(ctx workflow.Context, input *iotwireless.GetWirelessGatewayTaskDefinitionInput) (*iotwireless.GetWirelessGatewayTaskDefinitionOutput, error)
	GetWirelessGatewayTaskDefinitionAsync(ctx workflow.Context, input *iotwireless.GetWirelessGatewayTaskDefinitionInput) *GetWirelessGatewayTaskDefinitionFuture

	ListDestinations(ctx workflow.Context, input *iotwireless.ListDestinationsInput) (*iotwireless.ListDestinationsOutput, error)
	ListDestinationsAsync(ctx workflow.Context, input *iotwireless.ListDestinationsInput) *ListDestinationsFuture

	ListDeviceProfiles(ctx workflow.Context, input *iotwireless.ListDeviceProfilesInput) (*iotwireless.ListDeviceProfilesOutput, error)
	ListDeviceProfilesAsync(ctx workflow.Context, input *iotwireless.ListDeviceProfilesInput) *ListDeviceProfilesFuture

	ListPartnerAccounts(ctx workflow.Context, input *iotwireless.ListPartnerAccountsInput) (*iotwireless.ListPartnerAccountsOutput, error)
	ListPartnerAccountsAsync(ctx workflow.Context, input *iotwireless.ListPartnerAccountsInput) *ListPartnerAccountsFuture

	ListServiceProfiles(ctx workflow.Context, input *iotwireless.ListServiceProfilesInput) (*iotwireless.ListServiceProfilesOutput, error)
	ListServiceProfilesAsync(ctx workflow.Context, input *iotwireless.ListServiceProfilesInput) *ListServiceProfilesFuture

	ListTagsForResource(ctx workflow.Context, input *iotwireless.ListTagsForResourceInput) (*iotwireless.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iotwireless.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListWirelessDevices(ctx workflow.Context, input *iotwireless.ListWirelessDevicesInput) (*iotwireless.ListWirelessDevicesOutput, error)
	ListWirelessDevicesAsync(ctx workflow.Context, input *iotwireless.ListWirelessDevicesInput) *ListWirelessDevicesFuture

	ListWirelessGatewayTaskDefinitions(ctx workflow.Context, input *iotwireless.ListWirelessGatewayTaskDefinitionsInput) (*iotwireless.ListWirelessGatewayTaskDefinitionsOutput, error)
	ListWirelessGatewayTaskDefinitionsAsync(ctx workflow.Context, input *iotwireless.ListWirelessGatewayTaskDefinitionsInput) *ListWirelessGatewayTaskDefinitionsFuture

	ListWirelessGateways(ctx workflow.Context, input *iotwireless.ListWirelessGatewaysInput) (*iotwireless.ListWirelessGatewaysOutput, error)
	ListWirelessGatewaysAsync(ctx workflow.Context, input *iotwireless.ListWirelessGatewaysInput) *ListWirelessGatewaysFuture

	SendDataToWirelessDevice(ctx workflow.Context, input *iotwireless.SendDataToWirelessDeviceInput) (*iotwireless.SendDataToWirelessDeviceOutput, error)
	SendDataToWirelessDeviceAsync(ctx workflow.Context, input *iotwireless.SendDataToWirelessDeviceInput) *SendDataToWirelessDeviceFuture

	TagResource(ctx workflow.Context, input *iotwireless.TagResourceInput) (*iotwireless.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iotwireless.TagResourceInput) *TagResourceFuture

	TestWirelessDevice(ctx workflow.Context, input *iotwireless.TestWirelessDeviceInput) (*iotwireless.TestWirelessDeviceOutput, error)
	TestWirelessDeviceAsync(ctx workflow.Context, input *iotwireless.TestWirelessDeviceInput) *TestWirelessDeviceFuture

	UntagResource(ctx workflow.Context, input *iotwireless.UntagResourceInput) (*iotwireless.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iotwireless.UntagResourceInput) *UntagResourceFuture

	UpdateDestination(ctx workflow.Context, input *iotwireless.UpdateDestinationInput) (*iotwireless.UpdateDestinationOutput, error)
	UpdateDestinationAsync(ctx workflow.Context, input *iotwireless.UpdateDestinationInput) *UpdateDestinationFuture

	UpdatePartnerAccount(ctx workflow.Context, input *iotwireless.UpdatePartnerAccountInput) (*iotwireless.UpdatePartnerAccountOutput, error)
	UpdatePartnerAccountAsync(ctx workflow.Context, input *iotwireless.UpdatePartnerAccountInput) *UpdatePartnerAccountFuture

	UpdateWirelessDevice(ctx workflow.Context, input *iotwireless.UpdateWirelessDeviceInput) (*iotwireless.UpdateWirelessDeviceOutput, error)
	UpdateWirelessDeviceAsync(ctx workflow.Context, input *iotwireless.UpdateWirelessDeviceInput) *UpdateWirelessDeviceFuture

	UpdateWirelessGateway(ctx workflow.Context, input *iotwireless.UpdateWirelessGatewayInput) (*iotwireless.UpdateWirelessGatewayOutput, error)
	UpdateWirelessGatewayAsync(ctx workflow.Context, input *iotwireless.UpdateWirelessGatewayInput) *UpdateWirelessGatewayFuture
}

func NewClient() Client {
	return &stub{}
}
