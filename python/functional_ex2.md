### 함수형 프로그래밍
```python3
import inspect

def curry(f):
    args_cnt = len(inspect.signature(f).parameters)
    def curried(a, *args):
        if len(args)+1 == args_cnt:
            return f(a, *args)

        # 함수의 인수가 3개일때
        if args_cnt == 3:
            # reduce 용 curry. 인자 2개일때와 2개 커리하는 경우 처리
            if len(args) == 1:
                if args[0] == [] or args[0] == {}:
                    return lambda it: f(a, args[0], it)
                else:
                    it = iter(args[0])
                    return f(a, next(it), it)
            else:
                return lambda it: f(a, next(it), it)
        else:
            # 함수의 인수가 2개일때
            return lambda *args : f(a, *args)
            
    return curried


# 즉시평가 map
@curry
def i_map(f, it):
    args_cnt = len(inspect.signature(f).parameters)
    if args_cnt > 2 or args_cnt == 0:
        return

    rz = []
    i = 0
    for v in it:
        if args_cnt == 2:
            rz.append(f(v, i))
            i += 1
        else:
            rz.append(f(v))
    return rz


@curry
def map(f, it):
    args_cnt = len(inspect.signature(f).parameters)
    if args_cnt > 2 or args_cnt == 0:
        return

    i = 0
    for v in it:
        if args_cnt == 2:
            yield f(v, i)
            i += 1
        else:
            yield f(v)


@curry
def i_filter(f, it):
    args_cnt = len(inspect.signature(f).parameters)
    if args_cnt > 2 or args_cnt == 0:
        return

    rz = []
    i = 0
    for a in it:
        if args_cnt == 2:
            if f(a, i):
                rz.append(a)
            i += 1
        else:
            if f(a):
                rz.append(a)
    return rz


@curry
def filter(f, it):
    args_cnt = len(inspect.signature(f).parameters)
    if args_cnt > 2 or args_cnt == 0:
        return

    i = 0
    for a in it:
        if args_cnt == 2:
            if f(a, i):
                yield a
            i += 1
        else:
            if f(a):
                yield a


@curry
def reject(f, it):
    args_cnt = len(inspect.signature(f).parameters)
    if args_cnt > 2 or args_cnt == 0:
        return

    if args_cnt == 2:
        return filter(lambda a, i: not f(a, i), it)
    else:
        return filter(lambda a: not f(a), it)


# length 만큼만 반환
@curry
def take(length, it):
    for a in it:
        yield a
        length -= 1
        if length == 0:
            break


# 즉시평가 take
@curry
def i_take(length, it):
    rz = []
    for a in it:
        rz.append(a)
        length -= 1
        if length == 0:
            break
    return rz


# reduce
@curry
def reduce(f, acc, it):
    for a in it:
        acc = f(acc, a)
    return acc


# iter 값 하나씩 효과 주기
@curry
def each(f, it):
    for a in it:
        f(a)


def is_iterable(it):
    try:
        (e for e in it)
    except TypeError:
        return False
    else:
        return True


# [1, 2, 3, [4, 5], [6, 7]] => [1, 2, 3, 4, 5, 6, 7]
def flat(it):
    for a in it:
        if is_iterable(a):
            for b in a:
                yield b
        else:
            yield a


# 즉시평가 flat
def i_flat(it):
    rz = []
    for a in it:
        if is_iterable(a):
            for b in a:
                rz.append(b)
        else:
            rz.append(a)
    return rz


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


# entries {k1:v1, k2:v2} => [[k1, v1], [k2, v2]]
def entries(obj):
    for k in obj:
        yield [k, obj[k]]


def add_dict(a, b):
    a.update(b)
    return a


# object [['a',1], ['b',2]] => {'a': 1, 'b': 2}
def object(it):
    return go(it,
              map(lambda d: {d[0]: d[1]}),
              reduce(add_dict))


# 모든 map의 값들에 함수를 적용한다.
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


# uniq [1, 1, 2, 2, 3] => [1, 2, 3]
def uniq(it):
    tmp = []
    for a in it:
        if a not in tmp:
            tmp.append(a)
            yield a


# count [1, 1, 2, 2, 3] => [[1, 2], [2, 2], [3, 1]]
def count(it):
    for i in uniq(it):
        cnt = 0
        for j in it:
            if i == j:
                cnt += 1
        yield [i, cnt]


@curry
def count_by(f, it):
    cnt = 0
    for v in it:
        if f(v):
            cnt += 1
    return cnt


# sort_by [[3, 2], [1, 4]] => [[1, 4], [3, 2]]
sort_by = curry(lambda f,it: sorted(it, key=f))
sort_reverse_by = curry(lambda f,it: sorted(it, key=f, reverse=True))


# index_by [{'id':1, 'val':10 }, {'id':2, 'val':20}] => [1:{..}, 2:{..}]
@curry
def index_by(f, it):
    return reduce(lambda obj,a: add_dict(obj, {f(a): a}), {}, it)


# group_by [{'id':1, 'val':10 }, {'id':2, 'val':20}] => [1:{..}, 2:{..}]
@curry
def group_by(f, it):
    def _push(p, k, v):
        if k in p:
            p[k].append(v)
        else:
            p[k] = [v]
        return p

    return reduce(lambda group,a: _push(group, f(a), a), {}, it)


def is_uniq(it):
    return len(list(it)) == len(set(it))


def reverse(it):
    rz = list(it)
    rz.reverse()
    return rz


def find(f, it):
    for v in it:
        if f(v):
            return v

```
