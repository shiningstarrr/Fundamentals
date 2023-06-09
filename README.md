# Fundamentals

- [Naming Convention](#naming-convention)
    - [Reference: https://medium.com/@func25/naming-convention-in-go-master-the-art-2beb45fba15b](https://medium.com/@func25/naming-convention-in-go-master-the-art-2beb45fba15b)
- [Big O Cheet Sheet](#big-o-cheet-sheet)
    - [Reference: https://www.bigocheatsheet.com/pdf/big-o-cheatsheet.pdf](https://www.bigocheatsheet.com/pdf/big-o-cheatsheet.pdf)


## Naming Convention
### Variables name
- use "_" prefix underscore for variables and constants that are not exported.
- use variables and constants that are not exported with a lowercase first letter.
- start with uppercase letter if exported.
- If variable type is bool, its name should start with Has, Is, Can or Allow, etc.


### Constant name
- constant should use all capital letter and use underscore _ to seperate words.

### Function name
- camelCase: Functions should be named using camelCase (un-exported), where the first letter of the first word is lowercase and the first letter of subsequent words is uppercase. For example, “retrieveUser” or “fetchConfigs”.
- Use verb in the function name: The function name should begin with a verb indicating the action it performs.

### Package name
- lowercase: it’s best to keep it simple and use a single word or multiple words all smushed together without any separators like dots or underscores. 

### File name
- snake_lower_case: the convention for file names is to stick with lowercase letters and use snake_case or lowerCamelCase. 
- suffix “_test”: When naming files that contain test functions, it is convention to append the suffix “_test” to the file name.

### Error 
- Error variables should have a name starting with “Err” and should be capitalized, for example ErrInvalidInput, ErrNotFound.

## Big O Cheet Sheet
- Link: https://www.bigocheatsheet.com