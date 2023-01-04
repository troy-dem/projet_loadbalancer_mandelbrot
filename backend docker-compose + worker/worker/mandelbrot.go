package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"sync"
)

func png_generator(resolution_x, resolution_y, start_position_x, start_position_y, quantize_length, max_iteration float64, colormap [][3]int, wg *sync.WaitGroup) {
	img := image.NewNRGBA(image.Rect(0, 0, int(resolution_x), int(resolution_y)))
	for y := 0.0; y < resolution_y; y++ {
		for x := 0.0; x < resolution_x; x++ {
			new_x := (float64(x)/float64(resolution_x))*quantize_length + start_position_x
			new_y := start_position_y - (float64(y)/float64(resolution_y))*quantize_length
			iteration := mandelbrot(max_iteration, new_x, new_y)
			pixel_color := colorize(iteration, max_iteration, colormap)
			img.Set(int(x), int(y), color.NRGBA{
				R: uint8(pixel_color[0]),
				G: uint8(pixel_color[1]),
				B: uint8(pixel_color[2]),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	defer wg.Done()
}

func mandelbrot(max_iteration, c_real, c_imaginary float64) (iteration float64) {
	z_real := 0.0
	z_imaginary := 0.0
	new_xz := 0.0
	for iteration := 1.0; iteration < max_iteration; iteration++ {
		if iteration != 1.0 {
			new_xz = (z_real * z_real) - (z_imaginary * z_imaginary) + c_real
			z_imaginary = (2 * z_real * z_imaginary) + c_imaginary
			z_real = new_xz
		}
		modulus := math.Sqrt((z_real * z_real) + (z_imaginary * z_imaginary))
		if modulus > 2.0 {
			return iteration
		}
	}
	iteration = max_iteration
	return iteration
}

func colorize(iteration, max_iteration float64, colormap [][3]int) (pixel_color [3]int) {
	if iteration == max_iteration {
		pixel_color = [3]int{0, 0, 0}
		return pixel_color
	}
	pixel_color = colormap[int(iteration)]
	return pixel_color
	//ratio := iteration / max_iteration
	// if ratio >= 1.0 {
	// 	pixel_color = [3]int{0, 0, 0}
	// 	return pixel_color
	// }
	// if ratio >= 0.8 {
	// 	pixel_color = [3]int{255, 0, 0}
	// 	return pixel_color
	// }
	// if ratio >= 0.6 {
	// 	pixel_color = [3]int{0, 255, 0}
	// 	return pixel_color
	// }
	// if ratio >= 0.4 {
	// 	pixel_color = [3]int{0, 255, 255}
	// 	return pixel_color
	// }
	// if ratio >= 0.2 {
	// 	pixel_color = [3]int{255, 0, 255}
	// 	return pixel_color
	// }
	// pixel_color = [3]int{255, 255, 0}
	// return pixel_color
}

// func randomcolor(max_iteration float64) (colormap [][3]int) {
// 	for iteration := 0.0; iteration < max_iteration; iteration++ {
// 		colormap = append(colormap, [3]int{rand.Intn(230) + 24, rand.Intn(230) + 24, rand.Intn(230) + 24})
// 	}
// 	return colormap
// }
