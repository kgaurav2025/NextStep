# Solution Architect Mastery Guide
## From Technical Expert to Solution Architect

---

## Table of Contents
1. [Introduction to Solution Architecture](#introduction-to-solution-architecture)
2. [System Design Fundamentals](#system-design-fundamentals)
3. [Architectural Patterns](#architectural-patterns)
4. [Distributed Systems](#distributed-systems)
5. [Cloud Architecture](#cloud-architecture)
6. [Security Architecture](#security-architecture)
7. [Data Architecture](#data-architecture)
8. [Integration Architecture](#integration-architecture)
9. [DevOps & Platform Engineering](#devops--platform-engineering)
10. [Non-Functional Requirements](#non-functional-requirements)
11. [Architectural Decision Making](#architectural-decision-making)
12. [Case Studies & System Designs](#case-studies--system-designs)
13. [Practice Questions with Answers](#practice-questions-with-answers)

---

# Introduction to Solution Architecture

## What is a Solution Architect?

A Solution Architect is a technical leader who designs end-to-end solutions that meet business requirements while ensuring technical excellence, scalability, security, and maintainability.

### Key Responsibilities

| Area | Responsibilities |
|------|-----------------|
| **Technical Leadership** | Define technology stack, guide engineering teams, set standards |
| **System Design** | Design scalable, reliable, and secure systems |
| **Stakeholder Management** | Bridge business and technical teams |
| **Risk Management** | Identify and mitigate technical risks |
| **Documentation** | Create architecture documents, ADRs, runbooks |
| **Governance** | Ensure compliance with standards and best practices |

### Skills Required

```
┌─────────────────────────────────────────────────────────────────┐
│                    Solution Architect Skills                     │
├─────────────────┬─────────────────┬─────────────────────────────┤
│   Technical     │   Business      │   Soft Skills               │
├─────────────────┼─────────────────┼─────────────────────────────┤
│ Cloud Platforms │ Domain Knowledge│ Communication               │
│ Distributed Sys │ Cost Analysis   │ Leadership                  │
│ Security        │ Risk Assessment │ Problem Solving             │
│ Data Systems    │ Stakeholder Mgmt│ Decision Making             │
│ DevOps/SRE      │ Compliance      │ Documentation               │
│ Multiple Languages│ ROI Analysis  │ Mentoring                   │
└─────────────────┴─────────────────┴─────────────────────────────┘
```

---

# System Design Fundamentals

## 1. Scalability

### Horizontal vs Vertical Scaling

| Aspect | Horizontal Scaling | Vertical Scaling |
|--------|-------------------|------------------|
| Method | Add more machines | Add more resources to existing machine |
| Complexity | Higher | Lower |
| Cost | Linear | Exponential at high end |
| Limits | Nearly unlimited | Hardware limits |
| Downtime | None | May require restart |
| Example | Add web servers | Upgrade to larger instance |

### Scalability Patterns

```
┌─────────────────────────────────────────────────────────────────┐
│                      Load Balancer                               │
│                    (Round Robin/Weighted)                        │
└────────────────────────────┬────────────────────────────────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│  App Server  │    │  App Server  │    │  App Server  │
│     #1       │    │     #2       │    │     #N       │
└──────┬───────┘    └──────┬───────┘    └──────┬───────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
                           ▼
              ┌─────────────────────────┐
              │      Cache Layer        │
              │   (Redis Cluster)       │
              └────────────┬────────────┘
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
        ▼                  ▼                  ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│   Primary    │  │   Replica    │  │   Replica    │
│   Database   │──│   Database   │  │   Database   │
└──────────────┘  └──────────────┘  └──────────────┘
```

---

## 2. Reliability & Availability

### Availability Metrics

| Availability | Downtime/Year | Downtime/Month | Downtime/Week |
|-------------|---------------|----------------|---------------|
| 99% (two 9s) | 3.65 days | 7.31 hours | 1.68 hours |
| 99.9% (three 9s) | 8.77 hours | 43.83 minutes | 10.08 minutes |
| 99.99% (four 9s) | 52.60 minutes | 4.38 minutes | 1.01 minutes |
| 99.999% (five 9s) | 5.26 minutes | 26.30 seconds | 6.05 seconds |

### Reliability Patterns

```go
// Circuit Breaker Pattern
type CircuitBreaker struct {
    failureThreshold int
    successThreshold int
    timeout          time.Duration
    failures         int
    successes        int
    state            State
    lastFailure      time.Time
    mu               sync.RWMutex
}

type State int

const (
    Closed State = iota  // Normal operation
    Open                 // Failing, reject requests
    HalfOpen            // Testing if recovered
)

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    switch cb.state {
    case Open:
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = HalfOpen
            cb.successes = 0
        } else {
            return ErrCircuitOpen
        }
    }
    
    err := fn()
    
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.failureThreshold {
            cb.state = Open
        }
        return err
    }
    
    if cb.state == HalfOpen {
        cb.successes++
        if cb.successes >= cb.successThreshold {
            cb.state = Closed
            cb.failures = 0
        }
    }
    
    return nil
}
```

### Retry Pattern with Exponential Backoff

```go
func RetryWithBackoff(ctx context.Context, maxRetries int, fn func() error) error {
    var err error
    baseDelay := 100 * time.Millisecond
    maxDelay := 30 * time.Second
    
    for i := 0; i < maxRetries; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        
        if i == maxRetries-1 {
            break
        }
        
        // Exponential backoff with jitter
        delay := baseDelay * time.Duration(1<<uint(i))
        if delay > maxDelay {
            delay = maxDelay
        }
        jitter := time.Duration(rand.Int63n(int64(delay / 2)))
        delay = delay + jitter
        
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(delay):
        }
    }
    
    return fmt.Errorf("max retries exceeded: %w", err)
}
```

---

## 3. Performance

### Latency Optimization

| Technique | Impact | Implementation |
|-----------|--------|----------------|
| Caching | High | Redis, Memcached, CDN |
| Connection pooling | Medium | Database pools |
| Async processing | High | Message queues |
| Compression | Medium | gzip, brotli |
| Edge computing | High | CDN, edge functions |
| Query optimization | High | Indexes, query plans |

### Caching Strategies

```go
// Cache-Aside (Lazy Loading)
func GetUser(ctx context.Context, userID string) (*User, error) {
    // Try cache first
    cacheKey := fmt.Sprintf("user:%s", userID)
    cached, err := cache.Get(ctx, cacheKey)
    if err == nil {
        var user User
        json.Unmarshal(cached, &user)
        return &user, nil
    }
    
    // Cache miss - fetch from database
    user, err := db.GetUser(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    data, _ := json.Marshal(user)
    cache.Set(ctx, cacheKey, data, 1*time.Hour)
    
    return user, nil
}

// Write-Through Cache
func UpdateUser(ctx context.Context, user *User) error {
    // Update database first
    if err := db.UpdateUser(ctx, user); err != nil {
        return err
    }
    
    // Update cache
    cacheKey := fmt.Sprintf("user:%s", user.ID)
    data, _ := json.Marshal(user)
    return cache.Set(ctx, cacheKey, data, 1*time.Hour)
}

// Write-Behind (Write-Back) Cache
func UpdateUserAsync(ctx context.Context, user *User) error {
    // Update cache immediately
    cacheKey := fmt.Sprintf("user:%s", user.ID)
    data, _ := json.Marshal(user)
    if err := cache.Set(ctx, cacheKey, data, 1*time.Hour); err != nil {
        return err
    }
    
    // Queue database update
    return queue.Publish(ctx, "user-updates", user)
}
```

---

## 4. Consistency Models

### CAP Theorem

```
                        Consistency
                            /\
                           /  \
                          /    \
                         /  CP  \
                        /________\
                       /          \
                      /    CA      \
                     /_____________ \
                    /                \
                   /       AP         \
                  /____________________\
            Availability          Partition Tolerance
            
CP: Consistent but may be unavailable (MongoDB with strong reads)
AP: Available but eventually consistent (Cassandra, DynamoDB)
CA: Not realistic in distributed systems (single node)
```

### Consistency Levels

| Level | Description | Use Case |
|-------|-------------|----------|
| Strong | All reads return latest write | Financial transactions |
| Eventual | Data converges eventually | Social media feeds |
| Causal | Respects causality | Chat applications |
| Read-your-writes | See your own writes | User preferences |
| Session | Consistency within session | Shopping cart |

---

# Architectural Patterns

## 1. Microservices Architecture

### Microservices Design Principles

```
┌─────────────────────────────────────────────────────────────────┐
│                         API Gateway                              │
│              (Rate Limiting, Auth, Routing)                      │
└────────────────────────────┬────────────────────────────────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│    User      │    │    Order     │    │   Product    │
│   Service    │    │   Service    │    │   Service    │
└──────┬───────┘    └──────┬───────┘    └──────┬───────┘
       │                   │                   │
       ▼                   ▼                   ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│   User DB    │    │  Order DB    │    │  Product DB  │
│  (PostgreSQL)│    │   (MongoDB)  │    │ (PostgreSQL) │
└──────────────┘    └──────────────┘    └──────────────┘
```

### Service Communication Patterns

```go
// Synchronous - REST/gRPC
type OrderService struct {
    userClient  UserServiceClient
    productClient ProductServiceClient
}

func (s *OrderService) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    // Sync call to validate user
    user, err := s.userClient.GetUser(ctx, &GetUserRequest{ID: req.UserID})
    if err != nil {
        return nil, fmt.Errorf("user validation failed: %w", err)
    }
    
    // Sync call to check product
    product, err := s.productClient.GetProduct(ctx, &GetProductRequest{ID: req.ProductID})
    if err != nil {
        return nil, fmt.Errorf("product validation failed: %w", err)
    }
    
    // Create order
    order := &Order{
        ID:        uuid.New().String(),
        UserID:    user.ID,
        ProductID: product.ID,
        Amount:    product.Price * req.Quantity,
        Status:    "pending",
    }
    
    return order, s.repo.Save(ctx, order)
}

// Asynchronous - Event-Driven
func (s *OrderService) CreateOrderAsync(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    order := &Order{
        ID:        uuid.New().String(),
        UserID:    req.UserID,
        ProductID: req.ProductID,
        Quantity:  req.Quantity,
        Status:    "pending",
    }
    
    if err := s.repo.Save(ctx, order); err != nil {
        return nil, err
    }
    
    // Publish event for other services
    event := &OrderCreatedEvent{
        OrderID:   order.ID,
        UserID:    order.UserID,
        ProductID: order.ProductID,
        Timestamp: time.Now(),
    }
    
    if err := s.eventBus.Publish(ctx, "orders.created", event); err != nil {
        // Log but don't fail - compensating transaction later
        log.Printf("Failed to publish event: %v", err)
    }
    
    return order, nil
}
```

---

## 2. Event-Driven Architecture

### Event Sourcing

```go
// Event types
type Event interface {
    AggregateID() string
    EventType() string
    Timestamp() time.Time
}

type AccountCreatedEvent struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type MoneyDepositedEvent struct {
    AccountID string    `json:"account_id"`
    Amount    float64   `json:"amount"`
    Timestamp time.Time `json:"timestamp"`
}

type MoneyWithdrawnEvent struct {
    AccountID string    `json:"account_id"`
    Amount    float64   `json:"amount"`
    Timestamp time.Time `json:"timestamp"`
}

// Event Store
type EventStore interface {
    Save(ctx context.Context, events ...Event) error
    Load(ctx context.Context, aggregateID string) ([]Event, error)
    LoadFrom(ctx context.Context, aggregateID string, version int) ([]Event, error)
}

// Aggregate (rebuilds state from events)
type Account struct {
    ID      string
    Name    string
    Email   string
    Balance float64
    Version int
}

func (a *Account) Apply(event Event) {
    switch e := event.(type) {
    case *AccountCreatedEvent:
        a.ID = e.ID
        a.Name = e.Name
        a.Email = e.Email
        a.Balance = 0
    case *MoneyDepositedEvent:
        a.Balance += e.Amount
    case *MoneyWithdrawnEvent:
        a.Balance -= e.Amount
    }
    a.Version++
}

func LoadAccount(ctx context.Context, store EventStore, id string) (*Account, error) {
    events, err := store.Load(ctx, id)
    if err != nil {
        return nil, err
    }
    
    account := &Account{}
    for _, event := range events {
        account.Apply(event)
    }
    
    return account, nil
}
```

### CQRS (Command Query Responsibility Segregation)

```go
// Command Side
type CommandHandler interface {
    Handle(ctx context.Context, cmd Command) error
}

type CreateOrderCommand struct {
    OrderID   string
    UserID    string
    Items     []OrderItem
}

type CreateOrderHandler struct {
    repo     OrderRepository
    eventBus EventBus
}

func (h *CreateOrderHandler) Handle(ctx context.Context, cmd Command) error {
    createCmd := cmd.(*CreateOrderCommand)
    
    order := &Order{
        ID:     createCmd.OrderID,
        UserID: createCmd.UserID,
        Items:  createCmd.Items,
        Status: "created",
    }
    
    if err := h.repo.Save(ctx, order); err != nil {
        return err
    }
    
    // Publish event for read model update
    return h.eventBus.Publish(ctx, "orders", &OrderCreatedEvent{
        OrderID: order.ID,
        UserID:  order.UserID,
        Items:   order.Items,
    })
}

// Query Side (Read Model)
type OrderReadModel struct {
    ID          string    `json:"id"`
    UserID      string    `json:"user_id"`
    UserName    string    `json:"user_name"`  // Denormalized
    ItemCount   int       `json:"item_count"`
    TotalAmount float64   `json:"total_amount"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}

type OrderQueryService struct {
    readDB *sql.DB
}

func (s *OrderQueryService) GetOrdersByUser(ctx context.Context, userID string) ([]OrderReadModel, error) {
    query := `
        SELECT id, user_id, user_name, item_count, total_amount, status, created_at
        FROM order_read_model
        WHERE user_id = $1
        ORDER BY created_at DESC
    `
    // Execute optimized read query
    return s.executeQuery(ctx, query, userID)
}

// Event Handler for Read Model Projection
type OrderProjector struct {
    readDB *sql.DB
}

func (p *OrderProjector) Handle(ctx context.Context, event Event) error {
    switch e := event.(type) {
    case *OrderCreatedEvent:
        return p.projectOrderCreated(ctx, e)
    case *OrderShippedEvent:
        return p.projectOrderShipped(ctx, e)
    }
    return nil
}
```

---

## 3. Saga Pattern

### Choreography Saga

```go
// Each service listens for events and reacts
type OrderService struct {
    eventBus EventBus
}

func (s *OrderService) HandlePaymentCompleted(ctx context.Context, event *PaymentCompletedEvent) error {
    order, err := s.repo.GetByPaymentID(ctx, event.PaymentID)
    if err != nil {
        return err
    }
    
    order.Status = "paid"
    if err := s.repo.Save(ctx, order); err != nil {
        return err
    }
    
    // Trigger next step
    return s.eventBus.Publish(ctx, "orders.paid", &OrderPaidEvent{
        OrderID: order.ID,
    })
}

func (s *OrderService) HandlePaymentFailed(ctx context.Context, event *PaymentFailedEvent) error {
    // Compensating action
    order, err := s.repo.GetByPaymentID(ctx, event.PaymentID)
    if err != nil {
        return err
    }
    
    order.Status = "cancelled"
    return s.repo.Save(ctx, order)
}
```

### Orchestration Saga

```go
// Central orchestrator coordinates the saga
type OrderSaga struct {
    orderService   OrderServiceClient
    paymentService PaymentServiceClient
    inventoryService InventoryServiceClient
    shippingService  ShippingServiceClient
}

type SagaStep struct {
    Name       string
    Execute    func(ctx context.Context, state *SagaState) error
    Compensate func(ctx context.Context, state *SagaState) error
}

type SagaState struct {
    OrderID     string
    PaymentID   string
    ReservationID string
    ShipmentID  string
    Error       error
}

func (s *OrderSaga) Execute(ctx context.Context, orderReq *CreateOrderRequest) error {
    state := &SagaState{OrderID: uuid.New().String()}
    
    steps := []SagaStep{
        {
            Name: "CreateOrder",
            Execute: func(ctx context.Context, state *SagaState) error {
                _, err := s.orderService.Create(ctx, &Order{ID: state.OrderID})
                return err
            },
            Compensate: func(ctx context.Context, state *SagaState) error {
                return s.orderService.Cancel(ctx, state.OrderID)
            },
        },
        {
            Name: "ReserveInventory",
            Execute: func(ctx context.Context, state *SagaState) error {
                res, err := s.inventoryService.Reserve(ctx, orderReq.Items)
                if err != nil {
                    return err
                }
                state.ReservationID = res.ID
                return nil
            },
            Compensate: func(ctx context.Context, state *SagaState) error {
                return s.inventoryService.Release(ctx, state.ReservationID)
            },
        },
        {
            Name: "ProcessPayment",
            Execute: func(ctx context.Context, state *SagaState) error {
                payment, err := s.paymentService.Process(ctx, orderReq.PaymentInfo)
                if err != nil {
                    return err
                }
                state.PaymentID = payment.ID
                return nil
            },
            Compensate: func(ctx context.Context, state *SagaState) error {
                return s.paymentService.Refund(ctx, state.PaymentID)
            },
        },
    }
    
    // Execute saga with compensation on failure
    executedSteps := []SagaStep{}
    for _, step := range steps {
        if err := step.Execute(ctx, state); err != nil {
            state.Error = err
            // Compensate in reverse order
            for i := len(executedSteps) - 1; i >= 0; i-- {
                if compErr := executedSteps[i].Compensate(ctx, state); compErr != nil {
                    log.Printf("Compensation failed for %s: %v", executedSteps[i].Name, compErr)
                }
            }
            return err
        }
        executedSteps = append(executedSteps, step)
    }
    
    return nil
}
```

---

# Distributed Systems

## 1. Consensus Algorithms

### Raft Basics

```
┌─────────────────────────────────────────────────────────────────┐
│                         Raft Cluster                             │
│                                                                  │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │   Leader     │◄───│   Follower   │    │   Follower   │      │
│  │   (Node 1)   │───►│   (Node 2)   │    │   (Node 3)   │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
│        │                    ▲                   ▲               │
│        │                    │                   │               │
│        └────────────────────┴───────────────────┘               │
│                   Heartbeats & Log Replication                   │
└─────────────────────────────────────────────────────────────────┘

States:
- Leader: Handles all client requests, replicates log
- Follower: Passive, responds to leader
- Candidate: During election
```

### Distributed Locking

```go
// Redis-based distributed lock
type DistributedLock struct {
    client *redis.Client
    key    string
    value  string
    ttl    time.Duration
}

func (l *DistributedLock) Acquire(ctx context.Context) (bool, error) {
    l.value = uuid.New().String()
    result, err := l.client.SetNX(ctx, l.key, l.value, l.ttl).Result()
    return result, err
}

func (l *DistributedLock) Release(ctx context.Context) error {
    script := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `
    _, err := l.client.Eval(ctx, script, []string{l.key}, l.value).Result()
    return err
}

func (l *DistributedLock) Extend(ctx context.Context, additionalTTL time.Duration) error {
    script := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("pexpire", KEYS[1], ARGV[2])
        else
            return 0
        end
    `
    _, err := l.client.Eval(ctx, script, []string{l.key}, l.value, int64(additionalTTL/time.Millisecond)).Result()
    return err
}

// Usage with lock
func ProcessOrder(ctx context.Context, orderID string) error {
    lock := &DistributedLock{
        client: redisClient,
        key:    fmt.Sprintf("order-lock:%s", orderID),
        ttl:    30 * time.Second,
    }
    
    acquired, err := lock.Acquire(ctx)
    if err != nil {
        return fmt.Errorf("failed to acquire lock: %w", err)
    }
    if !acquired {
        return ErrOrderLocked
    }
    defer lock.Release(ctx)
    
    // Process order with lock held
    return processOrderInternal(ctx, orderID)
}
```

---

## 2. Consistent Hashing

```go
// Consistent hash ring for data distribution
type ConsistentHash struct {
    ring       map[uint32]string
    sortedKeys []uint32
    replicas   int
    mu         sync.RWMutex
}

func NewConsistentHash(replicas int) *ConsistentHash {
    return &ConsistentHash{
        ring:     make(map[uint32]string),
        replicas: replicas,
    }
}

func (c *ConsistentHash) hash(key string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(key))
    return h.Sum32()
}

func (c *ConsistentHash) AddNode(node string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    for i := 0; i < c.replicas; i++ {
        virtualKey := fmt.Sprintf("%s-%d", node, i)
        hash := c.hash(virtualKey)
        c.ring[hash] = node
        c.sortedKeys = append(c.sortedKeys, hash)
    }
    
    sort.Slice(c.sortedKeys, func(i, j int) bool {
        return c.sortedKeys[i] < c.sortedKeys[j]
    })
}

func (c *ConsistentHash) RemoveNode(node string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    for i := 0; i < c.replicas; i++ {
        virtualKey := fmt.Sprintf("%s-%d", node, i)
        hash := c.hash(virtualKey)
        delete(c.ring, hash)
    }
    
    // Rebuild sorted keys
    c.sortedKeys = c.sortedKeys[:0]
    for hash := range c.ring {
        c.sortedKeys = append(c.sortedKeys, hash)
    }
    sort.Slice(c.sortedKeys, func(i, j int) bool {
        return c.sortedKeys[i] < c.sortedKeys[j]
    })
}

func (c *ConsistentHash) GetNode(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    if len(c.sortedKeys) == 0 {
        return ""
    }
    
    hash := c.hash(key)
    
    // Binary search for the first node with hash >= key hash
    idx := sort.Search(len(c.sortedKeys), func(i int) bool {
        return c.sortedKeys[i] >= hash
    })
    
    // Wrap around to first node if we've passed all nodes
    if idx >= len(c.sortedKeys) {
        idx = 0
    }
    
    return c.ring[c.sortedKeys[idx]]
}
```

---

## 3. Rate Limiting

### Token Bucket Algorithm

```go
type TokenBucket struct {
    capacity   int64
    tokens     int64
    refillRate int64        // tokens per second
    lastRefill time.Time
    mu         sync.Mutex
}

func NewTokenBucket(capacity, refillRate int64) *TokenBucket {
    return &TokenBucket{
        capacity:   capacity,
        tokens:     capacity,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()
    
    // Refill tokens based on time elapsed
    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.tokens += int64(elapsed * float64(tb.refillRate))
    if tb.tokens > tb.capacity {
        tb.tokens = tb.capacity
    }
    tb.lastRefill = now
    
    // Check if we can consume a token
    if tb.tokens > 0 {
        tb.tokens--
        return true
    }
    
    return false
}

// Distributed Rate Limiting with Redis
type RedisRateLimiter struct {
    client     *redis.Client
    key        string
    limit      int64
    window     time.Duration
}

func (rl *RedisRateLimiter) Allow(ctx context.Context) (bool, error) {
    script := `
        local key = KEYS[1]
        local limit = tonumber(ARGV[1])
        local window = tonumber(ARGV[2])
        local current = redis.call("INCR", key)
        if current == 1 then
            redis.call("EXPIRE", key, window)
        end
        if current > limit then
            return 0
        end
        return 1
    `
    
    result, err := rl.client.Eval(ctx, script, []string{rl.key}, rl.limit, int(rl.window.Seconds())).Int64()
    if err != nil {
        return false, err
    }
    
    return result == 1, nil
}
```

### Sliding Window Rate Limiting

```go
type SlidingWindowLimiter struct {
    client    *redis.Client
    key       string
    limit     int64
    window    time.Duration
}

func (swl *SlidingWindowLimiter) Allow(ctx context.Context) (bool, error) {
    script := `
        local key = KEYS[1]
        local limit = tonumber(ARGV[1])
        local window = tonumber(ARGV[2])
        local now = tonumber(ARGV[3])
        
        -- Remove old entries outside the window
        redis.call("ZREMRANGEBYSCORE", key, "-inf", now - window)
        
        -- Count current entries
        local count = redis.call("ZCARD", key)
        
        if count < limit then
            -- Add new entry with current timestamp as score
            redis.call("ZADD", key, now, now .. "-" .. math.random())
            redis.call("EXPIRE", key, window)
            return 1
        end
        
        return 0
    `
    
    now := time.Now().UnixMilli()
    result, err := swl.client.Eval(ctx, script, []string{swl.key}, 
        swl.limit, int(swl.window.Milliseconds()), now).Int64()
    if err != nil {
        return false, err
    }
    
    return result == 1, nil
}
```

---

# Cloud Architecture

## 1. Multi-Cloud Strategy

### Cloud-Agnostic Design

```go
// Abstract cloud provider interface
type CloudProvider interface {
    ObjectStorage() ObjectStorage
    MessageQueue() MessageQueue
    Cache() CacheService
    Database() DatabaseService
}

type ObjectStorage interface {
    Upload(ctx context.Context, bucket, key string, data io.Reader) error
    Download(ctx context.Context, bucket, key string) (io.ReadCloser, error)
    Delete(ctx context.Context, bucket, key string) error
    ListObjects(ctx context.Context, bucket, prefix string) ([]ObjectInfo, error)
}

// AWS Implementation
type AWSProvider struct {
    s3Client  *s3.Client
    sqsClient *sqs.Client
    elasticache *elasticache.Client
}

func (p *AWSProvider) ObjectStorage() ObjectStorage {
    return &S3Storage{client: p.s3Client}
}

// GCP Implementation
type GCPProvider struct {
    storageClient *storage.Client
    pubsubClient  *pubsub.Client
}

func (p *GCPProvider) ObjectStorage() ObjectStorage {
    return &GCSStorage{client: p.storageClient}
}

// Azure Implementation
type AzureProvider struct {
    blobClient  *azblob.Client
    queueClient *azqueue.Client
}

func (p *AzureProvider) ObjectStorage() ObjectStorage {
    return &AzureBlobStorage{client: p.blobClient}
}
```

### Infrastructure as Code

```hcl
# Terraform - Multi-Cloud Kubernetes
module "aws_eks" {
  source          = "./modules/eks"
  cluster_name    = "production-aws"
  region          = "us-east-1"
  node_count      = 5
  instance_types  = ["m5.xlarge"]
}

module "gcp_gke" {
  source          = "./modules/gke"
  cluster_name    = "production-gcp"
  region          = "us-central1"
  node_count      = 5
  machine_type    = "e2-standard-4"
}

module "azure_aks" {
  source          = "./modules/aks"
  cluster_name    = "production-azure"
  region          = "eastus"
  node_count      = 5
  vm_size         = "Standard_D4s_v3"
}

# Cross-cloud load balancing
resource "cloudflare_load_balancer" "global" {
  name    = "global-lb"
  
  default_pool_ids = [
    cloudflare_load_balancer_pool.aws.id,
    cloudflare_load_balancer_pool.gcp.id,
    cloudflare_load_balancer_pool.azure.id,
  ]
  
  fallback_pool_id = cloudflare_load_balancer_pool.aws.id
  
  steering_policy = "geo"
  
  region_pools {
    region   = "ENAM"  # Eastern North America
    pool_ids = [cloudflare_load_balancer_pool.aws.id]
  }
  
  region_pools {
    region   = "WNAM"  # Western North America
    pool_ids = [cloudflare_load_balancer_pool.gcp.id]
  }
  
  region_pools {
    region   = "WEUR"  # Western Europe
    pool_ids = [cloudflare_load_balancer_pool.azure.id]
  }
}
```

---

## 2. Serverless Architecture

### Event-Driven Serverless

```go
// AWS Lambda Handler
type LambdaHandler struct {
    db      *dynamodb.Client
    sns     *sns.Client
    metrics *cloudwatch.Client
}

func (h *LambdaHandler) HandleSQSEvent(ctx context.Context, event events.SQSEvent) error {
    for _, record := range event.Records {
        var order Order
        if err := json.Unmarshal([]byte(record.Body), &order); err != nil {
            log.Printf("Failed to parse message: %v", err)
            continue
        }
        
        // Process order
        if err := h.processOrder(ctx, &order); err != nil {
            log.Printf("Failed to process order %s: %v", order.ID, err)
            // Message will be retried or sent to DLQ
            return err
        }
        
        // Publish notification
        h.notifyOrderProcessed(ctx, &order)
    }
    
    return nil
}

func (h *LambdaHandler) processOrder(ctx context.Context, order *Order) error {
    // Update order in DynamoDB
    _, err := h.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
        TableName: aws.String("orders"),
        Key: map[string]types.AttributeValue{
            "id": &types.AttributeValueMemberS{Value: order.ID},
        },
        UpdateExpression: aws.String("SET #status = :status, #updated = :updated"),
        ExpressionAttributeNames: map[string]string{
            "#status":  "status",
            "#updated": "updated_at",
        },
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":status":  &types.AttributeValueMemberS{Value: "processed"},
            ":updated": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
        },
    })
    
    return err
}
```

### Serverless Patterns

```yaml
# AWS SAM Template
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31

Resources:
  # API Gateway
  ApiGateway:
    Type: AWS::Serverless::Api
    Properties:
      StageName: prod
      Auth:
        DefaultAuthorizer: JWTAuthorizer
        Authorizers:
          JWTAuthorizer:
            JwtConfiguration:
              issuer: !Ref JwtIssuer
              audience:
                - !Ref JwtAudience

  # Order Processing Function
  OrderFunction:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: ./functions/order
      Handler: main
      Runtime: provided.al2
      MemorySize: 256
      Timeout: 30
      Events:
        Api:
          Type: Api
          Properties:
            RestApiId: !Ref ApiGateway
            Path: /orders
            Method: POST
        SQS:
          Type: SQS
          Properties:
            Queue: !GetAtt OrderQueue.Arn
            BatchSize: 10

  # DynamoDB Table
  OrdersTable:
    Type: AWS::DynamoDB::Table
    Properties:
      TableName: orders
      BillingMode: PAY_PER_REQUEST
      AttributeDefinitions:
        - AttributeName: id
          AttributeType: S
        - AttributeName: user_id
          AttributeType: S
      KeySchema:
        - AttributeName: id
          KeyType: HASH
      GlobalSecondaryIndexes:
        - IndexName: user-orders-index
          KeySchema:
            - AttributeName: user_id
              KeyType: HASH
          Projection:
            ProjectionType: ALL

  # SQS Queue with DLQ
  OrderQueue:
    Type: AWS::SQS::Queue
    Properties:
      QueueName: order-queue
      VisibilityTimeout: 60
      RedrivePolicy:
        deadLetterTargetArn: !GetAtt OrderDLQ.Arn
        maxReceiveCount: 3

  OrderDLQ:
    Type: AWS::SQS::Queue
    Properties:
      QueueName: order-dlq
```

---

# Security Architecture

## 1. Zero Trust Architecture

### Principles

```
┌─────────────────────────────────────────────────────────────────┐
│                    Zero Trust Principles                         │
├─────────────────────────────────────────────────────────────────┤
│ 1. Never trust, always verify                                    │
│ 2. Least privilege access                                        │
│ 3. Assume breach                                                 │
│ 4. Verify explicitly                                             │
│ 5. Secure all communications                                     │
└─────────────────────────────────────────────────────────────────┘

                    ┌─────────────────┐
                    │   Identity      │
                    │   Provider      │
                    └────────┬────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│    Users     │    │   Devices    │    │  Services    │
│              │    │              │    │              │
└──────┬───────┘    └──────┬───────┘    └──────┬───────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
                           ▼
              ┌─────────────────────────┐
              │   Policy Engine         │
              │   (Authorization)       │
              └────────────┬────────────┘
                           │
                           ▼
              ┌─────────────────────────┐
              │   Resources             │
              │   (Data, Services)      │
              └─────────────────────────┘
```

### Implementation

```go
// JWT-based Authentication Middleware
type AuthMiddleware struct {
    jwtValidator *jwt.Validator
    policyEngine *PolicyEngine
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract token
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        token = strings.TrimPrefix(token, "Bearer ")
        
        // Validate JWT
        claims, err := m.jwtValidator.Validate(r.Context(), token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        // Verify device (if available)
        deviceID := r.Header.Get("X-Device-ID")
        if deviceID != "" {
            if err := m.verifyDevice(r.Context(), claims.Subject, deviceID); err != nil {
                http.Error(w, "Device not authorized", http.StatusForbidden)
                return
            }
        }
        
        // Check authorization
        resource := r.URL.Path
        action := r.Method
        if !m.policyEngine.IsAllowed(claims, resource, action) {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        
        // Add claims to context
        ctx := context.WithValue(r.Context(), "claims", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Policy Engine (OPA-based)
type PolicyEngine struct {
    opa *opa.Client
}

type PolicyInput struct {
    Subject  string                 `json:"subject"`
    Resource string                 `json:"resource"`
    Action   string                 `json:"action"`
    Context  map[string]interface{} `json:"context"`
}

func (pe *PolicyEngine) IsAllowed(claims *Claims, resource, action string) bool {
    input := PolicyInput{
        Subject:  claims.Subject,
        Resource: resource,
        Action:   action,
        Context: map[string]interface{}{
            "roles":  claims.Roles,
            "groups": claims.Groups,
            "time":   time.Now().Unix(),
        },
    }
    
    result, err := pe.opa.Query(context.Background(), "data.authz.allow", input)
    if err != nil {
        return false
    }
    
    allowed, _ := result.(bool)
    return allowed
}
```

---

## 2. Secrets Management

### HashiCorp Vault Integration

```go
// Vault client wrapper
type VaultSecrets struct {
    client *vault.Client
    mount  string
}

func NewVaultSecrets(addr, token, mount string) (*VaultSecrets, error) {
    config := vault.DefaultConfig()
    config.Address = addr
    
    client, err := vault.NewClient(config)
    if err != nil {
        return nil, err
    }
    
    client.SetToken(token)
    
    return &VaultSecrets{
        client: client,
        mount:  mount,
    }, nil
}

func (v *VaultSecrets) GetSecret(ctx context.Context, path string) (map[string]interface{}, error) {
    secret, err := v.client.KVv2(v.mount).Get(ctx, path)
    if err != nil {
        return nil, fmt.Errorf("failed to get secret: %w", err)
    }
    
    return secret.Data, nil
}

func (v *VaultSecrets) GetDatabaseCredentials(ctx context.Context, role string) (*DatabaseCreds, error) {
    secret, err := v.client.Logical().ReadWithContext(ctx, 
        fmt.Sprintf("database/creds/%s", role))
    if err != nil {
        return nil, err
    }
    
    return &DatabaseCreds{
        Username:  secret.Data["username"].(string),
        Password:  secret.Data["password"].(string),
        LeaseTTL:  time.Duration(secret.LeaseDuration) * time.Second,
        LeaseID:   secret.LeaseID,
    }, nil
}

// Dynamic secrets with auto-renewal
type SecretManager struct {
    vault     *VaultSecrets
    secrets   sync.Map
    stopChan  chan struct{}
}

func (sm *SecretManager) StartRenewal(ctx context.Context) {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            sm.renewSecrets(ctx)
        case <-sm.stopChan:
            return
        case <-ctx.Done():
            return
        }
    }
}

func (sm *SecretManager) renewSecrets(ctx context.Context) {
    sm.secrets.Range(func(key, value interface{}) bool {
        secret := value.(*ManagedSecret)
        
        // Renew if less than 50% of TTL remaining
        if time.Until(secret.ExpiresAt) < secret.TTL/2 {
            if err := sm.renewSecret(ctx, key.(string), secret); err != nil {
                log.Printf("Failed to renew secret %s: %v", key, err)
            }
        }
        
        return true
    })
}
```

---

## 3. Security Patterns

### Input Validation

```go
import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=12,max=128"`
    FirstName string `json:"first_name" validate:"required,min=1,max=100,alpha"`
    LastName  string `json:"last_name" validate:"required,min=1,max=100,alpha"`
    Age       int    `json:"age" validate:"gte=18,lte=150"`
    Role      string `json:"role" validate:"required,oneof=user admin moderator"`
}

var validate *validator.Validate

func init() {
    validate = validator.New()
    
    // Register custom validators
    validate.RegisterValidation("password_strength", func(fl validator.FieldLevel) bool {
        password := fl.Field().String()
        
        hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
        hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
        hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
        hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
        
        return hasUpper && hasLower && hasNumber && hasSpecial
    })
}

func ValidateRequest(req interface{}) error {
    if err := validate.Struct(req); err != nil {
        var validationErrors validator.ValidationErrors
        if errors.As(err, &validationErrors) {
            var errMsgs []string
            for _, e := range validationErrors {
                errMsgs = append(errMsgs, formatValidationError(e))
            }
            return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, "; "))
        }
        return err
    }
    return nil
}

func formatValidationError(e validator.FieldError) string {
    switch e.Tag() {
    case "required":
        return fmt.Sprintf("%s is required", e.Field())
    case "email":
        return fmt.Sprintf("%s must be a valid email", e.Field())
    case "min":
        return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
    case "max":
        return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
    default:
        return fmt.Sprintf("%s failed validation: %s", e.Field(), e.Tag())
    }
}
```

---

# Data Architecture

## 1. Database Patterns

### Database Per Service

```
┌─────────────────────────────────────────────────────────────────┐
│                    Microservices Architecture                    │
│                                                                  │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │ User Service │    │Order Service │    │Product Service│     │
│  └──────┬───────┘    └──────┬───────┘    └──────┬───────┘      │
│         │                   │                   │               │
│         ▼                   ▼                   ▼               │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │   User DB    │    │   Order DB   │    │  Product DB  │      │
│  │ (PostgreSQL) │    │  (MongoDB)   │    │   (MySQL)    │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
└─────────────────────────────────────────────────────────────────┘
```

### Polyglot Persistence

```go
// Repository interfaces for different data stores
type UserRepository interface {
    Get(ctx context.Context, id string) (*User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

// PostgreSQL implementation for transactional data
type PostgresUserRepo struct {
    db *sql.DB
}

func (r *PostgresUserRepo) Get(ctx context.Context, id string) (*User, error) {
    query := `SELECT id, email, name, created_at FROM users WHERE id = $1`
    row := r.db.QueryRowContext(ctx, query, id)
    
    var user User
    if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt); err != nil {
        return nil, err
    }
    
    return &user, nil
}

// Redis implementation for caching/sessions
type RedisSessionStore struct {
    client *redis.Client
}

func (r *RedisSessionStore) Get(ctx context.Context, sessionID string) (*Session, error) {
    data, err := r.client.Get(ctx, "session:"+sessionID).Bytes()
    if err != nil {
        return nil, err
    }
    
    var session Session
    if err := json.Unmarshal(data, &session); err != nil {
        return nil, err
    }
    
    return &session, nil
}

// Elasticsearch for search
type ElasticsearchProductSearch struct {
    client *elasticsearch.Client
}

func (e *ElasticsearchProductSearch) Search(ctx context.Context, query string, filters map[string]interface{}) ([]Product, error) {
    searchBody := map[string]interface{}{
        "query": map[string]interface{}{
            "bool": map[string]interface{}{
                "must": []interface{}{
                    map[string]interface{}{
                        "multi_match": map[string]interface{}{
                            "query":  query,
                            "fields": []string{"name^3", "description", "category"},
                        },
                    },
                },
                "filter": buildFilters(filters),
            },
        },
    }
    
    res, err := e.client.Search(
        e.client.Search.WithContext(ctx),
        e.client.Search.WithIndex("products"),
        e.client.Search.WithBody(esutil.NewJSONReader(searchBody)),
    )
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    
    // Parse response
    var result struct {
        Hits struct {
            Hits []struct {
                Source Product `json:"_source"`
            } `json:"hits"`
        } `json:"hits"`
    }
    
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return nil, err
    }
    
    products := make([]Product, len(result.Hits.Hits))
    for i, hit := range result.Hits.Hits {
        products[i] = hit.Source
    }
    
    return products, nil
}
```

---

## 2. Data Streaming

### Apache Kafka Patterns

```go
// Kafka Producer with Idempotency
type KafkaProducer struct {
    producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true
    config.Producer.Idempotent = true
    config.Net.MaxOpenRequests = 1
    
    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }
    
    return &KafkaProducer{producer: producer}, nil
}

func (p *KafkaProducer) Publish(ctx context.Context, topic string, key string, event interface{}) error {
    data, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Key:   sarama.StringEncoder(key),
        Value: sarama.ByteEncoder(data),
        Headers: []sarama.RecordHeader{
            {Key: []byte("event-type"), Value: []byte(reflect.TypeOf(event).Name())},
            {Key: []byte("timestamp"), Value: []byte(time.Now().Format(time.RFC3339))},
            {Key: []byte("correlation-id"), Value: []byte(getCorrelationID(ctx))},
        },
    }
    
    _, _, err = p.producer.SendMessage(msg)
    return err
}

// Kafka Consumer with Exactly-Once Processing
type KafkaConsumer struct {
    consumer sarama.ConsumerGroup
    handlers map[string]EventHandler
}

type EventHandler func(ctx context.Context, msg *sarama.ConsumerMessage) error

func (c *KafkaConsumer) Start(ctx context.Context, topics []string) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            if err := c.consumer.Consume(ctx, topics, c); err != nil {
                log.Printf("Consumer error: %v", err)
            }
        }
    }
}

func (c *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for msg := range claim.Messages() {
        eventType := string(getHeader(msg.Headers, "event-type"))
        handler, ok := c.handlers[eventType]
        if !ok {
            log.Printf("No handler for event type: %s", eventType)
            session.MarkMessage(msg, "")
            continue
        }
        
        ctx := context.WithValue(context.Background(), 
            "correlation-id", string(getHeader(msg.Headers, "correlation-id")))
        
        if err := handler(ctx, msg); err != nil {
            log.Printf("Failed to process message: %v", err)
            // Don't mark as processed - will be redelivered
            return err
        }
        
        session.MarkMessage(msg, "")
    }
    
    return nil
}
```

---

## 3. Data Lake Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                      Data Sources                                │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐            │
│  │  Apps   │  │   IoT   │  │  Logs   │  │External │            │
│  └────┬────┘  └────┬────┘  └────┬────┘  └────┬────┘            │
└───────┼────────────┼────────────┼────────────┼──────────────────┘
        │            │            │            │
        └────────────┴────────────┴────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Ingestion Layer                               │
│            (Kafka, Kinesis, Event Hubs)                          │
└───────────────────────────┬─────────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        ▼                   ▼                   ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│    Bronze    │───▶│    Silver    │───▶│     Gold     │
│  (Raw Data)  │    │ (Cleaned)    │    │ (Aggregated) │
│              │    │              │    │              │
│ - As-is data │    │ - Deduped    │    │ - Business   │
│ - No schema  │    │ - Validated  │    │   metrics    │
│ - Immutable  │    │ - Typed      │    │ - Aggregates │
└──────────────┘    └──────────────┘    └──────────────┘
        │                   │                   │
        └───────────────────┼───────────────────┘
                            │
                            ▼
              ┌─────────────────────────┐
              │    Consumption Layer     │
              │  (BI, ML, Analytics)     │
              └─────────────────────────┘
```

---

# DevOps & Platform Engineering

## 1. CI/CD Pipeline Design

### GitOps Pipeline

```yaml
# .github/workflows/deploy.yaml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Run unit tests
        run: go test -v -race -coverprofile=coverage.out ./...
      
      - name: Run integration tests
        run: go test -v -tags=integration ./...
      
      - name: Security scan
        uses: securego/gosec@master
        with:
          args: ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      
      - name: Login to Registry
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            ghcr.io/${{ github.repository }}:${{ github.sha }}
            ghcr.io/${{ github.repository }}:latest
          cache-from: type=gha
          cache-to: type=gha,mode=max

  deploy:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Checkout manifests
        uses: actions/checkout@v4
        with:
          repository: org/k8s-manifests
          token: ${{ secrets.DEPLOY_TOKEN }}
      
      - name: Update image tag
        run: |
          cd apps/production
          kustomize edit set image app=ghcr.io/${{ github.repository }}:${{ github.sha }}
      
      - name: Commit and push
        run: |
          git config user.name "GitHub Actions"
          git config user.email "actions@github.com"
          git commit -am "Deploy ${{ github.repository }}:${{ github.sha }}"
          git push
```

### Progressive Delivery

```go
// Canary deployment controller
type CanaryController struct {
    client    kubernetes.Interface
    metrics   MetricsClient
}

type CanaryConfig struct {
    StableName    string
    CanaryName    string
    Namespace     string
    Steps         []CanaryStep
    SuccessRate   float64  // Minimum success rate to proceed
    LatencyP99    float64  // Maximum P99 latency to proceed
}

type CanaryStep struct {
    Weight   int
    Duration time.Duration
}

func (c *CanaryController) Execute(ctx context.Context, config *CanaryConfig) error {
    for _, step := range config.Steps {
        // Update traffic weight
        if err := c.updateWeight(ctx, config, step.Weight); err != nil {
            return c.rollback(ctx, config, err)
        }
        
        // Wait for stabilization
        time.Sleep(30 * time.Second)
        
        // Analyze metrics
        if err := c.analyzeMetrics(ctx, config, step.Duration); err != nil {
            return c.rollback(ctx, config, err)
        }
        
        log.Printf("Step completed: %d%% traffic to canary", step.Weight)
    }
    
    // Promote canary to stable
    return c.promote(ctx, config)
}

func (c *CanaryController) analyzeMetrics(ctx context.Context, config *CanaryConfig, duration time.Duration) error {
    deadline := time.Now().Add(duration)
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-ticker.C:
            // Check success rate
            successRate, err := c.metrics.GetSuccessRate(ctx, config.CanaryName, config.Namespace)
            if err != nil {
                return err
            }
            if successRate < config.SuccessRate {
                return fmt.Errorf("success rate %.2f%% below threshold %.2f%%", 
                    successRate*100, config.SuccessRate*100)
            }
            
            // Check latency
            latency, err := c.metrics.GetP99Latency(ctx, config.CanaryName, config.Namespace)
            if err != nil {
                return err
            }
            if latency > config.LatencyP99 {
                return fmt.Errorf("P99 latency %.2fms above threshold %.2fms", 
                    latency*1000, config.LatencyP99*1000)
            }
            
            if time.Now().After(deadline) {
                return nil
            }
        }
    }
}
```

---

## 2. Observability

### Three Pillars of Observability

```
                    ┌───────────────────────────────────────┐
                    │          Observability Platform       │
                    └───────────────────────────────────────┘
                                       │
        ┌──────────────────────────────┼──────────────────────────────┐
        │                              │                              │
        ▼                              ▼                              ▼
┌──────────────────┐        ┌──────────────────┐        ┌──────────────────┐
│     Metrics      │        │      Logs        │        │     Traces       │
│   (Prometheus)   │        │     (Loki)       │        │    (Jaeger)      │
├──────────────────┤        ├──────────────────┤        ├──────────────────┤
│ - Request rate   │        │ - Application    │        │ - Request flow   │
│ - Error rate     │        │   logs           │        │ - Latency        │
│ - Latency        │        │ - Access logs    │        │   breakdown      │
│ - Saturation     │        │ - Error traces   │        │ - Dependencies   │
└──────────────────┘        └──────────────────┘        └──────────────────┘
```

### OpenTelemetry Integration

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
    "go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
)

func InitTelemetry(ctx context.Context, serviceName string) (func(), error) {
    // Create resource
    res, err := resource.New(ctx,
        resource.WithAttributes(
            attribute.String("service.name", serviceName),
            attribute.String("service.version", version),
        ),
    )
    if err != nil {
        return nil, err
    }
    
    // Setup trace exporter
    traceExporter, err := otlptrace.New(ctx,
        otlptracegrpc.NewClient(
            otlptracegrpc.WithEndpoint("otel-collector:4317"),
            otlptracegrpc.WithInsecure(),
        ),
    )
    if err != nil {
        return nil, err
    }
    
    // Setup tracer provider
    tp := trace.NewTracerProvider(
        trace.WithBatcher(traceExporter),
        trace.WithResource(res),
        trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.1))),
    )
    otel.SetTracerProvider(tp)
    
    // Setup metric exporter
    metricExporter, err := otlpmetric.New(ctx,
        otlpmetricgrpc.NewClient(
            otlpmetricgrpc.WithEndpoint("otel-collector:4317"),
            otlpmetricgrpc.WithInsecure(),
        ),
    )
    if err != nil {
        return nil, err
    }
    
    mp := metric.NewMeterProvider(
        metric.WithReader(metric.NewPeriodicReader(metricExporter)),
        metric.WithResource(res),
    )
    otel.SetMeterProvider(mp)
    
    return func() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        tp.Shutdown(ctx)
        mp.Shutdown(ctx)
    }, nil
}

