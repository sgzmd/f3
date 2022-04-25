// package: flibustier
// file: flibustier.proto

import * as grpc from 'grpc';
import * as flibustier_pb from './flibustier_pb';

interface IFlibustierService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  globalSearch: IFlibustierService_IGlobalSearch;
  checkUpdates: IFlibustierService_ICheckUpdates;
  getSeriesBooks: IFlibustierService_IGetSeriesBooks;
  getAuthorBooks: IFlibustierService_IGetAuthorBooks;
  trackEntry: IFlibustierService_ITrackEntry;
  listTrackedEntries: IFlibustierService_IListTrackedEntries;
  untrackEntry: IFlibustierService_IUntrackEntry;
}

interface IFlibustierService_IGlobalSearch extends grpc.MethodDefinition<flibustier_pb.SearchRequest, flibustier_pb.SearchResponse> {
  path: '/flibustier.Flibustier/GlobalSearch'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.SearchRequest>;
  requestDeserialize: grpc.deserialize<flibustier_pb.SearchRequest>;
  responseSerialize: grpc.serialize<flibustier_pb.SearchResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.SearchResponse>;
}

interface IFlibustierService_ICheckUpdates extends grpc.MethodDefinition<flibustier_pb.UpdateCheckRequest, flibustier_pb.UpdateCheckResponse> {
  path: '/flibustier.Flibustier/CheckUpdates'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.UpdateCheckRequest>;
  requestDeserialize: grpc.deserialize<flibustier_pb.UpdateCheckRequest>;
  responseSerialize: grpc.serialize<flibustier_pb.UpdateCheckResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.UpdateCheckResponse>;
}

interface IFlibustierService_IGetSeriesBooks extends grpc.MethodDefinition<flibustier_pb.SequenceBooksRequest, flibustier_pb.EntityBookResponse> {
  path: '/flibustier.Flibustier/GetSeriesBooks'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.SequenceBooksRequest>;
  requestDeserialize: grpc.deserialize<flibustier_pb.SequenceBooksRequest>;
  responseSerialize: grpc.serialize<flibustier_pb.EntityBookResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.EntityBookResponse>;
}

interface IFlibustierService_IGetAuthorBooks extends grpc.MethodDefinition<flibustier_pb.AuthorBooksRequest, flibustier_pb.EntityBookResponse> {
  path: '/flibustier.Flibustier/GetAuthorBooks'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.AuthorBooksRequest>;
  requestDeserialize: grpc.deserialize<flibustier_pb.AuthorBooksRequest>;
  responseSerialize: grpc.serialize<flibustier_pb.EntityBookResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.EntityBookResponse>;
}

interface IFlibustierService_ITrackEntry extends grpc.MethodDefinition<flibustier_pb.TrackedEntry, flibustier_pb.TrackEntryResponse> {
  path: '/flibustier.Flibustier/TrackEntry'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.TrackedEntry>;
  requestDeserialize: grpc.deserialize<flibustier_pb.TrackedEntry>;
  responseSerialize: grpc.serialize<flibustier_pb.TrackEntryResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.TrackEntryResponse>;
}

interface IFlibustierService_IListTrackedEntries extends grpc.MethodDefinition<flibustier_pb.ListTrackedEntriesRequest, flibustier_pb.ListTrackedEntriesResponse> {
  path: '/flibustier.Flibustier/ListTrackedEntries'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.ListTrackedEntriesRequest>;
  requestDeserialize: grpc.deserialize<flibustier_pb.ListTrackedEntriesRequest>;
  responseSerialize: grpc.serialize<flibustier_pb.ListTrackedEntriesResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.ListTrackedEntriesResponse>;
}

interface IFlibustierService_IUntrackEntry extends grpc.MethodDefinition<flibustier_pb.TrackedEntryKey, flibustier_pb.UntrackEntryResponse> {
  path: '/flibustier.Flibustier/UntrackEntry'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<flibustier_pb.TrackedEntryKey>;
  requestDeserialize: grpc.deserialize<flibustier_pb.TrackedEntryKey>;
  responseSerialize: grpc.serialize<flibustier_pb.UntrackEntryResponse>;
  responseDeserialize: grpc.deserialize<flibustier_pb.UntrackEntryResponse>;
}

