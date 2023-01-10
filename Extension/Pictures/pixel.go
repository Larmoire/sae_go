package pictures

import (
		"fmt"
		"image"
		"image/jpeg"
		"os"
		"project-particles/config"
)


func Initpixel() {
		// damn important or else At(), Bounds() functions will
		// caused memory pointer error!!
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func Gettabpixels() (w int, h int,tab [][]int){
	imgfile, err := os.Open(config.General.Pictures)
	if err != nil {
			fmt.Println("image non trouvé")
			os.Exit(1)
	}	
	defer imgfile.Close()
	imgCfg, _, _ := image.DecodeConfig(imgfile)
	width := imgCfg.Width
	height := imgCfg.Height
	imgfile.Seek(0, 0)
	img, _, _ := image.Decode(imgfile)

	for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
					r, g, b, a := img.At(x, y).RGBA()
					tab = append(tab,[]int{x,y,int(r/257),int(g/257),int(b/257),int(a/257)})
			}
	}
	return  width, height,	tab
}
