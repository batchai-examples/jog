package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintVersion(t *testing.T) {
	// Arrange
	expectedOutput := static.AppVersion + "\n"

	// Act
	actualOutput := captureStdout(func() {
		PrintVersion()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintConfigTemplate(t *testing.T) {
	// Arrange
	expectedOutput := config.BuildDefaultConfigurationYAML() + "\n"

	// Act
	actualOutput := captureStdout(func() {
		PrintConfigTemplate()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelp(t *testing.T) {
	// Arrange
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in [defaultGrokLibraryDirs]
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to [defaultGrokLibraryDirs]
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithCustomGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs := "/custom/grok/library/dirs"
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in [/custom/grok/library/dirs]
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to [/custom/grok/library/dirs]
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		defaultGrokLibraryDirs = "/custom/grok/library/dirs"
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithEmptyGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs := ""
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in []
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to []
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		defaultGrokLibraryDirs = ""
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithInvalidGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs := "/invalid/grok/library/dirs"
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in []
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to []
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		defaultGrokLibraryDirs = "/invalid/grok/library/dirs"
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithNonexistentGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs := "/nonexistent/grok/library/dirs"
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in []
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to []
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		defaultGrokLibraryDirs = "/nonexistent/grok/library/dirs"
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithEmptyGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs = ""
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON log file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in []
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to []
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		defaultGrokLibraryDirs = ""
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithInvalidGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs = "/invalid/path"
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON log file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in []
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>                         Number of tail lines. 10 by default, for follow mode
       --reset-grok-library-dir                               Save default GROK patterns to []
  -t,  --template                                             Print a config YAML file template
  -V,  --version                                              Display app version information

`) + "\n"

	// Act
	actualOutput := captureStdout(func() {
		defaultGrokLibraryDirs = "/invalid/path"
		PrintHelp()
	})

	// Assert
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPrintHelpWithValidGrokLibraryDirs(t *testing.T) {
	// Arrange
	defaultGrokLibraryDirs = "/valid/path"
	expectedOutput := strings.TrimSpace(`
Convert and view structured (JSON) log
` + static.AppVersion + `
Usage:
  jog  [option...]  <your JSON log file path>
    or
  cat  <your JSON log file path>  |  jog  [option...]
Examples:
   1) follow with last 10 lines:         jog -f app-20200701-1.log
   2) follow with specified lines:       jog -n 100 -f app-20200701-1.log
   3) with specified config file:        jog -c another.jog.yml app-20200701-1.log
   4) view docker-compose log:           docker-compose logs | jog
   5) print the default template:        jog -t
   6) only shows WARN & ERROR level:     jog -l warn -l error app-20200701-1.log
   7) shows with timestamp range:        jog --after 2020-7-1 --before 2020-7-3 app-20200701-1.log
   8) natural timestamp range:           jog --after "1 week" --before "2 days" app-20200701-1.log
   9) output raw JSON and apply time range filter:      jog --after "1 week" --before "2 days" app-20200701-1.log --json
   10) disable colorization:             jog -cs colorization=false app-20200701-1.log
   11) view apache log, non-JSON log     jog -g COMMONAPACHELOG example_logs/grok_apache.log
Options:
  -a,  --after <timestamp>                                    'after' time filter. Auto-detect the timestamp format; can be natural datetime 
  -b,  --before <timestamp>                                   'before' time filter. Auto-detect the timestamp format; can be natural datetime 
  -c,  --config <config file path>                            Specify config YAML file path. The default is .jog.yaml or $HOME/.jog.yaml 
  -cs, --config-set <config item path>=<config item value>    Set value to specified config item 
  -cg, --config-get <config item path>                        Get value to specified config item 
  -d,  --debug                                                Print more error detail
  -f,  --follow                                               Follow mode - follow log output
  -g,  --grok <grok pattern name>                             For non-json log line. The default patterns are saved in []
  -h,  --help                                                 Display this information
  -j,  --json                                                 Output the raw JSON but then able to apply filters
  -l,  --level <level value>                                  Filter by log level. For ex. --level warn 
  -n,  --lines <number of tail lines>
