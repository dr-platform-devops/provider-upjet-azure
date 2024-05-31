// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this LinuxFunctionApp
func (mg *LinuxFunctionApp) GetTerraformResourceType() string {
	return "azurerm_linux_function_app"
}

// GetConnectionDetailsMapping for this LinuxFunctionApp
func (tr *LinuxFunctionApp) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"auth_settings[*].active_directory[*].client_secret": "authSettings[*].activeDirectory[*].clientSecretSecretRef", "auth_settings[*].facebook[*].app_secret": "authSettings[*].facebook[*].appSecretSecretRef", "auth_settings[*].github[*].client_secret": "authSettings[*].github[*].clientSecretSecretRef", "auth_settings[*].google[*].client_secret": "authSettings[*].google[*].clientSecretSecretRef", "auth_settings[*].microsoft[*].client_secret": "authSettings[*].microsoft[*].clientSecretSecretRef", "auth_settings[*].twitter[*].consumer_secret": "authSettings[*].twitter[*].consumerSecretSecretRef", "backup[*].storage_account_url": "backup[*].storageAccountUrlSecretRef", "connection_string[*].value": "connectionString[*].valueSecretRef", "custom_domain_verification_id": "status.atProvider.customDomainVerificationId", "site_config[*].application_insights_connection_string": "siteConfig[*].applicationInsightsConnectionStringSecretRef", "site_config[*].application_insights_key": "siteConfig[*].applicationInsightsKeySecretRef", "site_config[*].application_stack[*].docker[*].registry_password": "siteConfig[*].applicationStack[*].docker[*].registryPasswordSecretRef", "site_config[*].application_stack[*].docker[*].registry_username": "siteConfig[*].applicationStack[*].docker[*].registryUsernameSecretRef", "site_credential[*]": "status.atProvider.siteCredential[*]", "storage_account[*].access_key": "storageAccount[*].accessKeySecretRef", "storage_account_access_key": "storageAccountAccessKeySecretRef"}
}

// GetObservation of this LinuxFunctionApp
func (tr *LinuxFunctionApp) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LinuxFunctionApp
func (tr *LinuxFunctionApp) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LinuxFunctionApp
func (tr *LinuxFunctionApp) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LinuxFunctionApp
func (tr *LinuxFunctionApp) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LinuxFunctionApp
func (tr *LinuxFunctionApp) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this LinuxFunctionApp
func (tr *LinuxFunctionApp) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this LinuxFunctionApp
func (tr *LinuxFunctionApp) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this LinuxFunctionApp using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LinuxFunctionApp) LateInitialize(attrs []byte) (bool, error) {
	params := &LinuxFunctionAppParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("KeyVaultReferenceIdentityID"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LinuxFunctionApp) GetTerraformSchemaVersion() int {
	return 1
}
