# Docker Mastery Guide
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

## 1. Introduction to Docker

### What is Docker?
Docker is a platform for developing, shipping, and running applications in containers. Containers are lightweight, standalone, executable packages that include everything needed to run an application.

### Key Concepts

| Concept | Description |
|---------|-------------|
| **Image** | A read-only template with instructions for creating a container |
| **Container** | A runnable instance of an image |
| **Dockerfile** | A text file with instructions to build an image |
| **Registry** | A repository for storing Docker images (e.g., Docker Hub) |
| **Volume** | Persistent data storage for containers |
| **Network** | Communication channel between containers |

### Docker vs Virtual Machines

| Feature | Docker Containers | Virtual Machines |
|---------|------------------|------------------|
| Isolation | Process-level | Hardware-level |
| Size | Megabytes | Gigabytes |
| Boot Time | Seconds | Minutes |
| Resource Usage | Minimal | Significant |
| OS | Shares host kernel | Full OS per VM |

### Installation

```bash
# macOS
brew install --cask docker

# Or download Docker Desktop from docker.com

# Verify installation
docker --version
docker info
```

---

## 2. Docker Basic Commands

### Container Lifecycle

```bash
# Run a container
docker run hello-world

# Run with options
docker run -d --name my-nginx -p 8080:80 nginx
# -d: detached mode (background)
# --name: container name
# -p: port mapping (host:container)

# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Stop a container
docker stop my-nginx

# Start a stopped container
docker start my-nginx

# Restart a container
docker restart my-nginx

# Remove a container
docker rm my-nginx

# Remove running container (force)
docker rm -f my-nginx

# Remove all stopped containers
docker container prune
```

### Image Commands

```bash
# Pull an image
docker pull nginx:latest

# List images
docker images

# Remove an image
docker rmi nginx:latest

# Remove unused images
docker image prune

# Remove all unused images
docker image prune -a

# Tag an image
docker tag nginx:latest myregistry/nginx:v1

# Push an image
docker push myregistry/nginx:v1

# Search for images
docker search nginx
```

### Container Interaction

```bash
# Execute command in running container
docker exec -it my-nginx bash
# -i: interactive
# -t: pseudo-TTY

# View container logs
docker logs my-nginx

# Follow logs
docker logs -f my-nginx

# View last N lines
docker logs --tail 100 my-nginx

# Inspect container details
docker inspect my-nginx

# View container resource usage
docker stats my-nginx

# Copy files to/from container
docker cp file.txt my-nginx:/path/to/destination
docker cp my-nginx:/path/to/file.txt ./local-file.txt
```

---

## 3. Dockerfile Basics

### Basic Dockerfile

```dockerfile
# Base image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy application code
COPY . .

# Expose port
EXPOSE 3000

# Define environment variable
ENV NODE_ENV=production

# Run the application
CMD ["node", "server.js"]
```

### Dockerfile Instructions

| Instruction | Description |
|-------------|-------------|
| `FROM` | Base image |
| `WORKDIR` | Set working directory |
| `COPY` | Copy files from host to image |
| `ADD` | Copy files (supports URLs and tar extraction) |
| `RUN` | Execute commands during build |
| `CMD` | Default command to run (can be overridden) |
| `ENTRYPOINT` | Main command (harder to override) |
| `ENV` | Set environment variables |
| `EXPOSE` | Document which ports the container listens on |
| `VOLUME` | Create mount point for volumes |
| `ARG` | Build-time variables |
| `LABEL` | Add metadata |
| `USER` | Set user for RUN, CMD, ENTRYPOINT |

### Building Images

```bash
# Build an image
docker build -t myapp:v1 .

# Build with different Dockerfile
docker build -t myapp:v1 -f Dockerfile.prod .

# Build with build arguments
docker build --build-arg VERSION=1.0 -t myapp:v1 .

# Build without cache
docker build --no-cache -t myapp:v1 .

# Build for multiple platforms
docker buildx build --platform linux/amd64,linux/arm64 -t myapp:v1 .
```

---

## 4. Docker Volumes

### Types of Volumes

1. **Named Volumes** - Managed by Docker
2. **Bind Mounts** - Direct mapping to host filesystem
3. **tmpfs Mounts** - Stored in memory only

### Volume Commands

```bash
# Create a volume
docker volume create my-volume

# List volumes
docker volume ls

# Inspect a volume
docker volume inspect my-volume

# Remove a volume
docker volume rm my-volume

# Remove unused volumes
docker volume prune
```

### Using Volumes

```bash
# Named volume
docker run -d -v my-volume:/app/data nginx

# Bind mount
docker run -d -v /host/path:/container/path nginx

# Read-only volume
docker run -d -v my-volume:/app/data:ro nginx

# tmpfs mount
docker run -d --tmpfs /app/temp nginx
```

### Example: Persistent Database

```bash
# Create volume for PostgreSQL data
docker volume create postgres-data

# Run PostgreSQL with persistent data
docker run -d \
  --name postgres \
  -e POSTGRES_PASSWORD=secret \
  -v postgres-data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres:15
```

---

## 5. Docker Networking

### Network Types

| Type | Description |
|------|-------------|
| `bridge` | Default network for containers on same host |
| `host` | Use host's network directly |
| `none` | No networking |
| `overlay` | Multi-host networking (Swarm) |
| `macvlan` | Assign MAC address to container |

### Network Commands

```bash
# List networks
docker network ls

# Create a network
docker network create my-network

# Create with specific driver
docker network create --driver bridge my-bridge

# Inspect a network
docker network inspect my-network

# Connect container to network
docker network connect my-network my-container

# Disconnect container from network
docker network disconnect my-network my-container

# Remove a network
docker network rm my-network
```

### Container Communication

```bash
# Create a network
docker network create app-network

# Run containers on the same network
docker run -d --name db --network app-network postgres:15
docker run -d --name web --network app-network nginx

# Containers can communicate using container names
# From web container: ping db
```

---

## 6. Docker Compose Basics

### docker-compose.yml Structure

```yaml
version: '3.8'

services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - ./html:/usr/share/nginx/html
    networks:
      - app-network
    depends_on:
      - api

  api:
    build: ./api
    ports:
      - "3000:3000"
    environment:
      - DATABASE_URL=postgres://user:pass@db:5432/mydb
    networks:
      - app-network
    depends_on:
      - db

  db:
    image: postgres:15
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
      - POSTGRES_DB=mydb
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - app-network

networks:
  app-network:
    driver: bridge

volumes:
  db-data:
```

### Docker Compose Commands

```bash
# Start services
docker-compose up

# Start in detached mode
docker-compose up -d

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v

# Build and start
docker-compose up --build

# View logs
docker-compose logs

# Follow logs for specific service
docker-compose logs -f api

# Scale a service
docker-compose up -d --scale web=3

# Execute command in service
docker-compose exec api bash

# View running services
docker-compose ps
```

---

# Intermediate Level

## 7. Multi-Stage Builds

### Why Multi-Stage Builds?
- Reduce final image size
- Separate build and runtime environments
- Keep build tools out of production image

### Example: Go Application

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:3.18

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .

# Run as non-root user
RUN adduser -D appuser
USER appuser

EXPOSE 8080

