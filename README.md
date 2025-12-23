# Solution Architect Learning Roadmap

> **Goal:** Become a credible End-to-End Solution Architect within 6-9 months
> **Current Profile:** 15 years experience, strong in React Native, Node.js, AWS, IoT
> **Gaps to Address:** Backend depth, Databases, Event-driven systems, DevOps/K8s
> **Learning Decision:** Go with Gin (industry standard, net/http compatible)

---

## 📊 Tech Stack Mastery Framework

### Tier 1: MUST MASTER (Non-negotiable for SA)

| Area | Technologies | Depth Required |
|------|--------------|----------------|
| Backend Development | **Go + Gin**, REST, gRPC | Deep |
| Databases | PostgreSQL, Redis, MongoDB | Deep |
| Event-Driven | Kafka, AWS SQS/SNS, Event Sourcing | Deep |
| Cloud Platform | AWS (Primary) | Expert |
| API Design | OpenAPI, gRPC, API Gateway patterns | Deep |

### Tier 2: MUST UNDERSTAND (Architectural decisions)

| Area | Technologies | Depth Required |
|------|--------------|----------------|
| Container Orchestration | Kubernetes, Docker | Conceptual + Basic Hands-on |
| Infrastructure as Code | Terraform, AWS CDK | Working Knowledge |
| CI/CD | GitHub Actions, ArgoCD | Working Knowledge |
| Observability | Prometheus, Grafana, OpenTelemetry | Working Knowledge |
| Security | OAuth2, OIDC, WAF, IAM | Conceptual + Design |

### Tier 3: AWARENESS (Know when to use)

| Area | Technologies |
|------|--------------|
| Service Mesh | Istio, Linkerd |
| ML/AI Integration | LLM APIs, RAG patterns |
| Data Engineering | Spark, Airflow, Data Lakes |
| Edge Computing | CloudFront, Lambda@Edge |

---

## 🎯 Proof Artifacts Rule (Non-Negotiable)

> **Every learning pillar MUST produce at least 1 tangible artifact that can be shown in an interview.**

Architects are not hired for what they *know* — they're hired for what they can *demonstrate*.

| Month | Pillar | Required Artifact |
|-------|--------|-------------------|
| 1 | Go Foundations | Working CLI tool + Concurrent file processor (GitHub) |
| 2 | Go Backend | Production-ready User Service with tests (GitHub) |
| 3 | Database Mastery | E-commerce schema + Query optimization report |
| 4 | Event-Driven | Order processing system with Kafka (GitHub + Architecture diagram) |
| 5 | Cloud & Infra | Terraform modules + Infrastructure diagram |
| 6 | Portfolio | 3 complete case studies with C4 diagrams |

### Artifact Quality Checklist:
- [ ] Can I explain this in 5 minutes to a non-technical stakeholder?
- [ ] Does it include trade-offs I made and why?
- [ ] Does it show failure handling?
- [ ] Is there a cost estimate?
- [ ] Is there a diagram?

---

## 🚀 Flagship Case Study: The Anchor Project

> **One deeply owned system that everything maps back to.**

Solution Architects are hired because: *"They deeply owned THIS one system."*

### Your Flagship: Scalable Mobile + Backend Platform for 500K Users

This is your interview anchor. Every learning topic connects here.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    FLAGSHIP: Mobile Commerce Platform                        │
│                    "500K users, real-time inventory, multi-region"          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  Your Experience          What You're Adding          Interview Story       │
│  ─────────────────        ─────────────────           ───────────────       │
│  React Native App    ──▶  Go Backend APIs        ──▶  "Full stack owner"   │
│  Node.js basics      ──▶  Event-driven order     ──▶  "Handled 10K TPS"    │
│  AWS exposure        ──▶  Terraform + K8s        ──▶  "Designed infra"     │
│  IoT background      ──▶  Real-time events       ──▶  "Scale expertise"    │
│                                                                              │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Flagship Components (Build Throughout Roadmap):

| Component | Month Built | Technologies | Key Decision to Defend |
|-----------|-------------|--------------|------------------------|
| User Service | 2 | Go/Gin, PostgreSQL, JWT | Why Go over Node.js? |
| Product Catalog | 3 | PostgreSQL, Redis cache | Cache invalidation strategy |
| Order Processing | 4 | Kafka, Saga pattern | Why event-driven over sync? |
| Inventory Service | 4 | Redis locks, optimistic locking | How to prevent overselling? |
| Infrastructure | 5 | Terraform, EKS, RDS | Why K8s over ECS? |
| Observability | 5 | OpenTelemetry, Grafana | SLO definitions |

### Flagship Artifacts to Create:
- [ ] C4 Context Diagram (system in ecosystem)
- [ ] C4 Container Diagram (all services)
- [ ] C4 Component Diagram (Order Service deep dive)
- [ ] 5 ADRs (key architectural decisions)
- [ ] Cost model ($X at 100K users, $Y at 500K users)
- [ ] Failure mode analysis (what breaks first?)
- [ ] Rollout plan (how to deploy safely)
- [ ] 3-minute elevator pitch
- [ ] 15-minute deep dive presentation

---

## 📢 Visibility & Networking (Weekly Actions)

> **Architect networks grow from: Sharing thinking, Reviewing designs, Helping others.**

Networking is not "intent" — it's **weekly output**.

### Weekly Networking Commitments:

| Day | Action | Platform | Time |
|-----|--------|----------|------|
| Monday | Publish 1 architecture insight/learning | LinkedIn/Twitter | 30 min |
| Wednesday | Comment thoughtfully on 3 architecture posts | LinkedIn | 20 min |
| Friday | Share 1 trade-off or failure story | LinkedIn/Blog | 30 min |
| Weekend | Review 1 open-source architecture or RFC | GitHub/Company blogs | 1 hr |

### Content Ideas (Mapped to Learning):

| Month | Learning Topic | Content to Share |
|-------|---------------|------------------|
| 1 | Go basics | "Coming from Node.js to Go: First impressions" |
| 2 | Clean Architecture | "How I structure Go services for testability" |
| 3 | Database design | "PostgreSQL indexing mistakes I made (and fixed)" |
| 4 | Event-driven | "Saga pattern: When to use and when to avoid" |
| 5 | Kubernetes | "My first Terraform module: lessons learned" |
| 6 | Case studies | "How I designed a 500K user platform" |

### Networking Metrics (Track Monthly):
- [ ] LinkedIn connections in tech/architecture: ___
- [ ] Posts published: ___
- [ ] Meaningful comments made: ___
- [ ] DMs/conversations started: ___
- [ ] Architecture discussions participated in: ___

### Where to Engage:
- **LinkedIn:** Architecture posts, SA job discussions
- **Twitter/X:** Tech Twitter, #SystemDesign, #SoftwareArchitecture
- **Dev.to/Medium:** Long-form architecture breakdowns
- **GitHub:** Comment on architectural RFCs, design docs
- **Local Meetups:** Present your flagship case study

---

## 🗓️ 6-Month Detailed Learning Plan

---

## MONTH 1: Go Language Foundations

### Week 1: Core Language Basics

**Why Go?** Cloud-native default language, excellent for high-performance backends, strong backend credibility.

#### Topics to Master:

**Variables & Types:**
- [ ] Variable declaration (var, :=, const)
- [ ] Basic types (int, int64, float64, string, bool, byte, rune)
- [ ] Type conversions (explicit only, no implicit)
- [ ] Zero values (every type has one - CRITICAL concept)

**Data Structures (YOU WILL USE DAILY):**
- [ ] Arrays (fixed size: `[5]int`)
- [ ] Slices (dynamic: `[]int`) - append, copy, slicing
- [ ] Maps (`map[string]int`) - create, read, delete, check existence
- [ ] Slice vs Array (know the difference!)
- [ ] Nil slices vs empty slices

