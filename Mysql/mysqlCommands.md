# Mysql Basic Commands

## Databases
- SHOW DATABASES;
- CREATE DATABASE [NAME];
- DROP DATABASE [NAME];
- USE [DATABSE NAME]
- SELECT DATABASE ();

## Table
- CREATE TABLE [TABLENAME]
    (
        [column_name data_type],
        [column_name data_type]
    );

- SHOW TABLES;
- SHOW COLUMNS FROM [TABLENAME]; Or DESC [TABLENAME];
- SELECT * FROM [TABLENAME];
- ORDER BY [NAME] DESC;
- DROP TABLE [TABLENAME];

## Data Type 
- int
- varchar([size])

## Comment
use -- 

## Data 
- INSERT INTO cats(name,age) VALUES ('Jetson', 7), ('sadie',3);

## NOT Null values
- CREATE TABLE [TABLENAME]
    (
       [column_name data_type NOT NULL],
       [column_name data_type NOT NULL]
    );
    
- INSERT INTO cats(name) VALUES('Jetson'); or INSERT INTO cats() VALUES();
