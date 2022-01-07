# GoTool
some tools of golang

## Base64
### usage
```azure
import (
    "fmt"
    tools "github.com/ByBullet/GoTool/codec"
)
    
func main() {
    base64 := tools.Base64{}
    str := "abc"
    // encode
    ret1 := base64.EncodeStr(str)
    fmt.Println("encode:", ret1)
    // decode
    ret2, err := base64.DecodeStr(ret1)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("decode:", ret2)
}
```