**Control Flow:**
- [ ] if/else (with init statement)
- [ ] for loops (Go's only loop - 3 forms)
- [ ] switch (no break needed, case fallthrough)
- [ ] range (iterating slices, maps, strings)

**Functions:**
- [ ] Multiple return values
- [ ] Named return values
- [ ] Variadic functions (`...`)
- [ ] Functions as values (closures)
- [ ] Anonymous functions

**Pointers:**
- [ ] & (address of) and * (dereference)
- [ ] When to use pointers vs values
- [ ] Pointer receivers vs value receivers
- [ ] nil pointers

#### Quick Self-Test (Week 1):
```go
// Can you explain what each line does?
nums := []int{1, 2, 3}
nums = append(nums, 4)
m := make(map[string]int)
m["key"] = 1
if val, ok := m["key"]; ok { fmt.Println(val) }
```

---

### Week 2: Go Patterns & Idioms

#### Structs & Methods:
- [ ] Defining structs
- [ ] Struct literals (named vs positional)
- [ ] Embedded structs (composition)
- [ ] Methods (value receiver vs pointer receiver)
- [ ] Constructor patterns (NewXxx functions)

#### Interfaces:
- [ ] Implicit implementation (no `implements` keyword)
- [ ] Empty interface (`interface{}` / `any`)
- [ ] Type assertions (`value.(Type)`)
- [ ] Type switches
- [ ] Common interfaces (io.Reader, io.Writer, error, Stringer)

#### Error Handling:
- [ ] The `error` interface
- [ ] Creating errors (`errors.New`, `fmt.Errorf`)
- [ ] Error wrapping (`%w` and `errors.Is`, `errors.As`)
- [ ] Sentinel errors (`var ErrNotFound = errors.New(...)`)
- [ ] When to panic vs return error

#### Important Patterns:
- [ ] defer (cleanup, runs LIFO)
- [ ] panic and recover (rare, for truly exceptional cases)
- [ ] init() functions (runs before main)
- [ ] Packages and modules (go mod init, go mod tidy)
- [ ] Public vs private (capitalization)

#### Hands-on Project (Week 1-2):
```
Project 1: CLI Task Manager
- Add, list, complete, delete tasks
- Store in JSON file
- Use flag package for CLI args
- Proper error handling throughout
- Demonstrates: structs, slices, maps, file I/O, JSON

Project 2: Concurrent File Processor
- Read multiple files concurrently
- Count words/lines in each
- Aggregate results
- Demonstrates: goroutines, channels, sync.WaitGroup
```

#### Key Differences from JavaScript/TypeScript:
| Concept | TypeScript | Go |
|---------|------------|----|
| Error handling | try/catch | Multiple return values |
| Null safety | undefined/null | nil + explicit checks |
| OOP | Classes | Structs + Interfaces |
| Async | Promises/async-await | Goroutines + Channels |
| Generics | Built-in | Added in Go 1.18 |
| Zero values | undefined | Type-specific (0, "", nil) |
| Collections | Array, Object | Slice, Map |

#### Resources:
- **Official:** https://go.dev/tour (A Tour of Go) - DO THIS FIRST
- **Practice:** https://gobyexample.com (Concise examples)
- **Book:** "Learning Go" by Jon Bodner
- **Course:** "Go: The Complete Developer's Guide" - Udemy
- **Exercises:** https://exercism.io/tracks/go

---

### Week 3: Concurrency Deep Dive

#### Goroutines:
- [ ] Starting goroutines (`go func()`)
- [ ] Goroutine lifecycle
- [ ] Race conditions (and how to detect: `go run -race`)
- [ ] Goroutine leaks (and how to avoid)

#### Channels:
- [ ] Unbuffered channels (synchronous)
- [ ] Buffered channels (async up to capacity)
- [ ] Sending and receiving
- [ ] Closing channels
- [ ] Range over channels
- [ ] Directional channels (`chan<-`, `<-chan`)

#### Select Statement:
- [ ] Multiplexing channels
- [ ] Default case (non-blocking)
- [ ] Timeouts with `time.After`

#### Sync Package:
- [ ] sync.Mutex (protect shared state)
- [ ] sync.RWMutex (multiple readers, single writer)
- [ ] sync.WaitGroup (wait for goroutines)
- [ ] sync.Once (one-time initialization)

#### Context Package (CRITICAL FOR BACKEND):
- [ ] context.Background() and context.TODO()
- [ ] context.WithCancel
- [ ] context.WithTimeout
- [ ] context.WithDeadline
- [ ] context.WithValue (use sparingly)
- [ ] Passing context through function calls

#### Common Concurrency Patterns:
- [ ] Worker pools
- [ ] Fan-out, fan-in
- [ ] Pipeline
- [ ] Semaphore with buffered channel

#### Hands-on Project:
```
Project 3: Concurrent URL Checker
- Check multiple URLs concurrently
- Timeout after 5 seconds per URL
- Cancel all requests if user interrupts (Ctrl+C)
- Report results as they complete
- Demonstrates: goroutines, channels, context, select
```

---

### Week 4: Go Intermediate + Gin Intro

#### Standard Library Essentials:
- [ ] `io` (readers, writers, utilities)
- [ ] `os` (files, env vars, args, ReadFile/WriteFile)
- [ ] `strings` and `strconv`
- [ ] `time` (durations, parsing, formatting)
- [ ] `encoding/json` (Marshal, Unmarshal, struct tags)
- [ ] `net/http` (basic server, client)
- [ ] `log` and `log/slog` (structured logging)

#### Testing:
- [ ] Writing tests (`_test.go` files)
- [ ] Table-driven tests
- [ ] Subtests (`t.Run`)
- [ ] Test helpers
- [ ] Mocking (interfaces for testability)
- [ ] Benchmarks (`func BenchmarkXxx`)
- [ ] `go test -cover` for coverage

#### Generics (Go 1.18+):
- [ ] Type parameters
- [ ] Type constraints
- [ ] Common use cases (slices, maps utilities)

#### Project Structure:
- [ ] Standard layout (`cmd/`, `internal/`, `pkg/`)
- [ ] When to use `internal/`
- [ ] Organizing by feature vs layer

#### Gin Framework Basics:
> **Note on Gin:** You are using Gin because it is the most popular Go web framework, uses the standard `net/http` library (unlike Fiber), and has a massive ecosystem of middleware. It strikes the perfect balance between performance and productivity.

- [ ] Routing and handlers (`gin.Context`)
- [ ] Middleware (logging, recovery, CORS)
- [ ] Request parsing (`ShouldBindJSON`, `ShouldBindQuery`)
- [ ] Response handling (`c.JSON`)
- [ ] Validation (built-in `go-playground/validator`)
- [ ] Configuration management (Viper)

#### Hands-on Project:
```
Build: User Management API (Gin)
- POST /users (create with validation)
- GET /users/:id
- PUT /users/:id
- DELETE /users/:id
- Proper error handling
- Request logging middleware
- Input validation
- Structured logging (zerolog/zap)
```

#### Project Structure:
```
user-service/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── config/              # Configuration
│   ├── handler/             # HTTP handlers
│   ├── middleware/          # Gin middleware
│   ├── model/               # Domain models
│   ├── repository/          # Database access
│   └── service/             # Business logic
├── pkg/                     # Reusable packages
├── go.mod
├── go.sum
└── Makefile
```

#### Resources:
- **Gin Docs:** https://gin-gonic.com/docs/
- **Project Structure:** https://github.com/golang-standards/project-layout

---

## MONTH 2: Go Backend Deep Dive

### Week 1: Authentication & Security

#### JWT Authentication:
- [ ] JWT structure (header, payload, signature)
- [ ] Access tokens (short-lived, 15min-1hr)
- [ ] Refresh tokens (long-lived, stored securely)
- [ ] Refresh token rotation (security best practice)
- [ ] Token revocation strategies (blacklist, token versioning)
- [ ] `golang-jwt/jwt` library usage
- [ ] Claims (standard + custom)
- [ ] Middleware for token validation

#### Password Security:
- [ ] bcrypt hashing (cost factor selection)
- [ ] Password validation rules
- [ ] Timing attack prevention
- [ ] Password reset flow

#### Authorization:
- [ ] RBAC (Role-Based Access Control)
- [ ] Permission-based guards
- [ ] Middleware composition
- [ ] API key authentication
- [ ] Multi-tenancy considerations

#### Security Essentials:
- [ ] Input validation (go-playground/validator)
- [ ] SQL injection prevention (parameterized queries ALWAYS)
- [ ] XSS prevention (output encoding)
- [ ] CORS configuration
- [ ] Rate limiting (`gin-contrib/rate-limiter`)
- [ ] Request ID / correlation ID
- [ ] Secure headers

#### Code Example - Complete JWT Middleware:
```go
func JWTMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract token from header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(401, gin.H{
                "error": "Missing authorization header",
            })
            return
        }
        
        // Parse "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(401, gin.H{
                "error": "Invalid authorization format",
            })
            return
        }
        
        // Validate token
        claims, err := ValidateToken(parts[1], secret)
        if err != nil {
            c.AbortWithStatusJSON(401, gin.H{
                "error": "Invalid or expired token",
            })
            return
        }
        
        // Check if token is blacklisted (for logout support)
        if IsTokenBlacklisted(parts[1]) {
            c.AbortWithStatusJSON(401, gin.H{
                "error": "Token has been revoked",
            })
            return
        }
        
        // Store user info in context
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}

// Role-based authorization middleware
func RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole := c.GetString("role")
        for _, role := range roles {
            if userRole == role {
                c.Next()
                return
            }
        }
        c.AbortWithStatusJSON(403, gin.H{
            "error": "Insufficient permissions",
        })
    }
}
```

#### Hands-on Project:
```
Build: Complete Auth System
- POST /auth/register (with validation)
- POST /auth/login (returns access + refresh tokens)
- POST /auth/refresh (rotate refresh token)
- POST /auth/logout (blacklist token)
- GET /auth/me (protected route)
- Password reset flow (email token)
```

---

### Week 2: Database Integration

> **Infrastructure as Code (IaC) Tip:** Instead of creating your RDS/Postgres instance in the AWS Console, try to use Terraform now (see Month 5). This builds IaC muscle memory early.

#### Database Fundamentals in Go:
- [ ] `database/sql` package (standard library)
- [ ] Connection string management (env vars, secrets)
- [ ] Connection pooling (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`)
- [ ] Context with ALL queries (timeouts, cancellation)
- [ ] Prepared statements (performance + security)
- [ ] `lib/pq` (PostgreSQL driver)

#### Query Patterns:
- [ ] `QueryRow` vs `Query` vs `Exec` (know the difference!)
- [ ] Scanning into structs
- [ ] Handling NULL values (`sql.NullString`, `sql.NullInt64`, etc.)
- [ ] SQLX for easier struct scanning
- [ ] When to use ORM vs raw SQL

#### Code Example - Proper Database Usage:
```go
// Repository with proper patterns
type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Always use context for timeout/cancellation
func (r *UserRepository) GetByID(ctx context.Context, id string) (*User, error) {
    query := `
        SELECT id, email, name, role, created_at 
        FROM users 
        WHERE id = $1 AND deleted_at IS NULL
    `
    
    var user User
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID,
        &user.Email,
        &user.Name,
        &user.Role,
        &user.CreatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, ErrUserNotFound
    }
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return &user, nil
}

// Transaction example
func (r *UserRepository) CreateWithProfile(ctx context.Context, user *User, profile *Profile) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback() // Rollback if not committed
    
    // Insert user
    _, err = tx.ExecContext(ctx, 
        `INSERT INTO users (id, email, name) VALUES ($1, $2, $3)`,
        user.ID, user.Email, user.Name,
    )
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }
    
    // Insert profile
    _, err = tx.ExecContext(ctx,
        `INSERT INTO profiles (user_id, bio) VALUES ($1, $2)`,
        user.ID, profile.Bio,
    )
    if err != nil {
        return fmt.Errorf("failed to insert profile: %w", err)
    }
    
    return tx.Commit()
}
```

#### GORM (ORM option):
- [ ] Model definition with struct tags
- [ ] CRUD operations
- [ ] Relationships (HasOne, HasMany, BelongsTo)
- [ ] Preloading (avoiding N+1 problem)
- [ ] Hooks (BeforeCreate, AfterUpdate)
- [ ] Raw SQL when needed
- [ ] When NOT to use ORM (complex queries, performance)

#### Migrations:
- [ ] `golang-migrate` setup
- [ ] Up/Down migrations
- [ ] Migration versioning best practices
- [ ] CI/CD integration for migrations
- [ ] Rollback strategies

#### Error Handling:
- [ ] Wrapping database errors with context
- [ ] Custom error types (`ErrNotFound`, `ErrConflict`, `ErrDuplicate`)
- [ ] Retry logic for transient failures
- [ ] Query logging in debug mode

---

### Week 3: Architecture Patterns (CRITICAL FOR SA)

#### Clean Architecture:
- [ ] Layers: Handler → Service → Repository
- [ ] Dependency direction (always inward)
- [ ] Interface-based design
- [ ] Why this matters for testing and maintainability

#### Repository Pattern:
- [ ] Interface definition
- [ ] Concrete implementation
- [ ] Why interfaces enable mocking
- [ ] Unit of Work pattern (optional)

#### Service Layer:
- [ ] Business logic isolation
- [ ] Transaction boundaries
- [ ] Orchestrating multiple repositories
- [ ] Cross-cutting concerns

#### Dependency Injection:
- [ ] Constructor injection (preferred in Go)
- [ ] Interface-based DI
- [ ] Manual DI vs Wire (Google's DI tool)
- [ ] Wiring in main.go

#### Code Example - Clean Architecture:
```go
// 1. Define interfaces (in service or domain package)
type UserRepository interface {
    GetByID(ctx context.Context, id string) (*User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
}

type UserService interface {
    GetUser(ctx context.Context, id string) (*UserDTO, error)
    CreateUser(ctx context.Context, req CreateUserRequest) (*UserDTO, error)
}

// 2. Implement repository
type postgresUserRepo struct {
    db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) UserRepository {
    return &postgresUserRepo{db: db}
}

func (r *postgresUserRepo) GetByID(ctx context.Context, id string) (*User, error) {
    // Implementation
}

// 3. Implement service with injected repository
type userService struct {
    repo   UserRepository
    cache  CacheService
    logger *slog.Logger
}

func NewUserService(repo UserRepository, cache CacheService, logger *slog.Logger) UserService {
    return &userService{repo: repo, cache: cache, logger: logger}
}

func (s *userService) GetUser(ctx context.Context, id string) (*UserDTO, error) {
    // Check cache first
    if cached, err := s.cache.Get(ctx, "user:"+id); err == nil {
        return cached.(*UserDTO), nil
    }
    
    // Get from database
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    dto := mapUserToDTO(user)
    
    // Cache for next time
    s.cache.Set(ctx, "user:"+id, dto, 5*time.Minute)
    
    return dto, nil
}

// 4. Wire everything in main.go
func main() {
    // Setup
    db := setupDatabase()
    cache := setupRedis()
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    
    // Dependency injection
    userRepo := NewPostgresUserRepo(db)
    userService := NewUserService(userRepo, cache, logger)
    userHandler := NewUserHandler(userService)
    
    // Setup routes
    r := gin.Default()
    r.GET("/users/:id", userHandler.GetUser)
    
    r.Run(":3000")
}
```

#### Project Structure (Production-Ready):
```
my-service/
├── cmd/
│   └── api/
│       └── main.go              # Entry point, DI wiring
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration loading
│   ├── handler/                  # HTTP handlers (controllers)
│   │   ├── user_handler.go
│   │   └── auth_handler.go
│   ├── middleware/               # HTTP middleware
│   │   ├── auth.go
│   │   ├── logging.go
│   │   └── ratelimit.go
│   ├── service/                  # Business logic
│   │   ├── user_service.go
│   │   └── auth_service.go
│   ├── repository/               # Database access
│   │   ├── interfaces.go         # Repository interfaces
│   │   ├── user_repository.go
│   │   └── postgres/             # PostgreSQL implementations
│   ├── model/                    # Domain models
│   │   └── user.go
│   ├── dto/                      # Request/Response objects
│   │   ├── request.go
│   │   └── response.go
│   └── pkg/                      # Internal shared packages
│       ├── validator/
│       └── errors/
├── pkg/                          # Public reusable packages
├── migrations/
│   ├── 000001_create_users.up.sql
│   └── 000001_create_users.down.sql
├── tests/
│   ├── integration/
│   └── e2e/
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
└── go.sum
```

---

### Week 4: Testing & Production Readiness

#### Unit Testing:
- [ ] Testing service layer (mock repository)
- [ ] Table-driven tests
- [ ] Test helpers and fixtures
- [ ] Mocking with interfaces
- [ ] `testify` library (assert, mock, suite)

#### Integration Testing:
- [ ] Testing with real database
- [ ] `testcontainers-go` for PostgreSQL
- [ ] Test database setup/teardown
- [ ] Parallel test execution

#### Code Example - Testing with Mocks:
```go
// Mock repository for testing
type mockUserRepo struct {
    mock.Mock
}

func (m *mockUserRepo) GetByID(ctx context.Context, id string) (*User, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

// Test the service
func TestUserService_GetUser(t *testing.T) {
    tests := []struct {
        name        string
        userID      string
        mockUser    *User
        mockError   error
        expected    *UserDTO
        expectedErr error
    }{
        {
            name:   "success",
            userID: "123",
            mockUser: &User{
                ID:    "123",
                Email: "test@example.com",
                Name:  "Test User",
            },
            expected: &UserDTO{
                ID:    "123",
                Email: "test@example.com",
                Name:  "Test User",
            },
        },
        {
            name:        "user not found",
            userID:      "999",
            mockError:   ErrUserNotFound,
            expectedErr: ErrUserNotFound,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup mock
            mockRepo := new(mockUserRepo)
            mockRepo.On("GetByID", mock.Anything, tt.userID).
                Return(tt.mockUser, tt.mockError)
            
            // Create service with mock
            svc := NewUserService(mockRepo, nil, slog.Default())
            
            // Execute
            result, err := svc.GetUser(context.Background(), tt.userID)
            
            // Assert
            if tt.expectedErr != nil {
                assert.ErrorIs(t, err, tt.expectedErr)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expected, result)
            }
            
            mockRepo.AssertExpectations(t)
        })
    }
}
```

#### Configuration Management:
- [ ] Viper for config loading
- [ ] Environment-based configuration
- [ ] Config validation at startup
- [ ] Secrets management (env vars, not files)
- [ ] Feature flags

#### Logging & Observability:
- [ ] Structured logging (`slog`, `zerolog`, or `zap`)
- [ ] Log levels (debug, info, warn, error)
- [ ] Request logging middleware
- [ ] Correlation IDs across requests
- [ ] Basic metrics preparation

#### Error Handling (API Responses):
- [ ] Consistent error response format
- [ ] Error codes for client handling
- [ ] Validation error details
- [ ] Never leak internal errors to clients

#### Code Example - Error Response Pattern:
```go
// Standard error response
type ErrorResponse struct {
    Error   string            `json:"error"`
    Code    string            `json:"code,omitempty"`
    Details map[string]string `json:"details,omitempty"`
}

