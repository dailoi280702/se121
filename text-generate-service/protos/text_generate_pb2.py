# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: text_generate.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13text_generate.proto\x12\rtext_generate\"\x8c\x01\n\x11GenerateReviewReq\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x62rand\x18\x02 \x01(\t\x12\x0e\n\x06series\x18\x03 \x01(\t\x12\x12\n\nhorsePower\x18\x04 \x01(\x05\x12\x0e\n\x06torque\x18\x05 \x01(\x05\x12\x14\n\x0ctransmission\x18\x06 \x01(\t\x12\x10\n\x08\x66uelType\x18\x07 \x01(\t\";\n\x1cGenerateBlogSummarizationReq\x12\r\n\x05title\x18\x01 \x01(\t\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t\"\x19\n\tResString\x12\x0c\n\x04text\x18\x01 \x01(\t2\xca\x01\n\x13TextGenerateService\x12O\n\x11GenerateCarReview\x12 .text_generate.GenerateReviewReq\x1a\x18.text_generate.ResString\x12\x62\n\x19GenerateBlogSummarization\x12+.text_generate.GenerateBlogSummarizationReq\x1a\x18.text_generate.ResStringb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'text_generate_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _GENERATEREVIEWREQ._serialized_start=39
  _GENERATEREVIEWREQ._serialized_end=179
  _GENERATEBLOGSUMMARIZATIONREQ._serialized_start=181
  _GENERATEBLOGSUMMARIZATIONREQ._serialized_end=240
  _RESSTRING._serialized_start=242
  _RESSTRING._serialized_end=267
  _TEXTGENERATESERVICE._serialized_start=270
  _TEXTGENERATESERVICE._serialized_end=472
# @@protoc_insertion_point(module_scope)
