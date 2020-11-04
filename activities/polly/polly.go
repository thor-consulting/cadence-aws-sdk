// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/polly/pollyiface"

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
	client pollyiface.PollyAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := polly.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (pollyiface.PollyAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return polly.New(sess), nil
}

func (a *Activities) DeleteLexicon(ctx context.Context, input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLexiconWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVoices(ctx context.Context, input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVoicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetLexicon(ctx context.Context, input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetLexiconWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSpeechSynthesisTask(ctx context.Context, input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSpeechSynthesisTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListLexicons(ctx context.Context, input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListLexiconsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSpeechSynthesisTasks(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSpeechSynthesisTasksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutLexicon(ctx context.Context, input *polly.PutLexiconInput) (*polly.PutLexiconOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutLexiconWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartSpeechSynthesisTask(ctx context.Context, input *polly.StartSpeechSynthesisTaskInput) (*polly.StartSpeechSynthesisTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartSpeechSynthesisTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SynthesizeSpeech(ctx context.Context, input *polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SynthesizeSpeechWithContext(ctx, input)

	return output, internal.EncodeError(err)
}