package freq

import "testing"

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

func TestSplit001(t *testing.T) {
	var inText string
	var got, want []string

	inText = `bla bla`
	want = []string{"bla", "bla"}
	got = StringToCleanSlice(inText)

	if !Equal(got, want) {
		t.Errorf("StringToCleanSlice(inText) == %q, want %q", got, want)
	}
}

func TestSplit002(t *testing.T) {
	var inText string
	var got, want []string

	inText = `bla, bla: bla`
	want = []string{"bla", "bla", "bla"}
	got = StringToCleanSlice(inText)

	if !Equal(got, want) {
		t.Errorf("StringToCleanSlice(inText) == %q, want %q", got, want)
	}
}

func TestSplit003(t *testing.T) {
	var inText string
	var got, want []string

	inText = `Bla, bla: BLA.`
	want = []string{"bla", "bla", "bla"}
	got = StringToCleanSlice(inText)

	if !Equal(got, want) {
		t.Errorf("StringToCleanSlice(inText) == %q, want %q", got, want)
	}
}