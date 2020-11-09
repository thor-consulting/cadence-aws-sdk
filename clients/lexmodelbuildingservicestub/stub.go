// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package lexmodelbuildingservicestub

import (
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateBotVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBotVersionFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
	var output lexmodelbuildingservice.CreateBotVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateIntentVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateIntentVersionFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
	var output lexmodelbuildingservice.CreateIntentVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSlotTypeVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSlotTypeVersionFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
	var output lexmodelbuildingservice.CreateSlotTypeVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotOutput, error) {
	var output lexmodelbuildingservice.DeleteBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotAliasFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
	var output lexmodelbuildingservice.DeleteBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotChannelAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotChannelAssociationFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
	var output lexmodelbuildingservice.DeleteBotChannelAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotVersionFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
	var output lexmodelbuildingservice.DeleteBotVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteIntentFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
	var output lexmodelbuildingservice.DeleteIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteIntentVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteIntentVersionFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
	var output lexmodelbuildingservice.DeleteIntentVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
	var output lexmodelbuildingservice.DeleteSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSlotTypeVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSlotTypeVersionFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
	var output lexmodelbuildingservice.DeleteSlotTypeVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteUtterancesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteUtterancesFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
	var output lexmodelbuildingservice.DeleteUtterancesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotOutput, error) {
	var output lexmodelbuildingservice.GetBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotAliasFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
	var output lexmodelbuildingservice.GetBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotAliasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotAliasesFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
	var output lexmodelbuildingservice.GetBotAliasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotChannelAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotChannelAssociationFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
	var output lexmodelbuildingservice.GetBotChannelAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotChannelAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotChannelAssociationsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
	var output lexmodelbuildingservice.GetBotChannelAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotVersionsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
	var output lexmodelbuildingservice.GetBotVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBotsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBotsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotsOutput, error) {
	var output lexmodelbuildingservice.GetBotsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBuiltinIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBuiltinIntentFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
	var output lexmodelbuildingservice.GetBuiltinIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBuiltinIntentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBuiltinIntentsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
	var output lexmodelbuildingservice.GetBuiltinIntentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBuiltinSlotTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBuiltinSlotTypesFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
	var output lexmodelbuildingservice.GetBuiltinSlotTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetExportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetExportFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetExportOutput, error) {
	var output lexmodelbuildingservice.GetExportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetImportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetImportFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetImportOutput, error) {
	var output lexmodelbuildingservice.GetImportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetIntentFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetIntentOutput, error) {
	var output lexmodelbuildingservice.GetIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetIntentVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetIntentVersionsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
	var output lexmodelbuildingservice.GetIntentVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetIntentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetIntentsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetIntentsOutput, error) {
	var output lexmodelbuildingservice.GetIntentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
	var output lexmodelbuildingservice.GetSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSlotTypeVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSlotTypeVersionsFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
	var output lexmodelbuildingservice.GetSlotTypeVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSlotTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSlotTypesFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
	var output lexmodelbuildingservice.GetSlotTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetUtterancesViewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetUtterancesViewFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
	var output lexmodelbuildingservice.GetUtterancesViewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
	var output lexmodelbuildingservice.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutBotFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutBotOutput, error) {
	var output lexmodelbuildingservice.PutBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutBotAliasFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
	var output lexmodelbuildingservice.PutBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutIntentFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutIntentOutput, error) {
	var output lexmodelbuildingservice.PutIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
	var output lexmodelbuildingservice.PutSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartImportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartImportFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.StartImportOutput, error) {
	var output lexmodelbuildingservice.StartImportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.TagResourceOutput, error) {
	var output lexmodelbuildingservice.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*lexmodelbuildingservice.UntagResourceOutput, error) {
	var output lexmodelbuildingservice.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
	var output lexmodelbuildingservice.CreateBotVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-CreateBotVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) *CreateBotVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-CreateBotVersion", input)
	return &CreateBotVersionFuture{Future: future}
}

