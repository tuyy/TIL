### k8s 예제
* https://kubernetes.io/ko/docs/reference/kubectl/cheatsheet/

### 기본 명령어
```
$ kubectl get pods
$ kubectl describe pods ${NAME}
$ kubectl create -f xxxx.yaml
$ kubectl apply -f xxxx.yaml
$ kubectl delete -f xxxx.yaml
```


### centos7 이미지를 run하고 shell 접속
```
$ kubectl run -i -t shell-test --image=base.registry.navercorp.com/centos:centos7
$ exit
```


### 이미 떠있는 pods에 접속
```
$ kubectl exec -it ${NAME} /bin/bash
```


### stdout 로그 확인
```
$ kubectl logs -f ${NAME}
```

.. TODO

