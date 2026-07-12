# Todo App

Build docker image with `docker build -t todoapp .`

Import image into k3d with `k3d image import todoapp`

The port the server listens on can be changed by editing the
PORT environment variable inside `./manifests/deployment.yaml`
and the targetPort field in `./manifests/service.yaml`

Deploy app with `kubectl apply -f manifests
