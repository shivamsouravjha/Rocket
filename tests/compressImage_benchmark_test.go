package test

import (
	"fmt"
	imagecompress "rocket/helper/imageCompress"
	"testing"
)

func BenchmarkCompressImages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := imagecompress.CompressImages("208")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

//APP_ENV=test go test -run=10000 -bench=BenchmarkCompressImages -benchmem
