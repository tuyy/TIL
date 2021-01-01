### apache php module 연동

* php 설치하고 load module 해야한다.
* httpd.conf
```
LoadModule php5_module modules/libphp5.so
AddType application/x-httpd-php .php .nhn .phtml
AddType application/x-httpd-php-source .phps
```
