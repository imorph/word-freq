package freq

import (
	"sort"
	"testing"
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
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
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
		"1":  14,
		"2":  13,
		"3":  12,
		"4":  11,
		"5":  10,
		"6":  9,
		"7":  8,
		"8":  7,
		"9":  6,
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

func TestTotal003(t *testing.T) {
	text := `Go is expressive, concise, clean, and efficient.
Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines,
while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet
has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled
language that feels like a dynamically typed, interpreted language.
One of Go's key design goals is code adaptability; that it should
be easy to take a simple design and build upon it in a clean and
natural way. In this talk Andrew Gerrand describes a simple chat roulette
server that matches pairs of incoming TCP connections, and then use Go's
concurrency mechanisms, interfaces, and standard library to extend it with
a web interface and other features. While the function of the program changes
dramatically, Go's flexibility preserves the original design as it grows.
Concurrency is the key to designing high performance network services.
Go's concurrency primitives (goroutines and channels) provide a simple
and efficient means of expressing concurrent execution.
In this talk we see how tricky concurrency problems can be solved gracefully with simple Go code.
An import path is a string that uniquely identifies a package. A package's import path corresponds
to its location inside a workspace or in a remote repository (explained below).
The packages from the standard library are given short import paths such as fmt and net/http.
For your own packages, must choose a base path that is unlikely to collide with future additions
to the standard library or other external libraries.
If you keep your code in a source repository somewhere,
then you should use the root of that source repository as your base path. For instance, if you have a GitHub account
at github.com/user, that should be your base path.
Note that you don't need to publish your code to a remote repository before you can build it.
It's just a good habit to organize your code as if you will publish someday. In practice you can choose any
arbitrary path name, as long as it is unique to the standard library and greater Go ecosystem.
We'll use github.com/user as our base path. Create a directory inside your workspace in which to keep source code:
In cryptanalysis, frequency analysis (also known as counting letters) is the study of the frequency of letters or
groups of letters in a ciphertext. The method is used as an aid to breaking classical ciphers.
Frequency analysis is based on the fact that, in any given stretch of written language, certain
letters and combinations of letters occur with varying frequencies. Moreover, there is a characteristic
distribution of letters that is roughly the same for almost all samples of that language. For instance,
given a section of English language, E, T, A and O are the most common, while Z, Q and X are rare. Likewise,
TH, ER, ON, and AN are the most common pairs of letters (termed bigrams or digraphs), and SS, EE, TT, and FF
are the most common repeats.[1] The nonsense phrase "ETAOIN SHRDLU" represents the 12 most frequent letters
in typical English language text.
In some ciphers, such properties of the natural language plaintext are preserved in the ciphertext, and these
patterns have the potential to be exploited in a ciphertext-only attack.`

	want := []string{"the", "a", "and", "of", "to", "in", "that", "is", `as`, "letters"}
	got := Top10(StatMapFromSlice(StringToCleanSlice(text)))

	if !EqualSlice(got, want) {
		t.Errorf("Top10(StatMapFromSlice(StringToCleanSlice(text))) == %q, want %q", got, want)
	}
}
