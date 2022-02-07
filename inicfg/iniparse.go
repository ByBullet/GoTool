package inicfg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//解析状态
const (
	PNothing         = iota //未解析
	PSection                //正在解析 PSection
	PSectionError           //PSection 异常
	PParameters             //正在解析 PParameters
	PParametersError        //PParameters 异常
	PComment                //注释
	PSuccess                //解析成功
)

type Parameters = map[string]Value

type IniReader struct {
	reader   *bufio.Reader
	Sections map[string]Parameters
}

func NewIniReader(reader io.Reader) *IniReader {
	r := new(IniReader)
	r.Sections = make(map[string]Parameters)
	r.reader = bufio.NewReader(reader)
	return r
}

func (ini *IniReader) Get(key string) map[string]Value {
	return ini.Sections[key]
}

//读取一行数据，兼容'\r\n','\n'和'\r'
func (ini *IniReader) getLine() (string, error) {
	var line strings.Builder
	var err error = nil
	var eof bool

	for !eof {
		c, e := ini.reader.ReadByte()
		if e != nil {
			err = e
			break
		}

		switch c {
		case '\n':
			eof = true
			break
		case '\r':
			eof = true
			cc, er := ini.reader.Peek(1)
			if er != nil {
				err = er
				break
			}

			if cc[0] == '\n' {
				c, _ = ini.reader.ReadByte()
			}
			break
		default:
			line.WriteByte(c)
		}
	}

	return line.String(), err
}

func (ini *IniReader) Parse() error {
	var retErr error = nil

	var lineNum int
	var lastPSection string
	//读取一行字符串
	for {
		line, err := ini.getLine()
		if err != nil {
			if err != io.EOF {
				retErr = err
			}

			if line == "" {
				break
			}
		}

		if line == "" {
			continue
		}

		line = strings.Trim(line, " ") //去除首尾空格
		lineNum++

		var index int
		lineLen := len(line)
		//分析一行数据
		parseStatus := PNothing
		for parseStatus != PSuccess {
			switch parseStatus {
			case PNothing:
				{
					for ; index < lineLen; index++ {
						if line[index] == ';' {
							parseStatus = PComment
							break
						} else if line[index] == '[' {
							parseStatus = PSection
							break
						} else if line[index] == ']' { //']'比'['先出现
							parseStatus = PSectionError
							break
						} else if line[index] != ' ' { //忽略前置空格
							parseStatus = PParameters
							break
						}
					}
					break
				}
			case PSection:
				{
					content := line[index:]
					commentIndex, splitIndex := strings.IndexByte(content, ';'), strings.IndexByte(content, ']')

					//注释';'比']'先出现
					if commentIndex != -1 && commentIndex <= splitIndex {
						parseStatus = PSectionError
						break
					}

					if splitIndex == -1 { //'['没有闭合
						parseStatus = PSectionError
						break
					}

					lastPSection = content[index+1 : splitIndex]

					if ini.Sections[lastPSection] == nil {
						ini.Sections[lastPSection] = make(Parameters)
					}

					parseStatus = PSuccess

					break
				}
			case PSectionError:
				{
					return fmt.Errorf("[PSection] format invalid in the file %d line. ", lineNum)
				}
			case PParameters:
				{
					commentIndex, splitIndex := strings.IndexByte(line, ';'), strings.IndexByte(line, '=')
					//';'出现在'='前面
					if commentIndex != -1 && commentIndex <= splitIndex {
						parseStatus = PParametersError
						break
					}

					//没有'='号或者'='号处于开头或结尾
					if splitIndex == -1 || splitIndex == index || splitIndex == lineLen-1 {
						parseStatus = PParametersError
						break
					}

					key, value := strings.Trim(line[index:splitIndex], " "), strings.Trim(line[splitIndex+1:lineLen], " ")
					commentIndex = strings.IndexByte(value, ';')
					//将注释部分抛弃
					if commentIndex != -1 {
						value = value[0:commentIndex]
					}

					ini.Sections[lastPSection][key] = Value{value}

					parseStatus = PSuccess
					break
				}
			case PParametersError:
				{
					return fmt.Errorf("PParameters format invalid in the file %d line. ", lineNum)
				}
			case PComment:
				{
					parseStatus = PSuccess
					break
				}
			}
		}
	}

	return retErr
}
