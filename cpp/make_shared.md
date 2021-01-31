### std::shared_ptr\<T\>(..) vs std::make_shared<T>(..)

#### std::shared_ptr\<T\>
* std::shared_ptr holds only two pointers:
    * the stored pointer (one returned by get());
    * a pointer to control block.
* The control block is a dynamically-allocated object that holds:
    * either a pointer to the managed object or the managed object itself;
    * the deleter (type-erased);
    * the allocator (type-erased);
    * the number of shared_ptrs that own the managed object;
    * the number of weak_ptrs that refer to the managed object.
* control block에서 shared_ptr count와 weak_ptr count가 모두 0이되어야 managed pointer가 delete된다.
* ```std::shared_ptr<T> ptr(new T);``` 로 shared_ptr을 생성하는 경우 두 번의 alloc이 발생한다.
* std::unique_ptr은 std::shared_ptr로 형 변환이 가능하다.

#### std::make_shared\<T\>(..) 로 생성하면?
* std::make_shared\<T\>(..) 로 생성될 땐 control block 생성 한번으로 managed pointer도 함께 생성한다.
* std::make_shared\<T\>(..) 는 기본 deleter가 호출되기 때문에 array를 생성하면 안된다. (std::make_unique<T[]>(..)는 가능)
* exception 발생 문제를 해결할 수 있는 좋은 습관이다. 
```C++
template <typename T>
void foo(std::shared_ptr<T> a, std::shared_ptr<T> b);  
  
foo(new T(1), new T(2)); // 문제의 코드, 호출되는 순서가 보장되지 않아 leak이 발생할 수 있다.
```

#### 참고    
* https://stackoverflow.com/questions/20895648/difference-in-make-shared-and-normal-shared-ptr-in-c
* https://en.cppreference.com/w/cpp/memory/shared_ptr
