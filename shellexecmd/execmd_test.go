package shellexecmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDag_ScriptRunner(t *testing.T) {

	s := "./test01.sh"

	cmd, r := ScriptRunner(s)
	StartThenWait(cmd)
	/*go func(cmd *exec.Cmd) {
		if cmd != nil {
			if err := cmd.Start(); err != nil {
				log.Printf("Error starting Cmd: %v", err)
				return
			}
			if err := cmd.Wait(); err != nil {
				log.Printf("Error waiting for Cmd: %v", err)
				return
			}
		}
	}(cmd)*/

	ch := Reply(r)
	// 습관의 차이 일것 같지만 비교해보는 것도 좋은 주제가 되지 않을까?
	// https://stackoverflow.com/questions/37599302/string-contains-vs-string-equals-or-string-performance
	for m := range ch {
		if strings.Contains(m, "FINISHED") {
			log.Println("Exit Ok")
			os.Exit(0)
		}
		if strings.Contains(m, "ERRORS") {
			log.Println("Exit Error")
			os.Exit(1)
		}

		fmt.Println(">", m)
	}

}

func TestScriptRunnerString(t *testing.T) {
	//s := "./test01.sh"

	script := `
	set -e
	sleep 1
	echo "Hello World"
	sleep 1
	echo "one"
	sleep 1
	echo "two"
	sleep 1
	echo "three"
	sleep 1
	echo "four"
	sleep 1
	echo "Sleep 10s"
	sleep 10s
	echo "End"`

	cmd, r := ScriptRunnerString(script)
	StartThenWait(cmd)

	ch := Reply(r)
	// 습관의 차이 일것 같지만 비교해보는 것도 좋은 주제가 되지 않을까?
	// https://stackoverflow.com/questions/37599302/string-contains-vs-string-equals-or-string-performance
	for m := range ch {
		if strings.Contains(m, "FINISHED") {
			log.Println("Exit Ok")
			os.Exit(0)
		}
		if strings.Contains(m, "ERRORS") {
			log.Println("Exit Error")
			os.Exit(1)
		}

		fmt.Println(">", m)
	}
}

func TestRunner(t *testing.T) {
	script := `
	set -e
	sleep 1
	echo "Hello World"
	sleep 1
	echo "one"
	sleep 1
	echo "two"
	sleep 1
	echo "three"
	sleep 1
	echo "four"
	sleep 1
	echo "Sleep 10s"
	sleep 10s
	echo "End"`

	Runner(script)
}
