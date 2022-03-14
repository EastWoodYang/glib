package glib

/*
import (
	"bytes"
	"errors"
	"image/png"
)

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)
*/

/* ================================================================================
 * QR
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取二维码图片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func QrCode(content string, args ...int) ([]byte, error) {
	if len(content) == 0 {
		return nil, errors.New("qr content is empty")
	}
	var dataBuf bytes.Buffer
	qrCode, err := qr.Encode(content, qr.M, qr.Unicode)
	if err != nil {
		return nil, err
	}

	//默认宽高
	var w, h int
	w = 300
	h = 300

	//宽高参数
	argsCount := len(args)
	if argsCount == 1 {
		w = args[0]
		h = w
	} else if argsCount == 2 {
		w = args[0]
		h = args[1]
	}

	//二维码图片数据
	imageData, err := barcode.Scale(qrCode, w, h)
	if err != nil {
		return nil, err
	}

	//Png图像编码
	if err := png.Encode(&dataBuf, imageData); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}
*/

