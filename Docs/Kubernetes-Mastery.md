# Kubernetes Mastery Guide
## From Beginner to Expert & Solution Architect Level

---

## Table of Contents
1. [Beginner Level](#beginner-level)
2. [Intermediate Level](#intermediate-level)
3. [Advanced Level](#advanced-level)
4. [Expert Level](#expert-level)
5. [Solution Architect Level](#solution-architect-level)
6. [Practice Questions with Answers](#practice-questions-with-answers)

---

# Beginner Level

## 1. Introduction to Kubernetes

### What is Kubernetes?
Kubernetes (K8s) is an open-source container orchestration platform that automates deployment, scaling, and management of containerized applications. Originally developed by Google, it's now maintained by the Cloud Native Computing Foundation (CNCF).

### Why Kubernetes?
- **Automatic scaling** - Scale applications up/down based on demand
- **Self-healing** - Restart failed containers, replace nodes
- **Load balancing** - Distribute traffic across containers
- **Rolling updates** - Zero-downtime deployments
- **Secret management** - Manage sensitive data
- **Storage orchestration** - Mount storage systems automatically

### Kubernetes Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                        CONTROL PLANE                            │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌───────────┐  │
│  │ API Server  │ │   etcd      │ │ Scheduler   │ │Controller │  │
│  │             │ │  (store)    │ │             │ │  Manager  │  │
│  └─────────────┘ └─────────────┘ └─────────────┘ └───────────┘  │
└─────────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────────┐
│                         WORKER NODES                            │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │ Node 1                                                      ││
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐                     ││
│  │  │  kubelet │ │kube-proxy│ │Container │                     ││
│  │  │          │ │          │ │ Runtime  │                     ││
│  │  └──────────┘ └──────────┘ └──────────┘                     ││
│  │  ┌─────────────────────────────────────┐                    ││
│  │  │              Pods                   │                    ││
│  │  │  ┌─────┐ ┌─────┐ ┌─────┐            │                    ││
│  │  │  │ Pod │ │ Pod │ │ Pod │            │                    ││
│  │  │  └─────┘ └─────┘ └─────┘            │                    ││
│  │  └─────────────────────────────────────┘                    ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
```

### Control Plane Components

| Component | Description |
|-----------|-------------|
| **API Server** | Frontend for Kubernetes control plane |
| **etcd** | Key-value store for cluster data |
| **Scheduler** | Assigns pods to nodes |
| **Controller Manager** | Runs controller processes |
| **Cloud Controller** | Interacts with cloud provider |

### Node Components

| Component | Description |
|-----------|-------------|
| **kubelet** | Agent that runs on each node |
| **kube-proxy** | Network proxy, implements Services |
| **Container Runtime** | Runs containers (containerd, CRI-O) |

---

## 2. Setting Up Kubernetes

### Local Development Options

```bash
# Minikube - Single node cluster
brew install minikube
minikube start
minikube status

# Kind - Kubernetes in Docker
brew install kind
kind create cluster
kind get clusters

# Docker Desktop - Enable Kubernetes in settings

# k3d - Lightweight K3s in Docker
brew install k3d
k3d cluster create mycluster
```

### kubectl Installation

```bash
# macOS
brew install kubectl

# Verify installation
kubectl version --client

# Configure kubectl
kubectl config view
kubectl config get-contexts
kubectl config use-context minikube
```

### Basic kubectl Commands

```bash
# Cluster info
kubectl cluster-info
kubectl get nodes

# Get resources
kubectl get pods
kubectl get services
kubectl get deployments
kubectl get all

# Describe resources
kubectl describe pod <pod-name>
kubectl describe node <node-name>

# Create resources
kubectl create -f manifest.yaml
kubectl apply -f manifest.yaml

# Delete resources
kubectl delete -f manifest.yaml
kubectl delete pod <pod-name>

# Logs and exec
kubectl logs <pod-name>
kubectl logs -f <pod-name>  # Follow
kubectl exec -it <pod-name> -- /bin/sh

# Port forwarding
kubectl port-forward <pod-name> 8080:80
```

---

## 3. Pods

### What is a Pod?
A Pod is the smallest deployable unit in Kubernetes. It can contain one or more containers that share storage, network, and specifications.

### Pod Manifest

```yaml
# pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
  labels:
    app: nginx
    environment: development
spec:
  containers:
  - name: nginx
    image: nginx:1.25
    ports:
    - containerPort: 80
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  restartPolicy: Always
```

### Pod Commands

```bash
# Create pod
kubectl apply -f pod.yaml

# List pods
kubectl get pods
kubectl get pods -o wide
kubectl get pods --show-labels

# Describe pod
kubectl describe pod nginx-pod

# View logs
kubectl logs nginx-pod
kubectl logs nginx-pod -c <container-name>  # Multi-container

# Execute command
kubectl exec -it nginx-pod -- /bin/bash
kubectl exec nginx-pod -- ls /usr/share/nginx/html

# Delete pod
kubectl delete pod nginx-pod
```

### Multi-Container Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: multi-container-pod
spec:
  containers:
  - name: app
    image: myapp:latest
    ports:
    - containerPort: 8080
    volumeMounts:
    - name: shared-logs
      mountPath: /var/log/app
  
  - name: log-sidecar
    image: busybox
    command: ['sh', '-c', 'tail -f /var/log/app/app.log']
    volumeMounts:
    - name: shared-logs
      mountPath: /var/log/app
  
  volumes:
  - name: shared-logs
    emptyDir: {}
```

---

## 4. ReplicaSets and Deployments

### ReplicaSet

A ReplicaSet ensures a specified number of pod replicas are running.

```yaml
# replicaset.yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: nginx-replicaset
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.25
        ports:
        - containerPort: 80
```

### Deployment

A Deployment manages ReplicaSets and provides declarative updates.

```yaml
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.25
        ports:
        - containerPort: 80
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /
            port: 80
          initialDelaySeconds: 15
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /
            port: 80
          initialDelaySeconds: 5
          periodSeconds: 5
```

### Deployment Commands

```bash
# Create deployment
kubectl apply -f deployment.yaml

# List deployments
kubectl get deployments
kubectl get deploy

# Scale deployment
kubectl scale deployment nginx-deployment --replicas=5

# Update image
kubectl set image deployment/nginx-deployment nginx=nginx:1.26

# Check rollout status
kubectl rollout status deployment/nginx-deployment

# View rollout history
kubectl rollout history deployment/nginx-deployment

# Rollback
kubectl rollout undo deployment/nginx-deployment
kubectl rollout undo deployment/nginx-deployment --to-revision=2

# Pause/Resume rollout
kubectl rollout pause deployment/nginx-deployment
kubectl rollout resume deployment/nginx-deployment
```

---

## 5. Services

### Service Types

| Type | Description |
|------|-------------|
| **ClusterIP** | Internal cluster IP (default) |
| **NodePort** | Exposes on each node's IP at static port |
| **LoadBalancer** | Cloud load balancer |
| **ExternalName** | Maps to external DNS name |

### ClusterIP Service

```yaml
# service-clusterip.yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  type: ClusterIP
  selector:
    app: nginx
  ports:
  - protocol: TCP
    port: 80        # Service port
    targetPort: 80  # Container port
```

### NodePort Service

```yaml
# service-nodeport.yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-nodeport
spec:
  type: NodePort
  selector:
    app: nginx
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80
    nodePort: 30080  # 30000-32767
```

### LoadBalancer Service

```yaml
# service-loadbalancer.yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-lb
spec:
  type: LoadBalancer
  selector:
    app: nginx
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80
```

### Service Commands

```bash
# Create service
kubectl apply -f service.yaml

# List services
kubectl get services
kubectl get svc

# Describe service
kubectl describe service nginx-service

# Get endpoints
kubectl get endpoints nginx-service

# Delete service
kubectl delete service nginx-service
```

---

## 6. Namespaces

### What are Namespaces?
Namespaces provide a way to divide cluster resources between multiple users or teams.

### Default Namespaces

| Namespace | Purpose |
|-----------|---------|
| `default` | Default namespace for resources |
| `kube-system` | Kubernetes system components |
| `kube-public` | Publicly accessible resources |
| `kube-node-lease` | Node heartbeat data |

### Namespace Commands

```bash
# List namespaces
kubectl get namespaces
kubectl get ns

# Create namespace
kubectl create namespace development
kubectl create ns production

# Get resources in namespace
kubectl get pods -n development
kubectl get all -n kube-system

# Set default namespace
kubectl config set-context --current --namespace=development

# Delete namespace (deletes all resources in it)
kubectl delete namespace development
```

### Namespace Manifest

```yaml
# namespace.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: development
  labels:
    environment: dev
```

### Resource with Namespace

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  namespace: development  # Specify namespace
spec:
  # ...
```

---

## 7. ConfigMaps and Secrets

### ConfigMap

ConfigMaps store non-sensitive configuration data.

```yaml
# configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  DATABASE_HOST: "postgres.default.svc.cluster.local"
  DATABASE_PORT: "5432"
  LOG_LEVEL: "info"
  config.json: |
    {
      "feature_flags": {
        "new_ui": true
      }
    }
```

### Using ConfigMap

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  template:
    spec:
      containers:
      - name: app
        image: myapp:latest
        # As environment variables
        envFrom:
        - configMapRef:
            name: app-config
        # Individual keys
        env:
        - name: DB_HOST
          valueFrom:
            configMapKeyRef:
              name: app-config
              key: DATABASE_HOST
        # As volume
        volumeMounts:
        - name: config-volume
          mountPath: /etc/config
      volumes:
      - name: config-volume
        configMap:
          name: app-config
```

### Secret

Secrets store sensitive data (base64 encoded).

```yaml
# secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
type: Opaque
data:
  # base64 encoded values
  DB_PASSWORD: cGFzc3dvcmQxMjM=
  API_KEY: YXBpLWtleS1zZWNyZXQ=
---
# Using stringData (plain text, auto-encoded)
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets-string
type: Opaque
stringData:
  DB_PASSWORD: password123
  API_KEY: api-key-secret
```

### Using Secrets

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  template:
    spec:
      containers:
      - name: app
        image: myapp:latest
        envFrom:
        - secretRef:
            name: app-secrets
        env:
        - name: DATABASE_PASSWORD
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: DB_PASSWORD
        volumeMounts:
        - name: secret-volume
          mountPath: /etc/secrets
          readOnly: true
      volumes:
      - name: secret-volume
        secret:
          secretName: app-secrets
```

### ConfigMap/Secret Commands

```bash
# Create ConfigMap
kubectl create configmap app-config --from-literal=key1=value1
kubectl create configmap app-config --from-file=config.properties
kubectl create configmap app-config --from-env-file=.env

# Create Secret
kubectl create secret generic app-secrets --from-literal=password=secret
kubectl create secret generic app-secrets --from-file=ssh-key=~/.ssh/id_rsa

# View (secrets are base64 encoded)
kubectl get configmaps
kubectl get secrets
kubectl describe secret app-secrets

# Decode secret
kubectl get secret app-secrets -o jsonpath='{.data.DB_PASSWORD}' | base64 --decode
```

---

# Intermediate Level

## 8. Storage

### Volumes

```yaml
# Pod with volumes
apiVersion: v1
kind: Pod
metadata:
  name: pod-with-volume
spec:
  containers:
  - name: app
    image: myapp:latest
    volumeMounts:
    - name: data
      mountPath: /data
    - name: config
      mountPath: /etc/config
    - name: cache
      mountPath: /tmp/cache
  volumes:
  # EmptyDir - temporary storage
  - name: cache
    emptyDir: {}
  
  # HostPath - node filesystem
  - name: data
    hostPath:
      path: /var/data
      type: DirectoryOrCreate
  
  # ConfigMap as volume
  - name: config
    configMap:
      name: app-config
```

### PersistentVolume (PV)

```yaml
# persistentvolume.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-storage
spec:
  capacity:
    storage: 10Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: manual
  hostPath:
    path: /mnt/data
```

### PersistentVolumeClaim (PVC)

```yaml
# persistentvolumeclaim.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-storage
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: manual
  resources:
    requests:
      storage: 5Gi
```

### Using PVC in Pod

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
spec:
  template:
    spec:
      containers:
      - name: postgres
        image: postgres:15
        volumeMounts:
        - name: postgres-storage
          mountPath: /var/lib/postgresql/data
      volumes:
      - name: postgres-storage
        persistentVolumeClaim:
          claimName: pvc-storage
```

### StorageClass

```yaml
# storageclass.yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp3
  iopsPerGB: "10"
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
```

### Dynamic Provisioning

```yaml
# PVC with StorageClass
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: dynamic-pvc
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: fast  # Reference StorageClass
  resources:
    requests:
      storage: 10Gi
```

---

## 9. Ingress

### What is Ingress?
Ingress exposes HTTP/HTTPS routes from outside the cluster to services within the cluster.

### Ingress Controller

```bash
# Install NGINX Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml

# Verify installation
kubectl get pods -n ingress-nginx
kubectl get svc -n ingress-nginx
```

### Basic Ingress

```yaml
# ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: app-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  rules:
  - host: app.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: frontend-service
            port:
              number: 80
      - path: /api
        pathType: Prefix
        backend:
          service:
            name: api-service
            port:
              number: 8080
```

### TLS Ingress

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: tls-ingress
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - app.example.com
    secretName: tls-secret
  rules:
  - host: app.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: app-service
            port:
              number: 80
```

### Create TLS Secret

```bash
# Generate self-signed certificate
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout tls.key -out tls.crt \
  -subj "/CN=app.example.com"

# Create secret
kubectl create secret tls tls-secret --key tls.key --cert tls.crt
```

---

## 10. StatefulSets

### What is StatefulSet?
StatefulSets manage stateful applications with stable network identities and persistent storage.

### StatefulSet Manifest

```yaml
# statefulset.yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: postgres
spec:
  serviceName: postgres
  replicas: 3
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:15
        ports:
        - containerPort: 5432
          name: postgres
        env:
        - name: POSTGRES_PASSWORD
          valueFrom:
            secretKeyRef:
              name: postgres-secret
              key: password
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql/data
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: ["ReadWriteOnce"]
      storageClassName: fast
      resources:
        requests:
          storage: 10Gi
```

### Headless Service for StatefulSet

```yaml
apiVersion: v1
kind: Service
metadata:
  name: postgres
spec:
  clusterIP: None  # Headless
  selector:
    app: postgres
  ports:
  - port: 5432
    targetPort: 5432
```

### StatefulSet Features

- **Stable Network Identity**: `postgres-0`, `postgres-1`, `postgres-2`
- **DNS Names**: `postgres-0.postgres.default.svc.cluster.local`
- **Ordered Deployment**: Pods created in order (0, 1, 2)
- **Ordered Deletion**: Pods deleted in reverse order (2, 1, 0)
- **Persistent Storage**: Each pod gets its own PVC

---

## 11. DaemonSets

### What is DaemonSet?
A DaemonSet ensures that all (or some) nodes run a copy of a Pod.

### DaemonSet Manifest

```yaml
# daemonset.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd
  namespace: kube-system
spec:
  selector:
    matchLabels:
      name: fluentd
  template:
    metadata:
      labels:
        name: fluentd
    spec:
      tolerations:
      - key: node-role.kubernetes.io/control-plane
        operator: Exists
        effect: NoSchedule
      containers:
      - name: fluentd
        image: fluent/fluentd:v1.16
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - name: varlog
          mountPath: /var/log
        - name: varlibdockercontainers
          mountPath: /var/lib/docker/containers
          readOnly: true
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
      - name: varlibdockercontainers
        hostPath:
          path: /var/lib/docker/containers
```

### Common DaemonSet Use Cases
- Log collectors (Fluentd, Filebeat)
- Node monitoring (Prometheus Node Exporter)
- Network plugins (Calico, Cilium)
- Storage daemons (Ceph, GlusterFS)

---

## 12. Jobs and CronJobs

### Job

A Job creates one or more pods and ensures they successfully complete.

```yaml
# job.yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: backup-job
spec:
  completions: 1
  parallelism: 1
  backoffLimit: 3
  activeDeadlineSeconds: 600
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: backup
        image: backup-tool:latest
        command: ["/bin/sh", "-c"]
        args:
        - |
          pg_dump -h $DB_HOST -U $DB_USER mydb > /backup/dump.sql
          aws s3 cp /backup/dump.sql s3://my-bucket/backups/
        env:
        - name: DB_HOST
          value: "postgres"
        - name: DB_USER
          value: "admin"
```

### CronJob

A CronJob creates Jobs on a time-based schedule.

```yaml
# cronjob.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: backup-cronjob
spec:
  schedule: "0 2 * * *"  # Every day at 2 AM
  concurrencyPolicy: Forbid  # Allow, Forbid, Replace
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1
  startingDeadlineSeconds: 300
  jobTemplate:
    spec:
      template:
        spec:
          restartPolicy: OnFailure
          containers:
          - name: backup
            image: backup-tool:latest
            command: ["/scripts/backup.sh"]
```

### Job/CronJob Commands

```bash
# Create job
kubectl apply -f job.yaml

# List jobs
kubectl get jobs

# View job logs
kubectl logs job/backup-job

# Delete job
kubectl delete job backup-job

# List cronjobs
kubectl get cronjobs

# Manually trigger cronjob
kubectl create job --from=cronjob/backup-cronjob manual-backup
```

---

## 13. Resource Management

### Resource Requests and Limits

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      containers:
      - name: app
        image: myapp:latest
        resources:
          requests:        # Guaranteed resources
            memory: "256Mi"
            cpu: "250m"
          limits:          # Maximum resources
            memory: "512Mi"
            cpu: "500m"
```

### LimitRange

Sets default resource limits for a namespace.

```yaml
# limitrange.yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: default-limits
  namespace: development
spec:
  limits:
  - type: Container
    default:
      memory: "256Mi"
      cpu: "500m"
    defaultRequest:
      memory: "128Mi"
      cpu: "250m"
    min:
      memory: "64Mi"
      cpu: "100m"
    max:
      memory: "1Gi"
      cpu: "1"
  - type: Pod
    max:
      memory: "2Gi"
      cpu: "2"
```

### ResourceQuota

Limits total resources in a namespace.

```yaml
# resourcequota.yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-quota
  namespace: development
spec:
  hard:
    requests.cpu: "10"
    requests.memory: "20Gi"
    limits.cpu: "20"
    limits.memory: "40Gi"
    pods: "50"
    services: "10"
    persistentvolumeclaims: "10"
    secrets: "20"
    configmaps: "20"
```

---

## 14. Health Checks

### Liveness Probe

Determines if a container is running. If it fails, the container is restarted.

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 15
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3
```

### Readiness Probe

Determines if a container is ready to receive traffic.

```yaml
readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 5
  successThreshold: 1
  failureThreshold: 3
```

### Startup Probe

Used for slow-starting containers. Disables liveness/readiness until it succeeds.

```yaml
startupProbe:
  httpGet:
    path: /healthz
    port: 8080
  failureThreshold: 30
  periodSeconds: 10
```

### Probe Types

```yaml
# HTTP probe
httpGet:
  path: /health
  port: 8080
  httpHeaders:
  - name: Custom-Header
    value: "value"

# TCP probe
tcpSocket:
  port: 3306

# Exec probe
exec:
  command:
  - cat
  - /tmp/healthy

# gRPC probe
grpc:
  port: 50051
```

### Complete Example

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      containers:
      - name: app
        image: myapp:latest
        ports:
        - containerPort: 8080
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 15
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        startupProbe:
          httpGet:
            path: /healthz
            port: 8080
          failureThreshold: 30
          periodSeconds: 10
```

---

# Advanced Level

## 15. RBAC (Role-Based Access Control)

### RBAC Components

| Component | Scope | Description |
|-----------|-------|-------------|
| Role | Namespace | Permissions within a namespace |
| ClusterRole | Cluster | Permissions across the cluster |
| RoleBinding | Namespace | Binds Role to users/groups |
| ClusterRoleBinding | Cluster | Binds ClusterRole to users/groups |

### Role

```yaml
# role.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: development
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
- apiGroups: [""]
  resources: ["pods/log"]
  verbs: ["get"]
```

### RoleBinding

```yaml
# rolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
  namespace: development
subjects:
- kind: User
  name: developer
  apiGroup: rbac.authorization.k8s.io
- kind: ServiceAccount
  name: app-sa
  namespace: development
- kind: Group
  name: developers
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

### ClusterRole

```yaml
# clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: node-reader
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "watch", "list"]
- apiGroups: [""]
  resources: ["persistentvolumes"]
  verbs: ["get", "list"]
```

### ClusterRoleBinding

```yaml
# clusterrolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: read-nodes
subjects:
- kind: Group
  name: sre-team
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: node-reader
  apiGroup: rbac.authorization.k8s.io
```

### ServiceAccount

```yaml
# serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: app-sa
  namespace: development
---
# Use in Pod
apiVersion: v1
kind: Pod
metadata:
  name: app
spec:
  serviceAccountName: app-sa
  containers:
  - name: app
    image: myapp:latest
```

### RBAC Commands

```bash
# Check permissions
kubectl auth can-i get pods --as developer -n development
kubectl auth can-i create deployments --as developer -n development
kubectl auth can-i '*' '*' --as system:admin

# List roles
kubectl get roles -n development
kubectl get clusterroles

# Describe role
kubectl describe role pod-reader -n development
```

---

## 16. Network Policies

### What are Network Policies?
Network Policies control traffic flow between pods, namespaces, and external endpoints.

### Default Deny All

```yaml
# deny-all.yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
  namespace: production
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  - Egress
```

### Allow Specific Traffic

```yaml
# allow-frontend-to-api.yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: api-network-policy
  namespace: production
spec:
  podSelector:
    matchLabels:
      app: api
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
    - namespaceSelector:
        matchLabels:
          name: monitoring
    ports:
    - protocol: TCP
      port: 8080
  egress:
  - to:
    - podSelector:
        matchLabels:
          app: database
    ports:
    - protocol: TCP
      port: 5432
  - to:
    - namespaceSelector: {}
      podSelector:
        matchLabels:
          k8s-app: kube-dns
    ports:
    - protocol: UDP
      port: 53
```

### Allow External Traffic

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-external
spec:
  podSelector:
    matchLabels:
      app: frontend
  policyTypes:
  - Ingress
  ingress:
  - from:
    - ipBlock:
        cidr: 0.0.0.0/0
        except:
        - 10.0.0.0/8
    ports:
    - protocol: TCP
      port: 443
```

---

## 17. Horizontal Pod Autoscaler (HPA)

### Metrics Server

```bash
# Install metrics server
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

# Verify
kubectl top nodes
kubectl top pods
```

### HPA Manifest

```yaml
# hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: app-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: app
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 10
        periodSeconds: 60
    scaleUp:
      stabilizationWindowSeconds: 0
      policies:
      - type: Percent
        value: 100
        periodSeconds: 15
      - type: Pods
        value: 4
        periodSeconds: 15
      selectPolicy: Max
```

### HPA Commands

```bash
# Create HPA
kubectl apply -f hpa.yaml

# Create HPA imperatively
kubectl autoscale deployment app --min=2 --max=10 --cpu-percent=70

# Get HPA
kubectl get hpa

# Describe HPA
kubectl describe hpa app-hpa
```

---

## 18. Vertical Pod Autoscaler (VPA)

### VPA Manifest

```yaml
# vpa.yaml
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
  name: app-vpa
spec:
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: app
  updatePolicy:
    updateMode: "Auto"  # Off, Initial, Recreate, Auto
  resourcePolicy:
    containerPolicies:
    - containerName: app
      minAllowed:
        cpu: 100m
        memory: 128Mi
      maxAllowed:
        cpu: 2
        memory: 2Gi
      controlledResources: ["cpu", "memory"]
```

---

## 19. Pod Disruption Budgets (PDB)

### PDB Manifest

```yaml
# pdb.yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: app-pdb
spec:
  minAvailable: 2    # or use maxUnavailable
  # maxUnavailable: 1
  selector:
    matchLabels:
      app: myapp
```

### PDB Commands

```bash
# Create PDB
kubectl apply -f pdb.yaml

# List PDBs
kubectl get pdb

# Describe PDB
kubectl describe pdb app-pdb
```

---

## 20. Taints and Tolerations

### Taints

```bash
# Add taint to node
kubectl taint nodes node1 key=value:NoSchedule
kubectl taint nodes node1 dedicated=gpu:NoExecute

# Remove taint
kubectl taint nodes node1 key=value:NoSchedule-

# Taint effects:
# - NoSchedule: Don't schedule new pods
# - PreferNoSchedule: Avoid scheduling new pods
# - NoExecute: Evict existing pods and don't schedule new ones
```

### Tolerations

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gpu-app
spec:
  template:
    spec:
      tolerations:
      - key: "dedicated"
        operator: "Equal"
        value: "gpu"
        effect: "NoSchedule"
      - key: "node.kubernetes.io/not-ready"
        operator: "Exists"
        effect: "NoExecute"
        tolerationSeconds: 300
      containers:
      - name: app
        image: gpu-app:latest
```

---

## 21. Node Affinity and Pod Affinity

### Node Affinity

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  template:
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/os
                operator: In
                values:
                - linux
              - key: node-type
                operator: In
                values:
                - compute
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            preference:
              matchExpressions:
              - key: zone
                operator: In
                values:
                - us-east-1a
      containers:
      - name: app
        image: myapp:latest
```

### Pod Affinity and Anti-Affinity

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web
spec:
  template:
    spec:
      affinity:
        # Co-locate with cache pods
        podAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - cache
            topologyKey: kubernetes.io/hostname
        # Spread across zones
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - web
              topologyKey: topology.kubernetes.io/zone
      containers:
      - name: web
        image: web:latest
```

---

# Expert Level

## 22. Custom Resource Definitions (CRDs)

### Create CRD

```yaml
# crd.yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: databases.example.com
spec:
  group: example.com
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            required: ["engine", "version"]
            properties:
              engine:
                type: string
                enum: ["postgres", "mysql", "mongodb"]
              version:
                type: string
              replicas:
                type: integer
                minimum: 1
                maximum: 10
                default: 1
              storage:
                type: string
                pattern: "^[0-9]+Gi$"
          status:
            type: object
            properties:
              state:
                type: string
              endpoint:
                type: string
    subresources:
      status: {}
    additionalPrinterColumns:
    - name: Engine
      type: string
      jsonPath: .spec.engine
    - name: Version
      type: string
      jsonPath: .spec.version
    - name: State
      type: string
      jsonPath: .status.state
  scope: Namespaced
  names:
    plural: databases
    singular: database
    kind: Database
    shortNames:
    - db
```

### Create Custom Resource

```yaml
# database.yaml
apiVersion: example.com/v1
kind: Database
metadata:
  name: my-postgres
spec:
  engine: postgres
  version: "15"
  replicas: 3
  storage: "100Gi"
```

### CRD Commands

```bash
# Apply CRD
kubectl apply -f crd.yaml

# List CRDs
kubectl get crd

# Create custom resource
kubectl apply -f database.yaml

# List custom resources
kubectl get databases
kubectl get db

# Describe
kubectl describe database my-postgres
```

---

## 23. Operators

### Operator Pattern

Operators are software extensions that use custom resources to manage applications.

### Simple Operator with kubebuilder

```bash
# Install kubebuilder
brew install kubebuilder

# Create operator project
kubebuilder init --domain example.com --repo github.com/example/database-operator
kubebuilder create api --group apps --version v1 --kind Database

# Generate manifests
make manifests

# Build and push operator
make docker-build docker-push IMG=<registry>/database-operator:v1

# Deploy operator
make deploy IMG=<registry>/database-operator:v1
```

### Controller Logic (Go)

```go
// controllers/database_controller.go
package controllers

import (
    "context"
    
    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/runtime"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    
    examplev1 "github.com/example/database-operator/api/v1"
)

type DatabaseReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

func (r *DatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    log := ctrl.LoggerFrom(ctx)
    
    // Fetch the Database instance
    var database examplev1.Database
    if err := r.Get(ctx, req.NamespacedName, &database); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }
    
    // Create StatefulSet if not exists
    // Create Service if not exists
    // Update status
    
    log.Info("Reconciled database", "name", database.Name)
    return ctrl.Result{}, nil
}

func (r *DatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&examplev1.Database{}).
        Owns(&appsv1.StatefulSet{}).
        Owns(&corev1.Service{}).
        Complete(r)
}
```

---

## 24. Admission Controllers

### ValidatingWebhookConfiguration

```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: pod-validator
webhooks:
- name: pod-validator.example.com
  rules:
  - apiGroups: [""]
    apiVersions: ["v1"]
    operations: ["CREATE", "UPDATE"]
    resources: ["pods"]
  clientConfig:
    service:
      name: webhook-service
      namespace: webhook-system
      path: /validate-pods
    caBundle: <base64-encoded-ca>
  admissionReviewVersions: ["v1"]
  sideEffects: None
  failurePolicy: Fail
```

### MutatingWebhookConfiguration

```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  name: pod-mutator
webhooks:
- name: pod-mutator.example.com
  rules:
  - apiGroups: [""]
    apiVersions: ["v1"]
    operations: ["CREATE"]
    resources: ["pods"]
  clientConfig:
    service:
      name: webhook-service
      namespace: webhook-system
      path: /mutate-pods
    caBundle: <base64-encoded-ca>
  admissionReviewVersions: ["v1"]
  sideEffects: None
  reinvocationPolicy: Never
```

### Webhook Server (Go)

```go
package main

import (
    "encoding/json"
    "net/http"
    
    admissionv1 "k8s.io/api/admission/v1"
    corev1 "k8s.io/api/core/v1"
)

func validatePod(w http.ResponseWriter, r *http.Request) {
    var admissionReview admissionv1.AdmissionReview
    json.NewDecoder(r.Body).Decode(&admissionReview)
    
    var pod corev1.Pod
    json.Unmarshal(admissionReview.Request.Object.Raw, &pod)
    
    // Validate pod
    allowed := true
    message := "Pod is valid"
    
    // Example: Require resource limits
    for _, container := range pod.Spec.Containers {
        if container.Resources.Limits.Cpu().IsZero() {
            allowed = false
            message = "CPU limit is required"
            break
        }
    }
    
    response := admissionv1.AdmissionReview{
        Response: &admissionv1.AdmissionResponse{
            UID:     admissionReview.Request.UID,
            Allowed: allowed,
            Result: &metav1.Status{
                Message: message,
            },
        },
    }
    
    json.NewEncoder(w).Encode(response)
}
```

---

## 25. Service Mesh (Istio)

### Install Istio

```bash
# Download Istio
curl -L https://istio.io/downloadIstio | sh -
cd istio-*
export PATH=$PWD/bin:$PATH

# Install Istio
istioctl install --set profile=demo -y

# Enable sidecar injection
kubectl label namespace default istio-injection=enabled
```

### VirtualService

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: app-routing
spec:
  hosts:
  - app.example.com
  gateways:
  - app-gateway
  http:
  - match:
    - headers:
        x-version:
          exact: "v2"
    route:
    - destination:
        host: app
        subset: v2
  - route:
    - destination:
        host: app
        subset: v1
      weight: 90
    - destination:
        host: app
        subset: v2
      weight: 10
```

### DestinationRule

```yaml
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: app-destination
spec:
  host: app
  trafficPolicy:
    connectionPool:
      tcp:
        maxConnections: 100
      http:
        h2UpgradePolicy: UPGRADE
        http1MaxPendingRequests: 100
    loadBalancer:
      simple: ROUND_ROBIN
    outlierDetection:
      consecutive5xxErrors: 5
      interval: 30s
      baseEjectionTime: 30s
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```

### Gateway

```yaml
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: app-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
  - port:
      number: 443
      name: https
      protocol: HTTPS
    tls:
      mode: SIMPLE
      credentialName: app-tls
    hosts:
    - app.example.com
```

---

## 26. GitOps with ArgoCD

### Install ArgoCD

```bash
# Create namespace
kubectl create namespace argocd

# Install ArgoCD
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Access UI
kubectl port-forward svc/argocd-server -n argocd 8080:443

# Get admin password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

### Application CRD

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: myapp
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/org/myapp-manifests.git
    targetRevision: HEAD
    path: overlays/production
  destination:
    server: https://kubernetes.default.svc
    namespace: production
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
      allowEmpty: false
    syncOptions:
    - CreateNamespace=true
    retry:
      limit: 5
      backoff:
        duration: 5s
        factor: 2
        maxDuration: 3m
```

### ApplicationSet

```yaml
apiVersion: argoproj.io/v1alpha1
kind: ApplicationSet
metadata:
  name: cluster-apps
  namespace: argocd
spec:
  generators:
  - list:
      elements:
      - cluster: prod-us-east
        url: https://prod-us-east.example.com
      - cluster: prod-us-west
        url: https://prod-us-west.example.com
  template:
    metadata:
      name: '{{cluster}}-app'
    spec:
      project: default
      source:
        repoURL: https://github.com/org/apps.git
        targetRevision: HEAD
        path: '{{cluster}}'
      destination:
        server: '{{url}}'
        namespace: default
```

---

# Solution Architect Level

## 27. Multi-Cluster Management

### Cluster Federation

```yaml
# Kubefed - Federated Deployment
apiVersion: types.kubefed.io/v1beta1
kind: FederatedDeployment
metadata:
  name: app
  namespace: default
spec:
  template:
    metadata:
      labels:
        app: nginx
    spec:
      replicas: 3
      selector:
        matchLabels:
          app: nginx
      template:
        metadata:
          labels:
            app: nginx
        spec:
          containers:
          - name: nginx
            image: nginx:latest
  placement:
    clusters:
    - name: cluster1
    - name: cluster2
  overrides:
  - clusterName: cluster1
    clusterOverrides:
    - path: "/spec/replicas"
      value: 5
```

### Multi-Cluster Service Discovery

```yaml
# ServiceExport
apiVersion: multicluster.x-k8s.io/v1alpha1
kind: ServiceExport
metadata:
  name: myservice
  namespace: default
---
# ServiceImport
apiVersion: multicluster.x-k8s.io/v1alpha1
kind: ServiceImport
metadata:
  name: myservice
  namespace: default
spec:
  type: ClusterSetIP
  ports:
  - port: 80
    protocol: TCP
```

---

## 28. Disaster Recovery

### Velero Backup

```bash
# Install Velero
velero install \
  --provider aws \
  --plugins velero/velero-plugin-for-aws:v1.8.0 \
  --bucket my-backup-bucket \
  --backup-location-config region=us-east-1 \
  --secret-file ./credentials-velero

# Create backup
velero backup create full-backup --include-namespaces default,production

# Schedule regular backups
velero schedule create daily-backup --schedule="0 2 * * *"

# Restore from backup
velero restore create --from-backup full-backup
```

### Backup Configuration

```yaml
apiVersion: velero.io/v1
kind: Backup
metadata:
  name: production-backup
  namespace: velero
spec:
  includedNamespaces:
  - production
  - staging
  excludedResources:
  - events
  - pods
  storageLocation: default
  volumeSnapshotLocations:
  - default
  ttl: 720h  # 30 days
  hooks:
    resources:
    - name: my-hook
      includedNamespaces:
      - production
      pre:
      - exec:
          container: app
          command:
          - /bin/sh
          - -c
          - "pg_dump -U postgres mydb > /backup/dump.sql"
```

---

## 29. Production Best Practices

### Complete Production Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: production-app
  namespace: production
  labels:
    app: myapp
    version: v1.2.3
spec:
  replicas: 3
  revisionHistoryLimit: 5
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
        version: v1.2.3
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
    spec:
      serviceAccountName: myapp-sa
      securityContext:
        runAsNonRoot: true
        runAsUser: 1000
        fsGroup: 1000
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchLabels:
                  app: myapp
              topologyKey: topology.kubernetes.io/zone
      topologySpreadConstraints:
      - maxSkew: 1
        topologyKey: topology.kubernetes.io/zone
        whenUnsatisfiable: ScheduleAnyway
        labelSelector:
          matchLabels:
            app: myapp
      containers:
      - name: app
        image: myapp:v1.2.3
        imagePullPolicy: Always
        ports:
        - containerPort: 8080
          name: http
        - containerPort: 9090
          name: metrics
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        envFrom:
        - configMapRef:
            name: app-config
        - secretRef:
            name: app-secrets
        resources:
          requests:
            cpu: 250m
            memory: 256Mi
          limits:
            cpu: 500m
            memory: 512Mi
        securityContext:
          allowPrivilegeEscalation: false
          readOnlyRootFilesystem: true
          capabilities:
            drop:
            - ALL
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 15
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
          successThreshold: 1
          failureThreshold: 3
        volumeMounts:
        - name: tmp
          mountPath: /tmp
        - name: config
          mountPath: /etc/config
          readOnly: true
      volumes:
      - name: tmp
        emptyDir: {}
      - name: config
        configMap:
          name: app-config
      terminationGracePeriodSeconds: 30
---
apiVersion: v1
kind: Service
metadata:
  name: myapp
  namespace: production
spec:
  type: ClusterIP
  selector:
    app: myapp
  ports:
  - name: http
    port: 80
    targetPort: 8080
---
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: myapp-pdb
  namespace: production
spec:
  minAvailable: 2
  selector:
    matchLabels:
      app: myapp
---
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: myapp-hpa
  namespace: production
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: production-app
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

---

## 30. Monitoring Stack

### Prometheus Stack

```yaml
# Using kube-prometheus-stack Helm chart
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --set grafana.adminPassword=admin \
  --set prometheus.prometheusSpec.retention=30d \
  --set prometheus.prometheusSpec.storageSpec.volumeClaimTemplate.spec.resources.requests.storage=50Gi
```

### ServiceMonitor

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: myapp-monitor
  namespace: monitoring
spec:
  selector:
    matchLabels:
      app: myapp
  namespaceSelector:
    matchNames:
    - production
  endpoints:
  - port: metrics
    interval: 30s
    path: /metrics
```

### PrometheusRule (Alerts)

```yaml
apiVersion: monitoring.coreos.com/v1
kind: PrometheusRule
metadata:
  name: myapp-alerts
  namespace: monitoring
spec:
  groups:
  - name: myapp
    rules:
    - alert: HighErrorRate
      expr: |
        rate(http_requests_total{status=~"5.."}[5m])
        / rate(http_requests_total[5m]) > 0.1
      for: 5m
      labels:
        severity: critical
      annotations:
        summary: High error rate detected
        description: Error rate is above 10% for 5 minutes
    
    - alert: HighLatency
      expr: |
        histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m])) > 1
      for: 5m
      labels:
        severity: warning
      annotations:
        summary: High latency detected
        description: P99 latency is above 1 second
