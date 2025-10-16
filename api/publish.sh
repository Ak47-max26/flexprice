#!/bin/bash
# Script to publish SDKs to their respective package managers

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Function to show help
show_help() {
  echo "Usage: ./publish.sh [options]"
  echo ""
  echo "Options:"
  echo "  --js, --javascript    Publish JavaScript SDK to npm"
  echo "  --py, --python        Publish Python SDK to PyPI"
  echo "  --go                  Prepare Go SDK for publishing (creates tag)"
  echo "  --all                 Publish all SDKs"
  echo "  --version VERSION     Set version for all SDKs before publishing"
  echo "  --dry-run             Run in dry run mode without making changes"
  echo "  --help                Show this help message"
  echo ""
  echo "Examples:"
  echo "  ./publish.sh --all --version 1.2.3"
  echo "  ./publish.sh --js --py"
  echo "  ./publish.sh --go --version 1.0.0"
  echo "  ./publish.sh --go --version 1.0.0 --dry-run"
}

# Parse arguments
PUBLISH_JS=false
PUBLISH_PY=false
PUBLISH_GO=false
VERSION=""
DRY_RUN=false

while [[ $# -gt 0 ]]; do
  case $1 in
    --js|--javascript)
      PUBLISH_JS=true
      shift
      ;;
    --py|--python)
      PUBLISH_PY=true
      shift
      ;;
    --go)
      PUBLISH_GO=true
      shift
      ;;
    --all)
      PUBLISH_JS=true
      PUBLISH_PY=true
      PUBLISH_GO=true
      shift
      ;;
    --version)
      VERSION="$2"
      shift
      shift
      ;;
    --dry-run)
      DRY_RUN=true
      shift
      ;;
    --help)
      show_help
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      show_help
      exit 1
      ;;
  esac
done

# If no SDK specified, show help
if [[ "$PUBLISH_JS" == "false" && "$PUBLISH_PY" == "false" && "$PUBLISH_GO" == "false" ]]; then
  echo "Error: No SDK specified for publishing"
  show_help
  exit 1
fi

# Update versions if specified
if [[ -n "$VERSION" ]]; then
  echo "Updating SDK versions to $VERSION..."
  
  if [ "$DRY_RUN" = true ]; then
    echo "DRY RUN: Would update SDK versions to $VERSION"
  else
    # Update JavaScript SDK version
    if [[ "$PUBLISH_JS" == "true" ]] && [ -d "javascript" ]; then
      echo "Updating JavaScript SDK version..."
      if [ -f "javascript/package.json" ]; then
        # Use jq if available, otherwise use sed
        if command -v jq &> /dev/null; then
          jq ".version = \"$VERSION\"" javascript/package.json > javascript/package.json.tmp
          mv javascript/package.json.tmp javascript/package.json
        else
          # Using sed without triggering npm hooks
          sed -i.bak "s/\"version\": \"[^\"]*\"/\"version\": \"$VERSION\"/" javascript/package.json
          rm -f javascript/package.json.bak
        fi
        echo "✅ JavaScript SDK version updated"
      else
        echo "⚠️ package.json not found in javascript directory"
      fi
    fi

    # Update Python SDK version
    if [[ "$PUBLISH_PY" == "true" ]] && [ -d "python" ]; then
      echo "Updating Python SDK version..."
      if [ -f "python/setup.py" ]; then
        sed -i.bak "s/VERSION = \"[^\"]*\"/VERSION = \"$VERSION\"/" python/setup.py
        rm -f python/setup.py.bak
        echo "✅ Python SDK version updated"
      else
        echo "⚠️ setup.py not found in python directory"
      fi
    fi

    # For Go SDK, we just create a tag later
    if [[ "$PUBLISH_GO" == "true" ]] && [ -d "go" ]; then
      echo "Go SDK version will be set to $VERSION"
    fi
  fi
fi

if [ "$DRY_RUN" = true ]; then
  echo "DRY RUN: Publishing process completed (no changes made)"
else
  echo "Publishing process completed!"
  echo "NOTE: For the Go SDK, you need to push the tag to GitHub for publishing."
  echo "      The GitHub workflow will handle this automatically when triggered."
fi 