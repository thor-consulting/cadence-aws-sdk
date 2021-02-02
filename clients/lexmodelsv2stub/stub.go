// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package lexmodelsv2stub

import (
	"github.com/aws/aws-sdk-go/service/lexmodelsv2"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type BuildBotLocaleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BuildBotLocaleFuture) Get(ctx workflow.Context) (*lexmodelsv2.BuildBotLocaleOutput, error) {
	var output lexmodelsv2.BuildBotLocaleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBotFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateBotOutput, error) {
	var output lexmodelsv2.CreateBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBotAliasFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateBotAliasOutput, error) {
	var output lexmodelsv2.CreateBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBotLocaleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBotLocaleFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateBotLocaleOutput, error) {
	var output lexmodelsv2.CreateBotLocaleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBotVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBotVersionFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateBotVersionOutput, error) {
	var output lexmodelsv2.CreateBotVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateIntentFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateIntentOutput, error) {
	var output lexmodelsv2.CreateIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSlotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSlotFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateSlotOutput, error) {
	var output lexmodelsv2.CreateSlotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelsv2.CreateSlotTypeOutput, error) {
	var output lexmodelsv2.CreateSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteBotOutput, error) {
	var output lexmodelsv2.DeleteBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotAliasFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteBotAliasOutput, error) {
	var output lexmodelsv2.DeleteBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotLocaleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotLocaleFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteBotLocaleOutput, error) {
	var output lexmodelsv2.DeleteBotLocaleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBotVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBotVersionFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteBotVersionOutput, error) {
	var output lexmodelsv2.DeleteBotVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteIntentFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteIntentOutput, error) {
	var output lexmodelsv2.DeleteIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSlotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSlotFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteSlotOutput, error) {
	var output lexmodelsv2.DeleteSlotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelsv2.DeleteSlotTypeOutput, error) {
	var output lexmodelsv2.DeleteSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBotFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeBotOutput, error) {
	var output lexmodelsv2.DescribeBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBotAliasFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeBotAliasOutput, error) {
	var output lexmodelsv2.DescribeBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBotLocaleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBotLocaleFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeBotLocaleOutput, error) {
	var output lexmodelsv2.DescribeBotLocaleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBotVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBotVersionFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeBotVersionOutput, error) {
	var output lexmodelsv2.DescribeBotVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeIntentFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeIntentOutput, error) {
	var output lexmodelsv2.DescribeIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeSlotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeSlotFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeSlotOutput, error) {
	var output lexmodelsv2.DescribeSlotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelsv2.DescribeSlotTypeOutput, error) {
	var output lexmodelsv2.DescribeSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBotAliasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBotAliasesFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListBotAliasesOutput, error) {
	var output lexmodelsv2.ListBotAliasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBotLocalesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBotLocalesFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListBotLocalesOutput, error) {
	var output lexmodelsv2.ListBotLocalesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBotVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBotVersionsFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListBotVersionsOutput, error) {
	var output lexmodelsv2.ListBotVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBotsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBotsFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListBotsOutput, error) {
	var output lexmodelsv2.ListBotsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBuiltInIntentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBuiltInIntentsFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListBuiltInIntentsOutput, error) {
	var output lexmodelsv2.ListBuiltInIntentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBuiltInSlotTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBuiltInSlotTypesFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListBuiltInSlotTypesOutput, error) {
	var output lexmodelsv2.ListBuiltInSlotTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListIntentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListIntentsFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListIntentsOutput, error) {
	var output lexmodelsv2.ListIntentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSlotTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSlotTypesFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListSlotTypesOutput, error) {
	var output lexmodelsv2.ListSlotTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSlotsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSlotsFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListSlotsOutput, error) {
	var output lexmodelsv2.ListSlotsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*lexmodelsv2.ListTagsForResourceOutput, error) {
	var output lexmodelsv2.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*lexmodelsv2.TagResourceOutput, error) {
	var output lexmodelsv2.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*lexmodelsv2.UntagResourceOutput, error) {
	var output lexmodelsv2.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBotFuture) Get(ctx workflow.Context) (*lexmodelsv2.UpdateBotOutput, error) {
	var output lexmodelsv2.UpdateBotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBotAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBotAliasFuture) Get(ctx workflow.Context) (*lexmodelsv2.UpdateBotAliasOutput, error) {
	var output lexmodelsv2.UpdateBotAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBotLocaleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBotLocaleFuture) Get(ctx workflow.Context) (*lexmodelsv2.UpdateBotLocaleOutput, error) {
	var output lexmodelsv2.UpdateBotLocaleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateIntentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateIntentFuture) Get(ctx workflow.Context) (*lexmodelsv2.UpdateIntentOutput, error) {
	var output lexmodelsv2.UpdateIntentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateSlotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateSlotFuture) Get(ctx workflow.Context) (*lexmodelsv2.UpdateSlotOutput, error) {
	var output lexmodelsv2.UpdateSlotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateSlotTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateSlotTypeFuture) Get(ctx workflow.Context) (*lexmodelsv2.UpdateSlotTypeOutput, error) {
	var output lexmodelsv2.UpdateSlotTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BuildBotLocale(ctx workflow.Context, input *lexmodelsv2.BuildBotLocaleInput) (*lexmodelsv2.BuildBotLocaleOutput, error) {
	var output lexmodelsv2.BuildBotLocaleOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-BuildBotLocale", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BuildBotLocaleAsync(ctx workflow.Context, input *lexmodelsv2.BuildBotLocaleInput) *BuildBotLocaleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-BuildBotLocale", input)
	return &BuildBotLocaleFuture{Future: future}
}

