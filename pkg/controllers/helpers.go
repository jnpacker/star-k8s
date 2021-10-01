package controllers

import (
	"context"

	starK8s "github.com/open-cluster-management/star-k8s/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/klog"
)

// Logging types
const ERROR = -2
const WARN = -1
const INFO = 0
const DEBUG = 1

func recordStatusCondition(
	r *StarK8sDeploymentReconciler,
	sd *starK8s.StarK8sDeployment,
	newCondition v1.ComponentCondition) error {

	/*var newCondition = v1.ComponentCondition{
		Type:    containerName,
		Status:  conditionStatus,
		Error:  reason,
		Message: message,
	}*/

	meta.SetStatusCondition(&sd.Status.Conditions, newCondition)

	if err := r.Update(context.TODO(), sd); err != nil {
		return err
	}
	klog.V(4).Infof("newCondition: %v", newCondition)
	return nil
}