```

---

# Practice Questions with Answers

## Beginner Level Questions

### Q1: What is the difference between a Pod and a Deployment?

**Answer:**

| Feature | Pod | Deployment |
|---------|-----|------------|
| Purpose | Smallest deployable unit | Manages ReplicaSets and Pods |
| Self-healing | No | Yes (recreates failed pods) |
| Scaling | Manual | Automatic |
| Rolling updates | No | Yes |
| Rollback | No | Yes |

```yaml
# Pod - single instance, no management
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx

# Deployment - managed, scalable, updatable
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    # Pod template
```

---

### Q2: How do you expose a Deployment to external traffic?

**Answer:**

```bash
# Method 1: NodePort Service
kubectl expose deployment nginx --type=NodePort --port=80

# Method 2: LoadBalancer Service (cloud providers)
kubectl expose deployment nginx --type=LoadBalancer --port=80

# Method 3: Ingress (recommended for HTTP/HTTPS)
```

```yaml
# Ingress example
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx-ingress
spec:
  ingressClassName: nginx
  rules:
  - host: nginx.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: nginx
            port:
              number: 80
```

---

### Q3: What is the purpose of namespaces?

**Answer:**
Namespaces provide:
1. **Resource isolation** - Separate resources for different teams/environments
2. **Access control** - Apply RBAC per namespace
3. **Resource quotas** - Limit resources per namespace
4. **Network policies** - Control traffic between namespaces

```bash
# Create namespace
kubectl create namespace development

