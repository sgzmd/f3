// package: flibustier
// file: flibuserver/proto/flibustier.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as flibuserver_proto_flibustier_pb from "../../flibuserver/proto/flibustier_pb";

interface IFlibustierService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    globalSearch: IFlibustierService_IGlobalSearch;
    checkUpdates: IFlibustierService_ICheckUpdates;
    getSeriesBooks: IFlibustierService_IGetSeriesBooks;
    getAuthorBooks: IFlibustierService_IGetAuthorBooks;
    trackEntry: IFlibustierService_ITrackEntry;
    listTrackedEntries: IFlibustierService_IListTrackedEntries;
    untrackEntry: IFlibustierService_IUntrackEntry;
}

interface IFlibustierService_IGlobalSearch extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.SearchRequest, flibuserver_proto_flibustier_pb.SearchResponse> {
    path: "/flibustier.Flibustier/GlobalSearch";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.SearchRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.SearchRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.SearchResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.SearchResponse>;
}
interface IFlibustierService_ICheckUpdates extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.UpdateCheckRequest, flibuserver_proto_flibustier_pb.UpdateCheckResponse> {
    path: "/flibustier.Flibustier/CheckUpdates";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.UpdateCheckRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.UpdateCheckRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.UpdateCheckResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.UpdateCheckResponse>;
}
interface IFlibustierService_IGetSeriesBooks extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.SequenceBooksRequest, flibuserver_proto_flibustier_pb.EntityBookResponse> {
    path: "/flibustier.Flibustier/GetSeriesBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.SequenceBooksRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.SequenceBooksRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.EntityBookResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.EntityBookResponse>;
}
interface IFlibustierService_IGetAuthorBooks extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.AuthorBooksRequest, flibuserver_proto_flibustier_pb.EntityBookResponse> {
    path: "/flibustier.Flibustier/GetAuthorBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.AuthorBooksRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.AuthorBooksRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.EntityBookResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.EntityBookResponse>;
}
interface IFlibustierService_ITrackEntry extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.TrackedEntry, flibuserver_proto_flibustier_pb.TrackEntryResponse> {
    path: "/flibustier.Flibustier/TrackEntry";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.TrackedEntry>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.TrackedEntry>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.TrackEntryResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.TrackEntryResponse>;
}
interface IFlibustierService_IListTrackedEntries extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse> {
    path: "/flibustier.Flibustier/ListTrackedEntries";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse>;
}
interface IFlibustierService_IUntrackEntry extends grpc.MethodDefinition<flibuserver_proto_flibustier_pb.TrackedEntryKey, flibuserver_proto_flibustier_pb.UntrackEntryResponse> {
    path: "/flibustier.Flibustier/UntrackEntry";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.TrackedEntryKey>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.TrackedEntryKey>;
    responseSerialize: grpc.serialize<flibuserver_proto_flibustier_pb.UntrackEntryResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_flibustier_pb.UntrackEntryResponse>;
}

export const FlibustierService: IFlibustierService;

export interface IFlibustierServer extends grpc.UntypedServiceImplementation {
    globalSearch: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.SearchRequest, flibuserver_proto_flibustier_pb.SearchResponse>;
    checkUpdates: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.UpdateCheckRequest, flibuserver_proto_flibustier_pb.UpdateCheckResponse>;
    getSeriesBooks: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.SequenceBooksRequest, flibuserver_proto_flibustier_pb.EntityBookResponse>;
    getAuthorBooks: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.AuthorBooksRequest, flibuserver_proto_flibustier_pb.EntityBookResponse>;
    trackEntry: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.TrackedEntry, flibuserver_proto_flibustier_pb.TrackEntryResponse>;
    listTrackedEntries: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse>;
    untrackEntry: grpc.handleUnaryCall<flibuserver_proto_flibustier_pb.TrackedEntryKey, flibuserver_proto_flibustier_pb.UntrackEntryResponse>;
}

export interface IFlibustierClient {
    globalSearch(request: flibuserver_proto_flibustier_pb.SearchRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    globalSearch(request: flibuserver_proto_flibustier_pb.SearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    globalSearch(request: flibuserver_proto_flibustier_pb.SearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    checkUpdates(request: flibuserver_proto_flibustier_pb.UpdateCheckRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
    checkUpdates(request: flibuserver_proto_flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
    checkUpdates(request: flibuserver_proto_flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
    getSeriesBooks(request: flibuserver_proto_flibustier_pb.SequenceBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    getSeriesBooks(request: flibuserver_proto_flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    getSeriesBooks(request: flibuserver_proto_flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    getAuthorBooks(request: flibuserver_proto_flibustier_pb.AuthorBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    getAuthorBooks(request: flibuserver_proto_flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    getAuthorBooks(request: flibuserver_proto_flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    trackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntry, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    trackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntry, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    trackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntry, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    listTrackedEntries(request: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    listTrackedEntries(request: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    listTrackedEntries(request: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    untrackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntryKey, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    untrackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    untrackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
}

export class FlibustierClient extends grpc.Client implements IFlibustierClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public globalSearch(request: flibuserver_proto_flibustier_pb.SearchRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    public globalSearch(request: flibuserver_proto_flibustier_pb.SearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    public globalSearch(request: flibuserver_proto_flibustier_pb.SearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
    public checkUpdates(request: flibuserver_proto_flibustier_pb.UpdateCheckRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
    public checkUpdates(request: flibuserver_proto_flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
    public checkUpdates(request: flibuserver_proto_flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
    public getSeriesBooks(request: flibuserver_proto_flibustier_pb.SequenceBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    public getSeriesBooks(request: flibuserver_proto_flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    public getSeriesBooks(request: flibuserver_proto_flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    public getAuthorBooks(request: flibuserver_proto_flibustier_pb.AuthorBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    public getAuthorBooks(request: flibuserver_proto_flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    public getAuthorBooks(request: flibuserver_proto_flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
    public trackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntry, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    public trackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntry, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    public trackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntry, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    public listTrackedEntries(request: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    public listTrackedEntries(request: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    public listTrackedEntries(request: flibuserver_proto_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    public untrackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntryKey, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    public untrackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    public untrackEntry(request: flibuserver_proto_flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
}
