### centos 6 GLIB_2.14 설치
```
// 설치
mkdir ~/glibc_install; cd ~/glibc_install
wget http://ftp.gnu.org/gnu/glibc/glibc-2.14.tar.gz
tar zxvf glibc-2.14.tar.gz
cd glibc-2.14;mkdir build;cd build
export LD_LIBRARY_PATH=/opt/glibc-2.14/lib
../configure --prefix=/opt/glibc-2.14
make -j4
sudo make install

// ~/.bashrc에 추가
export LD_LIBRARY_PATH=/opt/glibc-2.14/lib
```
