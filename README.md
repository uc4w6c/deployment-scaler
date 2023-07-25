# deployment-scaler

## 使い方
```shell
$ docker build -t deployment-scaler:1.0.0 ./
$ kind load docker-image deployment-scaler:1.0.0 --name kind-service
$ k apply -f ./examples
```
