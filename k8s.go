package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getK8sConfig(kubeconfig *string) (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}
	if *kubeconfig != "" {
		// Do nothing.
	} else if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}

func getPodList(kubeconfig, namespaces string) (*corev1.PodList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	pods, err := clientset.CoreV1().Pods(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pods: %s", err)
	}
	return pods, nil
}

func getReplicaSetList(kubeconfig, namespaces string) (*appsv1.ReplicaSetList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	replicaSets, err := clientset.AppsV1().ReplicaSets(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get replicaSets: %s", err)
	}
	return replicaSets, nil
}

func getDeploymentList(kubeconfig, namespaces string) (*appsv1.DeploymentList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	deployments, err := clientset.AppsV1().Deployments(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployments: %s", err)
	}
	return deployments, nil
}

func getStatefulSetList(kubeconfig, namespaces string) (*appsv1.StatefulSetList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	statefulSets, err := clientset.AppsV1().StatefulSets(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get statefulSets: %s", err)
	}
	return statefulSets, nil
}

func getDaemonSetList(kubeconfig, namespaces string) (*appsv1.DaemonSetList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	daemonSets, err := clientset.AppsV1().DaemonSets(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get daemonSets: %s", err)
	}
	return daemonSets, nil
}

func getJobList(kubeconfig, namespaces string) (*batchv1.JobList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	job, err := clientset.BatchV1().Jobs(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get daemonSets: %s", err)
	}
	return job, nil
}

func getCronJobList(kubeconfig, namespaces string) (*batchv1.CronJobList, error) {
	config, err := getK8sConfig(&kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	cronJobs, err := clientset.BatchV1().CronJobs(namespaces).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get cronJobs: %s", err)
	}
	return cronJobs, nil
}
