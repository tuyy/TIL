## pm2 란?
* nodejs에서 사용하는 process manager 이다.
* nodejs로 웹 어플리케이션을 개발할때 watching을 걸어두면 코드가 수정되면 바로 데몬을 다시 띄어준다.

### 설치
```
npm install -g pm2
```

### 사용법
```
pm2 start ~.js --watch // ~로 이름이 지정된 프로세스가 실행된다.
pm2 list
pm2 stop ~.js
pm2 restart ~.js
pm2 delete ~.js
pm2 show

pm2 logs
pm2 flush // 로그 삭제

pm2 start ~.js -i 0(or 실행할 데몬 수) --name myname // cluster로 실행
pm2 reload // 클러스터로 구성된 경우 바로 실행 

pm2 monit // 모니터링

pm2 start config.json  // 설정파일을 만들어서 실행도 가능하다.
```