// Custom error types
var (
    ErrNotFound     = errors.New("resource not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrForbidden    = errors.New("forbidden")
    ErrConflict     = errors.New("resource already exists")
)

// Error handler middleware
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // Execute handlers first

        // Check if any errors were attached to the context
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            code := http.StatusInternalServerError
            response := ErrorResponse{Error: "Internal server error"}
            
            switch {
            case errors.Is(err, ErrNotFound):
                code = http.StatusNotFound
                response = ErrorResponse{Error: "Resource not found", Code: "NOT_FOUND"}
            case errors.Is(err, ErrUnauthorized):
                code = http.StatusUnauthorized
                response = ErrorResponse{Error: "Unauthorized", Code: "UNAUTHORIZED"}
            case errors.Is(err, ErrForbidden):
                code = http.StatusForbidden
                response = ErrorResponse{Error: "Forbidden", Code: "FORBIDDEN"}
            case errors.Is(err, ErrConflict):
                code = http.StatusConflict
                response = ErrorResponse{Error: "Resource already exists", Code: "CONFLICT"}
            }
            
            c.JSON(code, response)
        }
    }
}
```

#### Hands-on Project (Month 2 Capstone):
```
Build: Production-Ready User Service
- Complete CRUD with PostgreSQL
- JWT auth with refresh token rotation
- Role-based access control
- Clean architecture (handler → service → repository)
- Database migrations
- Unit tests with mocks
- Integration tests with testcontainers
- Structured logging
- Consistent error handling
- Docker + docker-compose
- Makefile for common tasks
```

#### Architecture Artifact:
Create an **Authentication Architecture Document**:
- Flow diagrams for login, refresh, logout
- Token storage decisions (client-side)
- Security considerations
- Rate limiting strategy

---

## MONTH 3: Database Mastery

> **Reality Check:** This month covers deep topics (internals, locking, isolation) that often take years to master. Don't get discouraged if this takes 6-8 weeks. It is better to understand *one* database deeply than three superficially.

### Week 1: PostgreSQL Deep Dive - Schema Design

#### Database Design Fundamentals:
- [ ] Normalization (1NF → 2NF → 3NF → BCNF)
- [ ] When to denormalize (read-heavy workloads)
- [ ] Primary keys (UUID vs serial vs ULID)
- [ ] Foreign keys and referential integrity
- [ ] Constraints (NOT NULL, UNIQUE, CHECK)
- [ ] Data types (choosing the right type matters!)
- [ ] Naming conventions (snake_case, consistent pluralization)

#### Schema Design Patterns:
- [ ] Soft deletes (`deleted_at` column)
- [ ] Audit columns (`created_at`, `updated_at`, `created_by`)
- [ ] Versioning for optimistic locking
- [ ] Multi-tenancy strategies (schema per tenant, row-level)
- [ ] Polymorphic associations
- [ ] Self-referential relationships (trees, hierarchies)

#### Hands-on Exercise - E-Commerce Schema:
```sql
-- Users with proper constraints
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'customer',
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    CONSTRAINT valid_role CHECK (role IN ('customer', 'admin', 'vendor'))
);

-- Partial index for active users only
CREATE INDEX idx_users_email_active ON users(email) WHERE deleted_at IS NULL;

-- Products with full-text search
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vendor_id UUID NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    search_vector TSVECTOR,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT positive_price CHECK (price > 0),
    CONSTRAINT valid_status CHECK (status IN ('draft', 'active', 'archived'))
);

-- Full-text search index
CREATE INDEX idx_products_search ON products USING GIN(search_vector);

-- Trigger to update search vector
CREATE OR REPLACE FUNCTION update_product_search_vector()
RETURNS TRIGGER AS $$
BEGIN
    NEW.search_vector := to_tsvector('english', COALESCE(NEW.name, '') || ' ' || COALESCE(NEW.description, ''));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_search_vector
    BEFORE INSERT OR UPDATE ON products
    FOR EACH ROW EXECUTE FUNCTION update_product_search_vector();

-- Orders partitioned by date (for scale)
CREATE TABLE orders (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    shipping_address JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    PRIMARY KEY (id, created_at),
    CONSTRAINT valid_order_status CHECK (status IN ('pending', 'confirmed', 'shipped', 'delivered', 'cancelled'))
) PARTITION BY RANGE (created_at);

-- Create partitions
CREATE TABLE orders_2024_q1 PARTITION OF orders
    FOR VALUES FROM ('2024-01-01') TO ('2024-04-01');
CREATE TABLE orders_2024_q2 PARTITION OF orders
    FOR VALUES FROM ('2024-04-01') TO ('2024-07-01');

-- Inventory with optimistic locking
CREATE TABLE inventory (
    product_id UUID PRIMARY KEY REFERENCES products(id),
    available_quantity INTEGER NOT NULL DEFAULT 0,
    reserved_quantity INTEGER NOT NULL DEFAULT 0,
    version INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT non_negative_available CHECK (available_quantity >= 0),
    CONSTRAINT non_negative_reserved CHECK (reserved_quantity >= 0)
);
```

---

### Week 2: PostgreSQL Deep Dive - Performance

#### Indexing Strategies:
- [ ] B-tree indexes (default, for equality and range)
- [ ] GIN indexes (for arrays, JSONB, full-text search)
- [ ] GiST indexes (for geometric, full-text)
- [ ] BRIN indexes (for large sorted tables)
- [ ] Partial indexes (index only relevant rows)
- [ ] Expression indexes (index on function results)
- [ ] Covering indexes (INCLUDE columns)
- [ ] Multi-column indexes (column order matters!)

#### Index Strategy Guide:
```sql
-- B-tree: Equality and range queries
CREATE INDEX idx_orders_user ON orders(user_id);
CREATE INDEX idx_orders_created ON orders(created_at DESC);

-- Partial: Only index what you query
CREATE INDEX idx_orders_pending ON orders(user_id) 
    WHERE status = 'pending';

-- Covering: Include columns to avoid table lookup
CREATE INDEX idx_orders_user_covering ON orders(user_id) 
    INCLUDE (status, total_amount);

-- Expression: Index computed values
CREATE INDEX idx_users_email_lower ON users(LOWER(email));

-- GIN for JSONB
CREATE INDEX idx_orders_shipping ON orders USING GIN(shipping_address);
```

#### Query Optimization:
- [ ] EXPLAIN vs EXPLAIN ANALYZE
- [ ] Reading execution plans (Seq Scan, Index Scan, Bitmap Scan)
- [ ] Identifying slow queries
- [ ] Query statistics (`pg_stat_statements`)
- [ ] Vacuuming and analyze
- [ ] Cost estimation

#### EXPLAIN ANALYZE Practice:
```sql
-- Bad: Full table scan
EXPLAIN ANALYZE
SELECT * FROM orders WHERE status = 'pending';

-- Good: Uses partial index
EXPLAIN ANALYZE
SELECT * FROM orders 
WHERE status = 'pending' AND user_id = '123e4567-e89b-12d3-a456-426614174000';

