# Nx Monorepo Mastery Guide
## From Beginner to Expert Level

---

## Table of Contents
1. [Beginner Level](#beginner-level)
2. [Intermediate Level](#intermediate-level)
3. [Advanced Level](#advanced-level)
4. [Expert Level](#expert-level)
5. [Real-World Use Cases](#real-world-use-cases)
6. [Practice Questions with Answers](#practice-questions-with-answers)

---

# Beginner Level

## 1. Introduction to Nx

### What is Nx?
Nx is a powerful build system with first-class monorepo support and powerful integrations. It helps you develop, test, build, and scale applications with ease.

### Why Use Nx?

| Feature | Benefit |
|---------|---------|
| **Smart Builds** | Only rebuild what's affected by changes |
| **Computation Caching** | Never rebuild the same code twice |
| **Dependency Graph** | Visualize project dependencies |
| **Code Generation** | Consistent scaffolding with generators |
| **Module Boundaries** | Enforce architectural constraints |
| **Remote Caching** | Share cache across team members |
| **Distributed Execution** | Parallelize CI across machines |

### Monorepo vs Polyrepo

```
┌─────────────────────────────────────────────────────────────────┐
│                         MONOREPO                                │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │                    Single Repository                        ││
│  │  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐         ││
│  │  │  App 1  │  │  App 2  │  │  Lib 1  │  │  Lib 2  │         ││
│  │  │ (React) │  │(Angular)│  │ (Shared)│  │  (UI)   │         ││
│  │  └─────────┘  └─────────┘  └─────────┘  └─────────┘         ││
│  │       │            │            ▲            ▲              ││
│  │       └────────────┴────────────┴────────────┘              ││
│  │              Shared code, single version                    ││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                         POLYREPO                                │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐         │
│  │  Repo 1  │  │  Repo 2  │  │  Repo 3  │  │  Repo 4  │         │
│  │  App 1   │  │  App 2   │  │  Lib 1   │  │  Lib 2   │         │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘         │
│       ↓             ↓             ↓             ↓               │
│  npm install   npm install   npm install   npm install          │
│  Separate versions, separate CI/CD pipelines                    │
└─────────────────────────────────────────────────────────────────┘
```

---

## 2. Setting Up Nx Workspace

### Creating a New Workspace

```bash
# Create a new Nx workspace
npx create-nx-workspace@latest my-org

# Options during setup:
# - Stack: React, Angular, Vue, Node, None (integrated or package-based)
# - Application name
# - Stylesheet format
# - Enable Nx Cloud (recommended)

# Create workspace with specific preset
npx create-nx-workspace@latest my-org --preset=react-monorepo
npx create-nx-workspace@latest my-org --preset=angular-monorepo
npx create-nx-workspace@latest my-org --preset=ts
```

### Workspace Structure

```
my-org/
├── apps/                    # Application projects
│   ├── web-app/            # React/Angular/Vue app
│   └── api/                # Node.js backend
├── libs/                    # Shared libraries
│   ├── shared/
│   │   └── utils/          # Shared utilities
│   └── ui/                 # UI component library
├── tools/                   # Custom scripts and tools
├── nx.json                  # Nx configuration
├── package.json            # Root package.json
├── tsconfig.base.json      # Base TypeScript config
└── .eslintrc.json          # ESLint configuration
```

### Understanding nx.json

```json
// nx.json
{
  "$schema": "./node_modules/nx/schemas/nx-schema.json",
  "defaultBase": "main",
  "namedInputs": {
    "default": ["{projectRoot}/**/*", "sharedGlobals"],
    "production": [
      "default",
      "!{projectRoot}/**/?(*.)+(spec|test).[jt]s?(x)?(.snap)",
      "!{projectRoot}/tsconfig.spec.json",
      "!{projectRoot}/.eslintrc.json"
    ],
    "sharedGlobals": []
  },
  "targetDefaults": {
    "build": {
      "dependsOn": ["^build"],
      "inputs": ["production", "^production"],
      "cache": true
    },
    "test": {
      "inputs": ["default", "^production", "{workspaceRoot}/jest.preset.js"],
      "cache": true
    },
    "lint": {
      "inputs": ["default", "{workspaceRoot}/.eslintrc.json"],
      "cache": true
    }
  },
  "plugins": [
    {
      "plugin": "@nx/webpack/plugin",
      "options": {
        "buildTargetName": "build",
        "serveTargetName": "serve"
      }
    }
  ]
}
```

---

## 3. Core Nx Commands

### Running Tasks

```bash
# Run a target for a specific project
nx build my-app
nx test my-app
nx lint my-app
nx serve my-app

# Run multiple targets
nx run-many -t build test lint

# Run target for specific projects
nx run-many -t build -p app1 app2

# Run target for all projects
nx run-many -t build --all

# Run affected projects (based on git changes)
nx affected -t build
nx affected -t test
nx affected -t lint

# Parallel execution
nx run-many -t build --parallel=5
```

### Generating Code

```bash
# Generate a new application
nx generate @nx/react:application my-app
nx generate @nx/angular:application my-app
nx generate @nx/node:application my-api

# Generate a new library
nx generate @nx/react:library my-lib
nx generate @nx/js:library shared-utils

# Generate components
nx generate @nx/react:component my-component --project=my-lib

# List available generators
nx list @nx/react
nx list @nx/angular

# Shorthand
nx g @nx/react:app my-app
nx g @nx/react:lib my-lib
```

### Exploring the Workspace

```bash
# View project graph in browser
nx graph

# View task graph for a specific target
nx build my-app --graph

# Show project details
nx show project my-app
nx show project my-app --web

# List all projects
nx show projects

# Show affected projects
nx show projects --affected
```

---

## 4. Project Configuration

### project.json Structure

```json
// apps/my-app/project.json
{
  "name": "my-app",
  "projectType": "application",
  "sourceRoot": "apps/my-app/src",
  "prefix": "app",
  "tags": ["scope:app", "type:feature"],
  "targets": {
    "build": {
      "executor": "@nx/webpack:webpack",
      "outputs": ["{options.outputPath}"],
      "options": {
        "outputPath": "dist/apps/my-app",
        "main": "apps/my-app/src/main.tsx",
        "tsConfig": "apps/my-app/tsconfig.app.json",
        "assets": ["apps/my-app/src/assets"]
      },
      "configurations": {
        "production": {
          "optimization": true,
          "extractLicenses": true,
          "sourceMap": false
        },
        "development": {
          "optimization": false,
          "sourceMap": true
        }
      },
      "defaultConfiguration": "production"
    },
    "serve": {
      "executor": "@nx/webpack:dev-server",
      "options": {
        "buildTarget": "my-app:build:development"
      }
    },
    "test": {
      "executor": "@nx/jest:jest",
      "outputs": ["{workspaceRoot}/coverage/{projectRoot}"],
      "options": {
        "jestConfig": "apps/my-app/jest.config.ts"
      }
    },
    "lint": {
      "executor": "@nx/eslint:lint"
    }
  }
}
```

### Library Configuration

```json
// libs/shared/utils/project.json
{
  "name": "shared-utils",
  "projectType": "library",
  "sourceRoot": "libs/shared/utils/src",
  "tags": ["scope:shared", "type:util"],
  "targets": {
    "build": {
      "executor": "@nx/js:tsc",
      "outputs": ["{options.outputPath}"],
      "options": {
        "outputPath": "dist/libs/shared/utils",
        "main": "libs/shared/utils/src/index.ts",
        "tsConfig": "libs/shared/utils/tsconfig.lib.json"
      }
    },
    "test": {
      "executor": "@nx/jest:jest",
      "outputs": ["{workspaceRoot}/coverage/{projectRoot}"],
      "options": {
        "jestConfig": "libs/shared/utils/jest.config.ts"
      }
    },
    "lint": {
      "executor": "@nx/eslint:lint"
    }
  }
}
```

---

## 5. Understanding the Project Graph

### What is the Project Graph?

The project graph is Nx's understanding of your workspace structure and dependencies between projects.

```
┌─────────────────────────────────────────────────────────────────┐
│                      Project Graph                              │
│                                                                 │
│     ┌──────────┐           ┌──────────┐                         │
│     │  Web App │           │   Admin  │                         │
│     │          │           │   App    │                         │
│     └────┬─────┘           └────┬─────┘                         │
│          │                      │                               │
│          ▼                      ▼                               │
│     ┌──────────┐           ┌──────────┐                         │
│     │ Feature  │           │ Feature  │                         │
│     │   Auth   │           │  Admin   │                         │
│     └────┬─────┘           └────┬─────┘                         │
│          │                      │                               │
│          └──────────┬───────────┘                               │
│                     ▼                                           │
│               ┌──────────┐                                      │
│               │    UI    │                                      │
│               │Components│                                      │
│               └────┬─────┘                                      │
│                    │                                            │
│                    ▼                                            │
│               ┌──────────┐                                      │
│               │  Shared  │                                      │
│               │  Utils   │                                      │
│               └──────────┘                                      │
└─────────────────────────────────────────────────────────────────┘
```

### Viewing Dependencies

```bash
# Open interactive project graph
nx graph

# View affected projects graph
nx affected:graph

# Export graph as JSON
nx graph --file=output.json

# Focus on specific project
nx graph --focus=my-app

# Exclude projects
nx graph --exclude=e2e
```

---

# Intermediate Level

## 6. Caching

### Local Caching

Nx caches task outputs to avoid redundant computations.

```json
// nx.json - Configure caching
{
  "targetDefaults": {
    "build": {
      "cache": true,
      "inputs": ["production", "^production"],
      "outputs": ["{options.outputPath}"]
    },
    "test": {
      "cache": true,
      "inputs": ["default", "^production"]
    },
    "lint": {
      "cache": true,
      "inputs": ["default"]
    }
  }
}
```

### Understanding Inputs and Outputs

```json
// Named inputs in nx.json
{
  "namedInputs": {
    "default": [
      "{projectRoot}/**/*",
      "sharedGlobals"
    ],
    "production": [
      "default",
      "!{projectRoot}/**/*.spec.ts",
      "!{projectRoot}/jest.config.ts"
    ],
    "sharedGlobals": [
      "{workspaceRoot}/tsconfig.base.json"
    ]
  }
}
```

### Cache Invalidation

```bash
# Skip cache for a run
nx build my-app --skip-nx-cache

# Clear local cache
nx reset

# View cache folder
ls node_modules/.cache/nx
```

### Configuring Custom Caching

```json
// project.json - Custom inputs/outputs
{
  "targets": {
    "build": {
      "executor": "@nx/webpack:webpack",
      "inputs": [
        "production",
        "^production",
        { "env": "NODE_ENV" },
        { "externalDependencies": ["webpack"] }
      ],
      "outputs": [
        "{options.outputPath}",
        "{workspaceRoot}/dist/.cache/{projectRoot}"
      ],
      "cache": true
    }
  }
}
```

---

## 7. Task Pipeline Configuration

### Defining Task Dependencies

```json
// nx.json
{
  "targetDefaults": {
    "build": {
      "dependsOn": ["^build"]  // Build dependencies first
    },
    "test": {
      "dependsOn": ["build"]   // Build this project first
    },
    "deploy": {
      "dependsOn": ["build", "test", "lint"]
    }
  }
}
```

### Understanding dependsOn Syntax

```json
{
  "targetDefaults": {
    "build": {
      // ^ means dependencies (projects this depends on)
      "dependsOn": ["^build"]
    },
    "test": {
      // No ^ means same project
      "dependsOn": ["build"]
    },
    "e2e": {
      // Specific project target
      "dependsOn": [
        { "projects": ["api"], "target": "serve" }
      ]
    },
    "deploy": {
      // Multiple dependencies
      "dependsOn": [
        "build",
        "test",
        { "projects": "self", "target": "lint" }
      ]
    }
  }
}
```

### Visualizing Task Pipeline

```bash
# View task graph
nx build my-app --graph

# The graph shows execution order:
# 1. shared-utils:build
# 2. ui-components:build
# 3. feature-auth:build
# 4. my-app:build
```

---

## 8. Affected Commands

### How Affected Works

```bash
# Show affected projects
nx show projects --affected

# Run tests only for affected projects
nx affected -t test

# Run multiple targets
nx affected -t lint test build

# Compare against specific branch
nx affected -t test --base=main --head=HEAD

# Compare specific commits
nx affected -t test --base=abc123 --head=def456
```

### Configuring Base Branch

```json
// nx.json
{
  "defaultBase": "main",  // or "master", "develop"
  "affected": {
    "defaultBase": "main"
  }
}
```

### CI Integration Example

```yaml
# .github/workflows/ci.yml
name: CI

on:
  pull_request:
    branches: [main]

jobs:
  main:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0  # Important for affected
      
      - name: Setup Node
        uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: 'npm'
      
      - name: Install dependencies
        run: npm ci
      
      - name: Derive SHAs
        uses: nrwl/nx-set-shas@v4
      
      - name: Run affected tests
        run: npx nx affected -t lint test build
```

---

## 9. Module Boundaries

### Defining Project Tags

```json
// libs/shared/utils/project.json
{
  "name": "shared-utils",
  "tags": ["scope:shared", "type:util"]
}

// libs/feature/auth/project.json
{
  "name": "feature-auth",
  "tags": ["scope:feature", "type:feature"]
}

// apps/web-app/project.json
{
  "name": "web-app",
  "tags": ["scope:app", "type:app"]
}
```

### Configuring ESLint Boundaries

```javascript
// eslint.config.mjs
import nx from '@nx/eslint-plugin';

export default [
  ...nx.configs['flat/base'],
  ...nx.configs['flat/typescript'],
  ...nx.configs['flat/javascript'],
  {
    files: ['**/*.ts', '**/*.tsx', '**/*.js', '**/*.jsx'],
    rules: {
      '@nx/enforce-module-boundaries': [
        'error',
        {
          enforceBuildableLibDependency: true,
          allow: [],
          depConstraints: [
            // Apps can depend on features and shared
            {
              sourceTag: 'type:app',
              onlyDependOnLibsWithTags: ['type:feature', 'type:ui', 'type:util']
            },
            // Features can depend on UI and utils
            {
              sourceTag: 'type:feature',
              onlyDependOnLibsWithTags: ['type:ui', 'type:util', 'type:data-access']
            },
            // UI can only depend on utils
            {
              sourceTag: 'type:ui',
              onlyDependOnLibsWithTags: ['type:util']
            },
            // Utils cannot depend on anything
            {
              sourceTag: 'type:util',
              onlyDependOnLibsWithTags: []
            },
            // Scope restrictions
            {
              sourceTag: 'scope:admin',
              onlyDependOnLibsWithTags: ['scope:admin', 'scope:shared']
            },
            {
              sourceTag: 'scope:client',
              onlyDependOnLibsWithTags: ['scope:client', 'scope:shared']
            }
          ]
        }
      ]
    }
  }
];
```

### Library Types Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    Library Types Hierarchy                      │
│                                                                 │
│     ┌─────────────────────────────────────────────────────────┐ │
│     │                    type:app                             │ │
│     │              (Applications only)                        │ │
│     └───────────────────────┬─────────────────────────────────┘ │
│                             │                                   │
│                             ▼                                   │
│     ┌─────────────────────────────────────────────────────────┐ │
│     │                  type:feature                           │ │
│     │         (Feature modules, smart components)             │ │
│     └───────────────────────┬─────────────────────────────────┘ │
│                             │                                   │
│              ┌──────────────┼──────────────┐                    │
│              ▼              ▼              ▼                    │
│     ┌──────────────┐ ┌──────────────┐ ┌──────────────┐          │
│     │   type:ui    │ │type:data-acc │ │  type:util   │          │
│     │  (Dumb UI)   │ │ (API calls)  │ │  (Helpers)   │          │
│     └──────────────┘ └──────────────┘ └──────────────┘          │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## 10. Generators

### Using Built-in Generators

```bash
# React generators
nx g @nx/react:app web-app
nx g @nx/react:lib feature-auth
nx g @nx/react:component Button --project=ui

# Angular generators
nx g @nx/angular:app admin-app
nx g @nx/angular:lib data-access
nx g @nx/angular:component header --project=ui

# Node generators
nx g @nx/node:app api
nx g @nx/node:lib shared-utils

# Nest generators
nx g @nx/nest:app api
nx g @nx/nest:module users --project=api
nx g @nx/nest:service users --project=api
nx g @nx/nest:controller users --project=api
```

### Generator Options

```bash
# Dry run (preview changes)
nx g @nx/react:lib my-lib --dry-run

# Skip formatting
nx g @nx/react:lib my-lib --skipFormat

# Specify directory
nx g @nx/react:lib my-lib --directory=libs/shared

# With routing
nx g @nx/react:app my-app --routing

# Standalone components (Angular)
nx g @nx/angular:component button --project=ui --standalone
```

---

# Advanced Level

## 11. Remote Caching (Nx Cloud)

### Setting Up Nx Cloud

```bash
# Connect to Nx Cloud
npx nx connect

# Or during workspace creation
npx create-nx-workspace@latest --nxCloud=yes
```

### Configuration

```json
// nx.json with Nx Cloud
{
  "nxCloudAccessToken": "your-access-token",
  // or use environment variable
  "nxCloudId": "your-cloud-id"
}
```

### Environment Variable Setup

```bash
# .env or CI environment
NX_CLOUD_ACCESS_TOKEN=your-access-token

# For CI pipelines
export NX_CLOUD_ACCESS_TOKEN=${{ secrets.NX_CLOUD_TOKEN }}
```

### Benefits of Remote Caching

```
Without Remote Cache:
Developer A: nx build app  → 5 minutes
Developer B: nx build app  → 5 minutes (rebuilds)
CI Pipeline: nx build app  → 5 minutes (rebuilds)

With Remote Cache:
Developer A: nx build app  → 5 minutes (uploads cache)
Developer B: nx build app  → 5 seconds (cache hit!)
CI Pipeline: nx build app  → 5 seconds (cache hit!)
```

---

## 12. Distributed Task Execution (Nx Agents)

### Configuring Nx Agents

```yaml
# .github/workflows/ci.yml with Nx Agents
name: CI

on:
  pull_request:
    branches: [main]

jobs:
  main:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: 'npm'
      
      - run: npm ci
      
      - name: Start Nx Agents
        run: npx nx-cloud start-ci-run --distribute-on="3 linux-medium-js"
      
      - name: Run affected commands
        run: npx nx affected -t lint test build e2e
```

### Dynamic Agent Allocation

```yaml
# Distribute based on PR size
- name: Start Nx Agents
  run: |
    npx nx-cloud start-ci-run \
      --distribute-on="auto" \
      --stop-agents-after="e2e"
```

### Task Distribution Visualization

```
┌─────────────────────────────────────────────────────────────────┐
│                    Distributed Execution                        │
│                                                                 │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │                    Main CI Machine                        │  │
│  │                  (Orchestrator)                           │  │
│  └─────────────────────────┬─────────────────────────────────┘  │
│                            │                                    │
│         ┌──────────────────┼──────────────────┐                 │
│         ▼                  ▼                  ▼                 │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│  │   Agent 1    │  │   Agent 2    │  │   Agent 3    │           │
│  │ - lib-a:build│  │ - lib-b:build│  │ - lib-c:build│           │
│  │ - lib-a:test │  │ - lib-b:test │  │ - app:build  │           │
│  │ - lib-d:lint │  │ - app:test   │  │ - app:e2e    │           │
│  └──────────────┘  └──────────────┘  └──────────────┘           │
│                                                                 │
│  Tasks are distributed based on:                                │
│  • Historical execution time                                    │
│  • Task dependencies                                            │
│  • Available agents                                             │
└─────────────────────────────────────────────────────────────────┘
```

---

## 13. Creating Custom Generators

### Setting Up Local Plugin

```bash
# Generate a local plugin
nx g @nx/plugin:plugin my-plugin --directory=tools/my-plugin

# Generate a generator in the plugin
nx g @nx/plugin:generator my-generator --project=my-plugin
```

### Generator Structure

```
tools/my-plugin/
├── src/
│   ├── generators/
│   │   └── my-generator/
│   │       ├── generator.ts
│   │       ├── generator.spec.ts
│   │       ├── schema.json
│   │       └── files/
│   │           └── src/
│   │               └── __fileName__.ts.template
│   └── index.ts
├── generators.json
└── project.json
```

### Generator Implementation

```typescript
// tools/my-plugin/src/generators/my-generator/generator.ts
import {
  Tree,
  formatFiles,
  generateFiles,
  names,
  offsetFromRoot,
  readProjectConfiguration,
  updateProjectConfiguration,
} from '@nx/devkit';
import * as path from 'path';
import { MyGeneratorSchema } from './schema';

interface NormalizedSchema extends MyGeneratorSchema {
  projectName: string;
  projectRoot: string;
  fileName: string;
  className: string;
}

function normalizeOptions(
  tree: Tree,
  options: MyGeneratorSchema
): NormalizedSchema {
  const name = names(options.name).fileName;
  const projectRoot = `libs/${options.directory}/${name}`;

  return {
    ...options,
    projectName: name,
    projectRoot,
    fileName: name,
    className: names(options.name).className,
  };
}

export async function myGenerator(tree: Tree, options: MyGeneratorSchema) {
  const normalizedOptions = normalizeOptions(tree, options);

  // Generate files from templates
  generateFiles(
    tree,
    path.join(__dirname, 'files'),
    normalizedOptions.projectRoot,
    {
      ...normalizedOptions,
      ...names(options.name),
      offsetFromRoot: offsetFromRoot(normalizedOptions.projectRoot),
      template: '',
    }
  );

  // Update project configuration
  const projectConfig = readProjectConfiguration(tree, normalizedOptions.projectName);
  projectConfig.tags = [...(projectConfig.tags || []), 'custom-generated'];
  updateProjectConfiguration(tree, normalizedOptions.projectName, projectConfig);

  await formatFiles(tree);

  return () => {
    console.log(`Generated ${normalizedOptions.projectName} successfully!`);
  };
}

export default myGenerator;
```

### Generator Schema

```json
// tools/my-plugin/src/generators/my-generator/schema.json
{
  "$schema": "http://json-schema.org/schema",
  "$id": "MyGenerator",
  "title": "My Generator",
  "description": "Creates a custom library with predefined structure",
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "description": "The name of the library",
      "$default": {
        "$source": "argv",
        "index": 0
      },
      "x-prompt": "What name would you like to use for the library?"
    },
    "directory": {
      "type": "string",
      "description": "The directory to place the library",
      "default": "shared"
    },
    "tags": {
      "type": "string",
      "description": "Comma-separated tags for the library"
    },
    "skipTests": {
      "type": "boolean",
      "description": "Skip generating test files",
      "default": false
    }
  },
  "required": ["name"]
}
```

### Template Files

```typescript
// files/src/__fileName__.ts.template
/**
 * <%= className %> Library
 * Generated by my-generator
 */

export interface <%= className %>Options {
  name: string;
  enabled: boolean;
}

export function create<%= className %>(options: <%= className %>Options) {
  return {
    ...options,
    createdAt: new Date(),
  };
}

export default create<%= className %>;
```

### Using Custom Generator

```bash
# Run custom generator
nx g @my-org/my-plugin:my-generator awesome-lib --directory=features

# Dry run
nx g @my-org/my-plugin:my-generator awesome-lib --dry-run
```

---

## 14. Creating Custom Executors

### Executor Structure

```
tools/my-plugin/
├── src/
│   ├── executors/
│   │   └── my-executor/
│   │       ├── executor.ts
│   │       ├── executor.spec.ts
│   │       └── schema.json
│   └── index.ts
├── executors.json
└── project.json
```

### Executor Implementation

```typescript
// tools/my-plugin/src/executors/my-executor/executor.ts
import { ExecutorContext, logger } from '@nx/devkit';
import { execSync } from 'child_process';
import { MyExecutorSchema } from './schema';

export default async function runExecutor(
  options: MyExecutorSchema,
  context: ExecutorContext
): Promise<{ success: boolean }> {
  const projectName = context.projectName;
  const projectRoot = context.projectsConfigurations.projects[projectName].root;

  logger.info(`Running my-executor for ${projectName}`);
  logger.info(`Project root: ${projectRoot}`);
  logger.info(`Options: ${JSON.stringify(options)}`);

  try {
    // Example: Run a custom build command
    const command = options.command || 'echo "No command specified"';
    
    execSync(command, {
      cwd: projectRoot,
      stdio: 'inherit',
      env: {
        ...process.env,
        NODE_ENV: options.production ? 'production' : 'development',
      },
    });

    // Post-build steps
    if (options.postBuild) {
      logger.info('Running post-build steps...');
      execSync(options.postBuild, {
        cwd: projectRoot,
        stdio: 'inherit',
      });
    }

    return { success: true };
  } catch (error) {
    logger.error(`Executor failed: ${error.message}`);
    return { success: false };
  }
}
```

### Executor Schema

```json
// tools/my-plugin/src/executors/my-executor/schema.json
{
  "$schema": "http://json-schema.org/schema",
  "title": "My Executor",
  "description": "Runs custom build process",
  "type": "object",
  "properties": {
    "command": {
      "type": "string",
      "description": "The command to execute"
    },
    "production": {
      "type": "boolean",
      "description": "Run in production mode",
      "default": false
    },
    "postBuild": {
      "type": "string",
      "description": "Command to run after build"
    },
    "outputPath": {
      "type": "string",
      "description": "Output directory"
    }
  },
  "required": ["command"]
}
```

### Using Custom Executor

```json
// project.json
{
  "targets": {
    "custom-build": {
      "executor": "@my-org/my-plugin:my-executor",
      "options": {
        "command": "npm run build:custom",
        "outputPath": "dist/apps/my-app"
      },
      "configurations": {
        "production": {
          "production": true,
          "postBuild": "npm run compress"
        }
      }
    }
  }
}
```

---

## 15. Inferred Tasks (Project Crystal)

### How Inferred Tasks Work

Nx plugins can automatically detect and configure tasks based on configuration files.

```
┌─────────────────────────────────────────────────────────────────┐
│                    Inferred Tasks Flow                          │
│                                                                 │
│  ┌──────────────┐                                               │
│  │ vite.config.ts│──────┐                                       │
│  └──────────────┘       │                                       │
│                         │                                       │
│  ┌──────────────┐       │     ┌──────────────────┐              │
│  │ jest.config.ts│──────┼────▶│   @nx/vite       │              │
│  └──────────────┘       │     │   @nx/jest       │              │
│                         │     │   @nx/eslint     │              │
│  ┌──────────────┐       │     └────────┬─────────┘              │
│  │ .eslintrc.js │───────┘              │                        │
│  └──────────────┘                      ▼                        │
│                              ┌──────────────────┐               │
│                              │ Inferred Targets │               │
│                              │ • build          │               │
│                              │ • serve          │               │
│                              │ • test           │               │
│                              │ • lint           │               │
│                              └──────────────────┘               │
└─────────────────────────────────────────────────────────────────┘
```

### Configuring Inferred Tasks

```json
// nx.json
{
  "plugins": [
    {
      "plugin": "@nx/vite/plugin",
      "options": {
        "buildTargetName": "build",
        "serveTargetName": "dev",
        "previewTargetName": "preview",
        "testTargetName": "test"
      }
    },
    {
      "plugin": "@nx/eslint/plugin",
      "options": {
        "targetName": "lint"
      }
    },
    {
      "plugin": "@nx/jest/plugin",
      "options": {
        "targetName": "test"
      }
    }
  ]
}
```

### Viewing Inferred Targets

```bash
# See all targets for a project (including inferred)
nx show project my-app --web

# View in JSON format
nx show project my-app
```

---

# Expert Level

## 16. Nx Release

### Configuring Nx Release

```json
// nx.json
{
  "release": {
    "projects": ["packages/*"],
    "projectsRelationship": "independent",
    "changelog": {
      "projectChangelogs": {
        "createRelease": "github"
      }
    },
    "version": {
      "preVersionCommand": "npx nx run-many -t build",
      "conventionalCommits": true
    },
    "git": {
      "commit": true,
      "tag": true
    }
  }
}
```

### Release Commands

```bash
# Version packages (interactive)
nx release version

# Version with specific bump
nx release version minor
nx release version 1.2.3

# Generate changelog
nx release changelog

# Publish to npm
nx release publish

# Full release (version + changelog + publish)
nx release

# Dry run
nx release --dry-run

# First release (no previous versions)
nx release --first-release
```

### Release Groups

```json
// nx.json
{
  "release": {
    "groups": {
      "core": {
        "projects": ["packages/core", "packages/utils"],
        "version": {
          "generatorOptions": {
            "groupPrerelease": true
          }
        }
      },
      "plugins": {
        "projects": ["packages/plugin-*"],
        "projectsRelationship": "independent"
      }
    }
  }
}
```

---

## 17. Advanced CI/CD Patterns

### Complete CI Pipeline

```yaml
# .github/workflows/ci.yml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  NX_CLOUD_ACCESS_TOKEN: ${{ secrets.NX_CLOUD_ACCESS_TOKEN }}

jobs:
  main:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Set SHAs for affected
        uses: nrwl/nx-set-shas@v4

      - name: Start Nx Agents
        run: npx nx-cloud start-ci-run --distribute-on="5 linux-medium-js"

      - name: Run Nx Commands
        run: |
          npx nx-cloud record -- npx nx format:check
          npx nx affected -t lint test build --parallel=3
          npx nx affected -t e2e --parallel=1

      - name: Stop Nx Agents
        if: always()
        run: npx nx-cloud stop-all-agents

  deploy:
    needs: main
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Build for production
        run: npx nx run-many -t build --configuration=production

      - name: Deploy
        run: npx nx affected -t deploy
```

### Splitting E2E Tests

```yaml
# Run E2E tests in parallel with Nx Atomizer
- name: Run E2E tests
  run: |
    npx nx affected -t e2e \
      --split-e2e-tasks \
      --parallel=5
```

### Caching Dependencies

```yaml
- name: Cache node_modules
  uses: actions/cache@v4
  with:
    path: |
      node_modules
      ~/.cache/Cypress
    key: ${{ runner.os }}-node-${{ hashFiles('**/package-lock.json') }}
    restore-keys: |
      ${{ runner.os }}-node-
```

---

## 18. Workspace Presets and Templates

### Creating a Custom Preset

```typescript
// packages/my-preset/src/generators/preset/generator.ts
import {
  Tree,
  addDependenciesToPackageJson,
  generateFiles,
  installPackagesTask,
  updateJson,
} from '@nx/devkit';
import * as path from 'path';

export default async function presetGenerator(
  tree: Tree,
  options: { name: string }
) {
  // Update nx.json with organization defaults
  updateJson(tree, 'nx.json', (json) => ({
    ...json,
    defaultProject: options.name,
    targetDefaults: {
      build: {
        dependsOn: ['^build'],
        cache: true,
      },
      test: {
        cache: true,
      },
      lint: {
        cache: true,
      },
    },
    workspaceLayout: {
      appsDir: 'apps',
      libsDir: 'libs',
    },
  }));

  // Add default dependencies
  addDependenciesToPackageJson(
    tree,
    {
      // runtime dependencies
    },
    {
      '@nx/react': 'latest',
      '@nx/node': 'latest',
      '@nx/jest': 'latest',
      '@nx/eslint': 'latest',
    }
  );

  // Generate default files
  generateFiles(tree, path.join(__dirname, 'files'), '.', {
    name: options.name,
    template: '',
  });

  return () => {
    installPackagesTask(tree);
  };
}
```

### Using Custom Preset

```bash
# Create workspace with custom preset
npx create-nx-workspace@latest my-org \
  --preset=@my-org/my-preset \
  --name=my-app
```

---

## 19. Module Federation

### Setting Up Module Federation

```bash
# Generate host application
nx g @nx/react:host shell --remotes=shop,cart

# Generate remote applications
nx g @nx/react:remote shop --host=shell
nx g @nx/react:remote cart --host=shell
```

### Module Federation Configuration

```typescript
// apps/shell/module-federation.config.ts
import { ModuleFederationConfig } from '@nx/webpack';

const config: ModuleFederationConfig = {
  name: 'shell',
  remotes: ['shop', 'cart'],
  shared: (libraryName, sharedConfig) => {
    if (libraryName === 'react' || libraryName === 'react-dom') {
      return {
        singleton: true,
        strictVersion: true,
        requiredVersion: sharedConfig.requiredVersion,
      };
    }
    return sharedConfig;
  },
};

export default config;
```

### Dynamic Remotes

```typescript
// apps/shell/src/app/app.tsx
import { loadRemoteModule } from '@nx/react/mf';
import { Suspense, lazy } from 'react';

const RemoteShop = lazy(() =>
  loadRemoteModule('shop', './Module').then((m) => ({ default: m.RemoteEntry }))
);

export function App() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <RemoteShop />
    </Suspense>
  );
}
```

---

## 20. Performance Optimization

### Optimizing Build Performance

```json
// nx.json
{
  "parallel": 5,
  "cacheDirectory": ".nx/cache",
  "useDaemonProcess": true,
  "targetDefaults": {
    "build": {
      "cache": true,
      "inputs": ["production", "^production"],
      "outputs": ["{options.outputPath}"]
    }
  }
}
```

### Using Batch Mode

```bash
# Run tasks in batch mode for faster execution
nx run-many -t build --batch

# Configure batch mode in target
{
  "targets": {
    "build": {
      "executor": "@nx/js:tsc",
      "options": {
        "batchMode": true
      }
    }
  }
}
```

### Daemon Process

```bash
# Check daemon status
nx daemon

# Reset daemon
nx daemon --stop
nx reset
```

### Memory and Process Management

```bash
# Increase Node memory
NODE_OPTIONS="--max-old-space-size=8192" nx build my-app

# Limit parallel processes
nx run-many -t build --parallel=3
```

---

# Real-World Use Cases

## Use Case 1: E-commerce Platform

### Workspace Structure

```
my-ecommerce/
├── apps/
│   ├── storefront/          # Customer-facing React app
│   ├── admin-dashboard/     # Admin Angular app
│   ├── api-gateway/         # Node.js API gateway
│   ├── order-service/       # Microservice
│   ├── inventory-service/   # Microservice
│   └── e2e/                 # E2E tests
├── libs/
│   ├── shared/
│   │   ├── types/           # TypeScript interfaces
│   │   ├── utils/           # Common utilities
│   │   └── constants/       # Shared constants
│   ├── ui/
│   │   ├── components/      # Shared UI components
│   │   └── styles/          # Design system
│   ├── data-access/
│   │   ├── products/        # Product API client
│   │   ├── orders/          # Order API client
│   │   └── auth/            # Auth service
│   └── features/
│       ├── checkout/        # Checkout feature
│       ├── product-catalog/ # Product listing
│       └── cart/            # Shopping cart
└── tools/
    └── generators/          # Custom generators
```

### Module Boundary Configuration

```javascript
// eslint.config.mjs
{
  rules: {
    '@nx/enforce-module-boundaries': [
      'error',
      {
        depConstraints: [
          // Apps
          { sourceTag: 'type:app', onlyDependOnLibsWithTags: ['type:feature', 'type:ui', 'type:data-access', 'type:util'] },
          
          // Features
          { sourceTag: 'type:feature', onlyDependOnLibsWithTags: ['type:ui', 'type:data-access', 'type:util'] },
          
          // Data Access
          { sourceTag: 'type:data-access', onlyDependOnLibsWithTags: ['type:util'] },
          
          // UI
          { sourceTag: 'type:ui', onlyDependOnLibsWithTags: ['type:util'] },
          
          // Utils - no dependencies
          { sourceTag: 'type:util', onlyDependOnLibsWithTags: [] },
          
          // Scope restrictions
          { sourceTag: 'scope:storefront', onlyDependOnLibsWithTags: ['scope:shared', 'scope:storefront'] },
          { sourceTag: 'scope:admin', onlyDependOnLibsWithTags: ['scope:shared', 'scope:admin'] },
        ]
      }
    ]
  }
}
```

---

## Use Case 2: Design System

### Workspace Structure

```
design-system/
├── apps/
│   ├── docs/                # Documentation site (Storybook)
│   └── playground/          # Component playground
├── libs/
│   ├── components/
│   │   ├── atoms/           # Basic components (Button, Input)
│   │   ├── molecules/       # Composed components (FormField)
│   │   ├── organisms/       # Complex components (Header, Modal)
│   │   └── templates/       # Page templates
│   ├── tokens/              # Design tokens
│   ├── icons/               # Icon library
│   ├── hooks/               # Shared React hooks
│   └── utils/               # Utility functions
└── tools/
    ├── generators/
    │   └── component/       # Component generator
    └── executors/
        └── build-tokens/    # Token build executor
```

### Component Generator

```typescript
// tools/generators/component/generator.ts
import { Tree, generateFiles, names, formatFiles } from '@nx/devkit';
import * as path from 'path';

interface ComponentSchema {
  name: string;
  category: 'atoms' | 'molecules' | 'organisms' | 'templates';
  withStory: boolean;
  withTests: boolean;
}

export default async function componentGenerator(
  tree: Tree,
  schema: ComponentSchema
) {
  const componentNames = names(schema.name);
  const projectRoot = `libs/components/${schema.category}`;
  const componentPath = `${projectRoot}/src/lib/${componentNames.fileName}`;

  generateFiles(tree, path.join(__dirname, 'files'), componentPath, {
    ...componentNames,
    category: schema.category,
    withStory: schema.withStory,
    withTests: schema.withTests,
    template: '',
  });

  // Update barrel export
  const indexPath = `${projectRoot}/src/index.ts`;
  const indexContent = tree.read(indexPath, 'utf-8');
  const newExport = `export * from './lib/${componentNames.fileName}/${componentNames.fileName}';`;
  
  if (!indexContent.includes(newExport)) {
    tree.write(indexPath, `${indexContent}\n${newExport}`);
  }

  await formatFiles(tree);
}
```

---

## Use Case 3: Full-Stack Application

### Integrated Frontend and Backend

```
full-stack-app/
├── apps/
│   ├── web/                 # React frontend
│   ├── api/                 # NestJS backend
│   ├── web-e2e/            # Playwright E2E tests
│   └── mobile/             # React Native app
├── libs/
│   ├── api/
│   │   ├── core/           # API core (guards, interceptors)
│   │   ├── users/          # Users module
│   │   └── products/       # Products module
│   ├── web/
│   │   ├── core/           # Web core (routing, state)
│   │   ├── features/       # Feature modules
│   │   └── ui/             # UI components
│   └── shared/
│       ├── types/          # Shared TypeScript types
│       ├── validators/     # Shared validation
│       └── constants/      # Shared constants
└── tools/
    └── scripts/
```

### Shared Types Between Frontend and Backend

```typescript
// libs/shared/types/src/lib/user.interface.ts
export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  role: UserRole;
  createdAt: Date;
}

export enum UserRole {
  ADMIN = 'admin',
  USER = 'user',
  GUEST = 'guest',
}

export interface CreateUserDto {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
}

export interface UserResponse {
  user: User;
  token: string;
}
```

```typescript
// libs/api/users/src/lib/users.controller.ts
import { Controller, Post, Body } from '@nestjs/common';
import { CreateUserDto, UserResponse } from '@my-org/shared/types';
import { UsersService } from './users.service';

@Controller('users')
export class UsersController {
  constructor(private usersService: UsersService) {}

  @Post()
  async create(@Body() dto: CreateUserDto): Promise<UserResponse> {
    return this.usersService.create(dto);
  }
}
```

```typescript
// libs/web/features/auth/src/lib/use-auth.ts
import { CreateUserDto, UserResponse } from '@my-org/shared/types';

export function useAuth() {
  const register = async (data: CreateUserDto): Promise<UserResponse> => {
    const response = await fetch('/api/users', {
      method: 'POST',
      body: JSON.stringify(data),
      headers: { 'Content-Type': 'application/json' },
    });
    return response.json();
  };

  return { register };
}
```

---

# Practice Questions with Answers

## Beginner Level Questions

### Q1: What is the difference between `nx affected` and `nx run-many`?

**Answer:**

| Command | Description | Use Case |
|---------|-------------|----------|
| `nx affected` | Runs tasks only on projects affected by changes | CI pipelines, targeted testing |
| `nx run-many` | Runs tasks on specified or all projects | Full builds, running all tests |

```bash
# affected - only runs on changed projects
nx affected -t test  # Runs tests for projects affected by git changes

# run-many - runs on all or specified projects
nx run-many -t test  # Runs tests for ALL projects
nx run-many -t test -p app1 lib1  # Runs tests for specific projects
```

**When to use:**
- `nx affected`: In CI for PRs, during development
- `nx run-many`: Full workspace verification, release builds

---

### Q2: How does Nx caching work?

**Answer:**

Nx caching works by computing a hash of all inputs and storing outputs:

```
┌─────────────────────────────────────────────────────────────────┐
│                       Nx Caching Flow                           │
│                                                                 │
│  Inputs (Hash Calculation):                                     │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │ • Source files                                           │   │
│  │ • Dependencies                                           │   │
│  │ • Environment variables                                  │   │
│  │ • Runtime (Node version)                                 │   │
│  │ • Command arguments                                      │   │
│  └──────────────────────────────────────────────────────────┘   │
│                            │                                    │
│                            ▼                                    │
│                    ┌──────────────┐                             │
│                    │   Hash: abc  │                             │
│                    └──────┬───────┘                             │
│                           │                                     │
│              ┌────────────┴────────────┐                        │
│              ▼                         ▼                        │
│     ┌──────────────┐          ┌──────────────┐                  │
│     │ Cache Miss   │          │  Cache Hit   │                  │
│     │ Run Task     │          │ Restore      │                  │
│     │ Store Output │          │ Output       │                  │
│     └──────────────┘          └──────────────┘                  │
└─────────────────────────────────────────────────────────────────┘
```

**Configuration:**
```json
// nx.json
{
  "targetDefaults": {
    "build": {
      "cache": true,
      "inputs": ["production", "^production"],
      "outputs": ["{options.outputPath}"]
    }
  }
}
```

---

### Q3: What are project tags and why are they important?

**Answer:**

Tags are labels applied to projects that enable:
1. **Module boundary enforcement** - Restrict which projects can depend on others
2. **Selective task execution** - Run tasks on projects with specific tags
3. **Organization** - Categorize projects by scope, type, or team

```json
// project.json
{
  "name": "feature-auth",
  "tags": ["scope:shared", "type:feature", "team:platform"]
}
```

**Usage:**
```bash
# Run tests for projects with specific tag
nx run-many -t test --projects=tag:scope:shared

# Enforce boundaries
# "type:feature" can only depend on "type:util" and "type:ui"
```

---

### Q4: How do you create a new library in Nx?

**Answer:**

```bash
# Using Nx generators
nx generate @nx/react:library my-lib
nx generate @nx/angular:library my-lib
nx generate @nx/js:library my-lib

# With options
nx g @nx/react:library my-lib \
  --directory=libs/shared \
  --tags="scope:shared,type:util" \
  --buildable \
  --publishable --importPath=@my-org/my-lib

# Shorthand
nx g @nx/react:lib my-lib
```

**Generator creates:**
```
libs/shared/my-lib/
├── src/
│   ├── lib/
│   │   └── my-lib.ts
│   └── index.ts
├── project.json
├── tsconfig.json
├── tsconfig.lib.json
└── jest.config.ts
```

---

### Q5: What is the project graph and how do you view it?

**Answer:**

The project graph is Nx's representation of your workspace structure and dependencies.

```bash
# View interactive graph in browser
nx graph

# View affected projects
nx affected:graph

# Export as JSON
nx graph --file=graph.json

# Focus on specific project
nx graph --focus=my-app
```

**Project Graph Information:**
- All projects in workspace
- Dependencies between projects
- Which projects are affected by changes
- Task execution order

---

## Intermediate Level Questions

### Q6: Explain the task pipeline and how to configure it.

**Answer:**

The task pipeline defines task dependencies and execution order.

```json
// nx.json
{
  "targetDefaults": {
    "build": {
      // ^ means dependency projects
      "dependsOn": ["^build"],
      "cache": true
    },
    "test": {
      // Same project dependency
      "dependsOn": ["build"]
    },
    "deploy": {
      // Multiple dependencies
      "dependsOn": [
        "build",
        "test",
        { "projects": "self", "target": "lint" }
      ]
    },
    "e2e": {
      // Specific project target
      "dependsOn": [
        { "projects": ["api"], "target": "serve" }
      ]
    }
  }
}
```

**Execution Flow:**
```
nx build my-app
    │
    ├── ^build (dependency libs)
    │   ├── shared-utils:build
    │   ├── ui-components:build
    │   └── feature-auth:build
    │
    └── my-app:build (after dependencies)
```

---

### Q7: How do you set up module boundaries?

**Answer:**

**Step 1: Define Tags**
```json
// libs/shared/utils/project.json
{ "tags": ["scope:shared", "type:util"] }

// libs/feature/auth/project.json
{ "tags": ["scope:feature", "type:feature"] }

// apps/web/project.json
{ "tags": ["scope:app", "type:app"] }
```

**Step 2: Configure ESLint**
```javascript
// eslint.config.mjs
{
  rules: {
    '@nx/enforce-module-boundaries': [
      'error',
      {
        depConstraints: [
          {
            sourceTag: 'type:app',
            onlyDependOnLibsWithTags: ['type:feature', 'type:ui', 'type:util']
          },
          {
            sourceTag: 'type:feature',
            onlyDependOnLibsWithTags: ['type:ui', 'type:util', 'type:data-access']
          },
          {
            sourceTag: 'type:util',
            onlyDependOnLibsWithTags: []  // No dependencies allowed
          }
        ]
      }
    ]
  }
}
```

**Step 3: Run Lint**
```bash
nx lint my-app
# Error if boundaries are violated
```

---

### Q8: How does remote caching with Nx Cloud work?

**Answer:**

```
┌─────────────────────────────────────────────────────────────────┐
│                    Remote Caching Flow                          │
│                                                                 │
│  Developer A                      Developer B                   │
│  ┌──────────┐                    ┌──────────┐                   │
│  │ nx build │                    │ nx build │                   │
│  └────┬─────┘                    └────┬─────┘                   │
│       │                               │                         │
│       ▼                               ▼                         │
│  ┌──────────┐                    ┌──────────┐                   │
│  │  Hash    │                    │  Hash    │                   │
│  │  abc123  │                    │  abc123  │                   │
│  └────┬─────┘                    └────┬─────┘                   │
│       │                               │                         │
│       ▼                               ▼                         │
│  Cache Miss                      Cache Hit!                     │
│  Run Build                       Download                       │
│  Upload to Cloud                 from Cloud                     │
│       │                               │                         │
│       └───────────┬───────────────────┘                         │
│                   ▼                                             │
│          ┌────────────────┐                                     │
│          │   Nx Cloud     │                                     │
│          │  Cache Storage │                                     │
│          └────────────────┘                                     │
└─────────────────────────────────────────────────────────────────┘
```

**Setup:**
```bash
# Connect to Nx Cloud
npx nx connect

# Or add token to nx.json
{
  "nxCloudAccessToken": "your-token"
}
```

---

### Q9: How do you run only affected E2E tests in CI?

**Answer:**

```yaml
# .github/workflows/ci.yml
name: CI

on: [pull_request]

jobs:
  e2e:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0  # Required for affected
      
      - uses: actions/setup-node@v4
        with:
          node-version: 20
      
      - run: npm ci
      
      - uses: nrwl/nx-set-shas@v4  # Sets NX_BASE and NX_HEAD
      
      - name: Run affected E2E tests
        run: npx nx affected -t e2e --parallel=1
```

**Key Points:**
- `fetch-depth: 0` - Required to get full git history
- `nrwl/nx-set-shas@v4` - Sets base/head for affected calculation
- `--parallel=1` - E2E tests typically run sequentially

---

### Q10: What are inputs and outputs in caching configuration?

**Answer:**

**Inputs** - Files/values that affect task output (invalidate cache if changed):
```json
{
  "inputs": [
    "default",                    // Named input
    "{projectRoot}/**/*",         // All project files
    "{workspaceRoot}/jest.config.ts", // Workspace file
    "^production",                // Dependency production files
    { "env": "NODE_ENV" },        // Environment variable
    { "runtime": "node -v" },     // Runtime command
    { "externalDependencies": ["webpack"] }  // Specific dependency version
  ]
}
```

**Outputs** - Files/folders produced by task (stored in cache):
```json
{
  "outputs": [
    "{options.outputPath}",           // From executor options
    "{projectRoot}/dist",             // Project-relative path
    "{workspaceRoot}/coverage/{projectRoot}"  // Workspace-relative path
  ]
}
```

---

## Advanced Level Questions

### Q11: How do you create a custom generator?

**Answer:**

**Step 1: Create Plugin**
```bash
nx g @nx/plugin:plugin my-plugin --directory=tools/my-plugin
nx g @nx/plugin:generator my-generator --project=my-plugin
```

**Step 2: Implement Generator**
```typescript
// tools/my-plugin/src/generators/my-generator/generator.ts
import {
  Tree,
  formatFiles,
  generateFiles,
  names,
  readProjectConfiguration,
  updateProjectConfiguration,
  addDependenciesToPackageJson,
} from '@nx/devkit';
import * as path from 'path';

export interface MyGeneratorSchema {
  name: string;
  directory?: string;
}

export default async function myGenerator(
  tree: Tree,
  options: MyGeneratorSchema
) {
  const normalizedOptions = {
    ...options,
    projectName: names(options.name).fileName,
    className: names(options.name).className,
  };

  // Generate files from templates
  generateFiles(
    tree,
    path.join(__dirname, 'files'),
    `libs/${options.directory || ''}/${normalizedOptions.projectName}`,
    { ...normalizedOptions, template: '' }
  );

  // Add dependencies if needed
  addDependenciesToPackageJson(
    tree,
    { lodash: '^4.17.21' },
    {}
  );

  await formatFiles(tree);

  return () => {
    console.log(`Run: nx test ${normalizedOptions.projectName}`);
  };
}
```

**Step 3: Define Schema**
```json
// schema.json
{
  "$schema": "http://json-schema.org/schema",
  "properties": {
    "name": {
      "type": "string",
      "description": "Library name",
      "x-prompt": "What is the library name?"
    },
    "directory": {
      "type": "string",
      "description": "Directory path"
    }
  },
  "required": ["name"]
}
```

**Step 4: Use Generator**
```bash
nx g @my-org/my-plugin:my-generator awesome-lib
```

---

### Q12: Explain Nx Agents and distributed task execution.

**Answer:**

Nx Agents distribute tasks across multiple machines for faster CI:

```yaml
# .github/workflows/ci.yml
- name: Start Nx Agents
  run: |
    npx nx-cloud start-ci-run \
      --distribute-on="5 linux-medium-js" \
      --stop-agents-after="e2e"

- name: Run tasks
  run: npx nx affected -t lint test build e2e
```

**How It Works:**
1. Main machine orchestrates task distribution
2. Agents receive and execute tasks
3. Results are collected and reported
4. Cache is shared across all agents

**Agent Types:**
- `linux-small-js` - 2 vCPUs, 4GB RAM
- `linux-medium-js` - 4 vCPUs, 8GB RAM
- `linux-large-js` - 8 vCPUs, 16GB RAM

**Dynamic Allocation:**
```yaml
--distribute-on="auto"  # Nx determines optimal agent count
```

---

### Q13: How do you implement Module Federation with Nx?

**Answer:**

```bash
# Generate host and remotes
nx g @nx/react:host shell --remotes=shop,cart
nx g @nx/react:remote shop --host=shell
nx g @nx/react:remote cart --host=shell
```

**Host Configuration:**
```typescript
// apps/shell/module-federation.config.ts
import { ModuleFederationConfig } from '@nx/webpack';

const config: ModuleFederationConfig = {
  name: 'shell',
  remotes: ['shop', 'cart'],
  shared: (libraryName, sharedConfig) => {
    if (['react', 'react-dom', 'react-router-dom'].includes(libraryName)) {
      return { singleton: true, strictVersion: true };
    }
    return sharedConfig;
  },
};

export default config;
```

**Remote Configuration:**
```typescript
// apps/shop/module-federation.config.ts
import { ModuleFederationConfig } from '@nx/webpack';

const config: ModuleFederationConfig = {
  name: 'shop',
  exposes: {
    './Module': './src/remote-entry.ts',
  },
};

export default config;
```

**Running:**
```bash
# Serve all together
nx serve shell

# Serve with specific remotes
nx serve shell --devRemotes=shop
```

---

### Q14: How do you configure Nx Release for package publishing?

**Answer:**

```json
// nx.json
{
  "release": {
    "projects": ["packages/*"],
    "version": {
      "preVersionCommand": "npx nx run-many -t build",
      "conventionalCommits": true
    },
    "changelog": {
      "projectChangelogs": {
        "createRelease": "github"
      },
      "workspaceChangelog": {
        "createRelease": "github"
      }
    },
    "git": {
      "commit": true,
      "commitMessage": "chore(release): publish {version}",
      "tag": true
    }
  }
}
```

**Release Commands:**
```bash
# First release
nx release --first-release

# Regular release
nx release

# Specific version
nx release version 1.2.0

# Dry run
nx release --dry-run

# Publish only
nx release publish
```

---

### Q15: How do you optimize Nx for large workspaces?

**Answer:**

**1. Use Daemon Process:**
```bash
# Ensure daemon is running
nx daemon
```

**2. Optimize Inputs:**
```json
{
  "namedInputs": {
    "production": [
      "default",
      "!{projectRoot}/**/*.spec.ts",
      "!{projectRoot}/**/*.stories.ts"
    ]
  }
}
```

**3. Parallelize Effectively:**
```bash
# Adjust based on machine
nx run-many -t build --parallel=5
```

**4. Use Batch Mode:**
```json
{
  "targets": {
    "build": {
      "executor": "@nx/js:tsc",
      "options": { "batchMode": true }
    }
  }
}
```

**5. Remote Caching:**
```bash
npx nx connect
```

**6. Distributed Execution:**
```yaml
npx nx-cloud start-ci-run --distribute-on="5 linux-medium-js"
```

**7. Selective Testing:**
```bash
# Only test affected
nx affected -t test

# Skip unchanged
nx test my-lib --skip-nx-cache=false
```

---

## Expert Level Questions

### Q16: How do you create a custom Nx plugin with both generators and executors?

**Answer:**

**Complete Plugin Structure:**
```
tools/my-plugin/
├── src/
│   ├── generators/
│   │   ├── library/
│   │   │   ├── generator.ts
│   │   │   ├── schema.json
│   │   │   └── files/
│   │   └── component/
│   │       ├── generator.ts
│   │       └── schema.json
│   ├── executors/
│   │   ├── build/
│   │   │   ├── executor.ts
│   │   │   └── schema.json
│   │   └── deploy/
│   │       ├── executor.ts
│   │       └── schema.json
│   └── index.ts
├── generators.json
├── executors.json
├── package.json
└── project.json
```

**generators.json:**
```json
{
  "generators": {
    "library": {
      "factory": "./src/generators/library/generator",
      "schema": "./src/generators/library/schema.json",
      "description": "Create a library"
    }
  }
}
```

**executors.json:**
```json
{
  "executors": {
    "build": {
      "implementation": "./src/executors/build/executor",
      "schema": "./src/executors/build/schema.json",
      "description": "Build the project"
    }
  }
}
```

**Publishing:**
```json
// package.json
{
  "name": "@my-org/my-plugin",
  "version": "1.0.0",
  "generators": "./generators.json",
  "executors": "./executors.json"
}
```

---

### Q17: How do you implement inferred targets in a custom plugin?

**Answer:**

```typescript
// tools/my-plugin/src/plugins/plugin.ts
import {
  CreateNodesV2,
  CreateNodesContext,
  TargetConfiguration,
} from '@nx/devkit';
import { existsSync } from 'fs';
import { dirname, join } from 'path';

export interface MyPluginOptions {
  buildTargetName?: string;
  testTargetName?: string;
}

export const createNodesV2: CreateNodesV2<MyPluginOptions> = [
  '**/my-config.json',  // Pattern to match
  async (configFiles, options, context) => {
    const results = [];

    for (const configFile of configFiles) {
      const projectRoot = dirname(configFile);
      
      const targets: Record<string, TargetConfiguration> = {};

      // Infer build target
      targets[options?.buildTargetName ?? 'build'] = {
        command: 'my-build-tool build',
        options: {
          cwd: projectRoot,
        },
        cache: true,
        inputs: ['production', '^production'],
        outputs: [`{projectRoot}/dist`],
      };

      // Infer test target if test config exists
      if (existsSync(join(projectRoot, 'test.config.js'))) {
        targets[options?.testTargetName ?? 'test'] = {
          command: 'my-test-tool',
          options: {
            cwd: projectRoot,
          },
          cache: true,
        };
      }

      results.push([
        configFile,
        {
          projects: {
            [projectRoot]: {
              targets,
            },
          },
        },
      ]);
    }

    return results;
  },
];
```

**Register in nx.json:**
```json
{
  "plugins": [
    {
      "plugin": "@my-org/my-plugin",
      "options": {
        "buildTargetName": "build",
        "testTargetName": "test"
      }
    }
  ]
}
```

---

### Q18: How do you design a workspace for a large enterprise?

**Answer:**

**Recommended Structure:**
```
enterprise-workspace/
├── apps/
│   ├── domain-a/
│   │   ├── app-1/
│   │   └── app-2/
│   └── domain-b/
│       └── app-3/
├── libs/
│   ├── shared/              # Cross-domain shared
│   │   ├── ui/
│   │   ├── utils/
│   │   └── types/
│   ├── domain-a/            # Domain-specific
│   │   ├── feature-x/
│   │   ├── data-access/
│   │   └── ui/
│   └── domain-b/
│       ├── feature-y/
│       └── data-access/
├── tools/
│   ├── generators/
│   └── scripts/
└── nx.json
```

**Tagging Strategy:**
```json
// Scope by domain
{ "tags": ["scope:domain-a", "type:feature"] }

// Type hierarchy
{ "tags": ["type:app", "type:feature", "type:ui", "type:data-access", "type:util"] }

// Team ownership
{ "tags": ["team:platform", "team:checkout"] }
```

**Boundary Rules:**
```javascript
depConstraints: [
  // Domain isolation
  { sourceTag: 'scope:domain-a', onlyDependOnLibsWithTags: ['scope:domain-a', 'scope:shared'] },
  { sourceTag: 'scope:domain-b', onlyDependOnLibsWithTags: ['scope:domain-b', 'scope:shared'] },
  
  // Type hierarchy
  { sourceTag: 'type:feature', onlyDependOnLibsWithTags: ['type:ui', 'type:data-access', 'type:util'] },
  { sourceTag: 'type:ui', onlyDependOnLibsWithTags: ['type:util'] },
]
```

**CI Strategy:**
```yaml
# Run affected per domain for parallel pipelines
- name: Domain A
  run: nx affected -t test --projects=tag:scope:domain-a

- name: Domain B  
  run: nx affected -t test --projects=tag:scope:domain-b
```

---

### Q19: How do you migrate an existing project to Nx?

**Answer:**

**Option 1: Add Nx to Existing Repo**
```bash
npx nx@latest init
```

**Option 2: Create Nx Wrapper**
```bash
# Initialize Nx
npx nx@latest init

# Add existing projects
npx nx add @nx/react  # or @nx/angular, @nx/node

# Import existing project
npx nx g @nx/workspace:move --project=existing-app --destination=apps/existing-app
```

**Option 3: Gradual Migration**
```json
// nx.json - Package-based monorepo
{
  "workspaceLayout": {
    "appsDir": "packages",
    "libsDir": "packages"
  }
}
```

**Migration Steps:**
1. Install Nx: `npm add -D nx @nx/workspace`
2. Create nx.json configuration
3. Add project.json to each project (or use package.json)
4. Configure caching and task dependencies
5. Set up module boundaries
6. Enable remote caching

---

### Q20: How do you troubleshoot Nx performance issues?

**Answer:**

**1. Profile Task Execution:**
```bash
# Enable profiling
NX_PROFILE=profile.json nx build my-app

# Analyze profile
npx nx-cloud analyze profile.json
```

**2. Check Cache Hit Rate:**
```bash
# View cache statistics
nx run-many -t build --verbose
```

**3. Analyze Project Graph:**
```bash
# Check for circular dependencies
nx graph

# View task dependencies
nx build my-app --graph
```

**4. Debug Daemon:**
```bash
# Check daemon status
nx daemon

# Restart daemon
nx daemon --stop
nx reset
```

**5. Optimize Configuration:**
```json
// nx.json
{
  "parallel": 3,  // Reduce if memory issues
  "cacheDirectory": ".nx/cache",
  "useDaemonProcess": true,
  "namedInputs": {
    "production": [
      "default",
      "!{projectRoot}/**/*.spec.ts"  // Exclude test files
    ]
  }
}
```

**6. Memory Issues:**
```bash
# Increase Node memory
NODE_OPTIONS="--max-old-space-size=8192" nx build my-app
```

**7. Check Affected Calculation:**
```bash
# Verify affected projects
nx show projects --affected

# Debug affected
NX_VERBOSE_LOGGING=true nx affected -t build
```

---

## Summary

**Key Nx Concepts:**

1. **Workspace** - Monorepo containing apps and libs
2. **Project Graph** - Dependency visualization
3. **Affected** - Run only changed projects
4. **Caching** - Local and remote task caching
5. **Generators** - Code scaffolding
6. **Executors** - Task runners
7. **Module Boundaries** - Architectural constraints
8. **Task Pipeline** - Dependency-based execution
9. **Nx Cloud** - Remote caching and distribution
10. **Inferred Tasks** - Automatic target detection

**Best Practices:**

1. Use consistent library structure
2. Define clear module boundaries with tags
3. Enable caching for all cacheable targets
4. Use affected commands in CI
5. Leverage remote caching for team efficiency
6. Create custom generators for consistency
7. Visualize project graph regularly
8. Keep libraries small and focused
9. Use conventional commits for releases
10. Monitor and optimize CI performance

---

*Document Version: 1.0*
*Last Updated: February 2026*
