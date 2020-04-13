/*
Copyright 2020 Google LLC

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"knative.dev/eventing/pkg/apis/duck"
	"knative.dev/pkg/apis"
)

var brokerCondSet = apis.NewLivingConditionSet(
	BrokerConditionIngress,
	BrokerConditionAddressable,

	BrokerConditionTopic,

	BrokerConditionSubscription,
)

const (
	BrokerConditionReady = apis.ConditionReady
	// BrokerConditionIngress reports the availability of the Broker's ingress
	// service.
	BrokerConditionIngress apis.ConditionType = "IngressReady"
	// BrokerConditionAddressable reports the status of the Broker's ingress
	// address.
	BrokerConditionAddressable apis.ConditionType = "Addressable"
	// BrokerConditionTopic reports the status of the Broker's PubSub topic.
	// THis condition is specific to the Google Cloud Broker.
	BrokerConditionTopic apis.ConditionType = "TopicReady"
	// BrokerConditionSubscription reports the status of the Broker's PubSub
	// subscription. This condition is specific to the Google Cloud Broker.
	BrokerConditionSubscription apis.ConditionType = "SubscriptionReady"
)

// GetCondition returns the condition currently associated with the given type, or nil.
func (bs *BrokerStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return brokerCondSet.Manage(bs).GetCondition(t)
}

// GetTopLevelCondition returns the top level Condition.
func (bs *BrokerStatus) GetTopLevelCondition() *apis.Condition {
	return brokerCondSet.Manage(bs).GetTopLevelCondition()
}

// IsReady returns true if the resource is ready overall.
func (bs *BrokerStatus) IsReady() bool {
	return brokerCondSet.Manage(bs).IsHappy()
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (bs *BrokerStatus) InitializeConditions() {
	brokerCondSet.Manage(bs).InitializeConditions()
}

func (bs *BrokerStatus) PropagateIngressAvailability(ep *corev1.Endpoints) {
	if duck.EndpointsAreAvailable(ep) {
		brokerCondSet.Manage(bs).MarkTrue(BrokerConditionIngress)
	} else {
		bs.MarkIngressFailed("EndpointsUnavailable", "Endpoints %q are unavailable.", ep.Name)
	}
}

// SetAddress makes this Broker addressable by setting the hostname. It also
// sets the BrokerConditionAddressable to true.
func (bs *BrokerStatus) SetAddress(url *apis.URL) {
	bs.Address.URL = url
	if url != nil {
		brokerCondSet.Manage(bs).MarkTrue(BrokerConditionAddressable)
	} else {
		brokerCondSet.Manage(bs).MarkFalse(BrokerConditionAddressable, "emptyURL", "URL is empty")
	}
}

func (bs *BrokerStatus) MarkTopicFailed(reason, format string, args ...interface{}) {
	brokerCondSet.Manage(bs).MarkFalse(BrokerConditionTopic, reason, format, args...)
}

func (bs *BrokerStatus) MarkTopicReady() {
	brokerCondSet.Manage(bs).MarkTrue(BrokerConditionTopic)
}

func (bs *BrokerStatus) MarkSubscriptionFailed(reason, format string, args ...interface{}) {
	brokerCondSet.Manage(bs).MarkFalse(BrokerConditionSubscription, reason, format, args...)
}

func (bs *BrokerStatus) MarkSubscriptionReady() {
	brokerCondSet.Manage(bs).MarkTrue(BrokerConditionSubscription)
}

func (bs *BrokerStatus) MarkIngressFailed(reason, format string, args ...interface{}) {
	brokerCondSet.Manage(bs).MarkFalse(BrokerConditionIngress, reason, format, args...)
}