-- Identify N+1 queries in your Go code!
```

#### Transactions & Isolation:
- [ ] ACID properties
- [ ] Isolation levels (Read Committed, Repeatable Read, Serializable)
- [ ] Deadlock detection and prevention
- [ ] Row-level locking (`SELECT ... FOR UPDATE`)
- [ ] Advisory locks for application-level locking

#### Code Example - Proper Transaction in Go:
```go
func (r *OrderRepository) CreateOrder(ctx context.Context, order *Order, items []OrderItem) error {
    tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
        Isolation: sql.LevelRepeatableRead,
    })
    if err != nil {
        return fmt.Errorf("begin transaction: %w", err)
    }
    defer tx.Rollback()
    
    // 1. Insert order
    _, err = tx.ExecContext(ctx, `
        INSERT INTO orders (id, user_id, status, total_amount)
        VALUES ($1, $2, $3, $4)
    `, order.ID, order.UserID, order.Status, order.TotalAmount)
    if err != nil {
        return fmt.Errorf("insert order: %w", err)
    }
    
    // 2. Reserve inventory with locking
    for _, item := range items {
        result, err := tx.ExecContext(ctx, `
            UPDATE inventory 
            SET available_quantity = available_quantity - $1,
                reserved_quantity = reserved_quantity + $1,
                version = version + 1
            WHERE product_id = $2 
              AND available_quantity >= $1
        `, item.Quantity, item.ProductID)
        
        if err != nil {
            return fmt.Errorf("update inventory: %w", err)
        }
        
        rows, _ := result.RowsAffected()
        if rows == 0 {
            return ErrInsufficientInventory
        }
    }
    
    // 3. Insert order items
    for _, item := range items {
        _, err = tx.ExecContext(ctx, `
            INSERT INTO order_items (order_id, product_id, quantity, unit_price)
            VALUES ($1, $2, $3, $4)
        `, order.ID, item.ProductID, item.Quantity, item.UnitPrice)
        if err != nil {
            return fmt.Errorf("insert order item: %w", err)
        }
    }
    
    return tx.Commit()
}
```

#### Scaling PostgreSQL:
- [ ] Connection pooling with PgBouncer
- [ ] Read replicas (streaming replication)
- [ ] Partitioning strategies (range, list, hash)
- [ ] When to partition (>10M rows, time-series)
- [ ] Archiving old data
- [ ] pg_dump/pg_restore for backups

#### Connection Pooling Configuration:
```go
func SetupDatabase(cfg *Config) (*sql.DB, error) {
    db, err := sql.Open("postgres", cfg.DatabaseURL)
    if err != nil {
        return nil, err
    }
    
    // Connection pool settings (CRITICAL for production)
    db.SetMaxOpenConns(25)                  // Max connections to DB
    db.SetMaxIdleConns(5)                   // Keep some connections ready
    db.SetConnMaxLifetime(5 * time.Minute)  // Recycle connections
    db.SetConnMaxIdleTime(1 * time.Minute)  // Close idle connections
    
    // Verify connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := db.PingContext(ctx); err != nil {
        return nil, fmt.Errorf("database ping failed: %w", err)
    }
    
    return db, nil
}
```

#### Architecture Artifact:
Create **Database Design Document**:
- ER diagram
- Table schemas with rationale
- Index strategy
- Partitioning plan
- Backup/recovery plan
- Scaling strategy

---

### Week 3-4: Redis & Caching Strategies

#### Redis Data Structures:
- [ ] Strings (simplest - key/value, counters)
- [ ] Hashes (objects, user sessions)
- [ ] Lists (queues, activity feeds)
- [ ] Sets (unique collections, tags)
- [ ] Sorted Sets (leaderboards, rate limiting)
- [ ] Streams (event log, message queue)
- [ ] HyperLogLog (approximate counting)

#### When to Use Each:
| Data Structure | Use Case | Example |
|----------------|----------|---------|
| String | Simple cache, counters | `GET user:123:profile` |
| Hash | Object storage | `HGETALL session:abc123` |
| List | Queues, recent items | `LPUSH notifications:user:123` |
| Set | Unique items, tags | `SADD product:123:tags "electronics"` |
| Sorted Set | Leaderboards, rate limiting | `ZADD leaderboard 1000 "user:123"` |
| Stream | Event log | `XADD orders * event "created"` |

#### Caching Patterns:

**1. Cache-Aside (Lazy Loading):**
```go
func (s *ProductService) GetProduct(ctx context.Context, id string) (*Product, error) {
    // 1. Try cache first
    cacheKey := fmt.Sprintf("product:%s", id)
    cached, err := s.redis.Get(ctx, cacheKey).Result()
    if err == nil {
        var product Product
        json.Unmarshal([]byte(cached), &product)
        return &product, nil
    }
    
    // 2. Cache miss - get from database
    product, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // 3. Populate cache for next time
    data, _ := json.Marshal(product)
    s.redis.Set(ctx, cacheKey, data, 15*time.Minute)
    
    return product, nil
}
```

**2. Write-Through:**
```go
func (s *ProductService) UpdateProduct(ctx context.Context, product *Product) error {
    // 1. Update database
    if err := s.repo.Update(ctx, product); err != nil {
        return err
    }
    
    // 2. Update cache immediately
    cacheKey := fmt.Sprintf("product:%s", product.ID)
    data, _ := json.Marshal(product)
    s.redis.Set(ctx, cacheKey, data, 15*time.Minute)
    
    return nil
}
```

**3. Write-Behind (Async):**
```go
func (s *ProductService) UpdateProductAsync(ctx context.Context, product *Product) error {
    // 1. Update cache immediately (fast response)
    cacheKey := fmt.Sprintf("product:%s", product.ID)
    data, _ := json.Marshal(product)
    s.redis.Set(ctx, cacheKey, data, 15*time.Minute)
    
    // 2. Queue database update for async processing
    s.queue.Publish(ctx, "product:updates", data)
    
    return nil
}
```

#### Cache Invalidation Strategies:
- [ ] TTL-based expiration (simplest)
- [ ] Event-based invalidation (publish on change)
- [ ] Version-based (cache key includes version)
- [ ] Tag-based invalidation (group related keys)

#### Code Example - Event-Based Invalidation:
```go
// When product is updated
func (s *ProductService) UpdateProduct(ctx context.Context, product *Product) error {
    if err := s.repo.Update(ctx, product); err != nil {
        return err
    }
    
    // Publish invalidation event
    s.redis.Publish(ctx, "cache:invalidate", fmt.Sprintf("product:%s", product.ID))
    
    // Also invalidate related caches
    s.redis.Del(ctx, 
        fmt.Sprintf("product:%s", product.ID),
        fmt.Sprintf("category:%s:products", product.CategoryID),
        "products:featured",
    )
    
    return nil
}
```

#### Cache Failure Patterns:

**1. Cache Stampede Prevention:**
```go
func (s *ProductService) GetProductWithLock(ctx context.Context, id string) (*Product, error) {
    cacheKey := fmt.Sprintf("product:%s", id)
    lockKey := fmt.Sprintf("lock:product:%s", id)
    
    // Try cache
    cached, err := s.redis.Get(ctx, cacheKey).Result()
    if err == nil {
        var product Product
        json.Unmarshal([]byte(cached), &product)
        return &product, nil
    }
    
    // Try to acquire lock
    acquired, err := s.redis.SetNX(ctx, lockKey, "1", 10*time.Second).Result()
    if err != nil || !acquired {
        // Another request is fetching, wait and retry
        time.Sleep(100 * time.Millisecond)
        return s.GetProductWithLock(ctx, id)
    }
    defer s.redis.Del(ctx, lockKey)
    
    // We got the lock - fetch from DB
    product, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Cache it
    data, _ := json.Marshal(product)
    s.redis.Set(ctx, cacheKey, data, 15*time.Minute)
    
    return product, nil
}
```

**2. Circuit Breaker for Cache:**
```go
// If Redis is down, don't keep hammering it
type CacheWithCircuitBreaker struct {
    redis    *redis.Client
    failures int
    lastFail time.Time
    mu       sync.Mutex
}

func (c *CacheWithCircuitBreaker) Get(ctx context.Context, key string) (string, error) {
    c.mu.Lock()
    if c.failures >= 5 && time.Since(c.lastFail) < 30*time.Second {
        c.mu.Unlock()
        return "", ErrCircuitOpen // Skip Redis entirely
    }
    c.mu.Unlock()
    
    val, err := c.redis.Get(ctx, key).Result()
    if err != nil && err != redis.Nil {
        c.mu.Lock()
        c.failures++
        c.lastFail = time.Now()
        c.mu.Unlock()
    } else {
        c.mu.Lock()
        c.failures = 0
        c.mu.Unlock()
    }
    
    return val, err
}
```

#### Distributed Locking:
```go
func (s *InventoryService) ReserveWithLock(ctx context.Context, productID string, quantity int) error {
    lockKey := fmt.Sprintf("lock:inventory:%s", productID)
    lockValue := uuid.New().String()
    
    // Acquire lock with expiration
    acquired, err := s.redis.SetNX(ctx, lockKey, lockValue, 10*time.Second).Result()
    if err != nil {
        return fmt.Errorf("failed to acquire lock: %w", err)
    }
    if !acquired {
        return ErrLockNotAcquired
    }
    
    // Ensure we release the lock (only if we own it)
    defer func() {
        // Lua script for atomic check-and-delete
        script := `
            if redis.call("get", KEYS[1]) == ARGV[1] then
                return redis.call("del", KEYS[1])
            else
                return 0
            end
        `
        s.redis.Eval(ctx, script, []string{lockKey}, lockValue)
    }()
    
    // Do the actual reservation
    return s.repo.Reserve(ctx, productID, quantity)
}
```

#### Rate Limiting with Redis:
```go
func RateLimitMiddleware(redis *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := fmt.Sprintf("ratelimit:%s", c.ClientIP())
        
        // Increment counter
        count, err := redis.Incr(c.Request.Context(), key).Result()
        if err != nil {
            c.Next() // Fail open
            return
        }
        
        // Set expiry on first request
        if count == 1 {
            redis.Expire(c.Request.Context(), key, window)
        }
        
        // Check limit
        if count > int64(limit) {
            c.AbortWithStatusJSON(429, gin.H{
                "error": "Rate limit exceeded",
                "retry_after": redis.TTL(c.Request.Context(), key).Val().Seconds(),
            })
            return
        }
        
        c.Next()
    }
}
```

#### Hands-on Project:
```
Implement Caching Layer:
1. Product catalog cache (cache-aside)
2. User session storage (Hash)
3. Rate limiting (Sorted Set with sliding window)
4. Leaderboard (Sorted Set)
5. Distributed locks for inventory
6. Cache invalidation on product updates
7. Circuit breaker for Redis failures
```

#### Architecture Artifact:
Create **Caching Strategy Document**:
- What to cache (hot data identification)
- What NOT to cache (frequently changing, large objects)
- TTL decisions by data type
- Invalidation strategy
- Failure handling (what if Redis is down?)
- Memory estimation and sizing

---

## MONTH 4: Event-Driven Architecture

### Week 1-2: Kafka Fundamentals

#### Kafka Architecture:
- [ ] Brokers (servers in the cluster)
- [ ] Topics (categories of messages)
- [ ] Partitions (parallelism unit)
- [ ] Replication (fault tolerance)
- [ ] ZooKeeper/KRaft (cluster coordination)
- [ ] Log compaction vs retention

#### Core Concepts:
```
┌─────────────────────────────────────────────────────────────┐
│                        Kafka Cluster                         │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  Topic: orders (3 partitions, replication factor 2)         │
│                                                              │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐                   │
│  │Partition 0│  │Partition 1│  │Partition 2│                 │
│  │ [0,1,2,3]│  │ [0,1,2]  │  │ [0,1,2,3,4]│                 │
│  │ Leader:B1│  │ Leader:B2│  │ Leader:B3│                   │
│  │Replica:B2│  │Replica:B3│  │Replica:B1│                   │
│  └──────────┘  └──────────┘  └──────────┘                   │
│                                                              │
└─────────────────────────────────────────────────────────────┘

Producer ──▶ Partition (based on key hash or round-robin)

Consumer Group: order-processors
├── Consumer 1 ──▶ Partition 0
├── Consumer 2 ──▶ Partition 1
└── Consumer 3 ──▶ Partition 2
```

#### Producers:
- [ ] Sync vs async sends
- [ ] Acks configuration (0, 1, all)
- [ ] Retries and idempotency
- [ ] Partitioning strategies
- [ ] Batching for throughput
- [ ] Compression

#### Code Example - Go Kafka Producer:
```go
import "github.com/segmentio/kafka-go"

type OrderEventPublisher struct {
    writer *kafka.Writer
}

func NewOrderEventPublisher(brokers []string) *OrderEventPublisher {
    return &OrderEventPublisher{
        writer: &kafka.Writer{
            Addr:         kafka.TCP(brokers...),
            Topic:        "orders",
            Balancer:     &kafka.Hash{}, // Partition by key
            RequiredAcks: kafka.RequireAll,
            Async:        false, // Sync for reliability
            Compression:  kafka.Snappy,
        },
    }
}

type OrderCreatedEvent struct {
    EventID   string    `json:"event_id"`
    EventType string    `json:"event_type"`
    Timestamp time.Time `json:"timestamp"`
    OrderID   string    `json:"order_id"`
    UserID    string    `json:"user_id"`
    Items     []Item    `json:"items"`
    Total     float64   `json:"total"`
}