export const FlibustierService: IFlibustierService;
export interface IFlibustierServer extends grpc.UntypedServiceImplementation {
  globalSearch: grpc.handleUnaryCall<flibustier_pb.SearchRequest, flibustier_pb.SearchResponse>;
  checkUpdates: grpc.handleUnaryCall<flibustier_pb.UpdateCheckRequest, flibustier_pb.UpdateCheckResponse>;
  getSeriesBooks: grpc.handleUnaryCall<flibustier_pb.SequenceBooksRequest, flibustier_pb.EntityBookResponse>;
  getAuthorBooks: grpc.handleUnaryCall<flibustier_pb.AuthorBooksRequest, flibustier_pb.EntityBookResponse>;
  trackEntry: grpc.handleUnaryCall<flibustier_pb.TrackedEntry, flibustier_pb.TrackEntryResponse>;
  listTrackedEntries: grpc.handleUnaryCall<flibustier_pb.ListTrackedEntriesRequest, flibustier_pb.ListTrackedEntriesResponse>;
  untrackEntry: grpc.handleUnaryCall<flibustier_pb.TrackedEntryKey, flibustier_pb.UntrackEntryResponse>;
}

export interface IFlibustierClient {
  globalSearch(request: flibustier_pb.SearchRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
  globalSearch(request: flibustier_pb.SearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
  globalSearch(request: flibustier_pb.SearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
  checkUpdates(request: flibustier_pb.UpdateCheckRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
  checkUpdates(request: flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
  checkUpdates(request: flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
  getSeriesBooks(request: flibustier_pb.SequenceBooksRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  getSeriesBooks(request: flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  getSeriesBooks(request: flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  getAuthorBooks(request: flibustier_pb.AuthorBooksRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  getAuthorBooks(request: flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  getAuthorBooks(request: flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  trackEntry(request: flibustier_pb.TrackedEntry, callback: (error: grpc.ServiceError | null, response: flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
  trackEntry(request: flibustier_pb.TrackedEntry, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
  trackEntry(request: flibustier_pb.TrackedEntry, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
  listTrackedEntries(request: flibustier_pb.ListTrackedEntriesRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
  listTrackedEntries(request: flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
  listTrackedEntries(request: flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
  untrackEntry(request: flibustier_pb.TrackedEntryKey, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
  untrackEntry(request: flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
  untrackEntry(request: flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
}

export class FlibustierClient extends grpc.Client implements IFlibustierClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
  public globalSearch(request: flibustier_pb.SearchRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
  public globalSearch(request: flibustier_pb.SearchRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
  public globalSearch(request: flibustier_pb.SearchRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.SearchResponse) => void): grpc.ClientUnaryCall;
  public checkUpdates(request: flibustier_pb.UpdateCheckRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
  public checkUpdates(request: flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
  public checkUpdates(request: flibustier_pb.UpdateCheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UpdateCheckResponse) => void): grpc.ClientUnaryCall;
  public getSeriesBooks(request: flibustier_pb.SequenceBooksRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  public getSeriesBooks(request: flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  public getSeriesBooks(request: flibustier_pb.SequenceBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  public getAuthorBooks(request: flibustier_pb.AuthorBooksRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  public getAuthorBooks(request: flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  public getAuthorBooks(request: flibustier_pb.AuthorBooksRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.EntityBookResponse) => void): grpc.ClientUnaryCall;
  public trackEntry(request: flibustier_pb.TrackedEntry, callback: (error: grpc.ServiceError | null, response: flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
  public trackEntry(request: flibustier_pb.TrackedEntry, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
  public trackEntry(request: flibustier_pb.TrackedEntry, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.TrackEntryResponse) => void): grpc.ClientUnaryCall;
  public listTrackedEntries(request: flibustier_pb.ListTrackedEntriesRequest, callback: (error: grpc.ServiceError | null, response: flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
  public listTrackedEntries(request: flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
  public listTrackedEntries(request: flibustier_pb.ListTrackedEntriesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.ListTrackedEntriesResponse) => void): grpc.ClientUnaryCall;
  public untrackEntry(request: flibustier_pb.TrackedEntryKey, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
  public untrackEntry(request: flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
  public untrackEntry(request: flibustier_pb.TrackedEntryKey, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: flibustier_pb.UntrackEntryResponse) => void): grpc.ClientUnaryCall;
}

