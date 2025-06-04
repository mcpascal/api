package validator

import (
	"api/internal/responses"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Trans ut.Translator
)

func Setup(locale string) (err error) {
	color.Red("** validator init start**")
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, Trans)
		}

		// 注册自定义验证规则
		registerValidatorFunc(v, "mobile", "非法手机号", ValidateMobile)

		return
	}
	return
}

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	//如果返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		responses.Fail(c, http.StatusBadRequest, "", err)
		return
	}
	errMap := removeTopStruct(errs.Translate(Trans))
	var errstr string
	for k, v := range errMap {
		if len(errstr) > 0 {
			errstr += ";"
		}
		errstr += k + ":" + v
	}
	responses.Fail(c, http.StatusBadRequest, "", errors.New(errstr))
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		// 从文本的逗号开始切分   处理后"mobile": "mobile为必填字段"  处理前: "PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// Func myvalidator.ValidateMobile
type Func func(fl validator.FieldLevel) bool

// RegisterValidatorFunc 注册自定义校验tag
func registerValidatorFunc(v *validator.Validate, tag string, msgStr string, fn Func) {
	// 注册tag自定义校验
	_ = v.RegisterValidation(tag, validator.Func(fn))
	//自定义错误内容
	_ = v.RegisterTranslation(tag, Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})

	return
}
