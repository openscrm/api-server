package util

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"github.com/iancoleman/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
	ecode2 "openscrm/common/ecode"
	log2 "openscrm/common/log"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

func GenerateModelsDocs(models ...interface{}) (docs []byte, err error) {
	for _, model := range models {
		md, err := GenerateModelDocs(model)
		if err != nil {
			err = errors.Wrap(err, "GenerateModelDocs failed")
			return md, err
		}
		docs = append(docs, md...)
	}
	return
}

func GenerateModelDocs(model interface{}) (md []byte, err error) {
	var t reflect.Type
	t = reflect.TypeOf(model)
	md = []byte(`

## ` + t.String() + `
**参数**
| 参数名      | 类型   | 说明                   |
| -------------------- | -------- | ------------------------ |
`)

	err = GetFieldsDocs(model, &md)
	if err != nil {
		err = errors.Wrap(err, "GetFieldsDocs failed")
		return
	}
	md = append(md, "\r\n"...)
	jsonStr, _ := jsoniter.MarshalIndent(model, "", " ")
	md = append(md, "\r\n##示例\r\n```\r\n"+string(jsonStr)+"\r\n```\r\n"...)
	return
}

func GetSubFieldDocs(field interface{}, md *[]byte) (err error) {
	var t reflect.Type
	t = reflect.TypeOf(field)
	if field == nil {
		err = errors.New("field is nil")
		return
	}
	if md == nil {
		err = errors.New("md is nil")
		return
	}
	var v reflect.Value
	v = reflect.ValueOf(field)

	exp, _ := regexp.Compile("comment:'(.+)'")
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") != "-" {
			gormTag := t.Field(i).Tag.Get("gorm")
			fieldName := strcase.ToSnake(t.Field(i).Name)
			fieldType := v.Field(i).Type().String()
			fieldComment := ""
			rs := exp.FindStringSubmatch(gormTag)
			if len(rs) == 2 {
				fieldComment = rs[1]
			}
			content := `| ` + fmt.Sprintf("%-20v", fieldName) + ` | ` +
				fmt.Sprintf("%-8v", fieldType) + ` | ` + fmt.Sprintf("%-24v", fieldComment) + " |\r\n"
			*md = append(*md, content...)
		}
	}
	return
}

func GetFieldsDocs(model interface{}, md *[]byte) (err error) {
	var t reflect.Type
	t = reflect.TypeOf(model)
	if model == nil {
		err = errors.New("model is nil")
		return
	}
	if t.Kind() != reflect.Struct {
		err = errors.New("only struct type")
		return
	}

	if md == nil {
		err = errors.New("md is nil")
		return
	}

	var v reflect.Value
	v = reflect.ValueOf(model)

	exp, _ := regexp.Compile("comment:'(.+)'")

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") != "-" {
			gormTag := t.Field(i).Tag.Get("gorm")
			docsTag := t.Field(i).Tag.Get("docs")
			fieldName := strcase.ToSnake(t.Field(i).Name)
			kind := v.Field(i).Type().Kind()
			if kind == reflect.Struct {
				ignore := false
				if _, ok := v.Field(i).Interface().(time.Time); ok {
					ignore = true
				}
				if strings.Contains(docsTag, "ignore") {
					ignore = true
				}
				if !ignore {
					err = GetSubFieldDocs(v.Field(i).Interface(), md)
					if err != nil {
						err = errors.Wrap(err, "GetFieldsDocs sub struct failed")
						return
					}
					continue
				}
			}
			fieldType := v.Field(i).Type().String()
			fieldComment := ""
			rs := exp.FindStringSubmatch(gormTag)
			if len(rs) == 2 {
				fieldComment = rs[1]
			}
			content := `| ` + fmt.Sprintf("%-20v", fieldName) + ` | ` +
				fmt.Sprintf("%-8v", fieldType) + ` | ` + fmt.Sprintf("%-24v", fieldComment) + " |\r\n"
			*md = append(*md, content...)
		}
	}
	return
}

func GenerateErrorCodeDocs() (md []byte) {
	header := `

## <a id="错误码说明">错误码说明</a>
| 错误码 | 中文说明           | 英文说明 |
| ------ | -------------- | ------- |`
	md = append(md, header...)
	m := ecode2.GetMessages()
	keys := funk.Keys(m).([]int)
	sort.Ints(keys)
	for _, code := range keys {
		msg := m[code]
		content := `
| ` + fmt.Sprintf("%d", code) + ` | ` + msg.Msg + `      | ` + msg.Detail + ` |`
		md = append(md, content...)
	}
	return
}

//JsonEncode 将任意数据JSON序列化
func JsonEncode(data interface{}) string {
	jsonStr, err := jsoniter.MarshalToString(data)
	if err != nil {
		log2.Sugar.Errorw("JsonEncode failed", "err", err, "data", data)
		return ""
	}
	return jsonStr
}

//GetCallerFile 获取调用者源码位置，
func GetCallerFile(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip+2, pc)
	file, line := runtime.FuncForPC(pc[0]).FileLine(pc[0])
	rootPath, _ := os.Getwd()
	if rootPath != "" {
		file = strings.ReplaceAll(file, rootPath, "")
		file = strings.TrimLeft(file, "/")
	}
	return fmt.Sprintf("%s:%d", file, line)
}

//GetCallerName 获取调用者函数名，
func GetCallerName(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip+2, pc)
	f := runtime.FuncForPC(pc[0])
	if gstr.Contains(f.Name(), "/") {
		segs := gstr.Split(f.Name(), "/")
		if len(segs) >= 1 {
			return segs[len(segs)-1]
		}
	}
	return f.Name()
}

//FuncTracer 记录函数的入口和出口，耗时
func FuncTracer(params ...interface{}) func(...interface{}) {
	funName := GetCallerName(1)
	params = append(params, "tracer_caller", GetCallerFile(1))
	log2.Sugar.Infow(funName+" start", params...)
	start := time.Now()
	logLevel := zap.DebugLevel
	return func(results ...interface{}) {
		results = append(results, "tracer_caller", GetCallerFile(1), "cost", time.Since(start).Seconds())
		for i, result := range results {
			if result == nil {
				continue
			}
			//zap日志库不会对空接口自动解指针，这里解了下指针
			v := reflect.ValueOf(result)
			if v.Kind() == reflect.Ptr && v.Elem().IsValid() {
				results[i] = v.Elem().Interface()
			}
			//如果有error类型，则打印error级别
			if e, ok := results[i].(error); ok && e != nil {
				logLevel = zap.ErrorLevel
			}
		}
		if logLevel == zap.ErrorLevel {
			log2.Sugar.Errorw(funName+" done", results...)
		} else {
			log2.Sugar.Infow(funName+" done", results...)
		}
	}
}

func GenBytesOrderByColumn(req interface{}) ([]byte, error) {
	t := reflect.TypeOf(req)

	columns := make([]string, t.NumField())
	for i := range columns {
		columns[i] = t.Field(i).Name
	}
	sort.Strings(columns)

	v := reflect.ValueOf(req)
	buf := &bytes.Buffer{}
	for _, column := range columns {
		val := v.FieldByName(column)
		if !val.CanInterface() {
			continue
		}
		_, err := fmt.Fprintf(buf, "%v=%v&", column, val.Interface())
		if err != nil {
			return nil, err
		}
	}
	s := buf.String()
	if len(columns) > 0 {
		s = strings.TrimRight(s, "&")
	}
	return []byte(s), nil
}
