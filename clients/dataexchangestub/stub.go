// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package dataexchangestub

import (
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CancelJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CancelJobFuture) Get(ctx workflow.Context) (*dataexchange.CancelJobOutput, error) {
	var output dataexchange.CancelJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDataSetFuture) Get(ctx workflow.Context) (*dataexchange.CreateDataSetOutput, error) {
	var output dataexchange.CreateDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateJobFuture) Get(ctx workflow.Context) (*dataexchange.CreateJobOutput, error) {
	var output dataexchange.CreateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateRevisionFuture) Get(ctx workflow.Context) (*dataexchange.CreateRevisionOutput, error) {
	var output dataexchange.CreateRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteAssetFuture) Get(ctx workflow.Context) (*dataexchange.DeleteAssetOutput, error) {
	var output dataexchange.DeleteAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDataSetFuture) Get(ctx workflow.Context) (*dataexchange.DeleteDataSetOutput, error) {
	var output dataexchange.DeleteDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteRevisionFuture) Get(ctx workflow.Context) (*dataexchange.DeleteRevisionOutput, error) {
	var output dataexchange.DeleteRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAssetFuture) Get(ctx workflow.Context) (*dataexchange.GetAssetOutput, error) {
	var output dataexchange.GetAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDataSetFuture) Get(ctx workflow.Context) (*dataexchange.GetDataSetOutput, error) {
	var output dataexchange.GetDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetJobFuture) Get(ctx workflow.Context) (*dataexchange.GetJobOutput, error) {
	var output dataexchange.GetJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetRevisionFuture) Get(ctx workflow.Context) (*dataexchange.GetRevisionOutput, error) {
	var output dataexchange.GetRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDataSetRevisionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDataSetRevisionsFuture) Get(ctx workflow.Context) (*dataexchange.ListDataSetRevisionsOutput, error) {
	var output dataexchange.ListDataSetRevisionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDataSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDataSetsFuture) Get(ctx workflow.Context) (*dataexchange.ListDataSetsOutput, error) {
	var output dataexchange.ListDataSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListJobsFuture) Get(ctx workflow.Context) (*dataexchange.ListJobsOutput, error) {
	var output dataexchange.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListRevisionAssetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListRevisionAssetsFuture) Get(ctx workflow.Context) (*dataexchange.ListRevisionAssetsOutput, error) {
	var output dataexchange.ListRevisionAssetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*dataexchange.ListTagsForResourceOutput, error) {
	var output dataexchange.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartJobFuture) Get(ctx workflow.Context) (*dataexchange.StartJobOutput, error) {
	var output dataexchange.StartJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*dataexchange.TagResourceOutput, error) {
	var output dataexchange.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*dataexchange.UntagResourceOutput, error) {
	var output dataexchange.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateAssetFuture) Get(ctx workflow.Context) (*dataexchange.UpdateAssetOutput, error) {
	var output dataexchange.UpdateAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDataSetFuture) Get(ctx workflow.Context) (*dataexchange.UpdateDataSetOutput, error) {
	var output dataexchange.UpdateDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateRevisionFuture) Get(ctx workflow.Context) (*dataexchange.UpdateRevisionOutput, error) {
	var output dataexchange.UpdateRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJob(ctx workflow.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
	var output dataexchange.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-CancelJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJobAsync(ctx workflow.Context, input *dataexchange.CancelJobInput) *CancelJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-CancelJob", input)
	return &CancelJobFuture{Future: future}
}

func (a *stub) CreateDataSet(ctx workflow.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
	var output dataexchange.CreateDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-CreateDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSetAsync(ctx workflow.Context, input *dataexchange.CreateDataSetInput) *CreateDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-CreateDataSet", input)
	return &CreateDataSetFuture{Future: future}
}

func (a *stub) CreateJob(ctx workflow.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
	var output dataexchange.CreateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-CreateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobAsync(ctx workflow.Context, input *dataexchange.CreateJobInput) *CreateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-CreateJob", input)
	return &CreateJobFuture{Future: future}
}

func (a *stub) CreateRevision(ctx workflow.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
	var output dataexchange.CreateRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-CreateRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRevisionAsync(ctx workflow.Context, input *dataexchange.CreateRevisionInput) *CreateRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-CreateRevision", input)
	return &CreateRevisionFuture{Future: future}
}

func (a *stub) DeleteAsset(ctx workflow.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
	var output dataexchange.DeleteAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-DeleteAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAssetAsync(ctx workflow.Context, input *dataexchange.DeleteAssetInput) *DeleteAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-DeleteAsset", input)
	return &DeleteAssetFuture{Future: future}
}

func (a *stub) DeleteDataSet(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
	var output dataexchange.DeleteDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-DeleteDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataSetAsync(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) *DeleteDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-DeleteDataSet", input)
	return &DeleteDataSetFuture{Future: future}
}

func (a *stub) DeleteRevision(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
	var output dataexchange.DeleteRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-DeleteRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRevisionAsync(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) *DeleteRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-DeleteRevision", input)
	return &DeleteRevisionFuture{Future: future}
}

func (a *stub) GetAsset(ctx workflow.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
	var output dataexchange.GetAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAssetAsync(ctx workflow.Context, input *dataexchange.GetAssetInput) *GetAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetAsset", input)
	return &GetAssetFuture{Future: future}
}

func (a *stub) GetDataSet(ctx workflow.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
	var output dataexchange.GetDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataSetAsync(ctx workflow.Context, input *dataexchange.GetDataSetInput) *GetDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetDataSet", input)
	return &GetDataSetFuture{Future: future}
}

func (a *stub) GetJob(ctx workflow.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
	var output dataexchange.GetJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobAsync(ctx workflow.Context, input *dataexchange.GetJobInput) *GetJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetJob", input)
	return &GetJobFuture{Future: future}
}

func (a *stub) GetRevision(ctx workflow.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
	var output dataexchange.GetRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRevisionAsync(ctx workflow.Context, input *dataexchange.GetRevisionInput) *GetRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-GetRevision", input)
	return &GetRevisionFuture{Future: future}
}

func (a *stub) ListDataSetRevisions(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
	var output dataexchange.ListDataSetRevisionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListDataSetRevisions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSetRevisionsAsync(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) *ListDataSetRevisionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListDataSetRevisions", input)
	return &ListDataSetRevisionsFuture{Future: future}
}

func (a *stub) ListDataSets(ctx workflow.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
	var output dataexchange.ListDataSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListDataSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSetsAsync(ctx workflow.Context, input *dataexchange.ListDataSetsInput) *ListDataSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListDataSets", input)
	return &ListDataSetsFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
	var output dataexchange.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *dataexchange.ListJobsInput) *ListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListJobs", input)
	return &ListJobsFuture{Future: future}
}

