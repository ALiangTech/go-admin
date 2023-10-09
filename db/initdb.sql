-- 创建账号
create role goadmin with login password '1234' superuser createdb;
-- 创建数据库
create database goAdmin;
-- 创建用户表
create table if not exists users (
    id serial primary key,
    name char(16) not null,
    pwd  varchar not null,
    createdOn timestamptz not null,
    updatedOn timestamptz null
 );
-- 其实updatedOn 是可以为null

alter table if exists users alter updatedOn drop not null;

-- 启动加密 pgcrypto模块
create extension pgcrypto
insert into users (name, pwd, createdOn) values (
    "admin",
    crypt("1234", gen_salt("bf")),
    CURRENT_TIMESTAMP
);