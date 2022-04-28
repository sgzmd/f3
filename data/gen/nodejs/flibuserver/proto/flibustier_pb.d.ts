// package: flibustier
// file: flibuserver/proto/flibustier.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class SearchRequest extends jspb.Message { 
    getSearchTerm(): string;
    setSearchTerm(value: string): SearchRequest;
    getEntryTypeFilter(): EntryType;
    setEntryTypeFilter(value: EntryType): SearchRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SearchRequest): SearchRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchRequest;
    static deserializeBinaryFromReader(message: SearchRequest, reader: jspb.BinaryReader): SearchRequest;
}

export namespace SearchRequest {
    export type AsObject = {
        searchTerm: string,
        entryTypeFilter: EntryType,
    }
}

export class FoundEntry extends jspb.Message { 
    getEntryType(): EntryType;
    setEntryType(value: EntryType): FoundEntry;
    getEntryName(): string;
    setEntryName(value: string): FoundEntry;
    getAuthor(): string;
    setAuthor(value: string): FoundEntry;
    getEntryId(): number;
    setEntryId(value: number): FoundEntry;
    getNumEntities(): number;
    setNumEntities(value: number): FoundEntry;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FoundEntry.AsObject;
    static toObject(includeInstance: boolean, msg: FoundEntry): FoundEntry.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FoundEntry, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FoundEntry;
    static deserializeBinaryFromReader(message: FoundEntry, reader: jspb.BinaryReader): FoundEntry;
}

export namespace FoundEntry {
    export type AsObject = {
        entryType: EntryType,
        entryName: string,
        author: string,
        entryId: number,
        numEntities: number,
    }
}

export class SearchResponse extends jspb.Message { 

    hasOriginalRequest(): boolean;
    clearOriginalRequest(): void;
    getOriginalRequest(): SearchRequest | undefined;
    setOriginalRequest(value?: SearchRequest): SearchResponse;
    clearEntryList(): void;
    getEntryList(): Array<FoundEntry>;
    setEntryList(value: Array<FoundEntry>): SearchResponse;
    addEntry(value?: FoundEntry, index?: number): FoundEntry;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SearchResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SearchResponse): SearchResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SearchResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SearchResponse;
    static deserializeBinaryFromReader(message: SearchResponse, reader: jspb.BinaryReader): SearchResponse;
}

export namespace SearchResponse {
    export type AsObject = {
        originalRequest?: SearchRequest.AsObject,
        entryList: Array<FoundEntry.AsObject>,
    }
}

export class Book extends jspb.Message { 
    getBookName(): string;
    setBookName(value: string): Book;
    getBookId(): number;
    setBookId(value: number): Book;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Book.AsObject;
    static toObject(includeInstance: boolean, msg: Book): Book.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Book, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Book;
    static deserializeBinaryFromReader(message: Book, reader: jspb.BinaryReader): Book;
}

export namespace Book {
    export type AsObject = {
        bookName: string,
        bookId: number,
    }
}

export class TrackedEntry extends jspb.Message { 
    getEntryType(): EntryType;
    setEntryType(value: EntryType): TrackedEntry;
    getEntryName(): string;
    setEntryName(value: string): TrackedEntry;
    getEntryId(): number;
    setEntryId(value: number): TrackedEntry;
    getNumEntries(): number;
    setNumEntries(value: number): TrackedEntry;
    getUserId(): string;
    setUserId(value: string): TrackedEntry;
    clearBookList(): void;
    getBookList(): Array<Book>;
    setBookList(value: Array<Book>): TrackedEntry;
    addBook(value?: Book, index?: number): Book;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackedEntry.AsObject;
    static toObject(includeInstance: boolean, msg: TrackedEntry): TrackedEntry.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackedEntry, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackedEntry;
    static deserializeBinaryFromReader(message: TrackedEntry, reader: jspb.BinaryReader): TrackedEntry;
}

export namespace TrackedEntry {
    export type AsObject = {
        entryType: EntryType,
        entryName: string,
        entryId: number,
        numEntries: number,
        userId: string,
        bookList: Array<Book.AsObject>,
    }
}

export class UpdateRequired extends jspb.Message { 

    hasTrackedEntry(): boolean;
    clearTrackedEntry(): void;
    getTrackedEntry(): TrackedEntry | undefined;
    setTrackedEntry(value?: TrackedEntry): UpdateRequired;
    getNewNumEntries(): number;
    setNewNumEntries(value: number): UpdateRequired;
    clearNewBookList(): void;
    getNewBookList(): Array<Book>;
    setNewBookList(value: Array<Book>): UpdateRequired;
    addNewBook(value?: Book, index?: number): Book;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRequired.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRequired): UpdateRequired.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRequired, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRequired;
    static deserializeBinaryFromReader(message: UpdateRequired, reader: jspb.BinaryReader): UpdateRequired;
}

