#!/usr/bin/bash

# Specifies the command to be executed. It is assumed that kubectl or oc is entered here.
k8sCmd=kubectl

# Define the resources you want to check as an array.
k8sResource=(
  "deployment"
  "statefulset"
  "daemonset"
  "cronjob"
  "pod"
)

# Define offsets to containers and InitContainers.
k8sOffset=(
  "spec.template.spec"                  # deployment
  "spec.template.spec"                  # statefulset
  "spec.template.spec"                  # daemonset
  "spec.jobTemplate.spec.template.spec" # cronjob
  "spec"                                # pod
)

# If a flag of 1 is entered, the output of that resource is skipped.
k8sResourceSkipFlag=(
  0 # deployment
  0 # statefulset
  0 # daemonset
  0 # cronjob
  1 # pod
)

# Split up the command because it is too long to describe it normally in the command.
k8sReqCPU="resources.requests.cpu"
k8sReqMem="resources.requests.memory"
k8sLimCPU="resources.limits.cpu"
k8sLimMem="resources.limits.memory"

# When this flag is set to 1, headers are output. Conversely,
# setting this flag to 0 skips the process of outputting headers.
HEADER_FLAG=1

# Specify the path to the file that outputs the results.
RESULT_FILE=/tmp/resources_check.csv


# This tool is designed to be run with old files always deleted.
# Care is taken to ensure that the results do not take an unintended form.
function cleanUpResultfile () {
  ls $RESULT_FILE > /dev/null 2>&1
  CMD_RESULT=$?
  if [ $CMD_RESULT -eq 0 ]; then
    rm -i $RESULT_FILE && echo "Successful a file deleted" || exit 1
  fi
}


# Output headers for CSV. HEADER_FLAG is used to determine whether to output.
function headerOutput () {
  if [ $HEADER_FLAG -eq 1 ]; then
    echo NAMESPACES,RESOURCES,METADATA_NAME,CONTAINER,REQUEST_CPU,REQUEST_MEM,LIMIT_CPU,LIMIT_MEM | tee $RESULT_FILE
    HEADER_FLAG=0
  fi
}


# Logs to console and file.
function resultOutput () {
  ns=$1
  resources=$2
  metadataName=$3
  conList=($4)
  reqCPUList=($5)
  reqMemList=($6)
  limCPUList=($7)
  limMemList=($8)

  for cnum in `seq 0 $((${#conList[*]}-1))`; do
    echo $ns,$resources,$metadataName,${conList[$cnum]},${reqCPUList[$cnum]},${reqMemList[$cnum]},${limCPUList[$cnum]},${limMemList[$cnum]} | tee -a $RESULT_FILE
  done
}


# Retrieve the resources declared in the received target on a per-container basis.
function retrievePerContainerInfo () {
  ns=$1
  rnum=$2
  metadataName=$3

  # Retrieve the container name of the container and initContainer.
  conList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.containers[*].name}"`)
  initConList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.initContainers[*].name}"`)

  if [ $((${#conList[*]})) -ne 0 ]; then
    conReqCPUList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.containers[*].$k8sReqCPU}"`)
    conReqMemList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.containers[*].$k8sReqMem}"`)
    conLimCPUList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.containers[*].$k8sLimCPU}"`)
    conLimMemList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.containers[*].$k8sLimMem}"`)
    resultOutput $ns ${k8sResource[$rnum]} $metadataName "${conList[*]}" "${conReqCPUList[*]}" "${conReqMemList[*]}" "${conLimCPUList[*]}" "${conLimMemList[*]}"
  fi

  if [ $((${#initConList[*]})) -ne 0 ]; then
    iconReqCPUList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.initContainers[*].$k8sReqCPU}"`)
    iconReqMemList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.initContainers[*].$k8sReqMem}"`)
    iconLimCPUList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.initContainers[*].$k8sLimCPU}"`)
    iconLimMemList=(`$k8sCmd get ${k8sResource[$rnum]} -n $ns $metadataName -o jsonpath="{.${k8sOffset[$rnum]}.initContainers[*].$k8sLimMem}"`)
    resultOutput $ns ${k8sResource[$rnum]} $metadataName "${iconList[*]}" "${iconReqCPUList[*]}" "${iconReqMemList[*]}" "${iconLimCPUList[*]}" "${iconLimMemList[*]}"
  fi
}


# Retrieve the namespaces declared in the received target on a per-resources basis.
function retrievePerResource () {
  ns=$1

  for rnum in `seq 0 $((${#k8sResource[*]}-1))`; do
    if [ ${k8sResourceSkipFlag[$rnum]} -eq 1 ]; then
      continue
    fi

    for metadataName in $($k8sCmd get ${k8sResource[$rnum]} -n $ns 2> /dev/null | grep -v NAME | awk '{print $1}'); do
      retrievePerContainerInfo $ns $rnum $metadataName
    done
  done
}


# The namespaces in the cluster are retrieved and processed one by one.
function retrievePerNamespaces () {
  for ns in $($k8sCmd get namespaces | grep -v NAME | awk '{print $1}'); do
    retrievePerResource $ns
  done
}


function main () {
  cleanUpResultfile
  headerOutput
  retrievePerNamespaces
}


main
