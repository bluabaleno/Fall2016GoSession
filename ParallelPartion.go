package main 

import "os"
import "fmt"
import "image"
import "image/jpeg"
import "strconv"
import "math"
import "time"

func resizing(input image.Image, name string, w, h int) error {
    
    startTime := time.Now().Unix()

	newImageModel := image.NewRGBA(image.Rect(0, 0, w, h))

	inBound := input.Bounds()

	maxX := (inBound.Max.X)
	maxY := (inBound.Max.Y)

	ratioX := (float32(maxX)/float32(w))
	ratioY := (float32(maxY)/float32(h))

    SignalChannel := make(chan int, 2)


	go func() {
        i := 0
        for y := newImageModel.Bounds().Min.Y; y < newImageModel.Bounds().Max.Y/2; y++{
    		for x := newImageModel.Bounds().Min.X; x < newImageModel.Bounds().Max.X; x++{

    			//Calculating horizontal ratios
    			coordX := ratioX * float32(x)
                coordY := ratioY * float32(y)

    			sampleAtX := int(coordX)
    			sampleAtY := int(coordY)

                leftCenter := float32(math.Floor(float64(coordX) - 0.5) + 0.5)
                distX := coordX - leftCenter

                leftRatio := 1 - distX
    			rightRatio := distX

                var redTop int
                var greenTop int
                var blueTop int
                var alphaTop int

    			//Interpolationg the top two pixels
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

            	//Interpolating the bottom two pixels

                var redBottom int
                var greenBottom int
                var blueBottom int
                var alphaBottom int

                {
                    r1,g1,b1,a1 := input.At(sampleAtX,sampleAtY+1).RGBA()
                    r2,g2,b2,a2 := input.At(sampleAtX+1, sampleAtY+1).RGBA()

                    leftRed := float32(r1/256)
                    rightRed := float32(r2/256)
                    leftGreen := float32(g1/256)
                    rightGreen := float32(g2/256)
                    leftBlue := float32(b1/256)
                    rightBlue := float32(b2/256)
                    leftAlpha := float32(a1/256)
                    rightAlpha := float32(a2/256)

                    redBottom = int(leftRed * leftRatio + rightRed * rightRatio)
                    greenBottom = int(leftGreen * leftRatio + rightGreen * rightRatio)
                    blueBottom = int(leftBlue * leftRatio + rightBlue * rightRatio)
                    alphaBottom = int(leftAlpha * leftRatio + rightAlpha * rightRatio)

                }
                // Calculating the ratio of the top and bottom ratios
                TopCenter := float32(math.Floor(float64(coordY) - 0.5) + 0.5)
                distY := coordY - TopCenter

                TopRatio := 1 - distY
                BottomRatio := distY

                renderRed := int(float32(redTop) * TopRatio + float32(redBottom) * BottomRatio)
                renderGreen := int(float32(greenTop) * TopRatio + float32(greenBottom) * BottomRatio)
                renderBlue := int(float32(blueTop) * TopRatio + float32(blueBottom) * BottomRatio)
                renderAlpha := int(float32(alphaTop) * TopRatio + float32(alphaBottom) * BottomRatio)

                newImageModel.Pix[i] = byte(renderRed)
                newImageModel.Pix[i+1] = byte(renderGreen)
                newImageModel.Pix[i+2] = byte(renderBlue)
                newImageModel.Pix[i+3] = byte(renderAlpha)

    			i += 4

    		}
    	}
        SignalChannel <- 1;

    }()

    go func() {
        i := newImageModel.Bounds().Max.Y/2 * newImageModel.Bounds().Max.X * 4
        for y := newImageModel.Bounds().Max.Y/2 + 1; y < newImageModel.Bounds().Max.Y; y++{
            for x := newImageModel.Bounds().Min.X; x < newImageModel.Bounds().Max.X; x++{

                //Calculating horizontal ratios
                coordX := ratioX * float32(x)
                coordY := ratioY * float32(y)

                sampleAtX := int(coordX)
                sampleAtY := int(coordY)

                leftCenter := float32(math.Floor(float64(coordX) - 0.5) + 0.5)
                distX := coordX - leftCenter

                leftRatio := 1 - distX
                rightRatio := distX

                var redTop int
                var greenTop int
                var blueTop int
                var alphaTop int

                //Interpolationg the top two pixels
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

                //Interpolating the bottom two pixels

                var redBottom int
                var greenBottom int
                var blueBottom int
                var alphaBottom int

                {
                    r1,g1,b1,a1 := input.At(sampleAtX,sampleAtY+1).RGBA()
                    r2,g2,b2,a2 := input.At(sampleAtX+1, sampleAtY+1).RGBA()

                    leftRed := float32(r1/256)
                    rightRed := float32(r2/256)
                    leftGreen := float32(g1/256)
                    rightGreen := float32(g2/256)
                    leftBlue := float32(b1/256)
                    rightBlue := float32(b2/256)
                    leftAlpha := float32(a1/256)
                    rightAlpha := float32(a2/256)

                    redBottom = int(leftRed * leftRatio + rightRed * rightRatio)
                    greenBottom = int(leftGreen * leftRatio + rightGreen * rightRatio)
                    blueBottom = int(leftBlue * leftRatio + rightBlue * rightRatio)
                    alphaBottom = int(leftAlpha * leftRatio + rightAlpha * rightRatio)

                }
                // Calculating the ratio of the top and bottom ratios
                TopCenter := float32(math.Floor(float64(coordY) - 0.5) + 0.5)
                distY := coordY - TopCenter

                TopRatio := 1 - distY
                BottomRatio := distY

                renderRed := int(float32(redTop) * TopRatio + float32(redBottom) * BottomRatio)
                renderGreen := int(float32(greenTop) * TopRatio + float32(greenBottom) * BottomRatio)
                renderBlue := int(float32(blueTop) * TopRatio + float32(blueBottom) * BottomRatio)
                renderAlpha := int(float32(alphaTop) * TopRatio + float32(alphaBottom) * BottomRatio)

                newImageModel.Pix[i] = byte(renderRed)
                newImageModel.Pix[i+1] = byte(renderGreen)
                newImageModel.Pix[i+2] = byte(renderBlue)
                newImageModel.Pix[i+3] = byte(renderAlpha)

                i += 4

            }
        }
        SignalChannel <- 1;
    }()

    <-SignalChannel
    <-SignalChannel

    endTime := time.Now().Unix()

    duration := endTime - startTime

    fmt.Println(duration)

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