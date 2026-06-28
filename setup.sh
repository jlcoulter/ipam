#!/usr/bin/env bash
# setup.sh — rename module and clean up template markers
# Usage: ./setup.sh my-service github.com/you/my-service

set -euo pipefail

if [ $# -lt 2 ]; then
  echo "Usage: $0 <service-name> <module-path>"
  echo "Example: $0 healthpinger github.com/jlcoulter/healthpinger"
  exit 1
fi

NAME="$1"
MODULE="$2"
OLD_MODULE="github.com/jlcoulter/go-api-template"

# Replace module path in all Go files
find . -name "*.go" -exec sed -i "s|${OLD_MODULE}|${MODULE}|g" {} +

# Replace in go.mod
sed -i "s|${OLD_MODULE}|${MODULE}|g" go.mod

# Replace in README
sed -i "s|go-api-template|${NAME}|g" README.md

# Remove setup script
rm -- "$0"

echo "Template configured: name=${NAME}, module=${MODULE}"
echo "Next steps:"
echo "  1. Run: go mod tidy"
echo "  2. Run: make test"
echo "  3. Delete internal/handler/example.go and replace with your domain"