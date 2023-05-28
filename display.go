package main

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

const (
	headerPodName       = "POD NAME"
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
	podNamePadding := paddingSet(headerPodName, len(headerPodName), maxLengths.podName)
	containerNamePadding := paddingSet(headerContainerName, len(headerContainerName), maxLengths.containerName)
	requestCPUPadding := paddingSet(headerRequestCPU, len(headerRequestCPU), maxLengths.requestCPU)
	requestMemoryPadding := paddingSet(headerRequestMemory, len(headerRequestMemory), maxLengths.requestMemory)
	limitCPUPadding := paddingSet(headerLimitCPU, len(headerLimitCPU), maxLengths.limitCPU)
	limitMemoryPadding := paddingSet(headerLimitMemory, len(headerLimitMemory), maxLengths.limitMemory)
	fmt.Printf("| %s%s ", headerPodName, strings.Repeat(" ", podNamePadding))
	fmt.Printf("| %s%s ", headerContainerName, strings.Repeat(" ", containerNamePadding))
	fmt.Printf("| %s%s ", headerRequestCPU, strings.Repeat(" ", requestCPUPadding))
	fmt.Printf("| %s%s ", headerRequestMemory, strings.Repeat(" ", requestMemoryPadding))
	fmt.Printf("| %s%s ", headerLimitCPU, strings.Repeat(" ", limitCPUPadding))
	fmt.Printf("| %s%s ", headerLimitMemory, strings.Repeat(" ", limitMemoryPadding))
	fmt.Printf("|\n")
}

func outputFrame(maxLengths maxLength) {
	fmt.Printf("+--%s", strings.Repeat("-", maxLengths.podName))
	fmt.Printf("+--%s", strings.Repeat("-", maxLengths.containerName))
	fmt.Printf("+--%s", strings.Repeat("-", maxLengths.requestCPU))
	fmt.Printf("+--%s", strings.Repeat("-", maxLengths.requestMemory))
	fmt.Printf("+--%s", strings.Repeat("-", maxLengths.limitCPU))
	fmt.Printf("+--%s", strings.Repeat("-", maxLengths.limitMemory))
	fmt.Printf("+\n")
}

func outputBody(maxLengths maxLength, container corev1.Container, podName string, resourceList map[string]containerResources) {
	podNamePadding := paddingSet(headerPodName, len(podName), maxLengths.podName)
	requestCPUPadding := paddingSet(headerRequestCPU, len(container.Resources.Requests.Cpu().String()), maxLengths.requestCPU)
	requestMemoryPadding := paddingSet(headerRequestMemory, len(container.Resources.Requests.Memory().String()), maxLengths.requestMemory)
	limitCPUPadding := paddingSet(headerLimitCPU, len(container.Resources.Limits.Cpu().String()), maxLengths.limitCPU)
	limitMemoryPadding := paddingSet(headerLimitMemory, len(container.Resources.Limits.Memory().String()), maxLengths.limitMemory)
	containerNamePadding := paddingSet(headerContainerName, len(container.Name), maxLengths.containerName)
	fmt.Printf("| %s%s ", podName, strings.Repeat(" ", podNamePadding))
	fmt.Printf("| %s%s ", container.Name, strings.Repeat(" ", containerNamePadding))
	fmt.Printf("| %s%s ", strings.Repeat(" ", requestCPUPadding), container.Resources.Requests.Cpu().String())
	fmt.Printf("| %s%s ", strings.Repeat(" ", requestMemoryPadding), container.Resources.Requests.Memory().String())
	fmt.Printf("| %s%s ", strings.Repeat(" ", limitCPUPadding), container.Resources.Limits.Cpu().String())
	fmt.Printf("| %s%s ", strings.Repeat(" ", limitMemoryPadding), container.Resources.Limits.Memory().String())
	fmt.Printf("|\n")
}
