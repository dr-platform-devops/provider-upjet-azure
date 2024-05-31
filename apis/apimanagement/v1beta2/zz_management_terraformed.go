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

// GetTerraformResourceType returns Terraform resource type for this Management
func (mg *Management) GetTerraformResourceType() string {
	return "azurerm_api_management"
}

// GetConnectionDetailsMapping for this Management
func (tr *Management) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"certificate[*].certificate_password": "certificate[*].certificatePasswordSecretRef", "certificate[*].encoded_certificate": "certificate[*].encodedCertificateSecretRef", "delegation[*].validation_key": "delegation[*].validationKeySecretRef", "hostname_configuration[*].developer_portal[*].certificate": "status.atProvider.hostnameConfiguration[*].developerPortal[*].certificate", "hostname_configuration[*].developer_portal[*].certificate_password": "status.atProvider.hostnameConfiguration[*].developerPortal[*].certificatePassword", "hostname_configuration[*].management[*].certificate": "status.atProvider.hostnameConfiguration[*].management[*].certificate", "hostname_configuration[*].management[*].certificate_password": "status.atProvider.hostnameConfiguration[*].management[*].certificatePassword", "hostname_configuration[*].portal[*].certificate": "status.atProvider.hostnameConfiguration[*].portal[*].certificate", "hostname_configuration[*].portal[*].certificate_password": "status.atProvider.hostnameConfiguration[*].portal[*].certificatePassword", "hostname_configuration[*].proxy[*].certificate": "status.atProvider.hostnameConfiguration[*].proxy[*].certificate", "hostname_configuration[*].proxy[*].certificate_password": "status.atProvider.hostnameConfiguration[*].proxy[*].certificatePassword", "hostname_configuration[*].scm[*].certificate": "status.atProvider.hostnameConfiguration[*].scm[*].certificate", "hostname_configuration[*].scm[*].certificate_password": "status.atProvider.hostnameConfiguration[*].scm[*].certificatePassword", "tenant_access[*].primary_key": "status.atProvider.tenantAccess[*].primaryKey", "tenant_access[*].secondary_key": "status.atProvider.tenantAccess[*].secondaryKey"}
}

// GetObservation of this Management
func (tr *Management) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Management
func (tr *Management) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Management
func (tr *Management) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Management
func (tr *Management) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Management
func (tr *Management) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Management
func (tr *Management) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this Management
func (tr *Management) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
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

// LateInitialize this Management using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Management) LateInitialize(attrs []byte) (bool, error) {
	params := &ManagementParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Management) GetTerraformSchemaVersion() int {
	return 0
}
