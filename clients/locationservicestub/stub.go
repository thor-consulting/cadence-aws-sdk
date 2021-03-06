// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package locationservicestub

import (
	"github.com/aws/aws-sdk-go/service/locationservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AssociateTrackerConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateTrackerConsumerFuture) Get(ctx workflow.Context) (*locationservice.AssociateTrackerConsumerOutput, error) {
	var output locationservice.AssociateTrackerConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDeleteGeofenceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDeleteGeofenceFuture) Get(ctx workflow.Context) (*locationservice.BatchDeleteGeofenceOutput, error) {
	var output locationservice.BatchDeleteGeofenceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchEvaluateGeofencesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchEvaluateGeofencesFuture) Get(ctx workflow.Context) (*locationservice.BatchEvaluateGeofencesOutput, error) {
	var output locationservice.BatchEvaluateGeofencesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchGetDevicePositionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchGetDevicePositionFuture) Get(ctx workflow.Context) (*locationservice.BatchGetDevicePositionOutput, error) {
	var output locationservice.BatchGetDevicePositionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchPutGeofenceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchPutGeofenceFuture) Get(ctx workflow.Context) (*locationservice.BatchPutGeofenceOutput, error) {
	var output locationservice.BatchPutGeofenceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchUpdateDevicePositionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchUpdateDevicePositionFuture) Get(ctx workflow.Context) (*locationservice.BatchUpdateDevicePositionOutput, error) {
	var output locationservice.BatchUpdateDevicePositionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateGeofenceCollectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateGeofenceCollectionFuture) Get(ctx workflow.Context) (*locationservice.CreateGeofenceCollectionOutput, error) {
	var output locationservice.CreateGeofenceCollectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateMapFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateMapFuture) Get(ctx workflow.Context) (*locationservice.CreateMapOutput, error) {
	var output locationservice.CreateMapOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreatePlaceIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreatePlaceIndexFuture) Get(ctx workflow.Context) (*locationservice.CreatePlaceIndexOutput, error) {
	var output locationservice.CreatePlaceIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateTrackerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateTrackerFuture) Get(ctx workflow.Context) (*locationservice.CreateTrackerOutput, error) {
	var output locationservice.CreateTrackerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteGeofenceCollectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteGeofenceCollectionFuture) Get(ctx workflow.Context) (*locationservice.DeleteGeofenceCollectionOutput, error) {
	var output locationservice.DeleteGeofenceCollectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteMapFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteMapFuture) Get(ctx workflow.Context) (*locationservice.DeleteMapOutput, error) {
	var output locationservice.DeleteMapOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeletePlaceIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeletePlaceIndexFuture) Get(ctx workflow.Context) (*locationservice.DeletePlaceIndexOutput, error) {
	var output locationservice.DeletePlaceIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteTrackerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteTrackerFuture) Get(ctx workflow.Context) (*locationservice.DeleteTrackerOutput, error) {
	var output locationservice.DeleteTrackerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeGeofenceCollectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeGeofenceCollectionFuture) Get(ctx workflow.Context) (*locationservice.DescribeGeofenceCollectionOutput, error) {
	var output locationservice.DescribeGeofenceCollectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeMapFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeMapFuture) Get(ctx workflow.Context) (*locationservice.DescribeMapOutput, error) {
	var output locationservice.DescribeMapOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribePlaceIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribePlaceIndexFuture) Get(ctx workflow.Context) (*locationservice.DescribePlaceIndexOutput, error) {
	var output locationservice.DescribePlaceIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeTrackerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeTrackerFuture) Get(ctx workflow.Context) (*locationservice.DescribeTrackerOutput, error) {
	var output locationservice.DescribeTrackerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateTrackerConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateTrackerConsumerFuture) Get(ctx workflow.Context) (*locationservice.DisassociateTrackerConsumerOutput, error) {
	var output locationservice.DisassociateTrackerConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDevicePositionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDevicePositionFuture) Get(ctx workflow.Context) (*locationservice.GetDevicePositionOutput, error) {
	var output locationservice.GetDevicePositionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDevicePositionHistoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDevicePositionHistoryFuture) Get(ctx workflow.Context) (*locationservice.GetDevicePositionHistoryOutput, error) {
	var output locationservice.GetDevicePositionHistoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetGeofenceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetGeofenceFuture) Get(ctx workflow.Context) (*locationservice.GetGeofenceOutput, error) {
	var output locationservice.GetGeofenceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMapGlyphsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMapGlyphsFuture) Get(ctx workflow.Context) (*locationservice.GetMapGlyphsOutput, error) {
	var output locationservice.GetMapGlyphsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMapSpritesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMapSpritesFuture) Get(ctx workflow.Context) (*locationservice.GetMapSpritesOutput, error) {
	var output locationservice.GetMapSpritesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMapStyleDescriptorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMapStyleDescriptorFuture) Get(ctx workflow.Context) (*locationservice.GetMapStyleDescriptorOutput, error) {
	var output locationservice.GetMapStyleDescriptorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMapTileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMapTileFuture) Get(ctx workflow.Context) (*locationservice.GetMapTileOutput, error) {
	var output locationservice.GetMapTileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListGeofenceCollectionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListGeofenceCollectionsFuture) Get(ctx workflow.Context) (*locationservice.ListGeofenceCollectionsOutput, error) {
	var output locationservice.ListGeofenceCollectionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListGeofencesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListGeofencesFuture) Get(ctx workflow.Context) (*locationservice.ListGeofencesOutput, error) {
	var output locationservice.ListGeofencesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListMapsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListMapsFuture) Get(ctx workflow.Context) (*locationservice.ListMapsOutput, error) {
	var output locationservice.ListMapsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPlaceIndexesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPlaceIndexesFuture) Get(ctx workflow.Context) (*locationservice.ListPlaceIndexesOutput, error) {
	var output locationservice.ListPlaceIndexesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTrackerConsumersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTrackerConsumersFuture) Get(ctx workflow.Context) (*locationservice.ListTrackerConsumersOutput, error) {
	var output locationservice.ListTrackerConsumersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTrackersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTrackersFuture) Get(ctx workflow.Context) (*locationservice.ListTrackersOutput, error) {
	var output locationservice.ListTrackersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutGeofenceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutGeofenceFuture) Get(ctx workflow.Context) (*locationservice.PutGeofenceOutput, error) {
	var output locationservice.PutGeofenceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SearchPlaceIndexForPositionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SearchPlaceIndexForPositionFuture) Get(ctx workflow.Context) (*locationservice.SearchPlaceIndexForPositionOutput, error) {
	var output locationservice.SearchPlaceIndexForPositionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SearchPlaceIndexForTextFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SearchPlaceIndexForTextFuture) Get(ctx workflow.Context) (*locationservice.SearchPlaceIndexForTextOutput, error) {
	var output locationservice.SearchPlaceIndexForTextOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateTrackerConsumer(ctx workflow.Context, input *locationservice.AssociateTrackerConsumerInput) (*locationservice.AssociateTrackerConsumerOutput, error) {
	var output locationservice.AssociateTrackerConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-AssociateTrackerConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateTrackerConsumerAsync(ctx workflow.Context, input *locationservice.AssociateTrackerConsumerInput) *AssociateTrackerConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-AssociateTrackerConsumer", input)
	return &AssociateTrackerConsumerFuture{Future: future}
}

func (a *stub) BatchDeleteGeofence(ctx workflow.Context, input *locationservice.BatchDeleteGeofenceInput) (*locationservice.BatchDeleteGeofenceOutput, error) {
	var output locationservice.BatchDeleteGeofenceOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchDeleteGeofence", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteGeofenceAsync(ctx workflow.Context, input *locationservice.BatchDeleteGeofenceInput) *BatchDeleteGeofenceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchDeleteGeofence", input)
	return &BatchDeleteGeofenceFuture{Future: future}
}