func (a *stub) CreateBot(ctx workflow.Context, input *lexmodelsv2.CreateBotInput) (*lexmodelsv2.CreateBotOutput, error) {
	var output lexmodelsv2.CreateBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBotAsync(ctx workflow.Context, input *lexmodelsv2.CreateBotInput) *CreateBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBot", input)
	return &CreateBotFuture{Future: future}
}

func (a *stub) CreateBotAlias(ctx workflow.Context, input *lexmodelsv2.CreateBotAliasInput) (*lexmodelsv2.CreateBotAliasOutput, error) {
	var output lexmodelsv2.CreateBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBotAliasAsync(ctx workflow.Context, input *lexmodelsv2.CreateBotAliasInput) *CreateBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBotAlias", input)
	return &CreateBotAliasFuture{Future: future}
}

func (a *stub) CreateBotLocale(ctx workflow.Context, input *lexmodelsv2.CreateBotLocaleInput) (*lexmodelsv2.CreateBotLocaleOutput, error) {
	var output lexmodelsv2.CreateBotLocaleOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBotLocale", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBotLocaleAsync(ctx workflow.Context, input *lexmodelsv2.CreateBotLocaleInput) *CreateBotLocaleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBotLocale", input)
	return &CreateBotLocaleFuture{Future: future}
}

func (a *stub) CreateBotVersion(ctx workflow.Context, input *lexmodelsv2.CreateBotVersionInput) (*lexmodelsv2.CreateBotVersionOutput, error) {
	var output lexmodelsv2.CreateBotVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBotVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBotVersionAsync(ctx workflow.Context, input *lexmodelsv2.CreateBotVersionInput) *CreateBotVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateBotVersion", input)
	return &CreateBotVersionFuture{Future: future}
}

func (a *stub) CreateIntent(ctx workflow.Context, input *lexmodelsv2.CreateIntentInput) (*lexmodelsv2.CreateIntentOutput, error) {
	var output lexmodelsv2.CreateIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateIntentAsync(ctx workflow.Context, input *lexmodelsv2.CreateIntentInput) *CreateIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateIntent", input)
	return &CreateIntentFuture{Future: future}
}

func (a *stub) CreateSlot(ctx workflow.Context, input *lexmodelsv2.CreateSlotInput) (*lexmodelsv2.CreateSlotOutput, error) {
	var output lexmodelsv2.CreateSlotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateSlot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSlotAsync(ctx workflow.Context, input *lexmodelsv2.CreateSlotInput) *CreateSlotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateSlot", input)
	return &CreateSlotFuture{Future: future}
}

func (a *stub) CreateSlotType(ctx workflow.Context, input *lexmodelsv2.CreateSlotTypeInput) (*lexmodelsv2.CreateSlotTypeOutput, error) {
	var output lexmodelsv2.CreateSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSlotTypeAsync(ctx workflow.Context, input *lexmodelsv2.CreateSlotTypeInput) *CreateSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-CreateSlotType", input)
	return &CreateSlotTypeFuture{Future: future}
}

func (a *stub) DeleteBot(ctx workflow.Context, input *lexmodelsv2.DeleteBotInput) (*lexmodelsv2.DeleteBotOutput, error) {
	var output lexmodelsv2.DeleteBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotAsync(ctx workflow.Context, input *lexmodelsv2.DeleteBotInput) *DeleteBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBot", input)
	return &DeleteBotFuture{Future: future}
}

