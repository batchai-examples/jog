package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/gookit/goutil/strutil"
	"github.com/qiangyt/jog/config"
	"github.com/tj/go-naturaldate"
)

// OptionsT ...
type OptionsT struct {
	LogFilePath     string
	ConfigFilePath  string
	Debug           bool
	ConfigItemPath  string
	ConfigItemValue string
	FollowMode      bool
	NumberOfLines   int

	levelFilterTexts []string
	levelFilters     []config.Enum

	beforeFilterText string
	BeforeFilter     *time.Time

	afterFilterText string
	AfterFilter     *time.Time
}

// Options ...
type Options = *OptionsT

// GetLevelFilters ...
func (i Options) GetLevelFilters() []config.Enum {
	return i.levelFilters
}

// InitLevelFilters ...
func (i Options) InitLevelFilters(levelFieldEnums config.EnumMap) {
	if len(i.levelFilterTexts) == 0 {
		i.levelFilters = make([]config.Enum, 0)
		return
	}

	for _, levelFilterText := range i.levelFilterTexts {
		levelFilter := levelFieldEnums.GetEnum(levelFilterText)
		i.levelFilters = append(i.levelFilters, levelFilter)
	}
}

// InitTimestampFilters ...
func (i Options) InitTimestampFilters(timestampField config.Field) {
	now := time.Now()

	if len(i.beforeFilterText) > 0 {
		f, err := naturaldate.Parse(i.beforeFilterText, now, naturaldate.WithDirection(naturaldate.Past))
		if err != nil {
			log.Printf("failed to parse before filter %s as natural timestamp, so try absolute parse\n", i.beforeFilterText)
			f = ParseTimestamp(timestampField, i.beforeFilterText)
		}
		log.Printf("before filter: %v", f)
		i.BeforeFilter = &f
	}
	if len(i.afterFilterText) > 0 {
		f, err := naturaldate.Parse(i.afterFilterText, now, naturaldate.WithDirection(naturaldate.Past))
		if err != nil {
			log.Printf("failed to parse after filter %s as natural timestamp, so try absolute parse\n", i.afterFilterText)
			f = ParseTimestamp(timestampField, i.afterFilterText)
		}
		log.Printf("after filter: %v", f)
		i.AfterFilter = &f

		if i.BeforeFilter != nil {
			if i.BeforeFilter.Before(*i.AfterFilter) {
				panic(fmt.Errorf("before filter (%s) shouldn't be before after filter (%s)", i.beforeFilterText, i.afterFilterText))
			}
		}
	}
}

// HasTimestampFilter ...
func (i Options) HasTimestampFilter() bool {
	return i.BeforeFilter != nil || i.AfterFilter != nil
}

// OptionsWithCommandLine ...
func OptionsWithCommandLine() (bool, Options) {

	r := &OptionsT{
		Debug:            false,
		FollowMode:       false,
		NumberOfLines:    -1,
		levelFilterTexts: make([]string, 0),
	}
	var err error
	var hasNumberOfLines = false

	for i := 0; i < len(os.Args); i++ {
		if i == 0 {
			continue
		}

		arg := os.Args[i]

		if arg[0:1] == "-" {
			if arg == "-c" || arg == "--config" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing config file path\n")
					PrintHelp()
					return false, nil
				}

				r.ConfigFilePath = os.Args[i+1]
				i++
			} else if arg == "-cs" || arg == "--config-set" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing config item expression\n")
					PrintHelp()
					return false, nil
				}

				r.ConfigItemPath, r.ConfigItemValue, err = ParseConfigExpression(os.Args[i+1])
				if err != nil {
					color.Red.Println("%v\n", err)
					PrintHelp()
					return false, nil
				}
				i++
			} else if arg == "-cg" || arg == "--config-get" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing config item path\n")
					PrintHelp()
					return false, nil
				}

				r.ConfigItemPath = os.Args[i+1]
				i++
			} else if arg == "-f" || arg == "--follow" {
				r.FollowMode = true
			} else if arg == "-n" || arg == "--lines" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing lines argument\n")
					PrintHelp()
					return false, nil
				}

				r.NumberOfLines = strutil.MustInt(os.Args[i+1])
				hasNumberOfLines = true
				i++
			} else if arg == "-t" || arg == "--template" {
				PrintConfigTemplate()
				return false, nil
			} else if arg == "-h" || arg == "--help" {
				PrintHelp()
				return false, nil
			} else if arg == "-V" || arg == "--version" {
				PrintVersion()
				return false, nil
			} else if arg == "-d" || arg == "--debug" {
				r.Debug = true
			} else if arg == "-l" || arg == "--level" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing level argument\n")
					PrintHelp()
					return false, nil
				}

				r.levelFilterTexts = append(r.levelFilterTexts, os.Args[i+1])
				i++
			} else if arg == "-a" || arg == "--after" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing after argument\n")
					PrintHelp()
					return false, nil
				}

				r.afterFilterText = os.Args[i+1]
				i++
			} else if arg == "-b" || arg == "--before" {
				if i+1 >= len(os.Args) {
					color.Red.Println("Missing before argument\n")
					PrintHelp()
					return false, nil
				}

				r.beforeFilterText = os.Args[i+1]
				i++
			} else {
				color.Red.Printf("Unknown option: '%s'\n\n", arg)
				PrintHelp()
				return false, nil
			}
		} else {
			r.LogFilePath = arg
		}
	}

	if !hasNumberOfLines {
		if r.FollowMode {
			r.NumberOfLines = 10
		}
	}

	return true, r
}
