package main

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{}

type params struct {
	namespace           string
	allNamespace        bool
	kubeconfig          string
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
	cmd.Flags().StringVarP(&params.kubeconfig, "kubeconfig", "k", params.kubeconfig, " Path to kubeconfig file to use")
	cmd.Flags().BoolVarP(&params.allNamespace, "all-namespaces", "A", params.allNamespace, "")
	cmd.Flags().BoolVarP(&params.notDisplayPod, "not-display-pod", "p", params.notDisplayPod, "")
	cmd.Flags().BoolVarP(&params.notDisplayContainer, "not-display-container", "c", params.notDisplayContainer, "")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return exec(params)
	}
}
