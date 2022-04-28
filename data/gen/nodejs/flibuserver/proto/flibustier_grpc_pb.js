// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var flibuserver_proto_flibustier_pb = require('../../flibuserver/proto/flibustier_pb.js');

function serialize_flibustier_AuthorBooksRequest(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.AuthorBooksRequest)) {
    throw new Error('Expected argument of type flibustier.AuthorBooksRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_AuthorBooksRequest(buffer_arg) {
  return flibuserver_proto_flibustier_pb.AuthorBooksRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_EntityBookResponse(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.EntityBookResponse)) {
    throw new Error('Expected argument of type flibustier.EntityBookResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_EntityBookResponse(buffer_arg) {
  return flibuserver_proto_flibustier_pb.EntityBookResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_ListTrackedEntriesRequest(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest)) {
    throw new Error('Expected argument of type flibustier.ListTrackedEntriesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_ListTrackedEntriesRequest(buffer_arg) {
  return flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_ListTrackedEntriesResponse(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse)) {
    throw new Error('Expected argument of type flibustier.ListTrackedEntriesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_ListTrackedEntriesResponse(buffer_arg) {
  return flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_SearchRequest(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.SearchRequest)) {
    throw new Error('Expected argument of type flibustier.SearchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_SearchRequest(buffer_arg) {
  return flibuserver_proto_flibustier_pb.SearchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_SearchResponse(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.SearchResponse)) {
    throw new Error('Expected argument of type flibustier.SearchResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_SearchResponse(buffer_arg) {
  return flibuserver_proto_flibustier_pb.SearchResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_SequenceBooksRequest(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.SequenceBooksRequest)) {
    throw new Error('Expected argument of type flibustier.SequenceBooksRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_SequenceBooksRequest(buffer_arg) {
  return flibuserver_proto_flibustier_pb.SequenceBooksRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_TrackEntryResponse(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.TrackEntryResponse)) {
    throw new Error('Expected argument of type flibustier.TrackEntryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_TrackEntryResponse(buffer_arg) {
  return flibuserver_proto_flibustier_pb.TrackEntryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_TrackedEntry(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.TrackedEntry)) {
    throw new Error('Expected argument of type flibustier.TrackedEntry');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_TrackedEntry(buffer_arg) {
  return flibuserver_proto_flibustier_pb.TrackedEntry.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_TrackedEntryKey(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.TrackedEntryKey)) {
    throw new Error('Expected argument of type flibustier.TrackedEntryKey');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_TrackedEntryKey(buffer_arg) {
  return flibuserver_proto_flibustier_pb.TrackedEntryKey.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_UntrackEntryResponse(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.UntrackEntryResponse)) {
    throw new Error('Expected argument of type flibustier.UntrackEntryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_UntrackEntryResponse(buffer_arg) {
  return flibuserver_proto_flibustier_pb.UntrackEntryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_UpdateCheckRequest(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.UpdateCheckRequest)) {
    throw new Error('Expected argument of type flibustier.UpdateCheckRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_UpdateCheckRequest(buffer_arg) {
  return flibuserver_proto_flibustier_pb.UpdateCheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flibustier_UpdateCheckResponse(arg) {
  if (!(arg instanceof flibuserver_proto_flibustier_pb.UpdateCheckResponse)) {
    throw new Error('Expected argument of type flibustier.UpdateCheckResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flibustier_UpdateCheckResponse(buffer_arg) {
  return flibuserver_proto_flibustier_pb.UpdateCheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var FlibustierService = exports.FlibustierService = {
  globalSearch: {
    path: '/flibustier.Flibustier/GlobalSearch',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.SearchRequest,
    responseType: flibuserver_proto_flibustier_pb.SearchResponse,
    requestSerialize: serialize_flibustier_SearchRequest,
    requestDeserialize: deserialize_flibustier_SearchRequest,
    responseSerialize: serialize_flibustier_SearchResponse,
    responseDeserialize: deserialize_flibustier_SearchResponse,
  },
  checkUpdates: {
    path: '/flibustier.Flibustier/CheckUpdates',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.UpdateCheckRequest,
    responseType: flibuserver_proto_flibustier_pb.UpdateCheckResponse,
    requestSerialize: serialize_flibustier_UpdateCheckRequest,
    requestDeserialize: deserialize_flibustier_UpdateCheckRequest,
    responseSerialize: serialize_flibustier_UpdateCheckResponse,
    responseDeserialize: deserialize_flibustier_UpdateCheckResponse,
  },
  getSeriesBooks: {
    path: '/flibustier.Flibustier/GetSeriesBooks',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.SequenceBooksRequest,
    responseType: flibuserver_proto_flibustier_pb.EntityBookResponse,
    requestSerialize: serialize_flibustier_SequenceBooksRequest,
    requestDeserialize: deserialize_flibustier_SequenceBooksRequest,
    responseSerialize: serialize_flibustier_EntityBookResponse,
    responseDeserialize: deserialize_flibustier_EntityBookResponse,
  },
  getAuthorBooks: {
    path: '/flibustier.Flibustier/GetAuthorBooks',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.AuthorBooksRequest,
    responseType: flibuserver_proto_flibustier_pb.EntityBookResponse,
    requestSerialize: serialize_flibustier_AuthorBooksRequest,
    requestDeserialize: deserialize_flibustier_AuthorBooksRequest,
    responseSerialize: serialize_flibustier_EntityBookResponse,
    responseDeserialize: deserialize_flibustier_EntityBookResponse,
  },
  trackEntry: {
    path: '/flibustier.Flibustier/TrackEntry',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.TrackedEntry,
    responseType: flibuserver_proto_flibustier_pb.TrackEntryResponse,
    requestSerialize: serialize_flibustier_TrackedEntry,
    requestDeserialize: deserialize_flibustier_TrackedEntry,
    responseSerialize: serialize_flibustier_TrackEntryResponse,
    responseDeserialize: deserialize_flibustier_TrackEntryResponse,
  },
  listTrackedEntries: {
    path: '/flibustier.Flibustier/ListTrackedEntries',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest,
    responseType: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse,
    requestSerialize: serialize_flibustier_ListTrackedEntriesRequest,
    requestDeserialize: deserialize_flibustier_ListTrackedEntriesRequest,
    responseSerialize: serialize_flibustier_ListTrackedEntriesResponse,
    responseDeserialize: deserialize_flibustier_ListTrackedEntriesResponse,
  },
  untrackEntry: {
    path: '/flibustier.Flibustier/UntrackEntry',
    requestStream: false,
    responseStream: false,
    requestType: flibuserver_proto_flibustier_pb.TrackedEntryKey,
    responseType: flibuserver_proto_flibustier_pb.UntrackEntryResponse,
    requestSerialize: serialize_flibustier_TrackedEntryKey,
    requestDeserialize: deserialize_flibustier_TrackedEntryKey,
    responseSerialize: serialize_flibustier_UntrackEntryResponse,
    responseDeserialize: deserialize_flibustier_UntrackEntryResponse,
  },
};

exports.FlibustierClient = grpc.makeGenericClientConstructor(FlibustierService);
