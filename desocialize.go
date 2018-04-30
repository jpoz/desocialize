package desocialize

import (
	"image"
	"os"
	"path"
	"runtime"

	"github.com/jpoz/glitch"
	"github.com/lazywei/go-opencv/opencv"
)

// Desocalizer will corrupt faces in images
type Desocalizer struct {
	Effect string
}

func (d Desocalizer) Desocialize(filename string, output_filename string) {
	_, currentfile, _, _ := runtime.Caller(0)
	img := opencv.LoadImage(path.Join(path.Dir(currentfile), filename))

	cascade := opencv.LoadHaarClassifierCascade(
		path.Join(path.Dir(currentfile), "haarcascade_frontalface_alt.xml"),
	)
	faces := cascade.DetectObjects(img)

	gl, err := glitch.NewGlitch(filename)
	check(err)

	gl.Copy()

	for _, f := range faces {
		r := image.Rect(
			f.X(), f.Y(),
			f.X()+f.Width(),
			f.Y()+f.Height(),
		)
		gl.SetBounds(r)
		//gl.VerticalTransposeInput(f.Width(), f.Height(), true)
		//gl.TransposeInput(f.Width(), f.Height(), true)
		gl.PrismBurst()
		gl.HalfLifeRight(f.Height()*4, f.Width()/2)
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
