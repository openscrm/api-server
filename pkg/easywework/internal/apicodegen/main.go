package main

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Field struct {
	Name       string
	IsRequired bool
	Desc       string
}

type Api struct {
	DocURL     string
	Name       string
	StructName string
	Method     string
	MethodCaml string
	URL        string
	ReqJson    string
	ReqCode    string
	ReqFields  []Field
	RespJson   string
	RespCode   string
	RespFields []Field
}

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
	// unreachable
}

func pickSubMatchString(matches []string, mustTotal int, pickIndex int) string {
	if len(matches) == mustTotal {
		return matches[pickIndex]
	}
	return ""
}

func generateStruct(rawJson string, structName string, subStruct bool, fields []Field) (code string, err error) {
	if rawJson == "" {
		return
	}
	var rawCode []byte
	rawJson = strings.ReplaceAll(rawJson, "\",\n}", "\"\n}")
	//修改腾讯文档的json格式错误
	rawJson = regexp.MustCompile(`([\d"]+)\n([ ]+")`).ReplaceAllString(rawJson, "$1,\n$2")
	rawCode, err = gojson.Generate(strings.NewReader(rawJson),
		gojson.ParseJson,
		structName,
		"workwx",
		[]string{"json"}, subStruct, true,
	)
	if err != nil {
		err = errors.Wrap(err, "gojson.Generate failed")
		return
	}

	code = string(rawCode)
	structFieldRegexp := regexp.MustCompile("(\\w+?)([ ]+?)([\\w\\[\\]]+?)([ ]+?)`json:\"(.+?)\"`")
	allMatches := structFieldRegexp.FindAllStringSubmatch(code, -1)
	code = strings.Replace(code, "package workwx\n", "", 1)
	replacedFields := make(map[string]Field)
	for _, matches := range allMatches {
		for _, field := range fields {
			if field.Name == "errcode" || field.Name == "errmsg" {
				continue
			}
			requiredTips := ""
			jsonFlag := ""
			if strings.HasPrefix(structName, "req") {
				jsonFlag = ",omitempty"
			}
			if field.IsRequired {
				requiredTips = "，必填"
				jsonFlag = ""
			}
			// 支持文档里面字段名的点号语法
			if strings.Contains(field.Name, ".") {
				tmps := strings.Split(field.Name, ".")
				field.Name = tmps[len(tmps)-1]
			}
			if len(matches) == 6 && matches[5] == field.Name {
				code = strings.Replace(code,
					matches[0],
					fmt.Sprintf("// %s %s%s\n\t%s", matches[1], field.Desc, requiredTips, matches[0]),
					1,
				)
				code = strings.Replace(code,
					fmt.Sprintf("`json:\"%s\"`", field.Name),
					fmt.Sprintf("`json:\"%s%s\"`", field.Name, jsonFlag),
					1,
				)

				replacedFields[field.Name] = field
			} else {
				if _, ok := replacedFields[field.Name]; ok {
					continue
				}
				code = strings.Replace(code,
					fmt.Sprintf("} `json:\"%s\"`", field.Name),
					fmt.Sprintf("} `json:\"%s%s\"` //%s%s", field.Name, jsonFlag, requiredTips, field.Desc),
					1,
				)
				replacedFields[field.Name] = field
			}
		}
	}

	code = strings.ReplaceAll(code, "int64", "int")
	commonRespRegexp := regexp.MustCompile("Errcode.+\\n.+?Errmsg.+")
	code = commonRespRegexp.ReplaceAllString(code, "respCommon")
	return
}

func generateCode(apis []*Api) (result []byte, err error) {
	tpl, err := template.ParseFiles("./api_code.tmpl")
	if err != nil {
		err = errors.Wrap(err, "template.ParseFiles failed")
		return
	}
	for _, api := range apis {
		api.URL = strings.ReplaceAll(api.URL, "&amp;", "?")
		api.URL = strings.Split(api.URL, "?")[0]
		segs := strings.Split(api.URL, "/")
		api.StructName = strcase.ToCamel(segs[len(segs)-1]) + strcase.ToCamel(segs[len(segs)-2])
		api.ReqCode, err = generateStruct(api.ReqJson, "req"+api.StructName, false, api.ReqFields)
		if err != nil {
			fmt.Printf("generate reqStruct failed: %+v\n", err)
			continue
		}

		api.RespCode, err = generateStruct(api.RespJson, "resp"+api.StructName, false, api.RespFields)
		if err != nil {
			fmt.Printf("generate reqStruct failed: %+v\n", err)
			continue
		}
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, apis)
	if err != nil {
		err = errors.Wrap(err, "tpl.Execute failed")
		return
	}

	result = buf.Bytes()

	return
}

