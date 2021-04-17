package payload

import (
	"testing"
)

func TestErrorResponse_VngErrorFactory(t *testing.T) {
	er := NewError()
	f := er.VngErrorFactory(
		"testing error",
		VndRequestURI("https://stafes.co.jp/"),
		VndAboutURI("https://stafes.co.jp/"))
	if f.Message != "testing error" {
		t.Errorf("it is not the expected value. / %s", f.Message)
	}
	if f.Links.About.Href != "https://stafes.co.jp/" {
		t.Errorf("it is not the expected value. / %s", f.Links.About.Href)
	}
	if f.Path != "https://stafes.co.jp/" {
		t.Errorf("it is not the expected value. / %s", f.Path)
	}
}
