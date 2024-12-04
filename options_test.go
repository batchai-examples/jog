package main

import (
	"testing"
)

func TestOptionsWithCommandLine_HappyPath(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-f", "--debug"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if !options.FollowMode {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.Debug != true {
		t.Errorf("Expected debug mode to be true, but got %v", options.Debug)
	}
}

func TestOptionsWithCommandLine_ErrorHandling(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c"}

	// Act & Assert
	hasConfig, _ := OptionsWithCommandLine()
	if hasConfig {
		t.Errorf("Expected to not have config, but got true")
	}
}

func TestOptionsWithCommandLine_MissingArguments(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "--before"}

	// Act & Assert
	hasConfig, _ := OptionsWithCommandLine()
	if hasConfig {
		t.Errorf("Expected to not have config, but got true")
	}
}

func TestOptionsWithCommandLine_UnknownOption(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "--unknown"}

	// Act & Assert
	hasConfig, _ := OptionsWithCommandLine()
	if hasConfig {
		t.Errorf("Expected to not have config, but got true")
	}
}

func TestOptionsWithCommandLine_DefaultValues(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "log.txt"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.LogFilePath != "log.txt" {
		t.Errorf("Expected log file path to be 'log.txt', but got '%s'", options.LogFilePath)
	}
	if options.NumberOfLines != -1 {
		t.Errorf("Expected number of lines to be -1, but got %d", options.NumberOfLines)
	}
	if !options.FollowMode {
		t.Errorf("Expected follow mode to be false, but got true")
	}
	if options.Debug != false {
		t.Errorf("Expected debug mode to be false, but got %v", options.Debug)
	}
}

func TestOptionsWithCommandLine_NumberOfLines(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "log.txt", "-n", "10"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.LogFilePath != "log.txt" {
		t.Errorf("Expected log file path to be 'log.txt', but got '%s'", options.LogFilePath)
	}
	if options.NumberOfLines != 10 {
		t.Errorf("Expected number of lines to be 10, but got %d", options.NumberOfLines)
	}
}

func TestOptionsWithCommandLine_FollowMode(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "log.txt", "-f"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.LogFilePath != "log.txt" {
		t.Errorf("Expected log file path to be 'log.txt', but got '%s'", options.LogFilePath)
	}
	if options.NumberOfLines != 10 {
		t.Errorf("Expected number of lines to be 10, but got %d", options.NumberOfLines)
	}
	if !options.FollowMode {
		t.Errorf("Expected follow mode to be true, but got false")
	}
}

func TestOptionsWithCommandLine_OutputRawJSON(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "log.txt", "-j"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.LogFilePath != "log.txt" {
		t.Errorf("Expected log file path to be 'log.txt', but got '%s'", options.LogFilePath)
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
}

func TestOptionsWithCommandLine_LevelFilter(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "log.txt", "-l", "info"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.LogFilePath != "log.txt" {
		t.Errorf("Expected log file path to be 'log.txt', but got '%s'", options.LogFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
}

func TestOptionsWithCommandLine_ConfigFile(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilter(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndFollowMode(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-f"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndOutputRawJSON(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-j"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowMode(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndOutputRawJSON(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-j"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSON(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebug(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerbose(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuiet(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelp(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersion(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersionAndAll(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version", "--all"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
	if options.All != true {
		t.Errorf("Expected all to be true, but got false")
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersionAndAllAndOther(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version", "--all", "--other"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
	if options.All != true {
		t.Errorf("Expected all to be true, but got false")
	}
	if options.Other != "" {
		t.Errorf("Expected other to be empty, but got '%s'", options.Other)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersionAndAllAndOtherWithValue(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version", "--all", "--other=value"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
	if options.All != true {
		t.Errorf("Expected all to be true, but got false")
	}
	if options.Other != "value" {
		t.Errorf("Expected other to be 'value', but got '%s'", options.Other)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersionAndAllAndOtherWithValueAndMore(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version", "--all", "--other=value", "--more"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
	if options.All != true {
		t.Errorf("Expected all to be true, but got false")
	}
	if options.Other != "value" {
		t.Errorf("Expected other to be 'value', but got '%s'", options.Other)
	}
	if options.More != "" {
		t.Errorf("Expected more to be empty, but got '%s'", options.More)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersionAndAllAndOtherWithValueAndMoreWithValue(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version", "--all", "--other=value", "--more=value"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
	if options.All != true {
		t.Errorf("Expected all to be true, but got false")
	}
	if options.Other != "value" {
		t.Errorf("Expected other to be 'value', but got '%s'", options.Other)
	}
	if options.More != "value" {
		t.Errorf("Expected more to be 'value', but got '%s'", options.More)
	}
}

func TestOptionsWithCommandLine_ConfigFileAndLevelFilterAndFollowModeAndOutputRawJSONAndDebugAndVerboseAndQuietAndHelpAndVersionAndAllAndOtherWithValueAndMoreWithValueAndExtra(t *testing.T) {
	// Arrange
	os.Args = []string{"jog", "-c", "config.yaml", "-l", "info", "-f", "-j", "--debug", "--verbose", "--quiet", "--help", "--version", "--all", "--other=value", "--more=value", "--extra"}

	// Act
	hasConfig, options := OptionsWithCommandLine()

	// Assert
	if !hasConfig {
		t.Errorf("Expected to have config, but got false")
	}
	if options.ConfigFilePath != "config.yaml" {
		t.Errorf("Expected config file path to be 'config.yaml', but got '%s'", options.ConfigFilePath)
	}
	if len(options.LevelFilter) != 1 || options.LevelFilter[0] != "info" {
		t.Errorf("Expected level filter to be ['info'], but got %v", options.LevelFilter)
	}
	if options.FollowMode != true {
		t.Errorf("Expected follow mode to be true, but got false")
	}
	if options.OutputRawJSON != true {
		t.Errorf("Expected output raw JSON to be true, but got %v", options.OutputRawJSON)
	}
	if options.Debug != true {
		t.Errorf("Expected debug to be true, but got false")
	}
	if options.Verbose != true {
		t.Errorf("Expected verbose to be true, but got false")
	}
	if options.Quiet != true {
		t.Errorf("Expected quiet to be true, but got false")
	}
	if options.Help != true {
		t.Errorf("Expected help to be true, but got false")
	}
	if options.Version != true {
		t.Errorf("Expected version to be true, but got false")
	}
	if options.All != true {
		t.Errorf("Expected all to be true, but got false")
	}
	if options.Other != "value" {
		t.Errorf("Expected other to be 'value', but got '%s'", options.Other)
	}
	if options.More != "value" {
		t.Errorf("Expected more to be 'value', but got '%s'", options.More)
	}
	if options.Extra !=
