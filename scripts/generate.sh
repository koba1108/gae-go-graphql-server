#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f gql/generated.go \
    gql/models/generated.go \
    gql/resolvers/generated/generated.go
time go run -v github.com/99designs/gqlgen "$1"
printf "\nDone.\n\n"
