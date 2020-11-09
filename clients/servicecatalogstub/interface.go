// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package servicecatalogstub

import (
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptPortfolioShare(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error)
	AcceptPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) *AcceptPortfolioShareFuture

	AssociateBudgetWithResource(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error)
	AssociateBudgetWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) *AssociateBudgetWithResourceFuture

	AssociatePrincipalWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error)
	AssociatePrincipalWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) *AssociatePrincipalWithPortfolioFuture

	AssociateProductWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error)
	AssociateProductWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) *AssociateProductWithPortfolioFuture

	AssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error)
	AssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) *AssociateServiceActionWithProvisioningArtifactFuture

	AssociateTagOptionWithResource(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error)
	AssociateTagOptionWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) *AssociateTagOptionWithResourceFuture

	BatchAssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error)
	BatchAssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) *BatchAssociateServiceActionWithProvisioningArtifactFuture

	BatchDisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error)
	BatchDisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) *BatchDisassociateServiceActionFromProvisioningArtifactFuture

	CopyProduct(ctx workflow.Context, input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error)
	CopyProductAsync(ctx workflow.Context, input *servicecatalog.CopyProductInput) *CopyProductFuture

	CreateConstraint(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error)
	CreateConstraintAsync(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) *CreateConstraintFuture

	CreatePortfolio(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error)
	CreatePortfolioAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) *CreatePortfolioFuture

	CreatePortfolioShare(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error)
	CreatePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) *CreatePortfolioShareFuture

	CreateProduct(ctx workflow.Context, input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error)
	CreateProductAsync(ctx workflow.Context, input *servicecatalog.CreateProductInput) *CreateProductFuture

	CreateProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error)
	CreateProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) *CreateProvisionedProductPlanFuture

	CreateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error)
	CreateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) *CreateProvisioningArtifactFuture

	CreateServiceAction(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error)
	CreateServiceActionAsync(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) *CreateServiceActionFuture

	CreateTagOption(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error)
	CreateTagOptionAsync(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) *CreateTagOptionFuture

	DeleteConstraint(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error)
	DeleteConstraintAsync(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) *DeleteConstraintFuture

	DeletePortfolio(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error)
	DeletePortfolioAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) *DeletePortfolioFuture

	DeletePortfolioShare(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error)
	DeletePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) *DeletePortfolioShareFuture

	DeleteProduct(ctx workflow.Context, input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error)
	DeleteProductAsync(ctx workflow.Context, input *servicecatalog.DeleteProductInput) *DeleteProductFuture

	DeleteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error)
	DeleteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) *DeleteProvisionedProductPlanFuture

	DeleteProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error)
	DeleteProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) *DeleteProvisioningArtifactFuture

	DeleteServiceAction(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error)
	DeleteServiceActionAsync(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) *DeleteServiceActionFuture

	DeleteTagOption(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error)
	DeleteTagOptionAsync(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) *DeleteTagOptionFuture

	DescribeConstraint(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error)
	DescribeConstraintAsync(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) *DescribeConstraintFuture

	DescribeCopyProductStatus(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error)
	DescribeCopyProductStatusAsync(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) *DescribeCopyProductStatusFuture

	DescribePortfolio(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error)
	DescribePortfolioAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) *DescribePortfolioFuture

	DescribePortfolioShareStatus(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error)
	DescribePortfolioShareStatusAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) *DescribePortfolioShareStatusFuture

	DescribeProduct(ctx workflow.Context, input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error)
	DescribeProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProductInput) *DescribeProductFuture

	DescribeProductAsAdmin(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error)
	DescribeProductAsAdminAsync(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) *DescribeProductAsAdminFuture

	DescribeProductView(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error)
	DescribeProductViewAsync(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) *DescribeProductViewFuture

	DescribeProvisionedProduct(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error)
	DescribeProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) *DescribeProvisionedProductFuture

	DescribeProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error)
	DescribeProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) *DescribeProvisionedProductPlanFuture

	DescribeProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error)
	DescribeProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) *DescribeProvisioningArtifactFuture

	DescribeProvisioningParameters(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error)
	DescribeProvisioningParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) *DescribeProvisioningParametersFuture

	DescribeRecord(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error)
	DescribeRecordAsync(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) *DescribeRecordFuture

	DescribeServiceAction(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error)
	DescribeServiceActionAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) *DescribeServiceActionFuture

	DescribeServiceActionExecutionParameters(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error)
	DescribeServiceActionExecutionParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) *DescribeServiceActionExecutionParametersFuture

	DescribeTagOption(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error)
	DescribeTagOptionAsync(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) *DescribeTagOptionFuture

	DisableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error)
	DisableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) *DisableAWSOrganizationsAccessFuture

	DisassociateBudgetFromResource(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error)
	DisassociateBudgetFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) *DisassociateBudgetFromResourceFuture

	DisassociatePrincipalFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error)
	DisassociatePrincipalFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) *DisassociatePrincipalFromPortfolioFuture

	DisassociateProductFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error)
	DisassociateProductFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) *DisassociateProductFromPortfolioFuture

	DisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error)
	DisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) *DisassociateServiceActionFromProvisioningArtifactFuture

	DisassociateTagOptionFromResource(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error)
	DisassociateTagOptionFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) *DisassociateTagOptionFromResourceFuture

	EnableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error)
	EnableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) *EnableAWSOrganizationsAccessFuture

	ExecuteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error)
	ExecuteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) *ExecuteProvisionedProductPlanFuture

	ExecuteProvisionedProductServiceAction(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error)
	ExecuteProvisionedProductServiceActionAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) *ExecuteProvisionedProductServiceActionFuture

	GetAWSOrganizationsAccessStatus(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error)
	GetAWSOrganizationsAccessStatusAsync(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) *GetAWSOrganizationsAccessStatusFuture

	GetProvisionedProductOutputs(ctx workflow.Context, input *servicecatalog.GetProvisionedProductOutputsInput) (*servicecatalog.GetProvisionedProductOutputsOutput, error)
	GetProvisionedProductOutputsAsync(ctx workflow.Context, input *servicecatalog.GetProvisionedProductOutputsInput) *GetProvisionedProductOutputsFuture

	ListAcceptedPortfolioShares(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)
	ListAcceptedPortfolioSharesAsync(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) *ListAcceptedPortfolioSharesFuture

	ListBudgetsForResource(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error)
	ListBudgetsForResourceAsync(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) *ListBudgetsForResourceFuture

	ListConstraintsForPortfolio(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error)
	ListConstraintsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) *ListConstraintsForPortfolioFuture

	ListLaunchPaths(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error)
	ListLaunchPathsAsync(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) *ListLaunchPathsFuture

	ListOrganizationPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error)
	ListOrganizationPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) *ListOrganizationPortfolioAccessFuture

	ListPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error)
	ListPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) *ListPortfolioAccessFuture

	ListPortfolios(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error)
	ListPortfoliosAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) *ListPortfoliosFuture

	ListPortfoliosForProduct(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error)
	ListPortfoliosForProductAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) *ListPortfoliosForProductFuture

	ListPrincipalsForPortfolio(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)
	ListPrincipalsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) *ListPrincipalsForPortfolioFuture

	ListProvisionedProductPlans(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error)
	ListProvisionedProductPlansAsync(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) *ListProvisionedProductPlansFuture

	ListProvisioningArtifacts(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error)
	ListProvisioningArtifactsAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) *ListProvisioningArtifactsFuture

	ListProvisioningArtifactsForServiceAction(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error)
	ListProvisioningArtifactsForServiceActionAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) *ListProvisioningArtifactsForServiceActionFuture

	ListRecordHistory(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error)
	ListRecordHistoryAsync(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) *ListRecordHistoryFuture

	ListResourcesForTagOption(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error)
	ListResourcesForTagOptionAsync(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) *ListResourcesForTagOptionFuture

	ListServiceActions(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error)
	ListServiceActionsAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) *ListServiceActionsFuture

	ListServiceActionsForProvisioningArtifact(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error)
	ListServiceActionsForProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) *ListServiceActionsForProvisioningArtifactFuture

	ListStackInstancesForProvisionedProduct(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error)
	ListStackInstancesForProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) *ListStackInstancesForProvisionedProductFuture

	ListTagOptions(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error)
	ListTagOptionsAsync(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) *ListTagOptionsFuture

	ProvisionProduct(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error)
	ProvisionProductAsync(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) *ProvisionProductFuture

	RejectPortfolioShare(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error)
	RejectPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) *RejectPortfolioShareFuture

	ScanProvisionedProducts(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error)
	ScanProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) *ScanProvisionedProductsFuture

	SearchProducts(ctx workflow.Context, input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error)
	SearchProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProductsInput) *SearchProductsFuture

	SearchProductsAsAdmin(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error)
	SearchProductsAsAdminAsync(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) *SearchProductsAsAdminFuture

	SearchProvisionedProducts(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error)
	SearchProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) *SearchProvisionedProductsFuture

	TerminateProvisionedProduct(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error)
	TerminateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) *TerminateProvisionedProductFuture

	UpdateConstraint(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error)
	UpdateConstraintAsync(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) *UpdateConstraintFuture

	UpdatePortfolio(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error)
	UpdatePortfolioAsync(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) *UpdatePortfolioFuture

	UpdateProduct(ctx workflow.Context, input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error)
	UpdateProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProductInput) *UpdateProductFuture

	UpdateProvisionedProduct(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error)
	UpdateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) *UpdateProvisionedProductFuture

	UpdateProvisionedProductProperties(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error)
	UpdateProvisionedProductPropertiesAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) *UpdateProvisionedProductPropertiesFuture

	UpdateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error)
	UpdateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) *UpdateProvisioningArtifactFuture

	UpdateServiceAction(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error)
	UpdateServiceActionAsync(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) *UpdateServiceActionFuture

	UpdateTagOption(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error)
	UpdateTagOptionAsync(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) *UpdateTagOptionFuture
}

func NewClient() Client {
	return &stub{}
}
