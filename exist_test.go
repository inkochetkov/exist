package exist

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDir check dir
func TestDir(t *testing.T) {

	t.Log("Test check dir")
	checkPath := "/test"

	check := CheckDir(checkPath)

	assert.Equal(t, false, check)

	t.Log("Test create dir")
	createPath := "/test"

	err := CreateDir(createPath)
	assert.NoError(t, err)

	err = os.Remove(createPath)
	assert.NoError(t, err)

	t.Log("Test init dir")
	initPath := "/test"

	err = InitDir(initPath)
	assert.NoError(t, err)

	err = os.Remove(initPath)
	assert.NoError(t, err)

	t.Log("Test success")

}

// TestFile check file
func TestFile(t *testing.T) {
	t.Log("Test check file")

	checkPath := "exist_test.go"
	check := CheckFile(checkPath)

	assert.Equal(t, true, check)

	t.Log("Test create file")

	createPath := "test.go"

	file, err := CreateFile(createPath)
	assert.NoError(t, err)

	file.Close()

	err = os.Remove(createPath)
	assert.NoError(t, err)

	t.Log("Test create file")

	initPath := "test.go"

	err = InitFile(initPath)
	assert.NoError(t, err)

	err = os.Remove(initPath)
	assert.NoError(t, err)

	t.Log("Test success")
}