func (a *stub) ListRevisionAssets(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
	var output dataexchange.ListRevisionAssetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListRevisionAssets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRevisionAssetsAsync(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) *ListRevisionAssetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListRevisionAssets", input)
	return &ListRevisionAssetsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
	var output dataexchange.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) StartJob(ctx workflow.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
	var output dataexchange.StartJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-StartJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartJobAsync(ctx workflow.Context, input *dataexchange.StartJobInput) *StartJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-StartJob", input)
	return &StartJobFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
	var output dataexchange.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *dataexchange.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
	var output dataexchange.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *dataexchange.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateAsset(ctx workflow.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
	var output dataexchange.UpdateAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-UpdateAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAssetAsync(ctx workflow.Context, input *dataexchange.UpdateAssetInput) *UpdateAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-UpdateAsset", input)
	return &UpdateAssetFuture{Future: future}
}

func (a *stub) UpdateDataSet(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
	var output dataexchange.UpdateDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-UpdateDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDataSetAsync(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) *UpdateDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-UpdateDataSet", input)
	return &UpdateDataSetFuture{Future: future}
}

func (a *stub) UpdateRevision(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
	var output dataexchange.UpdateRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws-dataexchange-UpdateRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRevisionAsync(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) *UpdateRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-dataexchange-UpdateRevision", input)
	return &UpdateRevisionFuture{Future: future}
}
