package main

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func scale(deploymentConfig *Deployment) {
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		log.Panic(err.Error())
	}
	deploymentsClient := clientset.AppsV1().Deployments(deploymentConfig.Namespace)
	s, err := deploymentsClient.GetScale(context.TODO(), deploymentConfig.Name, v1.GetOptions{})
	if err != nil {
		log.Panic(err)
	}
	sc := *s
	sc.Spec.Replicas = int32(deploymentConfig.Replicas)

	_, err = deploymentsClient.UpdateScale(context.TODO(), deploymentConfig.Name, &sc, v1.UpdateOptions{})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("namespace:%s, deployment:%s scaled to %d", deploymentConfig.Namespace, deploymentConfig.Name, deploymentConfig.Replicas)
}
