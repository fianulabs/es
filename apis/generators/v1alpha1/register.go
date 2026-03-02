/*
Copyright © 2025 ESO Maintainer Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "generators.external-secrets.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
	AddToScheme   = SchemeBuilder.AddToScheme
)

var (
	ECRAuthorizationTokenKind = reflect.TypeFor[ECRAuthorizationToken]().Name()
	STSSessionTokenKind       = reflect.TypeFor[STSSessionToken]().Name()
	GCRAccessTokenKind        = reflect.TypeFor[GCRAccessToken]().Name()
	ACRAccessTokenKind        = reflect.TypeFor[ACRAccessToken]().Name()
	PasswordKind              = reflect.TypeFor[Password]().Name()
	SSHKeyKind                = reflect.TypeFor[SSHKey]().Name()
	WebhookKind               = reflect.TypeFor[Webhook]().Name()
	FakeKind                  = reflect.TypeFor[Fake]().Name()
	VaultDynamicSecretKind    = reflect.TypeFor[VaultDynamicSecret]().Name()
	GithubAccessTokenKind     = reflect.TypeFor[GithubAccessToken]().Name()
	QuayAccessTokenKind       = reflect.TypeFor[QuayAccessToken]().Name()
	CloudsmithAccessTokenKind = reflect.TypeFor[CloudsmithAccessToken]().Name()
	UUIDKind                  = reflect.TypeFor[UUID]().Name()
	GrafanaKind               = reflect.TypeFor[Grafana]().Name()
	MFAKind                   = reflect.TypeFor[MFA]().Name()
	ClusterGeneratorKind      = reflect.TypeFor[ClusterGenerator]().Name()
)

func init() {
	SchemeBuilder.Register(&GeneratorState{}, &GeneratorStateList{})

	/*
		===============================================================================
		 NOTE: when adding support for new kinds of generators:
		  1. register the struct types in `SchemeBuilder` (right below this note)
		  2. update the `kubebuilder:validation:Enum` annotation for GeneratorRef.Kind (apis/externalsecrets/v1beta1/externalsecret_types.go)
		  3. add it to the imports of (pkg/generator/register/register.go)
		  4. add it to the ClusterRole called "*-controller" (deploy/charts/external-secrets/templates/rbac.yaml)
		  5. support it in ClusterGenerator:
			  - add a new GeneratorKind enum value (apis/generators/v1alpha1/types_cluster.go)
			  - update the `kubebuilder:validation:Enum` annotation for the GeneratorKind enum
			  - add a spec field to GeneratorSpec (apis/generators/v1alpha1/types_cluster.go)
			  - update the clusterGeneratorToVirtual() function (pkg/utils/resolvers/generator.go)
		===============================================================================
	*/

	SchemeBuilder.Register(&ACRAccessToken{}, &ACRAccessTokenList{})
	SchemeBuilder.Register(&ClusterGenerator{}, &ClusterGeneratorList{})
	SchemeBuilder.Register(&CloudsmithAccessToken{}, &CloudsmithAccessTokenList{})
	SchemeBuilder.Register(&ECRAuthorizationToken{}, &ECRAuthorizationTokenList{})
	SchemeBuilder.Register(&Fake{}, &FakeList{})
	SchemeBuilder.Register(&GCRAccessToken{}, &GCRAccessTokenList{})
	SchemeBuilder.Register(&GithubAccessToken{}, &GithubAccessTokenList{})
	SchemeBuilder.Register(&QuayAccessToken{}, &QuayAccessTokenList{})
	SchemeBuilder.Register(&Password{}, &PasswordList{})
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
	SchemeBuilder.Register(&STSSessionToken{}, &STSSessionTokenList{})
	SchemeBuilder.Register(&UUID{}, &UUIDList{})
	SchemeBuilder.Register(&VaultDynamicSecret{}, &VaultDynamicSecretList{})
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
	SchemeBuilder.Register(&Grafana{}, &GrafanaList{})
	SchemeBuilder.Register(&MFA{}, &MFAList{})
}
