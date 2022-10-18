# DCaF
 DeCaffinate and Filter data files

## Build
Run 
```
go build .
```
from the dcaf folder

## Usage

### Data format filtration
Using the -c flag followed by a string will exlude all lines that does not conform to the filter string.   
```
... -c 'D;T;A;'
```
Will only accept lines that starts with a Date(D) followed by a ';' delimiter, followed by a Time(T) ending with a ';' delimiter, 
followed by any(A) type of data ending with a line ending or a ';' delimiter. Note that if there is data proceeding the last delimiter it will not be included in the output, but the valid data will be(assuming all data-types gives a valid match).      
```
1997-03-02;19:31:01;This will be included in the output; This will not be included in the output
```
processed by the 'D;T;A;' filter would result in:
```
1997-03-02,19:31:01,This will be included in the output
```
In this example the ';' delimiters were replaced by ',', this is the default behavior, if you want a different delimitor use the -j flag followed by the delimitor of your choosing.   
   
The datatypes currently implemented are:
| Char | Datatype | Example |
| --- | --- | --- |
| D | Date | 1997-03-02 |
| T | Time | 19:30:25 |
| N | Number | -3.64 |
| I | Integer | 84 |
| A | Any | A3 _ test |

You can use any non-alphanumeric symbol as a delimiter.