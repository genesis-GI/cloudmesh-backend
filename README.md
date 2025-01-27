# Cloudmesh-backend
This is the backend for cloudmesh. Websites, acccount system etc.

[![CodeQL Advanced analysis](https://github.com/genesis-GI/cloudmesh-backend/actions/workflows/codeql.yml/badge.svg)](https://github.com/genesis-GI/cloudmesh-backend/actions/workflows/codeql.yml)

[![Building docker image](https://github.com/genesis-GI/cloudmesh-backend/actions/workflows/docker-image.yml/badge.svg)](https://github.com/genesis-GI/cloudmesh-backend/actions/workflows/docker-image.yml)

[![Compiling project](https://github.com/genesis-GI/cloudmesh-backend/actions/workflows/go.yml/badge.svg)](https://github.com/genesis-GI/cloudmesh-backend/actions/workflows/go.yml)

## Run it locally
go run . release -> Runs the project in release mode
go run . debug -> Runs in debug mode
go run . -> Runs in debug mode

## This is not working currently, but will soon: 
To run it locally with docker:
- Make sure you have git and docker installed.
- Clone the repository
- In the directory of the repository, run "docker-compose up --build -d"