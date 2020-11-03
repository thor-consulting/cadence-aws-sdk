// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package textractstub

import (
	"github.com/aws/aws-sdk-go/service/textract"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AnalyzeDocument(ctx workflow.Context, input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error)
	AnalyzeDocumentAsync(ctx workflow.Context, input *textract.AnalyzeDocumentInput) *TextractAnalyzeDocumentFuture

	DetectDocumentText(ctx workflow.Context, input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error)
	DetectDocumentTextAsync(ctx workflow.Context, input *textract.DetectDocumentTextInput) *TextractDetectDocumentTextFuture

	GetDocumentAnalysis(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error)
	GetDocumentAnalysisAsync(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) *TextractGetDocumentAnalysisFuture

	GetDocumentTextDetection(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error)
	GetDocumentTextDetectionAsync(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) *TextractGetDocumentTextDetectionFuture

	StartDocumentAnalysis(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error)
	StartDocumentAnalysisAsync(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) *TextractStartDocumentAnalysisFuture

	StartDocumentTextDetection(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error)
	StartDocumentTextDetectionAsync(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) *TextractStartDocumentTextDetectionFuture
}

func NewClient() Client {
	return &stub{}
}
