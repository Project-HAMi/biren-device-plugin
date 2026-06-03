# Biren Device Plugin

## Prerequisites

The list of prerequisites for running the Biren device plugin is described below:

1.  Biren GPU Driver >= 1.2.2
2.  Kubernetes >=1.13

## Deployment

### Label the Node with `biren=on`
```bash
kubectl label node {biren-node} biren=on
```

### Deploy `biren-device-plugin`


```bash
kubectl apply -f deploy/biren-device-plugin.yaml
```

### Usage

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod1
spec:
  restartPolicy: OnFailure
  containers:
    - image: ubuntu
      name: pod1-ctr
      command: ["sleep"]
      args: ["infinity"]
      resources:
        limits:
          birentech.com/gpu: 1
```