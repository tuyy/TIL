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

#### env

#### file
```file -i filename.txt```

#### iconv
```iconv -f UTF-8 -t ISO-8859-1 hyena.pkg.sh > out.sh```

#### sed
```sed ’s/bad/happy/g’ inputfile ```

#### free
```free | grep Mem | awk '{print $3/$2 * 100.0}'```

### 프로세스별 메모리 사용량 
```ps -eo user,pid,pmem,pcpu,time,cmd --sort -rss | head -n 11```

### running process info with detail time
```ps -eo pid,lstart,cmd|grep parser```
