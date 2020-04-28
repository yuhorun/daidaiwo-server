package util

import (
	"github.com/afocus/captcha"
	"image/color"
)

func DefalutCaptcha() *captcha.Captcha {
	capImage := captcha.New()

	//设置字体
	if err := capImage.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	//设置图片大小
	capImage.SetSize(91, 41)
	//设置干扰强度
	capImage.SetDisturbance(captcha.NORMAL)
	//设置前景色
	capImage.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色
	capImage.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	return capImage
}
