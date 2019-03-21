* on centos 6
```
# 1. Install a package with repository for your system:
# On CentOS, install package centos-release-scl available in CentOS repository:
$ sudo yum install centos-release-scl

# On RHEL, enable RHSCL repository for you system:
$ sudo yum-config-manager --enable rhel-server-rhscl-7-rpms

# 2. Install the collection:
$ sudo yum install devtoolset-7

# 3. Start using software collections:
$ scl enable devtoolset-7 bash
```
