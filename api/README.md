# FlexPrice API SDKs

This directory contains the generated SDKs for the Flexprice API. The SDKs are generated from the OpenAPI specification in `/docs/swagger/swagger-3-0.json`.

## SDK Generation

The SDKs are generated using the OpenAPI Generator. To generate the SDKs, run:

```bash
make generate-sdk
```

This will generate the following SDKs:

- Go SDK in `api/go/`
- Python SDK in `api/python/`
- JavaScript SDK in `api/javascript/`

## Test Files

Each SDK includes test files in a `test` directory, which serve as:

1. Functional tests for the SDK
2. Examples of how to use the SDK

The test files are preserved when regenerating the SDKs through the GitHub Actions workflow, which backs up test files before generation and restores them afterward.

## SDK Publishing

The SDKs are published to their respective repositories:

- Go SDK: [github.com/flexprice/go-sdk](https://github.com/flexprice/go-sdk)
- Python SDK: [github.com/flexprice/python-sdk](https://github.com/flexprice/python-sdk)
- JavaScript SDK: [github.com/flexprice/javascript-sdk](https://github.com/flexprice/javascript-sdk)

### Publishing Process

1. The SDKs are generated from the OpenAPI specification
2. Test files are preserved during generation
3. Dependencies are updated and tests are run
4. The SDKs are published to their respective repositories

The publishing process is automated via GitHub Actions and can be triggered by:

- Pushing changes to the OpenAPI specification
- Manually running the workflow with a version number

## Local Development

When developing locally, it's recommended to:

1. Generate the SDKs using `make generate-sdk`
2. Run tests to ensure functionality
3. Test your changes locally before pushing

To avoid committing generated code to the main repository, the SDK directories are included in `.gitignore` but exclude the test directories.

## Available SDKs

- **Go**: `api/go`
- **Python**: `api/python`
- **JavaScript**: `api/javascript`

## Generating SDKs

The SDKs are generated from the OpenAPI 3.0 specification located at `docs/swagger/swagger-3-0.json` using the OpenAPI Generator CLI.

To generate all SDKs, run:

```bash
make generate-sdk
```

To generate a specific SDK, run one of the following commands:

```bash
make generate-go-sdk
make generate-python-sdk
make generate-javascript-sdk
```

## SDK Usage

### Go SDK

```go
import (
    "context"
    flexprice "github.com/your-org/flexprice/api/go"
)

func main() {
    cfg := flexprice.NewConfiguration()
    cfg.Host = "your-api-host"
    client := flexprice.NewAPIClient(cfg)
    
    // Use the client to make API calls
    // ...
}
```

### Python SDK

```python
import flexprice
from flexprice.api_client import ApiClient
from flexprice.configuration import Configuration

# Configure API client
configuration = Configuration(host="your-api-host")
api_client = ApiClient(configuration)

# Use the client to make API calls
# ...
```

### JavaScript SDK

```javascript
import * as flexprice from 'flexprice';

// Configure API client
const apiClient = new flexprice.ApiClient("your-api-host");

// Use the client to make API calls
// ...
```

## Customization

The SDK generation process can be customized by modifying the OpenAPI Generator configuration in the Makefile.

## CI/CD Integration

These SDKs are intended to be generated as part of the CI/CD pipeline. The Makefile targets can be integrated into your CI/CD workflow to automatically generate and publish the SDKs when the API specification changes.

## Publishing SDKs

The SDKs can be published to their respective package managers using the CI/CD workflow defined in `api/ci-cd-example.yml`. Here's how each SDK is published:

### JavaScript SDK

The JavaScript SDK is published to npm as the `flexprice` package. To install it:

```bash
npm install flexprice
```

### Python SDK

The Python SDK is published to PyPI as the `flexprice` package. To install it:

```bash
pip install flexprice
```

### Go SDK

The Go SDK is published as a Go module on GitHub. To use it in your Go project:

```go
import "github.com/your-org/flexprice/api/go"
```

And add it to your dependencies:

```bash
go get github.com/your-org/flexprice/api/go
```

### Manual Publishing

If you need to publish the SDKs manually, follow these steps:

#### JavaScript SDK (npm)

```bash
cd api/javascript
# Update version in package.json if needed
npm publish --access public
```

#### Python SDK (PyPI)

```bash
cd api/python
pip install build twine
python -m build
python -m twine upload dist/*
```

#### Go SDK (GitHub)

```bash
cd api/go
# Ensure go.mod exists with correct module path
go mod init github.com/your-org/flexprice/api/go
go mod tidy

# Tag a new version
git tag -a "go-sdk/v1.0.0" -m "Go SDK release v1.0.0"
git push origin go-sdk/v1.0.0

# Update Go module proxy
GOPROXY=proxy.golang.org go list -m github.com/your-org/flexprice/api/go@v1.0.0
```

## Required Secrets for CI/CD

To enable automatic publishing in the CI/CD workflow, you need to set up the following secrets in your GitHub repository:

- `NPM_TOKEN`: Access token for publishing to npm
- `PYPI_USERNAME`: Your PyPI username
- `PYPI_PASSWORD`: Your PyPI password or API token

For more information on setting up these secrets, refer to:
- [npm: Creating and viewing access tokens](https://docs.npmjs.com/creating-and-viewing-access-tokens)
- [PyPI: Creating API tokens](https://pypi.org/help/#apitoken) 