CMD ["./main"]
```

### Example: Node.js Application

```dockerfile
# Build stage
FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

# Runtime stage
FROM node:18-alpine

WORKDIR /app

# Copy only production dependencies
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./

USER node

EXPOSE 3000

CMD ["node", "dist/index.js"]
```

### Example: React Application

```dockerfile
# Build stage
FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

# Runtime stage - Using nginx
FROM nginx:alpine

# Copy custom nginx config
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built assets
COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

---

## 8. Docker Image Optimization

### Best Practices

#### 1. Use Minimal Base Images

```dockerfile
# Instead of
FROM ubuntu:22.04

# Use
FROM alpine:3.18
# Or distroless
FROM gcr.io/distroless/static-debian11
```

#### 2. Reduce Layer Count

```dockerfile
# Bad - Multiple RUN commands
RUN apt-get update
RUN apt-get install -y package1
RUN apt-get install -y package2
RUN apt-get clean

# Good - Single RUN command
RUN apt-get update && \
    apt-get install -y \
        package1 \
        package2 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
```

#### 3. Order Layers by Change Frequency

```dockerfile
# Dependencies change less frequently
COPY package*.json ./
RUN npm install

# Source code changes more frequently
COPY . .
```

#### 4. Use .dockerignore

```
# .dockerignore
node_modules
npm-debug.log
Dockerfile*
docker-compose*
.git
.gitignore
.env
*.md
test
coverage
.vscode
```

#### 5. Don't Install Unnecessary Packages

```dockerfile
# Install only production dependencies
RUN npm ci --only=production

# Skip recommended packages in apt
RUN apt-get install -y --no-install-recommends package
```

### Analyzing Image Size

```bash
# View image size
docker images myapp

# View layer sizes
docker history myapp:latest

# Use dive to analyze layers
dive myapp:latest
```

---

## 9. Container Security

### Security Best Practices

#### 1. Run as Non-Root User

```dockerfile
# Create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Change ownership of app files
COPY --chown=appuser:appgroup . .

# Switch to non-root user
USER appuser
```

#### 2. Use Read-Only Filesystem

```bash
docker run --read-only --tmpfs /tmp myapp
```

#### 3. Drop Capabilities

```bash
docker run --cap-drop ALL --cap-add NET_BIND_SERVICE myapp
```

#### 4. Limit Resources

```bash
docker run \
  --memory="512m" \
  --memory-swap="512m" \
  --cpus="0.5" \
  --pids-limit 100 \
  myapp
```

#### 5. Scan Images for Vulnerabilities

```bash
# Docker Scout
docker scout cves myapp:latest

# Trivy
trivy image myapp:latest

# Snyk
snyk container test myapp:latest
```

#### 6. Use Specific Tags

```dockerfile
# Bad - Tag can change
FROM node:latest

# Good - Specific version
FROM node:18.19.0-alpine3.18
```

#### 7. Verify Image Integrity

```bash
# Enable Docker Content Trust
export DOCKER_CONTENT_TRUST=1

# Pull verified images
docker pull nginx:latest
```

---

## 10. Docker Compose Advanced

### Environment Variables

```yaml
version: '3.8'

services:
  api:
    image: myapi:latest
    environment:
      - NODE_ENV=production
      - DATABASE_URL=${DATABASE_URL}
    env_file:
      - .env
      - .env.local
```

### Health Checks

```yaml
services:
  web:
    image: nginx:latest
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

### Profiles

```yaml
version: '3.8'

services:
  web:
    image: nginx:latest
    profiles:
      - production
      - development

  debug:
    image: busybox
    profiles:
      - development
```

```bash
# Start only production profile
docker-compose --profile production up

# Start development profile
docker-compose --profile development up
```

### Extend and Override

```yaml
# docker-compose.yml (base)
version: '3.8'

services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
```

```yaml
# docker-compose.override.yml (development)
version: '3.8'

services:
  web:
    volumes:
      - ./html:/usr/share/nginx/html
    ports:
      - "8080:80"
```

```yaml
# docker-compose.prod.yml (production)
version: '3.8'

services:
  web:
    deploy:
      replicas: 3
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
```

```bash
# Development (uses override automatically)
docker-compose up

# Production
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up
```

### Secrets

```yaml
version: '3.8'

services:
  db:
    image: postgres:15
    secrets:
      - db_password
    environment:
      - POSTGRES_PASSWORD_FILE=/run/secrets/db_password

secrets:
  db_password:
    file: ./secrets/db_password.txt
```

---

## 11. Docker Registry

### Running Private Registry

```bash
# Run registry
docker run -d \
  -p 5000:5000 \
  --name registry \
  -v registry-data:/var/lib/registry \
  registry:2

# Tag and push image
docker tag myapp:latest localhost:5000/myapp:latest
docker push localhost:5000/myapp:latest

# Pull from private registry
docker pull localhost:5000/myapp:latest
```

### Registry with Authentication

```bash
# Create password file
mkdir auth
docker run --entrypoint htpasswd \
  httpd:2 -Bbn username password > auth/htpasswd

# Run secured registry
docker run -d \
  -p 5000:5000 \
  --name registry \
  -v $(pwd)/auth:/auth \
  -e "REGISTRY_AUTH=htpasswd" \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e "REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd" \
  -v registry-data:/var/lib/registry \
  registry:2

# Login to registry
docker login localhost:5000
```

### Registry with TLS

```yaml
# docker-compose.yml for secure registry
version: '3.8'

services:
  registry:
    image: registry:2
    ports:
      - "5000:5000"
    environment:
      REGISTRY_HTTP_TLS_CERTIFICATE: /certs/domain.crt
      REGISTRY_HTTP_TLS_KEY: /certs/domain.key
      REGISTRY_AUTH: htpasswd
      REGISTRY_AUTH_HTPASSWD_REALM: Registry Realm
      REGISTRY_AUTH_HTPASSWD_PATH: /auth/htpasswd
    volumes:
      - ./certs:/certs
      - ./auth:/auth
      - registry-data:/var/lib/registry

volumes:
  registry-data:
```

---

## 12. Logging and Monitoring

### Docker Logging Drivers

```bash
# View available logging drivers
docker info --format '{{.Plugins.Log}}'

# Run with specific logging driver
docker run -d \
  --log-driver json-file \
  --log-opt max-size=10m \
  --log-opt max-file=3 \
  nginx

# Run with syslog
docker run -d \
  --log-driver syslog \
  --log-opt syslog-address=udp://logs.example.com:514 \
  nginx
```

### Docker Compose Logging

```yaml
version: '3.8'

services:
  web:
    image: nginx:latest
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
```

### Container Metrics

```bash
# View live stats
docker stats

# View stats for specific container
docker stats my-container

# Format output
docker stats --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}"
```

### Prometheus Metrics

```yaml
# docker-compose.yml with cAdvisor
version: '3.8'

services:
  cadvisor:
    image: gcr.io/cadvisor/cadvisor:latest
    container_name: cadvisor
    ports:
      - "8080:8080"
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:ro
      - /sys:/sys:ro
      - /var/lib/docker/:/var/lib/docker:ro

  prometheus:
    image: prom/prometheus:latest
    container_name: prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
