package glib

import (
	"strings"
)

/* ================================================================================
 * 用户浏览代理
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

const (
	UA_DEVICE_PC string = "pc"

	UA_DEVICE_TABLE             string = "table"
	UA_DEVICE_TABLE_IPAD        string = "ipad"
	UA_DEVICE_TABLE_IPOD        string = "ipod"
	UA_DEVICE_TABLE_ANDROID_PAD string = "android pad"

	UA_DEVICE_MOBILE         string = "mobile"
	UA_DEVICE_MOBILE_IPHONE  string = "iphone"
	UA_DEVICE_MOBILE_ANDROID string = "android phone"
	UA_DEVICE_MOBILE_WINDOWS string = "windows phone"

	UA_OS_WINDOWS string = "windows"
	UA_OS_MAC     string = "mac"
	UA_OS_LINUX   string = "linux"

	UA_UNKNOW string = "unknow"
)

type (
	UserAgent struct {
		device  *UserAgentDevice //设备
		os      *UserAgentOs     //操作系统
		content string           //原始内容
	}

	UserAgentDevice struct {
		name      string // 设备名称（pc | table | mobile | unknow）
		childName string // 子名称（iphone | ipad | ipod | android | windows phone | nexus | unknow）
		content   string //设备内容
	}

	UserAgentOs struct {
		name    string // 操作系统名称（windows | mac | linux | unknow）
		content string //操作系统内容
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化UserAgent
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewUserAgent(content string) *UserAgent {
	userAgent := &UserAgent{
		device: &UserAgentDevice{
			name:      UA_UNKNOW,
			childName: UA_UNKNOW,
		},
		os: &UserAgentOs{
			name: UA_UNKNOW,
		},
		content: content,
	}

	userAgent.parse()

	return userAgent
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 解析
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) parse() {
	if len(s.content) == 0 {
		return
	}

	s.parseDevice()

	s.parseOs()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 解析设备
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) parseDevice() {
	deviceName := UA_UNKNOW
	if s.IsPc() {
		deviceName = UA_DEVICE_PC
	} else if s.IsTable() {
		deviceName = UA_DEVICE_TABLE
	} else if s.IsMobile() {
		deviceName = UA_DEVICE_MOBILE
	}

	s.device.name = deviceName
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 解析操作系统
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) parseOs() {
	operationName := UA_UNKNOW
	if strings.Index(s.content, "Windows") > 0 {
		operationName = UA_OS_WINDOWS
	} else if strings.Index(s.content, "Mac OS") > 0 {
		operationName = UA_OS_MAC
	} else if strings.Index(s.content, "Linux") > 0 {
		operationName = UA_OS_LINUX
	}

	s.os.name = operationName
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Pc
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsPc() bool {
	isPc := s.IsIeBrowser() ||
		s.IsSafariBrowser() ||
		s.IsChromeBrowser() ||
		s.IsFirfoxBrowser() ||
		s.IsOperaBrowser()

	return isPc
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Table
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsTable() bool {
	isTable := s.IsIpad() || s.IsAndroidPad()

	return isTable
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Mobile
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsMobile() bool {
	isMobile := s.IsIphone() ||
		s.IsIpad() ||
		s.IsIpod() ||
		s.IsAndroid() ||
		s.IsWindowsPhone() ||
		s.IsBlackBerry()

	return isMobile
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Iphone
 * Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsIphone() bool {
	return strings.Index(s.content, "iPhone") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Ipad
 * Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsIpad() bool {
	return strings.Index(s.content, "iPad") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Ipod
 * Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsIpod() bool {
	return strings.Index(s.content, "iPod") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Android
 * Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsAndroid() bool {
	return strings.Index(s.content, "Android") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Android Pad
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsAndroidPad() bool {
	return false
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Windows Phone
 * Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Titan)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsWindowsPhone() bool {
	return strings.Index(s.content, "Windows Phone") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否BlackBerry
 * Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, like Gecko) Version/6.0.0.337 Mobile Safari/534.1+
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsBlackBerry() bool {
	return strings.Index(s.content, "BlackBerry") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否IE Browser
 * Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsIeBrowser() bool {
	return strings.Index(s.content, "MSIE") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Safari Browser
 * Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsSafariBrowser() bool {
	return strings.Index(s.content, "Safari") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Chrome Browser
 * Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsChromeBrowser() bool {
	return strings.Index(s.content, "Chrome") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Firefox Browser
 * Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:66.0) Gecko/20100101 Firefox/66.0
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsFirfoxBrowser() bool {
	return strings.Index(s.content, "Firefox") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否Opera Browser
 * Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) IsOperaBrowser() bool {
	return strings.Index(s.content, "Opera") > 0
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取设备
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) GetDevice() *UserAgentDevice {
	return s.device
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取操作系统
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) GetOs() *UserAgentOs {
	return s.os
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取原始信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgent) GetContent() string {
	return s.content
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取设备名称
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgentDevice) GetName() string {
	return s.name
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取设备子名称
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgentDevice) GetChildName() string {
	return s.childName
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取设备信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgentDevice) GetContent() string {
	return s.content
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取操作系统名称
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgentOs) GetName() string {
	return s.name
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取操作系统信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *UserAgentOs) GetContent() string {
	return s.content
}