# Deploy to namespace
kubectl apply -f deployment.yaml -n development

# Set default namespace
kubectl config set-context --current --namespace=development
```

---

### Q4: How do you update a Deployment's image?

**Answer:**

```bash
# Method 1: kubectl set image
kubectl set image deployment/nginx nginx=nginx:1.26

# Method 2: kubectl edit
kubectl edit deployment nginx

# Method 3: kubectl apply with updated manifest
kubectl apply -f deployment.yaml

# Method 4: kubectl patch
kubectl patch deployment nginx -p '{"spec":{"template":{"spec":{"containers":[{"name":"nginx","image":"nginx:1.26"}]}}}}'
```

Check rollout status:
```bash
kubectl rollout status deployment/nginx
kubectl rollout history deployment/nginx
```

---

### Q5: What is the difference between ConfigMap and Secret?

**Answer:**

| Feature | ConfigMap | Secret |
|---------|-----------|--------|
| Data type | Non-sensitive | Sensitive |
| Encoding | Plain text | Base64 encoded |
| Encryption at rest | No | Optional (with EncryptionConfiguration) |
| Use case | Config files, env vars | Passwords, API keys, certificates |

```yaml
# ConfigMap
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  LOG_LEVEL: "info"
  DATABASE_HOST: "postgres"

