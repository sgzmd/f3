package main

import "fmt"

const (
	AuthorQueryTemplate = `
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

	AuthorQueryTemplateById = `
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

	SequenceQueryTemplate = `
select	
	f.SeqName,
	f.Authors,
	f.SeqId,
	(select count(ls.BookId) from libseq ls where ls.SeqId = f.SeqId) NumBooks
from 
	sequence_fts f 
where f.sequence_fts match ("%s*")
	`

	SequenceQueryTemplateById = `
select	
	f.SeqName,
	f.Authors,
	f.SeqId,
	(select count(ls.BookId) from libseq ls where ls.SeqId = f.SeqId) NumBooks
from 
	sequence_fts f 
where f.SeqId = %d
	`

	BooksByAuthorId = `
select b.Title, b.BookId
from libbook b,
     libavtor a
where b.BookId = a.BookId
  and b.Deleted = 0
  and a.AvtorId = %d
`

	BooksBySequenceId = `
select b.Title, b.BookId
from libbook b,
     libseq s
where b.BookId = s.BookId
  and b.Deleted = 0
  and s.SeqId = %d
`
)

func CreateAuthorSearchQuery(author string) string {
	return fmt.Sprintf(AuthorQueryTemplate, author)
}

func CreateSequenceSearchQuery(seq string) string {
	return fmt.Sprintf(SequenceQueryTemplate, seq)
}

func CreateAuthorByIdQuery(authorId int) string {
	return fmt.Sprintf(AuthorQueryTemplateById, authorId)
}

func CreateSequenceByIdQuery(seqId int) string {
	return fmt.Sprintf(SequenceQueryTemplateById, seqId)
}

func CreateGetBooksForAuthor(authorId int) string {
	return fmt.Sprintf(BooksByAuthorId, authorId)
}

func CreateGetBooksBySequenceId(seqId int) string {
	return fmt.Sprintf(BooksBySequenceId, seqId)
}
