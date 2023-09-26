BEGIN;

set time zone 'UTC-0';

create table files
(
    file_id    integer generated always as identity primary key,
    filename   varchar(256) unique not null,
    created_at timestamptz default current_timestamp
);

END;