func (a *stub) CreateIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
	var output lexmodelbuildingservice.CreateIntentVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-CreateIntentVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) *CreateIntentVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-CreateIntentVersion", input)
	return &CreateIntentVersionFuture{Future: future}
}

func (a *stub) CreateSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
	var output lexmodelbuildingservice.CreateSlotTypeVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-CreateSlotTypeVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) *CreateSlotTypeVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-CreateSlotTypeVersion", input)
	return &CreateSlotTypeVersionFuture{Future: future}
}

func (a *stub) DeleteBot(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error) {
	var output lexmodelbuildingservice.DeleteBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) *DeleteBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBot", input)
	return &DeleteBotFuture{Future: future}
}

func (a *stub) DeleteBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
	var output lexmodelbuildingservice.DeleteBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) *DeleteBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBotAlias", input)
	return &DeleteBotAliasFuture{Future: future}
}

func (a *stub) DeleteBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
	var output lexmodelbuildingservice.DeleteBotChannelAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBotChannelAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) *DeleteBotChannelAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBotChannelAssociation", input)
	return &DeleteBotChannelAssociationFuture{Future: future}
}

func (a *stub) DeleteBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
	var output lexmodelbuildingservice.DeleteBotVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBotVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) *DeleteBotVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteBotVersion", input)
	return &DeleteBotVersionFuture{Future: future}
}

func (a *stub) DeleteIntent(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
	var output lexmodelbuildingservice.DeleteIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) *DeleteIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteIntent", input)
	return &DeleteIntentFuture{Future: future}
}

func (a *stub) DeleteIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
	var output lexmodelbuildingservice.DeleteIntentVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteIntentVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) *DeleteIntentVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteIntentVersion", input)
	return &DeleteIntentVersionFuture{Future: future}
}

func (a *stub) DeleteSlotType(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
	var output lexmodelbuildingservice.DeleteSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) *DeleteSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteSlotType", input)
	return &DeleteSlotTypeFuture{Future: future}
}

func (a *stub) DeleteSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
	var output lexmodelbuildingservice.DeleteSlotTypeVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteSlotTypeVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) *DeleteSlotTypeVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteSlotTypeVersion", input)
	return &DeleteSlotTypeVersionFuture{Future: future}
}

func (a *stub) DeleteUtterances(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
	var output lexmodelbuildingservice.DeleteUtterancesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteUtterances", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteUtterancesAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) *DeleteUtterancesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-DeleteUtterances", input)
	return &DeleteUtterancesFuture{Future: future}
}

func (a *stub) GetBot(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error) {
	var output lexmodelbuildingservice.GetBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) *GetBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBot", input)
	return &GetBotFuture{Future: future}
}

func (a *stub) GetBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
	var output lexmodelbuildingservice.GetBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) *GetBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotAlias", input)
	return &GetBotAliasFuture{Future: future}
}

func (a *stub) GetBotAliases(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
	var output lexmodelbuildingservice.GetBotAliasesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotAliases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotAliasesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) *GetBotAliasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotAliases", input)
	return &GetBotAliasesFuture{Future: future}
}

func (a *stub) GetBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
	var output lexmodelbuildingservice.GetBotChannelAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotChannelAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) *GetBotChannelAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotChannelAssociation", input)
	return &GetBotChannelAssociationFuture{Future: future}
}

func (a *stub) GetBotChannelAssociations(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
	var output lexmodelbuildingservice.GetBotChannelAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotChannelAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotChannelAssociationsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) *GetBotChannelAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotChannelAssociations", input)
	return &GetBotChannelAssociationsFuture{Future: future}
}

func (a *stub) GetBotVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
	var output lexmodelbuildingservice.GetBotVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) *GetBotVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBotVersions", input)
	return &GetBotVersionsFuture{Future: future}
}

func (a *stub) GetBots(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error) {
	var output lexmodelbuildingservice.GetBotsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBots", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBotsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) *GetBotsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBots", input)
	return &GetBotsFuture{Future: future}
}

