package main

import (
	"github.com/shady831213/jarvism/core/errors"
	"github.com/shady831213/jarvism/core/loader"
	"github.com/shady831213/jarvism/core/plugin"
	"github.com/shady831213/jarvism/core/runtime"
	"github.com/shady831213/jarvism/core/utils"
	"math/rand"
	"strings"
	"time"
)

type testRunner struct {
}

func newTestRunner() plugin.Plugin {
	return new(testRunner)
}
func (r *testRunner) Name() string {
	return "testRunner"
}

func (r *testRunner) Parse(cfg map[interface{}]interface{}) *errors.JVSAstError {
	return nil
}

func (r *testRunner) KeywordsChecker(key string) (bool, *utils.StringMapSet, string) {
	return true, nil, ""
}

func (r *testRunner) PrepareBuild(build *loader.AstBuild, cmdRunner loader.CmdRunner) *errors.JVSRuntimeResult {
	time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
	return cmdRunner(nil, "echo", " ")
}

func (r *testRunner) Build(build *loader.AstBuild, cmdRunner loader.CmdRunner) *errors.JVSRuntimeResult {
	time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
	echoError := rand.Intn(100)
	runtime.Println("simulator " + loader.GetCurSimulator().Name() + " compile_cmd:" + loader.GetCurSimulator().CompileCmd())
	runtime.Println("compile options :" + build.CompileOption())
	runtime.Println("testDiscoverer" + build.GetTestDiscoverer().Name() + " testList:" + strings.Join(build.GetTestDiscoverer().TestList(), "\n"))
	if echoError < 20 {
		return cmdRunner(nil, "echo", " Error here ", build.Name)
	} else {
		return cmdRunner(nil, "echo", " Pass here ", build.Name)
	}
}

func (r *testRunner) PrepareTest(testCase *loader.AstTestCase, cmdRunner loader.CmdRunner) *errors.JVSRuntimeResult {
	time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
	return cmdRunner(nil, "echo", "")
}

func (r *testRunner) RunTest(testCase *loader.AstTestCase, cmdRunner loader.CmdRunner) *errors.JVSRuntimeResult {
	time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
	runtime.Println("simulator " + loader.GetCurSimulator().Name() + " sim_cmd:" + loader.GetCurSimulator().SimCmd())
	runtime.Println("sim options :" + testCase.SimOption())
	return cmdRunner(nil, "echo", "UVM_WARNING @abc : ", testCase.Name)
}

func init() {
	loader.RegisterRunner(newTestRunner)
}
