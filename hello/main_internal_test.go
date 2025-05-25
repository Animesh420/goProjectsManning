package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	want := "Hello world"
	lang := language("en")
	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	want := "Bonjour le monde"
	lang := language("fr")
	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

// func TestGreet_Akkadian(t *testing.T) {
// 	// Akkadian is not implemented yet !
// 	lang := language("akk")
// 	want := ""
// 	got := greet(lang)

// 	if got != want {
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
