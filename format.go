package main

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

func displayPodList(podList *corev1.PodList) error {
	padding := padding{
		podName:       0,
		containerName: 0,
		requestCPU:    0,
		requestMemory: 0,
		LimitCPU:      0,
		LimitMemory:   0,
	}
	totalResourceList := make(map[string]containerResources)
	for _, v := range podList.Items {
		padding = retrievePadding(padding, v.Name, v.Spec.Containers, v.Spec.InitContainers)
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
