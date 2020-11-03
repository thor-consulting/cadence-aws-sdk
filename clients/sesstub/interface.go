// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package sesstub

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CloneReceiptRuleSet(ctx workflow.Context, input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error)
	CloneReceiptRuleSetAsync(ctx workflow.Context, input *ses.CloneReceiptRuleSetInput) *SESCloneReceiptRuleSetFuture

	CreateConfigurationSet(ctx workflow.Context, input *ses.CreateConfigurationSetInput) (*ses.CreateConfigurationSetOutput, error)
	CreateConfigurationSetAsync(ctx workflow.Context, input *ses.CreateConfigurationSetInput) *SESCreateConfigurationSetFuture

	CreateConfigurationSetEventDestination(ctx workflow.Context, input *ses.CreateConfigurationSetEventDestinationInput) (*ses.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *ses.CreateConfigurationSetEventDestinationInput) *SESCreateConfigurationSetEventDestinationFuture

	CreateConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.CreateConfigurationSetTrackingOptionsInput) (*ses.CreateConfigurationSetTrackingOptionsOutput, error)
	CreateConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *ses.CreateConfigurationSetTrackingOptionsInput) *SESCreateConfigurationSetTrackingOptionsFuture

	CreateCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.CreateCustomVerificationEmailTemplateInput) (*ses.CreateCustomVerificationEmailTemplateOutput, error)
	CreateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.CreateCustomVerificationEmailTemplateInput) *SESCreateCustomVerificationEmailTemplateFuture

	CreateReceiptFilter(ctx workflow.Context, input *ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error)
	CreateReceiptFilterAsync(ctx workflow.Context, input *ses.CreateReceiptFilterInput) *SESCreateReceiptFilterFuture

	CreateReceiptRule(ctx workflow.Context, input *ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error)
	CreateReceiptRuleAsync(ctx workflow.Context, input *ses.CreateReceiptRuleInput) *SESCreateReceiptRuleFuture

	CreateReceiptRuleSet(ctx workflow.Context, input *ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error)
	CreateReceiptRuleSetAsync(ctx workflow.Context, input *ses.CreateReceiptRuleSetInput) *SESCreateReceiptRuleSetFuture

	CreateTemplate(ctx workflow.Context, input *ses.CreateTemplateInput) (*ses.CreateTemplateOutput, error)
	CreateTemplateAsync(ctx workflow.Context, input *ses.CreateTemplateInput) *SESCreateTemplateFuture

	DeleteConfigurationSet(ctx workflow.Context, input *ses.DeleteConfigurationSetInput) (*ses.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetAsync(ctx workflow.Context, input *ses.DeleteConfigurationSetInput) *SESDeleteConfigurationSetFuture

	DeleteConfigurationSetEventDestination(ctx workflow.Context, input *ses.DeleteConfigurationSetEventDestinationInput) (*ses.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *ses.DeleteConfigurationSetEventDestinationInput) *SESDeleteConfigurationSetEventDestinationFuture

	DeleteConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.DeleteConfigurationSetTrackingOptionsInput) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error)
	DeleteConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *ses.DeleteConfigurationSetTrackingOptionsInput) *SESDeleteConfigurationSetTrackingOptionsFuture

	DeleteCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.DeleteCustomVerificationEmailTemplateInput) (*ses.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.DeleteCustomVerificationEmailTemplateInput) *SESDeleteCustomVerificationEmailTemplateFuture

	DeleteIdentity(ctx workflow.Context, input *ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)
	DeleteIdentityAsync(ctx workflow.Context, input *ses.DeleteIdentityInput) *SESDeleteIdentityFuture

	DeleteIdentityPolicy(ctx workflow.Context, input *ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error)
	DeleteIdentityPolicyAsync(ctx workflow.Context, input *ses.DeleteIdentityPolicyInput) *SESDeleteIdentityPolicyFuture

	DeleteReceiptFilter(ctx workflow.Context, input *ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error)
	DeleteReceiptFilterAsync(ctx workflow.Context, input *ses.DeleteReceiptFilterInput) *SESDeleteReceiptFilterFuture

	DeleteReceiptRule(ctx workflow.Context, input *ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error)
	DeleteReceiptRuleAsync(ctx workflow.Context, input *ses.DeleteReceiptRuleInput) *SESDeleteReceiptRuleFuture

	DeleteReceiptRuleSet(ctx workflow.Context, input *ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error)
	DeleteReceiptRuleSetAsync(ctx workflow.Context, input *ses.DeleteReceiptRuleSetInput) *SESDeleteReceiptRuleSetFuture

	DeleteTemplate(ctx workflow.Context, input *ses.DeleteTemplateInput) (*ses.DeleteTemplateOutput, error)
	DeleteTemplateAsync(ctx workflow.Context, input *ses.DeleteTemplateInput) *SESDeleteTemplateFuture

	DeleteVerifiedEmailAddress(ctx workflow.Context, input *ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)
	DeleteVerifiedEmailAddressAsync(ctx workflow.Context, input *ses.DeleteVerifiedEmailAddressInput) *SESDeleteVerifiedEmailAddressFuture

	DescribeActiveReceiptRuleSet(ctx workflow.Context, input *ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error)
	DescribeActiveReceiptRuleSetAsync(ctx workflow.Context, input *ses.DescribeActiveReceiptRuleSetInput) *SESDescribeActiveReceiptRuleSetFuture

	DescribeConfigurationSet(ctx workflow.Context, input *ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error)
	DescribeConfigurationSetAsync(ctx workflow.Context, input *ses.DescribeConfigurationSetInput) *SESDescribeConfigurationSetFuture

	DescribeReceiptRule(ctx workflow.Context, input *ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error)
	DescribeReceiptRuleAsync(ctx workflow.Context, input *ses.DescribeReceiptRuleInput) *SESDescribeReceiptRuleFuture

	DescribeReceiptRuleSet(ctx workflow.Context, input *ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error)
	DescribeReceiptRuleSetAsync(ctx workflow.Context, input *ses.DescribeReceiptRuleSetInput) *SESDescribeReceiptRuleSetFuture

	GetAccountSendingEnabled(ctx workflow.Context, input *ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error)
	GetAccountSendingEnabledAsync(ctx workflow.Context, input *ses.GetAccountSendingEnabledInput) *SESGetAccountSendingEnabledFuture

	GetCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error)
	GetCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.GetCustomVerificationEmailTemplateInput) *SESGetCustomVerificationEmailTemplateFuture

	GetIdentityDkimAttributes(ctx workflow.Context, input *ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error)
	GetIdentityDkimAttributesAsync(ctx workflow.Context, input *ses.GetIdentityDkimAttributesInput) *SESGetIdentityDkimAttributesFuture

	GetIdentityMailFromDomainAttributes(ctx workflow.Context, input *ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error)
	GetIdentityMailFromDomainAttributesAsync(ctx workflow.Context, input *ses.GetIdentityMailFromDomainAttributesInput) *SESGetIdentityMailFromDomainAttributesFuture

	GetIdentityNotificationAttributes(ctx workflow.Context, input *ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)
	GetIdentityNotificationAttributesAsync(ctx workflow.Context, input *ses.GetIdentityNotificationAttributesInput) *SESGetIdentityNotificationAttributesFuture

	GetIdentityPolicies(ctx workflow.Context, input *ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error)
	GetIdentityPoliciesAsync(ctx workflow.Context, input *ses.GetIdentityPoliciesInput) *SESGetIdentityPoliciesFuture

	GetIdentityVerificationAttributes(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)
	GetIdentityVerificationAttributesAsync(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) *SESGetIdentityVerificationAttributesFuture

	GetSendQuota(ctx workflow.Context, input *ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)
	GetSendQuotaAsync(ctx workflow.Context, input *ses.GetSendQuotaInput) *SESGetSendQuotaFuture

	GetSendStatistics(ctx workflow.Context, input *ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)
	GetSendStatisticsAsync(ctx workflow.Context, input *ses.GetSendStatisticsInput) *SESGetSendStatisticsFuture

	GetTemplate(ctx workflow.Context, input *ses.GetTemplateInput) (*ses.GetTemplateOutput, error)
	GetTemplateAsync(ctx workflow.Context, input *ses.GetTemplateInput) *SESGetTemplateFuture

	ListConfigurationSets(ctx workflow.Context, input *ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error)
	ListConfigurationSetsAsync(ctx workflow.Context, input *ses.ListConfigurationSetsInput) *SESListConfigurationSetsFuture

	ListCustomVerificationEmailTemplates(ctx workflow.Context, input *ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error)
	ListCustomVerificationEmailTemplatesAsync(ctx workflow.Context, input *ses.ListCustomVerificationEmailTemplatesInput) *SESListCustomVerificationEmailTemplatesFuture

	ListIdentities(ctx workflow.Context, input *ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)
	ListIdentitiesAsync(ctx workflow.Context, input *ses.ListIdentitiesInput) *SESListIdentitiesFuture

	ListIdentityPolicies(ctx workflow.Context, input *ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error)
	ListIdentityPoliciesAsync(ctx workflow.Context, input *ses.ListIdentityPoliciesInput) *SESListIdentityPoliciesFuture

	ListReceiptFilters(ctx workflow.Context, input *ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error)
	ListReceiptFiltersAsync(ctx workflow.Context, input *ses.ListReceiptFiltersInput) *SESListReceiptFiltersFuture

	ListReceiptRuleSets(ctx workflow.Context, input *ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error)
	ListReceiptRuleSetsAsync(ctx workflow.Context, input *ses.ListReceiptRuleSetsInput) *SESListReceiptRuleSetsFuture

	ListTemplates(ctx workflow.Context, input *ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error)
	ListTemplatesAsync(ctx workflow.Context, input *ses.ListTemplatesInput) *SESListTemplatesFuture

	ListVerifiedEmailAddresses(ctx workflow.Context, input *ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)
	ListVerifiedEmailAddressesAsync(ctx workflow.Context, input *ses.ListVerifiedEmailAddressesInput) *SESListVerifiedEmailAddressesFuture

	PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *ses.PutConfigurationSetDeliveryOptionsInput) (*ses.PutConfigurationSetDeliveryOptionsOutput, error)
	PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *ses.PutConfigurationSetDeliveryOptionsInput) *SESPutConfigurationSetDeliveryOptionsFuture

	PutIdentityPolicy(ctx workflow.Context, input *ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error)
	PutIdentityPolicyAsync(ctx workflow.Context, input *ses.PutIdentityPolicyInput) *SESPutIdentityPolicyFuture

	ReorderReceiptRuleSet(ctx workflow.Context, input *ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error)
	ReorderReceiptRuleSetAsync(ctx workflow.Context, input *ses.ReorderReceiptRuleSetInput) *SESReorderReceiptRuleSetFuture

	SendBounce(ctx workflow.Context, input *ses.SendBounceInput) (*ses.SendBounceOutput, error)
	SendBounceAsync(ctx workflow.Context, input *ses.SendBounceInput) *SESSendBounceFuture

	SendBulkTemplatedEmail(ctx workflow.Context, input *ses.SendBulkTemplatedEmailInput) (*ses.SendBulkTemplatedEmailOutput, error)
	SendBulkTemplatedEmailAsync(ctx workflow.Context, input *ses.SendBulkTemplatedEmailInput) *SESSendBulkTemplatedEmailFuture

	SendCustomVerificationEmail(ctx workflow.Context, input *ses.SendCustomVerificationEmailInput) (*ses.SendCustomVerificationEmailOutput, error)
	SendCustomVerificationEmailAsync(ctx workflow.Context, input *ses.SendCustomVerificationEmailInput) *SESSendCustomVerificationEmailFuture

	SendEmail(ctx workflow.Context, input *ses.SendEmailInput) (*ses.SendEmailOutput, error)
	SendEmailAsync(ctx workflow.Context, input *ses.SendEmailInput) *SESSendEmailFuture

	SendRawEmail(ctx workflow.Context, input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)
	SendRawEmailAsync(ctx workflow.Context, input *ses.SendRawEmailInput) *SESSendRawEmailFuture

	SendTemplatedEmail(ctx workflow.Context, input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error)
	SendTemplatedEmailAsync(ctx workflow.Context, input *ses.SendTemplatedEmailInput) *SESSendTemplatedEmailFuture

	SetActiveReceiptRuleSet(ctx workflow.Context, input *ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error)
	SetActiveReceiptRuleSetAsync(ctx workflow.Context, input *ses.SetActiveReceiptRuleSetInput) *SESSetActiveReceiptRuleSetFuture

	SetIdentityDkimEnabled(ctx workflow.Context, input *ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error)
	SetIdentityDkimEnabledAsync(ctx workflow.Context, input *ses.SetIdentityDkimEnabledInput) *SESSetIdentityDkimEnabledFuture

	SetIdentityFeedbackForwardingEnabled(ctx workflow.Context, input *ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)
	SetIdentityFeedbackForwardingEnabledAsync(ctx workflow.Context, input *ses.SetIdentityFeedbackForwardingEnabledInput) *SESSetIdentityFeedbackForwardingEnabledFuture

	SetIdentityHeadersInNotificationsEnabled(ctx workflow.Context, input *ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error)
	SetIdentityHeadersInNotificationsEnabledAsync(ctx workflow.Context, input *ses.SetIdentityHeadersInNotificationsEnabledInput) *SESSetIdentityHeadersInNotificationsEnabledFuture

	SetIdentityMailFromDomain(ctx workflow.Context, input *ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error)
	SetIdentityMailFromDomainAsync(ctx workflow.Context, input *ses.SetIdentityMailFromDomainInput) *SESSetIdentityMailFromDomainFuture

	SetIdentityNotificationTopic(ctx workflow.Context, input *ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)
	SetIdentityNotificationTopicAsync(ctx workflow.Context, input *ses.SetIdentityNotificationTopicInput) *SESSetIdentityNotificationTopicFuture

	SetReceiptRulePosition(ctx workflow.Context, input *ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error)
	SetReceiptRulePositionAsync(ctx workflow.Context, input *ses.SetReceiptRulePositionInput) *SESSetReceiptRulePositionFuture

	TestRenderTemplate(ctx workflow.Context, input *ses.TestRenderTemplateInput) (*ses.TestRenderTemplateOutput, error)
	TestRenderTemplateAsync(ctx workflow.Context, input *ses.TestRenderTemplateInput) *SESTestRenderTemplateFuture

	UpdateAccountSendingEnabled(ctx workflow.Context, input *ses.UpdateAccountSendingEnabledInput) (*ses.UpdateAccountSendingEnabledOutput, error)
	UpdateAccountSendingEnabledAsync(ctx workflow.Context, input *ses.UpdateAccountSendingEnabledInput) *SESUpdateAccountSendingEnabledFuture

	UpdateConfigurationSetEventDestination(ctx workflow.Context, input *ses.UpdateConfigurationSetEventDestinationInput) (*ses.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetEventDestinationInput) *SESUpdateConfigurationSetEventDestinationFuture

	UpdateConfigurationSetReputationMetricsEnabled(ctx workflow.Context, input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error)
	UpdateConfigurationSetReputationMetricsEnabledAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) *SESUpdateConfigurationSetReputationMetricsEnabledFuture

	UpdateConfigurationSetSendingEnabled(ctx workflow.Context, input *ses.UpdateConfigurationSetSendingEnabledInput) (*ses.UpdateConfigurationSetSendingEnabledOutput, error)
	UpdateConfigurationSetSendingEnabledAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetSendingEnabledInput) *SESUpdateConfigurationSetSendingEnabledFuture

	UpdateConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.UpdateConfigurationSetTrackingOptionsInput) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error)
	UpdateConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetTrackingOptionsInput) *SESUpdateConfigurationSetTrackingOptionsFuture

	UpdateCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.UpdateCustomVerificationEmailTemplateInput) (*ses.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.UpdateCustomVerificationEmailTemplateInput) *SESUpdateCustomVerificationEmailTemplateFuture

	UpdateReceiptRule(ctx workflow.Context, input *ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error)
	UpdateReceiptRuleAsync(ctx workflow.Context, input *ses.UpdateReceiptRuleInput) *SESUpdateReceiptRuleFuture

	UpdateTemplate(ctx workflow.Context, input *ses.UpdateTemplateInput) (*ses.UpdateTemplateOutput, error)
	UpdateTemplateAsync(ctx workflow.Context, input *ses.UpdateTemplateInput) *SESUpdateTemplateFuture

	VerifyDomainDkim(ctx workflow.Context, input *ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error)
	VerifyDomainDkimAsync(ctx workflow.Context, input *ses.VerifyDomainDkimInput) *SESVerifyDomainDkimFuture

	VerifyDomainIdentity(ctx workflow.Context, input *ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)
	VerifyDomainIdentityAsync(ctx workflow.Context, input *ses.VerifyDomainIdentityInput) *SESVerifyDomainIdentityFuture

	VerifyEmailAddress(ctx workflow.Context, input *ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)
	VerifyEmailAddressAsync(ctx workflow.Context, input *ses.VerifyEmailAddressInput) *SESVerifyEmailAddressFuture

	VerifyEmailIdentity(ctx workflow.Context, input *ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)
	VerifyEmailIdentityAsync(ctx workflow.Context, input *ses.VerifyEmailIdentityInput) *SESVerifyEmailIdentityFuture

	WaitUntilIdentityExists(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) error
	WaitUntilIdentityExistsAsync(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
