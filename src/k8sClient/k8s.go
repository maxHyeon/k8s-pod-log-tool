package main

import (
  "context"
  "fmt"
  "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
)

func main() {
  // kubeconfig에서 현재 콘텍스트를 사용한다
  // path-to-kubeconfig -- 예를 들어, /root/.kube/config
  config, _ := clientcmd.BuildConfigFromFlags("", "<path-to-kubeconfig>")
  // clientset을 생성한다
  clientset, _ := kubernetes.NewForConfig(config)
  // 파드를 나열하기 위해 API에 접근한다
  pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
  fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}