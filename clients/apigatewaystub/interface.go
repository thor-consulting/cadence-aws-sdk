// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package apigatewaystub

import (
	"github.com/aws/aws-sdk-go/service/apigateway"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApiKey(ctx workflow.Context, input *apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error)
	CreateApiKeyAsync(ctx workflow.Context, input *apigateway.CreateApiKeyInput) *CreateApiKeyFuture

	CreateAuthorizer(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error)
	CreateAuthorizerAsync(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) *CreateAuthorizerFuture

	CreateBasePathMapping(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error)
	CreateBasePathMappingAsync(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) *CreateBasePathMappingFuture

	CreateDeployment(ctx workflow.Context, input *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error)
	CreateDeploymentAsync(ctx workflow.Context, input *apigateway.CreateDeploymentInput) *CreateDeploymentFuture

	CreateDocumentationPart(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) (*apigateway.DocumentationPart, error)
	CreateDocumentationPartAsync(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) *CreateDocumentationPartFuture

	CreateDocumentationVersion(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	CreateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) *CreateDocumentationVersionFuture

	CreateDomainName(ctx workflow.Context, input *apigateway.CreateDomainNameInput) (*apigateway.DomainName, error)
	CreateDomainNameAsync(ctx workflow.Context, input *apigateway.CreateDomainNameInput) *CreateDomainNameFuture

	CreateModel(ctx workflow.Context, input *apigateway.CreateModelInput) (*apigateway.Model, error)
	CreateModelAsync(ctx workflow.Context, input *apigateway.CreateModelInput) *CreateModelFuture

	CreateRequestValidator(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	CreateRequestValidatorAsync(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) *CreateRequestValidatorFuture

	CreateResource(ctx workflow.Context, input *apigateway.CreateResourceInput) (*apigateway.Resource, error)
	CreateResourceAsync(ctx workflow.Context, input *apigateway.CreateResourceInput) *CreateResourceFuture

	CreateRestApi(ctx workflow.Context, input *apigateway.CreateRestApiInput) (*apigateway.RestApi, error)
	CreateRestApiAsync(ctx workflow.Context, input *apigateway.CreateRestApiInput) *CreateRestApiFuture

	CreateStage(ctx workflow.Context, input *apigateway.CreateStageInput) (*apigateway.Stage, error)
	CreateStageAsync(ctx workflow.Context, input *apigateway.CreateStageInput) *CreateStageFuture

	CreateUsagePlan(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) (*apigateway.UsagePlan, error)
	CreateUsagePlanAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) *CreateUsagePlanFuture

	CreateUsagePlanKey(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) (*apigateway.UsagePlanKey, error)
	CreateUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) *CreateUsagePlanKeyFuture

	CreateVpcLink(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	CreateVpcLinkAsync(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) *CreateVpcLinkFuture

	DeleteApiKey(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error)
	DeleteApiKeyAsync(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) *DeleteApiKeyFuture

	DeleteAuthorizer(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error)
	DeleteAuthorizerAsync(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) *DeleteAuthorizerFuture

	DeleteBasePathMapping(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error)
	DeleteBasePathMappingAsync(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) *DeleteBasePathMappingFuture

	DeleteClientCertificate(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error)
	DeleteClientCertificateAsync(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) *DeleteClientCertificateFuture

	DeleteDeployment(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error)
	DeleteDeploymentAsync(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) *DeleteDeploymentFuture

	DeleteDocumentationPart(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) (*apigateway.DeleteDocumentationPartOutput, error)
	DeleteDocumentationPartAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) *DeleteDocumentationPartFuture

	DeleteDocumentationVersion(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) (*apigateway.DeleteDocumentationVersionOutput, error)
	DeleteDocumentationVersionAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) *DeleteDocumentationVersionFuture

	DeleteDomainName(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error)
	DeleteDomainNameAsync(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) *DeleteDomainNameFuture

	DeleteGatewayResponse(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) (*apigateway.DeleteGatewayResponseOutput, error)
	DeleteGatewayResponseAsync(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) *DeleteGatewayResponseFuture

	DeleteIntegration(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error)
	DeleteIntegrationAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) *DeleteIntegrationFuture

	DeleteIntegrationResponse(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) *DeleteIntegrationResponseFuture

	DeleteMethod(ctx workflow.Context, input *apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error)
	DeleteMethodAsync(ctx workflow.Context, input *apigateway.DeleteMethodInput) *DeleteMethodFuture

	DeleteMethodResponse(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error)
	DeleteMethodResponseAsync(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) *DeleteMethodResponseFuture

	DeleteModel(ctx workflow.Context, input *apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *apigateway.DeleteModelInput) *DeleteModelFuture

	DeleteRequestValidator(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) (*apigateway.DeleteRequestValidatorOutput, error)
	DeleteRequestValidatorAsync(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) *DeleteRequestValidatorFuture

	DeleteResource(ctx workflow.Context, input *apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error)
	DeleteResourceAsync(ctx workflow.Context, input *apigateway.DeleteResourceInput) *DeleteResourceFuture

	DeleteRestApi(ctx workflow.Context, input *apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error)
	DeleteRestApiAsync(ctx workflow.Context, input *apigateway.DeleteRestApiInput) *DeleteRestApiFuture

	DeleteStage(ctx workflow.Context, input *apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error)
	DeleteStageAsync(ctx workflow.Context, input *apigateway.DeleteStageInput) *DeleteStageFuture

	DeleteUsagePlan(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) (*apigateway.DeleteUsagePlanOutput, error)
	DeleteUsagePlanAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) *DeleteUsagePlanFuture

	DeleteUsagePlanKey(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) (*apigateway.DeleteUsagePlanKeyOutput, error)
	DeleteUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) *DeleteUsagePlanKeyFuture

	DeleteVpcLink(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) (*apigateway.DeleteVpcLinkOutput, error)
	DeleteVpcLinkAsync(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) *DeleteVpcLinkFuture

	FlushStageAuthorizersCache(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error)
	FlushStageAuthorizersCacheAsync(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) *FlushStageAuthorizersCacheFuture

	FlushStageCache(ctx workflow.Context, input *apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error)
	FlushStageCacheAsync(ctx workflow.Context, input *apigateway.FlushStageCacheInput) *FlushStageCacheFuture

	GenerateClientCertificate(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error)
	GenerateClientCertificateAsync(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) *GenerateClientCertificateFuture

	GetAccount(ctx workflow.Context, input *apigateway.GetAccountInput) (*apigateway.Account, error)
	GetAccountAsync(ctx workflow.Context, input *apigateway.GetAccountInput) *GetAccountFuture

	GetApiKey(ctx workflow.Context, input *apigateway.GetApiKeyInput) (*apigateway.ApiKey, error)
	GetApiKeyAsync(ctx workflow.Context, input *apigateway.GetApiKeyInput) *GetApiKeyFuture

	GetApiKeys(ctx workflow.Context, input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error)
	GetApiKeysAsync(ctx workflow.Context, input *apigateway.GetApiKeysInput) *GetApiKeysFuture

	GetAuthorizer(ctx workflow.Context, input *apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error)
	GetAuthorizerAsync(ctx workflow.Context, input *apigateway.GetAuthorizerInput) *GetAuthorizerFuture

	GetAuthorizers(ctx workflow.Context, input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error)
	GetAuthorizersAsync(ctx workflow.Context, input *apigateway.GetAuthorizersInput) *GetAuthorizersFuture

	GetBasePathMapping(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error)
	GetBasePathMappingAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) *GetBasePathMappingFuture

	GetBasePathMappings(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error)
	GetBasePathMappingsAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) *GetBasePathMappingsFuture

	GetClientCertificate(ctx workflow.Context, input *apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error)
	GetClientCertificateAsync(ctx workflow.Context, input *apigateway.GetClientCertificateInput) *GetClientCertificateFuture

	GetClientCertificates(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error)
	GetClientCertificatesAsync(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) *GetClientCertificatesFuture

	GetDeployment(ctx workflow.Context, input *apigateway.GetDeploymentInput) (*apigateway.Deployment, error)
	GetDeploymentAsync(ctx workflow.Context, input *apigateway.GetDeploymentInput) *GetDeploymentFuture

	GetDeployments(ctx workflow.Context, input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error)
	GetDeploymentsAsync(ctx workflow.Context, input *apigateway.GetDeploymentsInput) *GetDeploymentsFuture

	GetDocumentationPart(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) (*apigateway.DocumentationPart, error)
	GetDocumentationPartAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) *GetDocumentationPartFuture

	GetDocumentationParts(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationPartsAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) *GetDocumentationPartsFuture

	GetDocumentationVersion(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	GetDocumentationVersionAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) *GetDocumentationVersionFuture

	GetDocumentationVersions(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDocumentationVersionsAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) *GetDocumentationVersionsFuture

	GetDomainName(ctx workflow.Context, input *apigateway.GetDomainNameInput) (*apigateway.DomainName, error)
	GetDomainNameAsync(ctx workflow.Context, input *apigateway.GetDomainNameInput) *GetDomainNameFuture

	GetDomainNames(ctx workflow.Context, input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error)
	GetDomainNamesAsync(ctx workflow.Context, input *apigateway.GetDomainNamesInput) *GetDomainNamesFuture

	GetExport(ctx workflow.Context, input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error)
	GetExportAsync(ctx workflow.Context, input *apigateway.GetExportInput) *GetExportFuture

	GetGatewayResponse(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	GetGatewayResponseAsync(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) *GetGatewayResponseFuture

	GetGatewayResponses(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error)
	GetGatewayResponsesAsync(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) *GetGatewayResponsesFuture

	GetIntegration(ctx workflow.Context, input *apigateway.GetIntegrationInput) (*apigateway.Integration, error)
	GetIntegrationAsync(ctx workflow.Context, input *apigateway.GetIntegrationInput) *GetIntegrationFuture

	GetIntegrationResponse(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	GetIntegrationResponseAsync(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) *GetIntegrationResponseFuture

	GetMethod(ctx workflow.Context, input *apigateway.GetMethodInput) (*apigateway.Method, error)
	GetMethodAsync(ctx workflow.Context, input *apigateway.GetMethodInput) *GetMethodFuture

	GetMethodResponse(ctx workflow.Context, input *apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error)
	GetMethodResponseAsync(ctx workflow.Context, input *apigateway.GetMethodResponseInput) *GetMethodResponseFuture

	GetModel(ctx workflow.Context, input *apigateway.GetModelInput) (*apigateway.Model, error)
	GetModelAsync(ctx workflow.Context, input *apigateway.GetModelInput) *GetModelFuture

	GetModelTemplate(ctx workflow.Context, input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error)
	GetModelTemplateAsync(ctx workflow.Context, input *apigateway.GetModelTemplateInput) *GetModelTemplateFuture

	GetModels(ctx workflow.Context, input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *apigateway.GetModelsInput) *GetModelsFuture

	GetRequestValidator(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	GetRequestValidatorAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) *GetRequestValidatorFuture

	GetRequestValidators(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error)
	GetRequestValidatorsAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) *GetRequestValidatorsFuture

	GetResource(ctx workflow.Context, input *apigateway.GetResourceInput) (*apigateway.Resource, error)
	GetResourceAsync(ctx workflow.Context, input *apigateway.GetResourceInput) *GetResourceFuture

	GetResources(ctx workflow.Context, input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error)
	GetResourcesAsync(ctx workflow.Context, input *apigateway.GetResourcesInput) *GetResourcesFuture

	GetRestApi(ctx workflow.Context, input *apigateway.GetRestApiInput) (*apigateway.RestApi, error)
	GetRestApiAsync(ctx workflow.Context, input *apigateway.GetRestApiInput) *GetRestApiFuture

	GetRestApis(ctx workflow.Context, input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error)
	GetRestApisAsync(ctx workflow.Context, input *apigateway.GetRestApisInput) *GetRestApisFuture

	GetSdk(ctx workflow.Context, input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error)
	GetSdkAsync(ctx workflow.Context, input *apigateway.GetSdkInput) *GetSdkFuture

	GetSdkType(ctx workflow.Context, input *apigateway.GetSdkTypeInput) (*apigateway.SdkType, error)
	GetSdkTypeAsync(ctx workflow.Context, input *apigateway.GetSdkTypeInput) *GetSdkTypeFuture

	GetSdkTypes(ctx workflow.Context, input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error)
	GetSdkTypesAsync(ctx workflow.Context, input *apigateway.GetSdkTypesInput) *GetSdkTypesFuture

	GetStage(ctx workflow.Context, input *apigateway.GetStageInput) (*apigateway.Stage, error)
	GetStageAsync(ctx workflow.Context, input *apigateway.GetStageInput) *GetStageFuture

	GetStages(ctx workflow.Context, input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error)
	GetStagesAsync(ctx workflow.Context, input *apigateway.GetStagesInput) *GetStagesFuture

	GetTags(ctx workflow.Context, input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *apigateway.GetTagsInput) *GetTagsFuture

	GetUsage(ctx workflow.Context, input *apigateway.GetUsageInput) (*apigateway.Usage, error)
	GetUsageAsync(ctx workflow.Context, input *apigateway.GetUsageInput) *GetUsageFuture

	GetUsagePlan(ctx workflow.Context, input *apigateway.GetUsagePlanInput) (*apigateway.UsagePlan, error)
	GetUsagePlanAsync(ctx workflow.Context, input *apigateway.GetUsagePlanInput) *GetUsagePlanFuture

	GetUsagePlanKey(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) (*apigateway.UsagePlanKey, error)
	GetUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) *GetUsagePlanKeyFuture

	GetUsagePlanKeys(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlanKeysAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) *GetUsagePlanKeysFuture

	GetUsagePlans(ctx workflow.Context, input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error)
	GetUsagePlansAsync(ctx workflow.Context, input *apigateway.GetUsagePlansInput) *GetUsagePlansFuture

	GetVpcLink(ctx workflow.Context, input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	GetVpcLinkAsync(ctx workflow.Context, input *apigateway.GetVpcLinkInput) *GetVpcLinkFuture

	GetVpcLinks(ctx workflow.Context, input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error)
	GetVpcLinksAsync(ctx workflow.Context, input *apigateway.GetVpcLinksInput) *GetVpcLinksFuture

	ImportApiKeys(ctx workflow.Context, input *apigateway.ImportApiKeysInput) (*apigateway.ImportApiKeysOutput, error)
	ImportApiKeysAsync(ctx workflow.Context, input *apigateway.ImportApiKeysInput) *ImportApiKeysFuture

	ImportDocumentationParts(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) (*apigateway.ImportDocumentationPartsOutput, error)
	ImportDocumentationPartsAsync(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) *ImportDocumentationPartsFuture

	ImportRestApi(ctx workflow.Context, input *apigateway.ImportRestApiInput) (*apigateway.RestApi, error)
	ImportRestApiAsync(ctx workflow.Context, input *apigateway.ImportRestApiInput) *ImportRestApiFuture

	PutGatewayResponse(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	PutGatewayResponseAsync(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) *PutGatewayResponseFuture

	PutIntegration(ctx workflow.Context, input *apigateway.PutIntegrationInput) (*apigateway.Integration, error)
	PutIntegrationAsync(ctx workflow.Context, input *apigateway.PutIntegrationInput) *PutIntegrationFuture

	PutIntegrationResponse(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	PutIntegrationResponseAsync(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) *PutIntegrationResponseFuture

	PutMethod(ctx workflow.Context, input *apigateway.PutMethodInput) (*apigateway.Method, error)
	PutMethodAsync(ctx workflow.Context, input *apigateway.PutMethodInput) *PutMethodFuture

	PutMethodResponse(ctx workflow.Context, input *apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error)
	PutMethodResponseAsync(ctx workflow.Context, input *apigateway.PutMethodResponseInput) *PutMethodResponseFuture

	PutRestApi(ctx workflow.Context, input *apigateway.PutRestApiInput) (*apigateway.RestApi, error)
	PutRestApiAsync(ctx workflow.Context, input *apigateway.PutRestApiInput) *PutRestApiFuture

	TagResource(ctx workflow.Context, input *apigateway.TagResourceInput) (*apigateway.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *apigateway.TagResourceInput) *TagResourceFuture

	TestInvokeAuthorizer(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error)
	TestInvokeAuthorizerAsync(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) *TestInvokeAuthorizerFuture

	TestInvokeMethod(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error)
	TestInvokeMethodAsync(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) *TestInvokeMethodFuture

	UntagResource(ctx workflow.Context, input *apigateway.UntagResourceInput) (*apigateway.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *apigateway.UntagResourceInput) *UntagResourceFuture

	UpdateAccount(ctx workflow.Context, input *apigateway.UpdateAccountInput) (*apigateway.Account, error)
	UpdateAccountAsync(ctx workflow.Context, input *apigateway.UpdateAccountInput) *UpdateAccountFuture

	UpdateApiKey(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error)
	UpdateApiKeyAsync(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) *UpdateApiKeyFuture

	UpdateAuthorizer(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error)
	UpdateAuthorizerAsync(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) *UpdateAuthorizerFuture

	UpdateBasePathMapping(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error)
	UpdateBasePathMappingAsync(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) *UpdateBasePathMappingFuture

	UpdateClientCertificate(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error)
	UpdateClientCertificateAsync(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) *UpdateClientCertificateFuture

	UpdateDeployment(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error)
	UpdateDeploymentAsync(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) *UpdateDeploymentFuture

	UpdateDocumentationPart(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) (*apigateway.DocumentationPart, error)
	UpdateDocumentationPartAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) *UpdateDocumentationPartFuture

	UpdateDocumentationVersion(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	UpdateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) *UpdateDocumentationVersionFuture

	UpdateDomainName(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error)
	UpdateDomainNameAsync(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) *UpdateDomainNameFuture

	UpdateGatewayResponse(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	UpdateGatewayResponseAsync(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) *UpdateGatewayResponseFuture

	UpdateIntegration(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) (*apigateway.Integration, error)
	UpdateIntegrationAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) *UpdateIntegrationFuture

	UpdateIntegrationResponse(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) *UpdateIntegrationResponseFuture

	UpdateMethod(ctx workflow.Context, input *apigateway.UpdateMethodInput) (*apigateway.Method, error)
	UpdateMethodAsync(ctx workflow.Context, input *apigateway.UpdateMethodInput) *UpdateMethodFuture

	UpdateMethodResponse(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error)
	UpdateMethodResponseAsync(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) *UpdateMethodResponseFuture

	UpdateModel(ctx workflow.Context, input *apigateway.UpdateModelInput) (*apigateway.Model, error)
	UpdateModelAsync(ctx workflow.Context, input *apigateway.UpdateModelInput) *UpdateModelFuture

	UpdateRequestValidator(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	UpdateRequestValidatorAsync(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) *UpdateRequestValidatorFuture

	UpdateResource(ctx workflow.Context, input *apigateway.UpdateResourceInput) (*apigateway.Resource, error)
	UpdateResourceAsync(ctx workflow.Context, input *apigateway.UpdateResourceInput) *UpdateResourceFuture

	UpdateRestApi(ctx workflow.Context, input *apigateway.UpdateRestApiInput) (*apigateway.RestApi, error)
	UpdateRestApiAsync(ctx workflow.Context, input *apigateway.UpdateRestApiInput) *UpdateRestApiFuture

	UpdateStage(ctx workflow.Context, input *apigateway.UpdateStageInput) (*apigateway.Stage, error)
	UpdateStageAsync(ctx workflow.Context, input *apigateway.UpdateStageInput) *UpdateStageFuture

	UpdateUsage(ctx workflow.Context, input *apigateway.UpdateUsageInput) (*apigateway.Usage, error)
	UpdateUsageAsync(ctx workflow.Context, input *apigateway.UpdateUsageInput) *UpdateUsageFuture

	UpdateUsagePlan(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) (*apigateway.UsagePlan, error)
	UpdateUsagePlanAsync(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) *UpdateUsagePlanFuture

	UpdateVpcLink(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	UpdateVpcLinkAsync(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) *UpdateVpcLinkFuture
}

func NewClient() Client {
	return &stub{}
}
