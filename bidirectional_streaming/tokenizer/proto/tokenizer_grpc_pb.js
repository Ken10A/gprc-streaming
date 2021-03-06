// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var tokenizer_pb = require('./tokenizer_pb.js');

function serialize_practical_grpc_v1_TokenizeRequest(arg) {
  if (!(arg instanceof tokenizer_pb.TokenizeRequest)) {
    throw new Error('Expected argument of type practical.grpc.v1.TokenizeRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_practical_grpc_v1_TokenizeRequest(buffer_arg) {
  return tokenizer_pb.TokenizeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_practical_grpc_v1_TokenizeResponse(arg) {
  if (!(arg instanceof tokenizer_pb.TokenizeResponse)) {
    throw new Error('Expected argument of type practical.grpc.v1.TokenizeResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_practical_grpc_v1_TokenizeResponse(buffer_arg) {
  return tokenizer_pb.TokenizeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TokenizerService = exports.TokenizerService = {
  tokenize: {
    path: '/practical.grpc.v1.Tokenizer/Tokenize',
    requestStream: true,
    responseStream: true,
    requestType: tokenizer_pb.TokenizeRequest,
    responseType: tokenizer_pb.TokenizeResponse,
    requestSerialize: serialize_practical_grpc_v1_TokenizeRequest,
    requestDeserialize: deserialize_practical_grpc_v1_TokenizeRequest,
    responseSerialize: serialize_practical_grpc_v1_TokenizeResponse,
    responseDeserialize: deserialize_practical_grpc_v1_TokenizeResponse,
  },
};

exports.TokenizerClient = grpc.makeGenericClientConstructor(TokenizerService);
