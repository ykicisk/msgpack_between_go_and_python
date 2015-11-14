# msgpack_between_go_and_python
transportation data between golang and python via msgpack 

## prepare

```bash
$ pip install msgpack-python
$ go get gopkg.in/redis.v3
$ go get gopkg.in/vmihailenco/msgpack.v2
$ redis-server
```
## run

### go -> python

```
$ cd go_to_python
$ go run put.go
$ python get.py
{u'gokey2': u'goval2', u'gokey1': u'goval1'}
[1, 3, 5, 8]
[2.3, 6.4, 5.5]
```

### python -> go

```
$ cd python_to_go
$ python put.py
$ go run get.go
map[pykey1:pyval1 pykey2:pyval2]
[0 1 2 3 4 5 6 7 8 9]
[-1 2.02 5.04]
```

