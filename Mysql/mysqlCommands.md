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
- SELECT [NAME],[NAME] FROM [TABLENAME];
- SELECT * FROM [TABLENAME],[TABLENAME] WHERE age = 4;
- SELECT [NAME] AS [NEWNAME] FROM [TABLENAME];            //This is temporary for changing column name

### "Select distinct"
- SELECT DISTINCT [COLUMNNAME] FROM [TABLENAME];          //Avoid duplicate
### "Order by"
- SELECT [COLUMNNAME] FROM [TABLENAME] ORDER BY [COLUMNNAME];
- SELECT [COLUMNNAME] FROM [TABLENAME] ORDER BY [COLUMNNAME] DESC;
- SELECT [COLUMNAME_1],[COLUMNNAME_2],[COLUMNNAME_3] FROM [TABLENAME] ORDER BY 2,3,4;          //Order by the second column, and then the third, then the fourth one
### "Limit"
- SELECT [COLUMNNAME] FROM [TABLENAME] ORDER BY [COLUMNNAME] LIMIT [COUNT];
- SELECT [COLUMNNAME] FROM [TABLENAME] ORDER BY [COLUMNNAME] LIMIT [STARTPOSITION], [COUNT];   
## "Update" 
- UPDATE [TABLENAME] SET [COLUMNNAME] = 'NEW NAME' WHERE [COLUMNNAME] = 'OLD NAME';
## "Delete"
- DELETE FROM [TABLENAME] WHERE [COLUMNNAME] = '[ITEMNAME]';
- DELETE FROM [TABLENAME];           //This will delete whole table


## String functions and operators
- [String functions ref link: ](https://dev.mysql.com/doc/refman/8.0/en/string-functions.html)

### "Concat"
- CONCAT (X,Y,Z);
- SELECT CONCAT ([COLUMNNAME], '!') FROM [TABLENAME];
- SELECT CONCAT_WS ('-', [COLUMNNAME],[COLUMNNAME]);            //Adds seperator between each element
- SELECT CONCAT(SUBSTRING(...), SUBSTRING(...)) AS [NAME] FROM [TABLENAME];
### "Substring"
- SELECT SUBSTRING ('HELLO WORLD', [START POSITION(int)], [NUMBER OF CHARACTERS (int)]) FROM [TABLENAME];         
- SELECT SUBSTRING ([STRING], [START POSITION]) FROM [TABLENAME]               //This will go from start position to the end of the string
- SELECT SUBSTRING ([STRING], -3, 1) FROM [TABLENAME] ;        //Count backwards
### "Replace"
- SELECT REPLACE([STR],[FROMSTR], [TOSTR])
### "Reverse"
- SELECT REVERSE ([STR]);
### "Char_length"
- SELECT CHAR_LENGTH([STR]);
### "Upper and Lower"
- SELECT UPPER([STR]);
- SELECT LOWER([STR]);
### "Insert"
- SELECT INSERT([STR],[POS],[LEN],[NEWSTR])        //Len is the number of character to replace
### "Left, Right, Repeat"
- SELECT LEFT([STR],[LEN])          //Return the leftmost len characters from str
- SELECT RIGHT([STR],[LEN])
- SELECT REPEAT([STR],[COUNT])
### "Trim"
- SELECT TRIM('   BAR   ');        //Remove space
- SELECT TRIM(lEADING 'X' FROM 'XXXBARXXX');
- SELECT TRIM(BOTH 'X' FROM 'XXXBARXXX');
- SEELCT TRIM(TRAILING 'XYZ' FROM 'BARXXYZ');
### "Like"
- SELECT [COLUMNNAME] FROM [TABLENAME] WHERE [COLUMNNAME] LIKE '%[ANYCHARACTER]%'           //LIKE '_ _ _ _' means four characters
- WHERE TITLE LIKE '%\%%'          //for searching title with % sign.




## Aggregate functions
### "Count"
- SELECT COUNT(*) FROM [TABLENAME];
- SELECT COUNT(DISTINCT [COLUMNNAME]) FROM [TABLENAME];
### "Group by"
- SELECT [COLUMNNAME], COUNT(*) FROM [TABLENAME] GROUP BY [COLUMNNAME];
### "Min" and "Max"
- MIN([COLUMNNAME]), MAX([COLUMNNAME])
- SELECT * FROM [TABLENAME] WHERE [COLUMNNAME] = (SELECT MIN([COLUMNNAME]) FROM BOOKS);
