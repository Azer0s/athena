package marshaller

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMarshalString(t *testing.T) {
	buff := &bytes.Buffer{}

	err := Marshal("test", buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, "test", unmarshal)
}

func TestMarshalFloat64(t *testing.T) {
	buff := &bytes.Buffer{}

	err := Marshal(float64(math.Pi), buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, float64(math.Pi), unmarshal)
}

func TestMarshalFloat32(t *testing.T) {
	buff := &bytes.Buffer{}

	err := Marshal(float32(math.Pi), buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, float32(math.Pi), unmarshal)
}

func TestMarshalBool(t *testing.T) {
	buff := &bytes.Buffer{}

	err := Marshal(true, buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, true, unmarshal)
}

func TestMarshalInt(t *testing.T) {
	buff := &bytes.Buffer{}

	err := Marshal(20000, buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, 20000, unmarshal)
}

func TestMarshalInt2(t *testing.T) {
	buff := &bytes.Buffer{}

	err := Marshal(-20000, buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, -20000, unmarshal)
}

func TestMarshalStringSlice(t *testing.T) {
	buff := &bytes.Buffer{}

	val := []string{"foo", "bar", "hello", "world"}

	err := Marshal(val, buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, val, unmarshal)
}

func TestMarshalStringSlice2(t *testing.T) {
	buff := &bytes.Buffer{}

	val := []string{}

	err := Marshal(val, buff)
	if err != nil {
		return
	}

	unmarshal, err := Unmarshal(buff.Bytes())
	if err != nil {
		return
	}

	assert.Equal(t, val, unmarshal)
}
