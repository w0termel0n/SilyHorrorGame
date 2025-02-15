package assets

import (
	"image"
	"image/color" // ✅ Correct import for color
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	frameDelay   = 5 // Change frame every 5 ticks
)

var frames []*ebiten.Image

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if len(frames) == 0 {
		return
	}

	screen.Fill(color.White) // Set background to white

	i := (g.count / frameDelay) % len(frames) // Cycle through frames
	img := frames[i]

	// Get the original size of the image
	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()

	// Calculate scaling factors
	scaleX := float64(screenWidth) / float64(imgWidth)
	scaleY := float64(screenHeight) / float64(imgHeight)

	// Keep the aspect ratio by taking the smaller scale factor
	scale := scaleX
	if scaleY < scaleX {
		scale = scaleY
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)                                                // Scale the image
	op.GeoM.Translate(-float64(imgWidth)*scale/2, -float64(imgHeight)*scale/2) // Center
	op.GeoM.Translate(screenWidth/2, screenHeight/2)                           // Move to screen center

	screen.DrawImage(img, op)
}

// ✅ Implement missing Layout method
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadImagesFromDir(path string) []*ebiten.Image {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var images []*ebiten.Image
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.Open(path + "/" + file.Name())
		if err != nil {
			log.Println("Failed to open file:", file.Name(), err)
			continue
		}
		img, _, err := image.Decode(f)
		f.Close()
		if err != nil {
			log.Println("Failed to decode image:", file.Name(), err)
			continue
		}
		images = append(images, ebiten.NewImageFromImage(img))
	}
	return images
}

func main() {
	frames = loadImagesFromDir("assets/TestAnimationForGoLang")
	if len(frames) == 0 {
		log.Fatal("No images found in assets/TestAnimationForGoLang")
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Custom Animation")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
