// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package frauddetectorstub

import (
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchCreateVariable(ctx workflow.Context, input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error)
	BatchCreateVariableAsync(ctx workflow.Context, input *frauddetector.BatchCreateVariableInput) *BatchCreateVariableFuture

	BatchGetVariable(ctx workflow.Context, input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error)
	BatchGetVariableAsync(ctx workflow.Context, input *frauddetector.BatchGetVariableInput) *BatchGetVariableFuture

	CreateDetectorVersion(ctx workflow.Context, input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error)
	CreateDetectorVersionAsync(ctx workflow.Context, input *frauddetector.CreateDetectorVersionInput) *CreateDetectorVersionFuture

	CreateModel(ctx workflow.Context, input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error)
	CreateModelAsync(ctx workflow.Context, input *frauddetector.CreateModelInput) *CreateModelFuture

	CreateModelVersion(ctx workflow.Context, input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error)
	CreateModelVersionAsync(ctx workflow.Context, input *frauddetector.CreateModelVersionInput) *CreateModelVersionFuture

	CreateRule(ctx workflow.Context, input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *frauddetector.CreateRuleInput) *CreateRuleFuture

	CreateVariable(ctx workflow.Context, input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error)
	CreateVariableAsync(ctx workflow.Context, input *frauddetector.CreateVariableInput) *CreateVariableFuture

	DeleteDetector(ctx workflow.Context, input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error)
	DeleteDetectorAsync(ctx workflow.Context, input *frauddetector.DeleteDetectorInput) *DeleteDetectorFuture

	DeleteDetectorVersion(ctx workflow.Context, input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error)
	DeleteDetectorVersionAsync(ctx workflow.Context, input *frauddetector.DeleteDetectorVersionInput) *DeleteDetectorVersionFuture

	DeleteEvent(ctx workflow.Context, input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error)
	DeleteEventAsync(ctx workflow.Context, input *frauddetector.DeleteEventInput) *DeleteEventFuture

	DeleteRule(ctx workflow.Context, input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *frauddetector.DeleteRuleInput) *DeleteRuleFuture

	DescribeDetector(ctx workflow.Context, input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error)
	DescribeDetectorAsync(ctx workflow.Context, input *frauddetector.DescribeDetectorInput) *DescribeDetectorFuture

	DescribeModelVersions(ctx workflow.Context, input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error)
	DescribeModelVersionsAsync(ctx workflow.Context, input *frauddetector.DescribeModelVersionsInput) *DescribeModelVersionsFuture

	GetDetectorVersion(ctx workflow.Context, input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error)
	GetDetectorVersionAsync(ctx workflow.Context, input *frauddetector.GetDetectorVersionInput) *GetDetectorVersionFuture

	GetDetectors(ctx workflow.Context, input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error)
	GetDetectorsAsync(ctx workflow.Context, input *frauddetector.GetDetectorsInput) *GetDetectorsFuture

	GetEntityTypes(ctx workflow.Context, input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error)
	GetEntityTypesAsync(ctx workflow.Context, input *frauddetector.GetEntityTypesInput) *GetEntityTypesFuture

	GetEventPrediction(ctx workflow.Context, input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error)
	GetEventPredictionAsync(ctx workflow.Context, input *frauddetector.GetEventPredictionInput) *GetEventPredictionFuture

	GetEventTypes(ctx workflow.Context, input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error)
	GetEventTypesAsync(ctx workflow.Context, input *frauddetector.GetEventTypesInput) *GetEventTypesFuture

	GetExternalModels(ctx workflow.Context, input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error)
	GetExternalModelsAsync(ctx workflow.Context, input *frauddetector.GetExternalModelsInput) *GetExternalModelsFuture

	GetKMSEncryptionKey(ctx workflow.Context, input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error)
	GetKMSEncryptionKeyAsync(ctx workflow.Context, input *frauddetector.GetKMSEncryptionKeyInput) *GetKMSEncryptionKeyFuture

	GetLabels(ctx workflow.Context, input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error)
	GetLabelsAsync(ctx workflow.Context, input *frauddetector.GetLabelsInput) *GetLabelsFuture

	GetModelVersion(ctx workflow.Context, input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error)
	GetModelVersionAsync(ctx workflow.Context, input *frauddetector.GetModelVersionInput) *GetModelVersionFuture

	GetModels(ctx workflow.Context, input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *frauddetector.GetModelsInput) *GetModelsFuture

	GetOutcomes(ctx workflow.Context, input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error)
	GetOutcomesAsync(ctx workflow.Context, input *frauddetector.GetOutcomesInput) *GetOutcomesFuture

	GetRules(ctx workflow.Context, input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error)
	GetRulesAsync(ctx workflow.Context, input *frauddetector.GetRulesInput) *GetRulesFuture

	GetVariables(ctx workflow.Context, input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error)
	GetVariablesAsync(ctx workflow.Context, input *frauddetector.GetVariablesInput) *GetVariablesFuture

	ListTagsForResource(ctx workflow.Context, input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *frauddetector.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutDetector(ctx workflow.Context, input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error)
	PutDetectorAsync(ctx workflow.Context, input *frauddetector.PutDetectorInput) *PutDetectorFuture

	PutEntityType(ctx workflow.Context, input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error)
	PutEntityTypeAsync(ctx workflow.Context, input *frauddetector.PutEntityTypeInput) *PutEntityTypeFuture

	PutEventType(ctx workflow.Context, input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error)
	PutEventTypeAsync(ctx workflow.Context, input *frauddetector.PutEventTypeInput) *PutEventTypeFuture

	PutExternalModel(ctx workflow.Context, input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error)
	PutExternalModelAsync(ctx workflow.Context, input *frauddetector.PutExternalModelInput) *PutExternalModelFuture

	PutKMSEncryptionKey(ctx workflow.Context, input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error)
	PutKMSEncryptionKeyAsync(ctx workflow.Context, input *frauddetector.PutKMSEncryptionKeyInput) *PutKMSEncryptionKeyFuture

	PutLabel(ctx workflow.Context, input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error)
	PutLabelAsync(ctx workflow.Context, input *frauddetector.PutLabelInput) *PutLabelFuture

	PutOutcome(ctx workflow.Context, input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error)
	PutOutcomeAsync(ctx workflow.Context, input *frauddetector.PutOutcomeInput) *PutOutcomeFuture

	TagResource(ctx workflow.Context, input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *frauddetector.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *frauddetector.UntagResourceInput) *UntagResourceFuture

	UpdateDetectorVersion(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error)
	UpdateDetectorVersionAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionInput) *UpdateDetectorVersionFuture

	UpdateDetectorVersionMetadata(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error)
	UpdateDetectorVersionMetadataAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) *UpdateDetectorVersionMetadataFuture

	UpdateDetectorVersionStatus(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error)
	UpdateDetectorVersionStatusAsync(ctx workflow.Context, input *frauddetector.UpdateDetectorVersionStatusInput) *UpdateDetectorVersionStatusFuture

	UpdateModel(ctx workflow.Context, input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error)
	UpdateModelAsync(ctx workflow.Context, input *frauddetector.UpdateModelInput) *UpdateModelFuture

	UpdateModelVersion(ctx workflow.Context, input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error)
	UpdateModelVersionAsync(ctx workflow.Context, input *frauddetector.UpdateModelVersionInput) *UpdateModelVersionFuture

	UpdateModelVersionStatus(ctx workflow.Context, input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error)
	UpdateModelVersionStatusAsync(ctx workflow.Context, input *frauddetector.UpdateModelVersionStatusInput) *UpdateModelVersionStatusFuture

	UpdateRuleMetadata(ctx workflow.Context, input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error)
	UpdateRuleMetadataAsync(ctx workflow.Context, input *frauddetector.UpdateRuleMetadataInput) *UpdateRuleMetadataFuture

	UpdateRuleVersion(ctx workflow.Context, input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error)
	UpdateRuleVersionAsync(ctx workflow.Context, input *frauddetector.UpdateRuleVersionInput) *UpdateRuleVersionFuture

	UpdateVariable(ctx workflow.Context, input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error)
	UpdateVariableAsync(ctx workflow.Context, input *frauddetector.UpdateVariableInput) *UpdateVariableFuture
}

func NewClient() Client {
	return &stub{}
}
