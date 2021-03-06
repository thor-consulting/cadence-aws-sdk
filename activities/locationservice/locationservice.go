// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package locationservice

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/locationservice"
	"github.com/aws/aws-sdk-go/service/locationservice/locationserviceiface"

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
	client locationserviceiface.LocationServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := locationservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (locationserviceiface.LocationServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return locationservice.New(sess), nil
}

func (a *Activities) AssociateTrackerConsumer(ctx context.Context, input *locationservice.AssociateTrackerConsumerInput) (*locationservice.AssociateTrackerConsumerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateTrackerConsumerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDeleteGeofence(ctx context.Context, input *locationservice.BatchDeleteGeofenceInput) (*locationservice.BatchDeleteGeofenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDeleteGeofenceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchEvaluateGeofences(ctx context.Context, input *locationservice.BatchEvaluateGeofencesInput) (*locationservice.BatchEvaluateGeofencesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchEvaluateGeofencesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetDevicePosition(ctx context.Context, input *locationservice.BatchGetDevicePositionInput) (*locationservice.BatchGetDevicePositionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetDevicePositionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchPutGeofence(ctx context.Context, input *locationservice.BatchPutGeofenceInput) (*locationservice.BatchPutGeofenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchPutGeofenceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchUpdateDevicePosition(ctx context.Context, input *locationservice.BatchUpdateDevicePositionInput) (*locationservice.BatchUpdateDevicePositionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchUpdateDevicePositionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGeofenceCollection(ctx context.Context, input *locationservice.CreateGeofenceCollectionInput) (*locationservice.CreateGeofenceCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGeofenceCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMap(ctx context.Context, input *locationservice.CreateMapInput) (*locationservice.CreateMapOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMapWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePlaceIndex(ctx context.Context, input *locationservice.CreatePlaceIndexInput) (*locationservice.CreatePlaceIndexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePlaceIndexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTracker(ctx context.Context, input *locationservice.CreateTrackerInput) (*locationservice.CreateTrackerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTrackerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGeofenceCollection(ctx context.Context, input *locationservice.DeleteGeofenceCollectionInput) (*locationservice.DeleteGeofenceCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGeofenceCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMap(ctx context.Context, input *locationservice.DeleteMapInput) (*locationservice.DeleteMapOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMapWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePlaceIndex(ctx context.Context, input *locationservice.DeletePlaceIndexInput) (*locationservice.DeletePlaceIndexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePlaceIndexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTracker(ctx context.Context, input *locationservice.DeleteTrackerInput) (*locationservice.DeleteTrackerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTrackerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGeofenceCollection(ctx context.Context, input *locationservice.DescribeGeofenceCollectionInput) (*locationservice.DescribeGeofenceCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGeofenceCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMap(ctx context.Context, input *locationservice.DescribeMapInput) (*locationservice.DescribeMapOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMapWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePlaceIndex(ctx context.Context, input *locationservice.DescribePlaceIndexInput) (*locationservice.DescribePlaceIndexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePlaceIndexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTracker(ctx context.Context, input *locationservice.DescribeTrackerInput) (*locationservice.DescribeTrackerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTrackerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateTrackerConsumer(ctx context.Context, input *locationservice.DisassociateTrackerConsumerInput) (*locationservice.DisassociateTrackerConsumerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateTrackerConsumerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevicePosition(ctx context.Context, input *locationservice.GetDevicePositionInput) (*locationservice.GetDevicePositionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDevicePositionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevicePositionHistory(ctx context.Context, input *locationservice.GetDevicePositionHistoryInput) (*locationservice.GetDevicePositionHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDevicePositionHistoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGeofence(ctx context.Context, input *locationservice.GetGeofenceInput) (*locationservice.GetGeofenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGeofenceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMapGlyphs(ctx context.Context, input *locationservice.GetMapGlyphsInput) (*locationservice.GetMapGlyphsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMapGlyphsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMapSprites(ctx context.Context, input *locationservice.GetMapSpritesInput) (*locationservice.GetMapSpritesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMapSpritesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMapStyleDescriptor(ctx context.Context, input *locationservice.GetMapStyleDescriptorInput) (*locationservice.GetMapStyleDescriptorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMapStyleDescriptorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMapTile(ctx context.Context, input *locationservice.GetMapTileInput) (*locationservice.GetMapTileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMapTileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGeofenceCollections(ctx context.Context, input *locationservice.ListGeofenceCollectionsInput) (*locationservice.ListGeofenceCollectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGeofenceCollectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGeofences(ctx context.Context, input *locationservice.ListGeofencesInput) (*locationservice.ListGeofencesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGeofencesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMaps(ctx context.Context, input *locationservice.ListMapsInput) (*locationservice.ListMapsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMapsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPlaceIndexes(ctx context.Context, input *locationservice.ListPlaceIndexesInput) (*locationservice.ListPlaceIndexesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPlaceIndexesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTrackerConsumers(ctx context.Context, input *locationservice.ListTrackerConsumersInput) (*locationservice.ListTrackerConsumersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTrackerConsumersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTrackers(ctx context.Context, input *locationservice.ListTrackersInput) (*locationservice.ListTrackersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTrackersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutGeofence(ctx context.Context, input *locationservice.PutGeofenceInput) (*locationservice.PutGeofenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutGeofenceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchPlaceIndexForPosition(ctx context.Context, input *locationservice.SearchPlaceIndexForPositionInput) (*locationservice.SearchPlaceIndexForPositionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchPlaceIndexForPositionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchPlaceIndexForText(ctx context.Context, input *locationservice.SearchPlaceIndexForTextInput) (*locationservice.SearchPlaceIndexForTextOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchPlaceIndexForTextWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
