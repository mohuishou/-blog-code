/*
Copyright 2021.

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

package v1

import (
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var nodepoollog = logf.Log.WithName("nodepool-resource")

func (r *NodePool) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-nodes-lailin-xyz-v1-nodepool,mutating=true,failurePolicy=fail,sideEffects=None,groups=nodes.lailin.xyz,resources=nodepools,verbs=create;update,versions=v1,name=mnodepool.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &NodePool{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *NodePool) Default() {
	nodepoollog.Info("default", "name", r.Name)

	// 默认 handler 为 runc
	r.Spec.Handler = "runc"
}

// change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-nodes-lailin-xyz-v1-nodepool,mutating=false,failurePolicy=fail,sideEffects=None,groups=nodes.lailin.xyz,resources=nodepools,verbs=create;update,versions=v1,name=vnodepool.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &NodePool{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *NodePool) ValidateCreate() error {
	nodepoollog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *NodePool) ValidateUpdate(old runtime.Object) error {
	nodepoollog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *NodePool) ValidateDelete() error {
	nodepoollog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// validate 验证
func (r *NodePool) validate() error {
	reg := regexp.MustCompile(`^node-pool.lailin.xyz/*[a-zA-z0-9]*$`)

	for k := range r.Labels {
		if !reg.MatchString(k) {
		}
	}

	return nil
}
