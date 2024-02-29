package resouse

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

// 保存所有静态资源
var Images = make(map[string]*ebiten.Image)

func init() {
	dir := "./assets"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 如果文件不是目录，尝试加载它
			img, err := loadImage(path)
			if err != nil {
				fmt.Println("Failed to load image:", path)
			} else {
				// 将加载的图像存储在映射中，键是文件名
				Images[filepath.Base(path)] = img
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func loadImage(filePath string) (*ebiten.Image, error) {
	// 打开图像文件
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 解码图像文件
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// 将图像转换为 Ebiten 图像
	eImg := ebiten.NewImageFromImage(img)

	return eImg, nil
}
