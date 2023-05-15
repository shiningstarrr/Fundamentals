# Mysql Basic Commands
https://dev.mysql.com/doc/refman/8.0/en/
## Databases
- SHOW DATABASES;
- CREATE DATABASE [NAME];
- DROP DATABASE [NAME];
- USE [DATABSE NAME]
- SELECT DATABASE ();


## Table
- CREATE TABLE [TABLENAME]([column_name data_type],[column_name data_type]);
### Table with NOT Null Values / Default Values / Keys / Extra
- Not Null
CREATE TABLE [TABLENAME]([column_name data_type NOT NULL], [column_name data_type NOT NULL]);
- IFNULL(A,B)
- INSERT INTO cats(name) VALUES('Jetson'); or INSERT INTO cats() VALUES();
- Default Value
CREATE TABLE [TABLENAME]([column_name data_type DEFAULT [name]],[column_name data_type DEFAULT [name]]);
- Primary Keys
Each ID needs to be unique.
CREATE TABLE [TABLENAME]([IDNAME] INT NOT NULL PRIMARY KEY,name VARCHAR(100),age INT);
- Primary keys cannot be null.
CREATE TABLE [TABLENAME]([IDNAME] INT AUTO_INCREMENT PRIMARY KEY,        //or use PRIMARY KEY ([IDNAME])name VARCHAR(100),age INT); 
- SHOW TABLES;
- SHOW COLUMNS FROM [TABLENAME]; Or DESC [TABLENAME];
- ORDER BY [NAME] DESC;
- DROP TABLE [TABLENAME];


## Data Type 
- varchar([size])
- char              //has a fixed length
- tinyint(127), smallint(32767), mediumint(8388607), int(2147483647), bigint(2^63-1)
- decimal([totalNumberOfDigits], [DigitsAfterDecimal])
- float (memory needed: 4 Bytes, Precision issues: ~7 digits)
- double (memory needed: 8 bytes, Precision issues: ~15 digits)
### Type of time and date 
https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html
- date([yyyy-mm-dd]), time([hh:mm:ss]), datetime([yyyy-mm-dd] [hh:mm:ss])
- SELECT CURDATE();   - SELECT CURTIME();   - SELECT CURRENT_TIMESTAMP() or use NOW(); 
- SELECT DAYOFMONTH('');  - SELECT DAYOFWEEK('');  - SELECT DAYOFYEAR('');
- SELECT MONTHNAME('');
- SELECT DATE_FORMAT([COLUMNNAME], '%a, %b, %x') FROM [TABLENAME];
- CREATE TABLE [TABLENAME] ([COLUMNNAME] TIMESTAMP DEFAULT CURRENT_TIMESTAMP, [COLUMNAME] TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);      //use update to set up the update timestamp.


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
- SELECT [COLUMNNAME], MIN([COLUMNNAME]), COUNT(*) FROM [TABLENAME] GROUP BY [COLUMNNAME];
### "Min" and "Max"
- MIN([COLUMNNAME]), MAX([COLUMNNAME])
- SELECT * FROM [TABLENAME] WHERE [COLUMNNAME] = (SELECT MIN([COLUMNNAME]) FROM BOOKS);
### Subqueries
- SELECT * FROM [TABLENAME] WHERE [COLUMNNAME] = (SELECT MIN([COLUMNNAME]) FROM [TABLENAME])
#### "Sum"
- SELECT SUM([COLUMNNAME]);
- SUM + GROUPBY
### "Average"
- SELECT AVG([COLUMNNAME]);
- AVG + GROUPBY



## Logical Operation
-  !=
-  WHERE [COLUMNNAME] LIKE/NOLIKE '% %'
-  > / >=
-  < / <=
-  AND (&&)
-  OR  (||)
-  BETWEEN [A] AND [B]
-  CAST([TYPE] AS [ANOTHERTYPE])
-  WHERE [COLUMNNAME] IN (' ', ' ', ' ');
-  SELECT [COLUMNNAME],[COLUMNNAME], 
     CASE WHEN [CONDITION] THEN '[ANYNAME]'
     ELSE '[ANOTHERNAME]'
     END AS [NAME]
   FROM [TABLENAME];
-  IS NULL

Exercise Example: 
mysql> select author_fname,author_lname,
    -> case when count(*) = 1 then '1 book'
    -> else concat(count(*), 'books')
    -> end as count from books
    -> where author_lname is not null
    -> group by author_fname, author_lname;
 

## Constraints
- UNIQUE   //Example: phoneNumber varchar(15) NOT NULL UNIQUE
- CHECK    //Example: age int CHECK (age > 18)
- CONSTRAINT + CHECK     //CONSTRAINT age_not_negative CHECK (age>18)
- CONSTRAINT [CONSTRAINTNAME] UNIQUE ([COLUMNNAME],[COLUMNNAME])
- ALTER TABLE [TABLENAME]
- ALTER TABLE [TABLENAME] ADD COLUMN [COLUMNNAME] [TYPE];
- ALTER TABLE [TABLENAME] DROP COLUMN [COLUMNNAME] [TYPE];
- RENAME TABLE [TABLENAME] TO [NEWNAME];
- ALTER TABLE [TABLENAME] MODIFY [COLUMNNAME] [NEWTYPE];    //change an existing column's type
- ALTER TABLE [TABLENAME] CHANGE [OLUMNNAME] [NEWTYPE];   //use change to rename and change its data type
- ALTER TABLE [TABLENAME] ADD CONSTRAINT [NAME] CHECK [CONDITION];

 

## Relationships and Joins
### Cross Join
- FOREIGN KEY ([columnname]) REFERENCES [tablename][(columnname)];
- FOREIGN KEY ([columnname]) REFERENCES [tablename][(columnname)] ON DELETE CASCADE;       //Synchronize delete
- EXAMPLE: select * from orders where customer_id = (select id from customers where last_name = 'George');
### Inner Join
- SELECT * FROM [TABLENAME] INNER JOIN [ANOTHER TABLE NAME] ON [TABLENAME].[COLUMNNAME] = [ANOTHER TABLE NAME].[COLUMNNAME] INNER JOIN...;
### Left Join / Right Join
- SELECT [COLUMNNAME...] FROM [TABLENAME] LEFT/RIGHT JOIN [ANOTHER TABLE NAME] ON [TABLENAME].[COLUMNNAME] = [ANOTHER TABLE NAME].[COLUMNNAME] ORDER BY [COLUMNNAME];



## Views - save as virtual table
- CREATE VIEW [VIEWNAME] AS SELECT...
- CREATE OR REPLACE VIEW [VIEWNAME] AS SELECT...
- ALTER VIEW [VIEWNAME] AS SELECT...
- DROP VIEW [VIEWNAME]
- HAVING [CONDITION] GROUP BY [COLUMNNAME]  // Specifies conditions on groups.  
  WHERE   // Where clause specifies selection conditions.
- GROUP BY [COLUMNNAME] WITH ROLLUP    //Display the total value of each group. 
### Modes
- SELECT @@GLOBAL.sql_mode;
- SELECT @@SESSION.sql_mode;
- SET GLOBAL sql_mode = 'modes';
- SET SESSION sql_mode = 'modes';