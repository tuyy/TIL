#### ls
```ls -la --time-style=full-iso```

```ll -h | grep 'Jan 22' |awk '{system("mv "$9" sogb")}'```

#### date
```
date +%s
date -d @1527217745 +%Y%m%d
date -d ’10 days ago’ +%Y%m%d
date -d ’1 days ago’ +%Y%m%d
date -d '20150704' +%s
```

#### find
```find / -name libpython2.7.so.1.0 2>/dev/null  // error msg 출력x```

#### awk
```
awk \'{print $2}\’
awk '{sum+=1 num++} END {print "sum=" sum " num=" num}'
```

#### sort
```sort -u -k3 -r```

#### thread count
```
$ ps hH p {PID} | wc -l
$ cat /proc/{PID}/status
```

#### file
```file -i filename.txt```

#### iconv
```iconv -f UTF-8 -t ISO-8859-1 hyena.pkg.sh > out.sh```

#### sed
```sed ’s/bad/happy/g’ inputfile ```

#### free
```free | grep Mem | awk '{print $3/$2 * 100.0}'```

### 프로세스 메모리,CPU 사용량 (Top10)
```ps -eo user,pid,pmem,pcpu,time,cmd --sort -rss | head -n 11```

### 프로세스 시작시간
```ps -eo pid,lstart,cmd| grep {process}```

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

#### tgz
* tgz 압축
```tar cvzpf A.tgz shop```

* tgz 압축풀기 
```tar -xvf A.tgz```

### pstack, core 분석
```pstack {PID}```

```gdb {bin} {core_file}```
