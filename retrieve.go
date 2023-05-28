package main

import corev1 "k8s.io/api/core/v1"

type padding struct {
	podName       int
	containerName int
	requestCPU    int
	requestMemory int
	LimitCPU      int
	LimitMemory   int
}

func compareLengths(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func retrievePaddingLength(lp padding, podName string, container corev1.Container) padding {
	return padding{
		podName:       compareLengths(lp.podName, len(podName)),
		containerName: compareLengths(lp.containerName, len(container.Name)),
		requestCPU:    compareLengths(lp.requestCPU, container.Resources.Requests.Cpu().Size()),
		requestMemory: compareLengths(lp.requestMemory, container.Resources.Requests.Memory().Size()),
		LimitCPU:      compareLengths(lp.LimitCPU, container.Resources.Limits.Cpu().Size()),
		LimitMemory:   compareLengths(lp.LimitMemory, container.Resources.Limits.Memory().Size()),
	}
}

func retrievePadding(lp padding, podName string, containers, initContainers []corev1.Container) padding {
	for _, container := range containers {
		lp = retrievePaddingLength(lp, podName, container)
	}
	for _, container := range initContainers {
		lp = retrievePaddingLength(lp, podName, container)
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
