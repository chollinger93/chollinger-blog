package main

import (
	"bytes"
	"fmt"
	"log"
	"math/cmplx"
	"os"
	"strings"
	"time"

	"encoding/binary"
)

func linspace(start, end float64, num int) []float64 {
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result
}

type triple struct {
	ix, iy int64
	c      int64
}

func calcRow(ix, iy int64, c complex128, threshold int64) triple {
	return triple{ix, iy, countIterationsUntilDivergent(c, threshold)}
}

func countIterationsUntilDivergent(c complex128, threshold int64) int64 {
	z := complex(0, 0)
	var ix int64 = 0
	for i := int64(0); i < threshold; i++ {
		ix = i
		z = (z * z) + c
		if cmplx.Abs(z) > 4 {
			return i
		}
	}
	return ix
}

func mandelbrot(threshold, density int64) [][]int64 {
	realAxis := linspace(-0.22, -0.219, 1000)
	imaginaryAxis := linspace(-0.70, -0.699, 1000)
	fmt.Printf("realAxis %v\n", len(realAxis))
	fmt.Printf("imaginaryAxis %v\n", len(imaginaryAxis))
	atlas := make([][]int64, len(realAxis))
	for i := range atlas {
		atlas[i] = make([]int64, len(imaginaryAxis))
	}
	// Make a buffered channel
	ch := make(chan triple, int64(len(realAxis))*int64(len(imaginaryAxis)))

	for ix, _ := range realAxis {
		go func(ix int) {
			for iy, _ := range imaginaryAxis {
				cx := realAxis[ix]
				cy := imaginaryAxis[iy]
				c := complex(cx, cy)
				res := calcRow(int64(ix), int64(iy), c, threshold)
				ch <- res
			}
		}(ix)
	}

	for i := int64(0); i < int64(len(realAxis))*int64(len(imaginaryAxis)); i++ {
		select {
		case res := <-ch:
			atlas[res.ix][res.iy] = res.c
		}
	}
	return atlas
}

func write(data []int64) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		log.Fatal(err)
	}
}

func write2(data []int64) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("data.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//buf := new(bytes.Buffer)
	//err = binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data)), " "), "[]") + "\n"); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	start := time.Now()
	data := mandelbrot(120, 500)
	fmt.Println(time.Since(start))
	//fmt.Printf("%v\n", data[0])
	for _, d := range data {
		write2(d)
	}
	fmt.Println("End")
}