// Instrumented HTTP Handler
func InstrumentedHandler(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, span := otel.Tracer("http").Start(r.Context(), r.URL.Path)
        defer span.End()
        
        span.SetAttributes(
            attribute.String("http.method", r.Method),
            attribute.String("http.url", r.URL.String()),
            attribute.String("http.user_agent", r.UserAgent()),
        )
        
        // Wrap response writer to capture status
        rw := &responseWriter{ResponseWriter: w, statusCode: 200}
        
        start := time.Now()
        handler.ServeHTTP(rw, r.WithContext(ctx))
        duration := time.Since(start)
        
        span.SetAttributes(
            attribute.Int("http.status_code", rw.statusCode),
            attribute.Float64("http.duration_ms", float64(duration.Milliseconds())),
        )
        
        if rw.statusCode >= 400 {
            span.SetStatus(codes.Error, http.StatusText(rw.statusCode))
        }
    })
}
```

---

# Non-Functional Requirements

## 1. NFR Categories

| Category | Metrics | Targets |
|----------|---------|---------|
| **Performance** | Response time, throughput | P95 < 200ms, 10k TPS |
| **Availability** | Uptime, MTTR | 99.99%, MTTR < 30min |
| **Scalability** | Concurrent users, data volume | 1M users, 100TB data |
| **Security** | Vulnerabilities, compliance | Zero critical CVEs, SOC2 |
| **Maintainability** | MTTR, deployment frequency | < 1hr fix, daily deploys |
| **Reliability** | Error rate, data durability | < 0.01%, 99.999999999% |

## 2. Capacity Planning

```go
// Capacity planning model
type CapacityModel struct {
    BaseLoad         float64 // Requests per second at base load
    PeakMultiplier   float64 // Peak vs base load ratio
    GrowthRate       float64 // Monthly growth rate
    PlanningHorizon  int     // Months to plan for
}

