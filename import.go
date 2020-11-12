package cbtest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// checkSystem returns true if the given path contains a system.
func checkSystem(path string) error {

	systemJSONPath := filepath.Join(path, "system.json")

	_, err := os.Stat(systemJSONPath)
	if err != nil {
		return fmt.Errorf("not a system: %s", err)
	}

	return nil
}

// ImportSystem the system given by merging the base system given by `systempath`
// and the extra files given by each of the `extraPaths`.
func ImportSystem(t *testing.T, systemPath string, extraPaths ...string) {

	config, err := ReadConfigFromPath(ConfigPath())
	if err != nil {
		t.Errorf("could not read config: %s", err)
		t.FailNow()
	}

	if config.HasSystem() && config.HasDeveloper() {

	} else {

	}

	err = checkSystem(systemPath)
	require.NoError(t, err)
}
