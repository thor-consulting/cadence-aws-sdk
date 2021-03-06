// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface"

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
	client cloudhsmiface.CloudHSMAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := cloudhsm.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (cloudhsmiface.CloudHSMAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return cloudhsm.New(sess), nil
}

func (a *Activities) AddTagsToResource(ctx context.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsToResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHapg(ctx context.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHapgWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHsm(ctx context.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateHsmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLunaClient(ctx context.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLunaClientWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteHapg(ctx context.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteHapgWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteHsm(ctx context.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteHsmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLunaClient(ctx context.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLunaClientWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeHapg(ctx context.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeHapgWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeHsm(ctx context.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeHsmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLunaClient(ctx context.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLunaClientWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConfig(ctx context.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAvailableZones(ctx context.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAvailableZonesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHapgs(ctx context.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHapgsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHsms(ctx context.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHsmsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListLunaClients(ctx context.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListLunaClientsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyHapg(ctx context.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyHapgWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyHsm(ctx context.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyHsmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyLunaClient(ctx context.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyLunaClientWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTagsFromResource(ctx context.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsFromResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
