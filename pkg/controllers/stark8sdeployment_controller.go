/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	starK8s "github.com/open-cluster-management/star-k8s/api/v1alpha1"
)

// StarK8sDeploymentReconciler reconciles a StarK8sDeployment object
type StarK8sDeploymentReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=stark8sdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cluster.open-cluster-management.io,resources=stark8sdeployments/status,verbs=get;update;patch

func (r *StarK8sDeploymentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("stark8sdeployment", req.NamespacedName)

	// your logic here
	var sd starK8s.StarK8sDeployment
	if err := r.Get(ctx, req.NamespacedName, &sd); err != nil {
		log.V(ERROR).Info("Did not find starK8sDeployment resource by the name of " + req.Name)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, reconcileAks(r, &sd)
	/*if sd.Spec.Platform.Eks != nil {
		return ctrl.Result{}, reconcileAks(r, &sd)
	}

	return ctrl.Result{}, nil*/
}

func (r *StarK8sDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&starK8s.StarK8sDeployment{}).
		Complete(r)
}
