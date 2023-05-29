package main

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

const (
	headerPodName       = "METADATA NAME"
	headerContainerName = "CONTAINER NAME"
	headerRequestCPU    = "REQUEST CPU"
	headerRequestMemory = "REQUEST MEMORY"
	headerLimitCPU      = "LIMIT CPU"
	headerLimitMemory   = "LIMIT MEMORY"
)

func paddingSet(name string, length int, max int) int {
	if max >= len(name) {
		return max - length
	} else if max <= len(name) {
		return len(name) - length
	}
	return max + 1
}

func outputHeader(maxLengths maxLength) {
	podNamePadding := paddingSet(headerPodName, len(headerPodName), maxLengths.metadataName)
	containerNamePadding := paddingSet(headerContainerName, len(headerContainerName), maxLengths.containerName)
	requestCPUPadding := paddingSet(headerRequestCPU, len(headerRequestCPU), maxLengths.requestCPU)
	requestMemoryPadding := paddingSet(headerRequestMemory, len(headerRequestMemory), maxLengths.requestMemory)
	limitCPUPadding := paddingSet(headerLimitCPU, len(headerLimitCPU), maxLengths.limitCPU)
	limitMemoryPadding := paddingSet(headerLimitMemory, len(headerLimitMemory), maxLengths.limitMemory)
	fmt.Printf(
		"| %s%s | %s%s | %s%s | %s%s | %s%s | %s%s |\n",
		headerPodName, strings.Repeat(" ", podNamePadding),
		headerContainerName, strings.Repeat(" ", containerNamePadding),
		headerRequestCPU, strings.Repeat(" ", requestCPUPadding),
		headerRequestMemory, strings.Repeat(" ", requestMemoryPadding),
		headerLimitCPU, strings.Repeat(" ", limitCPUPadding),
		headerLimitMemory, strings.Repeat(" ", limitMemoryPadding),
	)
}

func outputFrame(maxLengths maxLength) {
	fmt.Printf(
		"+--%s+--%s+--%s+--%s+--%s+--%s+\n",
		strings.Repeat("-", maxLengths.metadataName),
		strings.Repeat("-", maxLengths.containerName),
		strings.Repeat("-", maxLengths.requestCPU),
		strings.Repeat("-", maxLengths.requestMemory),
		strings.Repeat("-", maxLengths.limitCPU),
		strings.Repeat("-", maxLengths.limitMemory),
	)
}

func outputBody(maxLengths maxLength, container corev1.Container, podName string, resourceList map[string]containerResources) {
	podNamePadding := paddingSet(headerPodName, len(podName), maxLengths.metadataName)
	requestCPUPadding := paddingSet(headerRequestCPU, len(container.Resources.Requests.Cpu().String()), maxLengths.requestCPU)
	requestMemoryPadding := paddingSet(headerRequestMemory, len(container.Resources.Requests.Memory().String()), maxLengths.requestMemory)
	limitCPUPadding := paddingSet(headerLimitCPU, len(container.Resources.Limits.Cpu().String()), maxLengths.limitCPU)
	limitMemoryPadding := paddingSet(headerLimitMemory, len(container.Resources.Limits.Memory().String()), maxLengths.limitMemory)
	containerNamePadding := paddingSet(headerContainerName, len(container.Name), maxLengths.containerName)
	fmt.Printf(
		"| %s%s | %s%s | %s%s | %s%s | %s%s | %s%s |\n",
		podName, strings.Repeat(" ", podNamePadding),
		container.Name, strings.Repeat(" ", containerNamePadding),
		strings.Repeat(" ", requestCPUPadding), container.Resources.Requests.Cpu().String(),
		strings.Repeat(" ", requestMemoryPadding), container.Resources.Requests.Memory().String(),
		strings.Repeat(" ", limitCPUPadding), container.Resources.Limits.Cpu().String(),
		strings.Repeat(" ", limitMemoryPadding), container.Resources.Limits.Memory().String(),
	)
}
