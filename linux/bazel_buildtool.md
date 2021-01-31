## bazel
* 여러 언어를 아우르는 build 툴
* centos7 보다 낮은 OS에선 소스파일을 받아서 빌드해야함
* [C++ 기본 예제](https://docs.bazel.build/versions/master/tutorial/cpp.html)
```
Bazel is an open-source build and test tool similar to Make, Maven, and Gradle.
It uses a human-readable, high-level build language.

bazel build //main:hello-world
bazel-bin/main/hello-world
```
