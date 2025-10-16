# Contributing to Flexprice

Thank you for your interest in Flexprice and for taking the time to contribute to this project. 🙌 Flexprice is a project by developers for developers and there are a lot of ways you can contribute.


## Code of conduct

Contributors are expected to adhere to the [Code of Conduct](CODE_OF_CONDUCT.md).

## Prerequisites for the contributors

Contributors should have knowledge of git, go, and markdown for most projects since the project work heavily depends on them.
We encourage Contributors to set up Flexprice for local development and play around with the code and tests to get more comfortable with the project. 

Sections

- <a name="contributing"> General Contribution Flow</a>
  - <a name="#commit-signing">Developer Certificate of Origin</a>
- <a name="contributing-Flexprice">Flexprice Contribution Flow</a>
  - <a name="Flexprice-server">Flexprice Server</a>
  - <a name="Flexprice-docs">Flexprice Documentation</a>

# <a name="contributing">General Contribution Flow</a>

## <a name="commit-signing">Signing-off on Commits (Developer Certificate of Origin)</a>

To contribute to this project, you must agree to the Developer Certificate of
Origin (DCO) for each commit you make. The DCO is a simple statement that you,
as a contributor, have the legal right to make the contribution.

See the [DCO](https://developercertificate.org) file for the full text of what you must agree to
and how it works [here](https://github.com/probot/dco#how-it-works).
To signify that you agree to the DCO for contributions, you simply add a line to each of your
git commit messages:

```
Signed-off-by: Jane Smith <jane.smith@example.com>
```

In most cases, you can add this signoff to your commit automatically with the
`-s` or `--signoff` flag to `git commit`. You must use your real name and a reachable email
address (sorry, no pseudonyms or anonymous contributions). An example of signing off on a commit:

```
$ commit -s -m “my commit message w/signoff”
```

To ensure all your commits are signed, you may choose to add this alias to your global `.gitconfig`:

_~/.gitconfig_

```
[alias]
  amend = commit -s --amend
  cm = commit -s -m
  commit = commit -s
```

# FlexPrice Service Architecture

## Project Structure

```
internal/
├── domain/
│   ├── events/
│   │   ├── model.go         # Core event domain model
│   │   ├── repository.go    # Repository interface
│   └── meter/
│       ├── model.go         # Core meter domain model
│       ├── repository.go    # Repository interface
├── repository/
│   ├── clickhouse/
│   │   └── event.go
│   ├── postgres/
│   │   └── meter.go
|   └── factory.go          # Factory for creating repositories
├── service/
│   ├── event.go            # Event service implementation
│   └── meter.go            # Meter service implementation
├── api/
│   ├── v1/
│   │   ├── events.go       # Event API implementation
│   │   └── meter.go        # Meter API implementation
│   ├── dto/
│   │   ├── event.go
│   │   └── meter.go
│   └── router.go           # API router implementation
└── cmd/server/
    └── main.go             # Server application entry point
└── docs/
    └── ...                 # Documentation files
    ├── swagger
    │   ├── swagger.yaml    # Generated Swagger API specifications
```


## Layer Responsibilities

### Domain Layer
- Contains core business logic and domain models
- Defines interfaces for repositories
- No dependencies on external packages or other layers

### Repository Layer
- Implements data access interfaces defined in domain
- Handles database operations

### Service Layer
- Orchestrates business operations
- Implements business logic
- Uses repository interfaces for data access
- Handles cross-cutting concerns

### API Layer
- Handles HTTP requests/responses
- Converts between DTOs and domain models
- No business logic, only request validation and response formatting

### Key Design Principles

1. **Dependency Rule**: Dependencies only point inward. Domain layer has no outward dependencies.

2. **Interface Segregation**: Repository interfaces are defined in domain layer but implemented in repository layer.

3. **Dependency Injection**: Using [fx](https://github.com/uber-go/fx) for clean dependency management.

4. **Separation of Concerns**: Each layer has a specific responsibility.

### Example Flow

For an event ingestion:
1. API Layer (`/api/v1/events.go`) receives HTTP request
2. Converts DTO to domain model
3. Calls service layer
4. Service layer (`/service/event_service.go`) handles business logic
5. Repository layer persists data
6. Response flows back through the layers

### Adding New Features

1. Define domain models and interfaces in domain layer
2. Implement repository interfaces if needed
3. Add service layer logic
4. Create API handlers and DTOs
5. Register routes and dependencies

## Testing

We use [testutil](./internal/testutil) package to create test data and setup test environment.

### Testing Guidelines

1. **Unit Tests**
   - Each package should have corresponding `*_test.go` files
   - Use table-driven tests when testing multiple scenarios
   - Mock external dependencies using interfaces
   - Aim for >80% test coverage for critical paths

2. **Integration Tests**
   - Place integration tests in a separate `integration` package
   - Use test containers for database tests
   - Integration tests should use the same config structure as production

3. **Running Tests**

Use the following make commands:

```
make test # Run all tests
make test-verbose # Run all tests with verbose output
make test-coverage # Run all tests and generate coverage report
```


# How to contribute ?

We encourage contributions from the community.

**Create a [GitHub issue](https://github.com/flexprice/flexprice/issues) for any changes beyond typos and small fixes.**

We review GitHub issues and PRs on a regular schedule.

To ensure that each change is relevant and properly peer reviewed, please adhere to best practices for open-source contributions.
This means that if you are outside the Flexprice organization, you must fork the repository and create PRs from branches on your own fork.
The README in GitHub's [first-contributions repo](https://github.com/firstcontributions/first-contributions) provides an example.


## <a name="contributing-Flexprice">Flexprice Contribution Flow</a>

Flexprice is written in `Go` (Golang) and leverages Go Modules. Relevant coding style guidelines are the [Go Code Review Comments](https://code.google.com/p/go-wiki/wiki/CodeReviewComments) and the _Formatting and style_ section of Peter Bourgon's [Go: Best
Practices for Production Environments](https://peter.bourgon.org/go-in-production/#formatting-and-style).

There are many ways in which you can contribute to Flexprice.

###  <a name="Flexprice-server">Flexprice Server</a>

#### Report a Bug
Report all issues through GitHub Issues using the [Report a Bug](https://github.com/flexprice/flexprice/issues/new?assignees=&labels=&template=bug_report.md&title=) template.
To help resolve your issue as quickly as possible, read the template and provide all the requested information.

#### Feature request
We welcome all feature requests, whether it's to add new functionality to an existing extension or to offer an idea for a brand new extension.
File your feature request through GitHub Issues using the [Feature Request](https://github.com/flexprice/flexprice/issues/new?assignees=&labels=&template=feature_request.md&title=) template.

#### Close a Bug
We welcome contributions that help make Flexprice bug free & improve the experience of our users. You can also find issues tagged [Good First Issues](https://github.com/flexprice/flexprice/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22).
