// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"

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
	client neptuneiface.NeptuneAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := neptune.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (neptuneiface.NeptuneAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return neptune.New(sess), nil
}

func (a *Activities) AddRoleToDBCluster(ctx context.Context, input *neptune.AddRoleToDBClusterInput) (*neptune.AddRoleToDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddRoleToDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AddSourceIdentifierToSubscription(ctx context.Context, input *neptune.AddSourceIdentifierToSubscriptionInput) (*neptune.AddSourceIdentifierToSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddSourceIdentifierToSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AddTagsToResource(ctx context.Context, input *neptune.AddTagsToResourceInput) (*neptune.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsToResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ApplyPendingMaintenanceAction(ctx context.Context, input *neptune.ApplyPendingMaintenanceActionInput) (*neptune.ApplyPendingMaintenanceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ApplyPendingMaintenanceActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CopyDBClusterParameterGroup(ctx context.Context, input *neptune.CopyDBClusterParameterGroupInput) (*neptune.CopyDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyDBClusterParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CopyDBClusterSnapshot(ctx context.Context, input *neptune.CopyDBClusterSnapshotInput) (*neptune.CopyDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyDBClusterSnapshotWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CopyDBParameterGroup(ctx context.Context, input *neptune.CopyDBParameterGroupInput) (*neptune.CopyDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyDBParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBCluster(ctx context.Context, input *neptune.CreateDBClusterInput) (*neptune.CreateDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBClusterEndpoint(ctx context.Context, input *neptune.CreateDBClusterEndpointInput) (*neptune.CreateDBClusterEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBClusterEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBClusterParameterGroup(ctx context.Context, input *neptune.CreateDBClusterParameterGroupInput) (*neptune.CreateDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBClusterParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBClusterSnapshot(ctx context.Context, input *neptune.CreateDBClusterSnapshotInput) (*neptune.CreateDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBClusterSnapshotWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBInstance(ctx context.Context, input *neptune.CreateDBInstanceInput) (*neptune.CreateDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBParameterGroup(ctx context.Context, input *neptune.CreateDBParameterGroupInput) (*neptune.CreateDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDBSubnetGroup(ctx context.Context, input *neptune.CreateDBSubnetGroupInput) (*neptune.CreateDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDBSubnetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEventSubscription(ctx context.Context, input *neptune.CreateEventSubscriptionInput) (*neptune.CreateEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEventSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBCluster(ctx context.Context, input *neptune.DeleteDBClusterInput) (*neptune.DeleteDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBClusterEndpoint(ctx context.Context, input *neptune.DeleteDBClusterEndpointInput) (*neptune.DeleteDBClusterEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBClusterEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBClusterParameterGroup(ctx context.Context, input *neptune.DeleteDBClusterParameterGroupInput) (*neptune.DeleteDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBClusterParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBClusterSnapshot(ctx context.Context, input *neptune.DeleteDBClusterSnapshotInput) (*neptune.DeleteDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBClusterSnapshotWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBInstance(ctx context.Context, input *neptune.DeleteDBInstanceInput) (*neptune.DeleteDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBParameterGroup(ctx context.Context, input *neptune.DeleteDBParameterGroupInput) (*neptune.DeleteDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDBSubnetGroup(ctx context.Context, input *neptune.DeleteDBSubnetGroupInput) (*neptune.DeleteDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDBSubnetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEventSubscription(ctx context.Context, input *neptune.DeleteEventSubscriptionInput) (*neptune.DeleteEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEventSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBClusterEndpoints(ctx context.Context, input *neptune.DescribeDBClusterEndpointsInput) (*neptune.DescribeDBClusterEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBClusterEndpointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBClusterParameterGroups(ctx context.Context, input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBClusterParameterGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBClusterParameters(ctx context.Context, input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBClusterParametersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBClusterSnapshotAttributes(ctx context.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBClusterSnapshotAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBClusterSnapshots(ctx context.Context, input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBClusterSnapshotsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBClusters(ctx context.Context, input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBClustersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBEngineVersions(ctx context.Context, input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBEngineVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBInstances(ctx context.Context, input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBParameterGroups(ctx context.Context, input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBParameterGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBParameters(ctx context.Context, input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBParametersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDBSubnetGroups(ctx context.Context, input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDBSubnetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEngineDefaultClusterParameters(ctx context.Context, input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEngineDefaultClusterParametersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEngineDefaultParameters(ctx context.Context, input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEngineDefaultParametersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEventCategories(ctx context.Context, input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventCategoriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEventSubscriptions(ctx context.Context, input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventSubscriptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeOrderableDBInstanceOptions(ctx context.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePendingMaintenanceActions(ctx context.Context, input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePendingMaintenanceActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeValidDBInstanceModifications(ctx context.Context, input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeValidDBInstanceModificationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) FailoverDBCluster(ctx context.Context, input *neptune.FailoverDBClusterInput) (*neptune.FailoverDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.FailoverDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBCluster(ctx context.Context, input *neptune.ModifyDBClusterInput) (*neptune.ModifyDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBClusterEndpoint(ctx context.Context, input *neptune.ModifyDBClusterEndpointInput) (*neptune.ModifyDBClusterEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBClusterEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBClusterParameterGroup(ctx context.Context, input *neptune.ModifyDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBClusterParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBClusterSnapshotAttribute(ctx context.Context, input *neptune.ModifyDBClusterSnapshotAttributeInput) (*neptune.ModifyDBClusterSnapshotAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBClusterSnapshotAttributeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBInstance(ctx context.Context, input *neptune.ModifyDBInstanceInput) (*neptune.ModifyDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBParameterGroup(ctx context.Context, input *neptune.ModifyDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyDBSubnetGroup(ctx context.Context, input *neptune.ModifyDBSubnetGroupInput) (*neptune.ModifyDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyDBSubnetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyEventSubscription(ctx context.Context, input *neptune.ModifyEventSubscriptionInput) (*neptune.ModifyEventSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyEventSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PromoteReadReplicaDBCluster(ctx context.Context, input *neptune.PromoteReadReplicaDBClusterInput) (*neptune.PromoteReadReplicaDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PromoteReadReplicaDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RebootDBInstance(ctx context.Context, input *neptune.RebootDBInstanceInput) (*neptune.RebootDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RebootDBInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveRoleFromDBCluster(ctx context.Context, input *neptune.RemoveRoleFromDBClusterInput) (*neptune.RemoveRoleFromDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveRoleFromDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveSourceIdentifierFromSubscription(ctx context.Context, input *neptune.RemoveSourceIdentifierFromSubscriptionInput) (*neptune.RemoveSourceIdentifierFromSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveSourceIdentifierFromSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTagsFromResource(ctx context.Context, input *neptune.RemoveTagsFromResourceInput) (*neptune.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsFromResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResetDBClusterParameterGroup(ctx context.Context, input *neptune.ResetDBClusterParameterGroupInput) (*neptune.ResetDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResetDBClusterParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResetDBParameterGroup(ctx context.Context, input *neptune.ResetDBParameterGroupInput) (*neptune.ResetDBParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResetDBParameterGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreDBClusterFromSnapshot(ctx context.Context, input *neptune.RestoreDBClusterFromSnapshotInput) (*neptune.RestoreDBClusterFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreDBClusterFromSnapshotWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreDBClusterToPointInTime(ctx context.Context, input *neptune.RestoreDBClusterToPointInTimeInput) (*neptune.RestoreDBClusterToPointInTimeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreDBClusterToPointInTimeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartDBCluster(ctx context.Context, input *neptune.StartDBClusterInput) (*neptune.StartDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopDBCluster(ctx context.Context, input *neptune.StopDBClusterInput) (*neptune.StopDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopDBClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilDBInstanceAvailable(ctx context.Context, input *neptune.DescribeDBInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilDBInstanceAvailableWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilDBInstanceDeleted(ctx context.Context, input *neptune.DescribeDBInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilDBInstanceDeletedWithContext(ctx, input, options...))
	})
}
