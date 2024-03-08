package middleware

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"github.com/skip2/go-qrcode"
)

// 生成 QR code
func GenerateQRcode(code string, size int) ([]byte, int) {
	qrCode, err := qrcode.Encode(code, qrcode.Highest, size)
	if err != nil {
		return nil, errmsg.GenerateQRError
	}
	return qrCode, errmsg.SUCCESS
}
