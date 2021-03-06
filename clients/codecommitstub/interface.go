// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codecommitstub

import (
	"github.com/aws/aws-sdk-go/service/codecommit"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateApprovalRuleTemplateWithRepository(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error)
	AssociateApprovalRuleTemplateWithRepositoryAsync(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) *AssociateApprovalRuleTemplateWithRepositoryFuture

	BatchAssociateApprovalRuleTemplateWithRepositories(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error)
	BatchAssociateApprovalRuleTemplateWithRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) *BatchAssociateApprovalRuleTemplateWithRepositoriesFuture

	BatchDescribeMergeConflicts(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error)
	BatchDescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) *BatchDescribeMergeConflictsFuture

	BatchDisassociateApprovalRuleTemplateFromRepositories(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error)
	BatchDisassociateApprovalRuleTemplateFromRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) *BatchDisassociateApprovalRuleTemplateFromRepositoriesFuture

	BatchGetCommits(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error)
	BatchGetCommitsAsync(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) *BatchGetCommitsFuture

	BatchGetRepositories(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error)
	BatchGetRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) *BatchGetRepositoriesFuture

	CreateApprovalRuleTemplate(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) (*codecommit.CreateApprovalRuleTemplateOutput, error)
	CreateApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) *CreateApprovalRuleTemplateFuture

	CreateBranch(ctx workflow.Context, input *codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error)
	CreateBranchAsync(ctx workflow.Context, input *codecommit.CreateBranchInput) *CreateBranchFuture

	CreateCommit(ctx workflow.Context, input *codecommit.CreateCommitInput) (*codecommit.CreateCommitOutput, error)
	CreateCommitAsync(ctx workflow.Context, input *codecommit.CreateCommitInput) *CreateCommitFuture

	CreatePullRequestApprovalRule(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) (*codecommit.CreatePullRequestApprovalRuleOutput, error)
	CreatePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) *CreatePullRequestApprovalRuleFuture

	CreateRepository(ctx workflow.Context, input *codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error)
	CreateRepositoryAsync(ctx workflow.Context, input *codecommit.CreateRepositoryInput) *CreateRepositoryFuture

	CreateUnreferencedMergeCommit(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) (*codecommit.CreateUnreferencedMergeCommitOutput, error)
	CreateUnreferencedMergeCommitAsync(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) *CreateUnreferencedMergeCommitFuture

	DeleteApprovalRuleTemplate(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) (*codecommit.DeleteApprovalRuleTemplateOutput, error)
	DeleteApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) *DeleteApprovalRuleTemplateFuture

	DeleteBranch(ctx workflow.Context, input *codecommit.DeleteBranchInput) (*codecommit.DeleteBranchOutput, error)
	DeleteBranchAsync(ctx workflow.Context, input *codecommit.DeleteBranchInput) *DeleteBranchFuture

	DeleteCommentContent(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) (*codecommit.DeleteCommentContentOutput, error)
	DeleteCommentContentAsync(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) *DeleteCommentContentFuture

	DeleteFile(ctx workflow.Context, input *codecommit.DeleteFileInput) (*codecommit.DeleteFileOutput, error)
	DeleteFileAsync(ctx workflow.Context, input *codecommit.DeleteFileInput) *DeleteFileFuture

	DeletePullRequestApprovalRule(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) (*codecommit.DeletePullRequestApprovalRuleOutput, error)
	DeletePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) *DeletePullRequestApprovalRuleFuture

	DeleteRepository(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error)
	DeleteRepositoryAsync(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) *DeleteRepositoryFuture

	DescribeMergeConflicts(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error)
	DescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) *DescribeMergeConflictsFuture

	DescribePullRequestEvents(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error)
	DescribePullRequestEventsAsync(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) *DescribePullRequestEventsFuture

	DisassociateApprovalRuleTemplateFromRepository(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error)
	DisassociateApprovalRuleTemplateFromRepositoryAsync(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) *DisassociateApprovalRuleTemplateFromRepositoryFuture

	EvaluatePullRequestApprovalRules(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error)
	EvaluatePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) *EvaluatePullRequestApprovalRulesFuture

	GetApprovalRuleTemplate(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error)
	GetApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) *GetApprovalRuleTemplateFuture

	GetBlob(ctx workflow.Context, input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error)
	GetBlobAsync(ctx workflow.Context, input *codecommit.GetBlobInput) *GetBlobFuture

	GetBranch(ctx workflow.Context, input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error)
	GetBranchAsync(ctx workflow.Context, input *codecommit.GetBranchInput) *GetBranchFuture

	GetComment(ctx workflow.Context, input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error)
	GetCommentAsync(ctx workflow.Context, input *codecommit.GetCommentInput) *GetCommentFuture

	GetCommentReactions(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error)
	GetCommentReactionsAsync(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) *GetCommentReactionsFuture

	GetCommentsForComparedCommit(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error)
	GetCommentsForComparedCommitAsync(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) *GetCommentsForComparedCommitFuture

	GetCommit(ctx workflow.Context, input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error)
	GetCommitAsync(ctx workflow.Context, input *codecommit.GetCommitInput) *GetCommitFuture

	GetDifferences(ctx workflow.Context, input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error)
	GetDifferencesAsync(ctx workflow.Context, input *codecommit.GetDifferencesInput) *GetDifferencesFuture

	GetFile(ctx workflow.Context, input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error)
	GetFileAsync(ctx workflow.Context, input *codecommit.GetFileInput) *GetFileFuture

	GetFolder(ctx workflow.Context, input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error)
	GetFolderAsync(ctx workflow.Context, input *codecommit.GetFolderInput) *GetFolderFuture

	GetMergeCommit(ctx workflow.Context, input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error)
	GetMergeCommitAsync(ctx workflow.Context, input *codecommit.GetMergeCommitInput) *GetMergeCommitFuture

	GetMergeConflicts(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error)
	GetMergeConflictsAsync(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) *GetMergeConflictsFuture

	GetMergeOptions(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error)
	GetMergeOptionsAsync(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) *GetMergeOptionsFuture

	GetPullRequestApprovalStates(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error)
	GetPullRequestApprovalStatesAsync(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) *GetPullRequestApprovalStatesFuture

	GetPullRequestOverrideState(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error)
	GetPullRequestOverrideStateAsync(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) *GetPullRequestOverrideStateFuture

	GetRepository(ctx workflow.Context, input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error)
	GetRepositoryAsync(ctx workflow.Context, input *codecommit.GetRepositoryInput) *GetRepositoryFuture

	GetRepositoryTriggers(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error)
	GetRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) *GetRepositoryTriggersFuture

	ListApprovalRuleTemplates(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error)
	ListApprovalRuleTemplatesAsync(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) *ListApprovalRuleTemplatesFuture

	ListAssociatedApprovalRuleTemplatesForRepository(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error)
	ListAssociatedApprovalRuleTemplatesForRepositoryAsync(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) *ListAssociatedApprovalRuleTemplatesForRepositoryFuture

	ListBranches(ctx workflow.Context, input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error)
	ListBranchesAsync(ctx workflow.Context, input *codecommit.ListBranchesInput) *ListBranchesFuture

	ListPullRequests(ctx workflow.Context, input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error)
	ListPullRequestsAsync(ctx workflow.Context, input *codecommit.ListPullRequestsInput) *ListPullRequestsFuture

	ListRepositories(ctx workflow.Context, input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error)
	ListRepositoriesAsync(ctx workflow.Context, input *codecommit.ListRepositoriesInput) *ListRepositoriesFuture

	ListRepositoriesForApprovalRuleTemplate(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error)
	ListRepositoriesForApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) *ListRepositoriesForApprovalRuleTemplateFuture

	ListTagsForResource(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) *ListTagsForResourceFuture

	MergeBranchesByFastForward(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) (*codecommit.MergeBranchesByFastForwardOutput, error)
	MergeBranchesByFastForwardAsync(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) *MergeBranchesByFastForwardFuture

	MergeBranchesBySquash(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) (*codecommit.MergeBranchesBySquashOutput, error)
	MergeBranchesBySquashAsync(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) *MergeBranchesBySquashFuture

	MergeBranchesByThreeWay(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) (*codecommit.MergeBranchesByThreeWayOutput, error)
	MergeBranchesByThreeWayAsync(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) *MergeBranchesByThreeWayFuture

	MergePullRequestByFastForward(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) (*codecommit.MergePullRequestByFastForwardOutput, error)
	MergePullRequestByFastForwardAsync(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) *MergePullRequestByFastForwardFuture

	MergePullRequestBySquash(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) (*codecommit.MergePullRequestBySquashOutput, error)
	MergePullRequestBySquashAsync(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) *MergePullRequestBySquashFuture

	MergePullRequestByThreeWay(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) (*codecommit.MergePullRequestByThreeWayOutput, error)
	MergePullRequestByThreeWayAsync(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) *MergePullRequestByThreeWayFuture

	OverridePullRequestApprovalRules(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) (*codecommit.OverridePullRequestApprovalRulesOutput, error)
	OverridePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) *OverridePullRequestApprovalRulesFuture

	PostCommentForComparedCommit(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) (*codecommit.PostCommentForComparedCommitOutput, error)
	PostCommentForComparedCommitAsync(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) *PostCommentForComparedCommitFuture

	PostCommentReply(ctx workflow.Context, input *codecommit.PostCommentReplyInput) (*codecommit.PostCommentReplyOutput, error)
	PostCommentReplyAsync(ctx workflow.Context, input *codecommit.PostCommentReplyInput) *PostCommentReplyFuture

	PutCommentReaction(ctx workflow.Context, input *codecommit.PutCommentReactionInput) (*codecommit.PutCommentReactionOutput, error)
	PutCommentReactionAsync(ctx workflow.Context, input *codecommit.PutCommentReactionInput) *PutCommentReactionFuture

	PutFile(ctx workflow.Context, input *codecommit.PutFileInput) (*codecommit.PutFileOutput, error)
	PutFileAsync(ctx workflow.Context, input *codecommit.PutFileInput) *PutFileFuture

	PutRepositoryTriggers(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error)
	PutRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) *PutRepositoryTriggersFuture

	TagResource(ctx workflow.Context, input *codecommit.TagResourceInput) (*codecommit.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codecommit.TagResourceInput) *TagResourceFuture

	TestRepositoryTriggers(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error)
	TestRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) *TestRepositoryTriggersFuture

	UntagResource(ctx workflow.Context, input *codecommit.UntagResourceInput) (*codecommit.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codecommit.UntagResourceInput) *UntagResourceFuture

	UpdateApprovalRuleTemplateContent(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error)
	UpdateApprovalRuleTemplateContentAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) *UpdateApprovalRuleTemplateContentFuture

	UpdateApprovalRuleTemplateDescription(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error)
	UpdateApprovalRuleTemplateDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) *UpdateApprovalRuleTemplateDescriptionFuture

	UpdateApprovalRuleTemplateName(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error)
	UpdateApprovalRuleTemplateNameAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) *UpdateApprovalRuleTemplateNameFuture

	UpdateComment(ctx workflow.Context, input *codecommit.UpdateCommentInput) (*codecommit.UpdateCommentOutput, error)
	UpdateCommentAsync(ctx workflow.Context, input *codecommit.UpdateCommentInput) *UpdateCommentFuture

	UpdateDefaultBranch(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error)
	UpdateDefaultBranchAsync(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) *UpdateDefaultBranchFuture

	UpdatePullRequestApprovalRuleContent(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error)
	UpdatePullRequestApprovalRuleContentAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) *UpdatePullRequestApprovalRuleContentFuture

	UpdatePullRequestApprovalState(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) (*codecommit.UpdatePullRequestApprovalStateOutput, error)
	UpdatePullRequestApprovalStateAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) *UpdatePullRequestApprovalStateFuture

	UpdatePullRequestDescription(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) (*codecommit.UpdatePullRequestDescriptionOutput, error)
	UpdatePullRequestDescriptionAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) *UpdatePullRequestDescriptionFuture

	UpdatePullRequestStatus(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) (*codecommit.UpdatePullRequestStatusOutput, error)
	UpdatePullRequestStatusAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) *UpdatePullRequestStatusFuture

	UpdatePullRequestTitle(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) (*codecommit.UpdatePullRequestTitleOutput, error)
	UpdatePullRequestTitleAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) *UpdatePullRequestTitleFuture

	UpdateRepositoryDescription(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error)
	UpdateRepositoryDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) *UpdateRepositoryDescriptionFuture

	UpdateRepositoryName(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error)
	UpdateRepositoryNameAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) *UpdateRepositoryNameFuture
}

func NewClient() Client {
	return &stub{}
}
