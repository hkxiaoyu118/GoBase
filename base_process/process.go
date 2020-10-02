package base_process

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"time"
)

var Timeout = 240 * time.Second //超时时间

//以超时的方式执行命令行
func CommandWithTimeout(cmdName string, arg string, stdinData string) ([]byte, error) {
	stdin := bytes.NewBuffer([]byte(stdinData))
	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctxt, cmdName, arg)
	cmd.Stdin = stdin

	if out, err := cmd.Output(); err != nil {
		//检测报错是否由于超时引起
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			return nil, errors.New("command timeout")
		}
		//报错不是由于超时引起
		return out, err
	} else {
		return out, nil
	}
}
