package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func create(x int64) *int64 {
	return &x
}
func main() {
	// todo
	// 1. get imput user's kubeconfig
	// 2. if no input automatically use user's home/.kube/config
	config, _ := clientcmd.BuildConfigFromFlags("https://hmsptest01:6443", "/Users/parkmh90/.kube/config")
	clientset, _ := kubernetes.NewForConfig(config)

	timeoutTime := create(10)
	// 파드를 나열하기 위해 API에 접근한다
	deployments, _ := clientset.AppsV1().Deployments("hmsp").List(context.TODO(), v1.ListOptions{})
	pods, _ := clientset.CoreV1().Pods("hmsp").List(context.TODO(), v1.ListOptions{TimeoutSeconds: timeoutTime})
	fmt.Printf("There are %d deployments in the cluster\n", len(deployments.Items))
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for dep, name := range deployments.Items {

	}

}
