package main

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func scale(deploymentConfig *Deployment) {
	// creates the in-cluster config
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		panic(err.Error())
	}
	deploymentsClient := clientset.AppsV1().Deployments(deploymentConfig.Namespace)
	s, err := deploymentsClient.GetScale(context.TODO(), deploymentConfig.Name, v1.GetOptions{})
	if err != nil {
		log.Fatal(err)
		return
	}
	sc := *s
	sc.Spec.Replicas = int32(deploymentConfig.Replicas)

	us, err := deploymentsClient.GetScale(context.TODO(), deploymentConfig.Name, v1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Panicln(us)
}
