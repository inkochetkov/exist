package test

import (
	"os"
	"testing"

	"github.com/inkochetkov/exist/pkg/exist"
	"github.com/stretchr/testify/assert"
)

// TestDir check dir
func TestDir(t *testing.T) {

	t.Log("Test check dir")
	checkPath := "./../test"

	check := exist.CheckDir(checkPath)

	assert.Equal(t, true, check)

	t.Log("Test create dir")
	createPath := "./../test/test"

	err := exist.CreateDir(createPath)
	assert.NoError(t, err)

	err = os.Remove(createPath)
	assert.NoError(t, err)

	t.Log("Test init dir")
	initPath := "./../test/test"

	err = exist.InitDir(initPath)
	assert.NoError(t, err)

	err = os.Remove(initPath)
	assert.NoError(t, err)

	t.Log("Test success")

}

// TestFile check file
func TestFile(t *testing.T) {
	t.Log("Test check file")

	checkPath := "exist_test.go"
	check := exist.CheckFile(checkPath)

	assert.Equal(t, true, check)

	t.Log("Test create file")

	createPath := "test.go"

	file, err := exist.CreateFile(createPath)
	assert.NoError(t, err)

	file.Close()

	err = os.Remove(createPath)
	assert.NoError(t, err)

	t.Log("Test create file")

	initPath := "test.go"

	err = exist.InitFile(initPath)
	assert.NoError(t, err)

	err = os.Remove(initPath)
	assert.NoError(t, err)

	t.Log("Test success")
}
