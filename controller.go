package main

func exec(params params) error {
	if params.resource == "pods" || params.resource == "pod" || params.resource == "po" {
		return execPod(params)
	} else if params.resource == "replicaset" || params.resource == "rs" {
		return execReplicaSet(params)
	} else if params.resource == "deployments" || params.resource == "deploy" {
		return execDeployment(params)
	} else if params.resource == "statefulsets" || params.resource == "sts" {
		return execStatefulSet(params)
	} else if params.resource == "daemonsets" || params.resource == "ds" {
		return execDaemonSet(params)
	} else if params.resource == "jobs" || params.resource == "job" {
		return execJob(params)
	} else if params.resource == "cronjobs" || params.resource == "cronjob" || params.resource == "cj" {
		return execCronJob(params)
	} else if params.allNamespace {
		return errOptionNotSupport
	}
	return errResourceNotSupport
}

func execPod(params params) error {
	podList, err := getPodList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayPodList(podList)
}

func execReplicaSet(params params) error {
	rsList, err := getReplicaSetList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayReplicaSetList(rsList)
}

func execDeployment(params params) error {
	deployList, err := getDeploymentList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayDeploymentList(deployList)
}

func execStatefulSet(params params) error {
	stsList, err := getStatefulSetList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayStatefulSetList(stsList)
}

func execDaemonSet(params params) error {
	dsList, err := getDaemonSetList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayDaemonSetList(dsList)
}

func execJob(params params) error {
	jobList, err := getJobList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayJobList(jobList)
}

func execCronJob(params params) error {
	cjList, err := getCronJobList(params.kubeconfig, params.namespace)
	if err != nil {
		return err
	}
	return displayCronJobList(cjList)
}