```

---

# Advanced Level

## 13. Docker Swarm

### Initialize Swarm

```bash
# Initialize swarm on manager node
docker swarm init --advertise-addr <manager-ip>

# Get worker join token
docker swarm join-token worker

# Get manager join token
docker swarm join-token manager

# Join as worker
docker swarm join --token <token> <manager-ip>:2377

# Leave swarm
docker swarm leave
docker swarm leave --force  # For manager
```

### Service Management

```bash
# Create a service
docker service create \
  --name web \
  --replicas 3 \
  --publish 80:80 \
  nginx

# List services
docker service ls

# View service tasks
docker service ps web

# Scale service
docker service scale web=5

# Update service
docker service update \
  --image nginx:latest \
  --update-parallelism 2 \
  --update-delay 10s \
  web

# Remove service
docker service rm web
```

### Stack Deployment

```yaml
# stack.yml
version: '3.8'

services:
  web:
    image: nginx:latest
    deploy:
      replicas: 3
      update_config:
        parallelism: 2
        delay: 10s
        failure_action: rollback
      rollback_config:
        parallelism: 1
        delay: 10s
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 3
      placement:
        constraints:
          - node.role == worker
      resources:
        limits:
          cpus: '0.5'
          memory: 256M
        reservations:
          cpus: '0.1'
          memory: 128M
    ports:
      - "80:80"
    networks:
      - webnet

networks:
  webnet:
    driver: overlay
```

```bash
# Deploy stack
docker stack deploy -c stack.yml myapp

# List stacks
docker stack ls

# List stack services
docker stack services myapp

# Remove stack
docker stack rm myapp
```

---

## 14. Docker Buildx and Multi-Platform

### Setup Buildx

```bash
# Create builder
docker buildx create --name mybuilder --use

# Bootstrap builder
docker buildx inspect --bootstrap

# List builders
docker buildx ls
```

### Multi-Platform Build

```bash
# Build for multiple platforms
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t myapp:latest \
  --push \
  .
```

### Advanced Buildx Features

```bash
# Build with cache export
docker buildx build \
  --cache-from type=registry,ref=myrepo/cache \
  --cache-to type=registry,ref=myrepo/cache \
  -t myapp:latest \
  .

# Build with SSH forwarding
docker buildx build \
  --ssh default \
  -t myapp:latest \
  .

# Build with secrets
docker buildx build \
  --secret id=mysecret,src=secret.txt \
  -t myapp:latest \
  .
```

### Dockerfile with Secrets

```dockerfile
FROM alpine

# Use secret during build (not stored in layer)
RUN --mount=type=secret,id=mysecret \
    cat /run/secrets/mysecret > /app/config
```

---

## 15. Container Orchestration Patterns

### Sidecar Pattern

```yaml
version: '3.8'

services:
  app:
    image: myapp:latest
    volumes:
      - logs:/var/log/app

  log-shipper:
    image: fluent/fluentd:latest
    volumes:
      - logs:/var/log/app:ro
    depends_on:
      - app

volumes:
  logs:
```

### Ambassador Pattern

```yaml
version: '3.8'

services:
  app:
    image: myapp:latest
    environment:
      - DB_HOST=db-ambassador

  db-ambassador:
    image: haproxy:latest
    volumes:
      - ./haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg
```

### Adapter Pattern

```yaml
version: '3.8'

services:
  legacy-app:
    image: legacy:latest
    
  adapter:
    image: adapter:latest
    environment:
      - LEGACY_HOST=legacy-app
    ports:
      - "8080:8080"
```

---

## 16. Docker API and SDK

### Docker API Examples

```bash
# Enable Docker API
# Add to daemon.json
{
  "hosts": ["unix:///var/run/docker.sock", "tcp://0.0.0.0:2375"]
}

# List containers via API
curl http://localhost:2375/containers/json

# Create container via API
curl -X POST http://localhost:2375/containers/create \
  -H "Content-Type: application/json" \
  -d '{"Image": "nginx:latest"}'

# Start container
curl -X POST http://localhost:2375/containers/<id>/start
```

### Go SDK Example

```go
package main

import (
    "context"
    "fmt"
    "io"
    "os"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }
    
    // List containers
    containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
    if err != nil {
        panic(err)
    }
    
    for _, container := range containers {
        fmt.Printf("%s: %s\n", container.ID[:12], container.Image)
    }
    
    // Create container
    resp, err := cli.ContainerCreate(ctx,
        &container.Config{
            Image: "nginx:latest",
        },
        nil, nil, nil, "my-nginx")
    if err != nil {
        panic(err)
    }
    
    // Start container
    err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Container %s started\n", resp.ID[:12])
}
```

### Python SDK Example

```python
import docker

client = docker.from_env()

# List containers
for container in client.containers.list():
    print(f"{container.id[:12]}: {container.image.tags}")

# Run container
container = client.containers.run(
    "nginx:latest",
    detach=True,
    name="my-nginx",
    ports={'80/tcp': 8080}
)

print(f"Container {container.id[:12]} started")

# Get logs
logs = container.logs()
print(logs.decode('utf-8'))

# Stop and remove
container.stop()
container.remove()
```

---

## 17. Docker Performance Tuning

### Resource Limits

```yaml
version: '3.8'

services:
  app:
    image: myapp:latest
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: 2G
        reservations:
          cpus: '0.5'
          memory: 512M
    # For non-swarm deployments
    mem_limit: 2g
    mem_reservation: 512m
    cpus: 2
    cpu_shares: 1024
```

### Storage Driver Optimization

```bash
# Check current storage driver
docker info | grep "Storage Driver"

# Common drivers:
# - overlay2 (recommended)
# - devicemapper
# - btrfs
# - zfs

# Configure in daemon.json
{
  "storage-driver": "overlay2",
  "storage-opts": [
    "overlay2.override_kernel_check=true"
  ]
}
```

### Network Performance

```bash
# Use host network for performance-critical containers
docker run --network host myapp

# Increase default MTU if needed
docker network create --opt com.docker.network.driver.mtu=9000 mynetwork
```

### Docker Daemon Tuning

```json
// /etc/docker/daemon.json
{
  "storage-driver": "overlay2",
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  },
  "default-ulimits": {
    "nofile": {
      "Name": "nofile",
      "Hard": 65536,
      "Soft": 65536
    }
  },
  "max-concurrent-downloads": 10,
  "max-concurrent-uploads": 10
}
```

---

# Expert Level

## 18. Docker Internals

### Container Primitives

#### Namespaces
- **PID**: Process isolation
- **NET**: Network isolation
- **MNT**: Mount point isolation
- **UTS**: Hostname isolation
- **IPC**: Inter-process communication isolation
- **USER**: User ID isolation

#### Control Groups (cgroups)
- CPU limits
- Memory limits
- I/O limits
- Network bandwidth

#### Union File Systems
- OverlayFS (recommended)
- AUFS (legacy)

### Inspecting Container Internals

```bash
# View container namespaces
ls -la /proc/<container-pid>/ns/

# Enter container namespace
nsenter -t <container-pid> -n ip addr

# View cgroup limits
cat /sys/fs/cgroup/memory/docker/<container-id>/memory.limit_in_bytes

