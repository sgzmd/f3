package sqlite3

import "fmt"

const (
	AuthorQueryTemplateSqlite = `
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

	AuthorQueryTemplateMysql = `
select lan.FirstName, lan.MiddleName, lan.LastName, COUNT(1) num
from libavtor la,
     libavtorname lan,
     libbook lb
where la.AvtorId = lan.AvtorId
  and la.BookId = lb.BookId
  and lb.Deleted != '1'
  and match(lan.FirstName, lan.LastName, lan.MiddleName) against('%s*' in boolean mode)
group by 1, 2, 3
order by num desc;
`

	AuthorQueryTemplateByIdSqlite = `
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

	SequenceQueryTemplateSqlite = `
select	
	f.SeqId,	
	f.SeqName,
	f.Authors,
	(select count(ls.BookId) from libseq ls where ls.SeqId = f.SeqId) NumBooks
from 
	sequence_fts f 
where f.sequence_fts match ("%s*")
	`

	SequenceQueryTemplateByIdSqlite = `
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
  and b.Deleted != '1'
  and a.AvtorId = %d
`

	BooksBySequenceId = `
select b.Title, b.BookId
from libbook b,
     libseq s
where b.BookId = s.BookId
  and b.Deleted != '1'
  and s.SeqId = %d
order by s.SeqNumb
`
)

func CreateAuthorSearchQuery(author string) string {
	return fmt.Sprintf(AuthorQueryTemplateSqlite, author)
}

func CreateSequenceSearchQuery(seq string) string {
	return fmt.Sprintf(SequenceQueryTemplateSqlite, seq)
}

func CreateAuthorByIdQuery(authorId int) string {
	return fmt.Sprintf(AuthorQueryTemplateByIdSqlite, authorId)
}

func CreateSequenceByIdQuery(seqId int) string {
	return fmt.Sprintf(SequenceQueryTemplateByIdSqlite, seqId)
}

func CreateGetBooksForAuthor(authorId int) string {
	return fmt.Sprintf(BooksByAuthorId, authorId)
}

func CreateGetBooksBySequenceId(seqId int) string {
	return fmt.Sprintf(BooksBySequenceId, seqId)
}
