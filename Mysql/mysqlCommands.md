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
### Table with NOT Null Values / Default Values / Keys / Extra
- Not Null
CREATE TABLE [TABLENAME]
    (
       [column_name data_type NOT NULL],
       [column_name data_type NOT NULL]
    );
    
- INSERT INTO cats(name) VALUES('Jetson'); or INSERT INTO cats() VALUES();

- Default Value
CREATE TABLE [TABLENAME]
    (
       [column_name data_type DEFAULT [name]],
       [column_name data_type DEFAULT [name]]
    );

- Primary Keys
Each ID needs to be unique.
CREATE TABLE [TABLENAME]
    (
        [IDNAME] INT NOT NULL PRIMARY KEY,
        name VARCHAR(100),
        age INT
    );

Primary keys cannot be null.
CREATE TABLE [TABLENAME]
    (
        [IDNAME] INT AUTO_INCREMENT PRIMARY KEY,        //or use PRIMARY KEY ([IDNAME])
        name VARCHAR(100),
        age INT
    ); 


- SHOW TABLES;
- SHOW COLUMNS FROM [TABLENAME]; Or DESC [TABLENAME];
- ORDER BY [NAME] DESC;
- DROP TABLE [TABLENAME];

## Data Type 
- int
- varchar([size])
 
## Comment
use -- 

## Data 
- INSERT INTO cats(name,age) VALUES ('Jetson', 7), ('sadie',3);


## Loading sql file
- source file_name.sql          //Under base command use "mysql -u root -p" to enter mysql


## "Select"
- SELECT * FROM [TABLENAME];
- SELECT [NAME],[NAME] FROM [TABLENAME]
- SELECT * FROM [TABLENAME],[TABLENAME] WHERE age = 4
- SELECT [NAME] AS [NEWNAME] FROM [TABLENAME]            //This is temporary for changing column name

## "Update" 
- UPDATE [TABLENAME] SET [COLUMNNAME] = 'NEW NAME' WHERE [COLUMNNAME] = 'OLD NAME'

## "Delete"
- DELETE FROM [TABLENAME] WHERE [COLUMNNAME] = '[ITEMNAME]'
- DELETE FROM [TABLENAME]           //This will delete whole table


## "Concat"
- CONCAT (X,Y,Z)
- SELECT CONCAT ([COLUMNNAME], '!') FROM [TABLENAME]