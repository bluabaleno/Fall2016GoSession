package main 

import "os"
import "fmt"
import "image"
import "image/jpeg"
import "strconv"

// func determineType(filename string) (fileType string) {
// 	j := 0

// 	for i := 0; i < len(filename); i++ {
// 		if string(filename[i]) == "." {
// 			j = i
// 		}
// 	}
// 	return filename[j:]
// }

func resizing(input image.Image, name string, w, h int) error {
	newImageModel := image.NewRGBA(image.Rect(0, 0, w, h))

	bounds := input.Bounds()

	maxX := bounds.Max.X
	maxY := bounds.Max.Y

	ratioX := float32(((maxX)/w))
	ratioY := float32(((maxY)/h))

	i := 0
	for y := bounds.Min.Y; y < maxY; y++{
		for x := bounds.Min.X; x < maxX; x++{
			sampleAtX := int(ratioX * float32(x))
			sampleAtY := int(ratioY * float32(y))

			r,g,b,a := input.At(sampleAtX,sampleAtY).RGBA()
			newImageModel.Pix[i] = byte(r / 256)
			newImageModel.Pix[i+1] = byte(g / 256)
			newImageModel.Pix[i+2] = byte(b / 256)
			newImageModel.Pix[i+3] = byte(a / 1)
			i += 4
		}
	}

	destFile, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, os.ModePerm)
    if err != nil {
        return err
    }
    defer destFile.Close()

    if err = jpeg.Encode(destFile, newImageModel, &jpeg.Options{100}); err != nil {
        return err
    }
    return nil

}

func main() {
	path := os.Args
	// fileType := determineType(path[1])

	// // if fileType == ".jpeg" {
	// // 	fmt.Println("it is a " + fileType)

    f,err := os.Open(path[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    imageObj,err := jpeg.Decode(f)
    if err != nil {
        fmt.Println(err)
        return
    }

    w,_ := strconv.Atoi(path[3])
    h,_ := strconv.Atoi(path[4])
    resizing(imageObj, path[2], w, h)
	// }
}