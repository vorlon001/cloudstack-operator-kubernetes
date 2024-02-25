# guestbook
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.21.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/guestbook:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/guestbook:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/guestbook:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/guestbook/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

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






### README
```

# download kubebuilder and install locally.
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/


mkdir -p ~/projects/cloudstack
cd ~/projects/cloudstack
kubebuilder init --domain iblog.pro --license apache2 --repo gitlab.iblog.pro/globus/asura

>>> INFO Writing kustomize manifests for you to edit...
>>> INFO Writing scaffold for you to edit...
>>> INFO Get controller runtime:
>>> $ go get sigs.k8s.io/controller-runtime@v0.17.0
>>> INFO Update dependencies:
>>> $ go mod tidy
>>> Next: define a resource with:
>>> $ kubebuilder create api


kubebuilder create api --group cloudstack --version v1 --kind Guestbook

>>> INFO Create Resource [y/n]
>>> y
>>> INFO Create Controller [y/n]
>>> y
>>> INFO Writing kustomize manifests for you to edit...
>>> INFO Writing scaffold for you to edit...
>>> INFO api/v1/guestbook_types.go
>>> INFO api/v1/groupversion_info.go
>>> INFO internal/controller/suite_test.go
>>> INFO internal/controller/guestbook_controller.go
>>> INFO internal/controller/guestbook_controller_test.go
>>> INFO Update dependencies:
>>> $ go mod tidy
>>> INFO Running make:
>>> $ make generate
>>> mkdir -p /root/projects/guestbook/bin
>>> Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0
>>> /root/projects/guestbook/bin/controller-gen-v0.14.0 object:headerFile="hack/boilerplate.go.txt" paths="./..."
>>> Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
>>> $ make manifests


make manifests

make build-installer IMG=harbor.iblog.pro/test/asura:latest

make docker-build IMG=harbor.iblog.pro/test/asura:latest
make docker-push IMG=harbor.iblog.pro/test/asura:latest
make deploy IMG=harbor.iblog.pro/test/asura:latest

kubectl create ns cloudstack-system
kubectl apply -f install.yaml

root@node180:~# kubectl apply -f cloudstack_v1_guestbook.yaml
guestbook.cloudstack.iblog.pro/guestbook-sample created

kubectl apply -f cloudstack_v1_guestbook.yaml -n cloudstack-system


cat <<EOF>cloudstack_v1_guestbook.yaml
apiVersion: cloudstack.iblog.pro/v1
kind: Guestbook
metadata:
  labels:
    app.kubernetes.io/name: guestbook
    app.kubernetes.io/instance: guestbook-sample
    app.kubernetes.io/part-of: guestbook
    app.kubernetes.io/managed-by: kustomize
    app.kubernetes.io/created-by: guestbook
  name: guestbook-sample
spec:
  foo: foo test Guestbook
  firstname: firstname test Guestbook
  lastname: lastname test Guestbook
EOF

kubectl apply -f cloudstack_v1_guestbook.yaml -n guestbook-system
kubectl delete -f cloudstack_v1_guestbook.yaml -n guestbook-system

root@node180:~# kubectl logs -n cloudstack-system   pod/cloudstack-controller-manager-c7fc6d9c-h28zx
2024-02-25T07:07:12Z    INFO    setup   starting manager
2024-02-25T07:07:12Z    INFO    controller-runtime.metrics      Starting metrics server
2024-02-25T07:07:12Z    INFO    starting server {"kind": "health probe", "addr": "[::]:8081"}
2024-02-25T07:07:12Z    INFO    controller-runtime.metrics      Serving metrics server  {"bindAddress": "127.0.0.1:8080", "secure": false}
I0225 07:07:12.754112       1 leaderelection.go:250] attempting to acquire leader lease cloudstack-system/44ff3e94.iblog.pro...
I0225 07:07:12.776459       1 leaderelection.go:260] successfully acquired lease cloudstack-system/44ff3e94.iblog.pro
2024-02-25T07:07:12Z    DEBUG   events  cloudstack-controller-manager-c7fc6d9c-h28zx_73adb361-c584-4497-a152-5cf86aadf9af became leader {"type": "Normal", "object": {"kind":"Lease","namespace":"cloudstack-system","name":"44ff3e94.iblog.pro","uid":"e95f7df0-081e-4e8b-91cc-f93f82d68944","apiVersion":"coordination.k8s.io/v1","resourceVersion":"8380"}, "reason": "LeaderElection"}
2024-02-25T07:07:12Z    INFO    Starting EventSource    {"controller": "guestbook", "controllerGroup": "cloudstack.iblog.pro", "controllerKind": "Guestbook", "source": "kind source: *v1.Guestbook"}
2024-02-25T07:07:12Z    INFO    Starting Controller     {"controller": "guestbook", "controllerGroup": "cloudstack.iblog.pro", "controllerKind": "Guestbook"}
2024-02-25T07:07:12Z    INFO    Starting workers        {"controller": "guestbook", "controllerGroup": "cloudstack.iblog.pro", "controllerKind": "Guestbook", "worker count": 1}
root@node180:~#


root@node180:~# kubectl logs -n cloudstack-system   pod/cloudstack-controller-manager-c7fc6d9c-h28zx
.........
I0225 07:08:33.665015       1 guestbook_controller.go:60] POINT 0: EVENT: reconcile.Request{NamespacedName:types.NamespacedName{Namespace:"default", Name:"guestbook-sample"}}
I0225 07:08:33.665136       1 guestbook_controller.go:67] POINT 2: Geeting from Kubebuilder to &{{Guestbook cloudstack.iblog.pro/v1} {guestbook-sample  default  af927179-7d9f-43fb-a62d-8f426bd38c9b 8655 2 2024-02-25 07:08:24 +0000 UTC 2024-02-25 07:08:33 +0000 UTC 0xc000578800 map[app.kubernetes.io/created-by:guestbook app.kubernetes.io/instance:guestbook-sample app.kubernetes.io/managed-by:kustomize app.kubernetes.io/name:guestbook app.kubernetes.io/part-of:guestbook] map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"cloudstack.iblog.pro/v1","kind":"Guestbook","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"guestbook","app.kubernetes.io/instance":"guestbook-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"guestbook","app.kubernetes.io/part-of":"guestbook"},"name":"guestbook-sample","namespace":"default"},"spec":{"firstname":"firstname test Guestbook","foo":"foo test Guestbook","lastname":"lastname test Guestbook"}}
] [] [cloudstack.iblog.pro/finalizer] [{kubectl-client-side-apply Update cloudstack.iblog.pro/v1 2024-02-25 07:08:24 +0000 UTC FieldsV1 {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app.kubernetes.io/created-by":{},"f:app.kubernetes.io/instance":{},"f:app.kubernetes.io/managed-by":{},"f:app.kubernetes.io/name":{},"f:app.kubernetes.io/part-of":{}}},"f:spec":{".":{},"f:firstname":{},"f:foo":{},"f:lastname":{}}} } {manager Update cloudstack.iblog.pro/v1 2024-02-25 07:08:24 +0000 UTC FieldsV1 {"f:metadata":{"f:finalizers":{".":{},"v:\"cloudstack.iblog.pro/finalizer\"":{}}}} } {manager Update cloudstack.iblog.pro/v1 2024-02-25 07:08:24 +0000 UTC FieldsV1 {"f:status":{".":{},"f:Status":{}}} status}]} {foo test Guestbook firstname test Guestbook lastname test Guestbook} {Running []}}
I0225 07:08:33.665164       1 guestbook_controller.go:96] Performing Finalizer Operations for guestbooks before delete CR
I0225 07:08:33.678535       1 guestbook_controller.go:59] POINT -1: EVENT: &context.valueCtx{Context:(*context.valueCtx)(0xc000445e90), key:controller.reconcileIDKey{}, val:"059c554b-860d-4f70-bfdb-2203c5c7065f"}
I0225 07:08:33.678578       1 guestbook_controller.go:60] POINT 0: EVENT: reconcile.Request{NamespacedName:types.NamespacedName{Namespace:"default", Name:"guestbook-sample"}}
I0225 07:08:33.678639       1 guestbook_controller.go:64] POINT 1: Unable to fetch object: default/guestbook-sample
I0225 07:09:28.663081       1 guestbook_controller.go:59] POINT -1: EVENT: &context.valueCtx{Context:(*context.valueCtx)(0xc000435470), key:controller.reconcileIDKey{}, val:"8d846eda-8145-41f1-9716-b7282659ceef"}
I0225 07:09:28.663119       1 guestbook_controller.go:60] POINT 0: EVENT: reconcile.Request{NamespacedName:types.NamespacedName{Namespace:"default", Name:"guestbook-sample"}}
I0225 07:09:28.663241       1 guestbook_controller.go:67] POINT 2: Geeting from Kubebuilder to &{{Guestbook cloudstack.iblog.pro/v1} {guestbook-sample  default  0c29ddd7-4223-48ae-bfab-3456d391c1a3 8837 1 2024-02-25 07:09:28 +0000 UTC <nil> <nil> map[app.kubernetes.io/created-by:guestbook app.kubernetes.io/instance:guestbook-sample app.kubernetes.io/managed-by:kustomize app.kubernetes.io/name:guestbook app.kubernetes.io/part-of:guestbook] map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"cloudstack.iblog.pro/v1","kind":"Guestbook","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"guestbook","app.kubernetes.io/instance":"guestbook-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"guestbook","app.kubernetes.io/part-of":"guestbook"},"name":"guestbook-sample","namespace":"default"},"spec":{"firstname":"firstname test Guestbook","foo":"foo test Guestbook","lastname":"lastname test Guestbook"}}
] [] [] [{kubectl-client-side-apply Update cloudstack.iblog.pro/v1 2024-02-25 07:09:28 +0000 UTC FieldsV1 {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app.kubernetes.io/created-by":{},"f:app.kubernetes.io/instance":{},"f:app.kubernetes.io/managed-by":{},"f:app.kubernetes.io/name":{},"f:app.kubernetes.io/part-of":{}}},"f:spec":{".":{},"f:firstname":{},"f:foo":{},"f:lastname":{}}} }]} {foo test Guestbook firstname test Guestbook lastname test Guestbook} { []}}
I0225 07:09:28.663313       1 guestbook_controller.go:72] Adding Finalizer
I0225 07:09:28.675869       1 guestbook_controller.go:59] POINT -1: EVENT: &context.valueCtx{Context:(*context.valueCtx)(0xc000469410), key:controller.reconcileIDKey{}, val:"2fee9b3e-860a-49ad-ade6-88fe855e711a"}
I0225 07:09:28.675903       1 guestbook_controller.go:60] POINT 0: EVENT: reconcile.Request{NamespacedName:types.NamespacedName{Namespace:"default", Name:"guestbook-sample"}}
I0225 07:09:28.676059       1 guestbook_controller.go:67] POINT 2: Geeting from Kubebuilder to &{{Guestbook cloudstack.iblog.pro/v1} {guestbook-sample  default  0c29ddd7-4223-48ae-bfab-3456d391c1a3 8839 1 2024-02-25 07:09:28 +0000 UTC <nil> <nil> map[app.kubernetes.io/created-by:guestbook app.kubernetes.io/instance:guestbook-sample app.kubernetes.io/managed-by:kustomize app.kubernetes.io/name:guestbook app.kubernetes.io/part-of:guestbook] map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"cloudstack.iblog.pro/v1","kind":"Guestbook","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"guestbook","app.kubernetes.io/instance":"guestbook-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"guestbook","app.kubernetes.io/part-of":"guestbook"},"name":"guestbook-sample","namespace":"default"},"spec":{"firstname":"firstname test Guestbook","foo":"foo test Guestbook","lastname":"lastname test Guestbook"}}
] [] [cloudstack.iblog.pro/finalizer] [{kubectl-client-side-apply Update cloudstack.iblog.pro/v1 2024-02-25 07:09:28 +0000 UTC FieldsV1 {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app.kubernetes.io/created-by":{},"f:app.kubernetes.io/instance":{},"f:app.kubernetes.io/managed-by":{},"f:app.kubernetes.io/name":{},"f:app.kubernetes.io/part-of":{}}},"f:spec":{".":{},"f:firstname":{},"f:foo":{},"f:lastname":{}}} } {manager Update cloudstack.iblog.pro/v1 2024-02-25 07:09:28 +0000 UTC FieldsV1 {"f:metadata":{"f:finalizers":{".":{},"v:\"cloudstack.iblog.pro/finalizer\"":{}}}} } {manager Update cloudstack.iblog.pro/v1 2024-02-25 07:09:28 +0000 UTC FieldsV1 {"f:status":{".":{},"f:Status":{}}} status}]} {foo test Guestbook firstname test Guestbook lastname test Guestbook} {Running []}}
........

```
