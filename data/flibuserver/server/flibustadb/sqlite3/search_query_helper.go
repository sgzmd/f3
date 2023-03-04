package sqlite3

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
)

