package shellexecmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

// TODO 코드 정리 필요.
// TODO exec.Command 에서 파일 위치에 대한 검사를 하지만 오류를 리턴하지는 않는다.
// 하지만, 이 함수는 사용되지 않으므로 구현은 생략한다.
// 따라서 이부분은 따로 처리 해줘야 한다.
func ScriptRunner(s string) (*exec.Cmd, io.Reader) {
	cmd := exec.Command(s)

	// StdoutPipe 쓰면 Run 및 기타 Run 을 포함한 method 를 쓰면 에러난다.
	r, err := cmd.StdoutPipe()
	if err != nil {
		log.Panicf("Error stdout pipe for Cmd: %v", err)
	}

	return cmd, r
}

// 이름은 대충
// TODO
// https://yourbasic.org/golang/multiline-string/
// string으로 가지고 와서 shell command 만 오는 것이 아니라, \t\n 같은 escape letter 들도 온다. 물론 실행에는 문제 없지만
// 잠재적인 오류를 없애기 위해서 shell command 만 가지고 와야 하지 않을까??
func ScriptRunnerString(s string) (*exec.Cmd, io.Reader) {
	cmd := exec.Command("/bin/sh", "-c", s)

	// StdoutPipe 쓰면 Run 및 기타 Run 을 포함한 method 를 쓰면 에러난다.
	r, err := cmd.StdoutPipe()
	if err != nil {
		log.Panicf("Error stdout pipe for Cmd: %v", err)
	}

	return cmd, r
}

func StartThenWait(cmd *exec.Cmd) {
	go func(cmd *exec.Cmd) {
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
	}(cmd)
}

// 스크립트 실행을 기다리지 않고 실시간으로 결과를 출력하기 위해서 고루틴을 사용하고 있다.
func Reply(i io.Reader) <-chan string {
	r := make(chan string, 1)

	go func() {
		// TODO 여기서 close 할 경우 상관 없지 않을까???
		defer close(r)
		scan := bufio.NewScanner(i)

		for {
			b := scan.Scan()
			if b != true {
				if scan.Err() == nil {
					// grpc 에서는 스트림을 닫아버리자.
					r <- "FINISHED"
					break
				}
				// 그외 에러 표시하기.
				log.Println(scan.Err())
				r <- "ERRORS"
				break
			}

			s := scan.Text()
			r <- s
		}
	}()

	return r
}

// 일단 빠르게 작성하고 코드 정리는 나중에
func Runner(s string) bool {

	if len(strings.TrimSpace(s)) == 0 {
		return false
	}
	cmd, r := ScriptRunnerString(s)
	StartThenWait(cmd) // goroutine

	ch := Reply(r)

	PrintOutput(ch)

	return true
}

// TODO 리턴값을 만들어야 함.
// 로그 관련해서 정리하기
// TODO Reply 에서 close 시켜 주고 있음. 상관없을듯한데...
func PrintOutput(ch <-chan string) {

	for m := range ch {
		if strings.Contains(m, "FINISHED") {
			log.Println("Exit Ok")
			return
		}
		if strings.Contains(m, "ERRORS") {
			log.Println("Exit Error")
			return
		}
		fmt.Println(">", m)
	}
}
