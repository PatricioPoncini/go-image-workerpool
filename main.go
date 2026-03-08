package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	inputDir  = "images"
	outputDir = "output"
	quality   = 20
)

func optimizeImage(fileName string) error {
	filePath := filepath.Join(inputDir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	outPath := filepath.Join(outputDir, "opt_"+fileName)
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	options := jpeg.Options{Quality: quality}
	return jpeg.Encode(outFile, img, &options)
}

func main() {
	files, _ := filepath.Glob(filepath.Join(inputDir, "*.jpg"))
	if len(files) == 0 {
		fmt.Println("No images found in 'images' folder. Please add some .jpg files.")
		return
	}
	_ = os.MkdirAll(outputDir, os.ModePerm)

	fmt.Printf("System: %d cores | Images to process: %d\n\n", runtime.NumCPU(), len(files))

	fmt.Println("Starting Sequential Optimization...")
	start := time.Now()
	for _, file := range files {
		_ = optimizeImage(filepath.Base(file))
	}
	seqDuration := time.Since(start)
	fmt.Printf("Sequential took: %v\n\n", seqDuration)

	_ = os.RemoveAll(outputDir)
	_ = os.MkdirAll(outputDir, os.ModePerm)

	numWorkers := runtime.NumCPU()
	jobs := make(chan string, len(files))
	var wg sync.WaitGroup

	fmt.Printf("Starting Worker Pool with %d workers...\n", numWorkers)
	start = time.Now()

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for fileName := range jobs {
				_ = optimizeImage(fileName)
			}
		}()
	}

	for _, file := range files {
		jobs <- filepath.Base(file)
	}
	close(jobs)
	wg.Wait()

	poolDuration := time.Since(start)
	fmt.Printf("Worker Pool took: %v\n", poolDuration)

	speedup := float64(seqDuration) / float64(poolDuration)
	fmt.Printf("\nPerformance gain: %.2fx faster\n", speedup)
}
