package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

type parameter struct {
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
	params := parameter{
		kubeconfig:          "",
		namespace:           "default",
		resource:            "",
		outputFormat:        "table",
		allNamespace:        false,
		allResource:         false,
		notDisplayPod:       false,
		notDisplayContainer: false,
	}

	rootCmd.Use = "k8sr"
	rootCmd.Version = "0.0.1"
	rootCmd.SilenceUsage = true
	rootCmd.Short = "Displays the CPU and Memory Requests and Limits defined in the kubernetes workload resource"
	rootCmd.Flags().StringVarP(&params.namespace, "namespace", "n", params.namespace, "Kubernetes namespace to use. Default to namespace configured in Kubernetes context")
	rootCmd.Flags().StringVarP(&params.resource, "resource", "r", params.resource, "Choose from kubernetes workload resources.\nSpecifically, choose from Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet, Job, and CronJob.")
	rootCmd.Flags().StringVarP(&params.kubeconfig, "kubeconfig", "k", params.kubeconfig, "Path to kubeconfig file to use.")
	rootCmd.Flags().StringVarP(&params.outputFormat, "output", "o", params.outputFormat, "Select the format you want to output in,\nby default it is displayed in table format,\nbut it can be displayed in csv, json, or yaml format.")
	rootCmd.Flags().BoolVarP(&params.allNamespace, "all-namespaces", "A", params.allNamespace, "Enable only if you want to cover all namespaces.")
	rootCmd.Flags().BoolVarP(&params.allResource, "all-resource", "R", params.allResource, "Enable only if you want to see all workload resources.")
	rootCmd.Flags().BoolVarP(&params.notDisplayPod, "not-display-pod", "p", params.notDisplayPod, "Enable only if you want to display only the Container without displaying the Pod.\nThis option is valid only if you want to display Pod resources.")
	rootCmd.Flags().BoolVarP(&params.notDisplayContainer, "not-display-container", "c", params.notDisplayContainer, "Enable the flag if you do not want the container to be displayed.\nThis option is valid only if you want to display Pod resources.")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.resource == "" && !params.allResource {
			return rootCmd.Help()
		}
		if err := exec(params); err != nil {
			return err
		}
		return nil
	}
}