# Secret
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
type: Opaque
stringData:
  DB_PASSWORD: "secret123"
  API_KEY: "my-api-key"
```

---

## Intermediate Level Questions

### Q6: Explain the difference between StatefulSet and Deployment.

**Answer:**

| Feature | StatefulSet | Deployment |
|---------|-------------|------------|
| Pod Identity | Stable, ordered (pod-0, pod-1) | Random |
| DNS | Stable DNS names | Random |
| Storage | Persistent per pod | Shared or ephemeral |
| Scaling | Ordered (0→1→2) | Parallel |
| Deletion | Reverse order (2→1→0) | Parallel |
| Use case | Databases, stateful apps | Stateless apps |

```yaml
# StatefulSet with stable storage
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: postgres
spec:
  serviceName: postgres  # Required for DNS
  replicas: 3
  # Each pod gets:
  # - postgres-0.postgres.default.svc.cluster.local
  # - Its own PVC (data-postgres-0)
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 10Gi
```

---

### Q7: How do liveness and readiness probes differ?

**Answer:**

| Probe | Purpose | Failure Action |
|-------|---------|----------------|
| **Liveness** | Is container running? | Restart container |
| **Readiness** | Is container ready for traffic? | Remove from Service endpoints |
| **Startup** | Has container started? | Disable liveness/readiness until success |

```yaml
containers:
- name: app
  livenessProbe:
    httpGet:
      path: /healthz
      port: 8080
    initialDelaySeconds: 15
    periodSeconds: 10
    failureThreshold: 3
  
  readinessProbe:
    httpGet:
      path: /ready
      port: 8080
    initialDelaySeconds: 5
    periodSeconds: 5
    failureThreshold: 3
  
  startupProbe:
    httpGet:
      path: /healthz
      port: 8080
    failureThreshold: 30
    periodSeconds: 10
