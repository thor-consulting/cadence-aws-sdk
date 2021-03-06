// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client worklinkiface.WorkLinkAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := worklink.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (worklinkiface.WorkLinkAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return worklink.New(sess), nil
}

func (a *Activities) AssociateDomain(ctx context.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateWebsiteAuthorizationProvider(ctx context.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateWebsiteAuthorizationProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateWebsiteCertificateAuthority(ctx context.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateWebsiteCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFleet(ctx context.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFleet(ctx context.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAuditStreamConfiguration(ctx context.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAuditStreamConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCompanyNetworkConfiguration(ctx context.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCompanyNetworkConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDevice(ctx context.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDevicePolicyConfiguration(ctx context.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDevicePolicyConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDomain(ctx context.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleetMetadata(ctx context.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeIdentityProviderConfiguration(ctx context.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeIdentityProviderConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeWebsiteCertificateAuthority(ctx context.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeWebsiteCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateDomain(ctx context.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateWebsiteAuthorizationProvider(ctx context.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateWebsiteAuthorizationProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateWebsiteCertificateAuthority(ctx context.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateWebsiteCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDevices(ctx context.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDevicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDomains(ctx context.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDomainsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFleets(ctx context.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFleetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListWebsiteAuthorizationProviders(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListWebsiteAuthorizationProvidersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListWebsiteCertificateAuthorities(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListWebsiteCertificateAuthoritiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreDomainAccess(ctx context.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreDomainAccessWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RevokeDomainAccess(ctx context.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RevokeDomainAccessWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SignOutUser(ctx context.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SignOutUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateAuditStreamConfiguration(ctx context.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateAuditStreamConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateCompanyNetworkConfiguration(ctx context.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateCompanyNetworkConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDevicePolicyConfiguration(ctx context.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDevicePolicyConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDomainMetadata(ctx context.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDomainMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFleetMetadata(ctx context.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFleetMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateIdentityProviderConfiguration(ctx context.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateIdentityProviderConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
