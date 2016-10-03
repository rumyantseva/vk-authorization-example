# VK authorization example

## Preparation

Use govendor for vendoring:

    go get -u github.com/kardianos/govendor
    govendor sync

Use PostgreSQL as database:

    create database auth encoding='UTF8' lc_collate='C' lc_ctype='UTF-8' template='template0';

Set up database:

    go get github.com/mattes/migrate
    migrate -url "postgres://postgres:postgres@127.0.0.1:5432/auth?sslmode=disable" -path ./migrations up

## Configuration


