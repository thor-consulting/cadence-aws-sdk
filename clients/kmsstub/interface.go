// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package kmsstub

import (
	"github.com/aws/aws-sdk-go/service/kms"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelKeyDeletion(ctx workflow.Context, input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error)
	CancelKeyDeletionAsync(ctx workflow.Context, input *kms.CancelKeyDeletionInput) *CancelKeyDeletionFuture

	ConnectCustomKeyStore(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error)
	ConnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) *ConnectCustomKeyStoreFuture

	CreateAlias(ctx workflow.Context, input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error)
	CreateAliasAsync(ctx workflow.Context, input *kms.CreateAliasInput) *CreateAliasFuture

	CreateCustomKeyStore(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error)
	CreateCustomKeyStoreAsync(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) *CreateCustomKeyStoreFuture

	CreateGrant(ctx workflow.Context, input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error)
	CreateGrantAsync(ctx workflow.Context, input *kms.CreateGrantInput) *CreateGrantFuture

	CreateKey(ctx workflow.Context, input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error)
	CreateKeyAsync(ctx workflow.Context, input *kms.CreateKeyInput) *CreateKeyFuture

	Decrypt(ctx workflow.Context, input *kms.DecryptInput) (*kms.DecryptOutput, error)
	DecryptAsync(ctx workflow.Context, input *kms.DecryptInput) *DecryptFuture

	DeleteAlias(ctx workflow.Context, input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)
	DeleteAliasAsync(ctx workflow.Context, input *kms.DeleteAliasInput) *DeleteAliasFuture

	DeleteCustomKeyStore(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error)
	DeleteCustomKeyStoreAsync(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) *DeleteCustomKeyStoreFuture

	DeleteImportedKeyMaterial(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error)
	DeleteImportedKeyMaterialAsync(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) *DeleteImportedKeyMaterialFuture

	DescribeCustomKeyStores(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeCustomKeyStoresAsync(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) *DescribeCustomKeyStoresFuture

	DescribeKey(ctx workflow.Context, input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)
	DescribeKeyAsync(ctx workflow.Context, input *kms.DescribeKeyInput) *DescribeKeyFuture

	DisableKey(ctx workflow.Context, input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error)
	DisableKeyAsync(ctx workflow.Context, input *kms.DisableKeyInput) *DisableKeyFuture

	DisableKeyRotation(ctx workflow.Context, input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)
	DisableKeyRotationAsync(ctx workflow.Context, input *kms.DisableKeyRotationInput) *DisableKeyRotationFuture

	DisconnectCustomKeyStore(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error)
	DisconnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) *DisconnectCustomKeyStoreFuture

	EnableKey(ctx workflow.Context, input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error)
	EnableKeyAsync(ctx workflow.Context, input *kms.EnableKeyInput) *EnableKeyFuture

	EnableKeyRotation(ctx workflow.Context, input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)
	EnableKeyRotationAsync(ctx workflow.Context, input *kms.EnableKeyRotationInput) *EnableKeyRotationFuture

	Encrypt(ctx workflow.Context, input *kms.EncryptInput) (*kms.EncryptOutput, error)
	EncryptAsync(ctx workflow.Context, input *kms.EncryptInput) *EncryptFuture

	GenerateDataKey(ctx workflow.Context, input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)
	GenerateDataKeyAsync(ctx workflow.Context, input *kms.GenerateDataKeyInput) *GenerateDataKeyFuture

	GenerateDataKeyPair(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error)
	GenerateDataKeyPairAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) *GenerateDataKeyPairFuture

	GenerateDataKeyPairWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error)
	GenerateDataKeyPairWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) *GenerateDataKeyPairWithoutPlaintextFuture

	GenerateDataKeyWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)
	GenerateDataKeyWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) *GenerateDataKeyWithoutPlaintextFuture

	GenerateRandom(ctx workflow.Context, input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)
	GenerateRandomAsync(ctx workflow.Context, input *kms.GenerateRandomInput) *GenerateRandomFuture

	GetKeyPolicy(ctx workflow.Context, input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)
	GetKeyPolicyAsync(ctx workflow.Context, input *kms.GetKeyPolicyInput) *GetKeyPolicyFuture

	GetKeyRotationStatus(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)
	GetKeyRotationStatusAsync(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) *GetKeyRotationStatusFuture

	GetParametersForImport(ctx workflow.Context, input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error)
	GetParametersForImportAsync(ctx workflow.Context, input *kms.GetParametersForImportInput) *GetParametersForImportFuture

	GetPublicKey(ctx workflow.Context, input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error)
	GetPublicKeyAsync(ctx workflow.Context, input *kms.GetPublicKeyInput) *GetPublicKeyFuture

	ImportKeyMaterial(ctx workflow.Context, input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error)
	ImportKeyMaterialAsync(ctx workflow.Context, input *kms.ImportKeyMaterialInput) *ImportKeyMaterialFuture

	ListAliases(ctx workflow.Context, input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error)
	ListAliasesAsync(ctx workflow.Context, input *kms.ListAliasesInput) *ListAliasesFuture

	ListGrants(ctx workflow.Context, input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error)
	ListGrantsAsync(ctx workflow.Context, input *kms.ListGrantsInput) *ListGrantsFuture

	ListKeyPolicies(ctx workflow.Context, input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)
	ListKeyPoliciesAsync(ctx workflow.Context, input *kms.ListKeyPoliciesInput) *ListKeyPoliciesFuture

	ListKeys(ctx workflow.Context, input *kms.ListKeysInput) (*kms.ListKeysOutput, error)
	ListKeysAsync(ctx workflow.Context, input *kms.ListKeysInput) *ListKeysFuture

	ListResourceTags(ctx workflow.Context, input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error)
	ListResourceTagsAsync(ctx workflow.Context, input *kms.ListResourceTagsInput) *ListResourceTagsFuture

	ListRetirableGrants(ctx workflow.Context, input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error)
	ListRetirableGrantsAsync(ctx workflow.Context, input *kms.ListRetirableGrantsInput) *ListRetirableGrantsFuture

	PutKeyPolicy(ctx workflow.Context, input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)
	PutKeyPolicyAsync(ctx workflow.Context, input *kms.PutKeyPolicyInput) *PutKeyPolicyFuture

	ReEncrypt(ctx workflow.Context, input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error)
	ReEncryptAsync(ctx workflow.Context, input *kms.ReEncryptInput) *ReEncryptFuture

	RetireGrant(ctx workflow.Context, input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error)
	RetireGrantAsync(ctx workflow.Context, input *kms.RetireGrantInput) *RetireGrantFuture

	RevokeGrant(ctx workflow.Context, input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)
	RevokeGrantAsync(ctx workflow.Context, input *kms.RevokeGrantInput) *RevokeGrantFuture

	ScheduleKeyDeletion(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionAsync(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) *ScheduleKeyDeletionFuture

	Sign(ctx workflow.Context, input *kms.SignInput) (*kms.SignOutput, error)
	SignAsync(ctx workflow.Context, input *kms.SignInput) *SignFuture

	TagResource(ctx workflow.Context, input *kms.TagResourceInput) (*kms.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kms.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kms.UntagResourceInput) *UntagResourceFuture

	UpdateAlias(ctx workflow.Context, input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error)
	UpdateAliasAsync(ctx workflow.Context, input *kms.UpdateAliasInput) *UpdateAliasFuture

	UpdateCustomKeyStore(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error)
	UpdateCustomKeyStoreAsync(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) *UpdateCustomKeyStoreFuture

	UpdateKeyDescription(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
	UpdateKeyDescriptionAsync(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) *UpdateKeyDescriptionFuture

	Verify(ctx workflow.Context, input *kms.VerifyInput) (*kms.VerifyOutput, error)
	VerifyAsync(ctx workflow.Context, input *kms.VerifyInput) *VerifyFuture
}

func NewClient() Client {
	return &stub{}
}