func (a *stub) BatchEvaluateGeofences(ctx workflow.Context, input *locationservice.BatchEvaluateGeofencesInput) (*locationservice.BatchEvaluateGeofencesOutput, error) {
	var output locationservice.BatchEvaluateGeofencesOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchEvaluateGeofences", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchEvaluateGeofencesAsync(ctx workflow.Context, input *locationservice.BatchEvaluateGeofencesInput) *BatchEvaluateGeofencesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchEvaluateGeofences", input)
	return &BatchEvaluateGeofencesFuture{Future: future}
}

func (a *stub) BatchGetDevicePosition(ctx workflow.Context, input *locationservice.BatchGetDevicePositionInput) (*locationservice.BatchGetDevicePositionOutput, error) {
	var output locationservice.BatchGetDevicePositionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchGetDevicePosition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetDevicePositionAsync(ctx workflow.Context, input *locationservice.BatchGetDevicePositionInput) *BatchGetDevicePositionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchGetDevicePosition", input)
	return &BatchGetDevicePositionFuture{Future: future}
}

func (a *stub) BatchPutGeofence(ctx workflow.Context, input *locationservice.BatchPutGeofenceInput) (*locationservice.BatchPutGeofenceOutput, error) {
	var output locationservice.BatchPutGeofenceOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchPutGeofence", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchPutGeofenceAsync(ctx workflow.Context, input *locationservice.BatchPutGeofenceInput) *BatchPutGeofenceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchPutGeofence", input)
	return &BatchPutGeofenceFuture{Future: future}
}

