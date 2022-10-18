# DCaF
 DeCaffinate and Filter data files

## Build
Run 
```
go build .
```
from the dcaf folder

## Usage
dcaf requires an input file(`-i`) and an output file(`-o`):
```
dcaf -i test.ssv -o out.csv
``` 
This will go line by line through the `test.ssv` file and write each line to the `out.csv`(output)file. To copy a file can be fun, but that is not why I made this tool. 

### Data Format Filtration
Using the -c flag followed by a string will exlude all lines that does not conform to the filter string.   
```
... -c 'D;T;A;'
```
Will only accept lines that starts with a Date(D) followed by a `;` delimiter, followed by a Time(T) ending with a `;` delimiter, 
followed by any(A) type of data ending with a line ending or a `;` delimiter. Note that if there is data proceeding the last delimiter it will not be included in the output, but the valid data will be(assuming all data-types gives a valid match).      
```
1997-03-02;19:31:01;This will be included in the output; This will not be included in the output
```
processed by the 'D;T;A;' filter would result in:
```
1997-03-02,19:31:01,This will be included in the output
```
In this example the `;` delimiters were replaced by `,`, this is the default behavior, if you want a different delimitor use the `-j` flag followed by the delimitor of your choosing. The last delimitor is also removed completely since it may or may not exist in the original, if you prefer different behavior check out the Line Reconstruction section.    
   
The datatypes currently implemented are:
| Char | Datatype | Example |
| --- | --- | --- |
| D | Date | 1997-03-02 |
| T | Time | 19:30:25 |
| N | Number | -3.64 |
| I | Integer | 84 |
| A | Any | A3 _ test |

You can use any non-alphanumeric symbol as a delimiters, including whitespace. If a datatype was to be provided after another datatype without a dilimiter symbol, only the later is tested. `'TD;'` will only produce a valid line if the field contains a Date, and only a Date. Having consecutive delimiters will result in the data being match with the Any(A) Datatype, `';A;'` is hence equivalent to `';;'`.

### Line Reconstruction
THIS FUNCTION IS YET TO BE ADDED
```
... -r '1,0 2,'
```