export namespace UpdateRequired {
    export type AsObject = {
        trackedEntry?: TrackedEntry.AsObject,
        newNumEntries: number,
        newBookList: Array<Book.AsObject>,
    }
}

export class UpdateCheckRequest extends jspb.Message { 
    clearTrackedEntryList(): void;
    getTrackedEntryList(): Array<TrackedEntry>;
    setTrackedEntryList(value: Array<TrackedEntry>): UpdateCheckRequest;
    addTrackedEntry(value?: TrackedEntry, index?: number): TrackedEntry;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateCheckRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateCheckRequest): UpdateCheckRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateCheckRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateCheckRequest;
    static deserializeBinaryFromReader(message: UpdateCheckRequest, reader: jspb.BinaryReader): UpdateCheckRequest;
}

export namespace UpdateCheckRequest {
    export type AsObject = {
        trackedEntryList: Array<TrackedEntry.AsObject>,
    }
}

export class UpdateCheckResponse extends jspb.Message { 
    clearUpdateRequiredList(): void;
    getUpdateRequiredList(): Array<UpdateRequired>;
    setUpdateRequiredList(value: Array<UpdateRequired>): UpdateCheckResponse;
    addUpdateRequired(value?: UpdateRequired, index?: number): UpdateRequired;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateCheckResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateCheckResponse): UpdateCheckResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateCheckResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateCheckResponse;
    static deserializeBinaryFromReader(message: UpdateCheckResponse, reader: jspb.BinaryReader): UpdateCheckResponse;
}

export namespace UpdateCheckResponse {
    export type AsObject = {
        updateRequiredList: Array<UpdateRequired.AsObject>,
    }
}

export class SequenceBooksRequest extends jspb.Message { 
    getSequenceId(): number;
    setSequenceId(value: number): SequenceBooksRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SequenceBooksRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SequenceBooksRequest): SequenceBooksRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SequenceBooksRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SequenceBooksRequest;
    static deserializeBinaryFromReader(message: SequenceBooksRequest, reader: jspb.BinaryReader): SequenceBooksRequest;
}

export namespace SequenceBooksRequest {
    export type AsObject = {
        sequenceId: number,
    }
}

export class AuthorBooksRequest extends jspb.Message { 
    getAuthorId(): number;
    setAuthorId(value: number): AuthorBooksRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthorBooksRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AuthorBooksRequest): AuthorBooksRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthorBooksRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthorBooksRequest;
    static deserializeBinaryFromReader(message: AuthorBooksRequest, reader: jspb.BinaryReader): AuthorBooksRequest;
}

export namespace AuthorBooksRequest {
    export type AsObject = {
        authorId: number,
    }
}

export class EntityBookResponse extends jspb.Message { 
    getEntityId(): number;
    setEntityId(value: number): EntityBookResponse;
    clearBookList(): void;
    getBookList(): Array<Book>;
    setBookList(value: Array<Book>): EntityBookResponse;
    addBook(value?: Book, index?: number): Book;

    hasEntityName(): boolean;
    clearEntityName(): void;
    getEntityName(): EntityName | undefined;
    setEntityName(value?: EntityName): EntityBookResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EntityBookResponse.AsObject;
    static toObject(includeInstance: boolean, msg: EntityBookResponse): EntityBookResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EntityBookResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EntityBookResponse;
    static deserializeBinaryFromReader(message: EntityBookResponse, reader: jspb.BinaryReader): EntityBookResponse;
}

export namespace EntityBookResponse {
    export type AsObject = {
        entityId: number,
        bookList: Array<Book.AsObject>,
        entityName?: EntityName.AsObject,
    }
}

export class AuthorName extends jspb.Message { 
    getFirstName(): string;
    setFirstName(value: string): AuthorName;
    getMiddleName(): string;
    setMiddleName(value: string): AuthorName;
    getLastName(): string;
    setLastName(value: string): AuthorName;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthorName.AsObject;
    static toObject(includeInstance: boolean, msg: AuthorName): AuthorName.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthorName, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthorName;
    static deserializeBinaryFromReader(message: AuthorName, reader: jspb.BinaryReader): AuthorName;
}

export namespace AuthorName {
    export type AsObject = {
        firstName: string,
        middleName: string,
        lastName: string,
    }
}

export class EntityName extends jspb.Message { 

    hasAuthorName(): boolean;
    clearAuthorName(): void;
    getAuthorName(): AuthorName | undefined;
    setAuthorName(value?: AuthorName): EntityName;