```

---

### Q8: Explain Kubernetes networking model.

**Answer:**

**Kubernetes Networking Requirements:**
1. All pods can communicate with each other without NAT
2. All nodes can communicate with all pods without NAT
3. The IP a pod sees for itself is the same IP others see

**Service Types:**
```
                     ┌─────────────────────┐
                     │   External Traffic  │
                     └──────────┬──────────┘
                                │
        ┌───────────────────────┴───────────────────────┐
        │                  LoadBalancer                 │
        │              (Cloud Provider)                 │
        └───────────────────────┬───────────────────────┘
                                │
        ┌───────────────────────┴───────────────────────┐
        │                   NodePort                    │
        │               (Node IP:30000-32767)           │
        └───────────────────────┬───────────────────────┘
                                │
        ┌───────────────────────┴───────────────────────┐
        │                  ClusterIP                    │
        │              (Internal Only)                  │
        └───────────────────────┬───────────────────────┘
                                │
                    ┌───────────┴───────────┐
                    │         Pods          │
                    └───────────────────────┘
```

---

### Q9: How do you implement blue-green deployments in Kubernetes?

**Answer:**

```yaml
# Blue deployment (current)
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-blue
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
      version: blue
  template:
    metadata:
      labels:
        app: myapp
        version: blue
---
# Green deployment (new)
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-green
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
      version: green
  template:
    metadata:
      labels:
        app: myapp
        version: green
