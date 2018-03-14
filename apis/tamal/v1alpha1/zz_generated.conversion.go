// +build !ignore_autogenerated

/*
Copyright 2018 The Kubepack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	unsafe "unsafe"

	tamal "github.com/kubepack/packserver/apis/tamal"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Pack_To_tamal_Pack,
		Convert_tamal_Pack_To_v1alpha1_Pack,
		Convert_v1alpha1_PackList_To_tamal_PackList,
		Convert_tamal_PackList_To_v1alpha1_PackList,
		Convert_v1alpha1_PackSpec_To_tamal_PackSpec,
		Convert_tamal_PackSpec_To_v1alpha1_PackSpec,
		Convert_v1alpha1_PackStatus_To_tamal_PackStatus,
		Convert_tamal_PackStatus_To_v1alpha1_PackStatus,
		Convert_v1alpha1_User_To_tamal_User,
		Convert_tamal_User_To_v1alpha1_User,
		Convert_v1alpha1_UserList_To_tamal_UserList,
		Convert_tamal_UserList_To_v1alpha1_UserList,
	)
}

func autoConvert_v1alpha1_Pack_To_tamal_Pack(in *Pack, out *tamal.Pack, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PackSpec_To_tamal_PackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PackStatus_To_tamal_PackStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Pack_To_tamal_Pack is an autogenerated conversion function.
func Convert_v1alpha1_Pack_To_tamal_Pack(in *Pack, out *tamal.Pack, s conversion.Scope) error {
	return autoConvert_v1alpha1_Pack_To_tamal_Pack(in, out, s)
}

func autoConvert_tamal_Pack_To_v1alpha1_Pack(in *tamal.Pack, out *Pack, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_tamal_PackSpec_To_v1alpha1_PackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_tamal_PackStatus_To_v1alpha1_PackStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_tamal_Pack_To_v1alpha1_Pack is an autogenerated conversion function.
func Convert_tamal_Pack_To_v1alpha1_Pack(in *tamal.Pack, out *Pack, s conversion.Scope) error {
	return autoConvert_tamal_Pack_To_v1alpha1_Pack(in, out, s)
}

func autoConvert_v1alpha1_PackList_To_tamal_PackList(in *PackList, out *tamal.PackList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]tamal.Pack)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PackList_To_tamal_PackList is an autogenerated conversion function.
func Convert_v1alpha1_PackList_To_tamal_PackList(in *PackList, out *tamal.PackList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackList_To_tamal_PackList(in, out, s)
}

func autoConvert_tamal_PackList_To_v1alpha1_PackList(in *tamal.PackList, out *PackList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Pack)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_tamal_PackList_To_v1alpha1_PackList is an autogenerated conversion function.
func Convert_tamal_PackList_To_v1alpha1_PackList(in *tamal.PackList, out *PackList, s conversion.Scope) error {
	return autoConvert_tamal_PackList_To_v1alpha1_PackList(in, out, s)
}

func autoConvert_v1alpha1_PackSpec_To_tamal_PackSpec(in *PackSpec, out *tamal.PackSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_PackSpec_To_tamal_PackSpec is an autogenerated conversion function.
func Convert_v1alpha1_PackSpec_To_tamal_PackSpec(in *PackSpec, out *tamal.PackSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackSpec_To_tamal_PackSpec(in, out, s)
}

func autoConvert_tamal_PackSpec_To_v1alpha1_PackSpec(in *tamal.PackSpec, out *PackSpec, s conversion.Scope) error {
	return nil
}

// Convert_tamal_PackSpec_To_v1alpha1_PackSpec is an autogenerated conversion function.
func Convert_tamal_PackSpec_To_v1alpha1_PackSpec(in *tamal.PackSpec, out *PackSpec, s conversion.Scope) error {
	return autoConvert_tamal_PackSpec_To_v1alpha1_PackSpec(in, out, s)
}

func autoConvert_v1alpha1_PackStatus_To_tamal_PackStatus(in *PackStatus, out *tamal.PackStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_PackStatus_To_tamal_PackStatus is an autogenerated conversion function.
func Convert_v1alpha1_PackStatus_To_tamal_PackStatus(in *PackStatus, out *tamal.PackStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PackStatus_To_tamal_PackStatus(in, out, s)
}

func autoConvert_tamal_PackStatus_To_v1alpha1_PackStatus(in *tamal.PackStatus, out *PackStatus, s conversion.Scope) error {
	return nil
}

// Convert_tamal_PackStatus_To_v1alpha1_PackStatus is an autogenerated conversion function.
func Convert_tamal_PackStatus_To_v1alpha1_PackStatus(in *tamal.PackStatus, out *PackStatus, s conversion.Scope) error {
	return autoConvert_tamal_PackStatus_To_v1alpha1_PackStatus(in, out, s)
}

func autoConvert_v1alpha1_User_To_tamal_User(in *User, out *tamal.User, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisallowedPacks = *(*[]string)(unsafe.Pointer(&in.DisallowedPacks))
	return nil
}

// Convert_v1alpha1_User_To_tamal_User is an autogenerated conversion function.
func Convert_v1alpha1_User_To_tamal_User(in *User, out *tamal.User, s conversion.Scope) error {
	return autoConvert_v1alpha1_User_To_tamal_User(in, out, s)
}

func autoConvert_tamal_User_To_v1alpha1_User(in *tamal.User, out *User, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisallowedPacks = *(*[]string)(unsafe.Pointer(&in.DisallowedPacks))
	return nil
}

// Convert_tamal_User_To_v1alpha1_User is an autogenerated conversion function.
func Convert_tamal_User_To_v1alpha1_User(in *tamal.User, out *User, s conversion.Scope) error {
	return autoConvert_tamal_User_To_v1alpha1_User(in, out, s)
}

func autoConvert_v1alpha1_UserList_To_tamal_UserList(in *UserList, out *tamal.UserList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]tamal.User)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_UserList_To_tamal_UserList is an autogenerated conversion function.
func Convert_v1alpha1_UserList_To_tamal_UserList(in *UserList, out *tamal.UserList, s conversion.Scope) error {
	return autoConvert_v1alpha1_UserList_To_tamal_UserList(in, out, s)
}

func autoConvert_tamal_UserList_To_v1alpha1_UserList(in *tamal.UserList, out *UserList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]User)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_tamal_UserList_To_v1alpha1_UserList is an autogenerated conversion function.
func Convert_tamal_UserList_To_v1alpha1_UserList(in *tamal.UserList, out *UserList, s conversion.Scope) error {
	return autoConvert_tamal_UserList_To_v1alpha1_UserList(in, out, s)
}