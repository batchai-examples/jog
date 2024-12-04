package static

import (
	"testing"
)

func TestDefaultConfiguration(t *testing.T) {
	expected := `color: default
compress-prefix:
  action: remove
  enabled: false
  separators: /, \
default-field-colors:
  color: default
  timestamp: default
description: ""
enum-values: {}
fields:
  color:
    alias: []
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  compress-prefix:
    action: remove
    enabled: false
    separators: /, \
  default-field-colors:
    color: default
    timestamp: default
  description:
    alias: []
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  enum-values: {}
  fields: {}
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @process_id]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  host:
    alias: [hostname, host-name, host_name, @host, @hostname, @host-name, @host_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  line:
    alias: [lineno, line-no, line_no, linenum, line-num, line_num, linenumber, line-number, line_number, @lineno, @line-no, @line_no, @linenum, @line-num, @line_num, @linenumber, @line-number, @line_number]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  logger:
    alias: [id, logger_name, logger-name, loggername, @id, @logger_name, @logger-name, @loggername, @logger]
    case-sensitive: false
    color: default, underscore
    compress-prefix:
      action: remove-non-first-letter
      enabled: true
      separators: . , /
      white-list: com.wxcount
    description: ""
    enum-values: {}
    type: string
  method:
    alias: [methodname, method-name, method_name, func, funcname, func-name, func_name, function, functionname, function-name, function_name, @method, @methodname, @method-name, @method_name, @func, @funcname, @func-name, @func_name, @function, @functionname, @function-name, @function_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  response:
    alias: [res, resp, @res, @resp, @response]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  request:
    alias: [req, request]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  stacktrace:
    alias: [stack, stacktrace]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  duration:
    alias: [dur, duration]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  level:
    alias: [lvl, level]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @process_id]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  host:
    alias: [hostname, host-name, host_name, @host, @hostname, @host-name, @host_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  line:
    alias: [lineno, line-no, line_no, linenum, line-num, line_num, linenumber, line-number, line_number, @lineno, @line-no, @line_no, @linenum, @line-num, @line_num, @linenumber, @line-number, @line_number]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  logger:
    alias: [id, logger_name, logger-name, loggername, @id, @logger_name, @logger-name, @loggername, @logger]
    case-sensitive: false
    color: default, underscore
    compress-prefix:
      action: remove-non-first-letter
      enabled: true
      separators: . , /
      white-list: com.wxcount
    description: ""
    enum-values: {}
    type: string
  method:
    alias: [methodname, method-name, method_name, func, funcname, func-name, func_name, function, functionname, function-name, function_name, @method, @methodname, @method-name, @method_name, @func, @funcname, @func-name, @func_name, @function, @functionname, @function-name, @function_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  request:
    alias: [req, request]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  stacktrace:
    alias: [stack, stacktrace]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  duration:
    alias: [dur, duration]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  level:
    alias: [lvl, level]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @process_id]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  host:
    alias: [hostname, host-name, host_name, @host, @hostname, @host-name, @host_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  line:
    alias: [lineno, line-no, line_no, linenum, line-num, line_num, linenumber, line-number, line_number, @lineno, @line-no, @line_no, @linenum, @line-num, @line_num, @linenumber, @line-number, @line_number]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  logger:
    alias: [id, logger_name, logger-name, loggername, @id, @logger_name, @logger-name, @loggername, @logger]
    case-sensitive: false
    color: default, underscore
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  method:
    alias: [methodname, method-name, method_name, func, funcname, func-name, func_name, function, functionname, function-name, function_name, @method, @methodname, @method-name, @method_name, @func, @funcname, @func-name, @func_name, @function, @functionname, @function-name, @function_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  request:
    alias: [req, request]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  stacktrace:
    alias: [stack, stacktrace]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  duration:
    alias: [dur, duration]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  level:
    alias: [lvl, level]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @process_id]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  host:
    alias: [hostname, host-name, host_name, @host, @hostname, @host-name, @host_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  line:
    alias: [lineno, line-no, line_no, linenum, line-num, line_num, linenumber, line-number, line_number, @lineno, @line-no, @line_no, @linenum, @line-num, @line_num, @linenumber, @line-number, @line_number]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  logger:
    alias: [id, logger_name, logger-name, loggername, @id, @logger_name, @logger-name, @loggername, @logger]
    case-sensitive: false
    color: default, underscore
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  method:
    alias: [methodname, method-name, method_name, func, funcname, func-name, func_name, function, functionname, function-name, function_name, @method, @methodname, @method-name, @method_name, @func, @funcname, @func-name, @func_name, @function, @functionname, @function-name, @function_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  request:
    alias: [req, request]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  stacktrace:
    alias: [stack, stacktrace]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  duration:
    alias: [dur, duration]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  level:
    alias: [lvl, level]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @process_id]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  host:
    alias: [hostname, host-name, host_name, @host, @hostname, @host-name, @host_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  line:
    alias: [lineno, line-no, line_no, linenum, line-num, line_num, linenumber, line-number, line_number, @lineno, @line-no, @line_no, @linenum, @line-num, @line_num, @linenumber, @line-number, @line_number]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  logger:
    alias: [id, logger_name, logger-name, loggername, @id, @logger_name, @logger-name, @loggername, @logger]
    case-sensitive: false
    color: default, underscore
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  method:
    alias: [methodname, method-name, method_name, func, funcname, func-name, func_name, function, functionname, function-name, function_name, @method, @methodname, @method-name, @method_name, @func, @funcname, @func-name, @func_name, @function, @functionname, @function-name, @function_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  request:
    alias: [req, request]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  stacktrace:
    alias: [stack, stacktrace]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  duration:
    alias: [dur, duration]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  level:
    alias: [lvl, level]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @process_id]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  host:
    alias: [hostname, host-name, host_name, @host, @hostname, @host-name, @host_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  line:
    alias: [lineno, line-no, line_no, linenum, line-num, line_num, linenumber, line-number, line_number, @lineno, @line-no, @line_no, @linenum, @line-num, @line_num, @linenumber, @line-number, @line_number]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  logger:
    alias: [id, logger_name, logger-name, loggername, @id, @logger_name, @logger-name, @loggername, @logger]
    case-sensitive: false
    color: default, underscore
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  method:
    alias: [methodname, method-name, method_name, func, funcname, func-name, func_name, function, functionname, function-name, function_name, @method, @methodname, @method-name, @method_name, @func, @funcname, @func-name, @func_name, @function, @functionname, @function-name, @function_name]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  request:
    alias: [req, request]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  stacktrace:
    alias: [stack, stacktrace]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  duration:
    alias: [dur, duration]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  level:
    alias: [lvl, level]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  message:
    alias: [msg, @msg, @message]
    case-sensitive: false
    color: cyan
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  timestamp:
    alias: [time, date, datetime, date-time, date_time, @time, @timestamp, @date, @datetime, @date-time, @date_time]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  version:
    alias: [ver, @ver, @version]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  user:
    alias: [usr, username, user-name, user_name, @usr, @username, @user-name, @user_name, @user]
    case-sensitive: false
    color: default
    compress-prefix:
      action: remove
      enabled: false
      separators: /, \
    description: ""
    enum-values: {}
    type: string
  pid:
    alias: [process, process-id, processid, process_id, @pid, @process, @process-id, @processid, @