func (m *CapacityModel) CalculateCapacity() *CapacityPlan {
    // Calculate peak load with growth
    peakRPS := m.BaseLoad * m.PeakMultiplier
    
    for i := 0; i < m.PlanningHorizon; i++ {
        peakRPS *= (1 + m.GrowthRate)
    }
    
    // Calculate infrastructure needs
    instanceRPS := 1000.0 // RPS per instance
    instances := int(math.Ceil(peakRPS / instanceRPS))
    
    // Add 20% buffer
    instances = int(float64(instances) * 1.2)
    
    // Calculate resources
    cpuPerInstance := 4.0    // vCPUs
    memPerInstance := 8.0    // GB
    
    return &CapacityPlan{
        PeakRPS:      peakRPS,
        Instances:    instances,
        TotalCPU:     float64(instances) * cpuPerInstance,
        TotalMemory:  float64(instances) * memPerInstance,
        EstimatedCost: calculateCost(instances, cpuPerInstance, memPerInstance),
    }
}

type CapacityPlan struct {
    PeakRPS       float64
    Instances     int
    TotalCPU      float64
    TotalMemory   float64
    EstimatedCost float64
}
```

---

# Architectural Decision Making

## 1. Architecture Decision Records (ADR)

```markdown
# ADR-001: Use Event-Driven Architecture for Order Processing

