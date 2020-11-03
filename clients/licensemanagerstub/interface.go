// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package licensemanagerstub

import (
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateLicenseConfiguration(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error)
	CreateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) *LicenseManagerCreateLicenseConfigurationFuture

	DeleteLicenseConfiguration(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error)
	DeleteLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) *LicenseManagerDeleteLicenseConfigurationFuture

	GetLicenseConfiguration(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error)
	GetLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) *LicenseManagerGetLicenseConfigurationFuture

	GetServiceSettings(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error)
	GetServiceSettingsAsync(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) *LicenseManagerGetServiceSettingsFuture

	ListAssociationsForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
	ListAssociationsForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) *LicenseManagerListAssociationsForLicenseConfigurationFuture

	ListFailuresForLicenseConfigurationOperations(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
	ListFailuresForLicenseConfigurationOperationsAsync(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) *LicenseManagerListFailuresForLicenseConfigurationOperationsFuture

	ListLicenseConfigurations(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error)
	ListLicenseConfigurationsAsync(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) *LicenseManagerListLicenseConfigurationsFuture

	ListLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
	ListLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) *LicenseManagerListLicenseSpecificationsForResourceFuture

	ListResourceInventory(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error)
	ListResourceInventoryAsync(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) *LicenseManagerListResourceInventoryFuture

	ListTagsForResource(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) *LicenseManagerListTagsForResourceFuture

	ListUsageForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
	ListUsageForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) *LicenseManagerListUsageForLicenseConfigurationFuture

	TagResource(ctx workflow.Context, input *licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *licensemanager.TagResourceInput) *LicenseManagerTagResourceFuture

	UntagResource(ctx workflow.Context, input *licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *licensemanager.UntagResourceInput) *LicenseManagerUntagResourceFuture

	UpdateLicenseConfiguration(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error)
	UpdateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) *LicenseManagerUpdateLicenseConfigurationFuture

	UpdateLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
	UpdateLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) *LicenseManagerUpdateLicenseSpecificationsForResourceFuture

	UpdateServiceSettings(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error)
	UpdateServiceSettingsAsync(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) *LicenseManagerUpdateServiceSettingsFuture
}

func NewClient() Client {
	return &stub{}
}
