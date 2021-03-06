// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"

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
	client codebuildiface.CodeBuildAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := codebuild.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (codebuildiface.CodeBuildAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return codebuild.New(sess), nil
}

func (a *Activities) BatchDeleteBuilds(ctx context.Context, input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDeleteBuildsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetBuildBatches(ctx context.Context, input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetBuildBatchesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetBuilds(ctx context.Context, input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetBuildsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetProjects(ctx context.Context, input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetReportGroups(ctx context.Context, input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetReportGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetReports(ctx context.Context, input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetReportsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateProject(ctx context.Context, input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateReportGroup(ctx context.Context, input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateReportGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateWebhook(ctx context.Context, input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateWebhookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBuildBatch(ctx context.Context, input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBuildBatchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProject(ctx context.Context, input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReport(ctx context.Context, input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReportWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReportGroup(ctx context.Context, input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReportGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteResourcePolicy(ctx context.Context, input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteResourcePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSourceCredentials(ctx context.Context, input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSourceCredentialsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteWebhook(ctx context.Context, input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteWebhookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCodeCoverages(ctx context.Context, input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCodeCoveragesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTestCases(ctx context.Context, input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTestCasesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetReportGroupTrend(ctx context.Context, input *codebuild.GetReportGroupTrendInput) (*codebuild.GetReportGroupTrendOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetReportGroupTrendWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetResourcePolicy(ctx context.Context, input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetResourcePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportSourceCredentials(ctx context.Context, input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportSourceCredentialsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) InvalidateProjectCache(ctx context.Context, input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.InvalidateProjectCacheWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBuildBatches(ctx context.Context, input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBuildBatchesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBuildBatchesForProject(ctx context.Context, input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBuildBatchesForProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBuilds(ctx context.Context, input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBuildsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBuildsForProject(ctx context.Context, input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBuildsForProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListCuratedEnvironmentImages(ctx context.Context, input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListCuratedEnvironmentImagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProjects(ctx context.Context, input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListReportGroups(ctx context.Context, input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListReportGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListReports(ctx context.Context, input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListReportsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListReportsForReportGroup(ctx context.Context, input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListReportsForReportGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSharedProjects(ctx context.Context, input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSharedProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSharedReportGroups(ctx context.Context, input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSharedReportGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSourceCredentials(ctx context.Context, input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSourceCredentialsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutResourcePolicy(ctx context.Context, input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutResourcePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RetryBuild(ctx context.Context, input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RetryBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RetryBuildBatch(ctx context.Context, input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RetryBuildBatchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartBuild(ctx context.Context, input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartBuildBatch(ctx context.Context, input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartBuildBatchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopBuild(ctx context.Context, input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopBuildBatch(ctx context.Context, input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopBuildBatchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateProject(ctx context.Context, input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateReportGroup(ctx context.Context, input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateReportGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateWebhook(ctx context.Context, input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateWebhookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
