package cli

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"refr/internal/config"
	"refr/internal/content"
	"refr/internal/nav"
	"refr/internal/render"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func rawPrint(s string) {
	fmt.Print(strings.ReplaceAll(s, "\n", "\r\n"))
}

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse reference pages interactively",
	RunE:  runBrowse,
}

func init() {
	rootCmd.AddCommand(browseCmd)
	rootCmd.RunE = browseCmd.RunE
}

func runBrowse(cmd *cobra.Command, args []string) error {
	tree, err := buildTree(appConfig)
	if err != nil {
		return err
	}

	navigator := nav.New(tree)
	styles := render.NewStyles(appConfig.Theme)
	highlighter := render.NewHighlighter(appConfig.Theme.Syntax)

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("entering raw mode: %w", err)
	}
	defer term.Restore(fd, oldState)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		term.Restore(fd, oldState)
		os.Exit(0)
	}()
	defer signal.Stop(sigCh)

	buf := make([]byte, 1)
	var pendingDigit int
	hasPending := false

	for {
		cur := navigator.Current()

		fmt.Print("\033[H\033[2J")

		if !cur.IsCategory && cur.Page != nil {
			output := render.RenderPage(cur.Page, navigator.Breadcrumb(), styles, highlighter)
			if appConfig.Pager != "" {
				term.Restore(fd, oldState)
				render.PipeToPage(output, appConfig.Pager)
				term.MakeRaw(fd)
			} else {
				rawPrint(output)
				fmt.Print("\r\n\r\nPress any key to go back...")
				os.Stdin.Read(buf)
			}
			navigator.Back()
			continue
		}

		output := render.RenderCategory(cur, navigator.Breadcrumb(), appConfig.ShowNumbers, styles)
		rawPrint(output)
		fmt.Print("\r\n\r\n> ")

		_, err := os.Stdin.Read(buf)
		if err != nil {
			return nil
		}

		ch := buf[0]

		if ch == 'q' || ch == 'Q' || ch == 3 {
			return nil
		}

		if ch == 'b' || ch == 'B' {
			if !navigator.Back() {
				return nil
			}
			hasPending = false
			continue
		}

		if ch == 27 {
			hasPending = false
			continue
		}

		if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			if hasPending {
				selection := pendingDigit*10 + digit - 1
				hasPending = false
				selectChild(navigator, cur, selection, styles, highlighter, appConfig, fd, oldState)
				continue
			}

			if len(cur.Children) > 9 {
				pendingDigit = digit
				hasPending = true
				fmt.Printf("%c", ch)

				os.Stdin.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
				n, _ := os.Stdin.Read(buf)
				os.Stdin.SetReadDeadline(time.Time{})

				if n > 0 && buf[0] >= '0' && buf[0] <= '9' {
					selection := pendingDigit*10 + int(buf[0]-'0') - 1
					hasPending = false
					selectChild(navigator, cur, selection, styles, highlighter, appConfig, fd, oldState)
					continue
				}
				selection := digit - 1
				hasPending = false
				if n > 0 {
					selectChild(navigator, cur, selection, styles, highlighter, appConfig, fd, oldState)
					continue
				}
				selectChild(navigator, cur, selection, styles, highlighter, appConfig, fd, oldState)
				continue
			}

			selection := digit - 1
			selectChild(navigator, cur, selection, styles, highlighter, appConfig, fd, oldState)
			continue
		}
	}
}

func selectChild(navigator *nav.Navigator, cur *content.Node, index int, styles *render.Styles, highlighter *render.Highlighter, cfg *config.Config, fd int, oldState *term.State) {
	if index < 0 || index >= len(cur.Children) {
		return
	}

	child := cur.Children[index]
	if child.IsCategory {
		navigator.Select(index)
		return
	}

	if child.Page == nil {
		return
	}

	fmt.Print("\033[H\033[2J")
	bc := navigator.BreadcrumbFor(child)
	output := render.RenderPage(child.Page, bc, styles, highlighter)

	if cfg.Pager != "" {
		term.Restore(fd, oldState)
		render.PipeToPage(output, cfg.Pager)
		term.MakeRaw(fd)
	} else {
		rawPrint(output)
		fmt.Print("\r\n\r\nPress any key to go back...")
		buf := make([]byte, 1)
		os.Stdin.Read(buf)
	}
}
