package main

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

func displayPodList(podList *corev1.PodList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range podList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.Containers, v.Spec.InitContainers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.Containers, v.Spec.InitContainers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, pod := range podList.Items {
		for _, container := range pod.Spec.Containers {
			outputBody(padding, container, pod.Name, totalResourceList)
		}
		for _, container := range pod.Spec.InitContainers {
			outputBody(padding, container, pod.Name, totalResourceList)
		}
	}
	outputFrame(padding)
	return nil
}

func displayReplicaSetList(rsList *appsv1.ReplicaSetList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range rsList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, rs := range rsList.Items {
		for _, container := range rs.Spec.Template.Spec.Containers {
			outputBody(padding, container, rs.Name, totalResourceList)
		}
		for _, container := range rs.Spec.Template.Spec.InitContainers {
			outputBody(padding, container, rs.Name, totalResourceList)
		}
	}
	outputFrame(padding)
	return nil
}

func displayDeploymentList(deployList *appsv1.DeploymentList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range deployList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, deploy := range deployList.Items {
		for _, container := range deploy.Spec.Template.Spec.Containers {
			outputBody(padding, container, deploy.Name, totalResourceList)
		}
		for _, container := range deploy.Spec.Template.Spec.InitContainers {
			outputBody(padding, container, deploy.Name, totalResourceList)
		}
	}
	outputFrame(padding)
	return nil
}

func displayStatefulSetList(stsList *appsv1.StatefulSetList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range stsList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, sts := range stsList.Items {
		for _, container := range sts.Spec.Template.Spec.Containers {
			outputBody(padding, container, sts.Name, totalResourceList)
		}
		for _, container := range sts.Spec.Template.Spec.InitContainers {
			outputBody(padding, container, sts.Name, totalResourceList)
		}
	}
	outputFrame(padding)
	return nil
}

func displayDaemonSetList(dsList *appsv1.DaemonSetList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range dsList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.InitContainers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, ds := range dsList.Items {
		for _, container := range ds.Spec.Template.Spec.Containers {
			outputBody(padding, container, ds.Name, totalResourceList)
		}
		for _, container := range ds.Spec.Template.Spec.InitContainers {
			outputBody(padding, container, ds.Name, totalResourceList)
		}
	}
	outputFrame(padding)
	return nil
}

func displayJobList(jobList *batchv1.JobList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range jobList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.Containers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.Template.Spec.Containers, v.Spec.Template.Spec.Containers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, job := range jobList.Items {
		for _, container := range job.Spec.Template.Spec.Containers {
			outputBody(padding, container, job.Name, totalResourceList)
		}
		for _, container := range job.Spec.Template.Spec.InitContainers {
			outputBody(padding, container, job.Name, totalResourceList)
		}
	}

	outputFrame(padding)
	return nil
}

func displayCronJobList(cjList *batchv1.CronJobList) error {
	padding := maxLength{
		metadataName:  0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		limitCPU:      0,
		limitMemory:   0,
	}
	padding = retrieveHeaderPaddingLength(padding)
	totalResourceList := make(map[string]containerResources)
	for _, v := range cjList.Items {
		padding = retrieveMaxLengths(padding, v.Name, v.Spec.JobTemplate.Spec.Template.Spec.Containers, v.Spec.JobTemplate.Spec.Template.Spec.InitContainers)
		totalResourceList[v.Name] = retrieveContainerResources(v.Spec.JobTemplate.Spec.Template.Spec.Containers, v.Spec.JobTemplate.Spec.Template.Spec.InitContainers)
	}
	outputFrame(padding)
	outputHeader(padding)
	outputFrame(padding)
	for _, cj := range cjList.Items {
		for _, container := range cj.Spec.JobTemplate.Spec.Template.Spec.Containers {
			outputBody(padding, container, cj.Name, totalResourceList)
		}
		for _, container := range cj.Spec.JobTemplate.Spec.Template.Spec.InitContainers {
			outputBody(padding, container, cj.Name, totalResourceList)
		}
	}
	outputFrame(padding)
	return nil
}
