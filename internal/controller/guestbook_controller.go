/*
Copyright 2024.

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

package controller

import (
        "context"
        "fmt"
        "k8s.io/apimachinery/pkg/runtime"
        ctrl "sigs.k8s.io/controller-runtime"
        "sigs.k8s.io/controller-runtime/pkg/client"
        "sigs.k8s.io/controller-runtime/pkg/log"
        "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
        logger "k8s.io/klog/v2"
        "errors"

	cloudstackv1 "gitlab.iblog.pro/globus/asura/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const guestbooksFinalizer = "iblog.pro/asura/finalizer"

//+kubebuilder:rbac:groups=cloudstack.iblog.pro,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cloudstack.iblog.pro,resources=guestbooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cloudstack.iblog.pro,resources=guestbooks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Guestbook object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *GuestbookReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

        logger.Info(fmt.Sprintf("POINT -1: EVENT: %#v\n", ctx))
        logger.Info(fmt.Sprintf("POINT 0: EVENT: %#v\n", req))

        obj := &cloudstackv1.Guestbook{}
        if err := r.Get(context.Background(), req.NamespacedName, obj); err != nil {
                logger.Info(fmt.Sprintf("POINT 1: Unable to fetch object: %v", req.NamespacedName))
                return ctrl.Result{}, nil
        } else {
                logger.Info(fmt.Sprintf("POINT 2: Geeting from Kubebuilder to %v", obj))
        }


        if !controllerutil.ContainsFinalizer(obj, guestbooksFinalizer) {
                logger.Info("Adding Finalizer")
                if ok := controllerutil.AddFinalizer(obj, guestbooksFinalizer); !ok {
                        logger.Info("Failed to add finalizer into the custom resource %v", req.NamespacedName)
                        return ctrl.Result{Requeue: true}, nil
                }

                if err := r.Update(ctx, obj); err != nil {
                        logger.Info("Failed to update custom resource to add finalizer :%v", req.NamespacedName)
                        return ctrl.Result{}, err
                }

                obj.Status.Status = "Running"
                if err := r.Status().Update(context.Background(), obj); err != nil {
                        logger.Info(fmt.Sprintf("POINT 3: unable to update status"))
                        return ctrl.Result{}, err
                }
        }

        // Check if the guestbooks instance is marked to be deleted, which is
        // indicated by the deletion timestamp being set.
        isguestbooksMarkedToBeDeleted := obj.GetDeletionTimestamp() != nil
        if isguestbooksMarkedToBeDeleted {

                if controllerutil.ContainsFinalizer(obj, guestbooksFinalizer) {
                        logger.Info("Performing Finalizer Operations for guestbooks before delete CR")

                        obj.Status.Status = "Deleting"
                        if err := r.Status().Update(context.Background(), obj); err != nil {
                                logger.Info(fmt.Sprintf("POINT 55: unable to update status"))
                                return ctrl.Result{}, err
                        }

                        if ok := controllerutil.RemoveFinalizer(obj, guestbooksFinalizer); !ok {
                                logger.Info("Failed to remove finalizer for guestbooks")
                                return ctrl.Result{Requeue: true}, errors.New("Failed to remove finalizer for guestbooks")
                        }

                        if err := r.Update(ctx, obj); err != nil {
                                logger.Info( "Failed to remove finalizer for guestbooks")
                                return ctrl.Result{}, err
                        }

                }

        }


	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudstackv1.Guestbook{}).
		Complete(r)
}
