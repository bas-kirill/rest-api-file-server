create table files
(
    file_id    integer generated always as identity primary key,
    filename   text not null,
    created_at timestamptz default current_timestamp
)