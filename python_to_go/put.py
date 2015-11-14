#!/usr/bin/env python
# -*- coding: utf-8 -*-
import argparse
import msgpack
import redis


if __name__ == "__main__":
    r = redis.StrictRedis(host='localhost', port=6379, db=0)
    # packing
    packed_ls_int = msgpack.packb(range(10))
    packed_ls_float = msgpack.packb([-1.0, 2.02, 5.04])
    packed_dic = msgpack.packb({u'pykey1': u'pyval1', u'pykey2': u'pyval2'})
    # put
    r.set("py_ls_int", packed_ls_int)
    r.set("py_ls_float", packed_ls_float)
    r.set("py_map", packed_dic)


