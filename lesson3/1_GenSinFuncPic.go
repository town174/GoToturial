package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func GenSinFuncPic() {
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size))

	//填充背景色
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	//画图,延x轴开始
	for x := 0; x < size; x++ {
		//算出x
		s := float64(x) * 2 * math.Pi / size
		//算出Y
		y := size/2 - math.Sin(s)*size/2
		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	//创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatalln(err)
	}

	//将数据写入文件
	png.Encode(file, pic)

	//关闭文件
	file.Close()
}
