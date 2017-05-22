package sms

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * 阿里大鱼短信工具模块
 * author: mliu
 * ================================================================================ */
type (
	AlidayuSmsTemplateParam struct {
		Code    string `form:"code" json:"code"`
		Product string `form:"product" json:"product"`
	}

	AlidayuSmsResult struct {
		Code      int32  `form:"err_code" json:"err_code"`
		Message   string `form:"msg" json:"msg"`
		Model     string `form:"model" json:"model"`
		RequestId string `form:"request_id" json:"request_id"`
		IsSuccess bool   `form:"is_success" json:"is_success"`
	}

	AlidayuSmsSendSuccessResponse struct {
		RequestId string                      `form:"request_id" json:"request_id"`
		Result    AlidayuSmsSendSuccessResult `form:"result" json:"result"`
	}

	AlidayuSmsSendSuccessResult struct {
		Code    int32  `form:"err_code" json:"err_code"`
		Message string `form:"msg" json:"msg"`
		Model   string `form:"model" json:"model"`
		Success bool   `form:"success" json:"success"`
	}

	AlidayuSmsSendErrorResponse struct {
		Result AlidayuSmsSendErrorResult `form:"error_response" json:"error_response"`
	}

	AlidayuSmsSendErrorResult struct {
		Code      int32  `form:"code" json:"code"`
		Message   string `form:"msg" json:"msg"`
		RequestId string `form:"request_id" json:"request_id"`
	}
)

type alidayuSms struct {
	AppKey          string `form:"app_key" json:"app_key"`
	AppSecret       string `form:"app_secret" json:"app_secret"`
	Method          string `form:"method" json:"method"`
	Format          string `form:"format" json:"format"`
	RecNum          string `form:"rec_num" json:"rec_num"`
	Simplify        string `form:"simplify" json:"simplify"`
	SmsFreeSignName string `form:"sms_free_sign_name" json:"sms_free_sign_name"`
	SmsTemplateCode string `form:"sms_template_code" json:"sms_template_code"`
	SmsParam        string `form:"sms_param" json:"sms_param"`
	SmsType         string `form:"sms_type" json:"sms_type"`
	SignMethod      string `form:"sign_method" json:"sign_method"`
	Timestamp       string `form:"timestamp" json:"timestamp"`
	Version         string `form:"v" json:"v"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建阿里大鱼短信结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewAlidayunSms(appKey, appSecret, signName string) *aliyunSms {
	dayuSms := new(alidayuSms)
	dayuSms.AppKey = appKey
	dayuSms.AppSecret = appSecret
	dayuSms.Method = "alibaba.aliqin.fc.sms.num.send"
	dayuSms.Format = "json"
	dayuSms.Simplify = "true"
	dayuSms.SmsType = "normal"
	dayuSms.SmsFreeSignName = signName
	dayuSms.SignMethod = "md5"
	dayuSms.SmsTemplateCode = ""
	dayuSms.SmsParam = ""
	dayuSms.RecNum = ""
	dayuSms.Timestamp = glib.TimeToString(time.Now())
	dayuSms.Version = "2.0"

	return dayuSms
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 发送手机信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) Send(mobiles string) (*AlidayuSmsResult, error) {
	result := new(AlidayuSmsResult)
	result.IsSuccess = false

	if len(mobiles) == 0 {
		return result, errors.New("手机号不能为空")
	}

	if len(s.SmsTemplateCode) == 0 || len(s.SmsParam) == 0 || len(s.SmsFreeSignName) == 0 {
		return nil, errors.New("参数不正确")
	}

	//接收手机号码
	s.RecNum = mobiles

	//签名请求参数
	requestString := s.GetRequestString()

	//发起Http请求
	if response, err := glib.HttpPost("http://gw.api.taobao.com/router/rest", requestString); err != nil {
		result.Message = err.Error()
		return result, err
	} else {
		if isSuccess := !strings.Contains(response, "error_response"); isSuccess {
			successResponse := new(AlidayuSmsSendSuccessResponse)
			glib.FromJson(response, successResponse)

			result.Code = successResponse.Result.Code
			result.Message = successResponse.Result.Message
			result.Model = successResponse.Result.Model
			result.RequestId = successResponse.RequestId
			result.IsSuccess = successResponse.Result.Success
		} else {
			errorResponse := new(AlidayuSmsSendErrorResponse)
			glib.FromJson(response, errorResponse)

			result.Code = errorResponse.Result.Code
			result.Message = errorResponse.Result.Message
			result.RequestId = errorResponse.Result.RequestId
		}
	}

	return result, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置模版码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) SetTemplateCode(code string) {
	s.SmsTemplateCode = code
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置模版参数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) SetTemplateParam(templateParam AlidayuSmsTemplateParam) {
	if jsonString, err := glib.ToJson(templateParam); err == nil {
		s.SmsParam = jsonString
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置签名字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) SetSignName(signName string) {
	s.SmsFreeSignName = signName
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取请求字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) GetRequestString() string {
	params := s.toDict()

	//Md5签名串
	sign := s.Sign(params)

	//参数值url编码
	var options []string = make([]string, 0)
	for k, v := range params {
		item := fmt.Sprintf("%s=%s", k, url.QueryEscape(v))
		options = append(options, item)
	}

	//把签名拼接到参数
	options = append(options, fmt.Sprintf("%s=%s", "sign", url.QueryEscape(sign)))

	//用&链接请求参数
	return strings.Join(options, "&")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 签名算法
 * params里的每个Value都需要进行url编码
 * fmt.Sprintf("%s=%s", key, url.QueryEscape(value))
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) Sign(params map[string]string) string {
	var keys []string = make([]string, 0)
	var values []string = make([]string, 0)

	//请求参数排序（字母升序）
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	//拼接KeyValue字符串
	for _, key := range keys {
		if len(params[key]) > 0 {
			values = append(values, key)         //Key
			values = append(values, params[key]) //Value
		}
	}
	paramString := strings.Join(values, "")

	//Md5签名（在拼接的字符串头尾附加上api密匙，然后md5，md5串是大写）
	paramString = fmt.Sprintf("%s%s%s", s.AppSecret, paramString, s.AppSecret)
	sign := Md5(paramString)

	return strings.ToUpper(sign)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取参数字典
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *alidayuSms) toDict() {
	var params map[string]string = make(map[string]string, 0)
	params["app_key"] = s.AppKey
	params["method"] = s.Method
	params["format"] = s.Format
	params["simplify"] = s.Simplify
	params["sms_type"] = s.SmsType
	params["sms_free_sign_name"] = s.SmsFreeSignName
	params["sign_method"] = s.SignMethod
	params["sms_template_code"] = s.SmsTemplateCode
	params["sms_param"] = s.SmsParam
	params["rec_num"] = s.RecNum
	params["v"] = s.Version
	params["timestamp"] = s.Timestamp

	return params
}
