package glib

import (
	"errors"
)

/* ================================================================================
 * 错误异常
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 捕获函数执行时的异常
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Try(fnSource func(), fnError func(interface{})) (er error) {
	defer func() {
		if err := recover(); err != nil {
			fnError(err)
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("try func call errors")
			}
		}
	}()

	fnSource()

	return er
}



