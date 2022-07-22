package main

import (
	"testing"
)

func TestGetSearchEngine(t *testing.T) {
	testCases := []struct {
		Engine    string
		Expected  string
		ExpectErr bool
		Name      string
	}{
		{
			Engine:    "google",
			Expected:  "https://google.com/search",
			ExpectErr: false,
			Name:      "google engine",
		},
		{
			Engine:    "duckduckgo",
			Expected:  "https://duckduckgo.com",
			ExpectErr: false,
			Name:      "duckduckgo engine",
		},
		{
			Engine:    "ddg",
			Expected:  "https://duckduckgo.com",
			ExpectErr: false,
			Name:      "ddg engine",
		},
		{
			Engine:    "stackoverflow",
			Expected:  "https://stackoverflow.com/search",
			ExpectErr: false,
			Name:      "stackoverflow engine",
		},
		{
			Engine:    "github",
			Expected:  "https://github.com/search",
			ExpectErr: false,
			Name:      "github engine",
		},
		{
			Engine:    "bing",
			Expected:  "",
			ExpectErr: true,
			Name:      "bing engine",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			e, err := getSearchEngine(tc.Engine)

			if tc.ExpectErr {
				if err == nil {
					t.Errorf("expected error, got nil instead")
				}

				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if tc.Expected != e.URL() {
				t.Errorf("Expected %q, got %q instead", tc.Expected, e.URL())
			}
		})
	}
}
