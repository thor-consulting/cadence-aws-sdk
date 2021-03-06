// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package customerprofilesstub

import (
	"github.com/aws/aws-sdk-go/service/customerprofiles"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddProfileKey(ctx workflow.Context, input *customerprofiles.AddProfileKeyInput) (*customerprofiles.AddProfileKeyOutput, error)
	AddProfileKeyAsync(ctx workflow.Context, input *customerprofiles.AddProfileKeyInput) *AddProfileKeyFuture

	CreateDomain(ctx workflow.Context, input *customerprofiles.CreateDomainInput) (*customerprofiles.CreateDomainOutput, error)
	CreateDomainAsync(ctx workflow.Context, input *customerprofiles.CreateDomainInput) *CreateDomainFuture

	CreateProfile(ctx workflow.Context, input *customerprofiles.CreateProfileInput) (*customerprofiles.CreateProfileOutput, error)
	CreateProfileAsync(ctx workflow.Context, input *customerprofiles.CreateProfileInput) *CreateProfileFuture

	DeleteDomain(ctx workflow.Context, input *customerprofiles.DeleteDomainInput) (*customerprofiles.DeleteDomainOutput, error)
	DeleteDomainAsync(ctx workflow.Context, input *customerprofiles.DeleteDomainInput) *DeleteDomainFuture

	DeleteIntegration(ctx workflow.Context, input *customerprofiles.DeleteIntegrationInput) (*customerprofiles.DeleteIntegrationOutput, error)
	DeleteIntegrationAsync(ctx workflow.Context, input *customerprofiles.DeleteIntegrationInput) *DeleteIntegrationFuture

	DeleteProfile(ctx workflow.Context, input *customerprofiles.DeleteProfileInput) (*customerprofiles.DeleteProfileOutput, error)
	DeleteProfileAsync(ctx workflow.Context, input *customerprofiles.DeleteProfileInput) *DeleteProfileFuture

	DeleteProfileKey(ctx workflow.Context, input *customerprofiles.DeleteProfileKeyInput) (*customerprofiles.DeleteProfileKeyOutput, error)
	DeleteProfileKeyAsync(ctx workflow.Context, input *customerprofiles.DeleteProfileKeyInput) *DeleteProfileKeyFuture

	DeleteProfileObject(ctx workflow.Context, input *customerprofiles.DeleteProfileObjectInput) (*customerprofiles.DeleteProfileObjectOutput, error)
	DeleteProfileObjectAsync(ctx workflow.Context, input *customerprofiles.DeleteProfileObjectInput) *DeleteProfileObjectFuture

	DeleteProfileObjectType(ctx workflow.Context, input *customerprofiles.DeleteProfileObjectTypeInput) (*customerprofiles.DeleteProfileObjectTypeOutput, error)
	DeleteProfileObjectTypeAsync(ctx workflow.Context, input *customerprofiles.DeleteProfileObjectTypeInput) *DeleteProfileObjectTypeFuture

	GetDomain(ctx workflow.Context, input *customerprofiles.GetDomainInput) (*customerprofiles.GetDomainOutput, error)
	GetDomainAsync(ctx workflow.Context, input *customerprofiles.GetDomainInput) *GetDomainFuture

	GetIntegration(ctx workflow.Context, input *customerprofiles.GetIntegrationInput) (*customerprofiles.GetIntegrationOutput, error)
	GetIntegrationAsync(ctx workflow.Context, input *customerprofiles.GetIntegrationInput) *GetIntegrationFuture

	GetProfileObjectType(ctx workflow.Context, input *customerprofiles.GetProfileObjectTypeInput) (*customerprofiles.GetProfileObjectTypeOutput, error)
	GetProfileObjectTypeAsync(ctx workflow.Context, input *customerprofiles.GetProfileObjectTypeInput) *GetProfileObjectTypeFuture

	GetProfileObjectTypeTemplate(ctx workflow.Context, input *customerprofiles.GetProfileObjectTypeTemplateInput) (*customerprofiles.GetProfileObjectTypeTemplateOutput, error)
	GetProfileObjectTypeTemplateAsync(ctx workflow.Context, input *customerprofiles.GetProfileObjectTypeTemplateInput) *GetProfileObjectTypeTemplateFuture

	ListAccountIntegrations(ctx workflow.Context, input *customerprofiles.ListAccountIntegrationsInput) (*customerprofiles.ListAccountIntegrationsOutput, error)
	ListAccountIntegrationsAsync(ctx workflow.Context, input *customerprofiles.ListAccountIntegrationsInput) *ListAccountIntegrationsFuture

	ListDomains(ctx workflow.Context, input *customerprofiles.ListDomainsInput) (*customerprofiles.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *customerprofiles.ListDomainsInput) *ListDomainsFuture

	ListIntegrations(ctx workflow.Context, input *customerprofiles.ListIntegrationsInput) (*customerprofiles.ListIntegrationsOutput, error)
	ListIntegrationsAsync(ctx workflow.Context, input *customerprofiles.ListIntegrationsInput) *ListIntegrationsFuture

	ListProfileObjectTypeTemplates(ctx workflow.Context, input *customerprofiles.ListProfileObjectTypeTemplatesInput) (*customerprofiles.ListProfileObjectTypeTemplatesOutput, error)
	ListProfileObjectTypeTemplatesAsync(ctx workflow.Context, input *customerprofiles.ListProfileObjectTypeTemplatesInput) *ListProfileObjectTypeTemplatesFuture

	ListProfileObjectTypes(ctx workflow.Context, input *customerprofiles.ListProfileObjectTypesInput) (*customerprofiles.ListProfileObjectTypesOutput, error)
	ListProfileObjectTypesAsync(ctx workflow.Context, input *customerprofiles.ListProfileObjectTypesInput) *ListProfileObjectTypesFuture

	ListProfileObjects(ctx workflow.Context, input *customerprofiles.ListProfileObjectsInput) (*customerprofiles.ListProfileObjectsOutput, error)
	ListProfileObjectsAsync(ctx workflow.Context, input *customerprofiles.ListProfileObjectsInput) *ListProfileObjectsFuture

	ListTagsForResource(ctx workflow.Context, input *customerprofiles.ListTagsForResourceInput) (*customerprofiles.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *customerprofiles.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutIntegration(ctx workflow.Context, input *customerprofiles.PutIntegrationInput) (*customerprofiles.PutIntegrationOutput, error)
	PutIntegrationAsync(ctx workflow.Context, input *customerprofiles.PutIntegrationInput) *PutIntegrationFuture

	PutProfileObject(ctx workflow.Context, input *customerprofiles.PutProfileObjectInput) (*customerprofiles.PutProfileObjectOutput, error)
	PutProfileObjectAsync(ctx workflow.Context, input *customerprofiles.PutProfileObjectInput) *PutProfileObjectFuture

	PutProfileObjectType(ctx workflow.Context, input *customerprofiles.PutProfileObjectTypeInput) (*customerprofiles.PutProfileObjectTypeOutput, error)
	PutProfileObjectTypeAsync(ctx workflow.Context, input *customerprofiles.PutProfileObjectTypeInput) *PutProfileObjectTypeFuture

	SearchProfiles(ctx workflow.Context, input *customerprofiles.SearchProfilesInput) (*customerprofiles.SearchProfilesOutput, error)
	SearchProfilesAsync(ctx workflow.Context, input *customerprofiles.SearchProfilesInput) *SearchProfilesFuture

	TagResource(ctx workflow.Context, input *customerprofiles.TagResourceInput) (*customerprofiles.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *customerprofiles.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *customerprofiles.UntagResourceInput) (*customerprofiles.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *customerprofiles.UntagResourceInput) *UntagResourceFuture

	UpdateDomain(ctx workflow.Context, input *customerprofiles.UpdateDomainInput) (*customerprofiles.UpdateDomainOutput, error)
	UpdateDomainAsync(ctx workflow.Context, input *customerprofiles.UpdateDomainInput) *UpdateDomainFuture

	UpdateProfile(ctx workflow.Context, input *customerprofiles.UpdateProfileInput) (*customerprofiles.UpdateProfileOutput, error)
	UpdateProfileAsync(ctx workflow.Context, input *customerprofiles.UpdateProfileInput) *UpdateProfileFuture
}

func NewClient() Client {
	return &stub{}
}
