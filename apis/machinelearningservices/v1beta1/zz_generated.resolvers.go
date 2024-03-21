// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ComputeCluster) ResolveReferences( // ResolveReferences of this ComputeCluster.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.ForProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.ForProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetResourceID")
	}
	mg.Spec.ForProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.InitProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.InitProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.InitProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetResourceID")
	}
	mg.Spec.InitProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ComputeInstance.
func (mg *ComputeInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.ForProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.ForProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetResourceID")
	}
	mg.Spec.ForProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.InitProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetResourceID")
	}
	mg.Spec.InitProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SynapseSpark.
func (mg *SynapseSpark) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta1", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.ForProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("synapse.azure.upbound.io", "v1beta1", "SparkPool", "SparkPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseSparkPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SynapseSparkPoolIDRef,
			Selector:     mg.Spec.ForProvider.SynapseSparkPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseSparkPoolID")
	}
	mg.Spec.ForProvider.SynapseSparkPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseSparkPoolIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("synapse.azure.upbound.io", "v1beta1", "SparkPool", "SparkPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SynapseSparkPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SynapseSparkPoolIDRef,
			Selector:     mg.Spec.InitProvider.SynapseSparkPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SynapseSparkPoolID")
	}
	mg.Spec.InitProvider.SynapseSparkPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SynapseSparkPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workspace.
func (mg *Workspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("insights.azure.upbound.io", "v1beta1", "ApplicationInsights", "ApplicationInsightsList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationInsightsID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ApplicationInsightsIDRef,
			Selector:     mg.Spec.ForProvider.ApplicationInsightsIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationInsightsID")
	}
	mg.Spec.ForProvider.ApplicationInsightsID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationInsightsIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Encryption); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption[i3].KeyID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption[i3].KeyIDRef,
				Selector:     mg.Spec.ForProvider.Encryption[i3].KeyIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption[i3].KeyID")
		}
		mg.Spec.ForProvider.Encryption[i3].KeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption[i3].KeyIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Encryption); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption[i3].KeyVaultID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption[i3].KeyVaultIDRef,
				Selector:     mg.Spec.ForProvider.Encryption[i3].KeyVaultIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption[i3].KeyVaultID")
		}
		mg.Spec.ForProvider.Encryption[i3].KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption[i3].KeyVaultIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Encryption); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption[i3].UserAssignedIdentityID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption[i3].UserAssignedIdentityIDRef,
				Selector:     mg.Spec.ForProvider.Encryption[i3].UserAssignedIdentityIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption[i3].UserAssignedIdentityID")
		}
		mg.Spec.ForProvider.Encryption[i3].UserAssignedIdentityID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption[i3].UserAssignedIdentityIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrimaryUserAssignedIdentity),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrimaryUserAssignedIdentityRef,
			Selector:     mg.Spec.ForProvider.PrimaryUserAssignedIdentitySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrimaryUserAssignedIdentity")
	}
	mg.Spec.ForProvider.PrimaryUserAssignedIdentity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrimaryUserAssignedIdentityRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("insights.azure.upbound.io", "v1beta1", "ApplicationInsights", "ApplicationInsightsList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationInsightsID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ApplicationInsightsIDRef,
			Selector:     mg.Spec.InitProvider.ApplicationInsightsIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationInsightsID")
	}
	mg.Spec.InitProvider.ApplicationInsightsID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationInsightsIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Encryption); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption[i3].KeyID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption[i3].KeyIDRef,
				Selector:     mg.Spec.InitProvider.Encryption[i3].KeyIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption[i3].KeyID")
		}
		mg.Spec.InitProvider.Encryption[i3].KeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption[i3].KeyIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Encryption); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption[i3].KeyVaultID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption[i3].KeyVaultIDRef,
				Selector:     mg.Spec.InitProvider.Encryption[i3].KeyVaultIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption[i3].KeyVaultID")
		}
		mg.Spec.InitProvider.Encryption[i3].KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption[i3].KeyVaultIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Encryption); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption[i3].UserAssignedIdentityID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption[i3].UserAssignedIdentityIDRef,
				Selector:     mg.Spec.InitProvider.Encryption[i3].UserAssignedIdentityIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption[i3].UserAssignedIdentityID")
		}
		mg.Spec.InitProvider.Encryption[i3].UserAssignedIdentityID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption[i3].UserAssignedIdentityIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta1", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrimaryUserAssignedIdentity),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PrimaryUserAssignedIdentityRef,
			Selector:     mg.Spec.InitProvider.PrimaryUserAssignedIdentitySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrimaryUserAssignedIdentity")
	}
	mg.Spec.InitProvider.PrimaryUserAssignedIdentity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrimaryUserAssignedIdentityRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.StorageAccountIDRef,
			Selector:     mg.Spec.InitProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccountID")
	}
	mg.Spec.InitProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageAccountIDRef = rsp.ResolvedReference

	return nil
}
