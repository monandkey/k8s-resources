#!/usr/bin/bash

for NAMESPACES in $(oc get namespaces | grep -v NAME | awk '{print $1}'); do
    for POD in $(oc get pod -n $NAMESPACES | grep -v NAME | awk '{print $1}'); do
        CONTAINER=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.containers[*].name}"`
        ICONTAINER=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.initContainers[*].name}"`
        CONTAINER_NUM=($CONTAINER)
        ICONTAINER_NUM=($ICONTAINER)
        for i in (0..CONTAINER_NUM); do
            REQUEST_CPU=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.contaienrs[i].resources.requests.cpu}"`
            REQUEST_MEM=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.contaienrs[i].resources.requests.memory}"`
            LIMIT_CPU=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.contaienrs[i].resources.limits1.cpu}"`
            LIMIT_MEM=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.contaienrs[i].resources.limits.memory}"`
            echo $POD,${CONTAINER[i]},$REQUEST_CPU,$REQUEST_MEM,$LIMIT_CPU,$LIMIT_MEM
        done
        for i in (0..ICONTAINER_NUM); do
            REQUEST_CPU=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.initContainers[i].resources.requests.cpu}"`
            REQUEST_MEM=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.initContainers[i].resources.requests.memory}"`
            LIMIT_CPU=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.initContainers[i].resources.limits1.cpu}"`
            LIMIT_MEM=`oc get pods -n $NAMESPACES $POD -o jsonpath="{.spec.initContainers[i].resources.limits.memory}"`
            echo $POD,${ICONTAINER[i]},$REQUEST_CPU,$REQUEST_MEM,$LIMIT_CPU,$LIMIT_MEM
        done
    done
done
