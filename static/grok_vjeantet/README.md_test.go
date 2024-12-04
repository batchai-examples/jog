package grok_vjeantet

import (
	"testing"
)

// TestREADME_md checks if the README_md constant contains the expected content.
func TestREADME_md(t *testing.T) {
	expected := `The default GROK patterns in ~/.jog/grok_vjeantet are copy of https://github.com/vjeantet/grok/tree/master/patterns v1.0.0

https://github.com/vjeantet/grok LICENSE file: ~/.jog/grok_vjeantet.LICENSE`
	if README_md != expected {
		t.Errorf("README_md does not match the expected content. Expected:\n%s\nGot:\n%s", expected, README_md)
	}
}

// TestREADME_md_Empty checks if the README_md constant is not empty.
func TestREADME_md_Empty(t *testing.T) {
	if len(README_md) == 0 {
		t.Errorf("README_md should not be empty")
	}
}

// TestREADME_md_ContainsURL checks if the README_md constant contains a URL.
func TestREADME_md_ContainsURL(t *testing.T) {
	url := "https://github.com/vjeantet/grok/tree/master/patterns v1.0.0"
	if !strings.Contains(README_md, url) {
		t.Errorf("README_md should contain the URL: %s", url)
	}
}

// TestREADME_md_ContainsLicenseFile checks if the README_md constant contains a license file reference.
func TestREADME_md_ContainsLicenseFile(t *testing.T) {
	licenseFile := "https://github.com/vjeantet/grok LICENSE file: ~/.jog/grok_vjeantet.LICENSE"
	if !strings.Contains(README_md, licenseFile) {
		t.Errorf("README_md should contain the license file reference: %s", licenseFile)
	}
}
