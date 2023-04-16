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

## NOT Null Values / Default Values / Keys / Extra
### Not Null
- CREATE TABLE [TABLENAME]
    (
       [column_name data_type NOT NULL],
       [column_name data_type NOT NULL]
    );
    
- INSERT INTO cats(name) VALUES('Jetson'); or INSERT INTO cats() VALUES();

### Default Value
- CREATE TABLE [TABLENAME]
    (
       [column_name data_type DEFAULT [name]],
       [column_name data_type DEFAULT [name]]
    );

### Primary Keys
Each ID needs to be unique.
- CREATE TABLE [TABLENAME]
    (
        [IDNAME] INT NOT NULL PRIMARY KEY,
        name VARCHAR(100),
        age INT
    );

Primary keys cannot be null.
- CREATE TABLE [TABLENAME]
    (
        [IDNAME] INT AUTO_INCREMENT,
        name VARCHAR(100),
        age INT,
        PRIMARY KEY ([IDNAME])
    );
