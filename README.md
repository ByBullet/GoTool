# GoTool
some tools of golang

## Base64
### usage
```azure
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
|Asynchronous log|No |
|file output|Yes     |

### usage
Four levels int thid module; <br>
LOGGER_ERROR LOGGER_WARNING LOGGER_INFO LOGGER_DEBUG;<br>
Each level of logging is output to a different log file;<br>

```
/* log initialization settings; LoggerLevle */
/* Log output to log folder in the root directory */
//logger.SetConfig(logger.Config{LoggerLevle: logger.LOGGER_INFO, OutType: logger.LOGGER_FILE, OutDir: "log"})  
/* Log output to log console */ 
logger.SetConfig(logger.Config{LoggerLevle: logger.LOGGER_DEBUG, OutType: logger.LOGGER_CONSOLE}) 

logger.ErrorFormat("error msg: %s\n", err)
logger.Error("error msg")
```

