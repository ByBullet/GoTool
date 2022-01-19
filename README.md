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
|file output|Yes     |

### usage
Four levels int thid module; <br>
LOGGER_ERROR LOGGER_WARNING LOGGER_INFO LOGGER_DEBUG;<br>
Each level of logging is output to a different log file;<br>

```
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