# View overlay filesystem
docker inspect <container-id> --format '{{.GraphDriver.Data}}'
```

### Custom Seccomp Profile

```json
{
  "defaultAction": "SCMP_ACT_ERRNO",
  "architectures": [
    "SCMP_ARCH_X86_64"
  ],
  "syscalls": [
    {
      "names": ["read", "write", "open", "close", "stat", "fstat"],
      "action": "SCMP_ACT_ALLOW"
    }
  ]
}
```

```bash
docker run --security-opt seccomp=seccomp.json myapp
```

### AppArmor Profile

```bash
# Create AppArmor profile
sudo nano /etc/apparmor.d/docker-myapp

# Load profile
sudo apparmor_parser -r /etc/apparmor.d/docker-myapp

# Run with profile
docker run --security-opt apparmor=docker-myapp myapp
```

---

## 19. Docker in CI/CD

### GitHub Actions

```yaml
name: Docker Build and Push

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Login to Registry
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          tags: |
            type=sha
            type=ref,event=branch
            type=semver,pattern={{version}}

      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          platforms: linux/amd64,linux/arm64
          push: ${{ github.event_name != 'pull_request' }}
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

### GitLab CI

```yaml
stages:
  - build
  - test
  - deploy

variables:
  DOCKER_TLS_CERTDIR: "/certs"
  IMAGE_TAG: $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA

build:
  stage: build
  image: docker:24
  services:
    - docker:24-dind
  before_script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  script:
    - docker build -t $IMAGE_TAG .
    - docker push $IMAGE_TAG

test:
  stage: test
  image: $IMAGE_TAG
  script:
    - npm test

deploy:
  stage: deploy
  script:
    - docker pull $IMAGE_TAG
    - docker tag $IMAGE_TAG $CI_REGISTRY_IMAGE:latest
    - docker push $CI_REGISTRY_IMAGE:latest
  only:
    - main
```

### Jenkins Pipeline

```groovy
pipeline {
    agent any
    
    environment {
        REGISTRY = 'docker.io'
        IMAGE = 'myorg/myapp'
        REGISTRY_CREDENTIAL = 'docker-hub'
    }
    
    stages {
        stage('Build') {
            steps {
                script {
                    dockerImage = docker.build("${IMAGE}:${BUILD_NUMBER}")
                }
            }
        }
        
        stage('Test') {
            steps {
                script {
                    dockerImage.inside {
                        sh 'npm test'
                    }
                }
            }
        }
        
        stage('Push') {
            steps {
                script {
                    docker.withRegistry("https://${REGISTRY}", REGISTRY_CREDENTIAL) {
                        dockerImage.push()
                        dockerImage.push('latest')
                    }
                }
            }
        }
        
        stage('Deploy') {
            steps {
                sh 'docker-compose -f docker-compose.prod.yml up -d'
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
    }
}
```

---

## 20. Docker Desktop Extensions

### Creating an Extension

```json
// metadata.json
{
  "name": "my-extension",
  "version": "1.0.0",
  "description": "My Docker Desktop Extension",
  "icon": "docker.svg",
  "ui": {
    "dashboard-tab": {
      "title": "My Extension",
      "src": "index.html",
      "root": "ui"
    }
  },
  "vm": {
    "composefile": "docker-compose.yaml"
  }
}
```

### Extension SDK

```typescript
// Using Docker Desktop Extension SDK
import { createDockerDesktopClient } from '@docker/extension-api-client';

const ddClient = createDockerDesktopClient();

// List containers
const containers = await ddClient.docker.listContainers();

// Run command
const result = await ddClient.docker.cli.exec('ps', ['-a']);

// Show notification
ddClient.desktopUI.toast.success('Operation completed!');
```

---

# Solution Architect Level

## 21. Production Architecture Patterns

### Blue-Green Deployment

```yaml
# docker-compose.blue.yml
version: '3.8'

services:
  app-blue:
    image: myapp:1.0
    networks:
      - app-network
    deploy:
      labels:
        - "traefik.http.routers.app.rule=Host(`app.example.com`)"

networks:
  app-network:
    external: true
```

```yaml
# docker-compose.green.yml
version: '3.8'

services:
  app-green:
    image: myapp:2.0
    networks:
      - app-network
    deploy:
      labels:
        - "traefik.http.routers.app-new.rule=Host(`app.example.com`)"
        - "traefik.http.routers.app-new.priority=2"

networks:
  app-network:
    external: true
```

### Canary Deployment with Traefik

```yaml
version: '3.8'

services:
  traefik:
    image: traefik:v2.10
    command:
      - "--providers.docker=true"
      - "--entrypoints.web.address=:80"
    ports:
      - "80:80"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

  app-stable:
    image: myapp:1.0
    labels:
      - "traefik.http.routers.app.rule=Host(`app.example.com`)"
      - "traefik.http.services.app-stable.loadbalancer.server.port=8080"
      - "traefik.http.services.app.weighted.services.app-stable.weight=90"

  app-canary:
    image: myapp:2.0
    labels:
      - "traefik.http.services.app-canary.loadbalancer.server.port=8080"
      - "traefik.http.services.app.weighted.services.app-canary.weight=10"
```

---

## 22. High Availability Setup

### HAProxy Load Balancer

```yaml
version: '3.8'

services:
  haproxy:
    image: haproxy:2.8
    ports:
      - "80:80"
      - "443:443"
      - "8404:8404"  # Stats
    volumes:
      - ./haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg:ro
    networks:
      - frontend
      - backend

  app1:
    image: myapp:latest
    networks:
      - backend

  app2:
    image: myapp:latest
    networks:
      - backend

  app3:
    image: myapp:latest
    networks:
      - backend

networks:
  frontend:
  backend:
```

```
# haproxy.cfg
global
    daemon
    maxconn 4096

defaults
    mode http
    timeout connect 5s
    timeout client 30s
    timeout server 30s
    option httplog
    option dontlognull

frontend http_front
    bind *:80
    default_backend http_back

backend http_back
    balance roundrobin
    option httpchk GET /health
    server app1 app1:8080 check
    server app2 app2:8080 check
    server app3 app3:8080 check

listen stats
    bind *:8404
    stats enable
    stats uri /stats
    stats refresh 10s
```

---

## 23. Service Mesh with Docker

### Consul Connect

```yaml
version: '3.8'

services:
  consul:
    image: consul:1.15
    command: agent -server -bootstrap-expect=1 -ui -client=0.0.0.0
    ports:
      - "8500:8500"
    volumes:
      - consul-data:/consul/data

  api:
    image: myapi:latest
    environment:
      - CONSUL_HTTP_ADDR=consul:8500
    labels:
      - "consul.register/enabled=true"
      - "consul.register/service.name=api"

  web:
    image: myweb:latest
    environment:
      - CONSUL_HTTP_ADDR=consul:8500
    depends_on:
      - api

volumes:
  consul-data:
```

### Linkerd Sidecar

```yaml
version: '3.8'

services:
  app:
    image: myapp:latest
    ports:
      - "8080:8080"

  linkerd-proxy:
    image: ghcr.io/linkerd/proxy:stable-2.14.0
    environment:
      - LINKERD2_PROXY_DESTINATION_SVC_ADDR=linkerd-destination.linkerd.svc.cluster.local:8086
    network_mode: "service:app"
```