func (p *OrderEventPublisher) PublishOrderCreated(ctx context.Context, event OrderCreatedEvent) error {
    event.EventID = uuid.New().String()
    event.EventType = "OrderCreated"
    event.Timestamp = time.Now()
    
    data, err := json.Marshal(event)
    if err != nil {
        return fmt.Errorf("marshal event: %w", err)
    }
    
    err = p.writer.WriteMessages(ctx, kafka.Message{
        Key:   []byte(event.OrderID), // Same order always goes to same partition
        Value: data,
        Headers: []kafka.Header{
            {Key: "event_type", Value: []byte(event.EventType)},
            {Key: "correlation_id", Value: []byte(getCorrelationID(ctx))},
        },
    })
    
    if err != nil {
        return fmt.Errorf("publish event: %w", err)
    }
    
    return nil
}
```

#### Consumers:
- [ ] Consumer groups
- [ ] Partition assignment
- [ ] Offset management (auto vs manual commit)
- [ ] Rebalancing
- [ ] At-least-once vs at-most-once vs exactly-once

#### Code Example - Go Kafka Consumer:
```go
type OrderEventConsumer struct {
    reader  *kafka.Reader
    handler OrderEventHandler
}

func NewOrderEventConsumer(brokers []string, groupID string, handler OrderEventHandler) *OrderEventConsumer {
    return &OrderEventConsumer{
        reader: kafka.NewReader(kafka.ReaderConfig{
            Brokers:        brokers,
            Topic:          "orders",
            GroupID:        groupID,
            MinBytes:       1,
            MaxBytes:       10e6, // 10MB
            CommitInterval: time.Second,
            StartOffset:    kafka.FirstOffset,
        }),
        handler: handler,
    }
}

func (c *OrderEventConsumer) Start(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return c.reader.Close()
        default:
            msg, err := c.reader.FetchMessage(ctx)
            if err != nil {
                if errors.Is(err, context.Canceled) {
                    return nil
                }
                log.Printf("fetch error: %v", err)
                continue
            }
            
            // Process message with retry
            if err := c.processWithRetry(ctx, msg); err != nil {
                // Send to dead letter queue
                c.sendToDLQ(ctx, msg, err)
            }
            
            // Manual commit after successful processing
            if err := c.reader.CommitMessages(ctx, msg); err != nil {
                log.Printf("commit error: %v", err)
            }
        }
    }
}

func (c *OrderEventConsumer) processWithRetry(ctx context.Context, msg kafka.Message) error {
    var lastErr error
    for attempt := 0; attempt < 3; attempt++ {
        if err := c.handler.Handle(ctx, msg.Value); err != nil {
            lastErr = err
            time.Sleep(time.Duration(attempt+1) * time.Second) // Exponential backoff
            continue
        }
        return nil
    }
    return lastErr
}
```

#### Idempotency (CRITICAL):
```go
type IdempotentHandler struct {
    redis   *redis.Client
    handler OrderEventHandler
}

func (h *IdempotentHandler) Handle(ctx context.Context, data []byte) error {
    var event OrderCreatedEvent
    if err := json.Unmarshal(data, &event); err != nil {
        return err
    }
    
    // Check if already processed
    key := fmt.Sprintf("processed:%s", event.EventID)
    exists, err := h.redis.Exists(ctx, key).Result()
    if err != nil {
        return err
    }
    if exists > 0 {
        log.Printf("Event %s already processed, skipping", event.EventID)
        return nil // Already processed
    }
    
    // Process the event
    if err := h.handler.Handle(ctx, data); err != nil {
        return err
    }
    
    // Mark as processed (with TTL)
    h.redis.Set(ctx, key, "1", 7*24*time.Hour)
    
    return nil
}
```

#### Dead Letter Queue (DLQ):
```go
func (c *OrderEventConsumer) sendToDLQ(ctx context.Context, msg kafka.Message, processErr error) {
    dlqWriter := &kafka.Writer{
        Addr:  kafka.TCP(c.brokers...),
        Topic: "orders-dlq",
    }
    defer dlqWriter.Close()
    
    dlqMessage := kafka.Message{
        Key:   msg.Key,
        Value: msg.Value,
        Headers: append(msg.Headers,
            kafka.Header{Key: "error", Value: []byte(processErr.Error())},
            kafka.Header{Key: "original_topic", Value: []byte("orders")},
            kafka.Header{Key: "failed_at", Value: []byte(time.Now().Format(time.RFC3339))},
        ),
    }
    
    if err := dlqWriter.WriteMessages(ctx, dlqMessage); err != nil {
        log.Printf("failed to send to DLQ: %v", err)
    }
}
```

---

### Week 3-4: Event Sourcing & Saga Pattern

#### Event Sourcing:
- [ ] Events as source of truth
- [ ] Event store design
- [ ] Rebuilding state from events
- [ ] Projections (read models)
- [ ] Snapshots for performance
- [ ] When to use (and when NOT to use)

#### Saga Pattern for Distributed Transactions:
```
┌─────────────────────────────────────────────────────────────────────────┐
│                        Order Saga (Choreography)                         │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│  ┌──────────┐    ┌──────────────┐    ┌──────────────┐    ┌───────────┐ │
│  │  Order   │───▶│  Inventory   │───▶│   Payment    │───▶│  Shipping │ │
│  │ Service  │    │   Service    │    │   Service    │    │  Service  │ │
│  └────┬─────┘    └──────┬───────┘    └──────┬───────┘    └───────────┘ │
│       │                 │                   │                           │
│       │ OrderCreated    │ InventoryReserved │ PaymentCompleted          │
│       ▼                 ▼                   ▼                           │
│  ┌─────────────────────────────────────────────────────────────────┐    │
│  │                         Kafka Topics                             │    │
│  └─────────────────────────────────────────────────────────────────┘    │
│                                                                          │
│  Compensating Transactions (on failure):                                 │
│  PaymentFailed ──▶ ReleaseInventory ──▶ CancelOrder                     │
│                                                                          │
└─────────────────────────────────────────────────────────────────────────┘
```

#### Code Example - Saga Orchestrator:
```go
type OrderSaga struct {
    orderRepo     OrderRepository
    inventoryPub  InventoryEventPublisher
    paymentPub    PaymentEventPublisher
    redis         *redis.Client
}

type SagaState struct {
    OrderID           string    `json:"order_id"`
    Status            string    `json:"status"`
    InventoryReserved bool      `json:"inventory_reserved"`
    PaymentProcessed  bool      `json:"payment_processed"`
    CreatedAt         time.Time `json:"created_at"`
}

func (s *OrderSaga) StartSaga(ctx context.Context, order *Order) error {
    // Initialize saga state
    state := SagaState{
        OrderID:   order.ID,
        Status:    "started",
        CreatedAt: time.Now(),
    }
    s.saveSagaState(ctx, state)
    
    // Step 1: Request inventory reservation
    err := s.inventoryPub.PublishReserveInventory(ctx, ReserveInventoryCommand{
        OrderID:  order.ID,
        Items:    order.Items,
        SagaID:   order.ID,
    })
    if err != nil {
        return s.compensate(ctx, state, err)
    }
    
    state.Status = "inventory_pending"
    s.saveSagaState(ctx, state)
    
    return nil
}

func (s *OrderSaga) HandleInventoryReserved(ctx context.Context, event InventoryReservedEvent) error {
    state, _ := s.loadSagaState(ctx, event.OrderID)
    state.InventoryReserved = true
    state.Status = "payment_pending"
    s.saveSagaState(ctx, state)
    
    // Step 2: Request payment
    return s.paymentPub.PublishProcessPayment(ctx, ProcessPaymentCommand{
        OrderID: event.OrderID,
        Amount:  event.TotalAmount,
        SagaID:  event.OrderID,
    })
}

func (s *OrderSaga) HandlePaymentCompleted(ctx context.Context, event PaymentCompletedEvent) error {
    state, _ := s.loadSagaState(ctx, event.OrderID)
    state.PaymentProcessed = true
    state.Status = "completed"
    s.saveSagaState(ctx, state)
    
    // Update order status
    return s.orderRepo.UpdateStatus(ctx, event.OrderID, "confirmed")
}

func (s *OrderSaga) HandlePaymentFailed(ctx context.Context, event PaymentFailedEvent) error {
    state, _ := s.loadSagaState(ctx, event.OrderID)
    
    // Compensate: Release inventory
    if state.InventoryReserved {
        s.inventoryPub.PublishReleaseInventory(ctx, ReleaseInventoryCommand{
            OrderID: event.OrderID,
        })
    }
    
    state.Status = "failed"
    s.saveSagaState(ctx, state)
    
    return s.orderRepo.UpdateStatus(ctx, event.OrderID, "cancelled")
}
```

#### Hands-on Project:
```
Build: Order Processing System with Kafka

Services:
1. Order Service - Creates orders, orchestrates saga
2. Inventory Service - Reserves/releases stock
3. Payment Service - Processes payments
4. Notification Service - Sends emails/SMS

Events:
- OrderCreated
- InventoryReserved / InventoryReservationFailed
- PaymentCompleted / PaymentFailed
- OrderConfirmed / OrderCancelled
- NotificationSent

Requirements:
- Idempotent event handling
- Dead letter queue for failures
- Saga state persistence
- Compensating transactions
- Event replay capability
```

#### Architecture Artifact:
Create **Event-Driven Architecture ADR**:
- Event catalog (all events documented)
- Event schema versioning strategy
- Consumer group design
- Failure handling patterns
- Monitoring and alerting

---

## MONTH 5: Cloud & Infrastructure

### Week 1-2: Kubernetes Deep Dive

#### Core Concepts:
- [ ] Pods (smallest deployable unit)
- [ ] Deployments (declarative updates)
- [ ] Services (network abstraction)
- [ ] Namespaces (resource isolation)
- [ ] Labels and selectors

#### Pod Specification:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: user-service
  labels:
    app: user-service
    environment: production
spec:
  containers:
  - name: user-service
    image: myregistry/user-service:v1.2.3
    ports:
    - containerPort: 8080
    
    # Resource management (CRITICAL)
    resources:
      requests:
        memory: "128Mi"
        cpu: "100m"
      limits:
        memory: "256Mi"
        cpu: "500m"
    
    # Health checks
    livenessProbe:
      httpGet:
        path: /health/live
        port: 8080
      initialDelaySeconds: 10
      periodSeconds: 10
      failureThreshold: 3
    
    readinessProbe:
      httpGet:
        path: /health/ready
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 5
    
    # Environment variables
    env:
    - name: DATABASE_URL
      valueFrom:
        secretKeyRef:
          name: user-service-secrets
          key: database-url
    - name: LOG_LEVEL
      valueFrom:
        configMapKeyRef:
          name: user-service-config
          key: log-level
    
    # Volume mounts
    volumeMounts:
    - name: config-volume
      mountPath: /app/config
      readOnly: true
  
  volumes:
  - name: config-volume
    configMap:
      name: user-service-config
```

#### Deployments:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
  namespace: production
spec:
  replicas: 3
  
  # Deployment strategy
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1        # Can create 1 extra pod during update
      maxUnavailable: 0  # Always keep all pods running
  
  selector:
    matchLabels:
      app: user-service
  
  template:
    metadata:
      labels:
        app: user-service
        version: v1.2.3
    spec:
      # Pod anti-affinity: spread across nodes
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchLabels:
                  app: user-service
              topologyKey: kubernetes.io/hostname
      
      containers:
      - name: user-service
        image: myregistry/user-service:v1.2.3
        # ... (same as pod spec above)
```

#### Services & Networking:
```yaml
# ClusterIP (internal)
apiVersion: v1
kind: Service
metadata:
  name: user-service
