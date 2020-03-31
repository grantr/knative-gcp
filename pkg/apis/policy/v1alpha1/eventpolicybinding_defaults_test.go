/*
Copyright 2020 Google LLC.

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

package v1alpha1

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

func TestEventPolicyBindingDefaults(t *testing.T) {
	pb := EventPolicyBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-policy-binding",
			Namespace: "test-namespace",
		},
		Spec: PolicyBindingSpec{Policy: duckv1.KReference{}},
	}
	pb.SetDefaults(context.Background())
	if pb.Spec.Subject.Namespace != "test-namespace" {
		t.Errorf("spec.Subject.Namespace got=%s want=test-namespace", pb.Spec.Subject.Namespace)
	}
	if pb.Spec.Policy.Namespace != "test-namespace" {
		t.Errorf("spec.Policy.Namespace got=%s want=test-namespace", pb.Spec.Policy.Namespace)
	}
}
