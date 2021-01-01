## Item 19: Use std::shared_ptr for shared-ownership resource management.
### 소유권 공유 자원 관리에는 std::shared_ptr을 사용하라
* std::shared_ptr의 크기는 raw pointer + control block의 크다.
* control block
    * raw pointer(or 가리키는 포인터), shared_ptr 참조횟수, weak_ptr 참조횟수, 기타 자료(커스텀 삭제자, 할당자 등)
    * control block은 힙 메모리에 할당해야한다.
    * 참조 횟수 증감은 반드시 원자적 연산이어야 한다.
    * 복사 시에는 참조 횟수를 증가해야 하지만 이동 시에는 증가할 필요가 없다.
* 삭제자를 지정 할 수 있으며 삭제자를 지정해도 추가적인 메모리가 필요 없다. (control block에 할당되기 때문에)
```C++
auto loggingDel = [](Widget *pw)
{
    makeLogEntry(pw);
    delete pw;
};

std::unique_ptr<Widget, decltype(loggingDel)> upw(new Widget, loggingDel);
std::shared_ptr<Widget> spw(new Widget, loggingDel); // unique_ptr의 비해 설계가 유연하다.(control block에 저장하기 때문에 가능)
```
* 잘못된 사용의 예 : 같은 raw pointer로 shared_ptr을 두 번 할당하면, 두 번 delete되어 미정의 동작 발생
```C++
auto pw = new Widget;
...
std::shared_ptr<Widget> spw1(pw, loggingDel);
...
std::shared_ptr<Widget> spw2(pw, loggingDel);

// 결과적으로, std::make_shared를 이용하거나 raw pointer 변수를 거치지 말고 new의 결과를 직접 전달하라
```
* unique_ptr과는 다르게 배열을 관리 할 수 없다. 즉 std::shared_ptr<T[]> 같은건 없다. (delete[]가 호출되지 않는다.)

### 참고
* Effective Modern C++ by Scott Meyers