---

## 24. Disaster Recovery

### Backup Strategy

```bash
#!/bin/bash
# backup.sh

# Backup volumes
docker run --rm \
  -v my-volume:/data:ro \
  -v $(pwd)/backup:/backup \
  alpine tar czf /backup/my-volume-$(date +%Y%m%d).tar.gz -C /data .

# Backup container configurations
for container in $(docker ps -q); do
  docker inspect $container > backup/container-$container.json
done

# Backup images
docker save myapp:latest | gzip > backup/myapp-latest.tar.gz

# Backup docker-compose
cp docker-compose.yml backup/
cp .env backup/

# Upload to S3
aws s3 sync backup/ s3://my-backup-bucket/docker/$(date +%Y%m%d)/
```

### Restore Script

```bash
#!/bin/bash
# restore.sh

# Restore volume
docker volume create my-volume
docker run --rm \
  -v my-volume:/data \
  -v $(pwd)/backup:/backup \
  alpine tar xzf /backup/my-volume-latest.tar.gz -C /data

# Load image
docker load < backup/myapp-latest.tar.gz

# Start services
docker-compose up -d
```

### Automated Backup with Docker

```yaml
version: '3.8'

services:
  backup:
    image: offen/docker-volume-backup:v2
    environment:
      - BACKUP_CRON_EXPRESSION=0 2 * * *
      - BACKUP_FILENAME=backup-%Y%m%d.tar.gz
      - AWS_S3_BUCKET_NAME=my-backup-bucket
      - AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
      - AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}
    volumes:
      - my-data:/backup/my-data:ro
      - /var/run/docker.sock:/var/run/docker.sock:ro

volumes:
  my-data:
```

---

## 25. Enterprise Docker Practices

### Image Governance

```yaml
# Policy: Only allow images from approved registries
# Using OPA (Open Policy Agent)

package docker.authz

default allow = false

allow {
    input.RequestMethod == "POST"
    input.RequestURI == "/containers/create"
    approved_image(input.RequestBody.Image)
}

approved_image(image) {
    startswith(image, "gcr.io/my-company/")
}

approved_image(image) {
    startswith(image, "my-registry.example.com/")
}
```

### Audit Logging

```json
// daemon.json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "50m",
    "max-file": "10",
    "labels": "env,app",
    "env": "os,customer"
  },
  "audit-log-path": "/var/log/docker-audit.log",
  "audit-log-maxsize": 100
}
```

### Network Segmentation

```yaml
version: '3.8'

services:
  frontend:
    image: nginx:latest
    networks:
      - public
      - internal

  api:
    image: myapi:latest
    networks:
      - internal
      - database

  db:
    image: postgres:15
    networks:
      - database

networks:
  public:
    driver: bridge
  internal:
    driver: bridge
    internal: true  # No external access
  database:
    driver: bridge
    internal: true
```

---

# Practice Questions with Answers

## Beginner Level Questions

### Q1: What is the difference between CMD and ENTRYPOINT?

**Answer:**

| Feature | CMD | ENTRYPOINT |
|---------|-----|------------|
| Override | Easily overridden at runtime | Requires `--entrypoint` flag |
| Purpose | Default arguments | Main executable |
| Combination | Provides arguments to ENTRYPOINT | Receives arguments from CMD |

```dockerfile
# CMD can be overridden
FROM ubuntu
CMD ["echo", "Hello"]
# docker run myimage  -> Hello
# docker run myimage echo "World"  -> World

# ENTRYPOINT is fixed
FROM ubuntu
ENTRYPOINT ["echo"]
CMD ["Hello"]
# docker run myimage  -> Hello
# docker run myimage World  -> World
# docker run --entrypoint /bin/bash myimage  -> bash shell
```

**Best Practice:**
- Use ENTRYPOINT for the main command
- Use CMD for default arguments

---

### Q2: What is the difference between COPY and ADD?

**Answer:**

| Feature | COPY | ADD |
|---------|------|-----|
| Copy files | ✓ | ✓ |
| Copy from URL | ✗ | ✓ |
| Auto-extract tar | ✗ | ✓ |
| Best practice | Preferred | Use when needed |

```dockerfile
# COPY - simple and predictable
COPY file.txt /app/
COPY . /app/

# ADD - auto-extracts tar files
ADD archive.tar.gz /app/

# ADD - downloads from URL (not recommended)
ADD https://example.com/file.txt /app/
```

**Best Practice:** Use COPY unless you need ADD's specific features.

---

### Q3: How do you view logs from a container?

**Answer:**
```bash
# View all logs
docker logs container-name

# Follow logs (like tail -f)
docker logs -f container-name

# View last N lines
docker logs --tail 100 container-name

# View logs since specific time
docker logs --since 2h container-name
docker logs --since 2024-01-01T00:00:00 container-name

# View logs with timestamps
docker logs -t container-name

# Combine options
docker logs -f --tail 100 -t container-name
```

---

### Q4: What is a Docker volume and why use it?

**Answer:**
A Docker volume is a persistent storage mechanism managed by Docker. Data in volumes survives container restarts and removal.

**Benefits:**
- Data persistence beyond container lifecycle
- Sharing data between containers
- Better performance than bind mounts
- Easier backup and migration
- Works across different environments

```bash
# Create volume
docker volume create mydata

# Use volume
docker run -v mydata:/app/data myimage

# Volume persists after container removal
docker rm mycontainer
docker volume ls  # mydata still exists
```

---

### Q5: How do you expose a port from a container?

**Answer:**
```dockerfile
# In Dockerfile (documentation only)
EXPOSE 80
EXPOSE 443
```

```bash
# Map specific port
docker run -p 8080:80 nginx
# Host port 8080 -> Container port 80

# Map to random host port
docker run -P nginx
# Docker assigns random host port

# Map to specific IP
docker run -p 127.0.0.1:8080:80 nginx

# Map UDP port
docker run -p 53:53/udp dns-server

# Multiple ports
docker run -p 80:80 -p 443:443 nginx
```

---

## Intermediate Level Questions

### Q6: Explain Docker layer caching and how to optimize it.

**Answer:**
Docker builds images layer by layer. Each instruction creates a new layer. If a layer hasn't changed, Docker uses the cached version.

**Optimization Tips:**

1. **Order by frequency of change:**
```dockerfile
# Dependencies change less frequently
COPY package*.json ./
RUN npm install

# Source changes frequently - keep at end
COPY . .
```

2. **Combine RUN commands:**
```dockerfile
# Bad - 3 layers
RUN apt-get update
RUN apt-get install -y package1
RUN apt-get clean

# Good - 1 layer
RUN apt-get update && \
    apt-get install -y package1 && \
    apt-get clean
```

3. **Use .dockerignore:**
```
node_modules
.git
*.log
```

4. **Use multi-stage builds:**
```dockerfile
FROM node:18 AS builder
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
```

---

### Q7: What is the difference between docker-compose up and docker-compose run?

**Answer:**

| Command | Purpose |
|---------|---------|
| `docker-compose up` | Starts all services defined in compose file |
| `docker-compose run` | Runs a one-time command in a new container |

