# go-dms
A small library for converting Decimal Degrees to Degrees, Minutes, Seconds coordinates

Efficiently converting coordinates between DD and DMS

Calculate Autostar Lat Long

## Installation

`go get -u github.com/ddefrancesco/go-dms/dms`

**test.go:**
```go
package main

import (
    "github.com/ddefrancesco/go-dms/dms"
    "fmt"
    "time"
    "log"
)

func main() {
    start := time.Now()
    dmsCoordinate, err := dms.NewDMS(dms.LatLon{Latitude: 2.21893, Longitude: 1.213905})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("DMS coordinates: %+v\n", dmsCoordinate.String()) 
    end := time.Now()
    fmt.Printf("Function took %f seconds.\n", end.Sub(start).Seconds())

    dms, _ := NewDMS(DecimalDegrees{Latitude: 41.82326, Longitude: 12.44474})

	// Calc AS latitude
    autostarLat := dms.AutostarLatitude(dms.Latitude)
    fmt.Printf("Autostar latitude: %+v\n", autostarLat) 
    //assert.Equal(t, "41*49", autostarLat)
    // Calc AS longitude
    autostarLong := dms.AutostarLongitude(dms.Longitude)
    //assert.Equal(t, "348*26", autostarLong)
    fmt.Printf("Autostar longitude: %+v\n", autostarLong) 
}
```
**>> go run test.go**

**Output:**
```
    DMS coordinates:
    2°13'8.148000" N, 1°12'50.058000" E
    Function took 0.000049 seconds.

    Autostar latitude: 41*49
    Autostar longitude: 348*26
```

**>> GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .** (Compiling as WebAssembly module for Node, run command in the same directory as `test.go`)

Golang WebAssemlby Wiki: https://github.com/golang/go/wiki/WebAssembly

**Output:**
```
    DMS coordinates:
    2°13'8.148000" N, 1°12'50.058000" E
    Function took 0.000478 seconds.
```


