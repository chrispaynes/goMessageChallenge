package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConvertMapToJSON(t *testing.T) {
	tt := []struct {
		name     string
		m        map[string]string
		expected string
		err      error
	}{
		{"Empty Email", map[string]string{
			"Date":         "",
			"From":         "",
			"To":           "",
			"Subject":      "",
			"Message-ID":   "",
			"Content-Type": "",
			"Body":         "",
		}, "{\"Date\":\"\",\"From\":\"\",\"To\":\"\",\"Subject\":\"\",\"MessageId\":\"\",\"ContentType\":\"\",\"Body\":\"\"}", nil},
		{"Full Email", map[string]string{
			"Date":         "Alpha",
			"From":         "Bravo",
			"To":           "Charlie",
			"Subject":      "Delta",
			"Message-ID":   "Echo",
			"Content-Type": "Foxtrot",
			"Body":         "Golf",
		}, "{\"Date\":\"Alpha\",\"From\":\"Bravo\",\"To\":\"Charlie\",\"Subject\":\"Delta\",\"MessageId\":\"Echo\",\"ContentType\":\"Foxtrot\",\"Body\":\"Golf\"}", nil},
		{"Missing Map Fields", map[string]string{
			"Date":         "",
			"From":         "",
			"Content-Type": "",
			"Body":         "",
		}, "{\"Date\":\"\",\"From\":\"\",\"To\":\"\",\"Subject\":\"\",\"MessageId\":\"\",\"ContentType\":\"\",\"Body\":\"\"}", nil},
		{"Extra Map Fields", map[string]string{
			"Date":         "",
			"From":         "",
			"To":           "",
			"Subject":      "",
			"Message-ID":   "",
			"Content-Type": "",
			"Body":         "",
			"Body2":        "",
		}, "{\"Date\":\"\",\"From\":\"\",\"To\":\"\",\"Subject\":\"\",\"MessageId\":\"\",\"ContentType\":\"\",\"Body\":\"\"}", nil},
		{"No Map Fields", map[string]string{}, "{\"Date\":\"\",\"From\":\"\",\"To\":\"\",\"Subject\":\"\",\"MessageId\":\"\",\"ContentType\":\"\",\"Body\":\"\"}", nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := convertMapToJSON(tc.m)

			if string(actual) != tc.expected {
				t.Fatalf("convertMapToJSON of %v should be %v; got %v", tc.name, tc.expected, string(actual))
			}

			if err != tc.err {
				t.Fatalf("convertMapToJSON of %v error should be %v; got %v", tc.name, tc.err, actual)
			}
		})
	}
}

func TestGetHealth(t *testing.T) {
	resp := httptest.NewRecorder()
	resp.Write([]byte("OK"))
	resp.WriteHeader(200)

	req, err := http.NewRequest("GET", "healthz", nil)

	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(resp, req)

	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if !strings.Contains(string(p), "OK") {
			t.Fatalf("health end point should return an \"OK\" response")
		}
	}
}