func (a *stub) BatchUpdateDevicePosition(ctx workflow.Context, input *locationservice.BatchUpdateDevicePositionInput) (*locationservice.BatchUpdateDevicePositionOutput, error) {
	var output locationservice.BatchUpdateDevicePositionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchUpdateDevicePosition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchUpdateDevicePositionAsync(ctx workflow.Context, input *locationservice.BatchUpdateDevicePositionInput) *BatchUpdateDevicePositionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-BatchUpdateDevicePosition", input)
	return &BatchUpdateDevicePositionFuture{Future: future}
}

func (a *stub) CreateGeofenceCollection(ctx workflow.Context, input *locationservice.CreateGeofenceCollectionInput) (*locationservice.CreateGeofenceCollectionOutput, error) {
	var output locationservice.CreateGeofenceCollectionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-CreateGeofenceCollection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGeofenceCollectionAsync(ctx workflow.Context, input *locationservice.CreateGeofenceCollectionInput) *CreateGeofenceCollectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-CreateGeofenceCollection", input)
	return &CreateGeofenceCollectionFuture{Future: future}
}

func (a *stub) CreateMap(ctx workflow.Context, input *locationservice.CreateMapInput) (*locationservice.CreateMapOutput, error) {
	var output locationservice.CreateMapOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-CreateMap", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMapAsync(ctx workflow.Context, input *locationservice.CreateMapInput) *CreateMapFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-CreateMap", input)
	return &CreateMapFuture{Future: future}
}

func (a *stub) CreatePlaceIndex(ctx workflow.Context, input *locationservice.CreatePlaceIndexInput) (*locationservice.CreatePlaceIndexOutput, error) {
	var output locationservice.CreatePlaceIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-CreatePlaceIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePlaceIndexAsync(ctx workflow.Context, input *locationservice.CreatePlaceIndexInput) *CreatePlaceIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-CreatePlaceIndex", input)
	return &CreatePlaceIndexFuture{Future: future}
}

func (a *stub) CreateTracker(ctx workflow.Context, input *locationservice.CreateTrackerInput) (*locationservice.CreateTrackerOutput, error) {
	var output locationservice.CreateTrackerOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-CreateTracker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTrackerAsync(ctx workflow.Context, input *locationservice.CreateTrackerInput) *CreateTrackerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-CreateTracker", input)
	return &CreateTrackerFuture{Future: future}
}

