#### Docker란?
* 도커는 컨테이너 기반의 오픈소스 가상화 플랫폼이다.
* 도커는 컨테이너를 이미지 파일로 빌드하고 배포하여 어디서나 실행할 수 있도록 해주는 오픈소스이다.
* 컨테이너를 git 에 저장된 소스처럼 build/push/pull 할 수 있는 방법을 먼저 제공하면서 주목 받았다.

#### 컨테이너
* 컨테이너는 어플리케이션의 런타임 인스턴스이다.
* 컨테이너 단위로 OS, 라이브러리, 어플리케이션을 패키징 할 수 있다.
* 컨테이너는 OS 에 여러 어플리케이션을 독립적으로 실행할 수 있도록 해준다.
* 즉, 컨테이너는 여러 어플리케이션의 격리된 환경을 지원하는 가상화 기술로 볼 수 있다.

#### 컨테이너 vs VM
* 컨테이너는 하나의 OS 위에서 여러개가 실행된다.
* 각각의 컨테이너는 사용자 영역에서 격리된다.
* 컨테이너는 VM 보다 가볍고 빠르다.

#### 컨테이너를 도입하면서 기대할 수 있는 점
* 서비스에 필요한 OS, 라이브러리, 어플리케이션을 컨테이너로 묶을 수 있어서 업데이트나 롤백 시 안정적이다.
* 컨테이너는 가볍고 빠르기 때문에 서비스 배포 시 빠르게 반영할 수 있다.
* 컨테이너 실행 환경만 갖춰져 있으면 어디서나 실행되기 때문에 신속한 서비스 스케일링이 가능하다.
* 단일 호스트의 장애를 컨테이너의 이동으로 대응하여 보다 빠르게 복구할 수 있다.
* 서비스와 운영을 분리할 수 있어서 컨테이너 실행 환경을 더이상 신경쓰지 않아도 된다.

#### 도커 동작방식
* 도커 이미지를 빌드하기 위해 Dockerfile 을 작성한다.
* Dockerfile 과 함께 도커 빌드 요청을 보낸다.
* 도커 서버에서는 도커 이미지를 빌드하여 로컬 저장소에 저장한다.
* 도커 push 명령을 받으면 도커 서버는 로컬의 도커 이미지를 도커 레지스트리에 올린다.
* 도커 run 명령을 배포할 도커 서버에 전송한다.
* 도커 run 명령을 받은 도커 서버는 도커 레지스트리에 이미지를 로컬 저장소로 다운 받는다.
* 도커 이미지를 이용하여 컨테이너를 시작한다.



# 컨테이너 생성 및 실행
- -i -t 옵션은 컨테이너와 상호 입출력을 가능하게 함
- -d (백그라운드 모드) 옵션을 사용하려면 컨테이너에 포그라운드로 프로그램이 돌아야함
- --link 옵션은 내부 IP를 알 필요 없이 항상 컨테이너에 별명으로 접근하도록 설정
```
# 생성 및 실행
$ docker run -i -t ubuntu:14.04

# 생성 및 실행 (shift+ctrl+q 를 누르면 컨테이너 동작된 상태로 나옴)
$ docker create -i -t --name mycentos centos:7
$ docker start mycentos
$ docker attach mycentos

# 백그라운드로 실행 및 /bin/bash 접근
$ docker run -d --name wordpressdb mysql:5.7
$ docker exec -it workspressdb /bin/bash


$ --link 옵션 사용: wordpressdb 컨테이너를 mysql라는 이름으로 설정
$ docker run -d --name wordpress --link wordpressdb:mysql -p 80 wordpress
```

# 도커 컨테이너 or 이미지 정보 확인
```
$ docker container ls
$ docker ps -a
$ docker images
```

# pull image
```
$ docker pull centos:7
```

# 도커 컨테이너 이름 변경
```
$ docker rename mycentos tucentos
```

# 도커 컨테이너 삭제
```
$ docker rm mycentos

# 전체 삭제
$ docker rm $(docker ps -a -q)

or


$ docker container prune
```

# 외부에 포트 노출
```
$ docker run -i -t --name mywebserver -p 80:80 ubuntu:14.04

# 포트 정보 확인
$ docker port mywebserver
```

# 볼륨
```
# 호스트와 컨테이너 간 볼륨 공유
$ docker run -i -t --name file_volume -v /Users/user:/usr/local/hello ubuntu

# 이미 호스트와 볼륨 공유 중인 컨테이너를 통해 볼륨 공유
$ docker run -i -t --name volumes_from_container --volumes-from file_volume

# 도커 볼륨
$ docker volume create --name myvolume
$ docker volume ls
$ docker run -i -t --name myvolume_1 -v myvolume:/root/ ubuntu:14.04
$ docker run -i -t --name myvolume_2 -v myvolume:/root/ ubuntu:14.04

# inspect 명령
$ docker inspect --type volume myvolume
```

# 도커 네트워크
- 브리지 네트워크 : NAT 로 IP 생성해서 브리지를 통해 호스트와 연결, 기본
- 호스트 네트워크 : 네트워크를 호스트로 설정하면 네트워크 환경 그대로 사용 가능
- 논 네트워크 : 아무런 네트워크를 쓰지 않음
- 컨테이너 네트워크 : 다른 컨테이너의 네트워크 네임스페이스 환경 공유
```
$ docker network ls


# 호스트 네트워크
$ docker run -i -t --name network_host --net host ubuntu:14.04


# 논 네트워크
$ docker run -i -t --name hello_none --net none ubuntu:14.04


# 컨테이너 네트워크
$ docker run -i -t -d --name network_1 ubuntu:14.04
$ docker run -i -t -d --name network_2 --net container:network_1 ubuntu:14.04
$ docker exec network_1 ifconfig
$ docker exec network_1 ifconfig
```

# 도커 이미지 생성 및 추출
- 도커 이미지 별 공통 layer는 공유한다.
```
$ docker commit -a "tuyy" -m "commit message" commit_test commit_test:first

# save tar
$ docker save -o ubuntu_14_04.tar ubuntu:14.04

# load tar
$ docker load -i ubuntu_14_04.tar

# tag 명 수정 및 upload dockhup
$ docker tag commit_test:second whwkfyd91/dev-tuyy:0.0
$ docker login
$ docker push whwkfyd91/dev-tuyy:0.0
$ docker pull whwkfyd91/dev-tuyy:0.0
```

# Dockerfile 예시
- 컨테이너에서 수행할 작업을 명시
- Dockerfile을 읽어 들일 때 기본적으로 현재 디렉터리에 있는 Dockerfile 이라는 이름을 가진 파일을 선택
- FROM, RUN, ADD, WORKDIR, CMD, EXPOSE
```
FROM ubuntu:14.04
MAINTAINER whwkfyd91
LABEL "purpose"="practice"
RUN apt-get update
RUN apt-get install apache2 -y
ADD test.html /var/www/html
WORKDIR /var/www/html
RUN ["/bin/bash", "-c", "echo hello >> test2.html"]
EXPOSE 80
CMD apachectl -DFOREGROUND
```

# Dockerfile 빌드 및 실행
- .dockerignore : 이 파일에 명시된 이름의 파일을 컨텍스트에서 제외
```
# 이미지 생성
$ docker build -t mybuild:0.0 .
$ docker build --no-cache -t mybuild:0.0 .


# 실행 : -P 옵션은 EXPOSE 로 지정한 포트로 실행을 의미
$ docker run -d -P --name myserver mybuild:0.0


# LABEL 로 지정한 이미지 검색
$ docker images --filter "label=purpose=practice"
```
