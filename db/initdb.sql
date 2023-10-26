-- 创建账号
create role goadmin with login password '1234' superuser createdb;
-- 创建数据库
create database goAdmin;
-- uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
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
alter table if exists users add if not exists superadmin varchar null; -- 是否是超级管理员
alter table if exists users add if not exists roldId int not null; -- 角色id
-- 添加uuid 列
alter table if exists users add if not exists uuid UUID default gen_random_uuid() not null; 

-- 启动加密 pgcrypto模块
create extension pgcrypto
insert into users (name, pwd, createdOn) values (
    'admin',
    crypt('1234', gen_salt('bf')),
    CURRENT_TIMESTAMP
);


-- 权限
-- 用户表 角色表 权限表
-- 角色表
create table if not exists roles (
    id serial primary key,
    name varchar(20) not null,
    description varchar(50) null
)

alter table roles add if not exists policy varchar[] not null;

-- 权限表
create table if not exists permissions (
    id serial primary key,
    name varchar(20) not null,
    description varchar(50) null,
)

-- 菜单表