// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	ujconversion "github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this CapacityReservation to the hub type.
func (tr *CapacityReservation) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the CapacityReservation type.
func (tr *CapacityReservation) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this DiskEncryptionSet to the hub type.
func (tr *DiskEncryptionSet) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the DiskEncryptionSet type.
func (tr *DiskEncryptionSet) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this GalleryApplicationVersion to the hub type.
func (tr *GalleryApplicationVersion) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the GalleryApplicationVersion type.
func (tr *GalleryApplicationVersion) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Image to the hub type.
func (tr *Image) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Image type.
func (tr *Image) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this LinuxVirtualMachine to the hub type.
func (tr *LinuxVirtualMachine) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the LinuxVirtualMachine type.
func (tr *LinuxVirtualMachine) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this LinuxVirtualMachineScaleSet to the hub type.
func (tr *LinuxVirtualMachineScaleSet) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the LinuxVirtualMachineScaleSet type.
func (tr *LinuxVirtualMachineScaleSet) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this ManagedDisk to the hub type.
func (tr *ManagedDisk) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the ManagedDisk type.
func (tr *ManagedDisk) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this OrchestratedVirtualMachineScaleSet to the hub type.
func (tr *OrchestratedVirtualMachineScaleSet) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the OrchestratedVirtualMachineScaleSet type.
func (tr *OrchestratedVirtualMachineScaleSet) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this SharedImage to the hub type.
func (tr *SharedImage) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the SharedImage type.
func (tr *SharedImage) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this SharedImageGallery to the hub type.
func (tr *SharedImageGallery) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the SharedImageGallery type.
func (tr *SharedImageGallery) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this Snapshot to the hub type.
func (tr *Snapshot) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Snapshot type.
func (tr *Snapshot) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this VirtualMachineExtension to the hub type.
func (tr *VirtualMachineExtension) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the VirtualMachineExtension type.
func (tr *VirtualMachineExtension) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this VirtualMachineRunCommand to the hub type.
func (tr *VirtualMachineRunCommand) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the VirtualMachineRunCommand type.
func (tr *VirtualMachineRunCommand) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this WindowsVirtualMachine to the hub type.
func (tr *WindowsVirtualMachine) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the WindowsVirtualMachine type.
func (tr *WindowsVirtualMachine) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this WindowsVirtualMachineScaleSet to the hub type.
func (tr *WindowsVirtualMachineScaleSet) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the WindowsVirtualMachineScaleSet type.
func (tr *WindowsVirtualMachineScaleSet) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}