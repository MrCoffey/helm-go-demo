### Build and push the image

```bash
docker build -t demo:test .

docker tag demo:test coffey0container/demo:latest

docker push coffey0container/demo:latest
```


### Example on Kubernetes

1.) Generate the pod file

```bash
kubectl run demo --image=coffey0container/demo:latest --generator=run-pod/v1 --dry-run -o yaml > pod.yaml
```

2.) Create the pod

```bash
kubectl create -f pod.yaml
```

3.) Expose the pod:

```bash
kubectl expose pod/demo --name=demo-service
```

4.) Port forward the service:

```bash
k port-forward service/demo-service 8000:8000
```
