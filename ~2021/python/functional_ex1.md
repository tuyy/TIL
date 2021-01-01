### 함수형 프로그래밍 python ver 예시 작성해봄

```python3
#-*- coding:utf-8 -*-

arr = [1,2,3,4,5,6,7,8,9]
length = 3

#rz = 0
#for v in arr:
#    if v % 2:
#        v = v*v
#        rz += v
#        length -= 1
#        if length == 0:
#            break
#print(rz)

def _filter(f, arr):
    for v in arr:
        if f(v):
            yield v

#for v in _filter(lambda v : v % 2, arr):
#    v = v*v
#    print(v)
#    length -= 1
#    if length == 0:
#        break

def _map(f, arr):
    for v in arr:
        yield f(v)

#for v in _map(lambda v : v*v, _filter(lambda v : v % 2, arr)):
#    print(v)
#    length -= 1
#    if length == 0:
#        break

def _take(length, arr):
    for v in arr:
        yield v
        length -= 1
        if length == 0:
            break

#for v in _take(length, _map(lambda v : v*v, _filter(lambda v : v % 2, arr))):
#    print(v)

def _reduce(f, acc, arr=None):
    if arr == None:
        arr = iter(acc)
        acc = next(arr)

    for v in arr:
        acc = f(acc, v)
    return acc

#rz = _reduce(lambda acc, v : acc+v, 0, _take(length, _map(lambda v : v*v, _filter(lambda v : v % 2, arr))))
#rz = _reduce(lambda acc, v : acc+v,
#     _take(length,
#     _map(lambda v : v*v,
#     _filter(lambda v : v % 2, arr))))
#
#print(rz)

def _forEach(f, arr):
    for v in arr:
        f(v)

#_forEach(print, [_reduce(lambda acc, v : acc+v, 0, _take(length, _map(lambda v : v*v, _filter(lambda v : v % 2, arr))))])

def go(*args):
    return _reduce(lambda acc, f : f(acc), args)

#go(arr,
#   lambda arr : _filter(lambda v : v%2, arr),
#   lambda arr : _map(lambda v : v*v, arr),
#   lambda arr : _take(length, arr),
#   lambda arr : _reduce(lambda acc, v : acc+v, arr),
#   print)


def curry(f):
    def curried(a, *args):
        if len(args) == 0:
            return lambda *args : f(a, *args)
        else:
            return f(a, *args)
    return curried

_filter = curry(_filter)
_map = curry(_map)
_take = curry(_take)
_reduce = curry(_reduce)

#go(arr,
#   lambda arr : _filter(lambda v : v%2)(arr),
#   lambda arr : _map(lambda v : v*v)(arr),
#   lambda arr : _take(length)(arr),
#   lambda arr : _reduce(lambda acc, v : acc+v)(arr),
#   print)

go(arr,
   _filter(lambda v : v%2),
   _map(lambda v : v*v),
   _take(length),
   _reduce(lambda acc, v : acc+v),
   print)

```
