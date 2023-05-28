package main

import (
	corev1 "k8s.io/api/core/v1"
)

type maxLength struct {
	podName       int
	containerName int
	requestCPU    int
	requestMemory int
	limitCPU      int
	limitMemory   int
}

func compareLengths(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func retrieveHeaderPaddingLength(lp maxLength) maxLength {
	return maxLength{
		podName:       compareLengths(lp.podName, len(headerPodName)),
		containerName: compareLengths(lp.containerName, len(headerContainerName)),
		requestCPU:    compareLengths(lp.requestCPU, len(headerRequestCPU)),
		requestMemory: compareLengths(lp.requestMemory, len(headerRequestMemory)),
		limitCPU:      compareLengths(lp.limitCPU, len(headerLimitCPU)),
		limitMemory:   compareLengths(lp.limitMemory, len(headerLimitMemory)),
	}
}

func retrieveMaxLength(lp maxLength, podName string, container corev1.Container) maxLength {
	return maxLength{
		podName:       compareLengths(lp.podName, len(podName)),
		containerName: compareLengths(lp.containerName, len(container.Name)),
		requestCPU:    compareLengths(lp.requestCPU, len(container.Resources.Requests.Cpu().String())),
		requestMemory: compareLengths(lp.requestMemory, len(container.Resources.Requests.Memory().String())),
		limitCPU:      compareLengths(lp.limitCPU, len(container.Resources.Limits.Cpu().String())),
		limitMemory:   compareLengths(lp.limitMemory, len(container.Resources.Limits.Memory().String())),
	}
}

func retrieveMaxLengths(lp maxLength, podName string, containers, initContainers []corev1.Container) maxLength {
	for _, container := range containers {
		lp = retrieveMaxLength(lp, podName, container)
	}
	for _, container := range initContainers {
		lp = retrieveMaxLength(lp, podName, container)
	}
	return lp
}

func retrieveContainerResources(containers, initContainers []corev1.Container) containerResources {
	resources := containerResources{
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	for i := range containers {
		resources = sumContainerResources(resources, containers[i])
	}
	for i := range initContainers {
		resources = sumContainerResources(resources, initContainers[i])
	}
	return resources
}
