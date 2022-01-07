package tools

import "fmt"

/**
Base64 编解码器
*/

type Base64 struct {
	enCode string //编码结果
	deCode []byte //解码结果
}

/* Base64 编码表 */
/* 获取 Base64 字符 */
func getBase64Char(code byte) byte {
	if code <= 25 {
		code += 65
	} else if code > 25 && code <= 51 {
		code += 71
	} else if code > 51 && code <= 61 {
		code -= 4
	} else if code == 62 {
		code = '+'
	} else if code == 63 {
		code = '/'
	}
	return code
}

/* 获取 Base64 码值 */
func getBase64Code(char byte) (byte, error) {

	if char >= 48 && char <= 57 { /* 数字 1 - 9 */
		char += 4
	} else if char >= 65 && char <= 90 { /* 大写 A - Z */
		char -= 65
	} else if char >= 97 && char <= 122 { /* 小写 a - z */
		char -= 71
	} else if char == 43 { /* + 号 */
		char = 62
	} else if char == 47 { /* / 号 */
		char = 63
	} else if char == '=' {
		char = 64
	} else {
		return char, fmt.Errorf("invalid base64 text format")
	}

	return char, nil
}

// Encode 编码
func (base *Base64) Encode(src []byte) string {
	if len(src) <= 0 {
		return ""
	}

	/* 数据长度不足 3 的倍数补 0 */
	var zeroCnt int
	for len(src)%3 != 0 {
		src = append(src, 0)
		zeroCnt++
	}

	/* 按 3*8bit 一组处理 */
	srcLen := len(src)
	for i := 0; i < srcLen; i += 3 {
		var group uint32 //按 3*8bit 一组存入 4 字节的内存中,以便分组操作
		var code byte

		group |= uint32(src[i])
		group <<= 8
		group |= uint32(src[i+1])
		group <<= 8
		group |= uint32(src[i+2])

		/* 按 4*6=24bit 一组组成 4 个字符 */
		for k := 3; k >= 0; k-- {
			code = byte((group >> (k * 6)) & 0x3f) // & 0x3f表示取后 6 位
			base.enCode += string(getBase64Char(code))
		}
	}

	/* 将最后 zeroCnt 位置为 '=' 号 */
	if zeroCnt > 0 {
		codeLen := len(base.enCode)
		base.enCode = base.enCode[:codeLen-zeroCnt]
		for i := codeLen - zeroCnt; i < codeLen; i++ {
			base.enCode += "="
		}
	}

	return base.enCode
}

func (base *Base64) EncodeStr(str string) string {
	return base.Encode([]byte(str))
}

// Decode 解码
func (base *Base64) Decode(src []byte) ([]byte, error) {
	srcLen := len(src)

	if srcLen <= 0 {
		return src, fmt.Errorf("base64 code empty")
	}

	if srcLen%4 != 0 {
		return src, fmt.Errorf("base64 code invalid format")
	}

	base.deCode = make([]byte, srcLen*3/4)
	var index int

	/* 4 个字节为一组操作 */
	for i := 0; i < srcLen; {
		var opNum uint32

		for j := 0; j < 4; j++ {
			base64Code, err := getBase64Code(src[i])
			if err != nil {
				return nil, err
			}

			if base64Code != 64 {
				opNum |= uint32(base64Code)
			}

			if j < 3 {
				opNum <<= 8
			}
			i++
		}

		num := byte((opNum >> 22) | (opNum & 0x300000 >> 20))
		if num != 0 {
			base.deCode[index] = num
			index++
		}
		num = byte(((opNum >> 12) & 0xf0) | (opNum & 0x3c00 >> 10))
		if num != 0 {
			base.deCode[index] = num
			index++
		}
		num = byte(opNum | (opNum & 0x300 >> 2))
		if num != 0 {
			base.deCode[index] = num
			index++
		}

	}

	return base.deCode, nil
}

func (base *Base64) DecodeStr(str string) (string, error) {
	result, err := base.Decode([]byte(str))
	if err != nil {
		return "", err
	}

	return string(result), nil
}
