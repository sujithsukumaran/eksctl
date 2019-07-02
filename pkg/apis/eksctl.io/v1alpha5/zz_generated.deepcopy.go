// +build !ignore_autogenerated

/*
Copyright 2018 Weaveworks. All rights reserved.

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

package v1alpha5

import (
	ipnet "github.com/weaveworks/eksctl/pkg/utils/ipnet"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(ClusterMeta)
		(*in).DeepCopyInto(*out)
	}
	out.IAM = in.IAM
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(ClusterVPC)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]*NodeGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigList) DeepCopyInto(out *ClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigList.
func (in *ClusterConfigList) DeepCopy() *ClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIAM) DeepCopyInto(out *ClusterIAM) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIAM.
func (in *ClusterIAM) DeepCopy() *ClusterIAM {
	if in == nil {
		return nil
	}
	out := new(ClusterIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMeta) DeepCopyInto(out *ClusterMeta) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMeta.
func (in *ClusterMeta) DeepCopy() *ClusterMeta {
	if in == nil {
		return nil
	}
	out := new(ClusterMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNAT) DeepCopyInto(out *ClusterNAT) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNAT.
func (in *ClusterNAT) DeepCopy() *ClusterNAT {
	if in == nil {
		return nil
	}
	out := new(ClusterNAT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSubnets) DeepCopyInto(out *ClusterSubnets) {
	*out = *in
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = make(map[string]Network, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = make(map[string]Network, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSubnets.
func (in *ClusterSubnets) DeepCopy() *ClusterSubnets {
	if in == nil {
		return nil
	}
	out := new(ClusterSubnets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVPC) DeepCopyInto(out *ClusterVPC) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = new(ClusterSubnets)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraCIDRs != nil {
		in, out := &in.ExtraCIDRs, &out.ExtraCIDRs
		*out = make([]*ipnet.IPNet, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.AutoAllocateIPv6 != nil {
		in, out := &in.AutoAllocateIPv6, &out.AutoAllocateIPv6
		*out = new(bool)
		**out = **in
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(ClusterNAT)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVPC.
func (in *ClusterVPC) DeepCopy() *ClusterVPC {
	if in == nil {
		return nil
	}
	out := new(ClusterVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.CIDR != nil {
		in, out := &in.CIDR, &out.CIDR
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.InstancesDistribution != nil {
		in, out := &in.InstancesDistribution, &out.InstancesDistribution
		*out = new(NodeGroupInstancesDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = new(NodeGroupSGs)
		(*in).DeepCopyInto(*out)
	}
	if in.DesiredCapacity != nil {
		in, out := &in.DesiredCapacity, &out.DesiredCapacity
		*out = new(int)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
	if in.VolumeEncrypted != nil {
		in, out := &in.VolumeEncrypted, &out.VolumeEncrypted
		*out = new(bool)
		**out = **in
	}
	if in.VolumeKmsKeyID != nil {
		in, out := &in.VolumeKmsKeyID, &out.VolumeKmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetGroupARNs != nil {
		in, out := &in.TargetGroupARNs, &out.TargetGroupARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		*out = new(NodeGroupSSH)
		(*in).DeepCopyInto(*out)
	}
	if in.IAM != nil {
		in, out := &in.IAM, &out.IAM
		*out = new(NodeGroupIAM)
		(*in).DeepCopyInto(*out)
	}
	if in.PreBootstrapCommands != nil {
		in, out := &in.PreBootstrapCommands, &out.PreBootstrapCommands
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OverrideBootstrapCommand != nil {
		in, out := &in.OverrideBootstrapCommand, &out.OverrideBootstrapCommand
		*out = new(string)
		**out = **in
	}
	if in.KubeletExtraConfig != nil {
		in, out := &in.KubeletExtraConfig, &out.KubeletExtraConfig
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupIAM) DeepCopyInto(out *NodeGroupIAM) {
	*out = *in
	if in.AttachPolicyARNs != nil {
		in, out := &in.AttachPolicyARNs, &out.AttachPolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.WithAddonPolicies.DeepCopyInto(&out.WithAddonPolicies)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupIAM.
func (in *NodeGroupIAM) DeepCopy() *NodeGroupIAM {
	if in == nil {
		return nil
	}
	out := new(NodeGroupIAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupIAMAddonPolicies) DeepCopyInto(out *NodeGroupIAMAddonPolicies) {
	*out = *in
	if in.ImageBuilder != nil {
		in, out := &in.ImageBuilder, &out.ImageBuilder
		*out = new(bool)
		**out = **in
	}
	if in.AutoScaler != nil {
		in, out := &in.AutoScaler, &out.AutoScaler
		*out = new(bool)
		**out = **in
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = new(bool)
		**out = **in
	}
	if in.CertManager != nil {
		in, out := &in.CertManager, &out.CertManager
		*out = new(bool)
		**out = **in
	}
	if in.AppMesh != nil {
		in, out := &in.AppMesh, &out.AppMesh
		*out = new(bool)
		**out = **in
	}
	if in.EBS != nil {
		in, out := &in.EBS, &out.EBS
		*out = new(bool)
		**out = **in
	}
	if in.FSX != nil {
		in, out := &in.FSX, &out.FSX
		*out = new(bool)
		**out = **in
	}
	if in.EFS != nil {
		in, out := &in.EFS, &out.EFS
		*out = new(bool)
		**out = **in
	}
	if in.ALBIngress != nil {
		in, out := &in.ALBIngress, &out.ALBIngress
		*out = new(bool)
		**out = **in
	}
	if in.XRay != nil {
		in, out := &in.XRay, &out.XRay
		*out = new(bool)
		**out = **in
	}
	if in.CloudWatch != nil {
		in, out := &in.CloudWatch, &out.CloudWatch
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupIAMAddonPolicies.
func (in *NodeGroupIAMAddonPolicies) DeepCopy() *NodeGroupIAMAddonPolicies {
	if in == nil {
		return nil
	}
	out := new(NodeGroupIAMAddonPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupInstancesDistribution) DeepCopyInto(out *NodeGroupInstancesDistribution) {
	*out = *in
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		*out = new(float64)
		**out = **in
	}
	if in.OnDemandBaseCapacity != nil {
		in, out := &in.OnDemandBaseCapacity, &out.OnDemandBaseCapacity
		*out = new(int)
		**out = **in
	}
	if in.OnDemandPercentageAboveBaseCapacity != nil {
		in, out := &in.OnDemandPercentageAboveBaseCapacity, &out.OnDemandPercentageAboveBaseCapacity
		*out = new(int)
		**out = **in
	}
	if in.SpotInstancePools != nil {
		in, out := &in.SpotInstancePools, &out.SpotInstancePools
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupInstancesDistribution.
func (in *NodeGroupInstancesDistribution) DeepCopy() *NodeGroupInstancesDistribution {
	if in == nil {
		return nil
	}
	out := new(NodeGroupInstancesDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in NodeGroupKubeletConfig) DeepCopyInto(out *NodeGroupKubeletConfig) {
	{
		in := &in
		clone := in.DeepCopy()
		*out = *clone
		return
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSGs) DeepCopyInto(out *NodeGroupSGs) {
	*out = *in
	if in.AttachIDs != nil {
		in, out := &in.AttachIDs, &out.AttachIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WithShared != nil {
		in, out := &in.WithShared, &out.WithShared
		*out = new(bool)
		**out = **in
	}
	if in.WithLocal != nil {
		in, out := &in.WithLocal, &out.WithLocal
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSGs.
func (in *NodeGroupSGs) DeepCopy() *NodeGroupSGs {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSGs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSSH) DeepCopyInto(out *NodeGroupSSH) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = new(bool)
		**out = **in
	}
	if in.PublicKeyPath != nil {
		in, out := &in.PublicKeyPath, &out.PublicKeyPath
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(string)
		**out = **in
	}
	if in.PublicKeyName != nil {
		in, out := &in.PublicKeyName, &out.PublicKeyName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSSH.
func (in *NodeGroupSSH) DeepCopy() *NodeGroupSSH {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}
