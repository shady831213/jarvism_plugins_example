package main

import (
	"github.com/shady831213/jarvism/core"
	"github.com/shady831213/jarvism/core/errors"
	"github.com/shady831213/jarvism/core/loader"
	"github.com/shady831213/jarvism/core/plugin"
	"github.com/shady831213/jarvism/core/utils"
	"path"
)

type testSimulator struct {
}

func newTestSimulator() plugin.Plugin {
	return new(testSimulator)
}

func (s *testSimulator) Parse(cfg map[interface{}]interface{}) *errors.JVSAstError {
	return nil
}

func (s *testSimulator) KeywordsChecker(key string) (bool, *utils.StringMapSet, string) {
	return true, nil, ""
}

func (s *testSimulator) Name() string {
	return "testSimulator"
}

func (s *testSimulator) BuildInOptionFile() string {
	return path.Join(core.GetPluginsHome(), "simulators", s.Name(), "buildInOptions", "test_options.yaml")
}

func (s *testSimulator) CompileCmd() string {
	return "testCompile"
}

func (s *testSimulator) SimCmd() string {
	return "testSim"
}

func (s *testSimulator) SeedOption() string {
	return "+testSeed="
}

func (s *testSimulator) GetFileList(paths ...string) (string, error) {
	fileList := ""
	return fileList, nil
}

func init() {
	loader.RegisterSimulator(newTestSimulator)
}
