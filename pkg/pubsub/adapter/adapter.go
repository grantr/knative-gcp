/*
Copyright 2019 Google LLC

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

package adapter

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/http"
	cepubsub "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	"go.uber.org/zap"
	"knative.dev/pkg/logging"

	"github.com/google/knative-gcp/pkg/kncloudevents"
	"github.com/google/knative-gcp/pkg/pubsub/adapter/converters"
)

// Adapter implements the Pub/Sub adapter to deliver Pub/Sub messages from a
// pre-existing topic/subscription to a Sink.
type Adapter struct {
	// Environment variable containing project id.
	Project string `envconfig:"PROJECT_ID"`

	// Environment variable containing the sink URI.
	Sink string `envconfig:"SINK_URI" required:"true"`

	// Environment variable containing the transformer URI.
	Transformer string `envconfig:"TRANSFORMER_URI"`

	// Topic is the environment variable containing the PubSub Topic being
	// subscribed to's name. In the form that is unique within the project.
	// E.g. 'laconia', not 'projects/my-gcp-project/topics/laconia'.
	Topic string `envconfig:"PUBSUB_TOPIC_ID" required:"true"`

	// Subscription is the environment variable containing the name of the
	// subscription to use.
	Subscription string `envconfig:"PUBSUB_SUBSCRIPTION_ID" required:"true"`

	// SendMode describes how the adapter sends events.
	// One of [binary, structured, push]. Default: binary
	SendMode converters.ModeType `envconfig:"SEND_MODE" default:"binary" required:"true"`

	// inbound is the cloudevents client to use to receive events.
	inbound cloudevents.Client

	// outbound is the cloudevents client to use to send events.
	outbound cloudevents.Client

	// transformer is the cloudevents client to transform received events before sending.
	transformer cloudevents.Client
}

// Start starts the adapter. Note: Only call once, not thread safe.
func (a *Adapter) Start(ctx context.Context) error {
	var err error

	if a.SendMode == "" {
		a.SendMode = converters.DefaultSendMode
	}

	// Receive Events on Pub/Sub.
	if a.inbound == nil {
		if a.inbound, err = a.newPubSubClient(ctx); err != nil {
			return fmt.Errorf("failed to create inbound cloudevent client: %s", err.Error())
		}
	}

	// Send events on HTTP.
	if a.outbound == nil {
		if a.outbound, err = a.newHTTPClient(a.Sink); err != nil {
			return fmt.Errorf("failed to create outbound cloudevent client: %s", err.Error())
		}
	}

	// Make the transformer client in case the TransformerURI has been set.
	if a.Transformer != "" {
		if a.transformer == nil {
			if a.transformer, err = kncloudevents.NewDefaultClient(a.Transformer); err != nil {
				return fmt.Errorf("failed to create transformer cloudevent client: %s", err.Error())
			}
		}
	}

	return a.inbound.StartReceiver(ctx, a.receive)
}

func (a *Adapter) receive(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	logger := logging.FromContext(ctx).With(zap.Any("event.id", event.ID()), zap.Any("sink", a.Sink))

	// If a transformer has been configured, then transform the message.
	if a.transformer != nil {
		// TODO: I do not like the transformer as it is. It would be better to pass the transport context and the
		// message to the transformer function as a transform request. Better yet, only do it for conversion issues?
		transformedEvent, err := a.transformer.Send(ctx, event)
		if err != nil {
			logger.Errorf("error transforming cloud event %q", event.ID())
			return err
		}
		if transformedEvent == nil {
			logger.Warnf("cloud event %q was not transformed", event.ID())
			return nil
		}
		// Update the event with the transformed one.
		event = *transformedEvent
	}

	// If send mode is Push, convert to Pub/Sub Push payload style.
	if a.SendMode == converters.Push {
		event = ConvertToPush(ctx, event)
	}

	if r, err := a.outbound.Send(ctx, event); err != nil {
		return err
	} else if r != nil {
		resp.RespondWith(200, r)
	}

	return nil
}

func (a *Adapter) convert(ctx context.Context, m transport.Message, err error) (*cloudevents.Event, error) {
	logger := logging.FromContext(ctx)
	logger.Debug("Converting event from transport.")

	if msg, ok := m.(*cepubsub.Message); ok {
		return converters.Convert(ctx, msg, a.SendMode)
	}
	return nil, err
}

func (a *Adapter) newPubSubClient(ctx context.Context) (cloudevents.Client, error) {
	tOpts := []cepubsub.Option{
		cepubsub.WithProjectID(a.Project),
		cepubsub.WithTopicID(a.Topic),
		cepubsub.WithSubscriptionID(a.Subscription),
	}

	// Make a pubsub transport for the CloudEvents client.
	t, err := cepubsub.New(ctx, tOpts...)
	if err != nil {
		return nil, err
	}

	// Use the transport to make a new CloudEvents client.
	return cloudevents.NewClient(t,
		cloudevents.WithConverterFn(a.convert),
	)
}

func (a *Adapter) newHTTPClient(target string) (cloudevents.Client, error) {
	tOpts := []http.Option{
		cloudevents.WithTarget(target),
	}

	switch a.SendMode {
	case converters.Binary, converters.Push:
		tOpts = append(tOpts, cloudevents.WithBinaryEncoding())
	case converters.Structured:
		tOpts = append(tOpts, cloudevents.WithStructuredEncoding())
	}

	// Make an http transport for the CloudEvents client.
	t, err := cloudevents.NewHTTPTransport(tOpts...)
	if err != nil {
		return nil, err
	}

	// Use the transport to make a new CloudEvents client.
	return cloudevents.NewClient(t)
}
