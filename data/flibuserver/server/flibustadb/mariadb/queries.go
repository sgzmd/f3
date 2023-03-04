package mariadb

import (
	"fmt"
	"strings"
)

const SearchAuthorsFtsMysql = `
select concat_name(lan.FirstName, lan.MiddleName, lan.LastName) authorName, la.AvtorId, COUNT(1) num
from libavtor la,
     libavtorname lan,
     libbook lb
where la.AvtorId = lan.AvtorId
  and la.BookId = lb.BookId
  and lb.Deleted != '1'
  and match(lan.FirstName, lan.LastName, lan.MiddleName) against('%s*' in boolean mode)
group by 1, 2
order by num desc;`

const SearchSeriesFtsMysql = `
select lsn.SeqId,
       SeqName,
       GROUP_CONCAT(DISTINCT concat_name(FirstName, MiddleName, LastName) order by LastName ) as AuthorName,
       (select count(ls.BookId) from libseq ls where ls.SeqId = lsn.SeqId)         NumBooks
from libseqname lsn,
     libseq ls,
     libavtor la,
     libavtorname lan
where lsn.SeqId = ls.SeqId
  and ls.BookId = la.BookId
  and la.AvtorId = lan.AvtorId
  and match(lsn.SeqName) against('%s*' in boolean mode)
GROUP BY lsn.SeqId
order by (select count(ls.BookId) from libseq ls where ls.SeqId = lsn.SeqId) desc;`

func MakeAuthorName(firstName, middleName, lastName string) string {
	return strings.TrimSpace(fmt.Sprintf("%s, %s %s", lastName, firstName, middleName))
}
