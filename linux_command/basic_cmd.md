#### ls
```ls -la --time-style=full-iso```

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
```awk \'{print $2}\’```

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
