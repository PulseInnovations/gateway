// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/ptr"
)

// DefaultEnvoyProxySpec returns an EnvoyProxySpec with default settings
// for infrastructure deployment. These defaults are used when no
// EnvoyProxyTemplate is configured.
func DefaultEnvoyProxySpec() *EnvoyProxySpec {
	return &EnvoyProxySpec{
		Provider: &EnvoyProxyProvider{
			Type: EnvoyProxyProviderTypeKubernetes,
			Kubernetes: &EnvoyProxyKubernetesProvider{
				EnvoyDeployment: defaultEnvoyDeployment(),
				EnvoyDaemonSet:  defaultEnvoyDaemonSet(),
			},
		},
	}
}

// defaultEnvoyDeployment returns default KubernetesDeploymentSpec for Envoy Proxy
func defaultEnvoyDeployment() *KubernetesDeploymentSpec {
	return &KubernetesDeploymentSpec{
		Replicas: ptr.To(int32(DefaultDeploymentReplicas)),
		Strategy: &appsv1.DeploymentStrategy{
			Type: appsv1.RollingUpdateDeploymentStrategyType,
		},
		Pod:       defaultEnvoyPodSpec(),
		Container: defaultEnvoyContainerSpec(),
	}
}

// defaultEnvoyDaemonSet returns default KubernetesDaemonSetSpec for Envoy Proxy
func defaultEnvoyDaemonSet() *KubernetesDaemonSetSpec {
	return &KubernetesDaemonSetSpec{
		Pod:       defaultEnvoyPodSpec(),
		Container: defaultEnvoyContainerSpec(),
	}
}

// defaultEnvoyPodSpec returns default KubernetesPodSpec for Envoy Proxy pods
func defaultEnvoyPodSpec() *KubernetesPodSpec {
	return &KubernetesPodSpec{
		SecurityContext: &corev1.PodSecurityContext{},
	}
}

// defaultEnvoyContainerSpec returns default KubernetesContainerSpec for Envoy container
func defaultEnvoyContainerSpec() *KubernetesContainerSpec {
	return &KubernetesContainerSpec{
		Image: ptr.To(DefaultEnvoyProxyImage),
		Resources: &corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse(DefaultDeploymentCPUResourceRequests),
				corev1.ResourceMemory: resource.MustParse(DefaultDeploymentMemoryResourceRequests),
			},
		},
		SecurityContext: defaultEnvoyContainerSecurityContext(),
	}
}

// defaultEnvoyContainerSecurityContext returns the default security context for Envoy container
func defaultEnvoyContainerSecurityContext() *corev1.SecurityContext {
	return &corev1.SecurityContext{
		RunAsUser:  ptr.To(int64(65532)),
		RunAsGroup: ptr.To(int64(65532)),
		Capabilities: &corev1.Capabilities{
			Drop: []corev1.Capability{"ALL"},
		},
		AllowPrivilegeEscalation: ptr.To(false),
		Privileged:               ptr.To(false),
		RunAsNonRoot:             ptr.To(true),
		SeccompProfile: &corev1.SeccompProfile{
			Type: corev1.SeccompProfileTypeRuntimeDefault,
		},
		// ReadOnlyRootFilesystem is not set to allow Envoy to write to log files/UDS sockets
	}
}

// DefaultShutdownManagerContainerSpec returns default KubernetesContainerSpec for shutdown manager
func DefaultShutdownManagerContainerSpec() *KubernetesContainerSpec {
	return &KubernetesContainerSpec{
		Image: ptr.To(DefaultShutdownManagerImage),
		Resources: &corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse(DefaultShutdownManagerCPUResourceRequests),
				corev1.ResourceMemory: resource.MustParse(DefaultShutdownManagerMemoryResourceRequests),
			},
		},
		SecurityContext: defaultShutdownManagerSecurityContext(),
	}
}

// defaultShutdownManagerSecurityContext returns the default security context for shutdown manager
func defaultShutdownManagerSecurityContext() *corev1.SecurityContext {
	return &corev1.SecurityContext{
		RunAsUser:  ptr.To(int64(65532)),
		RunAsGroup: ptr.To(int64(65532)),
		Capabilities: &corev1.Capabilities{
			Drop: []corev1.Capability{"ALL"},
		},
		AllowPrivilegeEscalation: ptr.To(false),
		Privileged:               ptr.To(false),
		RunAsNonRoot:             ptr.To(true),
		SeccompProfile: &corev1.SeccompProfile{
			Type: corev1.SeccompProfileTypeRuntimeDefault,
		},
		// ReadOnlyRootFilesystem is not set to allow shutdown manager to create files
	}
}
