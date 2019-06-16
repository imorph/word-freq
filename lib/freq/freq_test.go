package freq

import (
	"testing"
	"sort"
)

// EqualSlice tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func EqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// EqualMap tells whether a and b contain the key-value pairs.
func EqualMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, val := range a {
		if val != b[key] {
			return false
		}
	}
	return true
}

func TestCleanup001(t *testing.T) {
	var inWord string
	var got, want string

	inWord = "a4bc2d5e"
	want = "a4bc2d5e"
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup002(t *testing.T) {
	var inWord string
	var got, want string

	inWord = "a4bc2d5e,"
	want = "a4bc2d5e"
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup003(t *testing.T) {
	var inWord string
	var got, want string

	inWord = "a4bc2d5e;"
	want = "a4bc2d5e"
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup004(t *testing.T) {
	var inWord string
	var got, want string

	inWord = "a4bc2d5e:"
	want = "a4bc2d5e"
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup005(t *testing.T) {
	var inWord string
	var got, want string

	inWord = "a4bc2d5e."
	want = "a4bc2d5e"
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup006(t *testing.T) {
	var inWord string
	var got, want string

	inWord = `a4bc2d5e
`
	want = `a4bc2d5e`
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup007(t *testing.T) {
	var inWord string
	var got, want string

	inWord = ``
	want = ``
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestCleanup008(t *testing.T) {
	var inWord string
	var got, want string

	inWord = `A4bC2d5e...
`
	want = `a4bc2d5e`
	got = CleanupWord(inWord)

	if got != want {
		t.Errorf("CleanupWord(inWord) == %q, want %q", got, want)
	}
}

func TestSplit001(t *testing.T) {
	var inText string
	var got, want []string

	inText = `bla bla`
	want = []string{"bla", "bla"}
	got = StringToCleanSlice(inText)

	if !EqualSlice(got, want) {
		t.Errorf("StringToCleanSlice(inText) == %q, want %q", got, want)
	}
}

func TestSplit002(t *testing.T) {
	var inText string
	var got, want []string

	inText = `bla, bla: bla`
	want = []string{"bla", "bla", "bla"}
	got = StringToCleanSlice(inText)

	if !EqualSlice(got, want) {
		t.Errorf("StringToCleanSlice(inText) == %q, want %q", got, want)
	}
}

func TestSplit003(t *testing.T) {
	var inText string
	var got, want []string

	inText = `Bla, bla: BLA.`
	want = []string{"bla", "bla", "bla"}
	got = StringToCleanSlice(inText)

	if !EqualSlice(got, want) {
		t.Errorf("StringToCleanSlice(inText) == %q, want %q", got, want)
	}
}

func TestStatMap001(t *testing.T) {
	words := []string{"bla", "bla", "bla"} 
	want := map[string]int{
		"bla": 3,
	}
	got := StatMapFromSlice(words)

	if !EqualMap(got, want) {
		t.Errorf("StringToCleanSlice(words) == %q, want %q", got, want)
	}
}

func TestStatMap002(t *testing.T) {
	words := []string{"bla", "bla", "bla", "asd", "qwe"} 
	want := map[string]int{
		"bla": 3,
		"asd": 1,
		"qwe": 1,
	}
	got := StatMapFromSlice(words)

	if !EqualMap(got, want) {
		t.Errorf("StringToCleanSlice(words) == %q, want %q", got, want)
	}
}

func TestStatMap003(t *testing.T) {
	words := []string{"", "", "", "", ""} 
	want := map[string]int{
		"": 5,
	}
	got := StatMapFromSlice(words)

	if !EqualMap(got, want) {
		t.Errorf("StringToCleanSlice(words) == %q, want %q", got, want)
	}
}

func TestStatMap004(t *testing.T) {
	words := []string{"ert", "hgf", "asf", "asd", "qwe"} 
	want := map[string]int{
		"ert": 1,
		"hgf": 1,
		"asf": 1,
		"asd": 1,
		"qwe": 1,
	}
	got := StatMapFromSlice(words)

	if !EqualMap(got, want) {
		t.Errorf("StringToCleanSlice(words) == %q, want %q", got, want)
	}
}

func TestStatMap005(t *testing.T) {
	var words []string 
	want := make(map[string]int)

	got := StatMapFromSlice(words)

	if !EqualMap(got, want) {
		t.Errorf("StringToCleanSlice(words) == %q, want %q", got, want)
	}
}

func TestTop001(t *testing.T) {
	stats := map[string]int{
		"ert": 1,
		"hgf": 1,
		"asf": 1,
		"asd": 1,
		"qwe": 1,
	}
	want := []string{"ert", "hgf", "asf", "asd", "qwe"}

	got := Top10(stats)

	sort.Strings(got)
	sort.Strings(want)

	if !EqualSlice(got, want) {
		t.Errorf("Top10(stats) == %q, want %q", got, want)
	}
}

func TestTop002(t *testing.T) {
	stats := map[string]int{
		"ert": 5,
		"hgf": 4,
		"asf": 3,
		"asd": 2,
		"qwe": 1,
	}
	want := []string{"ert", "hgf", "asf", "asd", "qwe"}

	got := Top10(stats)

	if !EqualSlice(got, want) {
		t.Errorf("Top10(stats) == %q, want %q", got, want)
	}
}

func TestTop003(t *testing.T) {
	stats := map[string]int{
		"ert": 1,
		"hgf": 2,
		"asf": 3,
		"asd": 4,
		"qwe": 5,
	}
	want := []string{"qwe", "asd", "asf", "hgf", "ert"}

	got := Top10(stats)

	if !EqualSlice(got, want) {
		t.Errorf("Top10(stats) == %q, want %q", got, want)
	}
}

func TestTop004(t *testing.T) {
	stats := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"10": 10,
		"11": 11,
		"12": 12,
		"13": 13,
		"14": 14,
	}
	want := []string{"14", "13", "12", "11", "10", "9", "8", "7", "6", "5"}

	got := Top10(stats)

	if !EqualSlice(got, want) {
		t.Errorf("Top10(stats) == %q, want %q", got, want)
	}
}

func TestTop005(t *testing.T) {
	stats := map[string]int{
		"1": 14,
		"2": 13,
		"3": 12,
		"4": 11,
		"5": 10,
		"6": 9,
		"7": 8,
		"8": 7,
		"9": 6,
		"10": 5,
		"11": 4,
		"12": 3,
		"13": 2,
		"14": 1,
	}
	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	got := Top10(stats)

	if !EqualSlice(got, want) {
		t.Errorf("Top10(stats) == %q, want %q", got, want)
	}
}

func TestTotal001(t *testing.T) {
	text := `1 1 1 1 1 1 1 1 1 1 1 1 1 1
2 2 2 2 2 2 2 2 2 2 2 2 2
3 3 3 3 3 3 3 3 3 3 3 3
4 4 4 4 4 4 4 4 4 4 4
5 5 5 5 5 5 5 5 5 5
6 6 6 6 6 6 6 6 6
7 7 7 7 7 7 7 7
8 8 8 8 8 8 8
9 9 9 9 9 9 10 10 10 10 10 11 11 11 11 12 12 12 13 13 14`

	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	got := Top10(StatMapFromSlice(StringToCleanSlice(text)))

	if !EqualSlice(got, want) {
		t.Errorf("Top10(StatMapFromSlice(StringToCleanSlice(text))) == %q, want %q", got, want)
	}
}

func TestTotal002(t *testing.T) {
	text := `1: 1 1 1 1 1 1 1 1 1, 1 1 1 1...
2 2 2 2 2 2 2 2 2 2 2 2 2
3 3 3 3 3 3 3 3 3 3 3 3
4 4 4 4; 4 4 4 4 4 4 4
5 5 5 5 5 5 5 5 5 5;
6 6 6 6 6 6 6 6 6
7 7 7 7 7 7 7 7;
8 8 8: 8 8 8 8
9 9 9 9 9 9 10 10 10 10 10 11 11 11 11 12 12 12 13 13 14...
`
	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	got := Top10(StatMapFromSlice(StringToCleanSlice(text)))

	if !EqualSlice(got, want) {
		t.Errorf("Top10(StatMapFromSlice(StringToCleanSlice(text))) == %q, want %q", got, want)
	}
}