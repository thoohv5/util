package command

import (
	"bufio"
	"context"
	"io"
	"strings"
	"testing"
)

func TestExec_reader(t *testing.T) {
	out, err := Exec(context.Background(), "bash", WithArgs("-c", "for i in 1 2 3 4;do echo $i'$';sleep 2;done"), WithWriter(func(reader io.Reader) {

		buf := bufio.NewReader(reader)

		for {
			data, err := buf.ReadBytes('$')
			if nil != err {
				if err == io.EOF {
					err = nil
				}
				return
			}
			t.Log(strings.TrimSpace(string(data)))
		}

	}))
	t.Log(out.String(), err)
}

func TestExec_out(t *testing.T) {
	out, err := Exec(context.Background(), "bash", WithArgs("-c", "for i in 1 2 3 4;do echo $i'$';sleep 2;done"))
	t.Log(out.String(), err)
}

func TestExec_out_bash(t *testing.T) {
	out, err := ExecBash(context.Background(), WithArgs("for i in 1 2 3 4;do echo $i'$';sleep 2;done"))
	t.Log(out.String(), err)
}
