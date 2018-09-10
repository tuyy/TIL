### tar
* GNU ‘tar’ saves many files together into a single tape or disk archive, and can restore individual files from the archive.

#### option
* c : 압축
* z : tar 압축후 gzip압축
* j : tar 압축후 bzip2압축
* v : verbose 압축과정을 출력
* p : 소유권등 퍼미션을 그대로 유지
* f : 내가 지정한 파일명으로 압축

#### tar.gz
* 압축하기
```tar cvzf filename.tar.gz download/```

* 압축풀기
```tar xzvpf filename.tar.gz Download/```

#### tar.bz2
* bz2. 압축
```tar -cvjf filename.tar.bz2 Download```

* bz2. 압축풀기
```tar -xvjf filename.tar.bz2```
