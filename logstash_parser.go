package main

// LogstashParserT implements LogParser interface
type LogstashParserT struct {
}

// LogstashParser is pointer of LogstashParserT
type LogstashParser = *LogstashParserT

// Parse the log event
func (me LogstashParser) Parse(event LogEvent) int {
	amountOfFieldsPopulated := 0

	for fieldName, fieldValue := range event.All {
		if fieldName == "@timestamp" {
			event.Timestamp = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}
		if fieldName == "@version" {
			event.Version = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}
		if fieldName == "message" {
			event.Message = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}
		if fieldName == "logger_name" {
			event.Logger = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}
		if fieldName == "thread_name" {
			event.Thread = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}
		if fieldName == "level" {
			event.Level = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}
		if fieldName == "level_value" {
			amountOfFieldsPopulated++
			// skip
			continue
		}
		if fieldName == "stack_trace" {
			event.StackTrace = fieldValue.(string)
			amountOfFieldsPopulated++
			continue
		}

		event.Others[fieldName] = fieldValue
	}

	return amountOfFieldsPopulated
}