```bash
# Start all services
docker-compose up

# Start in detached mode
docker-compose up -d

# Run one-time command
docker-compose run web npm test
docker-compose run db psql -U postgres

# Run without starting dependencies
docker-compose run --no-deps web npm test
```

---

### Q8: How do you limit container resources?

**Answer:**
```bash
# Memory limit
docker run -m 512m myimage
docker run --memory="1g" myimage

# CPU limit
docker run --cpus="1.5" myimage
docker run --cpu-shares=512 myimage

# Combined
docker run \
  --memory="512m" \
  --memory-swap="512m" \
  --cpus="0.5" \
  --pids-limit=100 \
  myimage
```

```yaml
# Docker Compose
services:
  app:
    image: myimage
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
        reservations:
          cpus: '0.1'
          memory: 256M
```

---

### Q9: How do you debug a container that keeps crashing?

**Answer:**

1. **Check logs:**
```bash
docker logs container-name
docker logs --tail 1000 container-name
```

2. **Inspect container:**
```bash
docker inspect container-name
docker inspect container-name --format '{{.State.ExitCode}}'
```

3. **Run with different command:**
```bash
# Override entrypoint to get shell
docker run -it --entrypoint /bin/sh myimage

# Keep container running for debugging
docker run -it myimage /bin/sh
```

4. **Check events:**
```bash
docker events --filter container=container-name
```

5. **Start in foreground:**
```bash
docker run myimage  # Without -d flag
```

6. **Use healthcheck:**
```dockerfile
HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD curl -f http://localhost/health || exit 1
```

---

### Q10: Explain Docker networking modes.

**Answer:**

| Mode | Description | Use Case |
|------|-------------|----------|
| `bridge` | Default isolated network | Standard container networking |
| `host` | Use host's network stack | Performance-critical apps |
| `none` | No networking | Security-sensitive containers |
| `overlay` | Multi-host networking | Docker Swarm |
| `macvlan` | Container gets own MAC | Legacy apps needing direct network access |

```bash
# Bridge (default)
docker run nginx

# Host
docker run --network host nginx

# None
docker run --network none nginx

# Custom network
docker network create mynet
docker run --network mynet nginx

# Connect to multiple networks
docker network connect mynet2 mycontainer
```

---

## Advanced Level Questions

### Q11: Explain multi-stage builds and their benefits.

**Answer:**
Multi-stage builds use multiple FROM statements. Each stage can use different base images and copy artifacts between stages.

**Benefits:**
- Smaller final images
- Separate build and runtime environments
- Keep build tools out of production
- Better security (less attack surface)
- Single Dockerfile for complex builds

```dockerfile
# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o myapp

# Test stage
FROM builder AS tester
RUN go test ./...

# Runtime stage
FROM alpine:3.18
COPY --from=builder /app/myapp /usr/local/bin/
CMD ["myapp"]
```

**Size comparison:**
```
golang:1.21 base: ~800MB
alpine:3.18 + binary: ~15MB
```

---

### Q12: How do you implement zero-downtime deployments with Docker?

**Answer:**

**1. Rolling Updates (Docker Swarm):**
```yaml
services:
  web:
    image: myapp:latest
    deploy:
      replicas: 3
      update_config:
        parallelism: 1
        delay: 10s
        failure_action: rollback
        order: start-first
      rollback_config:
        parallelism: 1
        delay: 10s
```

**2. Blue-Green with Load Balancer:**
```bash
# Start new version
docker-compose -f docker-compose.green.yml up -d

# Test new version
curl http://green.local:8080/health

# Switch traffic in load balancer
# Update nginx/haproxy config

# Stop old version
docker-compose -f docker-compose.blue.yml down
```

**3. Health Checks:**
```dockerfile
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1
```

**4. Graceful Shutdown:**
```dockerfile
# Handle SIGTERM
STOPSIGNAL SIGTERM
```

```go
// In application
signal.Notify(sigChan, syscall.SIGTERM)
// Finish current requests before shutting down
```

---

### Q13: How do you secure Docker in production?

**Answer:**

**1. Image Security:**
```bash
# Use minimal base images
FROM gcr.io/distroless/static

# Scan for vulnerabilities
trivy image myapp:latest

# Sign images
docker trust sign myapp:latest
```

**2. Runtime Security:**
```bash
# Run as non-root
docker run --user 1000:1000 myapp

# Read-only filesystem
docker run --read-only --tmpfs /tmp myapp

# Drop capabilities
docker run --cap-drop ALL --cap-add NET_BIND_SERVICE myapp

# No new privileges
docker run --security-opt no-new-privileges myapp

# Use seccomp
docker run --security-opt seccomp=profile.json myapp
```

**3. Resource Limits:**
```bash
docker run \
  --memory="512m" \
  --cpus="0.5" \
  --pids-limit=100 \
  --ulimit nofile=1024:1024 \
  myapp
```

**4. Network Security:**
```yaml
networks:
  internal:
    internal: true  # No external access
```

---

### Q14: Explain Docker's copy-on-write (CoW) strategy.

**Answer:**
Docker uses a layered filesystem with copy-on-write. Each layer is read-only except the top container layer.

**How it works:**
1. Base image is read-only layer
2. Each Dockerfile instruction adds a new layer
3. Container adds writable layer on top
4. When modifying a file, it's copied to the writable layer first

**Benefits:**
- Fast container startup
- Efficient storage (shared layers)
- Quick image creation

**Storage Drivers:**
- **overlay2** (recommended): Uses OverlayFS
- **devicemapper**: Uses device mapper thin provisioning
- **btrfs/zfs**: Uses filesystem features

```bash
# View layers
docker history myimage

# View storage driver
docker info | grep "Storage Driver"
```

---

### Q15: How do you implement Docker secrets management?

**Answer:**

**1. Docker Swarm Secrets:**
```bash
# Create secret
echo "my-password" | docker secret create db_password -

# Use in service
docker service create \
  --name web \
  --secret db_password \
  myimage
```

```yaml
# docker-compose.yml (Swarm mode)
services:
  db:
    image: postgres
    secrets:
      - db_password
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password

secrets:
  db_password:
    external: true
```

**2. External Secret Management:**
```yaml
# HashiCorp Vault
services:
  app:
    image: myapp
    environment:
      VAULT_ADDR: http://vault:8200
      VAULT_TOKEN: ${VAULT_TOKEN}
```

**3. Build-time Secrets (Buildx):**
```dockerfile
RUN --mount=type=secret,id=github_token \
    cat /run/secrets/github_token | npm login
```

```bash
docker buildx build --secret id=github_token,src=.github_token .
```

---

## Expert Level Questions

### Q16: How does Docker handle container networking internally?

**Answer:**

**Bridge Network (default):**
1. Docker creates `docker0` bridge interface on host
2. Each container gets a `veth` pair (virtual ethernet)
3. One end in container's network namespace
4. Other end attached to `docker0` bridge
5. iptables rules handle NAT for external access

**Network flow:**
```
Container -> veth -> docker0 bridge -> iptables NAT -> host eth0 -> external
```

**Inspecting:**
```bash
# View bridge
ip addr show docker0
brctl show docker0

# View container veth
docker exec container ip addr

# View iptables rules
iptables -t nat -L -n -v
```

