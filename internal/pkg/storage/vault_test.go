package storage

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"vault-cluster-replication/internal/pkg/storage/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSystem_PullSnapshot_OK(t *testing.T) {
	mockSys := mocks.NewSys(t)
	system := NewClient(nil, mockSys, nil)
	sys := NewSystem(system.Sys)
	var snapshot bytes.Buffer
	var writer io.Writer
	mockSys.On("RaftSnapshot", &snapshot).Once().Return(writer, nil)
	_, err := sys.PullSnapshot()
	assert.Equal(t, nil, err)
	mockSys.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSys)
}

func TestSystem_PullSnapshot_Err(t *testing.T) {
	mockSys := mocks.NewSys(t)
	system := NewClient(nil, mockSys, nil)
	sys := NewSystem(system.Sys)
	var snapshot bytes.Buffer
	mockSys.On("RaftSnapshot", &snapshot).Once().Return(fmt.Errorf("error"))
	_, err := sys.PullSnapshot()
	assert.Error(t, err)
	mockSys.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSys)
}

func TestSystem_PushSnapshot_OK(t *testing.T) {
	mockSys := mocks.NewSys(t)
	system := NewClient(nil, mockSys, nil)
	sys := NewSystem(system.Sys)
	// Create a test file
	err := os.WriteFile("test.txt", []byte("test"), 0o600)
	if err != nil {
		return
	}

	mockSys.On("RaftSnapshotRestore", mock.Anything, true).Once().Return(nil)
	err = sys.PushSnapshot("test.txt")
	assert.Equal(t, nil, err)
	mockSys.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSys)
}

func TestSystem_PushSnapshot_Err(t *testing.T) {
	mockSys := mocks.NewSys(t)
	system := NewClient(nil, mockSys, nil)
	sys := NewSystem(system.Sys)
	// Create a test file
	err := os.WriteFile("test.txt", []byte("test"), 0o600)
	if err != nil {
		return
	}

	mockSys.On("RaftSnapshotRestore", mock.Anything, true).Once().Return(fmt.Errorf("error"))
	err = sys.PushSnapshot("test.txt")
	assert.Error(t, err)
	mockSys.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSys)
}

func TestWriter(t *testing.T) {
	// Define test data
	backupFile := "test_backup.txt"
	data := []byte("Test data to be written to the backup file.")

	// Ensure the test backup file is removed at the end of the test
	defer func() {
		if err := os.Remove(backupFile); err != nil {
			t.Errorf("Error cleaning up test backup file: %v", err)
		}
	}()

	// Call the writerFile function
	err := writerFile(backupFile, data)
	if err != nil {
		t.Errorf("writerFile function returned an error: %v", err)
	}

	// Check if the file was written correctly
	contents, readErr := os.ReadFile(backupFile)
	if readErr != nil {
		t.Errorf("Error reading the test backup file: %v", readErr)
	}

	if string(contents) != string(data) {
		t.Errorf("Written data does not match the expected data")
	}
}

func TestWriterError(t *testing.T) {
	// Define a test case with an invalid file path
	backupFile := "" // An empty path is invalid
	data := []byte("Test data to be written to the backup file.")

	// Call the writerFile function with an invalid file path
	err := writerFile(backupFile, data)
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestOpenFile_Success(t *testing.T) {
	// Create a temporary test file
	backupFileName := "test_backup.txt"
	content := []byte("Test data in the file")
	err := os.WriteFile(backupFileName, content, 0600)
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			t.Errorf("Error cleaning up test backup file: %v", err)
		}
	}(backupFileName)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	// Call the openFile function
	file, err := openFile(backupFileName)

	// Check if the file was opened successfully
	if err != nil {
		t.Errorf("openFile failed to open the file: %v", err)
	}
	if file == nil {
		t.Error("openFile returned a nil file")
	}
	// Clean up the file
	err = file.Close()
	if err != nil {
		return
	}
}

func TestOpenFile_Error(t *testing.T) {
	// Try to open a non-existent file
	backupFileName := "non_existent_file.txt"

	// Call the openFile function
	file, err := openFile(backupFileName)

	// Check if the function returned an error and a nil file
	if err == nil {
		t.Error("openFile should have returned an error, but it didn't")
	}
	if file != nil {
		t.Error("openFile should have returned a nil file, but it didn't")
	}
}
