/*
uvm_test test discoverers

tests will be automatically discovered if testcases organized as follow:

. $TestDir(default is $JVS_PRJ_HOME/testcases)

|--- test1(test1 is valid, and pass to simulator through +C_TEST=test1)

|------test1.c

|--- test2(test2 is invalid)

|------test2.c

|--- test3(test3 is invalid)
*/

package main

import (
	"github.com/shady831213/jarvism/core"
	"github.com/shady831213/jarvism/core/errors"
	"github.com/shady831213/jarvism/core/loader"
	"github.com/shady831213/jarvism/core/plugin"
	"github.com/shady831213/jarvism/core/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type cDiscoverer struct {
	testDir string
	tests   map[string]interface{}
}

func (d *cDiscoverer) Parse(cfg map[interface{}]interface{}) *errors.JVSAstError {
	//AstParse test_dir
	if err := loader.CfgToAstItemOptional(cfg, "test_dir", func(item interface{}) *errors.JVSAstError {
		testDir, err := filepath.Abs(os.ExpandEnv(item.(string)))
		if err != nil {
			return errors.JVSAstParseError(d.Name(), err.Error())
		}
		//check path
		if _, err := os.Stat(testDir); err != nil {
			return errors.JVSAstParseError(d.Name(), err.Error())
		}
		d.testDir = testDir
		return nil
	}); err != nil {
		return errors.JVSAstParseError("test_dir of "+d.Name(), err.Error())
	}
	//use default
	if d.testDir == "" {
		d.testDir, _ = filepath.Abs(path.Join(core.GetPrjHome(), "testcases"))
	}
	return nil
}

func (d *cDiscoverer) KeywordsChecker(s string) (bool, *utils.StringMapSet, string) {
	keywords := utils.NewStringMapSet()
	keywords.AddKey("test_dir")
	if !loader.CheckKeyWord(s, keywords) {
		return false, keywords, "Error in " + d.Name() + ":"
	}
	return true, nil, ""
}

func (d *cDiscoverer) Name() string {
	return "c_test"
}

func (d *cDiscoverer) TestDir() string {
	return d.testDir
}

func (d *cDiscoverer) TestCmd() string {
	return "+C_TEST="
}

func (d *cDiscoverer) TestList() []string {
	if d.discoverTest() != nil {
		return []string{}
	}
	return utils.KeyOfStringMap(d.tests)
}

func (d *cDiscoverer) TestFileList() []string {
	if d.discoverTest() != nil {
		return []string{}
	}
	values := utils.ValueOfStringMap(d.tests)
	paths := make([]string, len(values))
	for i, v := range values {
		paths[i] = v.(string)
	}
	return paths
}

func (d *cDiscoverer) discoverTest() error {
	if d.tests == nil {
		d.tests = make(map[string]interface{})
		err := filepath.Walk(d.testDir, d.filter)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *cDiscoverer) filter(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	if filepath.Ext(path) != ".c" {
		return nil
	}
	fileName := strings.TrimSuffix(filepath.Base(path), ".c")
	if fileName == filepath.Base(filepath.Dir(path)) {
		d.tests[fileName] = path
	}
	return nil
}

func (d *cDiscoverer) IsValidTest(test string) bool {
	if d.discoverTest() != nil {
		return false
	}
	_, ok := d.tests[test]
	return ok
}

func newCDiscoverer() plugin.Plugin {
	inst := new(cDiscoverer)
	return inst
}

func init() {
	loader.RegisterTestDiscoverer(newCDiscoverer)
}
