create table files
(
    filename   text not null primary key,
    created_at timestamptz default current_timestamp
)
