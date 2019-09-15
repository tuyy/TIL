### 함수형 프로그래밍
```python3
def curry(f):
    def curried(a, *args):
        if len(args) == 0:
            return lambda *args : f(a, *args)
        else:
            return f(a, *args)
    return curried


def _map(f, it):
    for a in it:
        yield f(a)


def _filter(f, it):
    for a in it:
        if f(a):
            yield a


def _reject(f, it):
    return filter(lambda a: not f(a), it)
    
    
# length 만큼만 반환
def _take(length, it):
    for a in it:
        yield a
        length -= 1
        if length == 0:
            break


# reduce
def _reduce(f, acc, it=None):
    if it is None:
        it = iter(acc)
        acc = next(it)

    for a in it:
        acc = f(acc, a)
    return acc


# iter 값 하나씩 효과 주기
def _each(f, it):
    for a in it:
        f(a)


def is_iterable(it):
    try:
        _ = (e for e in it)
        return True
    except TypeError:
        return False


# [1, 2, 3, [4, 5], [6, 7]] to [1, 2, 3, 4, 5, 6, 7]
def flat(it):
    for a in it:
        if is_iterable(a):
            for b in a:
                yield b
        else:
            yield a


# curried function
filter = curry(_filter)
reject = curry(_reject)
map = curry(_map)
reduce = curry(_reduce)
take = curry(_take)
each = curry(_each)


def go(*args):
    return reduce(lambda a,f: f(a), args)


# 여러 함수의 조합 생성(순차적으로 함수 실행)
def pipe(*args):
    return lambda a: go(a, *args)


# 하나라도 참이면 True
some = curry(lambda f, it: go(filter(f, it),
                              take(1),
                              lambda a: len(list(a)) > 0))


# 모두 참인 경우 True
every = curry(lambda f, it: go(reject(f, it),
                               take(1),
                               lambda a: len(list(a)) == 0))


# 0, None 제외하여 iter 반환
compact = filter(lambda a: a)


# 사칙연산
add = lambda a,b: a+b
sub = lambda a,b: a-b
mul = lambda a,b: a*b
div = lambda a,b: a/b
mod = lambda a,b: a%b


# keys [k1, k2]
def keys(obj):
    for k in obj:
        yield k


# values [v1, v2]
def values(obj):
    for k in obj:
        yield obj[k]


# entries [[k1, v1], [k2, v2]]
def entries(obj):
    for k in obj:
        yield [k, obj[k]]
        
        
def add_dict(a, b):
    a.update(b)
    return a


# object [['a',1], ['b',2]] to {'a': 1, 'b': 2}
def object(it):
    return go(it,
              map(lambda d: {d[0]: d[1]}),
              reduce(add_dict))


# mapObject
# map_object(lambda v: v*10, {'a': 1, 'b': 2})
# {'a': 10, 'b': 20}
def map_object(f, obj):
    return go(obj,
              entries,
              map(lambda v: [v[0], f(v[1])]),
              object)


# pick 첫 번째 인자로 받은 키 값의 value를 반환
# pick(['b', 'c'], {'a':1, 'b':2, 'c':3, 'd':4})
# {'b': 2, 'c': 3}


def pick(keys, obj):
    return go(obj,
              entries,
              filter(lambda v: v[0] in keys),
              take(len(keys)),
              object)
```
