#!/usr/bin/env python
# -*- coding: utf-8 -*-
import argparse
import msgpack
import redis


if __name__ == "__main__":
    r = redis.StrictRedis(host='localhost', port=6379, db=0)

    packed = r.get("go_map")
    dic = msgpack.unpackb(packed, encoding='utf-8')
    packed = r.get("go_ls_int")
    ls_int = msgpack.unpackb(packed, encoding='utf-8')
    packed = r.get("go_ls_float")
    ls_float = msgpack.unpackb(packed, encoding='utf-8')

    print dic
    print ls_int
    print ls_float

