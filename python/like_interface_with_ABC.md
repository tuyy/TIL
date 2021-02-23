### ABC (Abstract Base Class)
 ABC를 이용하여 자바나 C++ 과 비슷한 인터페이스를 구현 할 수 있다.
 
 * https://docs.python.org/ko/2.7/library/abc.html
 
 
 ```python
 from abc import *
 
class MyBase(metaclass=ABCMeta):
    @abstractmethod
    def do_it(self):
        pass

 
class MySub(MyBase):
    # 구현을 강제한다.
    def do_it(self):
        print('just do it!')
 
obj = MySub()
obj.do_it()
```
