package main 

import "os"
import "fmt"
import "image"
import "image/jpeg"
import "strconv"
import "math"

func resizing(input image.Image, name string, w, h int) error {
	newImageModel := image.NewRGBA(image.Rect(0, 0, w, h))

	inBound := input.Bounds()

	maxX := (inBound.Max.X)
	maxY := (inBound.Max.Y)

	ratioX := (float32(maxX)/float32(w))
	ratioY := (float32(maxY)/float32(h))

	i := 0
	for y := newImageModel.Bounds().Min.Y; y < newImageModel.Bounds().Max.Y; y++{
		for x := newImageModel.Bounds().Min.X; x < newImageModel.Bounds().Max.X; x++{
			coordX := ratioX * float32(x)
            coordY := ratioY * float32(y)

			sampleAtX := int(coordX)
			sampleAtY := int(coordY)

            leftCenter := float32(math.Floor(float64(coordX) - 0.5) + 0.5)
            distX := coordX - leftCenter

            leftRatio := 1 - distX
			rightRatio := distX

			var redTop, greenTop, blueTop, alphaTop int

			{
	            r1,g1,b1,a1 := input.At(sampleAtX,sampleAtY).RGBA()
	            r2,g2,b2,a2 := input.At(sampleAtX+1, sampleAtY).RGBA()

	            leftRed := float32(r1/256)
	            rightRed := float32(r2/256)
	            leftGreen := float32(g1/256)
	            rightGreen := float32(g2/256)
	            leftBlue := float32(b1/256)
	            rightBlue := float32(b2/256)
	            leftAlpha := float32(a1/256)
	            rightAlpha := float32(a2/256)

	            redTop = int(leftRed * leftRatio + rightRed * rightRatio)
	            greenTop = int(leftGreen * leftRatio + rightGreen * rightRatio)
	            blueTop = int(leftBlue * leftRatio + rightBlue * rightRatio)
	            alphaTop = int(leftAlpha * leftRatio + rightAlpha * rightRatio)
        	}

			newImageModel.Pix[i] = byte(redTop)
			newImageModel.Pix[i+1] = byte(greenTop)
			newImageModel.Pix[i+2] = byte(blueTop)
			newImageModel.Pix[i+3] = byte(alphaTop)
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
}