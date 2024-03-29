# echoserver [![docker image version](https://img.shields.io/docker/v/lapwingcloud/echoserver/latest?logo=docker&color=blue)](https://hub.docker.com/r/lapwingcloud/echoserver) [![github build status](https://img.shields.io/github/actions/workflow/status/lapwingcloud/echoserver/push-img-docker-hub.yaml?logo=github)](https://github.com/lapwingcloud/echoserver/actions)


A grpc and http echo server returning connection and server information for debugging purpose. :ping_pong:

**Features**

- supports grpc and http
- supports specifying a delay in the request
- rich information in the response, e.g.
  - version: useful in testing deployment roll out
  - server hostname: useful in testing load balancing
  - client ip and port
  - request host / authority header etc
- server request logs
- server graceful shutdown

## Usage

Use docker to run the echo server, by default it listens on ports

- `8080`: http
- `9090`: grpc

```
docker run -p 8080:8080 -p 9090:9090 lapwingcloud/echoserver
```

### request params

For both grpc and http it supports 2 parameters in the request. Note for http it needs to be passed in the request body as json.

- `delaySeconds`: if specified, the request will sleep the number of `delaySeconds` before returning the response
- `payload`: if specified, the server will add the same field with the same value in the response

### example http request

```
$ curl -d '{"delaySeconds":0.1, "payload":"hello"}' localhost:8080
{
  "timestamp": "2024-02-04T01:57:36.115139978Z",
  "version": "dev",
  "hostname": "02f3d12d7951",
  "remoteIp": "172.17.0.1",
  "remotePort": 47082,
  "requestId": "7746d5b4-6ce8-4d7f-9a2c-833bbcffe9f9",
  "requestHost": "localhost:8080",
  "requestMethod": "POST",
  "requestPath": "/",
  "requestQuery": "",
  "requestTime": 0.100276072,
  "userAgent": "curl/7.81.0",
  "delaySeconds": 0.1,
  "payload": "hello"
}
```

### example grpc request

```
$ grpcurl -plaintext -d '{"delaySeconds":0.1, "payload":"hello"}' localhost:9090 echo.Echo/Ping
{
  "timestamp": "2024-02-04T02:06:52.025236095Z",
  "hostname": "02f3d12d7951",
  "version": "dev",
  "remoteIp": "172.17.0.1",
  "remotePort": 35962,
  "requestId": "d0baf289-3ddd-4051-b67c-36df9549870f",
  "authority": "localhost:9090",
  "requestMethod": "/echo.Echo/Ping",
  "requestTime": 0.100468275,
  "userAgent": "grpcurl/dev-build (no version set) grpc-go/1.57.0",
  "delaySeconds": 0.1,
  "payload": "hello"
}
```

### cli flags

```
$ docker run lapwingcloud/echoserver echoserver -h
Usage of echoserver:
  -grpc-bind string
        The grpc server listen address (default ":9090")
  -http-bind string
        The http server listen address (default ":8080")
  -log-format string
        The log format (text, json) (default "json")
```

## Example

### Deploy to Kubernetes

Here is an [example kubernetes manifest](example/deploy-k8s.yaml) that you can deploy to any cluster.

**prerequisites**

- the kubernetes cluster has an ingress controller  (e.g. [ingress-nginx](https://kubernetes.github.io/ingress-nginx/))
- (optional) you already have a valid TLS certificate for the domain you want the echoserver to serve on

**setup**

First you need to replace the domain name in the example/deploy-k8s.yaml with your own domain.

Then run the `kubectl` commands below to create a namespace and create the k8s deployment, service, ingress objects.

```
kubectl create echoserver
kubectl -n echoserver apply -f example/deploy-k8s.yaml
```

Optionally, if you have a TLS certificate and the private key, create the TLS secret in kubernetes so the ingress controller can use it to serve https and grpc over TLS.

Note you can obtain free TLS certificates via tools like [certbot](https://certbot.eff.org/) or [cert-manager](https://cert-manager.io/) using [letsencrypt](https://letsencrypt.org/).

```
kubectl -n echoserver create secret tls echoserver-tls --cert fullchain.pem --key privkey.pem
```

Now the echoserver should be accessible from the domain you setup, e.g.

```
curl https://echoserver-http.lapwingcloud.com
grpcurl echoserver-grpc.lapwingcloud.com:443 echo.Echo/Ping
```

If you didn't setup the TLS, just add the insecure flags to ignore invalid certificate.

```
curl --insecure https://echoserver-http.lapwingcloud.com
grpcurl -insecure echoserver-grpc.lapwingcloud.com:443 echo.Echo/Ping
```

Note, nginx doesn't support grpc over plaintext, so you cannot use `grpcurl -plaintext` to access the ingress domain.
