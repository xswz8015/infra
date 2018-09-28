# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/api_proto/issue_objects.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from api.api_proto import common_pb2 as api_dot_api__proto_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/api_proto/issue_objects.proto',
  package='monorail',
  syntax='proto3',
  serialized_pb=_b('\n!api/api_proto/issue_objects.proto\x12\x08monorail\x1a\x1a\x61pi/api_proto/common.proto\"\xe3\x01\n\x08\x41pproval\x12%\n\tfield_ref\x18\x01 \x01(\x0b\x32\x12.monorail.FieldRef\x12(\n\rapprover_refs\x18\x02 \x03(\x0b\x32\x11.monorail.UserRef\x12(\n\x06status\x18\x03 \x01(\x0e\x32\x18.monorail.ApprovalStatus\x12\x0e\n\x06set_on\x18\x04 \x01(\x07\x12%\n\nsetter_ref\x18\x05 \x01(\x0b\x32\x11.monorail.UserRef\x12%\n\tphase_ref\x18\x07 \x01(\x0b\x32\x12.monorail.PhaseRef\"N\n\tAmendment\x12\x12\n\nfield_name\x18\x01 \x01(\t\x12\x1a\n\x12new_or_delta_value\x18\x02 \x01(\t\x12\x11\n\told_value\x18\x03 \x01(\t\"\xac\x01\n\nAttachment\x12\x15\n\rattachment_id\x18\x01 \x01(\x04\x12\x10\n\x08\x66ilename\x18\x02 \x01(\t\x12\x0c\n\x04size\x18\x03 \x01(\x04\x12\x14\n\x0c\x63ontent_type\x18\x04 \x01(\t\x12\x12\n\nis_deleted\x18\x05 \x01(\x08\x12\x15\n\rthumbnail_url\x18\x06 \x01(\t\x12\x10\n\x08view_url\x18\x07 \x01(\t\x12\x14\n\x0c\x64ownload_url\x18\x08 \x01(\t\"\xe6\x02\n\x07\x43omment\x12\x14\n\x0cproject_name\x18\x01 \x01(\t\x12\x10\n\x08local_id\x18\x02 \x01(\r\x12\x14\n\x0csequence_num\x18\x03 \x01(\r\x12\x12\n\nis_deleted\x18\x04 \x01(\x08\x12$\n\tcommenter\x18\x05 \x01(\x0b\x32\x11.monorail.UserRef\x12\x11\n\ttimestamp\x18\x06 \x01(\x07\x12\x0f\n\x07\x63ontent\x18\x07 \x01(\t\x12\x17\n\x0finbound_message\x18\x08 \x01(\t\x12\'\n\namendments\x18\t \x03(\x0b\x32\x13.monorail.Amendment\x12)\n\x0b\x61ttachments\x18\n \x03(\x0b\x32\x14.monorail.Attachment\x12(\n\x0c\x61pproval_ref\x18\x0b \x01(\x0b\x32\x12.monorail.FieldRef\x12\x17\n\x0f\x64\x65scription_num\x18\x0c \x01(\r\x12\x0f\n\x07is_spam\x18\r \x01(\x08\"}\n\nFieldValue\x12%\n\tfield_ref\x18\x01 \x01(\x0b\x32\x12.monorail.FieldRef\x12\r\n\x05value\x18\x02 \x01(\t\x12\x12\n\nis_derived\x18\x03 \x01(\x08\x12%\n\tphase_ref\x18\x04 \x01(\x0b\x32\x12.monorail.PhaseRef\"\xa1\x06\n\x05Issue\x12\x14\n\x0cproject_name\x18\x01 \x01(\t\x12\x10\n\x08local_id\x18\x02 \x01(\r\x12\x0f\n\x07summary\x18\x03 \x01(\t\x12\'\n\nstatus_ref\x18\x04 \x01(\x0b\x32\x13.monorail.StatusRef\x12$\n\towner_ref\x18\x05 \x01(\x0b\x32\x11.monorail.UserRef\x12\"\n\x07\x63\x63_refs\x18\x06 \x03(\x0b\x32\x11.monorail.UserRef\x12&\n\nlabel_refs\x18\x07 \x03(\x0b\x32\x12.monorail.LabelRef\x12.\n\x0e\x63omponent_refs\x18\x08 \x03(\x0b\x32\x16.monorail.ComponentRef\x12\x31\n\x15\x62locked_on_issue_refs\x18\t \x03(\x0b\x32\x12.monorail.IssueRef\x12/\n\x13\x62locking_issue_refs\x18\n \x03(\x0b\x32\x12.monorail.IssueRef\x12\x34\n\x18\x64\x61ngling_blocked_on_refs\x18\x17 \x03(\x0b\x32\x12.monorail.IssueRef\x12\x31\n\x15merged_into_issue_ref\x18\x0b \x01(\x0b\x32\x12.monorail.IssueRef\x12*\n\x0c\x66ield_values\x18\x0c \x03(\x0b\x32\x14.monorail.FieldValue\x12\x12\n\nis_deleted\x18\r \x01(\x08\x12\'\n\x0creporter_ref\x18\x0e \x01(\x0b\x32\x11.monorail.UserRef\x12\x18\n\x10opened_timestamp\x18\x0f \x01(\x07\x12\x18\n\x10\x63losed_timestamp\x18\x10 \x01(\x07\x12\x1a\n\x12modified_timestamp\x18\x11 \x01(\x07\x12\x12\n\nstar_count\x18\x12 \x01(\r\x12\x0f\n\x07is_spam\x18\x13 \x01(\x08\x12\x18\n\x10\x61ttachment_count\x18\x14 \x01(\r\x12+\n\x0f\x61pproval_values\x18\x15 \x03(\x0b\x32\x12.monorail.Approval\x12\"\n\x06phases\x18\x16 \x03(\x0b\x32\x12.monorail.PhaseDef\"G\n\x0cIssueSummary\x12\x14\n\x0cproject_name\x18\x01 \x01(\t\x12\x10\n\x08local_id\x18\x02 \x01(\r\x12\x0f\n\x07summary\x18\x03 \x01(\t\"?\n\x08PhaseDef\x12%\n\tphase_ref\x18\x01 \x01(\x0b\x32\x12.monorail.PhaseRef\x12\x0c\n\x04rank\x18\x02 \x01(\r\"\x1e\n\x08PhaseRef\x12\x12\n\nphase_name\x18\x01 \x01(\t*\x90\x01\n\x0e\x41pprovalStatus\x12\x0b\n\x07NOT_SET\x10\x00\x12\x10\n\x0cNEEDS_REVIEW\x10\x01\x12\x06\n\x02NA\x10\x02\x12\x14\n\x10REVIEW_REQUESTED\x10\x03\x12\x12\n\x0eREVIEW_STARTED\x10\x04\x12\r\n\tNEED_INFO\x10\x05\x12\x0c\n\x08\x41PPROVED\x10\x06\x12\x10\n\x0cNOT_APPROVED\x10\x07*^\n\x0b\x43\x61nnedQuery\x12\x07\n\x03\x41LL\x10\x00\x12\x07\n\x03NEW\x10\x01\x12\x08\n\x04OPEN\x10\x02\x12\t\n\x05OWNED\x10\x03\x12\x0c\n\x08REPORTED\x10\x04\x12\x0b\n\x07STARRED\x10\x05\x12\r\n\tTO_VERIFY\x10\x06\x62\x06proto3')
  ,
  dependencies=[api_dot_api__proto_dot_common__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_APPROVALSTATUS = _descriptor.EnumDescriptor(
  name='ApprovalStatus',
  full_name='monorail.ApprovalStatus',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NOT_SET', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NEEDS_REVIEW', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NA', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REVIEW_REQUESTED', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REVIEW_STARTED', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NEED_INFO', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='APPROVED', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NOT_APPROVED', index=7, number=7,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=2023,
  serialized_end=2167,
)
_sym_db.RegisterEnumDescriptor(_APPROVALSTATUS)

ApprovalStatus = enum_type_wrapper.EnumTypeWrapper(_APPROVALSTATUS)
_CANNEDQUERY = _descriptor.EnumDescriptor(
  name='CannedQuery',
  full_name='monorail.CannedQuery',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ALL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NEW', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OPEN', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OWNED', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REPORTED', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STARRED', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TO_VERIFY', index=6, number=6,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=2169,
  serialized_end=2263,
)
_sym_db.RegisterEnumDescriptor(_CANNEDQUERY)

CannedQuery = enum_type_wrapper.EnumTypeWrapper(_CANNEDQUERY)
NOT_SET = 0
NEEDS_REVIEW = 1
NA = 2
REVIEW_REQUESTED = 3
REVIEW_STARTED = 4
NEED_INFO = 5
APPROVED = 6
NOT_APPROVED = 7
ALL = 0
NEW = 1
OPEN = 2
OWNED = 3
REPORTED = 4
STARRED = 5
TO_VERIFY = 6



_APPROVAL = _descriptor.Descriptor(
  name='Approval',
  full_name='monorail.Approval',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='field_ref', full_name='monorail.Approval.field_ref', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='approver_refs', full_name='monorail.Approval.approver_refs', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='monorail.Approval.status', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='set_on', full_name='monorail.Approval.set_on', index=3,
      number=4, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='setter_ref', full_name='monorail.Approval.setter_ref', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='phase_ref', full_name='monorail.Approval.phase_ref', index=5,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=76,
  serialized_end=303,
)


_AMENDMENT = _descriptor.Descriptor(
  name='Amendment',
  full_name='monorail.Amendment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='field_name', full_name='monorail.Amendment.field_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='new_or_delta_value', full_name='monorail.Amendment.new_or_delta_value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='old_value', full_name='monorail.Amendment.old_value', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=305,
  serialized_end=383,
)


_ATTACHMENT = _descriptor.Descriptor(
  name='Attachment',
  full_name='monorail.Attachment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='attachment_id', full_name='monorail.Attachment.attachment_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='filename', full_name='monorail.Attachment.filename', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='size', full_name='monorail.Attachment.size', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='content_type', full_name='monorail.Attachment.content_type', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_deleted', full_name='monorail.Attachment.is_deleted', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='thumbnail_url', full_name='monorail.Attachment.thumbnail_url', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='view_url', full_name='monorail.Attachment.view_url', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='download_url', full_name='monorail.Attachment.download_url', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=386,
  serialized_end=558,
)


_COMMENT = _descriptor.Descriptor(
  name='Comment',
  full_name='monorail.Comment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_name', full_name='monorail.Comment.project_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='local_id', full_name='monorail.Comment.local_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sequence_num', full_name='monorail.Comment.sequence_num', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_deleted', full_name='monorail.Comment.is_deleted', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='commenter', full_name='monorail.Comment.commenter', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='monorail.Comment.timestamp', index=5,
      number=6, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='content', full_name='monorail.Comment.content', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='inbound_message', full_name='monorail.Comment.inbound_message', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='amendments', full_name='monorail.Comment.amendments', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='attachments', full_name='monorail.Comment.attachments', index=9,
      number=10, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='approval_ref', full_name='monorail.Comment.approval_ref', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description_num', full_name='monorail.Comment.description_num', index=11,
      number=12, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_spam', full_name='monorail.Comment.is_spam', index=12,
      number=13, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=561,
  serialized_end=919,
)


_FIELDVALUE = _descriptor.Descriptor(
  name='FieldValue',
  full_name='monorail.FieldValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='field_ref', full_name='monorail.FieldValue.field_ref', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='monorail.FieldValue.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_derived', full_name='monorail.FieldValue.is_derived', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='phase_ref', full_name='monorail.FieldValue.phase_ref', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=921,
  serialized_end=1046,
)


_ISSUE = _descriptor.Descriptor(
  name='Issue',
  full_name='monorail.Issue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_name', full_name='monorail.Issue.project_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='local_id', full_name='monorail.Issue.local_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='summary', full_name='monorail.Issue.summary', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status_ref', full_name='monorail.Issue.status_ref', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='owner_ref', full_name='monorail.Issue.owner_ref', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cc_refs', full_name='monorail.Issue.cc_refs', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='label_refs', full_name='monorail.Issue.label_refs', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='component_refs', full_name='monorail.Issue.component_refs', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='blocked_on_issue_refs', full_name='monorail.Issue.blocked_on_issue_refs', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='blocking_issue_refs', full_name='monorail.Issue.blocking_issue_refs', index=9,
      number=10, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='dangling_blocked_on_refs', full_name='monorail.Issue.dangling_blocked_on_refs', index=10,
      number=23, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='merged_into_issue_ref', full_name='monorail.Issue.merged_into_issue_ref', index=11,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='field_values', full_name='monorail.Issue.field_values', index=12,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_deleted', full_name='monorail.Issue.is_deleted', index=13,
      number=13, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='reporter_ref', full_name='monorail.Issue.reporter_ref', index=14,
      number=14, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='opened_timestamp', full_name='monorail.Issue.opened_timestamp', index=15,
      number=15, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='closed_timestamp', full_name='monorail.Issue.closed_timestamp', index=16,
      number=16, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='modified_timestamp', full_name='monorail.Issue.modified_timestamp', index=17,
      number=17, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='star_count', full_name='monorail.Issue.star_count', index=18,
      number=18, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_spam', full_name='monorail.Issue.is_spam', index=19,
      number=19, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='attachment_count', full_name='monorail.Issue.attachment_count', index=20,
      number=20, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='approval_values', full_name='monorail.Issue.approval_values', index=21,
      number=21, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='phases', full_name='monorail.Issue.phases', index=22,
      number=22, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1049,
  serialized_end=1850,
)


_ISSUESUMMARY = _descriptor.Descriptor(
  name='IssueSummary',
  full_name='monorail.IssueSummary',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_name', full_name='monorail.IssueSummary.project_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='local_id', full_name='monorail.IssueSummary.local_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='summary', full_name='monorail.IssueSummary.summary', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1852,
  serialized_end=1923,
)


_PHASEDEF = _descriptor.Descriptor(
  name='PhaseDef',
  full_name='monorail.PhaseDef',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phase_ref', full_name='monorail.PhaseDef.phase_ref', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='rank', full_name='monorail.PhaseDef.rank', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1925,
  serialized_end=1988,
)


_PHASEREF = _descriptor.Descriptor(
  name='PhaseRef',
  full_name='monorail.PhaseRef',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phase_name', full_name='monorail.PhaseRef.phase_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1990,
  serialized_end=2020,
)

_APPROVAL.fields_by_name['field_ref'].message_type = api_dot_api__proto_dot_common__pb2._FIELDREF
_APPROVAL.fields_by_name['approver_refs'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_APPROVAL.fields_by_name['status'].enum_type = _APPROVALSTATUS
_APPROVAL.fields_by_name['setter_ref'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_APPROVAL.fields_by_name['phase_ref'].message_type = _PHASEREF
_COMMENT.fields_by_name['commenter'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_COMMENT.fields_by_name['amendments'].message_type = _AMENDMENT
_COMMENT.fields_by_name['attachments'].message_type = _ATTACHMENT
_COMMENT.fields_by_name['approval_ref'].message_type = api_dot_api__proto_dot_common__pb2._FIELDREF
_FIELDVALUE.fields_by_name['field_ref'].message_type = api_dot_api__proto_dot_common__pb2._FIELDREF
_FIELDVALUE.fields_by_name['phase_ref'].message_type = _PHASEREF
_ISSUE.fields_by_name['status_ref'].message_type = api_dot_api__proto_dot_common__pb2._STATUSREF
_ISSUE.fields_by_name['owner_ref'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_ISSUE.fields_by_name['cc_refs'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_ISSUE.fields_by_name['label_refs'].message_type = api_dot_api__proto_dot_common__pb2._LABELREF
_ISSUE.fields_by_name['component_refs'].message_type = api_dot_api__proto_dot_common__pb2._COMPONENTREF
_ISSUE.fields_by_name['blocked_on_issue_refs'].message_type = api_dot_api__proto_dot_common__pb2._ISSUEREF
_ISSUE.fields_by_name['blocking_issue_refs'].message_type = api_dot_api__proto_dot_common__pb2._ISSUEREF
_ISSUE.fields_by_name['dangling_blocked_on_refs'].message_type = api_dot_api__proto_dot_common__pb2._ISSUEREF
_ISSUE.fields_by_name['merged_into_issue_ref'].message_type = api_dot_api__proto_dot_common__pb2._ISSUEREF
_ISSUE.fields_by_name['field_values'].message_type = _FIELDVALUE
_ISSUE.fields_by_name['reporter_ref'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_ISSUE.fields_by_name['approval_values'].message_type = _APPROVAL
_ISSUE.fields_by_name['phases'].message_type = _PHASEDEF
_PHASEDEF.fields_by_name['phase_ref'].message_type = _PHASEREF
DESCRIPTOR.message_types_by_name['Approval'] = _APPROVAL
DESCRIPTOR.message_types_by_name['Amendment'] = _AMENDMENT
DESCRIPTOR.message_types_by_name['Attachment'] = _ATTACHMENT
DESCRIPTOR.message_types_by_name['Comment'] = _COMMENT
DESCRIPTOR.message_types_by_name['FieldValue'] = _FIELDVALUE
DESCRIPTOR.message_types_by_name['Issue'] = _ISSUE
DESCRIPTOR.message_types_by_name['IssueSummary'] = _ISSUESUMMARY
DESCRIPTOR.message_types_by_name['PhaseDef'] = _PHASEDEF
DESCRIPTOR.message_types_by_name['PhaseRef'] = _PHASEREF
DESCRIPTOR.enum_types_by_name['ApprovalStatus'] = _APPROVALSTATUS
DESCRIPTOR.enum_types_by_name['CannedQuery'] = _CANNEDQUERY

Approval = _reflection.GeneratedProtocolMessageType('Approval', (_message.Message,), dict(
  DESCRIPTOR = _APPROVAL,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.Approval)
  ))
_sym_db.RegisterMessage(Approval)

Amendment = _reflection.GeneratedProtocolMessageType('Amendment', (_message.Message,), dict(
  DESCRIPTOR = _AMENDMENT,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.Amendment)
  ))
_sym_db.RegisterMessage(Amendment)

Attachment = _reflection.GeneratedProtocolMessageType('Attachment', (_message.Message,), dict(
  DESCRIPTOR = _ATTACHMENT,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.Attachment)
  ))
_sym_db.RegisterMessage(Attachment)

Comment = _reflection.GeneratedProtocolMessageType('Comment', (_message.Message,), dict(
  DESCRIPTOR = _COMMENT,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.Comment)
  ))
_sym_db.RegisterMessage(Comment)

FieldValue = _reflection.GeneratedProtocolMessageType('FieldValue', (_message.Message,), dict(
  DESCRIPTOR = _FIELDVALUE,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.FieldValue)
  ))
_sym_db.RegisterMessage(FieldValue)

Issue = _reflection.GeneratedProtocolMessageType('Issue', (_message.Message,), dict(
  DESCRIPTOR = _ISSUE,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.Issue)
  ))
_sym_db.RegisterMessage(Issue)

IssueSummary = _reflection.GeneratedProtocolMessageType('IssueSummary', (_message.Message,), dict(
  DESCRIPTOR = _ISSUESUMMARY,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.IssueSummary)
  ))
_sym_db.RegisterMessage(IssueSummary)

PhaseDef = _reflection.GeneratedProtocolMessageType('PhaseDef', (_message.Message,), dict(
  DESCRIPTOR = _PHASEDEF,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.PhaseDef)
  ))
_sym_db.RegisterMessage(PhaseDef)

PhaseRef = _reflection.GeneratedProtocolMessageType('PhaseRef', (_message.Message,), dict(
  DESCRIPTOR = _PHASEREF,
  __module__ = 'api.api_proto.issue_objects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.PhaseRef)
  ))
_sym_db.RegisterMessage(PhaseRef)


# @@protoc_insertion_point(module_scope)
