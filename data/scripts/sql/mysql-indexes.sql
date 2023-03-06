alter table libavtorname add fulltext index (FirstName, LastName, MiddleName);
alter table libseqname add fulltext index (SeqName);

create or replace function concat_name (FirstName varchar(99), MiddleName varchar(99), LastName varchar(99))
    returns varchar(255) deterministic
    return trim(concat(LastName, ' ', FirstName, ' ', MiddleName))
