create table if not exists todos (
    id integer auto_increment primary key,
    title varchar(40),
    description varchar(40),
    status boolean
)