## Status
Accepted

## Context
Our e-commerce platform processes 50,000 orders per day with peaks of 500 orders per minute.
The current synchronous architecture causes timeout issues during peak loads.

## Decision
We will implement an event-driven architecture using Apache Kafka for order processing:
- Orders are published to Kafka topics
- Microservices consume and process events asynchronously
- Event sourcing for order state management

## Consequences

### Positive
- Improved resilience during peak loads
- Better decoupling between services
- Natural audit trail through event log
- Easier horizontal scaling

### Negative
- Increased complexity in debugging
- Eventual consistency (not immediate)
- New infrastructure (Kafka cluster)
- Team needs training on event-driven patterns

### Risks
- Data consistency between services
- Message ordering guarantees
- Dead letter queue management

## Alternatives Considered
1. **Scale synchronous architecture** - Rejected due to coupling issues
2. **Use AWS SQS** - Rejected due to ordering requirements
3. **RabbitMQ** - Considered but Kafka chosen for higher throughput

## Implementation
See ADR-002 for Kafka cluster setup
See ADR-003 for event schema design
```

## 2. Trade-off Analysis

```
                    Cost vs Performance Trade-offs
                    
    High Performance
         │
         │     ┌─────────────┐
         │     │  Premium    │ ←── Best performance, highest cost
         │     │  Solution   │     (Multi-region active-active)
         │     └─────────────┘
         │
         │          ┌─────────────┐
         │          │  Balanced   │ ←── Good performance, reasonable cost
         │          │  Solution   │     (Single region with DR)
         │          └─────────────┘
         │
         │               ┌─────────────┐
         │               │   Budget    │ ←── Acceptable performance, low cost
         │               │  Solution   │     (Single region, no redundancy)
         │               └─────────────┘
         │
    Low Performance ────────────────────────── High Cost
                    Low Cost
