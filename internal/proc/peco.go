package proc

import (
	"bytes"
	"os/exec"
)

const (
	CMD_PECO = "peco"
)

// pecoが使えるかを調べ、使えなければerrorを返します
func CheckCanUsePeco() error {
	_, err := exec.LookPath(CMD_PECO)
	if err != nil {
		return err
	}

	return nil
}

// pecoをbufを入力に実行し、その結果を返します
func ExecPeco(buf []byte) string {
	if err := CheckCanUsePeco(); err != nil {
		return ""
	}

	cmd := exec.Command(CMD_PECO)
	cmd.Stdin = bytes.NewBuffer(buf)

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(out)
}
