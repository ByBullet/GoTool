# GoTool
some tools of golang

## Base64
### usage
```go
import (
    "fmt"
    "github.com/ByBullet/GoTool/codec"
)

func main() {
    base64 := codec.Base64{}
    str := "abc"
    
    ret1 := base64.EncodeStr(str)
    fmt.Println("encode:", ret1)
    ret2, err := base64.DecodeStr(ret1)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("decode:", ret2)
}

```

## logger
### feature
| feature |  support |   
|  ----   |  ----    | 
| log leve|  Yes     | 
|file output|Yes     |

### usage
Four levels int thid module; <br>
LOGGER_ERROR LOGGER_WARNING LOGGER_INFO LOGGER_DEBUG;<br>
Each level of logging is output to a different log file;<br>

```go
import (
	"github.com/ByBullet/GoTool/logger"
)

/* log initialization settings*/
func init() {
	/* Log output to log folder in the root directory;This module must be initialized before other modules; */
	logger.SetConfig(logger.Config{LoggerLevel: logger.LOGGER_DEBUG, OutType: logger.LOGGER_FILE, OutDir: "log"})
}

func main() {
	logger.Debug("debug")
	logger.InfoFormat("%s\n", "info")
}
```

## inicfg
Ini configuration file parsing tool
### feature
| feature             |  support |   
|        ----         |   ----   | 
| file stream input   |   Yes    | 
| string stream input |   Yes    |
### usage
An ini configuration file is required as follows:
```ini
; MySQL config
[MySQL]
host = 127.0.0.1
port = 3306
username = root
password = 123456
```
```go
import (
    "fmt"
    "github.com/ByBullet/GoTool/inicfg"
    "os"
)

func main() {
    file, err := os.Open("filepath")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    defer file.Close()
    
    iniReader := inicfg.NewIniReader(file)
    
    //parsing
    err = iniReader.Parse()
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    //get [MySQL] section
    mysql := iniReader.Get("MySQL")
    if mysql == nil {
        fmt.Println("No such section is a [MySQL].")
        return
    }
    
    hostVal := mysql["host"]
    portVal := mysql["port"]
    nameVal := mysql["username"]
    pwdVal  := mysql["password"]
    
    //to string
    host := hostVal.AsString()
    //to int
    port, e := portVal.AsInt()
    if e != nil {
        fmt.Println(e)
        return
    }
    
    name := nameVal.AsString()
    pwd := pwdVal.AsString()
    
    fmt.Println("host: ", host)
    fmt.Println("port: ", port)
    fmt.Println("uname: ", name)
    fmt.Println("pwd: ", pwd)	
}
```