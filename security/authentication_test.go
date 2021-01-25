package security

import (
	"net/url"
	"testing"
)

func TestCalculateSignature(t *testing.T) {
	data := [][4]string{
		[4]string{"939fe2a517cde7f092cc8bd7c3027a51a061f0b2", "aa2b78ea42cc035671754ebf36e54d97715b24a9", "/details", "2021-01-01T23:00:00Z"},
		[4]string{"6b07cb7ff9d8edea704eab029805fb5ac3b8c565", "eb5e30decbfbd7dd38a627a99c30a576aacf45c8", "/usagestatistics", "2021-01-03T23:14:00Z"},
		[4]string{"d95523a6855733de17e5022bb1fb7c6d8262b60c", "177482d5aa370f1741e10729159101cd45a63e0b", "/devices", "2021-01-01T12:00:00Z"},
		[4]string{"556e2f179ed65bdc2022a2b71f118ac061834f4d", "72aa298719365add6515d1996df88d6967c8e8bc", "/autocomplete", "2021-01-02T00:08:08Z"},
	}
	params := []url.Values{
		url.Values{"pbid": []string{"1000000249489"}},
		nil,
		url.Values{"did": []string{"52000"}},
		url.Values{"spec": []string{"8"}, "query": []string{"foobar"}, "count": []string{"50"}},
	}
	expected := []string{
		"2e3283e367176df6d9bb96fb528e0144dc5f4d1d",
		"3846b641e12ca971b89f5297886c83a5197724e1",
		"22a009ed6e226466428e19dc3cd11c4dd6653038",
		"58c129f33288f22d22ba7a687f1884c3c7b44836",
	}
	for k, in := range data {
		if u := calculateSignature(in[0], in[1], in[2], in[3], &params[k]); u != expected[k] {
			t.Errorf("Client.makeURL('%s', '%s', '%s', '%s', '%v'): found '%s', expected '%s'", in[0], in[1], in[2], in[3], params[k], u, expected[k])
		}
	}
}
