// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"

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
	client elbv2iface.ELBV2API

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := elbv2.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (elbv2iface.ELBV2API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return elbv2.New(sess), nil
}

func (a *Activities) AddListenerCertificates(ctx context.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddListenerCertificatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AddTags(ctx context.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateListener(ctx context.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateListenerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLoadBalancer(ctx context.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRule(ctx context.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRuleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTargetGroup(ctx context.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTargetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteListener(ctx context.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteListenerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLoadBalancer(ctx context.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRule(ctx context.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRuleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTargetGroup(ctx context.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTargetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterTargets(ctx context.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterTargetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccountLimits(ctx context.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeListenerCertificates(ctx context.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeListenerCertificatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeListeners(ctx context.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeListenersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancerAttributes(ctx context.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancerAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancers(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRules(ctx context.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSSLPolicies(ctx context.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSSLPoliciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTags(ctx context.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTargetGroupAttributes(ctx context.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTargetGroupAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTargetGroups(ctx context.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTargetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTargetHealth(ctx context.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTargetHealthWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyListener(ctx context.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyListenerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyLoadBalancerAttributes(ctx context.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyLoadBalancerAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyRule(ctx context.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyRuleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyTargetGroup(ctx context.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyTargetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyTargetGroupAttributes(ctx context.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyTargetGroupAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterTargets(ctx context.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterTargetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveListenerCertificates(ctx context.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveListenerCertificatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTags(ctx context.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetIpAddressType(ctx context.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetIpAddressTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetRulePriorities(ctx context.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetRulePrioritiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetSecurityGroups(ctx context.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetSecurityGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetSubnets(ctx context.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetSubnetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilLoadBalancerAvailable(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilLoadBalancerAvailableWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilLoadBalancerExists(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilLoadBalancerExistsWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilLoadBalancersDeleted(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilLoadBalancersDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilTargetDeregistered(ctx context.Context, input *elbv2.DescribeTargetHealthInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTargetDeregisteredWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilTargetInService(ctx context.Context, input *elbv2.DescribeTargetHealthInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTargetInServiceWithContext(ctx, input, options...))
	})
}