```

---

# Case Studies & System Designs

## Case Study 1: Design a URL Shortener

### Requirements
- Shorten long URLs to short codes
- Handle 100M URLs created per month
- 10:1 read:write ratio
- URLs expire after configurable time
- Analytics on click-through

### Architecture

```
                    ┌─────────────────────────────────────┐
                    │           CDN (CloudFlare)          │
                    └─────────────────┬───────────────────┘
                                      │
                    ┌─────────────────┴───────────────────┐
                    │           Load Balancer             │
                    └─────────────────┬───────────────────┘
                                      │
        ┌─────────────────────────────┼─────────────────────────────┐
        │                             │                             │
        ▼                             ▼                             ▼
┌──────────────┐            ┌──────────────┐            ┌──────────────┐
│  API Server  │            │  API Server  │            │  API Server  │
└──────┬───────┘            └──────┬───────┘            └──────┬───────┘
       │                           │                           │
       └───────────────────────────┼───────────────────────────┘
                                   │
              ┌────────────────────┼────────────────────┐
              │                    │                    │
              ▼                    ▼                    ▼
       ┌──────────────┐    ┌──────────────┐    ┌──────────────┐
       │ Redis Cache  │    │  PostgreSQL  │    │    Kafka     │
       │  (Hot URLs)  │    │   (Primary)  │    │ (Analytics)  │
       └──────────────┘    └──────────────┘    └──────────────┘
