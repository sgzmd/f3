package main

import "fmt"

const (
	AUTHOR_QUERY_TEMPLATE = `
select 
	a.authorName, 
	a.authorId,
	COUNT(1) as Count
from 
	author_fts a,
	libavtor la,
	libbook lb
where     
	a.author_fts match("%s*")
	and la.AvtorId = a.authorId
	and la.BookId = lb.BookId
	and lb.Deleted != '1'
GROUP BY 1,2;
	`

	AUTHOR_QUERY_TEMPLATE_BY_ID = `
select
	a.authorName,
	a.authorId,
	COUNT(1) as Count
from
	libavtor la,
	author_fts a,
	libbook lb
where
	la.AvtorId = a.authorId
	and la.BookId = lb.BookId
	and lb.Deleted != '1'
	and la.AvtorId = %d
GROUP BY 1,2;
	`

	SEQUENCE_QUERY_TEMPLATE = `
select	
	f.SeqName,
	f.Authors,
	f.SeqId,
	(select count(ls.BookId) from libseq ls where ls.SeqId = f.SeqId) NumBooks
from 
	sequence_fts f 
where f.sequence_fts match ("%s*")
	`

	SEQUENCE_QUERY_TEMPLATE_BY_ID = `
select	
	f.SeqName,
	f.Authors,
	f.SeqId,
	(select count(ls.BookId) from libseq ls where ls.SeqId = f.SeqId) NumBooks
from 
	libseq f 
where f.SeqId = %d
	`
)

func CreateAuthorSearchQuery(author string) string {
	return fmt.Sprintf(AUTHOR_QUERY_TEMPLATE, author)
}

func CreateSequenceSearchQuery(seq string) string {
	return fmt.Sprintf(SEQUENCE_QUERY_TEMPLATE, seq)
}

func CreateAuthorByIdQuery(authorId int) string {
	return fmt.Sprintf(AUTHOR_QUERY_TEMPLATE_BY_ID, authorId)
}

func CreateSequenceByIdQuery(seqId int) string {
	return fmt.Sprintf(SEQUENCE_QUERY_TEMPLATE_BY_ID, seqId)
}