---
# Service - switch by changing selector
apiVersion: v1
kind: Service
metadata:
  name: myapp
spec:
  selector:
    app: myapp
    version: blue  # Change to "green" to switch
  ports:
  - port: 80
```

**Switch traffic:**
```bash
kubectl patch service myapp -p '{"spec":{"selector":{"version":"green"}}}'
```

---

### Q10: Explain resource requests vs limits.

**Answer:**

| Feature | Requests | Limits |
|---------|----------|--------|
| Purpose | Scheduling guarantee | Maximum allowed |
| Scheduler | Uses for placement | Not considered |
| Enforcement | Guaranteed minimum | Hard cap |
| OOM | Won't cause eviction | Can cause OOM kill |

```yaml
resources:
  requests:
    cpu: 250m      # 0.25 CPU cores
    memory: 256Mi  # 256 MiB
  limits:
    cpu: 500m      # 0.5 CPU cores
    memory: 512Mi  # 512 MiB
```

**QoS Classes:**
- **Guaranteed**: requests == limits
- **Burstable**: requests < limits
- **BestEffort**: No requests or limits

---

## Advanced Level Questions

### Q11: Explain how HPA works internally.

**Answer:**

**HPA Control Loop:**
1. Metrics Server collects pod metrics every 15 seconds
2. HPA controller queries metrics every 15 seconds (configurable)
3. Calculates desired replicas using formula:
   ```
   desiredReplicas = ceil(currentReplicas × (currentMetric / targetMetric))
   ```
4. Scales deployment if needed

**Scaling Behavior:**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
spec:
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 300  # Wait 5 min before scaling down
      policies:
      - type: Percent
        value: 10
        periodSeconds: 60  # Scale down 10% per minute max
    scaleUp:
      stabilizationWindowSeconds: 0
      policies:
      - type: Pods
        value: 4
        periodSeconds: 15  # Add up to 4 pods every 15s
```

---

### Q12: How do Network Policies work?

**Answer:**

Network Policies are implemented by the CNI plugin (Calico, Cilium, etc.).

**Default behavior:** All pods can communicate freely.

**With Network Policy:**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: api-policy
spec:
  podSelector:
    matchLabels:
      app: api
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
    - namespaceSelector:
        matchLabels:
          name: monitoring
    ports:
    - port: 8080
  egress:
  - to:
    - podSelector:
        matchLabels:
          app: database
    ports:
    - port: 5432
