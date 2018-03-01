package main

import (
	"bytes"
	"flag"
	"log"
	"os/exec"
	"path/filepath"
)

// global structure to hold the command line parameters
type clParametersType struct {
	src string
	dst string
}

func main() {
	clParams := getParamsFromCommandLine(nil)
	executeRoboCopyWithParams(clParams)
}

func getParamsFromCommandLine(clParams *clParametersType) clParametersType {
	// if the params are passed in, use them instead of the reading from command line

	src := flag.String("src", "", "source directory to be copy from")
	dst := flag.String("dst", "", "destination directory to copy to")

	if clParams != nil {
		src = &(clParams.src)
		dst = &(clParams.dst)
	} else {
		flag.Parse()
	}

	// TODO: ADD VALIDATION LOGIC TO SRC AND DST PATHS
	convertToUncPath(*src)

	return clParametersType{*src, *dst}
}

func convertToUncPath(taskDsPath string) (uncPath string) {
	// uncPath = path.Dir(taskDsPath)
	// url, err := url.Parse(taskDsPath)
	// if err != nil {
	// 	defer errorHandler(err)
	// }
	uncPath = taskDsPath // url.Host + url.Path
	uncPath = filepath.Dir(uncPath)
	println(uncPath)
	return uncPath
}

func executeRoboCopyWithParams(clParams clParametersType) {
	//params source: http://improve.dk/simple-file-synchronization-using-robocopy/
	cmd := exec.Command("robocopy", clParams.src, clParams.dst, "/MIR", "/R:5", "/W:10", "/Z", "/XA:H", "/FFT", "/LOG:test.log")
	var out, errB bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errB

	err := cmd.Start()
	if err != nil {
		defer errorHandler(err)
	}

	err = cmd.Wait()
	if err != nil {
		defer errorHandler(err)
	}
}

func errorHandler(err error) {
	log.Fatalf("%s", err.Error())
}
