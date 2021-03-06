/*
Copyright 2016 The Kubernetes Authors.

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

package nodeup

import (
//"k8s.io/kops/upup/pkg/fi"
)

type NodeUpConfig struct {
	// Tags enable/disable chunks of the model
	Tags []string `json:",omitempty"`
	// Assets are locations where we can find files to be installed
	// TODO: Remove once everything is in containers?
	Assets []string `json:",omitempty"`

	// Images are a list of images we should preload
	Images []*Image `json:"images,omitempty"`

	// ConfigBase is the base VFS path for config objects
	ConfigBase *string `json:",omitempty"`

	// ClusterLocation is the VFS path to the cluster spec
	// Deprecated: prefer ConfigBase
	ClusterLocation *string `json:",omitempty"`

	// InstanceGroupName is the name of the instance group
	InstanceGroupName string `json:",omitempty"`

	// ClusterName is the name of the cluster
	// Technically this is redundant - it is in ClusterLocation, but this can serve as a cross-check,
	// and it allows us to more easily identify the cluster, for example when we are deleting resources.
	ClusterName string `json:",omitempty"`

	// ProtokubeImage is the docker image to load for protokube (bootstrapping)
	ProtokubeImage *Image `json:"protokubeImage,omitempty"`

	// Channels is a list of channels that we should apply
	Channels []string `json:"channels,omitempty"`
}

// Image is a docker image we should pre-load
type Image struct {
	// Name is the name of the tagged image
	// This is the name we would pass to "docker run", whereas source could be a URL from which
	// we would download an image.
	Name string `json:"name,omitempty"`

	// Source is the URL from which we should download the image
	Source string `json:"source,omitempty"`

	// Hash is the hash of the file, to verify image integrity (even over http)
	Hash string `json:"hash,omitempty"`
}
