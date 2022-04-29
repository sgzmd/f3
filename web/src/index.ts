import {GlobalSearchRequest} from '../gen/nodejs/flibuserver/proto/v1/flibustier_pb';
import {FlibustierServiceClient} from "../gen/nodejs/flibuserver/proto/v1/flibustier_grpc_pb";
import * as grpc from '@grpc/grpc-js';


const world = 'world';

export function hello(world: string = "world"): string {
  var sr = new GlobalSearchRequest();
  sr.setSearchTerm("Маск");

  var client = new FlibustierServiceClient("172.23.22.238:9000", grpc.credentials.createInsecure() );

  return JSON.stringify(sr.toObject());
}

console.log(hello());