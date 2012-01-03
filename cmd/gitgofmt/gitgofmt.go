//
// gitgofmt checks that any go files in the git index are correctly formatted.
// If not, it writes the bad files to stderr and returns an error status of 1.
//
package main

import (
	"fmt"
	"mrl/utils"
	"os"
	"bytes"
	"exec"
)

func main() {
	// Get the list of files to be committed
	stdout, _, err := utils.Execute("git", "diff", "--cached", "--name-only", "--diff-filter=ACM")
	if err != nil {
		fmt.Fprintf(os.Stderr, "git diff failed with %s\n", err)
		os.Exit(1)
	}
	filelist := utils.SplitLines(stdout)
	goext := []byte(".go")
	ok := true
	for _, file := range filelist {
		if bytes.HasSuffix(file, goext) {
			// Gross way to do this!
			catcmd := exec.Command("git", "show", fmt.Sprintf(":%s", file))
			catout, err := catcmd.StdoutPipe()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in cat %s: %s\n", file, err)
				continue
			}
			fmtcmd := exec.Command("gofmt", "-l")
			fmtcmd.Stdin = catout
			outb := utils.NewByteWriter()
			errb := utils.NewByteWriter()
			fmtcmd.Stdout = outb
			fmtcmd.Stderr = errb
			catcmd.Start()
			fmterr := fmtcmd.Run()
			catcmd.Wait()
			if fmterr != nil {
				ok = false
				fmt.Fprintf(os.Stderr, "syntax error in %s\n", file)
			}
			if len(outb.Data()) != 0 {
				ok = false
				fmt.Fprintf(os.Stderr, "formatting error: run gofmt -w %s\n", file)
			}
		}
	}
	if !ok {
		os.Exit(1)
	}
}
