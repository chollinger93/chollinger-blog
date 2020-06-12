package main

import (
	"fmt"
	"math"
)

// A Resource we're trying to access
type CoordinateData interface {
	calculateDistance(latTo, lonTo float64) float64
}

type GeospatialData struct {
	lat, lon float64
}

const earthRadius = float64(6371)
func (d GeospatialData) calculateDistance(latTo, lonTo float64) float64 {
	// Haversine distance: https://play.golang.org/p/MZVh5bRWqN
	var deltaLat = (latTo - d.lat) * (math.Pi / 180)
	var deltaLon = (lonTo - d.lon) * (math.Pi / 180)
	
	var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) + 
		math.Cos(d.lat * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
		math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
	var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
	
	return earthRadius * c
}

type CartesianPlaneData struct {
	x, y float64
}

func (d CartesianPlaneData) calculateDistance(xTo, yTo float64) float64 {
	// Simple 2-dimensional Euclidean distance 
	dx := (xTo - d.x)
	dy := (yTo - d.y)
	return math.Sqrt( dx*dx + dy*dy )
}

func main() {
	atlanta := GeospatialData{33.753746, -84.386330}
	distance := atlanta.calculateDistance(33.957409, -83.376801) // to Athens, GA
	fmt.Printf("The Haversine distance between Atlanta, GA and Athens, GA is %v\n", distance)

	pointA := CartesianPlaneData{1, 1}
	distanceA := pointA.calculateDistance(5, 5) 
	fmt.Printf("The Pythagorean distance from (1,1) to (5,5) is %v\n", distanceA)
}