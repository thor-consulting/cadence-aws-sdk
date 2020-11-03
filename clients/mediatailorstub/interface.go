// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package mediatailorstub

import (
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeletePlaybackConfiguration(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error)
	DeletePlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) *MediaTailorDeletePlaybackConfigurationFuture

	GetPlaybackConfiguration(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error)
	GetPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) *MediaTailorGetPlaybackConfigurationFuture

	ListPlaybackConfigurations(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error)
	ListPlaybackConfigurationsAsync(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) *MediaTailorListPlaybackConfigurationsFuture

	ListTagsForResource(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) *MediaTailorListTagsForResourceFuture

	PutPlaybackConfiguration(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error)
	PutPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) *MediaTailorPutPlaybackConfigurationFuture

	TagResource(ctx workflow.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediatailor.TagResourceInput) *MediaTailorTagResourceFuture

	UntagResource(ctx workflow.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediatailor.UntagResourceInput) *MediaTailorUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
