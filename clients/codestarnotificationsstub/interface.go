// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package codestarnotificationsstub

import (
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateNotificationRule(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error)
	CreateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) *CodeStarNotificationsCreateNotificationRuleFuture

	DeleteNotificationRule(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error)
	DeleteNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) *CodeStarNotificationsDeleteNotificationRuleFuture

	DeleteTarget(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error)
	DeleteTargetAsync(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) *CodeStarNotificationsDeleteTargetFuture

	DescribeNotificationRule(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error)
	DescribeNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) *CodeStarNotificationsDescribeNotificationRuleFuture

	ListEventTypes(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error)
	ListEventTypesAsync(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) *CodeStarNotificationsListEventTypesFuture

	ListNotificationRules(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error)
	ListNotificationRulesAsync(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) *CodeStarNotificationsListNotificationRulesFuture

	ListTagsForResource(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) *CodeStarNotificationsListTagsForResourceFuture

	ListTargets(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error)
	ListTargetsAsync(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) *CodeStarNotificationsListTargetsFuture

	Subscribe(ctx workflow.Context, input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error)
	SubscribeAsync(ctx workflow.Context, input *codestarnotifications.SubscribeInput) *CodeStarNotificationsSubscribeFuture

	TagResource(ctx workflow.Context, input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codestarnotifications.TagResourceInput) *CodeStarNotificationsTagResourceFuture

	Unsubscribe(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error)
	UnsubscribeAsync(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) *CodeStarNotificationsUnsubscribeFuture

	UntagResource(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) *CodeStarNotificationsUntagResourceFuture

	UpdateNotificationRule(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error)
	UpdateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) *CodeStarNotificationsUpdateNotificationRuleFuture
}

func NewClient() Client {
	return &stub{}
}