```

### Implementation

```go
// URL Shortener Service
type URLShortener struct {
    db     *sql.DB
    cache  *redis.Client
    kafka  *KafkaProducer
    idGen  *SnowflakeGenerator
}

type URL struct {
    ID          int64     `json:"id"`
    ShortCode   string    `json:"short_code"`
    OriginalURL string    `json:"original_url"`
    UserID      string    `json:"user_id"`
    ExpiresAt   time.Time `json:"expires_at"`
    CreatedAt   time.Time `json:"created_at"`
}

func (s *URLShortener) Shorten(ctx context.Context, originalURL, userID string, ttl time.Duration) (*URL, error) {
    // Validate URL
    if _, err := url.ParseRequestURI(originalURL); err != nil {
        return nil, ErrInvalidURL
    }
    
    // Generate unique ID
    id := s.idGen.Generate()
    shortCode := base62Encode(id)
    
    url := &URL{
        ID:          id,
        ShortCode:   shortCode,
        OriginalURL: originalURL,
        UserID:      userID,
        ExpiresAt:   time.Now().Add(ttl),
        CreatedAt:   time.Now(),
    }
    
    // Store in database
    if err := s.saveURL(ctx, url); err != nil {
        return nil, err
    }
    
    // Cache for fast lookups
    s.cacheURL(ctx, url)
    
    return url, nil
}

func (s *URLShortener) Resolve(ctx context.Context, shortCode string) (string, error) {
    // Check cache first
    cached, err := s.cache.Get(ctx, "url:"+shortCode).Result()
    if err == nil {
        // Record click asynchronously
        go s.recordClick(shortCode)
        return cached, nil
    }
    
    // Fetch from database
    url, err := s.getURL(ctx, shortCode)
    if err != nil {
        return "", ErrURLNotFound
    }
    
    // Check expiration
    if time.Now().After(url.ExpiresAt) {
        return "", ErrURLExpired
    }
    
    // Cache for future requests
    s.cacheURL(ctx, url)
    
    // Record click
    go s.recordClick(shortCode)
    
    return url.OriginalURL, nil
}

// Base62 encoding for short codes
func base62Encode(num int64) string {
    const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    if num == 0 {
        return string(charset[0])
    }
    
    var result []byte
    for num > 0 {
        result = append([]byte{charset[num%62]}, result...)
        num /= 62
    }
    
    return string(result)
}
```

---

## Case Study 2: Design a Notification System

### Requirements
- Support multiple channels (email, SMS, push, in-app)
- Handle 1M notifications per hour
- Template-based messages
- Rate limiting per user
- Delivery tracking and analytics

### Architecture

```yaml
# Notification System Architecture

components:
  api_gateway:
    type: nginx
    features:
      - rate_limiting
      - authentication
      - request_routing
  
  notification_service:
    type: golang_service
    replicas: 10
    features:
      - template_rendering
      - priority_routing
      - deduplication
  
  channel_workers:
    email:
      replicas: 5
      provider: sendgrid
    sms:
      replicas: 3
      provider: twilio
    push:
      replicas: 3
      providers: [fcm, apns]
    in_app:
      replicas: 2
      storage: redis
  
  message_queue:
    type: kafka
    topics:
      - notifications.email
      - notifications.sms
      - notifications.push
      - notifications.in_app
    retention: 7d
  
  storage:
    templates:
      type: mongodb
    delivery_logs:
      type: clickhouse
    preferences:
      type: postgresql
