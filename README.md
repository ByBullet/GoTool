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