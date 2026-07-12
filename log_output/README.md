# Log output app

Build docker image with `docker build -t log-output-app .`

Import image into k3d with `k3d image import log-output-app`

The port the server listens on can be changed by editing the
PORT environment variable inside `./manifests/deployment.yaml`
and the targetPort field in `./manifests/service.yaml`

Deploy with `kubectl apply -f manifests`
