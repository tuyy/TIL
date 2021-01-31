### weak_ptr
* std::shared_ptr control block에 weakptr_count 정보가 담겨있다.
* 사용시 lock() 함수를 통해 std::shared_ptr로 형변환해야한다.
* cache 같은 곳에 유용하게 사용될 수 있다.

```C++
int32_t main()
{
    std::weak_ptr<Test> sWPtr;
    {
        auto sSPtr = std::make_shared<Test>();
        sWPtr = sSPtr;
        if (sWPtr.expired() == false)
        {
            sWPtr.lock()->print();
        }
    }

    if (sWPtr.expired() == false)
    {
        sWPtr.lock()->print();
    }

    return 0;
}

```

#### 참고
* https://en.cppreference.com/w/cpp/memory/weak_ptr
* https://stackoverflow.com/questions/12030650/when-is-stdweak-ptr-useful