spec:
  type: ClusterIP
  selector:
    app: user-service
  ports:
  - port: 80
    targetPort: 8080

---
# Ingress (external)
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: api-ingress
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - api.example.com
    secretName: api-tls
  rules:
  - host: api.example.com
    http:
      paths:
      - path: /users
        pathType: Prefix
        backend:
          service:
            name: user-service
            port:
              number: 80
      - path: /orders
        pathType: Prefix
        backend:
          service:
            name: order-service
            port:
              number: 80
```

#### ConfigMaps & Secrets:
```yaml
# ConfigMap for non-sensitive config
apiVersion: v1
kind: ConfigMap
metadata:
  name: user-service-config
data:
  log-level: "info"
  cache-ttl: "300"
  feature-flags: |
    {
      "new_checkout": true,
      "dark_mode": false
    }

---
# Secret for sensitive data
apiVersion: v1
kind: Secret
metadata:
  name: user-service-secrets
type: Opaque
stringData:
  database-url: "postgres://user:pass@host:5432/db"
  jwt-secret: "super-secret-key"
```

#### Horizontal Pod Autoscaler:
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: user-service-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: user-service
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
      stabilizationWindowSeconds: 300  # Wait 5 min before scaling down
```

#### RBAC (Role-Based Access Control):
```yaml
# ServiceAccount for the application
apiVersion: v1
kind: ServiceAccount
metadata:
  name: user-service-sa
  namespace: production

---
# Role with specific permissions
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: user-service-role
  namespace: production
rules:
- apiGroups: [""]
  resources: ["configmaps", "secrets"]
  verbs: ["get", "list"]

---
# Bind role to service account
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: user-service-rolebinding
  namespace: production
subjects:
- kind: ServiceAccount
  name: user-service-sa
  namespace: production
roleRef:
  kind: Role
  name: user-service-role
  apiGroup: rbac.authorization.k8s.io
```

#### Key K8s Commands:
```bash
# Deployment management
kubectl apply -f deployment.yaml
kubectl rollout status deployment/user-service
kubectl rollout history deployment/user-service
kubectl rollout undo deployment/user-service

# Debugging
kubectl get pods -l app=user-service
kubectl describe pod user-service-xxx
kubectl logs user-service-xxx -f
kubectl exec -it user-service-xxx -- /bin/sh

# Scaling
kubectl scale deployment user-service --replicas=5
kubectl autoscale deployment user-service --min=2 --max=10 --cpu-percent=80

# Port forwarding for local testing
kubectl port-forward svc/user-service 8080:80
```

---

### Week 3-4: Terraform Deep Dive

