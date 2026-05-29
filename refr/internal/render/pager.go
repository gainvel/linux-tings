package render

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func PipeToPage(content string, pager string) error {
	if pager == "" {
		_, err := fmt.Fprint(os.Stdout, content)
		return err
	}

	parts := strings.Fields(pager)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, _ = io.WriteString(stdin, content)
	stdin.Close()

	return cmd.Wait()
}
