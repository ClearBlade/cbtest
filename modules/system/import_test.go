package system

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest/config"

	mocks "github.com/clearblade/cbtest/mocks"
)

func makeMockT() *mocks.T {
	mockT := mocks.T{}
	mockT.On("Errorf", mock.Anything, mock.Anything).Return()
	mockT.On("FailNow").Return()
	mockT.On("Helper").Return()
	mockT.On("Log", mock.Anything).Return()
	mockT.On("Logf", mock.Anything, mock.Anything).Return()
	return &mockT

}

func TestRunImportWorkflow_ValidStepsSucceeds(t *testing.T) {

	config := config.GetDefaultConfig()
	mockT := makeMockT()
	mockSteps := &mockImportSteps{}
	mockSteps.On("Path").Return("some-path")
	mockSteps.On("Extra").Return([]string{"some-extra-path"})
	mockSteps.On("CheckSystem", "some-path").Return(nil)
	mockSteps.On("MakeTempDir").Return("some-temp-dir", func() {})
	mockSteps.On("MergeFolders", "some-temp-dir", "some-path", "some-extra-path").Return(nil)
	mockSteps.On("RegisterDeveloper", mockT, config).Return(nil)
	mockSteps.On("DoImport", mockT, config, "some-temp-dir").Return("some-system-key", "some-system-secret", nil)

	system, err := runImportWorkflow(mockT, config, mockSteps)

	require.NoError(t, err)
	assert.Equal(t, "some-system-key", system.SystemKey())
	assert.Equal(t, "some-system-secret", system.SystemSecret())
	assert.Equal(t, "some-temp-dir", system.LocalPath())
	assert.False(t, system.IsExternal())
	mockSteps.AssertExpectations(t)
}

func TestRunImportWorkflow_BadCheckSystemFails(t *testing.T) {

	config := config.GetDefaultConfig()
	mockT := makeMockT()
	mockSteps := &mockImportSteps{}
	mockSteps.On("Path").Return("some-path")
	mockSteps.On("Extra").Return([]string{"some-extra-path"})
	mockSteps.On("CheckSystem", "some-path").Return(fmt.Errorf("some-check-system-error"))

	system, err := runImportWorkflow(mockT, config, mockSteps)

	assert.Nil(t, system)
	assert.EqualError(t, err, "some-check-system-error")
	mockSteps.AssertExpectations(t)
}

func TestRunImportWorkflow_BadMergeFoldersFails(t *testing.T) {

	config := config.GetDefaultConfig()
	mockT := makeMockT()
	mockSteps := &mockImportSteps{}
	mockSteps.On("Path").Return("some-path")
	mockSteps.On("Extra").Return([]string{"some-extra-path"})
	mockSteps.On("CheckSystem", "some-path").Return(nil)
	mockSteps.On("MakeTempDir").Return("some-temp-dir", func() {})
	mockSteps.On("MergeFolders", "some-temp-dir", "some-path", "some-extra-path").Return(fmt.Errorf("some-merge-folders-error"))
	mockSteps.On("Cleanup", "some-temp-dir").Return()

	system, err := runImportWorkflow(mockT, config, mockSteps)

	assert.Nil(t, system)
	assert.EqualError(t, err, "some-merge-folders-error")
	mockSteps.AssertExpectations(t)
}

func TestRunImportWorkflow_BadRegisterDeveloperFails(t *testing.T) {

	config := config.GetDefaultConfig()
	mockT := makeMockT()
	mockSteps := &mockImportSteps{}
	mockSteps.On("Path").Return("some-path")
	mockSteps.On("Extra").Return([]string{"some-extra-path"})
	mockSteps.On("CheckSystem", "some-path").Return(nil)
	mockSteps.On("MakeTempDir").Return("some-temp-dir", func() {})
	mockSteps.On("MergeFolders", "some-temp-dir", "some-path", "some-extra-path").Return(nil)
	mockSteps.On("RegisterDeveloper", mockT, config).Return(fmt.Errorf("some-register-developer-error"))
	mockSteps.On("Cleanup", "some-temp-dir").Return()

	system, err := runImportWorkflow(mockT, config, mockSteps)

	assert.Nil(t, system)
	assert.EqualError(t, err, "some-register-developer-error")
	mockSteps.AssertExpectations(t)
}

func TestRunImportWorkflow_BadDoImportFails(t *testing.T) {

	config := config.GetDefaultConfig()
	mockT := makeMockT()
	mockSteps := &mockImportSteps{}
	mockSteps.On("Path").Return("some-path")
	mockSteps.On("Extra").Return([]string{"some-extra-path"})
	mockSteps.On("CheckSystem", "some-path").Return(nil)
	mockSteps.On("MakeTempDir").Return("some-temp-dir", func() {})
	mockSteps.On("MergeFolders", "some-temp-dir", "some-path", "some-extra-path").Return(nil)
	mockSteps.On("RegisterDeveloper", mockT, config).Return(nil)
	mockSteps.On("DoImport", mockT, config, "some-temp-dir").Return("", "", fmt.Errorf("some-do-import-error"))
	mockSteps.On("Cleanup", "some-temp-dir").Return()

	system, err := runImportWorkflow(mockT, config, mockSteps)

	assert.Nil(t, system)
	assert.EqualError(t, err, "some-do-import-error")
	mockSteps.AssertExpectations(t)
}