func (a *stub) DeleteGeofenceCollection(ctx workflow.Context, input *locationservice.DeleteGeofenceCollectionInput) (*locationservice.DeleteGeofenceCollectionOutput, error) {
	var output locationservice.DeleteGeofenceCollectionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DeleteGeofenceCollection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteGeofenceCollectionAsync(ctx workflow.Context, input *locationservice.DeleteGeofenceCollectionInput) *DeleteGeofenceCollectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DeleteGeofenceCollection", input)
	return &DeleteGeofenceCollectionFuture{Future: future}
}

func (a *stub) DeleteMap(ctx workflow.Context, input *locationservice.DeleteMapInput) (*locationservice.DeleteMapOutput, error) {
	var output locationservice.DeleteMapOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DeleteMap", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMapAsync(ctx workflow.Context, input *locationservice.DeleteMapInput) *DeleteMapFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DeleteMap", input)
	return &DeleteMapFuture{Future: future}
}

func (a *stub) DeletePlaceIndex(ctx workflow.Context, input *locationservice.DeletePlaceIndexInput) (*locationservice.DeletePlaceIndexOutput, error) {
	var output locationservice.DeletePlaceIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DeletePlaceIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePlaceIndexAsync(ctx workflow.Context, input *locationservice.DeletePlaceIndexInput) *DeletePlaceIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DeletePlaceIndex", input)
	return &DeletePlaceIndexFuture{Future: future}
}

func (a *stub) DeleteTracker(ctx workflow.Context, input *locationservice.DeleteTrackerInput) (*locationservice.DeleteTrackerOutput, error) {
	var output locationservice.DeleteTrackerOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DeleteTracker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTrackerAsync(ctx workflow.Context, input *locationservice.DeleteTrackerInput) *DeleteTrackerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DeleteTracker", input)
	return &DeleteTrackerFuture{Future: future}
}

func (a *stub) DescribeGeofenceCollection(ctx workflow.Context, input *locationservice.DescribeGeofenceCollectionInput) (*locationservice.DescribeGeofenceCollectionOutput, error) {
	var output locationservice.DescribeGeofenceCollectionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribeGeofenceCollection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeGeofenceCollectionAsync(ctx workflow.Context, input *locationservice.DescribeGeofenceCollectionInput) *DescribeGeofenceCollectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribeGeofenceCollection", input)
	return &DescribeGeofenceCollectionFuture{Future: future}
}

func (a *stub) DescribeMap(ctx workflow.Context, input *locationservice.DescribeMapInput) (*locationservice.DescribeMapOutput, error) {
	var output locationservice.DescribeMapOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribeMap", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeMapAsync(ctx workflow.Context, input *locationservice.DescribeMapInput) *DescribeMapFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribeMap", input)
	return &DescribeMapFuture{Future: future}
}

func (a *stub) DescribePlaceIndex(ctx workflow.Context, input *locationservice.DescribePlaceIndexInput) (*locationservice.DescribePlaceIndexOutput, error) {
	var output locationservice.DescribePlaceIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribePlaceIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePlaceIndexAsync(ctx workflow.Context, input *locationservice.DescribePlaceIndexInput) *DescribePlaceIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribePlaceIndex", input)
	return &DescribePlaceIndexFuture{Future: future}
}

func (a *stub) DescribeTracker(ctx workflow.Context, input *locationservice.DescribeTrackerInput) (*locationservice.DescribeTrackerOutput, error) {
	var output locationservice.DescribeTrackerOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribeTracker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTrackerAsync(ctx workflow.Context, input *locationservice.DescribeTrackerInput) *DescribeTrackerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DescribeTracker", input)
	return &DescribeTrackerFuture{Future: future}
}

func (a *stub) DisassociateTrackerConsumer(ctx workflow.Context, input *locationservice.DisassociateTrackerConsumerInput) (*locationservice.DisassociateTrackerConsumerOutput, error) {
	var output locationservice.DisassociateTrackerConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-DisassociateTrackerConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateTrackerConsumerAsync(ctx workflow.Context, input *locationservice.DisassociateTrackerConsumerInput) *DisassociateTrackerConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-DisassociateTrackerConsumer", input)
	return &DisassociateTrackerConsumerFuture{Future: future}
}

