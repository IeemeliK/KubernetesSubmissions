# Log output app

Build docker image with `docker build -t log-output-app .`

Import image into k3d with `k3d image import log-output-app`

Deploy with `kubectl apply -f manifests/deployment.yaml`
