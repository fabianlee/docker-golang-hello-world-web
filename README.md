# GoLang HTTP web server

Golang http web server running by default on port 8080 that is intended for testing.

Kubernetes compatible health check available at "/healthz".

Final image is based on busybox:1.32.1-glibc totaling ~11Mb because it takes advantage of multi-stage building.

I have this same project, but enhanced with Prometheus metrics at [docker-prom-golang-hello-world-web](https://github.com/fabianlee/docker-prom-golang-hello-world-web).

# Pulling image from GitHub Container Registry

```
docker pull ghcr.io/fabianlee/docker-golang-hello-world-web:latest
```

# Environment variables available to image

* GREETING - message displayed, defaults to "World"
* PORT - listen port, defaults to 8080
* APP_CONTEXT - base context path of app, defaults to '/'

# Environment variables populated from Downward API
* MY_NODE_NAME - name of k8s node
* MY_POD_NAME - name of k8s pod
* MY_POD_IP - k8s pod IP
* MY_POD_SERVICE_ACCOUNT - service account of k8s pod

# Prerequisites for local build of image
* make utility (sudo apt-get install make)

# Makefile targets for local build
* docker-build (builds image)
* docker-run-fg (runs container in foreground, ctrl-C to exit)
* docker-run-bg (runs container in background)
* k8s-apply (applies deployment to kubernetes cluster)
* k8s-delete (removes deployment on kubernetes cluster)

# Creating tag that invokes Github Action

```
newtag=v1.0.1
git commit -a -m "changes for new tag $newtag" && git push -o ci.skip
git tag $newtag && git push origin $newtag
```

# Deleting tag

```
# delete local tag, then remote
todel=v1.0.1
git tag -d $todel && git push -d origin $todel
```