```

**Key points:**
- Policies are additive (union of all matching policies)
- Empty selector `{}` matches all pods
- Policies are namespace-scoped
- Need CNI that supports Network Policies

---

### Q13: Explain Kubernetes RBAC.

**Answer:**

**Components:**
```
Subject (User/Group/ServiceAccount)
    ↓
RoleBinding / ClusterRoleBinding
    ↓
Role / ClusterRole
    ↓
API Resources (pods, services, etc.)
```

**Permission Model:**
```yaml
# Role defines permissions
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
rules:
- apiGroups: [""]           # Core API group
  resources: ["pods"]       # Resource type
  verbs: ["get", "list"]    # Allowed actions
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "list", "create", "update", "delete"]
  resourceNames: ["specific-deployment"]  # Optional: specific resources

# RoleBinding assigns role to subject
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
subjects:
- kind: User
  name: developer
  apiGroup: rbac.authorization.k8s.io
- kind: ServiceAccount
  name: app-sa
  namespace: default
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

**Verbs:** get, list, watch, create, update, patch, delete, deletecollection

---

### Q14: How do you troubleshoot a pod stuck in Pending state?

**Answer:**

**Common causes and solutions:**

1. **Insufficient Resources:**
```bash
kubectl describe pod <pod-name>
# Look for: "Insufficient cpu/memory"

# Solution: Add more nodes or reduce resource requests
```

2. **Node Selector/Affinity Mismatch:**
```bash
kubectl get nodes --show-labels
# Ensure nodes have required labels

# Solution: Update node labels or pod affinity
```

3. **Taints and Tolerations:**
```bash
kubectl describe node <node-name> | grep Taints
# Solution: Add tolerations to pod
```

4. **PVC Not Bound:**
```bash
kubectl get pvc
kubectl describe pvc <pvc-name>
# Solution: Create matching PV or StorageClass
```

5. **Resource Quota Exceeded:**
```bash
kubectl describe resourcequota -n <namespace>
# Solution: Increase quota or reduce requests
```

---

### Q15: Explain Pod Security Standards.

**Answer:**

**Three levels:**
1. **Privileged** - Unrestricted, for trusted workloads
2. **Baseline** - Minimal restrictions, prevents known privilege escalations
3. **Restricted** - Heavily restricted, security best practices

**Enforcement Modes:**
- **enforce** - Reject pods that violate policy
- **audit** - Log violations
- **warn** - Show warnings

```yaml
# Apply to namespace
apiVersion: v1
kind: Namespace
metadata:
  name: production
  labels:
    pod-security.kubernetes.io/enforce: restricted
    pod-security.kubernetes.io/enforce-version: latest
    pod-security.kubernetes.io/audit: restricted
    pod-security.kubernetes.io/warn: restricted
```

**Restricted pod example:**
```yaml
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    fsGroup: 1000
    seccompProfile:
      type: RuntimeDefault
  containers:
  - name: app
    securityContext:
      allowPrivilegeEscalation: false
      readOnlyRootFilesystem: true
      capabilities:
        drop:
        - ALL
```

---

## Expert Level Questions

### Q16: How does the Kubernetes scheduler work?

**Answer:**

**Scheduling Process:**

1. **Filtering (Predicates):**
   - NodeSelector match
   - Resource availability
   - Taints/Tolerations
   - Node affinity
   - Volume constraints

2. **Scoring (Priorities):**
   - Resource balance
   - Pod affinity/anti-affinity
   - Node affinity preferences
   - Image locality

3. **Binding:**
   - Scheduler binds pod to selected node

**Custom Scheduling:**
```yaml
# Pod with scheduler name
spec:
  schedulerName: my-custom-scheduler

# Scheduler Extender
apiVersion: kubescheduler.config.k8s.io/v1
kind: KubeSchedulerConfiguration
extenders:
- urlPrefix: "http://my-scheduler-extender:8888"
  filterVerb: "filter"
  prioritizeVerb: "prioritize"
```

---

### Q17: Explain etcd and its role in Kubernetes.

**Answer:**

**etcd is:**
- Distributed key-value store
- Stores all cluster state
- Uses Raft consensus algorithm
- Provides strong consistency

**What's stored:**
- Pods, Services, ConfigMaps, Secrets
- RBAC rules
- Custom resources
- Lease information

**Best practices:**
```bash
# Backup etcd
ETCDCTL_API=3 etcdctl snapshot save backup.db \
  --endpoints=https://127.0.0.1:2379 \
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \
  --cert=/etc/kubernetes/pki/etcd/server.crt \
  --key=/etc/kubernetes/pki/etcd/server.key

# Restore
ETCDCTL_API=3 etcdctl snapshot restore backup.db

# Check cluster health
ETCDCTL_API=3 etcdctl endpoint health
```

**Production recommendations:**
- Odd number of members (3 or 5)
- Fast SSD storage
- Dedicated nodes
- Regular backups

---

### Q18: How do you implement canary deployments?

**Answer:**

**Method 1: Multiple Deployments with Service Weights (Istio)**
```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: myapp
spec:
  hosts:
  - myapp
  http:
  - route:
    - destination:
        host: myapp
        subset: stable
      weight: 90
    - destination:
        host: myapp
        subset: canary
      weight: 10
```

**Method 2: Argo Rollouts**
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Rollout
metadata:
  name: myapp
spec:
  replicas: 10
  strategy:
    canary:
      steps:
      - setWeight: 10
      - pause: {duration: 1h}
      - setWeight: 20
      - pause: {duration: 1h}
      - setWeight: 50
      - pause: {duration: 1h}
      - setWeight: 100
      analysis:
        templates:
        - templateName: success-rate
```

**Method 3: Native Kubernetes**
```yaml
# 90% stable
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-stable
spec:
  replicas: 9
  
# 10% canary
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-canary
spec:
  replicas: 1
```

---

### Q19: How do you implement zero-trust security in Kubernetes?

**Answer:**

**1. Network Policies (Default Deny):**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-all
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  - Egress
```

**2. Mutual TLS (Service Mesh):**
```yaml
# Istio PeerAuthentication
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: istio-system
spec:
  mtls:
    mode: STRICT
```

**3. Pod Security Standards:**
```yaml
# Namespace-level enforcement
labels:
  pod-security.kubernetes.io/enforce: restricted
```

**4. RBAC (Least Privilege):**
```yaml
# Minimal permissions
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list"]
  # No wildcards
```

**5. Runtime Security (Falco):**
```yaml
# Detect suspicious behavior
- rule: Terminal shell in container
  condition: >
    spawned_process and
    container and
    shell_procs
  output: "Shell spawned in container"
```

---

### Q20: Explain Kubernetes API aggregation layer.

**Answer:**

API aggregation allows extending the Kubernetes API with custom API servers.

**Use cases:**
- Custom metrics APIs
- Service catalog
- Custom resource implementations

**APIService:**
```yaml
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  name: v1beta1.metrics.k8s.io
spec:
  service:
    name: metrics-server
    namespace: kube-system
  group: metrics.k8s.io
  version: v1beta1
  insecureSkipTLSVerify: true
  groupPriorityMinimum: 100
  versionPriority: 100
```

**Request flow:**
```
kubectl → API Server → Aggregation Layer → Custom API Server
```

---

## Solution Architect Level Questions

### Q21: Design a multi-tenant Kubernetes cluster.

**Answer:**

**Isolation Layers:**

1. **Namespace Isolation:**
```yaml
# Per-tenant namespace
apiVersion: v1
kind: Namespace
metadata:
  name: tenant-a
  labels:
    tenant: a
```

2. **Resource Quotas:**
```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: tenant-quota
  namespace: tenant-a
spec:
  hard:
    requests.cpu: "10"
    requests.memory: "20Gi"
    limits.cpu: "20"
    limits.memory: "40Gi"
    pods: "50"
```

