### install mysql source
```
wget http://www.atblog.co.kr/file/package/mysql-5.6.14.tar.gz
tar zxvf mysql-5.6.14.tar.gz
cd mysql-5.6.14

cmake \
-DCMAKE_INSTALL_PREFIX=/usr/local/mysql \   <-- 수정 
-DMYSQL_DATADIR=/var/mysql/data \   <-- 수정
-DMYSQL_UNIX_ADDR=/var/mysql/mysql.sock \   <-- 수정
-DSYSCONFDIR=/etc \
-DMYSQL_TCP_PORT=3306 \
-DMYSQL_USER=mysql \
-DDEFAULT_CHARSET=utf8 \
-DDEFAULT_COLLATION=utf8_general_ci \
-DWITH_EXTRA_CHARSETS=all \
-DENABLED_LOCAL_INFILE=1 \
-DWITH_INNOBASE_STORAGE_ENGINE=1 \
-DWITH_ARCHIVE_STORAGE_ENGINE=1 \
-DWITH_BLACKHOLE_STORAGE_ENGINE=1

make
sudo make install
```

* 이후 root su 없이 실행
    *https://superuser.com/questions/209203/how-can-i-install-mysql-on-centos-without-being-root-su
