package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
	"github.com/thomasboyt/go-selecta/selecta"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	app := cli.NewApp()
	app.Name = "selecta"
	app.Usage = "fuzzy find whatever you want"
	app.Flags = []cli.Flag {
		cli.StringFlag{"search, s", "", "initial search to fill in"},
	}
	app.Action = func(c *cli.Context) {
		// parse choices
		bytes, _ := ioutil.ReadAll(os.Stdin)
		choices := strings.Split(string(bytes), "\n")

		// create a search
		initialSearch := c.String("search")

		// set up termbox
		err := termbox.Init()
		if err != nil {
			panic(err)
		}

		_, height := termbox.Size()
		search := selecta.BlankSearch(choices, initialSearch, height - 1)

		termbox.SetInputMode(termbox.InputEsc)

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		DrawApp(search)
		EventLoop(search)

		termbox.Close()

		fmt.Fprint(os.Stdout, search.SelectedChoice())
	}

	app.Run(os.Args)
}

func EventLoop(s *selecta.Search) {
	for !s.Done {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				termbox.Close()
				os.Exit(0)
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				s.Backspace()
			case termbox.KeyEnter:
				s.Done = true
			case termbox.KeyCtrlN:
				s.Down()
			case termbox.KeyCtrlP:
				s.Up()
			case termbox.KeyCtrlW:
				s.DeleteWord()
			default:
				char := rune(ev.Ch)
				if !unicode.IsControl(char) {
					s.AppendQuery(string(char))
				} else if ev.Key == termbox.KeySpace {
					s.AppendQuery(" ")
				}
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			DrawApp(s)
		}
	}
}

// Rendering

func DrawApp(s *selecta.Search) {
	WriteLine(0, "> "+s.Query, false)
	termbox.SetCursor(2+len(s.Query), 0)

	for i, match := range s.Matches {
		if i >= s.VisibleChoices {
			break
		}
		choice := match.Value
		highlight := false
		if s.Index == i {
			highlight = true
		}
		WriteLine(i+1, choice, highlight)
	}

	termbox.Flush()
}

func WriteLine(row int, str string, highlight bool) {
	bgColor := termbox.ColorDefault
	width, _ := termbox.Size()
	if highlight {
		bgColor = termbox.ColorBlack
	}

	for col := 0; col < width; col++ {
		if col < len(str) {
			termbox.SetCell(col, row, rune(str[col]), termbox.ColorDefault, bgColor)
		} else {
			termbox.SetCell(col, row, ' ', termbox.ColorDefault, bgColor)
		}
	}
}
