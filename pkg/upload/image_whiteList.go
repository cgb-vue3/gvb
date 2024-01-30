package upload

import "akita/pkg/exlist"

// 记录图片后缀
var suffix = []any{
	".bmp",
	".jpg",
	".jpeg",
	".jpe",
	".jxr",
	".png",
	".tif",
	".tiff",
	".avif",
	".xbm",
	".pjp",
	".svgz",
	".ico",
	".svg",
	".jfif",
	".webp",
	".pjpeg",
	".gif",
	".iff",
	".ilbm",
	".ppm",
	".pcx",
	".xcf",
	".xpm",
	".psd",
	".sai",
	".psp",
	".ufo",
}

// White 白名单判断
func White(suf string) bool {
	// 判断某个值是否再某个列表中
	existence := exlist.Existence(suf, suffix)
	return existence
}