**Overlay Network (Swarm):**
1. Uses VXLAN for encapsulation
2. Control plane via gossip protocol
3. Data plane via VXLAN tunnels
4. Each node has ingress network for load balancing

---

### Q17: Explain Docker's OCI compliance.

**Answer:**

**OCI (Open Container Initiative) Standards:**

1. **Runtime Specification:**
   - Defines how to run a container
   - runc is reference implementation
   - Docker uses containerd + runc

2. **Image Specification:**
   - Defines image format
   - Layers, manifests, configs
   - Portable across registries

3. **Distribution Specification:**
   - Defines how to push/pull images
   - Registry API standard

**Container Runtime Architecture:**
```
Docker CLI
    ↓
Docker Daemon (dockerd)
    ↓
containerd
    ↓
runc (OCI runtime)
    ↓
Linux Kernel (namespaces, cgroups)
```

**Using alternative runtimes:**
```bash
# Use gVisor (secure)
docker run --runtime=runsc myimage

# Use kata containers (VM-based)
docker run --runtime=kata myimage
```

---

### Q18: How do you implement a Docker-based CI/CD pipeline?

**Answer:**

**Complete Pipeline Example:**

```yaml
# .github/workflows/ci-cd.yml
name: CI/CD Pipeline

on:
  push:
    branches: [main]
    tags: ['v*']
  pull_request:
    branches: [main]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Build test image
        run: docker build --target test -t test-image .
      
      - name: Run tests
        run: docker run test-image npm test
      
      - name: Run security scan
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: test-image

  build:
    needs: test
    runs-on: ubuntu-latest
    outputs:
      digest: ${{ steps.build.outputs.digest }}
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      
      - name: Login to Registry
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Build and push
        id: build
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:${{ github.sha }}
            ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:latest
          cache-from: type=gha
          cache-to: type=gha,mode=max

  deploy:
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Deploy to production
        run: |
          # Update Kubernetes deployment
          kubectl set image deployment/myapp \
            app=${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}@${{ needs.build.outputs.digest }}
```

---

### Q19: How do you implement container health monitoring?

**Answer:**

**1. Docker Health Checks:**
```dockerfile
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1
```

**2. Prometheus Monitoring:**
```yaml
version: '3.8'

services:
  app:
    image: myapp:latest
    labels:
      - "prometheus.scrape=true"
      - "prometheus.port=8080"

  cadvisor:
    image: gcr.io/cadvisor/cadvisor:latest
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:ro
      - /sys:/sys:ro
      - /var/lib/docker/:/var/lib/docker:ro

  prometheus:
    image: prom/prometheus:latest
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
```

**3. Application Metrics:**
```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
)

func init() {
    prometheus.MustRegister(requestCount)
}

func main() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}
```

---

### Q20: How do you handle stateful applications in Docker?

**Answer:**

**1. Use Named Volumes:**
```yaml
services:
  postgres:
    image: postgres:15
    volumes:
      - postgres-data:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
    secrets:
      - db_password

volumes:
  postgres-data:
    driver: local
    driver_opts:
      type: none
      device: /mnt/data/postgres
      o: bind
```

**2. Database Clustering:**
```yaml
services:
  postgres-primary:
    image: postgres:15
    environment:
      POSTGRES_REPLICATION_MODE: master
    volumes:
      - primary-data:/var/lib/postgresql/data

  postgres-replica:
    image: postgres:15
    environment:
      POSTGRES_REPLICATION_MODE: slave
      POSTGRES_MASTER_HOST: postgres-primary
    volumes:
      - replica-data:/var/lib/postgresql/data
    depends_on:
      - postgres-primary
```

**3. Backup Strategy:**
```bash
#!/bin/bash
# Backup script
docker exec postgres pg_dump -U user dbname > backup.sql

# With volume
docker run --rm \
  -v postgres-data:/data:ro \
  -v $(pwd)/backup:/backup \
  alpine tar czf /backup/postgres-$(date +%Y%m%d).tar.gz /data
```

**4. Storage Drivers for Production:**
```yaml
volumes:
  data:
    driver: rexray/ebs  # AWS EBS
    driver_opts:
      size: "100"
      volumetype: "gp3"
```

---

## Solution Architect Level Questions

### Q21: Design a microservices architecture with Docker.

**Answer:**

```yaml
# docker-compose.yml
version: '3.8'

services:
  # API Gateway
  traefik:
    image: traefik:v2.10
    command:
      - "--api.insecure=true"
      - "--providers.docker=true"
      - "--entrypoints.web.address=:80"
      - "--tracing.jaeger=true"
    ports:
      - "80:80"
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

  # Service Discovery
  consul:
    image: consul:1.15
    ports:
      - "8500:8500"

  # User Service
  user-service:
    image: user-service:latest
    labels:
      - "traefik.http.routers.user.rule=PathPrefix(`/api/users`)"
    environment:
      - CONSUL_HTTP_ADDR=consul:8500
      - JAEGER_AGENT_HOST=jaeger
    depends_on:
      - consul
      - user-db

  user-db:
    image: postgres:15
    volumes:
      - user-db-data:/var/lib/postgresql/data

  # Order Service
  order-service:
    image: order-service:latest
    labels:
      - "traefik.http.routers.order.rule=PathPrefix(`/api/orders`)"
    environment:
      - CONSUL_HTTP_ADDR=consul:8500
      - KAFKA_BROKERS=kafka:9092
    depends_on:
      - consul
      - kafka

  # Message Queue
  kafka:
    image: confluentinc/cp-kafka:7.5.0
    environment:
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
    depends_on:
      - zookeeper

  zookeeper:
    image: confluentinc/cp-zookeeper:7.5.0

  # Observability
  jaeger:
    image: jaegertracing/all-in-one:latest
    ports:
      - "16686:16686"

  prometheus:
    image: prom/prometheus:latest
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"

volumes:
  user-db-data:
```

---

### Q22: How do you implement container orchestration at scale?

**Answer:**

**Docker Swarm for Medium Scale:**
```bash
# Initialize swarm
docker swarm init

# Add manager nodes (for HA)
docker swarm join-token manager

# Add worker nodes
docker swarm join-token worker

# Deploy stack
docker stack deploy -c docker-compose.yml myapp

# Scale services
docker service scale myapp_web=10
```

**Stack Configuration:**
```yaml
version: '3.8'

services:
  web:
    image: myapp:latest
    deploy:
      mode: replicated
      replicas: 6
      placement:
        constraints:
          - node.role == worker
        preferences:
          - spread: node.labels.zone
      update_config:
        parallelism: 2
        delay: 10s
        failure_action: rollback
        order: start-first
      rollback_config:
        parallelism: 1
        delay: 10s
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
        reservations:
          cpus: '0.1'
          memory: 256M
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 3
        window: 120s
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

---

### Q23: Design a disaster recovery strategy for Docker.

**Answer:**

**1. Multi-Region Setup:**
```yaml
# Primary Region
services:
  app:
    image: myapp:latest
    deploy:
      placement:
        constraints:
          - node.labels.region == us-east

  db:
    image: postgres:15
    volumes:
      - db-primary:/var/lib/postgresql/data
    environment:
      POSTGRES_REPLICATION_MODE: master

