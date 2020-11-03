// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/groundstation/groundstationiface"

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
	client groundstationiface.GroundStationAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := groundstation.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (groundstationiface.GroundStationAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return groundstation.New(sess), nil
}

func (a *Activities) CancelContact(ctx context.Context, input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateConfig(ctx context.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDataflowEndpointGroup(ctx context.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDataflowEndpointGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMissionProfile(ctx context.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMissionProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConfig(ctx context.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDataflowEndpointGroup(ctx context.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDataflowEndpointGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMissionProfile(ctx context.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMissionProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeContact(ctx context.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConfig(ctx context.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDataflowEndpointGroup(ctx context.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDataflowEndpointGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMinuteUsage(ctx context.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMinuteUsageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMissionProfile(ctx context.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMissionProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSatellite(ctx context.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSatelliteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListConfigs(ctx context.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListConfigsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListContacts(ctx context.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListContactsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDataflowEndpointGroups(ctx context.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDataflowEndpointGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGroundStations(ctx context.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGroundStationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMissionProfiles(ctx context.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMissionProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSatellites(ctx context.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSatellitesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ReserveContact(ctx context.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ReserveContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateConfig(ctx context.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMissionProfile(ctx context.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMissionProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