func (a *stub) DeleteBotAlias(ctx workflow.Context, input *lexmodelsv2.DeleteBotAliasInput) (*lexmodelsv2.DeleteBotAliasOutput, error) {
	var output lexmodelsv2.DeleteBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotAliasAsync(ctx workflow.Context, input *lexmodelsv2.DeleteBotAliasInput) *DeleteBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBotAlias", input)
	return &DeleteBotAliasFuture{Future: future}
}

func (a *stub) DeleteBotLocale(ctx workflow.Context, input *lexmodelsv2.DeleteBotLocaleInput) (*lexmodelsv2.DeleteBotLocaleOutput, error) {
	var output lexmodelsv2.DeleteBotLocaleOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBotLocale", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotLocaleAsync(ctx workflow.Context, input *lexmodelsv2.DeleteBotLocaleInput) *DeleteBotLocaleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBotLocale", input)
	return &DeleteBotLocaleFuture{Future: future}
}

func (a *stub) DeleteBotVersion(ctx workflow.Context, input *lexmodelsv2.DeleteBotVersionInput) (*lexmodelsv2.DeleteBotVersionOutput, error) {
	var output lexmodelsv2.DeleteBotVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBotVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBotVersionAsync(ctx workflow.Context, input *lexmodelsv2.DeleteBotVersionInput) *DeleteBotVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteBotVersion", input)
	return &DeleteBotVersionFuture{Future: future}
}

func (a *stub) DeleteIntent(ctx workflow.Context, input *lexmodelsv2.DeleteIntentInput) (*lexmodelsv2.DeleteIntentOutput, error) {
	var output lexmodelsv2.DeleteIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIntentAsync(ctx workflow.Context, input *lexmodelsv2.DeleteIntentInput) *DeleteIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteIntent", input)
	return &DeleteIntentFuture{Future: future}
}

func (a *stub) DeleteSlot(ctx workflow.Context, input *lexmodelsv2.DeleteSlotInput) (*lexmodelsv2.DeleteSlotOutput, error) {
	var output lexmodelsv2.DeleteSlotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteSlot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSlotAsync(ctx workflow.Context, input *lexmodelsv2.DeleteSlotInput) *DeleteSlotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteSlot", input)
	return &DeleteSlotFuture{Future: future}
}

func (a *stub) DeleteSlotType(ctx workflow.Context, input *lexmodelsv2.DeleteSlotTypeInput) (*lexmodelsv2.DeleteSlotTypeOutput, error) {
	var output lexmodelsv2.DeleteSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSlotTypeAsync(ctx workflow.Context, input *lexmodelsv2.DeleteSlotTypeInput) *DeleteSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DeleteSlotType", input)
	return &DeleteSlotTypeFuture{Future: future}
}

func (a *stub) DescribeBot(ctx workflow.Context, input *lexmodelsv2.DescribeBotInput) (*lexmodelsv2.DescribeBotOutput, error) {
	var output lexmodelsv2.DescribeBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBotAsync(ctx workflow.Context, input *lexmodelsv2.DescribeBotInput) *DescribeBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBot", input)
	return &DescribeBotFuture{Future: future}
}

func (a *stub) DescribeBotAlias(ctx workflow.Context, input *lexmodelsv2.DescribeBotAliasInput) (*lexmodelsv2.DescribeBotAliasOutput, error) {
	var output lexmodelsv2.DescribeBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBotAliasAsync(ctx workflow.Context, input *lexmodelsv2.DescribeBotAliasInput) *DescribeBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBotAlias", input)
	return &DescribeBotAliasFuture{Future: future}
}

func (a *stub) DescribeBotLocale(ctx workflow.Context, input *lexmodelsv2.DescribeBotLocaleInput) (*lexmodelsv2.DescribeBotLocaleOutput, error) {
	var output lexmodelsv2.DescribeBotLocaleOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBotLocale", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBotLocaleAsync(ctx workflow.Context, input *lexmodelsv2.DescribeBotLocaleInput) *DescribeBotLocaleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBotLocale", input)
	return &DescribeBotLocaleFuture{Future: future}
}

func (a *stub) DescribeBotVersion(ctx workflow.Context, input *lexmodelsv2.DescribeBotVersionInput) (*lexmodelsv2.DescribeBotVersionOutput, error) {
	var output lexmodelsv2.DescribeBotVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBotVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBotVersionAsync(ctx workflow.Context, input *lexmodelsv2.DescribeBotVersionInput) *DescribeBotVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeBotVersion", input)
	return &DescribeBotVersionFuture{Future: future}
}

