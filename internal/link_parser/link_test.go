package link_parser

import (
	"os"
	"testing"
)

// TestParseLink tests the parsing of a single link.
func TestParseSingleLink(t *testing.T) {
	content, err := os.ReadFile("data/example_1.html")
	if err != nil {
		t.Fatal(err)
	}

	expected := Link{
		Href: "/other-page",
		Text: "A link to another page",
	}

	links := Parse(string(content))

	t.Run("should parse the link", func(t *testing.T) {
		if len(links) != 1 {
			t.Errorf("Expected at least one link, got %d", len(links))
		}
	})

	t.Run("should extract the link's href", func(t *testing.T) {
		if links[0].Href != expected.Href {
			t.Errorf("Expected href to be `%s`, got %s", expected.Href, links[0].Href)
		}
	})

	t.Run("should save the link's text", func(t *testing.T) {
		if links[0].Text != expected.Text {
			t.Errorf("Expected text to be %s, got %s", expected.Text, links[0].Text)
		}
	})
}

func TestParseNestedLinkText(t *testing.T) {
	content, err := os.ReadFile("data/example_2.html")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Gophercises is on Github!"
	output := Parse(string(content))

	t.Run("should get the deeply nested link's text", func(t *testing.T) {
		// since there are two links, we need to get the second one which has nested text
		if output[1].Text != expected {
			t.Errorf("Expected %s, got %s", expected, output[1].Text)
		}
	})

}

// example_1 has one link
// example_2 has two links
// example_3 has a link within a deep nested html structure
// example_4 has a link with a comment inside the tag
func TestParseLinkContent(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected []Link
	}{
		{
			name: "finds a link",
			file: "data/example_1.html",
			expected: []Link{
				{Href: "/other-page", Text: "A link to another page"},
			},
		},
		{
			name: "finds multiple links",
			file: "data/example_2.html",
			expected: []Link{
				{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
				{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
			},
		},
		{
			name: "finds a deeply nested link",
			file: "data/example_3.html",
			expected: []Link{
				{Href: "#", Text: "Login"},
				{Href: "/lost", Text: "Lost? Need help?"},
				{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
			},
		},
		{
			name: "excludes comments from the link text",
			file: "data/example_4.html",
			expected: []Link{
				{Href: "/dog-cat", Text: "dog cat"},
			},
		},
	}

	for _, test := range tests {
		content, err := os.ReadFile(test.file)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(test.name, func(t *testing.T) {
			output := Parse(string(content))

			for idx, link := range output {
				if link.Href != test.expected[idx].Href {
					t.Errorf("Expected href to be `%s`, got %s", test.expected[idx].Href, link.Href)
				}

				if link.Text != test.expected[idx].Text {
					t.Errorf("Expected text to be `%s`, got %s", test.expected[idx].Text, link.Text)
				}
			}
		})
	}
}