    hasSequenceName(): boolean;
    clearSequenceName(): void;
    getSequenceName(): string;
    setSequenceName(value: string): EntityName;

    getNameCase(): EntityName.NameCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EntityName.AsObject;
    static toObject(includeInstance: boolean, msg: EntityName): EntityName.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EntityName, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EntityName;
    static deserializeBinaryFromReader(message: EntityName, reader: jspb.BinaryReader): EntityName;
}

export namespace EntityName {
    export type AsObject = {
        authorName?: AuthorName.AsObject,
        sequenceName: string,
    }

    export enum NameCase {
        NAME_NOT_SET = 0,
        AUTHOR_NAME = 1,
        SEQUENCE_NAME = 2,
    }

}

export class TrackedEntryKey extends jspb.Message { 
    getEntityType(): EntryType;
    setEntityType(value: EntryType): TrackedEntryKey;
    getEntityId(): number;
    setEntityId(value: number): TrackedEntryKey;
    getUserId(): string;
    setUserId(value: string): TrackedEntryKey;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackedEntryKey.AsObject;
    static toObject(includeInstance: boolean, msg: TrackedEntryKey): TrackedEntryKey.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackedEntryKey, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackedEntryKey;
    static deserializeBinaryFromReader(message: TrackedEntryKey, reader: jspb.BinaryReader): TrackedEntryKey;
}

export namespace TrackedEntryKey {
    export type AsObject = {
        entityType: EntryType,
        entityId: number,
        userId: string,
    }
}

export class TrackEntryResponse extends jspb.Message { 

    hasKey(): boolean;
    clearKey(): void;
    getKey(): TrackedEntryKey | undefined;
    setKey(value?: TrackedEntryKey): TrackEntryResponse;
    getResult(): TrackEntryResult;
    setResult(value: TrackEntryResult): TrackEntryResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackEntryResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TrackEntryResponse): TrackEntryResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackEntryResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackEntryResponse;
    static deserializeBinaryFromReader(message: TrackEntryResponse, reader: jspb.BinaryReader): TrackEntryResponse;
}

export namespace TrackEntryResponse {
    export type AsObject = {
        key?: TrackedEntryKey.AsObject,
        result: TrackEntryResult,
    }
}

export class ListTrackedEntriesRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): ListTrackedEntriesRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListTrackedEntriesRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListTrackedEntriesRequest): ListTrackedEntriesRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListTrackedEntriesRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListTrackedEntriesRequest;
    static deserializeBinaryFromReader(message: ListTrackedEntriesRequest, reader: jspb.BinaryReader): ListTrackedEntriesRequest;
}

export namespace ListTrackedEntriesRequest {
    export type AsObject = {
        userId: string,
    }
}

export class ListTrackedEntriesResponse extends jspb.Message { 
    clearEntryList(): void;
    getEntryList(): Array<TrackedEntry>;
    setEntryList(value: Array<TrackedEntry>): ListTrackedEntriesResponse;
    addEntry(value?: TrackedEntry, index?: number): TrackedEntry;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListTrackedEntriesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListTrackedEntriesResponse): ListTrackedEntriesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListTrackedEntriesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListTrackedEntriesResponse;
    static deserializeBinaryFromReader(message: ListTrackedEntriesResponse, reader: jspb.BinaryReader): ListTrackedEntriesResponse;
}

export namespace ListTrackedEntriesResponse {
    export type AsObject = {
        entryList: Array<TrackedEntry.AsObject>,
    }
}

export class UntrackEntryResponse extends jspb.Message { 

    hasKey(): boolean;
    clearKey(): void;
    getKey(): TrackedEntryKey | undefined;
    setKey(value?: TrackedEntryKey): UntrackEntryResponse;
    getResult(): UntrackEntryResult;
    setResult(value: UntrackEntryResult): UntrackEntryResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UntrackEntryResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UntrackEntryResponse): UntrackEntryResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UntrackEntryResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UntrackEntryResponse;
    static deserializeBinaryFromReader(message: UntrackEntryResponse, reader: jspb.BinaryReader): UntrackEntryResponse;
}

export namespace UntrackEntryResponse {
    export type AsObject = {
        key?: TrackedEntryKey.AsObject,
        result: UntrackEntryResult,
    }
}

export enum EntryType {
    UNKNOWN = 0,
    SERIES = 1,
    AUTHOR = 2,
    BOOK = 3,
}

export enum TrackEntryResult {
    TRACK_OK = 0,
    TRACK_ALREADY_TRACKED = 1,
}

export enum UntrackEntryResult {
    UNTRACK_OK = 0,
    UNTRACK_NOT_TRACKED = 1,
}
