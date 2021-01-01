## Item 20: Use std::weak_ptr for std::shared_ptr like pointers that can dangle.
### std::shared_ptr처럼 동작하되 대상을 잃을 수도 있는 포인터가 필요하면 std::weak_ptr을 사용하라
* std::shared_ptr은 std::shared_ptr로 부터 생성된다.
```C++
auto spw = std::make_shared<Widget>();
...
std::weak_ptr<Widget> wpw = spw;
// or std::weak_ptr<Widget> wpw(spw);
...
spw = nullptr; // spw가 대상을 잃었다.

if (wpw.expired()) // wpw가 객체를 잃었으면, (expired되었다고 표현)
{
    ...
}
else
{
    std::shared_ptr<Widget> spw1 = wpw.lock(); // 대상을 잃지 않았으면 해당 shared_ptr 반환, 그렇지않으면 nullptr
    // or auto spw2 = wpw.lock();
    // or std::shared_ptr<Widget> spw3(wpw); // 만약 wpw가 만료이면 std::bad_weak_ptr이 발생
}
```
* 효율성 면에서 std::weak_ptr과 std::shared_ptr은 본질적으로 동일하다.
    * std::weak_ptr과 std::shared_ptr은 객체의 크기가 서로 같다.
    * 원자적 참조 횟수 조작도 관여한다.
* 언제 사용되나?
    * 대상이 만료될 수 있는 캐시 구현
    ```C++
    std::shared_ptr<const Widget> fastLoadWidget(WidgetID id)
    {
        // 캐시에 저장할 포인터는 자신이 대상을 잃었음을 감지할 수 있는 포인터여야 한다.
        static std::unordered_map<WIdgetID, std::weak_ptr<const WIdget>> cache;
        
        auto objPtr = cache[id].lock();
        
        if (objPtr == nullptr)
        {
            objPtr = loadWidget(id);
            cache[id] = objPtr;
        }
        return objPtr;
    }
    ```
    * Observer 패턴에서 subject가 관리하는 observer 목록을 weak_ptr로 구현
    * std::shared_ptr 순환 고리 방지가 있다.

### 참고
* Effective Modern C++ by Scott Meyers
