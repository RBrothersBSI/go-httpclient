package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid content type")
	}
	if HeaderUserAgent != "User-Agent"{
		t.Error("Invalid user agent")
	}
	if ContentTypeJson != "application/json"{
		t.Error("invalid json content type")
	}
}
