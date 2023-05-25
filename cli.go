package main

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{}

type params struct {
	kubeconfig          string
	namespace           string
	resource            string
	outputFormat        string
	allNamespace        bool
	allResource         bool
	notDisplayPod       bool
	notDisplayContainer bool
}

func init() {
	params := params{
		namespace: "default",
	}
	cmd.Use = "k8sr"
	cmd.Short = "Displays the CPU and Memory Requests and Limits defined in the kubernetes workload resource"
	cmd.Flags().StringVarP(&params.namespace, "namespace", "n", params.namespace, "Kubernetes namespace to use. Default to namespace configured in Kubernetes context")
	cmd.Flags().StringVarP(&params.resource, "resource", "r", params.resource, "Choose from kubernetes workload resources.\nSpecifically, choose from Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet, Job, and CronJob.")
	cmd.Flags().StringVarP(&params.kubeconfig, "kubeconfig", "k", params.kubeconfig, "Path to kubeconfig file to use.")
	cmd.Flags().StringVarP(&params.outputFormat, "output", "o", params.outputFormat, "Select the format you want to output in,\nby default it is displayed in table format,\nbut it can be displayed in csv, json, or yaml format.")
	cmd.Flags().BoolVarP(&params.allNamespace, "all-namespaces", "A", params.allNamespace, "Enable only if you want to cover all namespaces.")
	cmd.Flags().BoolVarP(&params.allResource, "all-resource", "R", params.allResource, "Enable only if you want to see all workload resources.")
	cmd.Flags().BoolVarP(&params.notDisplayPod, "not-display-pod", "p", params.notDisplayPod, "Enable only if you want to display only the Container without displaying the Pod.\nThis option is valid only if you want to display Pod resources.")
	cmd.Flags().BoolVarP(&params.notDisplayContainer, "not-display-container", "c", params.notDisplayContainer, "Enable the flag if you do not want the container to be displayed.\nThis option is valid only if you want to display Pod resources.")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return exec(params)
	}
}