#### Core Concepts:
- [ ] Providers (AWS, GCP, Azure, etc.)
- [ ] Resources (what you're creating)
- [ ] Data sources (read existing resources)
- [ ] Variables and outputs
- [ ] State management
- [ ] Modules (reusable components)

#### Project Structure:
```
infrastructure/
├── environments/
│   ├── dev/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── terraform.tfvars
│   ├── staging/
│   │   └── ...
│   └── prod/
│       └── ...
├── modules/
│   ├── vpc/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   ├── eks/
│   │   └── ...
│   ├── rds/
│   │   └── ...
│   └── redis/
│       └── ...
└── shared/
    └── backend.tf
```

#### VPC Module:
```hcl
# modules/vpc/main.tf

variable "name" {
  description = "VPC name"
  type        = string
}

variable "cidr_block" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "availability_zones" {
  description = "AZs to use"
  type        = list(string)
}

variable "environment" {
  description = "Environment name"
  type        = string
}

# VPC
resource "aws_vpc" "main" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = var.name
    Environment = var.environment
  }
}

# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${var.name}-igw"
  }
}

# Public Subnets
resource "aws_subnet" "public" {
  count                   = length(var.availability_zones)
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.cidr_block, 8, count.index)
  availability_zone       = var.availability_zones[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name = "${var.name}-public-${count.index + 1}"
    Type = "public"
  }
}

# Private Subnets
resource "aws_subnet" "private" {
  count             = length(var.availability_zones)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.cidr_block, 8, count.index + 10)
  availability_zone = var.availability_zones[count.index]

  tags = {
    Name = "${var.name}-private-${count.index + 1}"
    Type = "private"
  }
}

# NAT Gateway (for private subnet internet access)
resource "aws_eip" "nat" {
  count  = length(var.availability_zones)
  domain = "vpc"
}

resource "aws_nat_gateway" "main" {
  count         = length(var.availability_zones)
  allocation_id = aws_eip.nat[count.index].id
  subnet_id     = aws_subnet.public[count.index].id

  tags = {
    Name = "${var.name}-nat-${count.index + 1}"
  }
}

# Route Tables
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "${var.name}-public-rt"
  }
}

resource "aws_route_table" "private" {
  count  = length(var.availability_zones)
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.main[count.index].id
  }

  tags = {
    Name = "${var.name}-private-rt-${count.index + 1}"
  }
}

# Route Table Associations
resource "aws_route_table_association" "public" {
  count          = length(var.availability_zones)
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "private" {
  count          = length(var.availability_zones)
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private[count.index].id
}

# Outputs
output "vpc_id" {
  value = aws_vpc.main.id
}

output "public_subnet_ids" {
  value = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  value = aws_subnet.private[*].id
}
```

#### RDS Module:
```hcl
# modules/rds/main.tf

variable "identifier" {
  type = string
}

variable "engine_version" {
  type    = string
  default = "15.4"
}

variable "instance_class" {
  type    = string
  default = "db.t3.medium"
}

variable "allocated_storage" {
  type    = number
  default = 20
}

variable "database_name" {
  type = string
}

variable "username" {
  type      = string
  sensitive = true
}

variable "password" {
  type      = string
  sensitive = true
}

variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "allowed_security_groups" {
  type = list(string)
}

# Subnet Group
resource "aws_db_subnet_group" "main" {
  name       = "${var.identifier}-subnet-group"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "${var.identifier}-subnet-group"
  }
}

# Security Group
resource "aws_security_group" "rds" {
  name        = "${var.identifier}-rds-sg"
  description = "Security group for RDS"
  vpc_id      = var.vpc_id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = var.allowed_security_groups
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# RDS Instance
resource "aws_db_instance" "main" {
  identifier     = var.identifier
  engine         = "postgres"
  engine_version = var.engine_version
  instance_class = var.instance_class

  allocated_storage     = var.allocated_storage
  max_allocated_storage = var.allocated_storage * 5  # Autoscaling

  db_name  = var.database_name
  username = var.username
  password = var.password

  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.rds.id]

  # High availability
  multi_az = true

  # Backups
  backup_retention_period = 7
  backup_window          = "03:00-04:00"
  maintenance_window     = "Mon:04:00-Mon:05:00"

  # Performance
  performance_insights_enabled = true
  
  # Security
  storage_encrypted = true
  
  # Updates
  auto_minor_version_upgrade = true
  
  # Deletion protection
  deletion_protection = true
  skip_final_snapshot = false
  final_snapshot_identifier = "${var.identifier}-final-snapshot"

  tags = {
    Name = var.identifier
  }
}

output "endpoint" {
  value = aws_db_instance.main.endpoint
}

output "port" {
  value = aws_db_instance.main.port
}
```

#### Main Environment Configuration:
```hcl
# environments/prod/main.tf

terraform {
  required_version = ">= 1.5.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}

provider "aws" {
  region = var.aws_region
}

# VPC
module "vpc" {
  source = "../../modules/vpc"

  name               = "prod-vpc"
  cidr_block         = "10.0.0.0/16"
  availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c"]
  environment        = "production"
}

# RDS
module "rds" {
  source = "../../modules/rds"

  identifier        = "prod-database"
  database_name     = "myapp"
  username          = var.db_username
  password          = var.db_password
  vpc_id            = module.vpc.vpc_id
  subnet_ids        = module.vpc.private_subnet_ids
  allowed_security_groups = [module.eks.node_security_group_id]
}

# EKS
module "eks" {
  source = "../../modules/eks"

  cluster_name       = "prod-cluster"
  vpc_id             = module.vpc.vpc_id
  subnet_ids         = module.vpc.private_subnet_ids
  kubernetes_version = "1.28"
  
  node_groups = {
    general = {
      instance_types = ["t3.medium"]
      min_size       = 2
      max_size       = 10
      desired_size   = 3
    }
  }
}

# Outputs
output "database_endpoint" {
  value     = module.rds.endpoint
  sensitive = true
}

output "eks_cluster_endpoint" {
  value = module.eks.cluster_endpoint
}
```

#### State Management:
```hcl
# shared/backend.tf

# Create S3 bucket for state
resource "aws_s3_bucket" "terraform_state" {
  bucket = "my-terraform-state"

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_versioning" "terraform_state" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state" {
  bucket = aws_s3_bucket.terraform_state.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "aws:kms"
    }
  }
}

# DynamoDB for state locking
resource "aws_dynamodb_table" "terraform_locks" {
  name         = "terraform-locks"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
```

#### Key Terraform Commands:
```bash
# Initialize
terraform init

# Plan (preview changes)
terraform plan -var-file="terraform.tfvars"

# Apply changes
terraform apply -var-file="terraform.tfvars"

# Destroy (be careful!)
terraform destroy -var-file="terraform.tfvars"

# State management
terraform state list
terraform state show aws_vpc.main
terraform state mv aws_vpc.old aws_vpc.new

# Import existing resources
terraform import aws_vpc.main vpc-12345678

# Format code
terraform fmt -recursive

# Validate configuration
terraform validate
```

#### Hands-on Project:
```
Build: Complete AWS Infrastructure with Terraform

Components:
1. VPC with public/private subnets (3 AZs)
2. EKS cluster with managed node groups
3. RDS PostgreSQL (Multi-AZ)
4. ElastiCache Redis cluster
5. S3 buckets for assets
6. CloudFront distribution
7. Route53 DNS
8. ACM certificates

Requirements:
- Modular structure
- Environment separation (dev/staging/prod)
- Remote state with locking
- Secrets management (SSM Parameter Store)
- Tagging strategy
- Cost estimation with Infracost
```

#### Architecture Artifact:
Create **Infrastructure Architecture Document**:
- Network topology diagram
- Security group rules matrix
- Cost breakdown by service
- Disaster recovery plan
- Scaling strategy

---

## MONTH 5: Production Excellence

### Week 1-2: Observability

#### Topics to Master:
- [ ] Three pillars: Logs, Metrics, Traces
- [ ] OpenTelemetry integration
- [ ] Prometheus & Grafana
- [ ] Distributed tracing (Jaeger/X-Ray)
- [ ] Log aggregation (ELK, CloudWatch)
- [ ] Alerting strategies
- [ ] SLIs, SLOs, SLAs

#### Hands-on Project:
```
Add Observability to Order System:
- Structured logging (JSON)
- Custom metrics (order count, latency)
- Distributed traces across services
- Grafana dashboards
- PagerDuty/Slack alerts
```

---

### Week 3-4: Security Architecture

#### Topics to Master:
- [ ] OWASP Top 10
- [ ] API security (rate limiting, input validation)
- [ ] Secrets management (AWS Secrets Manager, Vault)
- [ ] Network security (WAF, Shield)
- [ ] Data encryption (at-rest, in-transit)
- [ ] IAM best practices (least privilege)
- [ ] Compliance basics (SOC2, GDPR)

#### Architecture Artifact:
Create **Security Threat Model**:
- STRIDE analysis
- Data flow diagrams
- Security controls by layer
- Incident response plan

---

## MONTH 6: Architecture Portfolio & Interview Prep

### Week 1-2: Case Study Documentation

Document your projects as **Solution Architecture Case Studies** (see templates below).

### Week 3-4: Interview Preparation

#### The Business Side (Crucial for SA):
- [ ] **Translate Tech to Value:** Practice explaining technical decisions in business terms.
    - *Bad:* "We need Kubernetes for auto-scaling."
    - *Good:* "We need to reduce our cloud bill by 30% during off-peak hours, which container orchestration allows us to do."
- [ ] **Negotiation:** Role-play scenarios where you have to say "No" to a feature request to save the architecture.

#### Interview Practice:
- [ ] Practice system design interviews
- [ ] Mock architecture reviews
- [ ] Behavioral questions (leadership, conflict resolution)

---

### 🎤 Interview Readiness Framework

> **Translate roadmap outputs into interview stories.**

Architect interviews test: **Narratives, Decision Defense, Trade-off Pressure, Ambiguity Handling**

#### The STAR-T Method for Architecture Stories:

| Element | Description | Example |
|---------|-------------|----------|
| **S**ituation | Business context | "E-commerce platform, 100K daily orders" |
| **T**ask | Your responsibility | "Design order processing to handle 10x scale" |
| **A**ction | What you designed | "Event-driven with Kafka, Saga for transactions" |
| **R**esult | Measurable outcome | "Reduced order failures from 2% to 0.1%" |
| **T**rade-off | What you sacrificed | "Added complexity, 2 more services to maintain" |

#### Prepare These 10 Stories:

| Story | Maps to Learning | Key Trade-off to Defend |
|-------|-----------------|------------------------|
| 1. Why Go over Node.js | Month 1-2 | Performance vs team familiarity |
| 2. Database schema design | Month 3 | Normalization vs query performance |
| 3. Caching strategy | Month 3 | Consistency vs latency |
| 4. Event-driven vs sync | Month 4 | Complexity vs scalability |
| 5. Saga vs 2PC | Month 4 | Eventual consistency vs simplicity |
| 6. K8s vs ECS | Month 5 | Flexibility vs operational overhead |
| 7. Multi-AZ design | Month 5 | Cost vs availability |
| 8. Observability investment | Month 5 | Upfront cost vs debugging time |
| 9. Security architecture | Month 5 | User friction vs protection |
| 10. Cost optimization | All months | Feature velocity vs cloud bill |

#### Decision Defense Practice:

For each major decision, prepare to answer:

1. **"Why did you choose X over Y?"**
   - State the constraints (time, team, budget)
   - List 3 options considered
   - Explain selection criteria
   - Acknowledge what you gave up

2. **"What would you do differently?"**
   - Never say "nothing"
   - Show growth mindset
   - Tie to lessons learned

3. **"How would this scale to 10x/100x?"**
   - Identify bottlenecks
   - Propose evolution path
   - Mention cost implications

4. **"What breaks first?"**
   - Database connections
   - Message queue backpressure
   - Cache stampedes
   - Network partitions

#### Ambiguity Handling Practice:

When given vague requirements:

```
1. Clarify scope: "Are we optimizing for latency or cost?"
2. State assumptions: "I'll assume 99.9% availability target"
3. Propose options: "We could do A, B, or C. Here's my recommendation."
4. Invite feedback: "Does this align with your priorities?"
```

#### Mock Interview Schedule (Week 3-4):

| Day | Activity | Duration |
|-----|----------|----------|
| Day 1 | Record yourself presenting Flagship case study | 1 hr |
| Day 2 | Practice with peer (system design) | 1.5 hr |
| Day 3 | Practice with peer (behavioral) | 1 hr |
| Day 4 | Review recordings, identify weak points | 1 hr |
| Day 5 | Practice weak areas | 1.5 hr |
| Day 6 | Full mock interview (friend/mentor) | 2 hr |
| Day 7 | Rest + light review | 30 min |

#### Interview Question Bank:

**System Design (Practice These):**
- [ ] Design a rate limiter
- [ ] Design a URL shortener
- [ ] Design a notification system
- [ ] Design an e-commerce order system
- [ ] Design a real-time chat application
- [ ] Design a metrics/monitoring system
- [ ] Design a multi-tenant SaaS platform

**Behavioral (SA-Specific):**
- [ ] "Tell me about a time you had to push back on a technical decision."
- [ ] "How do you handle disagreements with senior engineers?"
- [ ] "Describe a system you designed that failed. What did you learn?"
- [ ] "How do you balance technical debt with feature delivery?"
- [ ] "How do you communicate technical concepts to non-technical stakeholders?"

#### Your 30-Second Pitch:

> "I'm a technology leader with 15 years of experience across mobile, backend, and cloud. I've spent the last [X months] deepening my architecture skills—designing event-driven systems, mastering Go for high-performance backends, and building infrastructure with Terraform. My flagship project is a mobile commerce platform designed for 500K users, which I can walk you through end-to-end. I'm looking for a Solution Architect role where I can own technical direction and translate business goals into scalable systems."

---

## 📋 Case Study Templates

---

## CASE STUDY 1: Multi-Tenant SaaS Platform

### Business Context
Design a B2B SaaS platform serving 500+ enterprise customers with strict data isolation requirements.

### Requirements

**Functional:**
- Multi-tenant architecture (tenant isolation)
- User management per tenant
- Role-based access control
- API access for integrations
- Audit logging
- Subscription & billing

**Non-Functional:**
- 99.9% availability
- <200ms API latency (p95)
- Support 10,000 concurrent users
- GDPR & SOC2 compliance
- Cost: <$50k/month at scale

### Architecture Decisions

#### Decision 1: Tenant Isolation Strategy

| Option | Pros | Cons | Decision |
|--------|------|------|----------|
| Shared DB, Shared Schema | Cost-effective, simple | No isolation, noisy neighbor | ❌ |
| Shared DB, Separate Schema | Good isolation, moderate cost | Schema management complex | ✅ Selected |
| Separate DB per Tenant | Complete isolation | Expensive, operational overhead | ❌ |

**Rationale:** Schema-per-tenant provides strong isolation while keeping infrastructure costs manageable. Row-level security adds additional protection.

#### Decision 2: Authentication Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Client    │────▶│  Auth0/     │────▶│  API        │
│   (React)   │     │  Cognito    │     │  Gateway    │
└─────────────┘     └─────────────┘     └─────────────┘
                           │                    │
                           ▼                    ▼
                    ┌─────────────┐     ┌─────────────┐
                    │   Token     │     │  Services   │
                    │   (JWT)     │     │ (Go/Gin)    │
                    └─────────────┘     └─────────────┘
```

#### Decision 3: Data Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    PostgreSQL (RDS)                      │
├─────────────────────────────────────────────────────────┤
│  ┌──────────┐  ┌──────────┐  ┌──────────┐               │
│  │ tenant_a │  │ tenant_b │  │ tenant_c │  ...          │
│  │ (schema) │  │ (schema) │  │ (schema) │               │
│  └──────────┘  └──────────┘  └──────────┘               │
├─────────────────────────────────────────────────────────┤
│  ┌──────────────────────────────────────────────────┐   │
│  │               public (shared metadata)            │   │
│  │  - tenants, subscriptions, billing                │   │
│  └──────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

### High-Level Architecture Diagram

```
                                 ┌─────────────────┐
                                 │   CloudFront    │
                                 │     (CDN)       │
                                 └────────┬────────┘
                                          │
                                 ┌────────▼────────┐
                                 │   API Gateway   │
                                 │   + WAF         │
                                 └────────┬────────┘
                                          │
                    ┌─────────────────────┼─────────────────────┐
                    │                     │                     │
           ┌────────▼────────┐   ┌────────▼────────┐   ┌────────▼────────┐
           │  User Service   │   │  Core Service   │   │  Billing Service│
           │    (ECS)        │   │    (ECS)        │   │    (ECS)        │
           └────────┬────────┘   └────────┬────────┘   └────────┬────────┘
                    │                     │                     │
                    └─────────────────────┼─────────────────────┘
                                          │
                    ┌─────────────────────┼─────────────────────┐
                    │                     │                     │
           ┌────────▼────────┐   ┌────────▼────────┐   ┌────────▼────────┐
           │   PostgreSQL    │   │     Redis       │   │   EventBridge   │
           │    (RDS)        │   │  (ElastiCache)  │   │                 │
           └─────────────────┘   └─────────────────┘   └─────────────────┘
```

### Cost Estimate (At Scale)

| Component | Specification | Monthly Cost |
|-----------|---------------|--------------|
| ECS Fargate | 6 services × 2 tasks × 1vCPU/2GB | $800 |
| RDS PostgreSQL | db.r6g.xlarge, Multi-AZ | $1,200 |
| ElastiCache | cache.r6g.large, 2 nodes | $400 |
| ALB | 2 load balancers | $50 |
| Data Transfer | ~5TB/month | $400 |
| CloudWatch | Logs + Metrics | $200 |
| **Total** | | **~$3,050/month** |

### Failure Scenarios & Mitigations

| Failure | Impact | Mitigation |
|---------|--------|------------|
| Database failure | Complete outage | Multi-AZ RDS, automated failover |
| Service crash | Partial degradation | Multiple ECS tasks, health checks |
| Redis failure | Increased latency | ElastiCache Multi-AZ, fallback to DB |
| Region outage | Complete outage | DR plan with cross-region backup |

### What I Would Do Differently at 10x Scale

1. **Move to separate databases** for largest tenants (enterprise tier)
2. **Add read replicas** for reporting workloads
3. **Implement CQRS** for high-read scenarios
4. **Consider Aurora Serverless** for cost optimization
5. **Add global CDN** with regional API deployments

---

## CASE STUDY 2: Real-Time Order Processing System

### Business Context
Design an order processing system for an e-commerce platform handling 100k orders/day with real-time inventory updates.

### Requirements

**Functional:**
- Order placement & validation
- Real-time inventory management
- Payment processing
- Order status tracking
- Notifications (email, SMS, push)
- Returns & refunds

**Non-Functional:**
- Process 1,200 orders/minute (peak)
- <2 second order confirmation
- Zero overselling
- 99.95% availability
- Audit trail for all transactions

### Event-Driven Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Order API  │────▶│   Kafka     │────▶│  Inventory  │
│             │     │             │     │  Service    │
└─────────────┘     └──────┬──────┘     └─────────────┘
                           │
          ┌────────────────┼────────────────┐
          │                │                │
          ▼                ▼                ▼
   ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
   │  Payment    │  │  Shipping   │  │  Notification│
   │  Service    │  │  Service    │  │  Service    │
   └─────────────┘  └─────────────┘  └─────────────┘
```

### Event Flow

```
1. OrderCreated
   ├──▶ InventoryService: Reserve stock
   │    └──▶ InventoryReserved / InventoryFailed
   │
   ├──▶ PaymentService: Process payment
   │    └──▶ PaymentCompleted / PaymentFailed
   │
   └──▶ NotificationService: Send confirmation

2. If all succeed:
   └──▶ OrderConfirmed
        └──▶ ShippingService: Create shipment

3. If any fails:
   └──▶ Compensating transactions (Saga)
        ├──▶ Release inventory
        ├──▶ Refund payment
        └──▶ Notify customer
```

### Saga Pattern for Distributed Transactions

```
┌─────────────────────────────────────────────────────────────────┐
│                        Order Saga                                │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐   │
│  │ Reserve  │───▶│ Process  │───▶│ Create   │───▶│ Complete │   │
│  │Inventory │    │ Payment  │    │ Shipment │    │  Order   │   │
│  └────┬─────┘    └────┬─────┘    └────┬─────┘    └──────────┘   │
│       │               │               │                          │
│       │ Fail          │ Fail          │ Fail                     │
│       ▼               ▼               ▼                          │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐                    │
│  │  Cancel  │◀───│  Refund  │◀───│ Release  │                    │
│  │  Order   │    │ Payment  │    │Inventory │                    │
│  └──────────┘    └──────────┘    └──────────┘                    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

### Kafka Topic Design

| Topic | Partitions | Retention | Key |
|-------|------------|-----------|-----|
| orders.created | 12 | 7 days | orderId |
| orders.confirmed | 12 | 7 days | orderId |
| orders.failed | 6 | 30 days | orderId |
| inventory.reserved | 12 | 7 days | productId |
| inventory.released | 6 | 7 days | productId |
| payments.completed | 12 | 30 days | orderId |
| payments.failed | 6 | 30 days | orderId |
| notifications.send | 6 | 3 days | userId |

### Database Schema (PostgreSQL)

```sql
-- Orders table (partitioned by month)
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
) PARTITION BY RANGE (created_at);