//注意设置运行的主目录
//apicodegen https://work.weixin.qq.com/api/doc/90000/90135/92572 output/contact_way_api.go
func main() {
	//检查启动参数
	if len(os.Args) <= 1 {
		die("invalid param, example: apicodegen URL [SAVEPATH]")
		return
	}
	docURL := os.Args[1]
	savePath := os.Args[2]
	fmt.Printf("开始抓取和生成API代码，文档地址:%s，代码保存路径:%s\n", docURL, savePath)
	// get the fresh documentation!
	var doc *goquery.Document
	{
		resp, err := http.Get(os.Args[1])
		if err != nil {
			die("http get of errcode documentation failed: %+v\n", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			die("non-200 app: %s\n", resp)
		}

		tmp, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			die("parse document failed: %+v\n", err)
		}

		doc = tmp
	}

	rawHtml, err := doc.Find("#js_doc_apiShow_cnt").Html()
	if err != nil {
		die("failed to get html: %+v\n", err)
	}

	apis := make([]*Api, 0)
	rawHtml = regexp.MustCompile(`<h2 id="([0-9a-zA-Z\-_]+)">`).ReplaceAllString(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtml = regexp.MustCompile(`<h1 id="\w+" class="\w+">`).ReplaceAllString(rawHtml, `<h1 class="rawHtmlSection">`)
	rawHtmlSections := strings.Split(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtmlType := "h2" // 一个页面多个接口
	if len(rawHtmlSections) == 1 {
		rawHtmlType = "h1" // 一个页面一个接口
		rawHtmlSections = strings.Split(rawHtml, `<h1 class="rawHtmlSection">`)
	}
	for index, rawHtmlSection := range rawHtmlSections {
		fmt.Printf("\n\n开始处理第%d个接口\n", index)
		api := &Api{
			DocURL: docURL,
		}

		if rawHtmlType == "h1" {
			rawHtmlSection = `<h1 class="rawHtmlSection">` + rawHtmlSection
			apiNameRegexp := regexp.MustCompile(`<h1 class="rawHtmlSection">(.+?)</h1>`)
			api.Name = pickSubMatchString(apiNameRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)
		} else {
			rawHtmlSection = `<h2 class="rawHtmlSection">` + rawHtmlSection
			apiNameRegexp := regexp.MustCompile(`<a name="(.+?)" class="reference-link"></a>`)
			api.Name = pickSubMatchString(apiNameRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)
		}

		fmt.Printf("%s\n", api.Name)

		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<td style="text-align:left">`, `<td>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式</strong>:`, `<strong>请求方式:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求包体：</strong>`, `<strong>请求示例：</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>返回结果：</strong>`, `<strong>返回结果:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>POST(HTTP)</strong>`, `POST（<strong>HTTPS</strong>）`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>GET(HTTP)</strong>`, `GET（<strong>HTTPS</strong>）`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `：`, `:`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `: </strong>`, `:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `:</strong>`, `</strong>`)

		//过滤掉不是接口的节点
		if !strings.Contains(rawHtmlSection, `<strong>请求方式</strong>`) {
			fmt.Println("没有请求方式，跳过处理")
			continue
		}

		buf := bytes.NewBufferString(rawHtmlSection)
		subDoc, err := goquery.NewDocumentFromReader(buf)
		if err != nil {
			fmt.Printf("goquery.NewDocumentFromReader failed: %+v\n", err)
			continue
		}

		apiMethodRegexp := regexp.MustCompile(`<strong>请求方式</strong>(\w+)（<strong>HTTPS</strong>）`)
		api.Method = pickSubMatchString(apiMethodRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)

		apiURLRegexp := regexp.MustCompile(`<strong>请求地址</strong>(.+?)</p>`)
		api.URL = pickSubMatchString(apiURLRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)

		if strings.Contains(rawHtmlSection, `<strong>请求示例</strong>`) {
			api.ReqJson = subDoc.Find("pre > code").Eq(0).Text()
		}

		if strings.Contains(rawHtmlSection, `<strong>返回结果</strong>`) {
			if subDoc.Find("pre > code").Eq(1).Text() != "" {
				api.RespJson = subDoc.Find("pre > code").Eq(1).Text()
			} else if subDoc.Find("pre > code").Eq(0).Text() != "" {
				api.RespJson = subDoc.Find("pre > code").Eq(0).Text()
			}
		}

		api.MethodCaml = strcase.ToCamel(strings.ToLower(api.Method))
		api.URL = strings.ReplaceAll(api.URL, "https://qyapi.weixin.qq.com", "")

		//传入参数，每行有3个字段，利用字段数量来提取
		if strings.Contains(rawHtmlSection, `<strong>参数说明</strong>`) {
			fieldsRegexp := regexp.MustCompile(`<tr>\n<td>(.+?)</td>\n<td>(.+?)</td>\n<td>(.+?)</td>\n</tr>`)
			fields := make([]Field, 0)
			allMatches := fieldsRegexp.FindAllStringSubmatch(rawHtmlSection, -1)
			for _, matches := range allMatches {
				field := Field{
					Name:       pickSubMatchString(matches, 4, 1),
					IsRequired: false,
					Desc:       pickSubMatchString(matches, 4, 3),
				}
				if pickSubMatchString(matches, 4, 2) == "是" {
					field.IsRequired = true
				}
				fields = append(fields, field)
			}

			api.ReqFields = fields
		}

		//响应参数，每行有2个字段，利用字段数量来提取
		if strings.Contains(rawHtmlSection, `<strong>参数说明</strong>`) && len(api.ReqFields) == 0 {
			fieldsRegexp := regexp.MustCompile(`<tr>\n<td>(.+?)</td>\n<td>(.+?)</td>\n</tr>`)
			fields := make([]Field, 0)
			allMatches := fieldsRegexp.FindAllStringSubmatch(rawHtmlSection, -1)
			for _, matches := range allMatches {
				field := Field{
					Name: pickSubMatchString(matches, 3, 1),
					Desc: pickSubMatchString(matches, 3, 2),
				}
				fields = append(fields, field)
			}

			api.RespFields = fields
		}

		apis = append(apis, api)
		fmt.Println(jsoniter.MarshalToString(api))
		fmt.Printf("第%d个h2完成\n", index)
	}

	result, err := generateCode(apis)
	if err != nil {
		die("generateCode failed: %+v\n", err)
	}

	if savePath != "" {
		err = ioutil.WriteFile(savePath, result, os.ModePerm)
		if err != nil {
			die("ioutil.WriteFile failed: %+v\n", err)
		}
		fmt.Printf("保存文件成功:%s", savePath)
	} else {
		fmt.Println(result)
	}
}
