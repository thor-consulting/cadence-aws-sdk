// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package acmstub

import (
	"github.com/aws/aws-sdk-go/service/acm"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToCertificate(ctx workflow.Context, input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error)
	AddTagsToCertificateAsync(ctx workflow.Context, input *acm.AddTagsToCertificateInput) *AddTagsToCertificateFuture

	DeleteCertificate(ctx workflow.Context, input *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error)
	DeleteCertificateAsync(ctx workflow.Context, input *acm.DeleteCertificateInput) *DeleteCertificateFuture

	DescribeCertificate(ctx workflow.Context, input *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error)
	DescribeCertificateAsync(ctx workflow.Context, input *acm.DescribeCertificateInput) *DescribeCertificateFuture

	ExportCertificate(ctx workflow.Context, input *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error)
	ExportCertificateAsync(ctx workflow.Context, input *acm.ExportCertificateInput) *ExportCertificateFuture

	GetCertificate(ctx workflow.Context, input *acm.GetCertificateInput) (*acm.GetCertificateOutput, error)
	GetCertificateAsync(ctx workflow.Context, input *acm.GetCertificateInput) *GetCertificateFuture

	ImportCertificate(ctx workflow.Context, input *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error)
	ImportCertificateAsync(ctx workflow.Context, input *acm.ImportCertificateInput) *ImportCertificateFuture

	ListCertificates(ctx workflow.Context, input *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error)
	ListCertificatesAsync(ctx workflow.Context, input *acm.ListCertificatesInput) *ListCertificatesFuture

	ListTagsForCertificate(ctx workflow.Context, input *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error)
	ListTagsForCertificateAsync(ctx workflow.Context, input *acm.ListTagsForCertificateInput) *ListTagsForCertificateFuture

	RemoveTagsFromCertificate(ctx workflow.Context, input *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error)
	RemoveTagsFromCertificateAsync(ctx workflow.Context, input *acm.RemoveTagsFromCertificateInput) *RemoveTagsFromCertificateFuture

	RenewCertificate(ctx workflow.Context, input *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error)
	RenewCertificateAsync(ctx workflow.Context, input *acm.RenewCertificateInput) *RenewCertificateFuture

	RequestCertificate(ctx workflow.Context, input *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error)
	RequestCertificateAsync(ctx workflow.Context, input *acm.RequestCertificateInput) *RequestCertificateFuture

	ResendValidationEmail(ctx workflow.Context, input *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error)
	ResendValidationEmailAsync(ctx workflow.Context, input *acm.ResendValidationEmailInput) *ResendValidationEmailFuture

	UpdateCertificateOptions(ctx workflow.Context, input *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error)
	UpdateCertificateOptionsAsync(ctx workflow.Context, input *acm.UpdateCertificateOptionsInput) *UpdateCertificateOptionsFuture

	WaitUntilCertificateValidated(ctx workflow.Context, input *acm.DescribeCertificateInput) error
	WaitUntilCertificateValidatedAsync(ctx workflow.Context, input *acm.DescribeCertificateInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
