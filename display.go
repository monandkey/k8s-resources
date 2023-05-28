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
	if max > len(name) {
		return max - length
	} else if max < len(name) {
		return len(name) - length
	}
	return max + 1
}

func outputHeader(padding padding) {
	podNamePadding := paddingSet(headerPodName, len(headerPodName), padding.podName)
	containerNamePadding := paddingSet(headerContainerName, len(headerContainerName), padding.containerName)
	requestCPUPadding := paddingSet(headerRequestCPU, len(headerRequestCPU), padding.requestCPU)
	requestMemoryPadding := paddingSet(headerRequestMemory, len(headerRequestMemory), padding.requestMemory)
	limitCPUPadding := paddingSet(headerLimitCPU, len(headerLimitCPU), padding.LimitCPU)
	limitMemoryPadding := paddingSet(headerLimitMemory, len(headerLimitMemory), padding.LimitMemory)
	fmt.Printf("| %s%s ", headerPodName, strings.Repeat(" ", podNamePadding))
	fmt.Printf("| %s%s ", headerContainerName, strings.Repeat(" ", containerNamePadding))
	fmt.Printf("| %s%s ", headerRequestCPU, strings.Repeat(" ", requestCPUPadding))
	fmt.Printf("| %s%s ", headerRequestMemory, strings.Repeat(" ", requestMemoryPadding))
	fmt.Printf("| %s%s ", headerLimitCPU, strings.Repeat(" ", limitCPUPadding))
	fmt.Printf("| %s%s ", headerLimitMemory, strings.Repeat(" ", limitMemoryPadding))
	fmt.Printf("|\n")
}

func outputFrame(padding padding) {
	fmt.Printf("+%s", strings.Repeat("-", padding.podName))
	fmt.Printf("+%s", strings.Repeat("-", padding.containerName))
	fmt.Printf("+%s", strings.Repeat("-", padding.requestCPU))
	fmt.Printf("+%s", strings.Repeat("-", padding.requestMemory))
	fmt.Printf("+%s", strings.Repeat("-", padding.LimitCPU))
	fmt.Printf("+%s", strings.Repeat("-", padding.LimitMemory))
	fmt.Printf("+\n")
}

func outputBody(padding padding, container corev1.Container, podName string, resourceList map[string]containerResources) {
	podNamePadding := paddingSet(headerPodName, len(podName), padding.podName)
	requestCPUPadding := paddingSet(headerRequestCPU, resourceList[podName].requestCPU, padding.requestCPU)
	requestMemoryPadding := paddingSet(headerRequestMemory, resourceList[podName].requestMemory, padding.requestMemory)
	limitCPUPadding := paddingSet(headerLimitCPU, resourceList[podName].limitCPU, padding.LimitCPU)
	limitMemoryPadding := paddingSet(headerLimitMemory, resourceList[podName].limitMemory, padding.LimitMemory)
	containerNamePadding := paddingSet(headerContainerName, len(container.Name), padding.containerName)
	fmt.Printf("| %s%s ", podName, strings.Repeat(" ", podNamePadding))
	fmt.Printf("| %s%s ", container.Name, strings.Repeat(" ", containerNamePadding))
	fmt.Printf("| %d%s ", container.Resources.Requests.Cpu().Size(), strings.Repeat(" ", requestCPUPadding))
	fmt.Printf("| %d%s ", container.Resources.Requests.Memory().Size(), strings.Repeat(" ", requestMemoryPadding))
	fmt.Printf("| %d%s ", container.Resources.Limits.Cpu().Size(), strings.Repeat(" ", limitCPUPadding))
	fmt.Printf("| %d%s ", container.Resources.Limits.Memory().Size(), strings.Repeat(" ", limitMemoryPadding))
	fmt.Printf("|\n")
}
