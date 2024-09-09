# QuizGame
## Simple project to practise golang. Use next modules:
```golang
import (
  3         "bufio"
  4         "encoding/csv"
  5         "errors"
  6         "flag"
  7         "fmt"
  8         "os"
  9         "path"
 10         "strings"
 11         "time"
 12 )
```
### Example of csv file:
```cmd
$ cat problems.csv 
5+5,10
1+1,2
```
### Build:
```cmd
$ make -B
go build -ldflags "-w -s"
```
### Help:
```cmd
 ./QuizGame -h                        
Usage of ./QuizGame:
  -csv string
        a csv file which contains the task in the format of 'question,answer' (default "problems.csv")
  -limit uint
        the time limit for the quiz in seconds (default 30)
```
