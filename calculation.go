package main

import corev1 "k8s.io/api/core/v1"

type containerResources struct {
	requestCPU    int
	requestMemory int
	limitCPU      int
	limitMemory   int
}

func sumContainerResources(cr containerResources, container corev1.Container) containerResources {
	return containerResources{
		requestCPU:    cr.requestCPU + container.Resources.Requests.Cpu().Size(),
		requestMemory: cr.requestMemory + container.Resources.Requests.Memory().Size(),
		limitCPU:      cr.limitCPU + container.Resources.Limits.Cpu().Size(),
		limitMemory:   cr.limitMemory + container.Resources.Limits.Memory().Size(),
	}
}
