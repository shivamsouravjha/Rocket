package test

import (
	imagecompress "rocket/helper/imageCompress"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Integration Test
func TestCompressImagesIntegration(t *testing.T) {
	err := imagecompress.CompressImages("218")
	assert.NoError(t, err)
}
