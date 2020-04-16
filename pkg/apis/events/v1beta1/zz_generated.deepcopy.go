// +build !ignore_autogenerated

/*
Copyright 2020 Google LLC

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudAuditLogsSource) DeepCopyInto(out *CloudAuditLogsSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudAuditLogsSource.
func (in *CloudAuditLogsSource) DeepCopy() *CloudAuditLogsSource {
	if in == nil {
		return nil
	}
	out := new(CloudAuditLogsSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudAuditLogsSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudAuditLogsSourceList) DeepCopyInto(out *CloudAuditLogsSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudAuditLogsSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudAuditLogsSourceList.
func (in *CloudAuditLogsSourceList) DeepCopy() *CloudAuditLogsSourceList {
	if in == nil {
		return nil
	}
	out := new(CloudAuditLogsSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudAuditLogsSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudAuditLogsSourceSpec) DeepCopyInto(out *CloudAuditLogsSourceSpec) {
	*out = *in
	in.PubSubSpec.DeepCopyInto(&out.PubSubSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudAuditLogsSourceSpec.
func (in *CloudAuditLogsSourceSpec) DeepCopy() *CloudAuditLogsSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudAuditLogsSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudAuditLogsSourceStatus) DeepCopyInto(out *CloudAuditLogsSourceStatus) {
	*out = *in
	in.PubSubStatus.DeepCopyInto(&out.PubSubStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudAuditLogsSourceStatus.
func (in *CloudAuditLogsSourceStatus) DeepCopy() *CloudAuditLogsSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudAuditLogsSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPubSubSource) DeepCopyInto(out *CloudPubSubSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPubSubSource.
func (in *CloudPubSubSource) DeepCopy() *CloudPubSubSource {
	if in == nil {
		return nil
	}
	out := new(CloudPubSubSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudPubSubSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPubSubSourceList) DeepCopyInto(out *CloudPubSubSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudPubSubSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPubSubSourceList.
func (in *CloudPubSubSourceList) DeepCopy() *CloudPubSubSourceList {
	if in == nil {
		return nil
	}
	out := new(CloudPubSubSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudPubSubSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPubSubSourceSpec) DeepCopyInto(out *CloudPubSubSourceSpec) {
	*out = *in
	in.PubSubSpec.DeepCopyInto(&out.PubSubSpec)
	if in.AckDeadline != nil {
		in, out := &in.AckDeadline, &out.AckDeadline
		*out = new(string)
		**out = **in
	}
	if in.RetentionDuration != nil {
		in, out := &in.RetentionDuration, &out.RetentionDuration
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPubSubSourceSpec.
func (in *CloudPubSubSourceSpec) DeepCopy() *CloudPubSubSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudPubSubSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPubSubSourceStatus) DeepCopyInto(out *CloudPubSubSourceStatus) {
	*out = *in
	in.PubSubStatus.DeepCopyInto(&out.PubSubStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPubSubSourceStatus.
func (in *CloudPubSubSourceStatus) DeepCopy() *CloudPubSubSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudPubSubSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSchedulerSource) DeepCopyInto(out *CloudSchedulerSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSchedulerSource.
func (in *CloudSchedulerSource) DeepCopy() *CloudSchedulerSource {
	if in == nil {
		return nil
	}
	out := new(CloudSchedulerSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSchedulerSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSchedulerSourceList) DeepCopyInto(out *CloudSchedulerSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudSchedulerSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSchedulerSourceList.
func (in *CloudSchedulerSourceList) DeepCopy() *CloudSchedulerSourceList {
	if in == nil {
		return nil
	}
	out := new(CloudSchedulerSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudSchedulerSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSchedulerSourceSpec) DeepCopyInto(out *CloudSchedulerSourceSpec) {
	*out = *in
	in.PubSubSpec.DeepCopyInto(&out.PubSubSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSchedulerSourceSpec.
func (in *CloudSchedulerSourceSpec) DeepCopy() *CloudSchedulerSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSchedulerSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSchedulerSourceStatus) DeepCopyInto(out *CloudSchedulerSourceStatus) {
	*out = *in
	in.PubSubStatus.DeepCopyInto(&out.PubSubStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSchedulerSourceStatus.
func (in *CloudSchedulerSourceStatus) DeepCopy() *CloudSchedulerSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudSchedulerSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageSource) DeepCopyInto(out *CloudStorageSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageSource.
func (in *CloudStorageSource) DeepCopy() *CloudStorageSource {
	if in == nil {
		return nil
	}
	out := new(CloudStorageSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStorageSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageSourceList) DeepCopyInto(out *CloudStorageSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStorageSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageSourceList.
func (in *CloudStorageSourceList) DeepCopy() *CloudStorageSourceList {
	if in == nil {
		return nil
	}
	out := new(CloudStorageSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStorageSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageSourceSpec) DeepCopyInto(out *CloudStorageSourceSpec) {
	*out = *in
	in.PubSubSpec.DeepCopyInto(&out.PubSubSpec)
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageSourceSpec.
func (in *CloudStorageSourceSpec) DeepCopy() *CloudStorageSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStorageSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStorageSourceStatus) DeepCopyInto(out *CloudStorageSourceStatus) {
	*out = *in
	in.PubSubStatus.DeepCopyInto(&out.PubSubStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStorageSourceStatus.
func (in *CloudStorageSourceStatus) DeepCopy() *CloudStorageSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStorageSourceStatus)
	in.DeepCopyInto(out)
	return out
}
