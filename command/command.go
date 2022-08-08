package command

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
)

const (
	Bash    = "bash"
	BashOpt = "-c"
)

// Exec 执行命令
func Exec(ctx context.Context, name string, opts ...Option) (out bytes.Buffer, err error) {

	// 初始化参数
	o := options{
		env: os.Environ(),
	}
	for _, opt := range opts {
		opt(&o)
	}

	// 输入输出流
	stderr := bytes.NewBuffer(nil)
	in := bytes.NewBuffer(nil)
	exit := make(chan struct{})

	// 拼接命令
	cmd := exec.CommandContext(ctx, name, o.args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// 命令流
	cmd.Stdin = in
	// cmd.Stdout = out
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		return
	}
	cmd.Stderr = stderr

	// 输入流命令
	if len(o.inArgs) > 0 {
		// 开启协程
		go func() {
			// 防止golang崩溃
			defer func() {
				if rec := recover(); nil != rec {
					if nil != err {
						err = errors.WithMessagef(err, "rec:%v", rec)
					} else {
						err = errors.Errorf("rec:%v", rec)
					}
				}
			}()
			for _, arg := range o.inArgs {
				in.WriteString(fmt.Sprintf("%s\n", arg))
			}
			exit <- struct{}{}
		}()
		<-exit
	}

	// 执行
	if err = cmd.Start(); nil != err {
		err = errors.WithMessage(err, stderr.String())
		return
	}

	// 并发获取执行结果
	if reader := o.reader; reader != nil {
		go func() {
			// 防止golang崩溃
			defer func() {
				if rec := recover(); nil != rec {
					if nil != err {
						err = errors.WithMessagef(err, "rec:%v", rec)
					} else {
						err = errors.Errorf("rec:%v", rec)
					}
				}
			}()
			reader(stdout)
		}()
	} else {
		go func() {
			_, _ = io.Copy(&out, stdout)
		}()
	}

	// 结束
	if err = cmd.Wait(); nil != err {
		err = errors.WithMessage(err, stderr.String())
		return
	}

	// Kill进程及其子进程，以及处在sleep或者阻塞状态中的程序
	if err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); nil != err {
		return
	}

	return
}

func ExecBash(ctx context.Context, opts ...Option) (out bytes.Buffer, err error) {
	return Exec(ctx, Bash, append([]Option{WithArgs(BashOpt)}, opts...)...)
}