func (a *stub) GetDevicePosition(ctx workflow.Context, input *locationservice.GetDevicePositionInput) (*locationservice.GetDevicePositionOutput, error) {
	var output locationservice.GetDevicePositionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetDevicePosition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDevicePositionAsync(ctx workflow.Context, input *locationservice.GetDevicePositionInput) *GetDevicePositionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetDevicePosition", input)
	return &GetDevicePositionFuture{Future: future}
}

func (a *stub) GetDevicePositionHistory(ctx workflow.Context, input *locationservice.GetDevicePositionHistoryInput) (*locationservice.GetDevicePositionHistoryOutput, error) {
	var output locationservice.GetDevicePositionHistoryOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetDevicePositionHistory", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDevicePositionHistoryAsync(ctx workflow.Context, input *locationservice.GetDevicePositionHistoryInput) *GetDevicePositionHistoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetDevicePositionHistory", input)
	return &GetDevicePositionHistoryFuture{Future: future}
}

func (a *stub) GetGeofence(ctx workflow.Context, input *locationservice.GetGeofenceInput) (*locationservice.GetGeofenceOutput, error) {
	var output locationservice.GetGeofenceOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetGeofence", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetGeofenceAsync(ctx workflow.Context, input *locationservice.GetGeofenceInput) *GetGeofenceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetGeofence", input)
	return &GetGeofenceFuture{Future: future}
}

func (a *stub) GetMapGlyphs(ctx workflow.Context, input *locationservice.GetMapGlyphsInput) (*locationservice.GetMapGlyphsOutput, error) {
	var output locationservice.GetMapGlyphsOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapGlyphs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMapGlyphsAsync(ctx workflow.Context, input *locationservice.GetMapGlyphsInput) *GetMapGlyphsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapGlyphs", input)
	return &GetMapGlyphsFuture{Future: future}
}

func (a *stub) GetMapSprites(ctx workflow.Context, input *locationservice.GetMapSpritesInput) (*locationservice.GetMapSpritesOutput, error) {
	var output locationservice.GetMapSpritesOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapSprites", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMapSpritesAsync(ctx workflow.Context, input *locationservice.GetMapSpritesInput) *GetMapSpritesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapSprites", input)
	return &GetMapSpritesFuture{Future: future}
}

func (a *stub) GetMapStyleDescriptor(ctx workflow.Context, input *locationservice.GetMapStyleDescriptorInput) (*locationservice.GetMapStyleDescriptorOutput, error) {
	var output locationservice.GetMapStyleDescriptorOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapStyleDescriptor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMapStyleDescriptorAsync(ctx workflow.Context, input *locationservice.GetMapStyleDescriptorInput) *GetMapStyleDescriptorFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapStyleDescriptor", input)
	return &GetMapStyleDescriptorFuture{Future: future}
}

func (a *stub) GetMapTile(ctx workflow.Context, input *locationservice.GetMapTileInput) (*locationservice.GetMapTileOutput, error) {
	var output locationservice.GetMapTileOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapTile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMapTileAsync(ctx workflow.Context, input *locationservice.GetMapTileInput) *GetMapTileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-GetMapTile", input)
	return &GetMapTileFuture{Future: future}
}

func (a *stub) ListGeofenceCollections(ctx workflow.Context, input *locationservice.ListGeofenceCollectionsInput) (*locationservice.ListGeofenceCollectionsOutput, error) {
	var output locationservice.ListGeofenceCollectionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-ListGeofenceCollections", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGeofenceCollectionsAsync(ctx workflow.Context, input *locationservice.ListGeofenceCollectionsInput) *ListGeofenceCollectionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-ListGeofenceCollections", input)
	return &ListGeofenceCollectionsFuture{Future: future}
}

