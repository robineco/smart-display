# smart-display-backend

## Project Setup
You need a running mariadb with a table named "temperature"

table-setup.sql
```
create table temperature.temperature
(
    ID       int auto_increment primary key,
    Temp     float     null,
    Time     timestamp null,
    Location char(255) null
);
```
