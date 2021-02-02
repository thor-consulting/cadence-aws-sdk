// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package licensemanagerstub

import (
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptGrant(ctx workflow.Context, input *licensemanager.AcceptGrantInput) (*licensemanager.AcceptGrantOutput, error)
	AcceptGrantAsync(ctx workflow.Context, input *licensemanager.AcceptGrantInput) *AcceptGrantFuture

	CheckInLicense(ctx workflow.Context, input *licensemanager.CheckInLicenseInput) (*licensemanager.CheckInLicenseOutput, error)
	CheckInLicenseAsync(ctx workflow.Context, input *licensemanager.CheckInLicenseInput) *CheckInLicenseFuture

	CheckoutBorrowLicense(ctx workflow.Context, input *licensemanager.CheckoutBorrowLicenseInput) (*licensemanager.CheckoutBorrowLicenseOutput, error)
	CheckoutBorrowLicenseAsync(ctx workflow.Context, input *licensemanager.CheckoutBorrowLicenseInput) *CheckoutBorrowLicenseFuture

	CheckoutLicense(ctx workflow.Context, input *licensemanager.CheckoutLicenseInput) (*licensemanager.CheckoutLicenseOutput, error)
	CheckoutLicenseAsync(ctx workflow.Context, input *licensemanager.CheckoutLicenseInput) *CheckoutLicenseFuture

	CreateGrant(ctx workflow.Context, input *licensemanager.CreateGrantInput) (*licensemanager.CreateGrantOutput, error)
	CreateGrantAsync(ctx workflow.Context, input *licensemanager.CreateGrantInput) *CreateGrantFuture

	CreateGrantVersion(ctx workflow.Context, input *licensemanager.CreateGrantVersionInput) (*licensemanager.CreateGrantVersionOutput, error)
	CreateGrantVersionAsync(ctx workflow.Context, input *licensemanager.CreateGrantVersionInput) *CreateGrantVersionFuture

	CreateLicense(ctx workflow.Context, input *licensemanager.CreateLicenseInput) (*licensemanager.CreateLicenseOutput, error)
	CreateLicenseAsync(ctx workflow.Context, input *licensemanager.CreateLicenseInput) *CreateLicenseFuture

	CreateLicenseConfiguration(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error)
	CreateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) *CreateLicenseConfigurationFuture

	CreateLicenseVersion(ctx workflow.Context, input *licensemanager.CreateLicenseVersionInput) (*licensemanager.CreateLicenseVersionOutput, error)
	CreateLicenseVersionAsync(ctx workflow.Context, input *licensemanager.CreateLicenseVersionInput) *CreateLicenseVersionFuture

	CreateToken(ctx workflow.Context, input *licensemanager.CreateTokenInput) (*licensemanager.CreateTokenOutput, error)
	CreateTokenAsync(ctx workflow.Context, input *licensemanager.CreateTokenInput) *CreateTokenFuture

	DeleteGrant(ctx workflow.Context, input *licensemanager.DeleteGrantInput) (*licensemanager.DeleteGrantOutput, error)
	DeleteGrantAsync(ctx workflow.Context, input *licensemanager.DeleteGrantInput) *DeleteGrantFuture

	DeleteLicense(ctx workflow.Context, input *licensemanager.DeleteLicenseInput) (*licensemanager.DeleteLicenseOutput, error)
	DeleteLicenseAsync(ctx workflow.Context, input *licensemanager.DeleteLicenseInput) *DeleteLicenseFuture

	DeleteLicenseConfiguration(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error)
	DeleteLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) *DeleteLicenseConfigurationFuture

	DeleteToken(ctx workflow.Context, input *licensemanager.DeleteTokenInput) (*licensemanager.DeleteTokenOutput, error)
	DeleteTokenAsync(ctx workflow.Context, input *licensemanager.DeleteTokenInput) *DeleteTokenFuture

	ExtendLicenseConsumption(ctx workflow.Context, input *licensemanager.ExtendLicenseConsumptionInput) (*licensemanager.ExtendLicenseConsumptionOutput, error)
	ExtendLicenseConsumptionAsync(ctx workflow.Context, input *licensemanager.ExtendLicenseConsumptionInput) *ExtendLicenseConsumptionFuture

	GetAccessToken(ctx workflow.Context, input *licensemanager.GetAccessTokenInput) (*licensemanager.GetAccessTokenOutput, error)
	GetAccessTokenAsync(ctx workflow.Context, input *licensemanager.GetAccessTokenInput) *GetAccessTokenFuture

	GetGrant(ctx workflow.Context, input *licensemanager.GetGrantInput) (*licensemanager.GetGrantOutput, error)
	GetGrantAsync(ctx workflow.Context, input *licensemanager.GetGrantInput) *GetGrantFuture

	GetLicense(ctx workflow.Context, input *licensemanager.GetLicenseInput) (*licensemanager.GetLicenseOutput, error)
	GetLicenseAsync(ctx workflow.Context, input *licensemanager.GetLicenseInput) *GetLicenseFuture

	GetLicenseConfiguration(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error)
	GetLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) *GetLicenseConfigurationFuture

	GetLicenseUsage(ctx workflow.Context, input *licensemanager.GetLicenseUsageInput) (*licensemanager.GetLicenseUsageOutput, error)
	GetLicenseUsageAsync(ctx workflow.Context, input *licensemanager.GetLicenseUsageInput) *GetLicenseUsageFuture

	GetServiceSettings(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error)
	GetServiceSettingsAsync(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) *GetServiceSettingsFuture

	ListAssociationsForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
	ListAssociationsForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) *ListAssociationsForLicenseConfigurationFuture

	ListDistributedGrants(ctx workflow.Context, input *licensemanager.ListDistributedGrantsInput) (*licensemanager.ListDistributedGrantsOutput, error)
	ListDistributedGrantsAsync(ctx workflow.Context, input *licensemanager.ListDistributedGrantsInput) *ListDistributedGrantsFuture

	ListFailuresForLicenseConfigurationOperations(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
	ListFailuresForLicenseConfigurationOperationsAsync(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) *ListFailuresForLicenseConfigurationOperationsFuture

	ListLicenseConfigurations(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error)
	ListLicenseConfigurationsAsync(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) *ListLicenseConfigurationsFuture

	ListLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
	ListLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) *ListLicenseSpecificationsForResourceFuture

	ListLicenseVersions(ctx workflow.Context, input *licensemanager.ListLicenseVersionsInput) (*licensemanager.ListLicenseVersionsOutput, error)
	ListLicenseVersionsAsync(ctx workflow.Context, input *licensemanager.ListLicenseVersionsInput) *ListLicenseVersionsFuture

	ListLicenses(ctx workflow.Context, input *licensemanager.ListLicensesInput) (*licensemanager.ListLicensesOutput, error)
	ListLicensesAsync(ctx workflow.Context, input *licensemanager.ListLicensesInput) *ListLicensesFuture

	ListReceivedGrants(ctx workflow.Context, input *licensemanager.ListReceivedGrantsInput) (*licensemanager.ListReceivedGrantsOutput, error)
	ListReceivedGrantsAsync(ctx workflow.Context, input *licensemanager.ListReceivedGrantsInput) *ListReceivedGrantsFuture

	ListReceivedLicenses(ctx workflow.Context, input *licensemanager.ListReceivedLicensesInput) (*licensemanager.ListReceivedLicensesOutput, error)
	ListReceivedLicensesAsync(ctx workflow.Context, input *licensemanager.ListReceivedLicensesInput) *ListReceivedLicensesFuture

	ListResourceInventory(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error)
	ListResourceInventoryAsync(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) *ListResourceInventoryFuture

	ListTagsForResource(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTokens(ctx workflow.Context, input *licensemanager.ListTokensInput) (*licensemanager.ListTokensOutput, error)
	ListTokensAsync(ctx workflow.Context, input *licensemanager.ListTokensInput) *ListTokensFuture

	ListUsageForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
	ListUsageForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) *ListUsageForLicenseConfigurationFuture

	RejectGrant(ctx workflow.Context, input *licensemanager.RejectGrantInput) (*licensemanager.RejectGrantOutput, error)
	RejectGrantAsync(ctx workflow.Context, input *licensemanager.RejectGrantInput) *RejectGrantFuture

	TagResource(ctx workflow.Context, input *licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *licensemanager.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *licensemanager.UntagResourceInput) *UntagResourceFuture

	UpdateLicenseConfiguration(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error)
	UpdateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) *UpdateLicenseConfigurationFuture

	UpdateLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
	UpdateLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) *UpdateLicenseSpecificationsForResourceFuture

	UpdateServiceSettings(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error)
	UpdateServiceSettingsAsync(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) *UpdateServiceSettingsFuture
}

func NewClient() Client {
	return &stub{}
}
