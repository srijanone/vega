# Drupal Helm Charts

## Introduction

This chart bootstraps a drupal-proxy deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

It deploys a drupal application. Optionally, you can set up an Ingress resource to access your application.

## Prerequisites

- Kubernetes 1.12+
- Helm 2.11+ or Helm 3.0-beta3+

## Installing the Chart (Helm3)

To install the chart with the release name `my-release`:

```console
$ helm install my-release ./charts/drupal
```

These commands deploy drupal on the Kubernetes cluster in the default configuration. The [Parameters](#parameters) section lists the parameters that can be configured during installation. Also includes support for MariaDB chart out of the box.

Due that the Helm Chart clones the application on the /app volume while the container is initializing, a persistent volume is not required.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm uninstall my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the drupal chart and their default values.

| Parameter                               | Description                                                                 | Default                                                 |
| --------------------------------------- | --------------------------------------------------------------------------- | ------------------------------------------------------- |
| `global.imageRegistry`                  | Global Docker image registry                                                | `docker.io`                                             |
| `phpfpm.registry`                       | PHP-FPM image registry                                                      | `nil`                                                   |
| `phpfpm.repository`                     | PHP-FPM image name                                                          | `drupal:demo`                                            |
| `phpfpm.pullPolicy`                     | PHP-FPM image pull policy                                                   | `IfNotPresent`                                          |
| `phpfpm.extraEnv`                       | PHP-FPM container environment variables                                     | `nill`                                                  |
| `phpfpm.command`                        | PHP-FPM container entry point                                               | from image                                              |
| `phpfpm.arg`                            | PHP-FPM container arguments                                                 | from image                                              |
| `phpfpm.port`                           | PHP-FPM container listing port                                              | 9000                                                    |
| `webserver.registry`                    | Web Server image registry                                                   | `nil`                                                   |
| `webserver.repository`                  | Web Server image name                                                       | `drupal:8`                                            |
| `webserver.pullPolicy`                  | Web Server image pull policy                                                | `IfNotPresent`                                          |
| `webserver.extraEnv`                    | Web Server container environment variables                                  | `nill`                                                  |
| `webserver.command`                     | Web Server container entry point                                            | from image                                              |
| `webserver.arg`                         | Web Server container arguments                                              | from image                                              |
| `webserver.port`                        | Web Server container listing port                                           | 8080                                                    |
| `nameOverride`                          | String to partially override drupal.fullname template                       | `nil`                                                   |
| `fullnameOverride`                      | String to fully override drupal.fullname template                           | `nil`                                                   |
| `applicationKind`                       | Deployment or ReplicaSet                                                    | `Deployment`                                            |
| `replicas`                              | Number of replicas for the application                                      | `1`                                                     |
| `extraEnv`                              | Any extra environment variables to be pass to the pods                      | `{}`                                                    |
| `affinity`                              | Map of drupal/pod affinities                                                | `{}` (The value is evaluated as a template)             |
| `drupalSelector`                        | drupal labels for pod assignment                                            | `{}` (The value is evaluated as a template)             |
| `tolerations`                           | Tolerations for pod assignment                                              | `[]` (The value is evaluated as a template)             |
| `securityContext.enabled`               | Enable security context                                                     | `true`                                                  |
| `securityContext.fsGroup`               | Group ID for the container                                                  | `1001`                                                  |
| `securityContext.runAsUser`             | User ID for the container                                                   | `1001`                                                  |
| `resources`                             | Resource requests and limits                                                | `{}`                                                    |
| `service.type`                          | Kubernetes Service type                                                     | `ClusterIP`                                             |
| `service.port`                          | Kubernetes Service port                                                     | `80`                                                    |
| `service.annotations`                   | Annotations for the Service                                                 | {}                                                      |
| `service.loadBalancerIP`                | LoadBalancer IP if Service type is `LoadBalancer`                           | `nil`                                                   |
| `service.drupalPort`                    | drupalPort if Service type is `LoadBalancer` or `drupalPort`                | `nil`                                                   |
| `ingress.enabled`                       | Enable ingress controller resource                                          | `false`                                                 |
| `ingress.hosts[0].name`                 | Hostname to your drupal installation                                        | `drupal.local`                                          |
| `ingress.hosts[0].path`                 | Path within the url structure                                               | `/`                                                     |
| `ingress.hosts[0].tls`                  | Utilize TLS backend in ingress                                              | `false`                                                 |
| `ingress.hosts[0].certManager`          | Add annotations for cert-manager                                            | `false`                                                 |
| `ingress.hosts[0].tlsSecret`            | TLS Secret (certificates)                                                   | `drupal.local-tls-secret`                               |
| `ingress.hosts[0].annotations`          | Annotations for this host's ingress record                                  | `[]`                                                    |
| `ingress.secrets[0].name`               | TLS Secret Name                                                             | `nil`                                                   |
| `ingress.secrets[0].certificate`        | TLS Secret Certificate                                                      | `nil`                                                   |
| `ingress.secrets[0].key`                | TLS Secret Key                                                              | `nil`                                                   |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install my-release \
  --set replicas=2 \
    ./charts/drupal
```

The above command clones the remote git repository to the `/app/` directory  of the container. Additionally it sets the number of `replicas` to `2`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. For example,

```console
$ helm install my-release -f values.yaml ./charts/drupal
```

> **Tip**: You can use the default [values.yaml](values.yaml)

## Configuration and installation details

### Set up an Ingress controller

First install the nginx-ingress controller and then deploy the drupal helm chart with the following parameters:

```console
ingress.enabled=true
ingress.host=example.com
service.type=ClusterIP
```

### Configure TLS termination for your ingress controller

You must manually create a secret containing the certificate and key for your domain. Then ensure you deploy the Helm chart with the following ingress configuration:

```yaml
ingress:
  enabled: false
  path: /
  host: example.com
  annotations:
    kubernetes.io/ingress.class: nginx
  tls:
      hosts:
        - example.com
```

### Steps to manually put the helm charts and values.yaml to S3 bucket

- Put the helm charts folder and values.yaml to Bastion host (or to a place from where s3 bucket is accessible).
- Make sure that current directory is having `charts.yaml`.
- Run below command to create a helm package
  ```
  helm package .
  ```
- You should see a helm package named - `drupal-1.0.0.tgz`.
- Upload the `drupal-1.0.0.tgz` helm package to s3 bucket :
  ```
  aws s3 cp drupal-1.0.0.tgz s3://s3-helm/drupal/drupal-1.0.0.tgz --sse=AES256 --region=ap-southeast-1
  ```
- Upload the `values.yaml` to s3 bucket :
  ```
  aws s3 cp values-<env>.yaml s3://s3-helm/drupal/values/st-<env>-values.yaml --sse=AES256 --region=ap-southeast-1
  ```
