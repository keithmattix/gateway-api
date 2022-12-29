// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deployment

import (
	"fmt"

	"istio.io/istio/pkg/test/framework/components/echo/deployment"
	"istio.io/istio/pkg/test/framework/components/echo/match"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"sigs.k8s.io/gateway-api/conformance/mesh/echo"
	"sigs.k8s.io/gateway-api/conformance/mesh/resource"
)

const (
	ASvc           = "a"
	BSvc           = "b"
	CSvc           = "c"
	HeadlessSvc    = "headless"
	StatefulSetSvc = "statefulset"
)

// EchoNamespace contains the echo instances for a single namespace.
type EchoNamespace struct {
	// Namespace where the services are deployed.
	Namespace namespace.Instance
	// Standard echo app to be used by tests
	A echo.Instances
	// Standard echo app to be used by tests
	B echo.Instances
	// Standard echo app to be used by tests
	C echo.Instances
	// Headless echo app to be used by tests
	Headless echo.Instances
	// StatefulSet echo app to be used by tests
	StatefulSet echo.Instances
	// All echo apps in this namespace
	All echo.Services
}

func (n EchoNamespace) build(b deployment.Builder, cfg Config) deployment.Builder {
	for _, config := range cfg.Configs.Get() {
		if config.Namespace == nil {
			config.Namespace = n.Namespace
		}
		if config.Namespace.Name() == n.Namespace.Name() {
			b = b.WithConfig(config)
		}
	}
	return b
}

func (n *EchoNamespace) loadValues(t resource.Context, echos echo.Instances, d *Echos) error {
	ns := n.Namespace
	n.All = match.Namespace(ns).GetMatches(echos).Services()

	n.A = match.ServiceName(echo.NamespacedName{Name: ASvc, Namespace: ns}).GetMatches(echos)
	n.B = match.ServiceName(echo.NamespacedName{Name: BSvc, Namespace: ns}).GetMatches(echos)
	n.C = match.ServiceName(echo.NamespacedName{Name: CSvc, Namespace: ns}).GetMatches(echos)
	n.Headless = match.ServiceName(echo.NamespacedName{Name: HeadlessSvc, Namespace: ns}).GetMatches(echos)
	n.StatefulSet = match.ServiceName(echo.NamespacedName{Name: StatefulSetSvc, Namespace: ns}).GetMatches(echos)

	namespaces, err := namespace.GetAll(t)
	if err != nil {
		return fmt.Errorf("failed retrieving list of namespaces: %v", err)
	}

	return nil
}
