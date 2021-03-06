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

package main

import (
	"flag"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/signals"

	"github.com/google/knative-gcp/pkg/decorator"
)

func main() {
	flag.Parse()

	sl, _ := logging.NewLogger("", "INFO") // TODO: use logging config map.
	logger := sl.Desugar()
	defer logger.Sync()

	ctx := logging.WithLogger(signals.NewContext(), logger.Sugar())

	dec := decorator.Decorator{}
	if err := envconfig.Process("", &dec); err != nil {
		logger.Fatal("Failed to process env var", zap.Error(err))
	}
	if err := dec.Start(ctx); err != nil {
		logger.Fatal("Failed to run Decorator.", zap.Error(err))
	}
}
