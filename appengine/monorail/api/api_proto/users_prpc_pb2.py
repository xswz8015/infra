# Generated by the pRPC protocol buffer compiler plugin.  DO NOT EDIT!
# source: api/api_proto/users.proto

import base64
import zlib

from google.protobuf import descriptor_pb2

# Includes description of the api/api_proto/users.proto and all of its transitive
# dependencies. Includes source code info.
FILE_DESCRIPTOR_SET = descriptor_pb2.FileDescriptorSet()
FILE_DESCRIPTOR_SET.ParseFromString(zlib.decompress(base64.b64decode(
    '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')))
_INDEX = {
    f.name: {
      'descriptor': f,
      'services': {s.name: s for s in f.service},
    }
    for f in FILE_DESCRIPTOR_SET.file
}


UsersServiceDescription = {
  'file_descriptor_set': FILE_DESCRIPTOR_SET,
  'file_descriptor': _INDEX[u'api/api_proto/users.proto']['descriptor'],
  'service_descriptor': _INDEX[u'api/api_proto/users.proto']['services'][u'Users'],
}
