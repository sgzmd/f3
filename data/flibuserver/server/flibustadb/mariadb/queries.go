package mariadb

import (
	"fmt"
	"strings"
)

const SearchAuthorsFtsMysql = `
select lan.FirstName, lan.MiddleName, lan.LastName, la.AvtorId, COUNT(1) num
from libavtor la,
     libavtorname lan,
     libbook lb
where la.AvtorId = lan.AvtorId
  and la.BookId = lb.BookId
  and lb.Deleted != '1'
  and match(lan.FirstName, lan.LastName, lan.MiddleName) against('%s*' in boolean mode)
group by 1, 2, 3
order by num desc;`

func MakeAuthorName(firstName, middleName, lastName string) string {
	return strings.TrimSpace(fmt.Sprintf("%s, %s %s", lastName, firstName, middleName))
}
