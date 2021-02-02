// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package ecrpublic

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecrpublic"
	"github.com/aws/aws-sdk-go/service/ecrpublic/ecrpubliciface"

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
	client ecrpubliciface.ECRPublicAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := ecrpublic.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (ecrpubliciface.ECRPublicAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return ecrpublic.New(sess), nil
}

func (a *Activities) BatchCheckLayerAvailability(ctx context.Context, input *ecrpublic.BatchCheckLayerAvailabilityInput) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchCheckLayerAvailabilityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDeleteImage(ctx context.Context, input *ecrpublic.BatchDeleteImageInput) (*ecrpublic.BatchDeleteImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDeleteImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CompleteLayerUpload(ctx context.Context, input *ecrpublic.CompleteLayerUploadInput) (*ecrpublic.CompleteLayerUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CompleteLayerUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRepository(ctx context.Context, input *ecrpublic.CreateRepositoryInput) (*ecrpublic.CreateRepositoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRepositoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRepository(ctx context.Context, input *ecrpublic.DeleteRepositoryInput) (*ecrpublic.DeleteRepositoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRepositoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRepositoryPolicy(ctx context.Context, input *ecrpublic.DeleteRepositoryPolicyInput) (*ecrpublic.DeleteRepositoryPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRepositoryPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeImageTags(ctx context.Context, input *ecrpublic.DescribeImageTagsInput) (*ecrpublic.DescribeImageTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeImageTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeImages(ctx context.Context, input *ecrpublic.DescribeImagesInput) (*ecrpublic.DescribeImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeImagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRegistries(ctx context.Context, input *ecrpublic.DescribeRegistriesInput) (*ecrpublic.DescribeRegistriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRegistriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRepositories(ctx context.Context, input *ecrpublic.DescribeRepositoriesInput) (*ecrpublic.DescribeRepositoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRepositoriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAuthorizationToken(ctx context.Context, input *ecrpublic.GetAuthorizationTokenInput) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAuthorizationTokenWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRegistryCatalogData(ctx context.Context, input *ecrpublic.GetRegistryCatalogDataInput) (*ecrpublic.GetRegistryCatalogDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRegistryCatalogDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRepositoryCatalogData(ctx context.Context, input *ecrpublic.GetRepositoryCatalogDataInput) (*ecrpublic.GetRepositoryCatalogDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRepositoryCatalogDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRepositoryPolicy(ctx context.Context, input *ecrpublic.GetRepositoryPolicyInput) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRepositoryPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) InitiateLayerUpload(ctx context.Context, input *ecrpublic.InitiateLayerUploadInput) (*ecrpublic.InitiateLayerUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.InitiateLayerUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutImage(ctx context.Context, input *ecrpublic.PutImageInput) (*ecrpublic.PutImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRegistryCatalogData(ctx context.Context, input *ecrpublic.PutRegistryCatalogDataInput) (*ecrpublic.PutRegistryCatalogDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRegistryCatalogDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRepositoryCatalogData(ctx context.Context, input *ecrpublic.PutRepositoryCatalogDataInput) (*ecrpublic.PutRepositoryCatalogDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRepositoryCatalogDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetRepositoryPolicy(ctx context.Context, input *ecrpublic.SetRepositoryPolicyInput) (*ecrpublic.SetRepositoryPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetRepositoryPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UploadLayerPart(ctx context.Context, input *ecrpublic.UploadLayerPartInput) (*ecrpublic.UploadLayerPartOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UploadLayerPartWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
