// package: flibuserver.proto.v1
// file: flibuserver/proto/v1/flibustier.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as flibuserver_proto_v1_flibustier_pb from "../../../flibuserver/proto/v1/flibustier_pb";

interface IFlibustierServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    globalSearch: IFlibustierServiceService_IGlobalSearch;
    checkUpdates: IFlibustierServiceService_ICheckUpdates;
    getSeriesBooks: IFlibustierServiceService_IGetSeriesBooks;
    getAuthorBooks: IFlibustierServiceService_IGetAuthorBooks;
    trackEntry: IFlibustierServiceService_ITrackEntry;
    listTrackedEntries: IFlibustierServiceService_IListTrackedEntries;
    untrackEntry: IFlibustierServiceService_IUntrackEntry;
}

interface IFlibustierServiceService_IGlobalSearch extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/GlobalSearch";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse>;
}
interface IFlibustierServiceService_ICheckUpdates extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/CheckUpdates";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse>;
}
interface IFlibustierServiceService_IGetSeriesBooks extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/GetSeriesBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse>;
}
interface IFlibustierServiceService_IGetAuthorBooks extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/GetAuthorBooks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse>;
}
interface IFlibustierServiceService_ITrackEntry extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, flibuserver_proto_v1_flibustier_pb.TrackEntryResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/TrackEntry";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.TrackEntryRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.TrackEntryRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.TrackEntryResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.TrackEntryResponse>;
}
interface IFlibustierServiceService_IListTrackedEntries extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/ListTrackedEntries";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse>;
}
interface IFlibustierServiceService_IUntrackEntry extends grpc.MethodDefinition<flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse> {
    path: "/flibuserver.proto.v1.FlibustierService/UntrackEntry";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest>;
    requestDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest>;
    responseSerialize: grpc.serialize<flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse>;
    responseDeserialize: grpc.deserialize<flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse>;
}

export const FlibustierServiceService: IFlibustierServiceService;

export interface IFlibustierServiceServer extends grpc.UntypedServiceImplementation {
    globalSearch: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse>;
    checkUpdates: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse>;
    getSeriesBooks: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse>;
    getAuthorBooks: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse>;
    trackEntry: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, flibuserver_proto_v1_flibustier_pb.TrackEntryResponse>;
    listTrackedEntries: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse>;
    untrackEntry: grpc.handleUnaryCall<flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse>;
}

export interface IFlibustierServiceClient {
    globalSearch(request: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse) => void): grpc.ClientUnaryCall;
    globalSearch(request: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse) => void): grpc.ClientUnaryCall;
    globalSearch(request: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse) => void): grpc.ClientUnaryCall;
    checkUpdates(request: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse) => void): grpc.ClientUnaryCall;
    checkUpdates(request: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse) => void): grpc.ClientUnaryCall;
    checkUpdates(request: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse) => void): grpc.ClientUnaryCall;
    getSeriesBooks(request: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse) => void): grpc.ClientUnaryCall;
    getSeriesBooks(request: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse) => void): grpc.ClientUnaryCall;
    getSeriesBooks(request: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse) => void): grpc.ClientUnaryCall;
    getAuthorBooks(request: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse) => void): grpc.ClientUnaryCall;
    getAuthorBooks(request: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse) => void): grpc.ClientUnaryCall;
    getAuthorBooks(request: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse) => void): grpc.ClientUnaryCall;
    trackEntry(request: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    trackEntry(request: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    trackEntry(request: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    listTrackedEntries(request: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    listTrackedEntries(request: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    listTrackedEntries(request: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    untrackEntry(request: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    untrackEntry(request: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    untrackEntry(request: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
}

export class FlibustierServiceClient extends grpc.Client implements IFlibustierServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public globalSearch(request: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse) => void): grpc.ClientUnaryCall;
    public globalSearch(request: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse) => void): grpc.ClientUnaryCall;
    public globalSearch(request: flibuserver_proto_v1_flibustier_pb.GlobalSearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GlobalSearchResponse) => void): grpc.ClientUnaryCall;
    public checkUpdates(request: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse) => void): grpc.ClientUnaryCall;
    public checkUpdates(request: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse) => void): grpc.ClientUnaryCall;
    public checkUpdates(request: flibuserver_proto_v1_flibustier_pb.CheckUpdatesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.CheckUpdatesResponse) => void): grpc.ClientUnaryCall;
    public getSeriesBooks(request: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse) => void): grpc.ClientUnaryCall;
    public getSeriesBooks(request: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse) => void): grpc.ClientUnaryCall;
    public getSeriesBooks(request: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetSeriesBooksResponse) => void): grpc.ClientUnaryCall;
    public getAuthorBooks(request: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse) => void): grpc.ClientUnaryCall;
    public getAuthorBooks(request: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse) => void): grpc.ClientUnaryCall;
    public getAuthorBooks(request: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.GetAuthorBooksResponse) => void): grpc.ClientUnaryCall;
    public trackEntry(request: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    public trackEntry(request: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    public trackEntry(request: flibuserver_proto_v1_flibustier_pb.TrackEntryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
    public listTrackedEntries(request: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    public listTrackedEntries(request: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    public listTrackedEntries(request: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
    public untrackEntry(request: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    public untrackEntry(request: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
    public untrackEntry(request: flibuserver_proto_v1_flibustier_pb.UntrackEntryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibuserver_proto_v1_flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
}
