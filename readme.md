# Golang - gin, gorm, postgres 를 이용한 Restful API Example 

## 0. 사전 준비
- PostgresSQL 설치(docker)
- PostgresSQL 계정생성
- PostgreSQL 

### docker를 이용한 PostgreSQL 실행
```shell
❯ docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=1234 -d postgres
```

#### PostgresSQL 접속
```shell
❯ docker exec -it postgres /bin/bash
root@d51aa050fc5e:/# psql -U postgres
psql (14.1 (Debian 14.1-1.pgdg110+1))
Type "help" for help.

postgres=# 
```

#### PostgresSQL 계정생성 및 db생성
```shell
postgres=# CREATE USER potato PASSWORD '1234' SUPERUSER;
```
#### PostgresSQL db 생성
```shell
postgres=# CREATE DATABASE  test OWNER potato;
```

#### 1. golang 초기화
```shell

```
