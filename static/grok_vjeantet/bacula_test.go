package grok_vjeantet

import (
	"testing"
)

func TestBaculaLogMaxCapacity(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_MAX_CAPACITY User defined maximum volume capacity 10,20,30 exceeded on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_MAX_CAPACITY User defined maximum volume capacity 100,200,300 exceeded on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_MAX_CAPACITY User defined maximum volume capacity 1,2,3 exceeded on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_MAX_CAPACITY User defined maximum volume capacity 1000,2000,3000 exceeded on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogMaxCapacity(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogMaxCapacity(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogEndVolume(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_END_VOLUME End of volume reached on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_END_VOLUME End of volume reached on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_END_VOLUME End of volume reached on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_END_VOLUME End of volume reached on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogEndVolume(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogEndVolume(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogNewVolume(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_NEW_VOLUME New volume created on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NEW_VOLUME New volume created on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NEW_VOLUME New volume created on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_NEW_VOLUME New volume created on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogNewVolume(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogNewVolume(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogNewLabel(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_NEW_LABEL New label created on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NEW_LABEL New label created on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NEW_LABEL New label created on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_NEW_LABEL New label created on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogNewLabel(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogNewLabel(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogWroteLabel(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_WROTE_LABEL Label written on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WROTE_LABEL Label written on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WROTE_LABEL Label written on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_WROTE_LABEL Label written on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogWroteLabel(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogWroteLabel(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogMounted(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_MOUNTED Volume mounted on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_MOUNTED Volume mounted on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_MOUNTED Volume mounted on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_MOUNTED Volume mounted on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogMounted(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogMounted(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogUnmounted(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_UNMOUNTED Volume unmounted on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_UNMOUNTED Volume unmounted on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_UNMOUNTED Volume unmounted on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_UNMOUNTED Volume unmounted on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogUnmounted(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogUnmounted(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogReadError(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_READ_ERROR Read error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_READ_ERROR Read error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_READ_ERROR Read error on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_READ_ERROR Read error on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogReadError(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogReadError(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogWriteError(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_WRITE_ERROR Write error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WRITE_ERROR Write error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WRITE_ERROR Write error on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_WRITE_ERROR Write error on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogWriteError(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogWriteError(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogMediaChanged(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_MEDIA_CHANGED Media changed on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_MEDIA_CHANGED Media changed on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_MEDIA_CHANGED Media changed on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_MEDIA_CHANGED Media changed on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogMediaChanged(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogMediaChanged(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogNoMedia(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_NO_MEDIA No media on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NO_MEDIA No media on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NO_MEDIA No media on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_NO_MEDIA No media on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogNoMedia(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogNoMedia(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogFull(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_FULL Volume full on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_FULL Volume full on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_FULL Volume full on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_FULL Volume full on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogFull(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogFull(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogWarning(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_WARNING Warning on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WARNING Warning on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WARNING Warning on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_WARNING Warning on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogWarning(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogWarning(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogError(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_ERROR Error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ERROR Error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ERROR Error on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_ERROR Error on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogError(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogError(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogCritical(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogCritical(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogCritical(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogDebug(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_DEBUG Debug message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_DEBUG Debug message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_DEBUG Debug message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_DEBUG Debug message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogDebug(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogDebug(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogInfo(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_INFO Informational message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_INFO Informational message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_INFO Informational message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_INFO Informational message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogInfo(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogInfo(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogNotice(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_NOTICE Notice message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NOTICE Notice message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_NOTICE Notice message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_NOTICE Notice message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogNotice(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogNotice(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogEmergency(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_EMERGENCY Emergency message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_EMERGENCY Emergency message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_EMERGENCY Emergency message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_EMERGENCY Emergency message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogEmergency(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogEmergency(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogAlert(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogAlert(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogAlert(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogWarning(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogWarning(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogWarning(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogError(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogError(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogError(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogCritical(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogCritical(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogCritical(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogAlert(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogAlert(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogAlert(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogWarning(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_WARNING Warning message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogWarning(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogWarning(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogError(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_ERROR Error message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogError(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogError(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogCritical(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_CRITICAL Critical error on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogCritical(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogCritical(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBaculaLogAlert(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: false,
		},
		{
			input:    `BACULA_LOG_ALERT Alert message on device "device" (/path/to/device)`,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := IsBaculaLogAlert(tc.input)
		if result != tc.expected {
			t.Errorf("IsBaculaLogAlert(%q) = %v; want %v", tc.input, result, tc.expected)
