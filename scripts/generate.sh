#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/graphql-server/gqlgen/generated.go \
    internal/graphql-server/gqlgen/models/generated.go \
    internal/graphql-server/gqlgen/resolvers/generated/resolver.go
time go run -v github.com/99designs/gqlgen "$1"
printf "\nDone.\n\n"
