// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mediatailor

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"

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
	client mediatailoriface.MediaTailorAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mediatailor.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mediatailoriface.MediaTailorAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return mediatailor.New(sess), nil
}

func (a *Activities) DeletePlaybackConfiguration(ctx context.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePlaybackConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPlaybackConfiguration(ctx context.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPlaybackConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPlaybackConfigurations(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPlaybackConfigurationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutPlaybackConfiguration(ctx context.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutPlaybackConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
