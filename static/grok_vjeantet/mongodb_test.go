package grok_vjeantet

import (
	"testing"
)

// TestMongodbHappyPath tests the Mongodb constant with a valid log entry.
func TestMongodbHappyPath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbCornerCase tests the Mongodb constant with a log entry that has missing fields.
func TestMongodbCornerCase(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// TestMongodbNegativeCornerCase tests the Mongodb constant with a log entry that has missing fields.
func TestMongodbNegativeCornerCase(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{NONNEGINT:ntoreturn} %{WORD}:%{NONNEGINT:ntoskip} %{WORD}:%{NONNEGINT:nscanned}.*nreturned:%{NONNEGINT:nreturned}..+ (?<duration>[0-9]+)ms
MONGO_WORDDASH \b[\w-]+\b
MONGO3_SEVERITY \w
MONGO3_COMPONENT %{WORD}|-
MONGO3_LOG %{TIMESTAMP_ISO8601:timestamp} %{MONGO3_SEVERITY:severity} %{MONGO3_COMPONENT:component}%{SPACE}(?:\[%{DATA:context}\])? %{GREEDYDATA:message}
`
	result := Mongodb
	if result == expected {
		t.Errorf("Expected an error, but got no error")
	}
}

// TestMongodbNegativePath tests the Mongodb constant with an invalid log entry.
func TestMongodbNegativePath(t *testing.T) {
	logEntry := `MONGO_LOG 2023-04-10T12:34:56Z [mongod] { "query": { "find": "users", "filter": {}, "limit": 10, "skip": 0, "sort": { "_id": 1 }, "fields": null, "hint": null, "returnKey": false, "showDiskLoc": false, "snapshot": false, "readConcern": null, "writeConcern": null, "collation": null, "maxTimeMS": null, "$db": "test" } } ntoreturn:10 ntoskip:0 nscanned:20 nreturned:5..+ 10ms`
	expected := `MONGO_LOG %{SYSLOGTIMESTAMP:timestamp} \[%{WORD:component}\] %{GREEDYDATA:message}
MONGO_QUERY \{ (?<={ ).*(?= } ntoreturn:) \}
MONGO_SLOWQUERY %{WORD} %{MONGO_WORDDASH:database}\.%{MONGO_WORDDASH:collection} %{WORD}: %{MONGO_QUERY:query} %{WORD}:%{
