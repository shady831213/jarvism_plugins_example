package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestDiscovererPlugin(t *testing.T) {
	cmd := exec.Command("jarvism", "show_plugins", "testDiscoverer")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	cmd = exec.Command("jarvism", "show_tests", "build1")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestRun(t *testing.T) {
	cmd := exec.Command("jarvism", "run_test", "build1", "test1")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func init() {
	//build and load plugin
	prjHome, _ := filepath.Abs("testFiles")
	pluginHome := filepath.Dir(filepath.Dir(filepath.Dir(prjHome)))
	os.Setenv("JVS_PRJ_HOME", prjHome)
	os.Setenv("JVS_PLUGINS_HOME", pluginHome)
}
