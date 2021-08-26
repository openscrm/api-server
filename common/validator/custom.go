package validator

import (
	"database/sql/driver"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/thoas/go-funk"
	"openscrm/app/constants"
	"openscrm/app/models"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"time"
)

type CustomValidator struct {
	Once     sync.Once
	Validate *val.Validate
	Trans    ut.Translator
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.LazyInit()
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *CustomValidator) Engine() interface{} {
	v.LazyInit()
	return v.Validate
}

func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(constants.StringArrayField); ok {
		return []string(valuer)
	}

	if valuer, ok := field.Interface().(driver.Valuer); ok {
		v, err := valuer.Value()
		if err == nil {
			return v
		}
	}

	return ""
}

func (v *CustomValidator) LazyInit() {
	v.Once.Do(func() {
		v.Validate = val.New()
		v.Validate.SetTagName("validate")
		v.Validate.RegisterValidation("int64", ValidateInt64Filed)
		v.Validate.RegisterValidation("phone", ValidatePhoneFiled)
		v.Validate.RegisterValidation("corp_id", ValidateCorpIDFiled)
		v.Validate.RegisterValidation("word", ValidateWordFiled)
		v.Validate.RegisterValidation("ext_id", ValidateExtIDFiled)
		v.Validate.RegisterValidation("time", ValidateTimeFiled)
		v.Validate.RegisterValidation("date", ValidateDateFiled)
		v.Validate.RegisterValidation("boolean", ValidateBooleanFiled)
		v.Validate.RegisterValidation("weekday", ValidateWeekdayFiled)
		v.Validate.RegisterValidation("validAdmin", ValidateAdmins)

		v.Trans, _ = ut.New(en.New(), zh.New()).GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(v.Validate, v.Trans)

		//自定义错误内容
		v.Validate.RegisterTranslation("required", v.Trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe val.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})

		v.Validate.RegisterTranslation("int64", v.Trans, func(ut ut.Translator) error {
			return ut.Add("int64", "{0} 必须是int64整型", true)
		}, func(ut ut.Translator, fe val.FieldError) string {
			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				fmt.Printf("warning: error translating FieldError: %#v", fe)
				return fe.Error()
			}
			return t
		})

		v.Validate.RegisterTranslation("boolean", v.Trans,
			func(ut ut.Translator) (err error) {
				if err = ut.Add("boolean", "{0} 必须是1或者2，1代表是，2代表否", true); err != nil {
					return
				}
				return
			}, func(ut ut.Translator, fe val.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.Error()
				}
				return t
			})

		v.Validate.RegisterTranslation("weekday", v.Trans,
			func(ut ut.Translator) (err error) {
				if err = ut.Add("weekday", "{0} 必须是[\"周一\", \"周二\", \"周三\", \"周四\", \"周五\", \"周六\", \"周日\"]其中之一", false); err != nil {
					return
				}
				return
			}, func(ut ut.Translator, fe val.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.Error()
				}
				return t
			})

		v.Validate.RegisterTranslation("validAdmin", v.Trans,
			func(ut ut.Translator) (err error) {
				if err = ut.Add("validAdmin", "{0} 需均为管理员", false); err != nil {
					return
				}
				return
			}, func(ut ut.Translator, fe val.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.Error()
				}
				return t
			})

		//v.Validate.RegisterCustomTypeFunc(ValidateValuer, constants.TimeField{}, constants.StringArrayField{})
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func ValidateInt64Filed(fl val.FieldLevel) bool {
	_, err := strconv.ParseInt(fl.Field().String(), 10, 64)
	if err != nil {
		return false
	}
	return true
}

func ValidatePhoneFiled(fl val.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
	return ok
}

func ValidateWordFiled(fl val.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, fl.Field().String())
	return ok
}

const (
	ExtIDLen  = 32
	CorpIDLen = 18
)

// ValidateExtIDFiled
// len(extid)=18
func ValidateExtIDFiled(fl val.FieldLevel) bool {
	matchString, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, fl.Field().String())
	return matchString && (len(fl.Field().String()) == ExtIDLen)
}

// ValidateCorpIDFiled
// len(corpid)=18
func ValidateCorpIDFiled(fl val.FieldLevel) bool {
	matchString, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, fl.Field().String())
	return matchString && (len(fl.Field().String()) == CorpIDLen)
}

// ValidateTimeFiled 验证字符串时间
func ValidateTimeFiled(fl val.FieldLevel) bool {
	_, err := time.Parse(constants.TimeLayout, fl.Field().String())
	if err != nil {
		return false
	}

	return true
}

// ValidateBooleanFiled 验证数字布尔型
func ValidateBooleanFiled(fl val.FieldLevel) bool {
	return funk.ContainsInt64([]int64{1, 2}, fl.Field().Int())
}

// ValidateWeekdayFiled 验证星期
func ValidateWeekdayFiled(fl val.FieldLevel) bool {
	return funk.ContainsString([]string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}, fl.Field().String())
}

// ValidateAdmins 验证管理员列表中的id是否均存在
func ValidateAdmins(fl val.FieldLevel) bool {
	if admins, ok := fl.Field().Interface().([]string); ok {
		var total int64
		err := models.DB.Model(&models.Staff{}).
			Where("id in ? and role_type = ?", admins, constants.RoleTypeAdmin).Count(&total).Error
		if err != nil || total != int64(len(admins)) {
			return false
		}
	}
	return true
}

// ValidateDateFiled 验证字符串日期
func ValidateDateFiled(fl val.FieldLevel) bool {
	_, err := time.Parse(constants.DateLayout, fl.Field().String())
	if err != nil {
		return false
	}

	return true
}
