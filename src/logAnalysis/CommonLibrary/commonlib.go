package CommonLibrary

import (
	"fmt"
	"log"
	//"time"

	"gopkg.in/kataras/iris.v6"
	"strings"

	"logAnalysis/runglobal"
	"regexp"
	"strconv"
	"time"
)

//import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalln("error: %s", err.Error())
	}
}

func CheckPrintlnError(err error) {
	if err != nil {
		log.Println("error: ", err.Error())
	}
}

type diverror struct {
	s    string
	args []interface{}
}

func (e *diverror) Error() string {
	return fmt.Sprint(e.s, e.args)
}

func CatchExecption(args ...interface{}) { //必须要先声明defer，否则不能捕获到panic异常
	if err := recover(); err != nil {
		CheckPanicError(err.(error), args...)
	}

}

func CheckPanicError(err error, args ...interface{}) {
	if err != nil {
		err := &diverror{err.Error(), args}
		panic(err)
	}
}
func CheckHtmlError(err error, ctx *iris.Context) {
	if err != nil {
		ctx.Log(iris.DevMode, err.Error())
	}

}

func CheckErrorPrintln(filename string, method string, err error) bool {
	if err != nil {
		return true
		fmt.Println(" 文件名：", filename, " 方法：", method, " 异常：", err)
	}
	return false
}

func IsNumber(s string) bool {
	r := regexp.MustCompile("^[0-9]+([.]{1}[0-9]+){0,1}$")
	return r.MatchString(s)
}

//根据传入的标题来识别需要转换为什么样的类型，如果没有传入，默认不转换，直接存为字符串
func IsDataType(title string, da interface{}) interface{} {
	//log.Println("传入的类型标题-------------------", title)
	if da.(string) == "-" {
		return da
	}
	if strings.Contains(title, ":") {
		dtype := strings.Split(title, ":")[1]
		if dtype == "" {
			return da
		} else {
			switch dtype {
			case "int":
				if !IsNumber(da.(string)) {
					return da
				}
				res, err := strconv.Atoi(da.(string))
				CheckPanicError(err, da)
				return res
			case "date":
				if len(strings.Split(title, ":")) != 3 {
					return da
				}
				regname := strings.Split(title, ":")[2]
				if val, ok := runglobal.PulicRegex[regname]; ok {
					r := regexp.MustCompile(val)
					res := r.FindStringSubmatch(da.(string))
					//log.Println("正则表达式是", val)
					//log.Println("匹配的结果是", res)
					t := time.Date(S2I(res[3]), IsMonth(res[2]), S2I(res[1]), S2I(res[4]), S2I(res[5]), S2I(res[6]), 0, time.Local)
					//log.Println("生成的时间是", t.String())
					return t
				} else {
					return da
				}
			case "float":
				if !IsNumber(da.(string)) {
					return da
				}
				value, err := strconv.ParseFloat(da.(string), 64)
				CheckPanicError(err, da)
				return value
			case "string":
				return da.(string)
			default:
				return da.(string)

			}
		}
	}
	return da
}

func IsMonth(mon string) time.Month {
	switch mon {
	case "January":
		return time.January
	case "February":
		return time.February
	case "March":
		return time.March
	case "April":
		return time.April
	case "May":
		return time.May
	case "June":
		return time.June
	case "July":
		return time.July
	case "October":
		return time.October
	case "November":
		return time.November
	case "December":
		return time.December
	case "September":
		return time.September

	}
	return time.Month(0)
}

func S2I(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}