func (a *stub) ListGeofences(ctx workflow.Context, input *locationservice.ListGeofencesInput) (*locationservice.ListGeofencesOutput, error) {
	var output locationservice.ListGeofencesOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-ListGeofences", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGeofencesAsync(ctx workflow.Context, input *locationservice.ListGeofencesInput) *ListGeofencesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-ListGeofences", input)
	return &ListGeofencesFuture{Future: future}
}

func (a *stub) ListMaps(ctx workflow.Context, input *locationservice.ListMapsInput) (*locationservice.ListMapsOutput, error) {
	var output locationservice.ListMapsOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-ListMaps", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMapsAsync(ctx workflow.Context, input *locationservice.ListMapsInput) *ListMapsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-ListMaps", input)
	return &ListMapsFuture{Future: future}
}

func (a *stub) ListPlaceIndexes(ctx workflow.Context, input *locationservice.ListPlaceIndexesInput) (*locationservice.ListPlaceIndexesOutput, error) {
	var output locationservice.ListPlaceIndexesOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-ListPlaceIndexes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPlaceIndexesAsync(ctx workflow.Context, input *locationservice.ListPlaceIndexesInput) *ListPlaceIndexesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-ListPlaceIndexes", input)
	return &ListPlaceIndexesFuture{Future: future}
}

func (a *stub) ListTrackerConsumers(ctx workflow.Context, input *locationservice.ListTrackerConsumersInput) (*locationservice.ListTrackerConsumersOutput, error) {
	var output locationservice.ListTrackerConsumersOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-ListTrackerConsumers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTrackerConsumersAsync(ctx workflow.Context, input *locationservice.ListTrackerConsumersInput) *ListTrackerConsumersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-ListTrackerConsumers", input)
	return &ListTrackerConsumersFuture{Future: future}
}

func (a *stub) ListTrackers(ctx workflow.Context, input *locationservice.ListTrackersInput) (*locationservice.ListTrackersOutput, error) {
	var output locationservice.ListTrackersOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-ListTrackers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTrackersAsync(ctx workflow.Context, input *locationservice.ListTrackersInput) *ListTrackersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-ListTrackers", input)
	return &ListTrackersFuture{Future: future}
}

func (a *stub) PutGeofence(ctx workflow.Context, input *locationservice.PutGeofenceInput) (*locationservice.PutGeofenceOutput, error) {
	var output locationservice.PutGeofenceOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-PutGeofence", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutGeofenceAsync(ctx workflow.Context, input *locationservice.PutGeofenceInput) *PutGeofenceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-PutGeofence", input)
	return &PutGeofenceFuture{Future: future}
}

func (a *stub) SearchPlaceIndexForPosition(ctx workflow.Context, input *locationservice.SearchPlaceIndexForPositionInput) (*locationservice.SearchPlaceIndexForPositionOutput, error) {
	var output locationservice.SearchPlaceIndexForPositionOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-SearchPlaceIndexForPosition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SearchPlaceIndexForPositionAsync(ctx workflow.Context, input *locationservice.SearchPlaceIndexForPositionInput) *SearchPlaceIndexForPositionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-SearchPlaceIndexForPosition", input)
	return &SearchPlaceIndexForPositionFuture{Future: future}
}

func (a *stub) SearchPlaceIndexForText(ctx workflow.Context, input *locationservice.SearchPlaceIndexForTextInput) (*locationservice.SearchPlaceIndexForTextOutput, error) {
	var output locationservice.SearchPlaceIndexForTextOutput
	err := workflow.ExecuteActivity(ctx, "aws-locationservice-SearchPlaceIndexForText", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SearchPlaceIndexForTextAsync(ctx workflow.Context, input *locationservice.SearchPlaceIndexForTextInput) *SearchPlaceIndexForTextFuture {
	future := workflow.ExecuteActivity(ctx, "aws-locationservice-SearchPlaceIndexForText", input)
	return &SearchPlaceIndexForTextFuture{Future: future}
}
