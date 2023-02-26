package flibustadb

import pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"

type FlibustaDb interface {
	SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error)
	SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error)
	GetAuthorBooks(authorId int64) ([]*pb.Book, error)
	GetSeriesBooks(seriesId int64) ([]*pb.Book, error)

	GetAuthorName(authorId int64) (pb.AuthorName, error)
	GetSequenceName(seqId int64) (string, error)

	GetBookAuthor(bookId int64) (pb.AuthorName, error)

	Close() error
}
