package main

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	HEADER_POD_NAME       = "POD NAME"
	HEADER_CONTAINER_NAME = "CONTAINER NAME"
	HEADER_REQUEST_CPU    = "REQUEST CPU"
	HEADER_REQUEST_MEMORY = "REQUEST MEMORY"
	HEADER_LIMIT_CPU      = "LIMIT CPU"
	HEADER_LIMIT_MEMORY   = "LIMIT MEMORY"
)

func compareLengths(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func paddingSet(name string, length int, max int) int {
	if max > len(name) {
		return max - length
	} else if max < len(name) {
		return len(name) - length
	}
	return max + 1
}

func displayPodList(podList *corev1.PodList) error {
	padding := map[string]int{
		"podName":       0,
		"containerName": 0,
		"requestCPU":    0,
		"requestMemory": 0,
		"LimitCPU":      0,
		"LimitMemory":   0,
	}
	totalResourceList := map[string]map[string]int{}
	for _, v := range podList.Items {
		var (
			totalReqCPU = 0
			totalReqMem = 0
			totalLimCPU = 0
			totalLimMem = 0
		)
		for i := range v.Spec.Containers {
			container := v.Spec.Containers[i]
			padding["podName"] = compareLengths(padding["podName"], len(v.GetName()))
			padding["containerName"] = compareLengths(padding["containerName"], len(container.Name))
			padding["requestCPU"] = compareLengths(padding["requestCPU"], container.Resources.Requests.Cpu().Size())
			padding["requestMemory"] = compareLengths(padding["requestMemory"], container.Resources.Requests.Memory().Size())
			padding["LimitCPU"] = compareLengths(padding["LimitCPU"], container.Resources.Limits.Cpu().Size())
			padding["LimitMemory"] = compareLengths(padding["LimitMemory"], container.Resources.Limits.Memory().Size())
			totalReqCPU += container.Resources.Requests.Cpu().Size()
			totalReqMem += container.Resources.Requests.Memory().Size()
			totalLimCPU += container.Resources.Limits.Cpu().Size()
			totalLimMem += container.Resources.Limits.Memory().Size()
		}
		for i := range v.Spec.InitContainers {
			container := v.Spec.InitContainers[i]
			padding["podName"] = compareLengths(padding["podName"], len(v.GetName()))
			padding["containerName"] = compareLengths(padding["containerName"], len(container.Name))
			padding["requestCPU"] = compareLengths(padding["requestCPU"], container.Resources.Requests.Cpu().Size())
			padding["requestMemory"] = compareLengths(padding["requestMemory"], container.Resources.Requests.Memory().Size())
			padding["LimitCPU"] = compareLengths(padding["LimitCPU"], container.Resources.Limits.Cpu().Size())
			padding["LimitMemory"] = compareLengths(padding["LimitMemory"], container.Resources.Limits.Memory().Size())
			totalReqCPU += container.Resources.Requests.Cpu().Size()
			totalReqMem += container.Resources.Requests.Memory().Size()
			totalLimCPU += container.Resources.Limits.Cpu().Size()
			totalLimMem += container.Resources.Limits.Memory().Size()
		}
		totalResourceList[v.GetName()]["requestCPU"] = totalReqCPU
		totalResourceList[v.GetName()]["requestMemory"] = totalReqMem
		totalResourceList[v.GetName()]["LimitCPU"] = totalLimCPU
		totalResourceList[v.GetName()]["LimitMemory"] = totalLimMem
	}

	for _, pod := range podList.Items {
		podNamePadding := paddingSet(HEADER_POD_NAME, len(pod.GetName()), padding["podName"])
		containerNamePadding := paddingSet(HEADER_CONTAINER_NAME, len(""), padding["containerName"])
		requestCPUPadding := paddingSet(HEADER_REQUEST_CPU, totalResourceList[pod.Name]["requestCPU"], padding["requestCPU"])
		requestMemoryPadding := paddingSet(HEADER_REQUEST_MEMORY, totalResourceList[pod.Name]["requestMemory"], padding["requestMemory"])
		LimitCPUPadding := paddingSet(HEADER_LIMIT_CPU, totalResourceList[pod.Name]["LimitCPU"], padding["LimitCPU"])
		LimitMemoryPadding := paddingSet(HEADER_LIMIT_MEMORY, totalResourceList[pod.Name]["LimitMemory"], padding["LimitMemory"])
	}

	return nil
}

func displayReplicaSetList(*appsv1.ReplicaSetList) error {
	return nil
}

func displayDeploymentList(*appsv1.DeploymentList) error {
	return nil
}

func displayStatefulSetList(*appsv1.StatefulSetList) error {
	return nil
}

func displayDaemonSetList(*appsv1.DaemonSetList) error {
	return nil
}

func displayJobList(*batchv1.JobList) error {
	return nil
}

func displayCronJobList(*batchv1.CronJobList) error {
	return nil
}
