package cli

import (
	"fmt"
	"os"
	"strings"

	"refr/internal/content"
	"refr/internal/render"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search reference pages by title, tags, or content",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runSearch,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func runSearch(cmd *cobra.Command, args []string) error {
	tree, err := buildTree(appConfig)
	if err != nil {
		return err
	}

	query := strings.Join(args, " ")
	results := content.Search(tree, query)
	styles := render.NewStyles(appConfig.Theme)
	highlighter := render.NewHighlighter(appConfig.Theme.Syntax)

	if len(results) == 0 {
		fmt.Printf("No results for %q\n", query)
		return nil
	}

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		for i, r := range results {
			fmt.Printf("%d  %s\n", i+1, r.Name)
		}
		return nil
	}
	defer term.Restore(fd, oldState)

	buf := make([]byte, 1)

	for {
		fmt.Print("\033[H\033[2J")

		header := styles.Header.Render(fmt.Sprintf("Search: %q (%d results)", query, len(results)))
		fmt.Print(header)
		fmt.Print("\r\n\r\n")

		numWidth := len(fmt.Sprintf("%d", len(results)))
		maxName := 0
		for _, r := range results {
			if len(r.Name) > maxName {
				maxName = len(r.Name)
			}
		}

		for i, r := range results {
			num := fmt.Sprintf("%*d", numWidth, i+1)
			name := fmt.Sprintf("%-*s", maxName, r.Name)
			fmt.Printf("  %s  %s", styles.Number.Render(num), styles.Page.Render(name))
			if r.Description != "" {
				fmt.Printf("    %s", styles.Description.Render(r.Description))
			}
			fmt.Print("\r\n")
		}

		fmt.Print("\r\n")
		help := styles.HelpBar.Render("[number] view") +
			styles.Accent.Render(" · ") +
			styles.HelpBar.Render("[q] quit")
		fmt.Print(help)
		fmt.Print("\r\n\r\n> ")

		_, err := os.Stdin.Read(buf)
		if err != nil {
			return nil
		}

		ch := buf[0]

		if ch == 'q' || ch == 'Q' {
			return nil
		}

		if ch >= '1' && ch <= '9' {
			idx := int(ch-'0') - 1
			if idx < len(results) {
				r := results[idx]
				if r.Page != nil {
					fmt.Print("\033[H\033[2J")
					output := render.RenderPage(r.Page, r.Name, styles, highlighter)
					if appConfig.Pager != "" {
						term.Restore(fd, oldState)
						render.PipeToPage(output, appConfig.Pager)
						term.MakeRaw(fd)
					} else {
						rawPrint(output)
						fmt.Print("\r\n\r\nPress any key to go back...")
						os.Stdin.Read(buf)
					}
				}
			}
		}
	}
}