func (a *stub) DescribeIntent(ctx workflow.Context, input *lexmodelsv2.DescribeIntentInput) (*lexmodelsv2.DescribeIntentOutput, error) {
	var output lexmodelsv2.DescribeIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIntentAsync(ctx workflow.Context, input *lexmodelsv2.DescribeIntentInput) *DescribeIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeIntent", input)
	return &DescribeIntentFuture{Future: future}
}

func (a *stub) DescribeSlot(ctx workflow.Context, input *lexmodelsv2.DescribeSlotInput) (*lexmodelsv2.DescribeSlotOutput, error) {
	var output lexmodelsv2.DescribeSlotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeSlot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSlotAsync(ctx workflow.Context, input *lexmodelsv2.DescribeSlotInput) *DescribeSlotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeSlot", input)
	return &DescribeSlotFuture{Future: future}
}

func (a *stub) DescribeSlotType(ctx workflow.Context, input *lexmodelsv2.DescribeSlotTypeInput) (*lexmodelsv2.DescribeSlotTypeOutput, error) {
	var output lexmodelsv2.DescribeSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSlotTypeAsync(ctx workflow.Context, input *lexmodelsv2.DescribeSlotTypeInput) *DescribeSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-DescribeSlotType", input)
	return &DescribeSlotTypeFuture{Future: future}
}

func (a *stub) ListBotAliases(ctx workflow.Context, input *lexmodelsv2.ListBotAliasesInput) (*lexmodelsv2.ListBotAliasesOutput, error) {
	var output lexmodelsv2.ListBotAliasesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBotAliases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBotAliasesAsync(ctx workflow.Context, input *lexmodelsv2.ListBotAliasesInput) *ListBotAliasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBotAliases", input)
	return &ListBotAliasesFuture{Future: future}
}

func (a *stub) ListBotLocales(ctx workflow.Context, input *lexmodelsv2.ListBotLocalesInput) (*lexmodelsv2.ListBotLocalesOutput, error) {
	var output lexmodelsv2.ListBotLocalesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBotLocales", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBotLocalesAsync(ctx workflow.Context, input *lexmodelsv2.ListBotLocalesInput) *ListBotLocalesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBotLocales", input)
	return &ListBotLocalesFuture{Future: future}
}

func (a *stub) ListBotVersions(ctx workflow.Context, input *lexmodelsv2.ListBotVersionsInput) (*lexmodelsv2.ListBotVersionsOutput, error) {
	var output lexmodelsv2.ListBotVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBotVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBotVersionsAsync(ctx workflow.Context, input *lexmodelsv2.ListBotVersionsInput) *ListBotVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBotVersions", input)
	return &ListBotVersionsFuture{Future: future}
}

func (a *stub) ListBots(ctx workflow.Context, input *lexmodelsv2.ListBotsInput) (*lexmodelsv2.ListBotsOutput, error) {
	var output lexmodelsv2.ListBotsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBots", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBotsAsync(ctx workflow.Context, input *lexmodelsv2.ListBotsInput) *ListBotsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBots", input)
	return &ListBotsFuture{Future: future}
}

func (a *stub) ListBuiltInIntents(ctx workflow.Context, input *lexmodelsv2.ListBuiltInIntentsInput) (*lexmodelsv2.ListBuiltInIntentsOutput, error) {
	var output lexmodelsv2.ListBuiltInIntentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBuiltInIntents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBuiltInIntentsAsync(ctx workflow.Context, input *lexmodelsv2.ListBuiltInIntentsInput) *ListBuiltInIntentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBuiltInIntents", input)
	return &ListBuiltInIntentsFuture{Future: future}
}

func (a *stub) ListBuiltInSlotTypes(ctx workflow.Context, input *lexmodelsv2.ListBuiltInSlotTypesInput) (*lexmodelsv2.ListBuiltInSlotTypesOutput, error) {
	var output lexmodelsv2.ListBuiltInSlotTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBuiltInSlotTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBuiltInSlotTypesAsync(ctx workflow.Context, input *lexmodelsv2.ListBuiltInSlotTypesInput) *ListBuiltInSlotTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListBuiltInSlotTypes", input)
	return &ListBuiltInSlotTypesFuture{Future: future}
}

