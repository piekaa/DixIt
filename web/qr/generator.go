package qr

import "github.com/skip2/go-qrcode"

func Create(url string) (image []byte, err error) {
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	return png, err;
}
