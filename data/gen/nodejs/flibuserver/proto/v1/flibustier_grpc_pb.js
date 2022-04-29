// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var flibuserver_proto_v1_flibustier_pb = require('../../../flibuserver/proto/v1/flibustier_pb.js');

function serialize_flibuserver_proto_v1_CheckUpdatesRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.CheckUpdatesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_CheckUpdatesRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_CheckUpdatesResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.CheckUpdatesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_CheckUpdatesResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_GetAuthorBooksRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.GetAuthorBooksRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_GetAuthorBooksRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_GetAuthorBooksResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.GetAuthorBooksResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_GetAuthorBooksResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_GetSeriesBooksRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.GetSeriesBooksRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_GetSeriesBooksRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_GetSeriesBooksResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.GetSeriesBooksResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_GetSeriesBooksResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_GlobalSearchRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.GlobalSearchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_GlobalSearchRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_GlobalSearchResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.GlobalSearchResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_GlobalSearchResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_ListTrackedEntriesRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.ListTrackedEntriesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_ListTrackedEntriesRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_ListTrackedEntriesResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.ListTrackedEntriesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_ListTrackedEntriesResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_TrackEntryRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.TrackEntryRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.TrackEntryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_TrackEntryRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.TrackEntryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_TrackEntryResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.TrackEntryResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.TrackEntryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_TrackEntryResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.TrackEntryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_UntrackEntryRequest(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.UntrackEntryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_UntrackEntryRequest(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibuserver_proto_v1_UntrackEntryResponse(arg) {
  if (!(arg instanceof flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse)) {
    throw new Error('Expected argument of type flibuserver.proto.v1.UntrackEntryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibuserver_proto_v1_UntrackEntryResponse(buffer_arg) {
  return flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var FlibustierServiceService = exports.FlibustierServiceService = {
  globalSearch: {
    path: '/flibuserver.proto.v1.FlibustierService/GlobalSearch',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse,
    requestSerialize: serialize_flibuserver_proto_v1_GlobalSearchRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_GlobalSearchRequest,
    responseSerialize: serialize_flibuserver_proto_v1_GlobalSearchResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_GlobalSearchResponse,
  },
  checkUpdates: {
    path: '/flibuserver.proto.v1.FlibustierService/CheckUpdates',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse,
    requestSerialize: serialize_flibuserver_proto_v1_CheckUpdatesRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_CheckUpdatesRequest,
    responseSerialize: serialize_flibuserver_proto_v1_CheckUpdatesResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_CheckUpdatesResponse,
  },
  getSeriesBooks: {
    path: '/flibuserver.proto.v1.FlibustierService/GetSeriesBooks',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse,
    requestSerialize: serialize_flibuserver_proto_v1_GetSeriesBooksRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_GetSeriesBooksRequest,
    responseSerialize: serialize_flibuserver_proto_v1_GetSeriesBooksResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_GetSeriesBooksResponse,
  },
  getAuthorBooks: {
    path: '/flibuserver.proto.v1.FlibustierService/GetAuthorBooks',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse,
    requestSerialize: serialize_flibuserver_proto_v1_GetAuthorBooksRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_GetAuthorBooksRequest,
    responseSerialize: serialize_flibuserver_proto_v1_GetAuthorBooksResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_GetAuthorBooksResponse,
  },
  trackEntry: {
    path: '/flibuserver.proto.v1.FlibustierService/TrackEntry',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse,
    requestSerialize: serialize_flibuserver_proto_v1_TrackEntryRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_TrackEntryRequest,
    responseSerialize: serialize_flibuserver_proto_v1_TrackEntryResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_TrackEntryResponse,
  },
  listTrackedEntries: {
    path: '/flibuserver.proto.v1.FlibustierService/ListTrackedEntries',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse,
    requestSerialize: serialize_flibuserver_proto_v1_ListTrackedEntriesRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_ListTrackedEntriesRequest,
    responseSerialize: serialize_flibuserver_proto_v1_ListTrackedEntriesResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_ListTrackedEntriesResponse,
  },
  untrackEntry: {
    path: '/flibuserver.proto.v1.FlibustierService/UntrackEntry',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest,
    responseType: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse,
    requestSerialize: serialize_flibuserver_proto_v1_UntrackEntryRequest,
    requestDeserialize: deserialize_flibuserver_proto_v1_UntrackEntryRequest,
    responseSerialize: serialize_flibuserver_proto_v1_UntrackEntryResponse,
    responseDeserialize: deserialize_flibuserver_proto_v1_UntrackEntryResponse,
  },
};

exports.FlibustierServiceClient = grpc.makeGenericClientConstructor(FlibustierServiceService);
