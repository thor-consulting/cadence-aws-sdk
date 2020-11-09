// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package alexaforbusinessstub

import (
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ApproveSkill(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error)
	ApproveSkillAsync(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) *ApproveSkillFuture

	AssociateContactWithAddressBook(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error)
	AssociateContactWithAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) *AssociateContactWithAddressBookFuture

	AssociateDeviceWithNetworkProfile(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error)
	AssociateDeviceWithNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) *AssociateDeviceWithNetworkProfileFuture

	AssociateDeviceWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error)
	AssociateDeviceWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) *AssociateDeviceWithRoomFuture

	AssociateSkillGroupWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error)
	AssociateSkillGroupWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) *AssociateSkillGroupWithRoomFuture

	AssociateSkillWithSkillGroup(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error)
	AssociateSkillWithSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) *AssociateSkillWithSkillGroupFuture

	AssociateSkillWithUsers(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error)
	AssociateSkillWithUsersAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) *AssociateSkillWithUsersFuture

	CreateAddressBook(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error)
	CreateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) *CreateAddressBookFuture

	CreateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error)
	CreateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) *CreateBusinessReportScheduleFuture

	CreateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error)
	CreateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) *CreateConferenceProviderFuture

	CreateContact(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error)
	CreateContactAsync(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) *CreateContactFuture

	CreateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error)
	CreateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) *CreateGatewayGroupFuture

	CreateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error)
	CreateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) *CreateNetworkProfileFuture

	CreateProfile(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error)
	CreateProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) *CreateProfileFuture

	CreateRoom(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error)
	CreateRoomAsync(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) *CreateRoomFuture

	CreateSkillGroup(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error)
	CreateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) *CreateSkillGroupFuture

	CreateUser(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) *CreateUserFuture

	DeleteAddressBook(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error)
	DeleteAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) *DeleteAddressBookFuture

	DeleteBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error)
	DeleteBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) *DeleteBusinessReportScheduleFuture

	DeleteConferenceProvider(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error)
	DeleteConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) *DeleteConferenceProviderFuture

	DeleteContact(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error)
	DeleteContactAsync(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) *DeleteContactFuture

	DeleteDevice(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error)
	DeleteDeviceAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) *DeleteDeviceFuture

	DeleteDeviceUsageData(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error)
	DeleteDeviceUsageDataAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) *DeleteDeviceUsageDataFuture

	DeleteGatewayGroup(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error)
	DeleteGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) *DeleteGatewayGroupFuture

	DeleteNetworkProfile(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error)
	DeleteNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) *DeleteNetworkProfileFuture

	DeleteProfile(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error)
	DeleteProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) *DeleteProfileFuture

	DeleteRoom(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error)
	DeleteRoomAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) *DeleteRoomFuture

	DeleteRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error)
	DeleteRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) *DeleteRoomSkillParameterFuture

	DeleteSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error)
	DeleteSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) *DeleteSkillAuthorizationFuture

	DeleteSkillGroup(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error)
	DeleteSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) *DeleteSkillGroupFuture

	DeleteUser(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) *DeleteUserFuture

	DisassociateContactFromAddressBook(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error)
	DisassociateContactFromAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) *DisassociateContactFromAddressBookFuture

	DisassociateDeviceFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error)
	DisassociateDeviceFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) *DisassociateDeviceFromRoomFuture

	DisassociateSkillFromSkillGroup(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error)
	DisassociateSkillFromSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) *DisassociateSkillFromSkillGroupFuture

	DisassociateSkillFromUsers(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error)
	DisassociateSkillFromUsersAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) *DisassociateSkillFromUsersFuture

	DisassociateSkillGroupFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error)
	DisassociateSkillGroupFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) *DisassociateSkillGroupFromRoomFuture

	ForgetSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error)
	ForgetSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) *ForgetSmartHomeAppliancesFuture

	GetAddressBook(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error)
	GetAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) *GetAddressBookFuture

	GetConferencePreference(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error)
	GetConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) *GetConferencePreferenceFuture

	GetConferenceProvider(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error)
	GetConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) *GetConferenceProviderFuture

	GetContact(ctx workflow.Context, input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error)
	GetContactAsync(ctx workflow.Context, input *alexaforbusiness.GetContactInput) *GetContactFuture

	GetDevice(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error)
	GetDeviceAsync(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) *GetDeviceFuture

	GetGateway(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error)
	GetGatewayAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) *GetGatewayFuture

	GetGatewayGroup(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error)
	GetGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) *GetGatewayGroupFuture

	GetInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error)
	GetInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) *GetInvitationConfigurationFuture

	GetNetworkProfile(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error)
	GetNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) *GetNetworkProfileFuture

	GetProfile(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error)
	GetProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) *GetProfileFuture

	GetRoom(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error)
	GetRoomAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) *GetRoomFuture

	GetRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error)
	GetRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) *GetRoomSkillParameterFuture

	GetSkillGroup(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error)
	GetSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) *GetSkillGroupFuture

	ListBusinessReportSchedules(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error)
	ListBusinessReportSchedulesAsync(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) *ListBusinessReportSchedulesFuture

	ListConferenceProviders(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error)
	ListConferenceProvidersAsync(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) *ListConferenceProvidersFuture

	ListDeviceEvents(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error)
	ListDeviceEventsAsync(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) *ListDeviceEventsFuture

	ListGatewayGroups(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error)
	ListGatewayGroupsAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) *ListGatewayGroupsFuture

	ListGateways(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error)
	ListGatewaysAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) *ListGatewaysFuture

	ListSkills(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error)
	ListSkillsAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) *ListSkillsFuture

	ListSkillsStoreCategories(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error)
	ListSkillsStoreCategoriesAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) *ListSkillsStoreCategoriesFuture

	ListSkillsStoreSkillsByCategory(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error)
	ListSkillsStoreSkillsByCategoryAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) *ListSkillsStoreSkillsByCategoryFuture

	ListSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error)
	ListSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) *ListSmartHomeAppliancesFuture

	ListTags(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) *ListTagsFuture

	PutConferencePreference(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error)
	PutConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) *PutConferencePreferenceFuture

	PutInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error)
	PutInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) *PutInvitationConfigurationFuture

	PutRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error)
	PutRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) *PutRoomSkillParameterFuture

	PutSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error)
	PutSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) *PutSkillAuthorizationFuture

	RegisterAVSDevice(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error)
	RegisterAVSDeviceAsync(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) *RegisterAVSDeviceFuture

	RejectSkill(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error)
	RejectSkillAsync(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) *RejectSkillFuture

	ResolveRoom(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error)
	ResolveRoomAsync(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) *ResolveRoomFuture

	RevokeInvitation(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error)
	RevokeInvitationAsync(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) *RevokeInvitationFuture

	SearchAddressBooks(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error)
	SearchAddressBooksAsync(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) *SearchAddressBooksFuture

	SearchContacts(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error)
	SearchContactsAsync(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) *SearchContactsFuture

	SearchDevices(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error)
	SearchDevicesAsync(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) *SearchDevicesFuture

	SearchNetworkProfiles(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error)
	SearchNetworkProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) *SearchNetworkProfilesFuture

	SearchProfiles(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error)
	SearchProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) *SearchProfilesFuture

	SearchRooms(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error)
	SearchRoomsAsync(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) *SearchRoomsFuture

	SearchSkillGroups(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error)
	SearchSkillGroupsAsync(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) *SearchSkillGroupsFuture

	SearchUsers(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error)
	SearchUsersAsync(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) *SearchUsersFuture

	SendAnnouncement(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error)
	SendAnnouncementAsync(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) *SendAnnouncementFuture

	SendInvitation(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error)
	SendInvitationAsync(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) *SendInvitationFuture

	StartDeviceSync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error)
	StartDeviceSyncAsync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) *StartDeviceSyncFuture

	StartSmartHomeApplianceDiscovery(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error)
	StartSmartHomeApplianceDiscoveryAsync(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) *StartSmartHomeApplianceDiscoveryFuture

	TagResource(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) *UntagResourceFuture

	UpdateAddressBook(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error)
	UpdateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) *UpdateAddressBookFuture

	UpdateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error)
	UpdateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) *UpdateBusinessReportScheduleFuture

	UpdateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error)
	UpdateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) *UpdateConferenceProviderFuture

	UpdateContact(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error)
	UpdateContactAsync(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) *UpdateContactFuture

	UpdateDevice(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error)
	UpdateDeviceAsync(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) *UpdateDeviceFuture

	UpdateGateway(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error)
	UpdateGatewayAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) *UpdateGatewayFuture

	UpdateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error)
	UpdateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) *UpdateGatewayGroupFuture

	UpdateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error)
	UpdateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) *UpdateNetworkProfileFuture

	UpdateProfile(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error)
	UpdateProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) *UpdateProfileFuture

	UpdateRoom(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error)
	UpdateRoomAsync(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) *UpdateRoomFuture

	UpdateSkillGroup(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error)
	UpdateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) *UpdateSkillGroupFuture
}

func NewClient() Client {
	return &stub{}
}