```

### Implementation

```go
// Notification Service
type NotificationService struct {
    templateStore  TemplateStore
    preferenceRepo PreferenceRepository
    producer       *KafkaProducer
    rateLimiter    *RateLimiter
}

type NotificationRequest struct {
    UserID      string                 `json:"user_id"`
    Type        string                 `json:"type"`
    TemplateID  string                 `json:"template_id"`
    Data        map[string]interface{} `json:"data"`
    Priority    Priority               `json:"priority"`
    Channels    []Channel              `json:"channels"`
    ScheduledAt *time.Time             `json:"scheduled_at,omitempty"`
}

func (s *NotificationService) Send(ctx context.Context, req *NotificationRequest) (*NotificationResponse, error) {
    // Check rate limit
    if !s.rateLimiter.Allow(req.UserID, req.Type) {
        return nil, ErrRateLimited
    }
    
    // Get user preferences
    prefs, err := s.preferenceRepo.Get(ctx, req.UserID)
    if err != nil {
        return nil, err
    }
    
    // Filter channels based on preferences
    channels := s.filterChannels(req.Channels, prefs)
    if len(channels) == 0 {
        return &NotificationResponse{Status: "skipped", Reason: "user_opted_out"}, nil
    }
    
    // Get and render template
    template, err := s.templateStore.Get(ctx, req.TemplateID)
    if err != nil {
        return nil, err
    }
    
    // Create notification events for each channel
    notificationID := uuid.New().String()
    events := make([]*NotificationEvent, len(channels))
    
    for i, channel := range channels {
        content, err := s.renderTemplate(template, channel, req.Data)
        if err != nil {
            return nil, err
        }
        
        events[i] = &NotificationEvent{
            ID:             uuid.New().String(),
            NotificationID: notificationID,
            UserID:         req.UserID,
            Channel:        channel,
            Content:        content,
            Priority:       req.Priority,
            ScheduledAt:    req.ScheduledAt,
            CreatedAt:      time.Now(),
        }
    }
    
    // Publish to appropriate topics
    for _, event := range events {
        topic := fmt.Sprintf("notifications.%s", event.Channel)
        if err := s.producer.Publish(ctx, topic, event.UserID, event); err != nil {
            log.Printf("Failed to publish notification: %v", err)
        }
    }
    
    return &NotificationResponse{
        NotificationID: notificationID,
        Status:         "queued",
        Channels:       channels,
    }, nil
}

// Email Worker
type EmailWorker struct {
    emailClient  EmailClient
    deliveryRepo DeliveryRepository
}

func (w *EmailWorker) Process(ctx context.Context, event *NotificationEvent) error {
    // Send email
    messageID, err := w.emailClient.Send(ctx, &Email{
        To:      event.Content.Recipient,
        Subject: event.Content.Subject,
        Body:    event.Content.Body,
    })
    
    // Record delivery attempt
    delivery := &DeliveryLog{
        NotificationID: event.NotificationID,
        Channel:        "email",
        Status:         "delivered",
        MessageID:      messageID,
        Timestamp:      time.Now(),
    }
    
    if err != nil {
        delivery.Status = "failed"
        delivery.Error = err.Error()
    }
    
    return w.deliveryRepo.Save(ctx, delivery)
}
```

---

# Practice Questions with Answers

## System Design Questions

### Q1: Design a rate limiter for an API gateway.

**Answer:**

**Requirements Analysis:**
- Handle 10,000+ RPS
- Multiple rate limiting strategies (per user, per IP, global)
- Distributed across multiple nodes
- Minimal latency impact

**Design:**

```go
// Distributed Rate Limiter using Redis
type DistributedRateLimiter struct {
    client *redis.Client
}

type RateLimitConfig struct {
    Key      string        // e.g., "user:123" or "ip:10.0.0.1"
    Limit    int64         // Requests per window
    Window   time.Duration // Time window
    Strategy string        // "sliding_window" or "token_bucket"
}

func (rl *DistributedRateLimiter) SlidingWindowAllow(ctx context.Context, config RateLimitConfig) (*RateLimitResult, error) {
    now := time.Now().UnixMilli()
    windowStart := now - config.Window.Milliseconds()
    
    script := `
        local key = KEYS[1]
        local now = tonumber(ARGV[1])
        local window_start = tonumber(ARGV[2])
        local limit = tonumber(ARGV[3])
        local window_ms = tonumber(ARGV[4])
        
        -- Remove old entries
        redis.call("ZREMRANGEBYSCORE", key, "-inf", window_start)
        
        -- Count current entries
        local count = redis.call("ZCARD", key)
        
        if count < limit then
            -- Add new entry
            redis.call("ZADD", key, now, now .. "-" .. math.random())
            redis.call("PEXPIRE", key, window_ms)
            return {1, limit - count - 1, 0}
        end
        
        -- Get oldest entry for retry-after
        local oldest = redis.call("ZRANGE", key, 0, 0, "WITHSCORES")
        local retry_after = 0
        if #oldest > 0 then
            retry_after = tonumber(oldest[2]) + window_ms - now
        end
        
        return {0, 0, retry_after}
    `
    
    result, err := rl.client.Eval(ctx, script, []string{config.Key},
        now, windowStart, config.Limit, config.Window.Milliseconds()).Slice()
    if err != nil {
        return nil, err
    }
    
    return &RateLimitResult{
        Allowed:    result[0].(int64) == 1,
        Remaining:  result[1].(int64),
        RetryAfter: time.Duration(result[2].(int64)) * time.Millisecond,
    }, nil
}

