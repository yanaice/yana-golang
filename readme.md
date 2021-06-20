# Project: Yana Study :)
## Architecture
- micro services
- google cloud k8s

## Project folder structure
- Business logic is in folder service
- Business logic handler all response error
- Interface name = folder name

## Code pattern
- hexagonal concept : port & adapter : talking to each other via interface (port)

## Common function / standard uses
- error handler
- logs
- middleware

## Config local

## Run local
- go run .

## Mock / Test local
- mockery
- go run test...

## Deploy
- before deploy (go mod vendor, sync common etc.)
- Git flow
- Code Review
- CI/CD
- project config, secret on cloud (each env)
- other config ex. ingress

## Logging / Monitoring on ENV
- vpn
- LENS config k8s
- how to debug, view logs errors