3. **Network Isolation:**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: isolate-tenant
  namespace: tenant-a
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          tenant: a
  egress:
  - to:
    - namespaceSelector:
        matchLabels:
          tenant: a
```

4. **RBAC:**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: tenant-admin
  namespace: tenant-a
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: tenant-a-admins
  namespace: tenant-a
subjects:
- kind: Group
  name: tenant-a-admins
roleRef:
  kind: Role
  name: tenant-admin
```

5. **Pod Security:**
```yaml
# Enforce restricted security
labels:
  pod-security.kubernetes.io/enforce: restricted
```

---

### Q22: Design a disaster recovery strategy for Kubernetes.

**Answer:**

**1. Backup Strategy:**
```bash
# Velero scheduled backup
velero schedule create daily \
  --schedule="0 2 * * *" \
  --ttl 720h

# Include specific namespaces
velero backup create app-backup \
  --include-namespaces production,staging \
  --storage-location aws-backup
```

**2. Multi-Region Setup:**
```yaml
# ArgoCD ApplicationSet for multi-cluster
apiVersion: argoproj.io/v1alpha1
kind: ApplicationSet
metadata:
  name: app-set
spec:
  generators:
  - clusters:
      selector:
        matchLabels:
          region: primary
  - clusters:
      selector:
        matchLabels:
          region: dr
```

**3. Database Replication:**
```yaml
# CrunchyData PostgreSQL Operator
apiVersion: postgres-operator.crunchydata.com/v1beta1
kind: PostgresCluster
metadata:
  name: postgres
spec:
  standby:
    enabled: true
    repoName: repo2
    host: primary-cluster-pgbouncer.primary.svc
```

**4. Recovery Runbook:**
```bash
#!/bin/bash
# 1. Verify primary is down
# 2. Promote DR cluster
kubectl config use-context dr-cluster
kubectl scale deployment app --replicas=10

# 3. Update DNS
aws route53 change-resource-record-sets ...

# 4. Verify services
kubectl get pods -n production

# 5. Notify stakeholders
```

---

### Q23: Design a GitOps workflow for Kubernetes.

**Answer:**

**Architecture:**
```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Dev       │───▶│   Git Repo  │───▶│   ArgoCD    │
│   Push      │    │   (Source)  │    │   (Sync)    │
└─────────────┘    └─────────────┘    └─────────────┘
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │    Kubernetes   │
                                    │    Cluster      │
                                    └─────────────────┘
```

**Repository Structure:**
```
├── apps/
│   ├── base/
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   └── kustomization.yaml
│   └── overlays/
│       ├── development/
│       ├── staging/
│       └── production/
├── infrastructure/
│   ├── monitoring/
│   ├── ingress/
│   └── cert-manager/
└── clusters/
    ├── development/
    ├── staging/
    └── production/
```

**ArgoCD Application:**
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: myapp
  namespace: argocd
  finalizers:
  - resources-finalizer.argocd.argoproj.io
spec:
  project: default
  source:
    repoURL: https://github.com/org/k8s-manifests.git
    targetRevision: main
    path: apps/overlays/production
  destination:
    server: https://kubernetes.default.svc
    namespace: production
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
    - CreateNamespace=true
```

---

### Q24: How do you optimize Kubernetes costs?

**Answer:**

**1. Right-sizing Resources:**
```yaml
# Use VPA recommendations
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
  name: app-vpa
spec:
  targetRef:
    kind: Deployment
    name: app
  updatePolicy:
    updateMode: "Off"  # Just recommendations
```

**2. Cluster Autoscaling:**
```yaml
# Node pool with autoscaling
apiVersion: karpenter.sh/v1alpha5
kind: Provisioner
metadata:
  name: default
spec:
  requirements:
  - key: karpenter.sh/capacity-type
    operator: In
    values: ["spot", "on-demand"]
  - key: kubernetes.io/instance-type
    operator: In
    values: ["m5.large", "m5.xlarge"]
  limits:
    resources:
      cpu: 100
```

**3. Spot Instances:**
```yaml
# Tolerate spot termination
tolerations:
- key: "kubernetes.azure.com/scalesetpriority"
  operator: "Equal"
  value: "spot"
  effect: "NoSchedule"
```

**4. Pod Priority:**
```yaml
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: critical
value: 1000000
globalDefault: false
description: "Critical workloads"
---
spec:
  priorityClassName: critical
```

**5. Resource Quotas:**
```yaml
# Prevent over-provisioning
apiVersion: v1
kind: ResourceQuota
metadata:
  name: team-quota
spec:
  hard:
    requests.cpu: "50"
    requests.memory: "100Gi"
```

---

### Q25: Design a complete production Kubernetes platform.

**Answer:**

**Platform Components:**

```
┌──────────────────────────────────────────────────────────────────┐
│                         Platform Layer                           │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐             │
│  │ Ingress  │ │  Cert    │ │ External │ │  Vault   │             │
│  │ (NGINX)  │ │ Manager  │ │   DNS    │ │(Secrets) │             │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘             │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐             │
│  │Prometheus│ │ Grafana  │ │  Loki    │ │  Jaeger  │             │
│  │ (Metrics)│ │  (Viz)   │ │ (Logs)   │ │(Tracing) │             │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘             │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐                          │
│  │  ArgoCD  │ │  Velero  │ │  Falco   │                          │
│  │ (GitOps) │ │ (Backup) │ │(Security)│                          │
│  └──────────┘ └──────────┘ └──────────┘                          │
└──────────────────────────────────────────────────────────────────┘
                              │
┌──────────────────────────────────────────────────────────────────┐
│                      Application Workloads                       │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │                    Production Namespace                     │ │
│  │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐            │ │
│  │  │Frontend │ │   API   │ │ Workers │ │Database │            │ │
│  │  │ (HPA)   │ │ (HPA)   │ │ (KEDA)  │ │(Stateful)│           │ │
│  │  └─────────┘ └─────────┘ └─────────┘ └─────────┘            │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │                    Staging Namespace                        │ │
│  └─────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────────┘
                              │
┌──────────────────────────────────────────────────────────────────┐
│                      Kubernetes Cluster                          │
│  ┌──────────────┐ ┌──────────────┐ ┌──────────────┐              │
│  │Control Plane │ │ Node Pool 1  │ │ Node Pool 2  │              │
│  │   (HA: 3)    │ │  (General)   │ │   (Compute)  │              │
│  └──────────────┘ └──────────────┘ └──────────────┘              │
└──────────────────────────────────────────────────────────────────┘
```

**Helm Chart for Platform:**
```yaml
# values.yaml
ingress-nginx:
  enabled: true
  controller:
    replicaCount: 2
    
cert-manager:
  enabled: true
  installCRDs: true
  
prometheus-stack:
  enabled: true
  grafana:
    adminPassword: ${GRAFANA_PASSWORD}
  prometheus:
    retention: 30d
    
argocd:
  enabled: true
  server:
    ingress:
      enabled: true
      
velero:
  enabled: true
  schedules:
    daily:
      schedule: "0 2 * * *"
```

---

## Summary

This comprehensive Kubernetes guide covers:

- **Beginner:** Core concepts, pods, deployments, services, namespaces
- **Intermediate:** Storage, ingress, StatefulSets, Jobs, resource management
- **Advanced:** RBAC, Network Policies, HPA/VPA, affinity, taints
- **Expert:** CRDs, Operators, Admission Controllers, Service Mesh
- **Solution Architect:** Multi-cluster, DR, GitOps, Security, Platform Design

**Key Takeaways:**
1. Start with Deployments, not naked Pods
2. Always set resource requests and limits
3. Use namespaces for logical separation
4. Implement RBAC with least privilege
5. Use Network Policies for security
6. Implement proper health checks
7. Use GitOps for declarative deployments
8. Plan for disaster recovery from day one
9. Monitor everything (metrics, logs, traces)
10. Automate with Operators when appropriate

---

*Document Version: 1.0*
*Last Updated: January 2026*