-- Order items
CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id),
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    CONSTRAINT positive_quantity CHECK (quantity > 0)
);

-- Inventory with optimistic locking
CREATE TABLE inventory (
    product_id UUID PRIMARY KEY,
    available_quantity INTEGER NOT NULL DEFAULT 0,
    reserved_quantity INTEGER NOT NULL DEFAULT 0,
    version INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT non_negative_available CHECK (available_quantity >= 0),
    CONSTRAINT non_negative_reserved CHECK (reserved_quantity >= 0)
);

-- Idempotency keys
CREATE TABLE idempotency_keys (
    key VARCHAR(255) PRIMARY KEY,
    response JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Index for order lookups
CREATE INDEX idx_orders_user_status ON orders(user_id, status);
CREATE INDEX idx_orders_created_at ON orders(created_at);
```

### Handling Race Conditions (Inventory)

```go
// Optimistic locking for inventory updates in Go
func (s *InventoryService) ReserveInventory(ctx context.Context, productID string, quantity int) error {
    const maxRetries = 3
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        inventory, err := s.GetInventory(ctx, productID)
        if err != nil {
            return fmt.Errorf("failed to get inventory: %w", err)
        }
        
        if inventory.AvailableQuantity < quantity {
            return ErrInsufficientStock
        }
        
        result, err := s.db.ExecContext(ctx, `
            UPDATE inventory 
            SET 
                available_quantity = available_quantity - $1,
                reserved_quantity = reserved_quantity + $1,
                version = version + 1,
                updated_at = NOW()
            WHERE product_id = $2 
              AND version = $3
              AND available_quantity >= $1
        `, quantity, productID, inventory.Version)
        
        if err != nil {
            return fmt.Errorf("failed to update inventory: %w", err)
        }
        
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected > 0 {
            return nil // Success
        }
        
        // Retry on conflict with exponential backoff
        time.Sleep(time.Duration(math.Pow(2, float64(attempt))) * 10 * time.Millisecond)
    }
    
    return ErrConflict
}
```

### Cost Estimate

| Component | Specification | Monthly Cost |
|-----------|---------------|--------------|
| MSK (Kafka) | 3 brokers, kafka.m5.large | $1,200 |
| ECS Fargate | 8 services × 3 tasks | $1,500 |
| RDS PostgreSQL | db.r6g.2xlarge, Multi-AZ | $2,400 |
| ElastiCache | cache.r6g.xlarge cluster | $800 |
| SES (Email) | 1M emails | $100 |
| SNS (SMS) | 100k SMS | $750 |
| **Total** | | **~$6,750/month** |

---

## CASE STUDY 3: IoT Data Ingestion Platform

### Business Context
Design a platform to ingest, process, and analyze data from 1M+ IoT devices sending telemetry every 30 seconds.

### Scale Requirements

- 1,000,000 devices
- Data every 30 seconds
- = 33,000 messages/second
- = 2M messages/minute
- = ~86B messages/month

### Architecture

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           IoT Devices (1M+)                              │
└───────────────────────────────────┬─────────────────────────────────────┘
                                    │ MQTT
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                         AWS IoT Core                                     │
│                    (Managed MQTT Broker)                                 │
└───────────────────────────────────┬─────────────────────────────────────┘
                                    │ IoT Rules
                    ┌───────────────┼───────────────┐
                    ▼               ▼               ▼
             ┌───────────┐   ┌───────────┐   ┌───────────┐
             │  Kinesis  │   │    S3     │   │  Lambda   │
             │  Streams  │   │ (Raw Data)│   │ (Alerts)  │
             └─────┬─────┘   └───────────┘   └─────┬─────┘
                   │                               │
                   ▼                               ▼
             ┌───────────┐                   ┌───────────┐
             │  Kinesis  │                   │    SNS    │
             │ Analytics │                   │ (Notify)  │
             └─────┬─────┘                   └───────────┘
                   │
         ┌─────────┴─────────┐
         ▼                   ▼
   ┌───────────┐       ┌───────────┐
   │ Timestream│       │ OpenSearch│
   │ (Metrics) │       │  (Logs)   │
   └─────┬─────┘       └─────┬─────┘
         │                   │
         └─────────┬─────────┘
                   ▼
             ┌───────────┐
             │  Grafana  │
             │(Dashboard)│
             └───────────┘
```

### Data Flow

1. **Ingestion Layer**
   - Devices connect via MQTT to AWS IoT Core
   - IoT Rules route messages to Kinesis

2. **Processing Layer**
   - Kinesis Data Analytics for real-time aggregation
   - Lambda for threshold-based alerts

3. **Storage Layer**
   - TimeStream for time-series metrics (hot data: 7 days)
   - S3 for raw data archive (cold storage)
   - OpenSearch for log search

4. **Presentation Layer**
   - Grafana dashboards
   - API for historical queries

### Message Schema

```json
{
  "deviceId": "device-123456",
  "timestamp": "2024-01-15T10:30:00Z",
  "type": "telemetry",
  "payload": {
    "temperature": 23.5,
    "humidity": 45.2,
    "battery": 87,
    "signalStrength": -65
  },
  "metadata": {
    "firmwareVersion": "2.1.0",
    "location": {
      "lat": 37.7749,
      "lng": -122.4194
    }
  }
}
```

### Cost Estimate (1M Devices)

| Component | Specification | Monthly Cost |
|-----------|---------------|--------------|
| IoT Core | 86B messages | $3,000 |
| Kinesis | 10 shards | $1,500 |
| Timestream | 10TB writes, 50TB storage | $4,000 |
| S3 | 100TB storage | $2,300 |
| Lambda | Alert processing | $500 |
| OpenSearch | 3 nodes | $1,200 |
| **Total** | | **~$12,500/month** |

---

## 🎯 Architecture Artifacts Checklist

For each case study, create:

- [ ] **Context Diagram** (C4 Level 1)
- [ ] **Container Diagram** (C4 Level 2)
- [ ] **Component Diagram** (C4 Level 3) for key services
- [ ] **Sequence Diagrams** for critical flows
- [ ] **Database Schema Diagram**
- [ ] **Infrastructure Diagram** (AWS)
- [ ] **ADRs** (3-5 key decisions)
- [ ] **NFR Table** (latency, throughput, availability targets)
- [ ] **Cost Breakdown**
- [ ] **Risk Register**
- [ ] **Runbook** for operations

---

## 📚 Learning Resources

### Books (Must Read)
1. **Designing Data-Intensive Applications** - Martin Kleppmann
2. **Fundamentals of Software Architecture** - Mark Richards
3. **Building Microservices** - Sam Newman
4. **System Design Interview Vol 1 & 2** - Alex Xu

### Go-Specific Books
1. **Learning Go** - Jon Bodner (Start here)
2. **Concurrency in Go** - Katherine Cox-Buday
3. **100 Go Mistakes and How to Avoid Them** - Teiva Harsanyi
4. **Let's Go** - Alex Edwards (Web development)

### Courses
1. **AWS Solutions Architect Professional** - A Cloud Guru
2. **Distributed Systems** - MIT OpenCourseWare
3. **Kafka for Developers** - Confluent
4. **Go: The Complete Developer's Guide** - Udemy
5. **Ardan Labs Go Training** - https://www.ardanlabs.com (Advanced)

### Go Resources
1. **A Tour of Go** - https://go.dev/tour (Official)
2. **Go by Example** - https://gobyexample.com
3. **Exercism Go Track** - https://exercism.io/tracks/go
4. **Gin Docs** - https://gin-gonic.com/docs/
5. **Go Project Layout** - https://github.com/golang-standards/project-layout

### Practice Platforms
1. **System Design Primer** - GitHub
2. **DesignGurus.io** - System Design
3. **ByteByteGo** - Newsletter

---

## ✅ Weekly Progress Tracker

### Week 1-2: Go Fundamentals
- [ ] Complete A Tour of Go
- [ ] Build CLI tool project
- [ ] Build concurrent downloader
- [ ] Understand error handling patterns

### Week 3-4: Go + Gin
- [ ] Learn Gin framework basics
- [ ] Build User Management API
- [ ] Implement middleware patterns
- [ ] Write API documentation

### Week 5-6: Auth + Database
- [ ] Implement JWT authentication
- [ ] PostgreSQL integration with Go
- [ ] Database migrations setup
- [ ] Rate limiting implementation

### Week 7-8: Database Deep Dive
- [ ] Design e-commerce schema
- [ ] Practice query optimization
- [ ] Implement Redis caching layer
- [ ] Connection pooling setup

### Week 9-10: Event-Driven
- [ ] Set up local Kafka
- [ ] Build order processing in Go
- [ ] Implement saga pattern
- [ ] Kafka consumer groups

### Week 11-12: AWS
- [ ] Deploy Go services to AWS
- [ ] Write Terraform
- [ ] Set up monitoring
- [ ] ECS/EKS deployment

### Week 13-14: Observability
- [ ] Add distributed tracing (OpenTelemetry)
- [ ] Create Grafana dashboards
- [ ] Define SLOs
- [ ] Prometheus metrics in Go

### Week 15-16: Portfolio
- [ ] Document Case Study 1
- [ ] Document Case Study 2
- [ ] Document Case Study 3
- [ ] Update resume with Go experience

---

## 🏆 Success Metrics

You're ready for Solution Architect roles when you can:

### Technical Proof (Execution)
1. ✅ Design any system in 45 minutes with clear trade-offs
2. ✅ Present 3 documented case studies with C4 diagrams
3. ✅ Explain cost implications of architecture decisions
4. ✅ Discuss failure modes and mitigations fluently
5. ✅ Build production-ready backend services in Go
6. ✅ Explain concurrency patterns (goroutines, channels)
7. ✅ Confidently discuss Go vs. other languages for different use cases

### Credibility Proof (Visibility)
8. ✅ Have your Flagship case study documented with all artifacts
9. ✅ Published 10+ architecture posts/insights on LinkedIn
10. ✅ Can deliver your 30-second pitch without hesitation
11. ✅ Have 3 practiced stories ready for behavioral questions
12. ✅ Received feedback on your case study from at least 2 peers

### Certification (Optional but Helpful)
13. ⬜ AWS Solutions Architect Professional (validates cloud depth)

---

## 📊 Alignment Scorecard

Track your progress monthly:

| Dimension | Target | Month 1 | Month 2 | Month 3 | Month 4 | Month 5 | Month 6 |
|-----------|--------|---------|---------|---------|---------|---------|----------|
| E2E Ownership | ⭐⭐⭐⭐⭐ | | | | | | |
| Technical Depth | ⭐⭐⭐⭐⭐ | | | | | | |
| Business Thinking | ⭐⭐⭐⭐⭐ | | | | | | |
| Real-world Credibility | ⭐⭐⭐⭐⭐ | | | | | | |
| Networking & Visibility | ⭐⭐⭐⭐⭐ | | | | | | |
| Interview Readiness | ⭐⭐⭐⭐⭐ | | | | | | |

---

## 💡 Why Go Will Pay Off

| Benefit | Impact |
|---------|--------|
| Cloud-native credibility | K8s, Docker, Terraform all in Go |
| Performance conversations | Can discuss memory, GC, latency |
| Backend depth | Shows you're not "just a frontend dev" |
| Hiring advantage | Go demand growing rapidly |
| Simplicity | Easy to read others' code in reviews |

---

*Last Updated: December 2024*
*Learning Path: Go + Gin (industry standard)*
*Next Review: Monthly progress check*
