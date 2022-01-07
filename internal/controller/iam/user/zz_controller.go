/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package user

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	tjcontroller "github.com/crossplane/terrajet/pkg/controller"
	"github.com/crossplane/terrajet/pkg/terraform"
	ctrl "sigs.k8s.io/controller-runtime"

	v1alpha2 "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2"
)

// Setup adds a controller that reconciles User managed resources.
func Setup(mgr ctrl.Manager, o controller.Options, s terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider) error {
	name := managed.ControllerName(v1alpha2.User_GroupVersionKind.String())
	var initializers managed.InitializerChain
	for _, i := range cfg.Resources["aws_iam_user"].InitializerFns {
		initializers = append(initializers, i(mgr.GetClient()))
	}
	initializers = append(initializers, managed.NewNameAsExternalName(mgr.GetClient()))
	r := managed.NewReconciler(mgr,
		xpresource.ManagedKind(v1alpha2.User_GroupVersionKind),
		managed.WithExternalConnecter(tjcontroller.NewConnector(mgr.GetClient(), ws, s, cfg.Resources["aws_iam_user"])),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(ws, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3*time.Minute),
		managed.WithInitializers(initializers),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha2.User{}).
		Complete(r)
}
