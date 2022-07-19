"""
pip install protobuf
"""
from .pb_pb2 import Teacher


t=Teacher()
t.ParseFromString(res.content)

t1.ParseFromString(res.content)
t.SerializeToString() #b'\n\x03pan\x101\x18\x02"\x04math'