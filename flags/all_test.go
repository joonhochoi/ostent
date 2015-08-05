package flags

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestPeriodStr(t *testing.T) {
	for _, v := range []struct {
		dstr string
		cmp  string
		json []byte
	}{
		{"1h", "1h", []byte("\"1h\"")},
	} {
		td, err := time.ParseDuration(v.dstr)
		if err != nil {
			t.Error(err)
		}
		p := Period{Duration: td}
		cmp := p.String()
		if cmp != v.cmp {
			t.Errorf("Mismatch: %q vs %q", cmp, v.cmp)
		}
		j, err := p.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(j, v.json) {
			t.Errorf("Mismatch: %q vs %q", j, v.json)
		}
	}
}

func TestPeriodSet(t *testing.T) {
	for _, v := range []struct {
		dstr  string
		above string
		str   string
		err   error
	}{
		{"abc", "", "", fmt.Errorf("time: invalid duration abc")},
		{"0.5s", "", "", fmt.Errorf("Less than a second: 500ms")},
		{"1.5s", "", "", fmt.Errorf("Not a multiple of a second: 1.5s")},
		{"1s1s", "", "2s", nil},
		{"3s", "4s", "", fmt.Errorf("Should be above 4s: 3s")},
	} {
		td, err := time.ParseDuration(v.dstr)
		if err != nil && err.Error() != v.err.Error() {
			t.Error(err)
		}
		p := Period{Duration: td}
		if v.above != "" {
			if ad, err := time.ParseDuration(v.above); err != nil {
				t.Error(err)
			} else {
				p.Above = &ad
			}
		}
		err = p.Set(v.dstr)
		if err != nil {
			if err.Error() == v.err.Error() {
				continue
			}
			t.Error(err)
		}
		if p.String() != v.str {
			t.Errorf("Mismatch: %q vs %q", p.String(), v.str)
		}
	}
}

func TestBindSet(t *testing.T) {
	const defportint = 9050
	defport := fmt.Sprintf("%d", defportint)
	for _, v := range []struct {
		a    string
		cmp  string
		errs []error
	}{
		{"a:b:", "", []error{fmt.Errorf("too many colons in address a:b:")}},
		{"localhost:nonexistent", "", []error{
			fmt.Errorf("lookup tcp/nonexistent: Servname not supported for ai_socktype"),
			fmt.Errorf("lookup tcp/nonexistent: nodename nor servname provided, or not known"),
			fmt.Errorf("unknown port tcp/nonexistent"),
		}},
		{"localhost", "localhost:9050", nil},
		{"", ":9050", nil},
		{"8001", ":8001", nil},
		{"8001", ":8001", nil},
		{":8001", ":8001", nil},
		{"*:8001", ":8001", nil},
		{"127.1:8001", "127.1:8001", nil},
		{"127.0.0.1:8001", "127.0.0.1:8001", nil},
		{"127.0.0.1", "127.0.0.1:" + defport, nil},
		{"127", "127.0.0.1:" + defport, nil},
		{"127.1", "127.1:" + defport, nil},
	} {
		b := NewBind(defportint)
		if err := b.Set(v.a); err != nil {
			unknownerr := true
			for _, x := range v.errs {
				if err.Error() == x.Error() {
					unknownerr = false
					break
				}
			}
			if unknownerr {
				t.Errorf("Error: %q\nExpected errors: %+v\n", err, v.errs)
			}
			continue
		}
		if b.string != v.cmp {
			t.Errorf("Mismatch: Bind %v == %v != %v\n", v.a, v.cmp, b.string)
		}
	}
}