func (a *stub) GetBuiltinIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
	var output lexmodelbuildingservice.GetBuiltinIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBuiltinIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBuiltinIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) *GetBuiltinIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBuiltinIntent", input)
	return &GetBuiltinIntentFuture{Future: future}
}

func (a *stub) GetBuiltinIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
	var output lexmodelbuildingservice.GetBuiltinIntentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBuiltinIntents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBuiltinIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) *GetBuiltinIntentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBuiltinIntents", input)
	return &GetBuiltinIntentsFuture{Future: future}
}

func (a *stub) GetBuiltinSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
	var output lexmodelbuildingservice.GetBuiltinSlotTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBuiltinSlotTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBuiltinSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) *GetBuiltinSlotTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetBuiltinSlotTypes", input)
	return &GetBuiltinSlotTypesFuture{Future: future}
}

func (a *stub) GetExport(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error) {
	var output lexmodelbuildingservice.GetExportOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetExport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetExportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) *GetExportFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetExport", input)
	return &GetExportFuture{Future: future}
}

func (a *stub) GetImport(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error) {
	var output lexmodelbuildingservice.GetImportOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetImport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) *GetImportFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetImport", input)
	return &GetImportFuture{Future: future}
}

func (a *stub) GetIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error) {
	var output lexmodelbuildingservice.GetIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) *GetIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetIntent", input)
	return &GetIntentFuture{Future: future}
}

func (a *stub) GetIntentVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
	var output lexmodelbuildingservice.GetIntentVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetIntentVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIntentVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) *GetIntentVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetIntentVersions", input)
	return &GetIntentVersionsFuture{Future: future}
}

func (a *stub) GetIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error) {
	var output lexmodelbuildingservice.GetIntentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetIntents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) *GetIntentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetIntents", input)
	return &GetIntentsFuture{Future: future}
}

func (a *stub) GetSlotType(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
	var output lexmodelbuildingservice.GetSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) *GetSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetSlotType", input)
	return &GetSlotTypeFuture{Future: future}
}

func (a *stub) GetSlotTypeVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
	var output lexmodelbuildingservice.GetSlotTypeVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetSlotTypeVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSlotTypeVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) *GetSlotTypeVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetSlotTypeVersions", input)
	return &GetSlotTypeVersionsFuture{Future: future}
}

func (a *stub) GetSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
	var output lexmodelbuildingservice.GetSlotTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetSlotTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) *GetSlotTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetSlotTypes", input)
	return &GetSlotTypesFuture{Future: future}
}

func (a *stub) GetUtterancesView(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
	var output lexmodelbuildingservice.GetUtterancesViewOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetUtterancesView", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetUtterancesViewAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) *GetUtterancesViewFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-GetUtterancesView", input)
	return &GetUtterancesViewFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
	var output lexmodelbuildingservice.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PutBot(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error) {
	var output lexmodelbuildingservice.PutBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) *PutBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutBot", input)
	return &PutBotFuture{Future: future}
}

func (a *stub) PutBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
	var output lexmodelbuildingservice.PutBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) *PutBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutBotAlias", input)
	return &PutBotAliasFuture{Future: future}
}

func (a *stub) PutIntent(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error) {
	var output lexmodelbuildingservice.PutIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) *PutIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutIntent", input)
	return &PutIntentFuture{Future: future}
}

func (a *stub) PutSlotType(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
	var output lexmodelbuildingservice.PutSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) *PutSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-PutSlotType", input)
	return &PutSlotTypeFuture{Future: future}
}

func (a *stub) StartImport(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error) {
	var output lexmodelbuildingservice.StartImportOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-StartImport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) *StartImportFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-StartImport", input)
	return &StartImportFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error) {
	var output lexmodelbuildingservice.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error) {
	var output lexmodelbuildingservice.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelbuildingservice-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
