package desocialize

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	pigo "github.com/esimov/pigo/core"
	"github.com/jpoz/glitch"
)

// Desocalizer will corrupt faces in images
type Desocalizer struct {
	Effect string
}

func (d Desocalizer) Desocialize(filename string, output_filename string) {
	_, currentfile, _, _ := runtime.Caller(0)
	cascadeFilePath := path.Join(path.Dir(currentfile), "facefinder")

	cascadeFile, err := ioutil.ReadFile(cascadeFilePath)
	check(err)

	src, err := pigo.GetImage(filename)
	check(err)

	sampleImg := pigo.RgbToGrayscale(src)

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,
	}
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y
	imgParams := pigo.ImageParams{sampleImg, rows, cols, cols}

	pigo := pigo.NewPigo()

	classifier, err := pigo.Unpack(cascadeFile)
	check(err)

	faces := classifier.RunCascade(imgParams, cParams)
	faces = classifier.ClusterDetections(faces, 0.2)

	fmt.Println(faces)

	gl, err := glitch.NewGlitch(filename)
	check(err)

	gl.Copy()

	for _, f := range faces {
		var qThresh float32 = 3.0

		if f.Q > qThresh {
			x := f.Col-f.Scale/2
			y := f.Row-f.Scale/2

			r := image.Rect(
				x, y,
				x + f.Scale,
				y + f.Scale,
			)

			gl.SetBounds(r)
			//gl.VerticalTransposeInput(f.Width(), f.Height(), true)
			//gl.TransposeInput(f.Width(), f.Height(), true)
			gl.HalfLifeRight(f.Scale, f.Scale)
			gl.PrismBurst()
		}
	}

	f, err := os.Create(output_filename)
	check(err)

	gl.Write(f)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
