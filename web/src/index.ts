import {SearchRequest} from '../gen/flibuserver/proto/flibustier_pb';

const world = 'world';

export function hello(world: string = "world"): string {
  var sr = new SearchRequest();
  sr.setSearchTerm("Test");

  return JSON.stringify(sr.toObject());
}

console.log(hello());