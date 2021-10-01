package controllers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	starK8s "github.com/open-cluster-management/star-k8s/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

func reconcileAks(r *StarK8sDeploymentReconciler, sd *starK8s.StarK8sDeployment) error {
	ctx := context.Background()
	typesNamespacedName := types.NamespacedName{Namespace: sd.Namespace, Name: sd.Spec.ProviderSecretRef.Name}
	log := r.Log.WithValues("stark8sdeployment", typesNamespacedName)

	var providerSecret v1.Secret

	if err := r.Get(ctx, typesNamespacedName, &providerSecret); err != nil {
		log.Error(err, "Could not retreive providerSecretRef")
		return err
	}

	defaultClusterList := eks.ListClustersInput{}
	sess := session.Must(session.NewSession())

	svc := eks.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	list, err := svc.ListClusters(&defaultClusterList)
	if err != nil {
		log.Error(err, "Could not list EKS clusters")
		return err
	}
	fmt.Printf("%v", list)

	return nil
}
