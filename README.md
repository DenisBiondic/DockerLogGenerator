# Docker Log Generator

This is a simple containarized Go application which outputs several log messages (in different formats) every couple of seconds. Can be used to test out different logging systems which work with Docker way of logging to stdout / stderr.

# Running locally

`docker build -t log-generator .`

`docker run -it --rm log-generator`

# Deploy to Kubernetes

`kubectl apply -f deploy.yaml`