func (a *stub) ListIntents(ctx workflow.Context, input *lexmodelsv2.ListIntentsInput) (*lexmodelsv2.ListIntentsOutput, error) {
	var output lexmodelsv2.ListIntentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListIntents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListIntentsAsync(ctx workflow.Context, input *lexmodelsv2.ListIntentsInput) *ListIntentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListIntents", input)
	return &ListIntentsFuture{Future: future}
}

func (a *stub) ListSlotTypes(ctx workflow.Context, input *lexmodelsv2.ListSlotTypesInput) (*lexmodelsv2.ListSlotTypesOutput, error) {
	var output lexmodelsv2.ListSlotTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListSlotTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSlotTypesAsync(ctx workflow.Context, input *lexmodelsv2.ListSlotTypesInput) *ListSlotTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListSlotTypes", input)
	return &ListSlotTypesFuture{Future: future}
}

func (a *stub) ListSlots(ctx workflow.Context, input *lexmodelsv2.ListSlotsInput) (*lexmodelsv2.ListSlotsOutput, error) {
	var output lexmodelsv2.ListSlotsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListSlots", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSlotsAsync(ctx workflow.Context, input *lexmodelsv2.ListSlotsInput) *ListSlotsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListSlots", input)
	return &ListSlotsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *lexmodelsv2.ListTagsForResourceInput) (*lexmodelsv2.ListTagsForResourceOutput, error) {
	var output lexmodelsv2.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *lexmodelsv2.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *lexmodelsv2.TagResourceInput) (*lexmodelsv2.TagResourceOutput, error) {
	var output lexmodelsv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *lexmodelsv2.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *lexmodelsv2.UntagResourceInput) (*lexmodelsv2.UntagResourceOutput, error) {
	var output lexmodelsv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *lexmodelsv2.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateBot(ctx workflow.Context, input *lexmodelsv2.UpdateBotInput) (*lexmodelsv2.UpdateBotOutput, error) {
	var output lexmodelsv2.UpdateBotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateBot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBotAsync(ctx workflow.Context, input *lexmodelsv2.UpdateBotInput) *UpdateBotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateBot", input)
	return &UpdateBotFuture{Future: future}
}

func (a *stub) UpdateBotAlias(ctx workflow.Context, input *lexmodelsv2.UpdateBotAliasInput) (*lexmodelsv2.UpdateBotAliasOutput, error) {
	var output lexmodelsv2.UpdateBotAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateBotAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBotAliasAsync(ctx workflow.Context, input *lexmodelsv2.UpdateBotAliasInput) *UpdateBotAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateBotAlias", input)
	return &UpdateBotAliasFuture{Future: future}
}

func (a *stub) UpdateBotLocale(ctx workflow.Context, input *lexmodelsv2.UpdateBotLocaleInput) (*lexmodelsv2.UpdateBotLocaleOutput, error) {
	var output lexmodelsv2.UpdateBotLocaleOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateBotLocale", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBotLocaleAsync(ctx workflow.Context, input *lexmodelsv2.UpdateBotLocaleInput) *UpdateBotLocaleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateBotLocale", input)
	return &UpdateBotLocaleFuture{Future: future}
}

func (a *stub) UpdateIntent(ctx workflow.Context, input *lexmodelsv2.UpdateIntentInput) (*lexmodelsv2.UpdateIntentOutput, error) {
	var output lexmodelsv2.UpdateIntentOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateIntent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateIntentAsync(ctx workflow.Context, input *lexmodelsv2.UpdateIntentInput) *UpdateIntentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateIntent", input)
	return &UpdateIntentFuture{Future: future}
}

func (a *stub) UpdateSlot(ctx workflow.Context, input *lexmodelsv2.UpdateSlotInput) (*lexmodelsv2.UpdateSlotOutput, error) {
	var output lexmodelsv2.UpdateSlotOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateSlot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSlotAsync(ctx workflow.Context, input *lexmodelsv2.UpdateSlotInput) *UpdateSlotFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateSlot", input)
	return &UpdateSlotFuture{Future: future}
}

func (a *stub) UpdateSlotType(ctx workflow.Context, input *lexmodelsv2.UpdateSlotTypeInput) (*lexmodelsv2.UpdateSlotTypeOutput, error) {
	var output lexmodelsv2.UpdateSlotTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateSlotType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSlotTypeAsync(ctx workflow.Context, input *lexmodelsv2.UpdateSlotTypeInput) *UpdateSlotTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lexmodelsv2-UpdateSlotType", input)
	return &UpdateSlotTypeFuture{Future: future}
}