# Secondary Region (DR)
  db-replica:
    image: postgres:15
    deploy:
      placement:
        constraints:
          - node.labels.region == us-west
    environment:
      POSTGRES_REPLICATION_MODE: slave
      POSTGRES_MASTER_HOST: db
```

**2. Backup Automation:**
```yaml
services:
  backup:
    image: backup-service:latest
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      BACKUP_SCHEDULE: "0 */6 * * *"
      S3_BUCKET: my-backup-bucket
      ENCRYPTION_KEY: ${BACKUP_ENCRYPTION_KEY}
    deploy:
      placement:
        constraints:
          - node.role == manager
```

**3. Recovery Runbook:**
```bash
#!/bin/bash
# dr-failover.sh

# 1. Verify primary is down
if ! curl -f http://primary.example.com/health; then
  echo "Primary is down, initiating failover"
  
  # 2. Promote replica to primary
  docker exec db-replica pg_ctl promote
  
  # 3. Update DNS
  aws route53 change-resource-record-sets \
    --hosted-zone-id Z123 \
    --change-batch file://dns-failover.json
  
  # 4. Scale up DR region
  docker service scale myapp_web=10
  
  # 5. Notify team
  curl -X POST $SLACK_WEBHOOK -d '{"text":"DR failover complete"}'
fi
```

---

### Q24: How do you implement container security at enterprise level?

**Answer:**

**1. Image Pipeline Security:**
```yaml
# GitHub Actions security pipeline
name: Secure Build

on: push

jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      # Static code analysis
      - name: SonarQube Scan
        uses: sonarqube/sonarcloud-github-action@master
      
      # Dockerfile linting
      - name: Hadolint
        uses: hadolint/hadolint-action@v3
      
      # Build image
      - name: Build
        run: docker build -t myapp:${{ github.sha }} .
      
      # Vulnerability scan
      - name: Trivy Scan
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: myapp:${{ github.sha }}
          severity: 'CRITICAL,HIGH'
          exit-code: '1'
      
      # Sign image
      - name: Sign Image
        run: |
          cosign sign --key cosign.key myapp:${{ github.sha }}
```

**2. Runtime Security:**
```yaml
# seccomp profile
{
  "defaultAction": "SCMP_ACT_ERRNO",
  "syscalls": [
    {
      "names": ["accept", "bind", "clone", "close", "connect", ...],
      "action": "SCMP_ACT_ALLOW"
    }
  ]
}
```

```yaml
# docker-compose with security
services:
  app:
    image: myapp:latest
    security_opt:
      - no-new-privileges:true
      - seccomp:seccomp-profile.json
    cap_drop:
      - ALL
    cap_add:
      - NET_BIND_SERVICE
    read_only: true
    tmpfs:
      - /tmp
    user: "1000:1000"
```

**3. Network Policies:**
```yaml
services:
  frontend:
    networks:
      - public
      - internal
  
  api:
    networks:
      - internal
      - database
  
  db:
    networks:
      - database

networks:
  public:
  internal:
    internal: true
  database:
    internal: true
```

---

### Q25: Design a complete production Docker platform.

**Answer:**

**Architecture Overview:**
```
┌─────────────────────────────────────────────────────────────┐
│                     Load Balancer (HAProxy/Traefik)         │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                      Docker Swarm Cluster                    │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │   Manager   │  │   Manager   │  │   Manager   │         │
│  │    Node 1   │  │    Node 2   │  │    Node 3   │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │   Worker    │  │   Worker    │  │   Worker    │         │
│  │    Node 1   │  │    Node 2   │  │    Node N   │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                     Shared Services                          │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐       │
│  │ Registry │ │  Consul  │ │  Vault   │ │  Kafka   │       │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘       │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐                     │
│  │Prometheus│ │  Grafana │ │  Jaeger  │                     │
│  └──────────┘ └──────────┘ └──────────┘                     │
└─────────────────────────────────────────────────────────────┘
```

**Complete Platform Stack:**
```yaml
version: '3.8'

services:
  # Ingress
  traefik:
    image: traefik:v2.10
    command:
      - "--api.dashboard=true"
      - "--providers.docker.swarmMode=true"
      - "--entrypoints.web.address=:80"
      - "--entrypoints.websecure.address=:443"
      - "--certificatesresolvers.le.acme.email=admin@example.com"
      - "--certificatesresolvers.le.acme.storage=/letsencrypt/acme.json"
      - "--certificatesresolvers.le.acme.httpchallenge.entrypoint=web"
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - letsencrypt:/letsencrypt
    deploy:
      mode: replicated
      replicas: 2
      placement:
        constraints:
          - node.role == manager

  # Service Discovery & Config
  consul:
    image: consul:1.15
    command: agent -server -bootstrap-expect=3 -ui
    volumes:
      - consul-data:/consul/data
    deploy:
      mode: global
      placement:
        constraints:
          - node.role == manager

  # Secrets Management
  vault:
    image: vault:1.15
    environment:
      VAULT_ADDR: http://127.0.0.1:8200
    volumes:
      - vault-data:/vault/data
    deploy:
      placement:
        constraints:
          - node.role == manager

  # Private Registry
  registry:
    image: registry:2
    volumes:
      - registry-data:/var/lib/registry
    deploy:
      placement:
        constraints:
          - node.labels.storage == ssd

  # Monitoring Stack
  prometheus:
    image: prom/prometheus:latest
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
      - prometheus-data:/prometheus
    deploy:
      placement:
        constraints:
          - node.role == manager

  grafana:
    image: grafana/grafana:latest
    environment:
      - GF_SERVER_ROOT_URL=https://grafana.example.com
    volumes:
      - grafana-data:/var/lib/grafana

  # Logging
  loki:
    image: grafana/loki:latest
    volumes:
      - loki-data:/loki

  promtail:
    image: grafana/promtail:latest
    volumes:
      - /var/log:/var/log:ro
      - /var/lib/docker/containers:/var/lib/docker/containers:ro
    deploy:
      mode: global

  # Tracing
  jaeger:
    image: jaegertracing/all-in-one:latest
    environment:
      - COLLECTOR_OTLP_ENABLED=true

volumes:
  letsencrypt:
  consul-data:
  vault-data:
  registry-data:
  prometheus-data:
  grafana-data:
  loki-data:

networks:
  default:
    driver: overlay
    attachable: true
```

---

## Summary

This comprehensive Docker guide covers:

- **Beginner:** Basic commands, Dockerfile, volumes, networking, Compose
- **Intermediate:** Multi-stage builds, security, logging, registry
- **Advanced:** Swarm, Buildx, orchestration patterns, API/SDK
- **Expert:** Internals, CI/CD, monitoring, security hardening
- **Solution Architect:** Production patterns, HA, DR, enterprise security

**Key Takeaways:**
1. Always use multi-stage builds for smaller, secure images
2. Never run containers as root in production
3. Implement health checks for all services
4. Use proper logging and monitoring from day one
5. Automate security scanning in CI/CD
6. Plan for disaster recovery with proper backups
7. Use orchestration (Swarm/Kubernetes) for production workloads

---

*Document Version: 1.0*
*Last Updated: January 2026*