// HTTP Middleware
func RateLimitMiddleware(limiter *DistributedRateLimiter, config RateLimitConfig) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Determine rate limit key
            key := config.Key
            if key == "" {
                key = getUserID(r)
                if key == "" {
                    key = getClientIP(r)
                }
            }
            
            result, err := limiter.SlidingWindowAllow(r.Context(), RateLimitConfig{
                Key:    key,
                Limit:  config.Limit,
                Window: config.Window,
            })
            
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
            
            // Set rate limit headers
            w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", config.Limit))
            w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", result.Remaining))
            
            if !result.Allowed {
                w.Header().Set("Retry-After", fmt.Sprintf("%d", int(result.RetryAfter.Seconds())))
                http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

---

### Q2: How would you design a distributed cache system?

**Answer:**

**Architecture:**

```
┌─────────────────────────────────────────────────────────────────┐
│                      Application Layer                           │
└───────────────────────────┬─────────────────────────────────────┘
                            │
                            ▼
              ┌─────────────────────────┐
              │     Cache Client        │
              │  (Consistent Hashing)   │
              └────────────┬────────────┘
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
        ▼                  ▼                  ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│  Cache Node  │  │  Cache Node  │  │  Cache Node  │
│     #1       │  │     #2       │  │     #3       │
│  (Primary)   │  │  (Primary)   │  │  (Primary)   │
└──────┬───────┘  └──────┬───────┘  └──────┬───────┘
       │                 │                 │
       ▼                 ▼                 ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│   Replica    │  │   Replica    │  │   Replica    │
│    #1        │  │    #2        │  │    #3        │
└──────────────┘  └──────────────┘  └──────────────┘
```

**Implementation:**

```go
type DistributedCache struct {
    ring       *ConsistentHash
    nodes      map[string]*CacheNode
    replicaFactor int
    mu         sync.RWMutex
}

type CacheNode struct {
    address string
    client  *redis.Client
    primary bool
    replica *CacheNode
}

func (c *DistributedCache) Get(ctx context.Context, key string) ([]byte, error) {
    node := c.getNode(key)
    
    // Try primary first
    val, err := node.client.Get(ctx, key).Bytes()
    if err == nil {
        return val, nil
    }
    
    // Fallback to replica
    if node.replica != nil && errors.Is(err, redis.Nil) == false {
        return node.replica.client.Get(ctx, key).Bytes()
    }
    
    return nil, err
}

func (c *DistributedCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
    node := c.getNode(key)
    
    // Write to primary
    if err := node.client.Set(ctx, key, value, ttl).Err(); err != nil {
        return err
    }
    
    // Async replicate
    if node.replica != nil {
        go func() {
            if err := node.replica.client.Set(context.Background(), key, value, ttl).Err(); err != nil {
                log.Printf("Replication failed for key %s: %v", key, err)
            }
        }()
    }
    
    return nil
}

func (c *DistributedCache) getNode(key string) *CacheNode {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    nodeAddress := c.ring.GetNode(key)
    return c.nodes[nodeAddress]
}

// Cache Invalidation Strategy
type CacheInvalidator struct {
    cache    *DistributedCache
    pubsub   *redis.PubSub
    patterns map[string][]string // tag -> keys mapping
}

func (i *CacheInvalidator) InvalidateByTag(ctx context.Context, tag string) error {
    keys := i.patterns[tag]
    if len(keys) == 0 {
        return nil
    }
    
    // Publish invalidation message
    msg := &InvalidationMessage{
        Tag:  tag,
        Keys: keys,
        Time: time.Now(),
    }
    
    data, _ := json.Marshal(msg)
    return i.pubsub.Publish(ctx, "cache:invalidate", data).Err()
}
```

---

### Q3: Design a system to handle 1 million concurrent WebSocket connections.

**Answer:**

**Architecture:**

```
┌─────────────────────────────────────────────────────────────────┐
│                      Load Balancer                               │
│                   (Layer 4, Sticky Sessions)                     │
└───────────────────────────┬─────────────────────────────────────┘
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
        ▼                   ▼                   ▼
┌──────────────┐    ┌──────────────┐    ┌──────────────┐
│  WS Server   │    │  WS Server   │    │  WS Server   │
│  (100k conn) │    │  (100k conn) │    │  (100k conn) │
└──────┬───────┘    └──────┬───────┘    └──────┬───────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
              ┌────────────┴────────────┐
              │      Redis Pub/Sub      │
              │   (Message Fanout)      │
              └─────────────────────────┘
```

**Implementation:**

```go
type WebSocketHub struct {
    connections sync.Map           // userID -> *Connection
    rooms       sync.Map           // roomID -> map[string]*Connection
    pubsub      *redis.PubSub
    serverID    string
    maxConns    int
    currentConns atomic.Int64
}

type Connection struct {
    ID       string
    UserID   string
    Conn     *websocket.Conn
    RoomIDs  []string
    Send     chan []byte
    mu       sync.Mutex
}

func (h *WebSocketHub) HandleConnection(w http.ResponseWriter, r *http.Request) {
    // Check connection limit
    if h.currentConns.Load() >= int64(h.maxConns) {
        http.Error(w, "Server at capacity", http.StatusServiceUnavailable)
        return
    }
    
    // Upgrade connection
    upgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin:     func(r *http.Request) bool { return true },
    }
    
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    
    userID := getUserIDFromToken(r)
    c := &Connection{
        ID:     uuid.New().String(),
        UserID: userID,
        Conn:   conn,
        Send:   make(chan []byte, 256),
    }
    
    h.registerConnection(c)
    h.currentConns.Add(1)
    
    // Start read/write goroutines
    go h.readPump(c)
    go h.writePump(c)
}

func (h *WebSocketHub) readPump(c *Connection) {
    defer func() {
        h.unregisterConnection(c)
        h.currentConns.Add(-1)
        c.Conn.Close()
    }()
    
    c.Conn.SetReadLimit(4096)
    c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    c.Conn.SetPongHandler(func(string) error {
        c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })
    
    for {
        _, message, err := c.Conn.ReadMessage()
        if err != nil {
            break
        }
        
        h.handleMessage(c, message)
    }
}

func (h *WebSocketHub) writePump(c *Connection) {
    ticker := time.NewTicker(30 * time.Second)
    defer func() {
        ticker.Stop()
        c.Conn.Close()
    }()
    
    for {
        select {
        case message, ok := <-c.Send:
            if !ok {
                c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            
            c.mu.Lock()
            c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            err := c.Conn.WriteMessage(websocket.TextMessage, message)
            c.mu.Unlock()
            
            if err != nil {
                return
            }
            
        case <-ticker.C:
            c.mu.Lock()
            c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            err := c.Conn.WriteMessage(websocket.PingMessage, nil)
            c.mu.Unlock()
            
            if err != nil {
                return
            }
        }
    }
}

// Cross-server messaging via Redis Pub/Sub
func (h *WebSocketHub) BroadcastToRoom(roomID string, message []byte) {
    // Publish to Redis for other servers
    payload := &BroadcastMessage{
        RoomID:   roomID,
        Message:  message,
        ServerID: h.serverID,
    }
    
    data, _ := json.Marshal(payload)
    h.pubsub.Publish(context.Background(), "ws:broadcast:"+roomID, data)
    
    // Send to local connections
    h.sendToLocalRoom(roomID, message)
}
```

---

### Q4: Explain how to ensure data consistency in a distributed system.

**Answer:**

**Consistency Strategies:**

1. **Synchronous Replication:**
```go
func (s *Service) Write(ctx context.Context, data *Data) error {
    // Write to primary
    if err := s.primary.Write(ctx, data); err != nil {
        return err
    }
    
    // Write to all replicas synchronously
    var wg sync.WaitGroup
    errors := make(chan error, len(s.replicas))
    
    for _, replica := range s.replicas {
        wg.Add(1)
        go func(r *Replica) {
            defer wg.Done()
            if err := r.Write(ctx, data); err != nil {
                errors <- err
            }
        }(replica)
    }
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    for err := range errors {
        // Rollback or compensate
        s.rollback(ctx, data)
        return fmt.Errorf("replication failed: %w", err)
    }
    
    return nil
}
```

2. **Two-Phase Commit (2PC):**
```go
type TwoPhaseCommit struct {
    coordinator *Coordinator
    participants []*Participant
}

func (tpc *TwoPhaseCommit) Execute(ctx context.Context, tx *Transaction) error {
    // Phase 1: Prepare
    prepareCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    votes := make([]bool, len(tpc.participants))
    var wg sync.WaitGroup
    
    for i, p := range tpc.participants {
        wg.Add(1)
        go func(idx int, participant *Participant) {
            defer wg.Done()
            votes[idx] = participant.Prepare(prepareCtx, tx)
        }(i, p)
    }
    
    wg.Wait()
    
    // Check all votes
    allPrepared := true
    for _, vote := range votes {
        if !vote {
            allPrepared = false
            break
        }
    }
    
    // Phase 2: Commit or Abort
    if allPrepared {
        for _, p := range tpc.participants {
            if err := p.Commit(ctx, tx); err != nil {
                log.Printf("Commit failed: %v", err)
            }
        }
        return nil
    }
    
    // Abort all
    for _, p := range tpc.participants {
        p.Abort(ctx, tx)
    }
    return ErrTransactionAborted
}
```

3. **Saga Pattern with Compensation:**
```go
type Saga struct {
    steps []SagaStep
}

type SagaStep struct {
    Name       string
    Execute    func(ctx context.Context) error
    Compensate func(ctx context.Context) error
}

func (s *Saga) Run(ctx context.Context) error {
    executed := []SagaStep{}
    
    for _, step := range s.steps {
        if err := step.Execute(ctx); err != nil {
            // Compensate in reverse order
            for i := len(executed) - 1; i >= 0; i-- {
                if compErr := executed[i].Compensate(ctx); compErr != nil {
                    log.Printf("Compensation failed for %s: %v", 
                        executed[i].Name, compErr)
                }
            }
            return fmt.Errorf("saga failed at %s: %w", step.Name, err)
        }
        executed = append(executed, step)
    }
    
    return nil
}
```

---

### Q5: How do you design for high availability?

**Answer:**

**HA Architecture Patterns:**

```
┌─────────────────────────────────────────────────────────────────┐
│                    Multi-Region Architecture                     │
│                                                                  │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │                       Global DNS                             ││
│  │               (Route53 / Cloudflare)                        ││
│  │               Latency-based routing                         ││
│  └───────────────────────────┬─────────────────────────────────┘│
│                              │                                   │
│          ┌───────────────────┼───────────────────┐              │
│          │                   │                   │              │
│          ▼                   ▼                   ▼              │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │  US-East     │    │  EU-West     │    │  AP-South    │      │
│  │  (Active)    │◄──▶│  (Active)    │◄──▶│  (Active)    │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
│                                                                  │
│  Data Replication: Async multi-master with conflict resolution  │
└─────────────────────────────────────────────────────────────────┘
```

**Implementation:**

```go
// Health Check System
type HealthChecker struct {
    services map[string]*ServiceHealth
    alerter  Alerter
}

type ServiceHealth struct {
    Name           string
    Endpoints      []string
    HealthPath     string
    CheckInterval  time.Duration
    Timeout        time.Duration
    FailThreshold  int
    consecutiveFails int
    healthy        bool
    mu             sync.RWMutex
}

func (hc *HealthChecker) Start(ctx context.Context) {
    for name, service := range hc.services {
        go hc.monitorService(ctx, name, service)
    }
}

func (hc *HealthChecker) monitorService(ctx context.Context, name string, service *ServiceHealth) {
    ticker := time.NewTicker(service.CheckInterval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            healthy := hc.checkHealth(ctx, service)
            
            service.mu.Lock()
            if healthy {
                service.consecutiveFails = 0
                if !service.healthy {
                    service.healthy = true
                    hc.alerter.Send(Alert{
                        Severity: "info",
                        Message:  fmt.Sprintf("Service %s recovered", name),
                    })
                }
            } else {
                service.consecutiveFails++
                if service.consecutiveFails >= service.FailThreshold && service.healthy {
                    service.healthy = false
                    hc.alerter.Send(Alert{
                        Severity: "critical",
                        Message:  fmt.Sprintf("Service %s is unhealthy", name),
                    })
                }
            }
            service.mu.Unlock()
        }
    }
}

// Failover Manager
type FailoverManager struct {
    primary   *Endpoint
    secondary *Endpoint
    current   *Endpoint
    mu        sync.RWMutex
}

func (fm *FailoverManager) GetEndpoint() *Endpoint {
    fm.mu.RLock()
    defer fm.mu.RUnlock()
    return fm.current
}

func (fm *FailoverManager) Failover(ctx context.Context) error {
    fm.mu.Lock()
    defer fm.mu.Unlock()
    
    if fm.current == fm.primary {
        fm.current = fm.secondary
        log.Println("Failover: switched to secondary")
    } else {
        fm.current = fm.primary
        log.Println("Failback: switched to primary")
    }
    
    return nil
}
```

---

### Summary

**Key Solution Architect Skills:**

1. **System Design** - Design scalable, reliable systems
2. **Trade-off Analysis** - Balance competing requirements
3. **Technology Selection** - Choose appropriate technologies
4. **Security Architecture** - Implement defense in depth
5. **Data Architecture** - Design for consistency and performance
6. **Cloud Native** - Leverage cloud services effectively
7. **DevOps Integration** - Enable continuous delivery
8. **Documentation** - Create clear architectural artifacts
9. **Communication** - Bridge technical and business stakeholders
10. **Risk Management** - Identify and mitigate technical risks

**Architecture Principles:**

- Design for failure
- Keep it simple (KISS)
- Build for scale
- Automate everything
- Monitor and observe
- Security by design
- Document decisions (ADRs)
- Embrace change

---

*Document Version: 1.0*
*Last Updated: January 2025*
