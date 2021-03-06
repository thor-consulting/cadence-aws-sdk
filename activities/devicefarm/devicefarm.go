// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"

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
	client devicefarmiface.DeviceFarmAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := devicefarm.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (devicefarmiface.DeviceFarmAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return devicefarm.New(sess), nil
}

func (a *Activities) CreateDevicePool(ctx context.Context, input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDevicePoolWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateInstanceProfile(ctx context.Context, input *devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateInstanceProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateNetworkProfile(ctx context.Context, input *devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateProject(ctx context.Context, input *devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRemoteAccessSession(ctx context.Context, input *devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRemoteAccessSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTestGridProject(ctx context.Context, input *devicefarm.CreateTestGridProjectInput) (*devicefarm.CreateTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTestGridProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTestGridUrl(ctx context.Context, input *devicefarm.CreateTestGridUrlInput) (*devicefarm.CreateTestGridUrlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTestGridUrlWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUpload(ctx context.Context, input *devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVPCEConfiguration(ctx context.Context, input *devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateVPCEConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDevicePool(ctx context.Context, input *devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDevicePoolWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteInstanceProfile(ctx context.Context, input *devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteInstanceProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNetworkProfile(ctx context.Context, input *devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProject(ctx context.Context, input *devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRemoteAccessSession(ctx context.Context, input *devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRemoteAccessSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRun(ctx context.Context, input *devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTestGridProject(ctx context.Context, input *devicefarm.DeleteTestGridProjectInput) (*devicefarm.DeleteTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTestGridProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUpload(ctx context.Context, input *devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVPCEConfiguration(ctx context.Context, input *devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVPCEConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAccountSettings(ctx context.Context, input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAccountSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevice(ctx context.Context, input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDeviceInstance(ctx context.Context, input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDeviceInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevicePool(ctx context.Context, input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDevicePoolWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevicePoolCompatibility(ctx context.Context, input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDevicePoolCompatibilityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetInstanceProfile(ctx context.Context, input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetInstanceProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetJob(ctx context.Context, input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetNetworkProfile(ctx context.Context, input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetOfferingStatus(ctx context.Context, input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetOfferingStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetProject(ctx context.Context, input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRemoteAccessSession(ctx context.Context, input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRemoteAccessSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRun(ctx context.Context, input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSuite(ctx context.Context, input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSuiteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTest(ctx context.Context, input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTestWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTestGridProject(ctx context.Context, input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTestGridProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTestGridSession(ctx context.Context, input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTestGridSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetUpload(ctx context.Context, input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetVPCEConfiguration(ctx context.Context, input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetVPCEConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) InstallToRemoteAccessSession(ctx context.Context, input *devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.InstallToRemoteAccessSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListArtifacts(ctx context.Context, input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListArtifactsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDeviceInstances(ctx context.Context, input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDeviceInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDevicePools(ctx context.Context, input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDevicePoolsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDevices(ctx context.Context, input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDevicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInstanceProfiles(ctx context.Context, input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInstanceProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobs(ctx context.Context, input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListNetworkProfiles(ctx context.Context, input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListNetworkProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListOfferingPromotions(ctx context.Context, input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListOfferingPromotionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListOfferingTransactions(ctx context.Context, input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListOfferingTransactionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListOfferings(ctx context.Context, input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListOfferingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProjects(ctx context.Context, input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRemoteAccessSessions(ctx context.Context, input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRemoteAccessSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRuns(ctx context.Context, input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRunsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSamples(ctx context.Context, input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSamplesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSuites(ctx context.Context, input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSuitesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTestGridProjects(ctx context.Context, input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTestGridProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTestGridSessionActions(ctx context.Context, input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTestGridSessionActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTestGridSessionArtifacts(ctx context.Context, input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTestGridSessionArtifactsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTestGridSessions(ctx context.Context, input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTestGridSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTests(ctx context.Context, input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTestsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUniqueProblems(ctx context.Context, input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUniqueProblemsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUploads(ctx context.Context, input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUploadsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVPCEConfigurations(ctx context.Context, input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVPCEConfigurationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PurchaseOffering(ctx context.Context, input *devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PurchaseOfferingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RenewOffering(ctx context.Context, input *devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RenewOfferingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ScheduleRun(ctx context.Context, input *devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ScheduleRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopJob(ctx context.Context, input *devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopRemoteAccessSession(ctx context.Context, input *devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopRemoteAccessSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopRun(ctx context.Context, input *devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *devicefarm.TagResourceInput) (*devicefarm.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *devicefarm.UntagResourceInput) (*devicefarm.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDeviceInstance(ctx context.Context, input *devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDeviceInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDevicePool(ctx context.Context, input *devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDevicePoolWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateInstanceProfile(ctx context.Context, input *devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateInstanceProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateNetworkProfile(ctx context.Context, input *devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateProject(ctx context.Context, input *devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateTestGridProject(ctx context.Context, input *devicefarm.UpdateTestGridProjectInput) (*devicefarm.UpdateTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateTestGridProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUpload(ctx context.Context, input *devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVPCEConfiguration(ctx context.Context, input *devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateVPCEConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
