package lexer

import(
	"testing"
	"bufio"
	"os"
	"unicode/utf8"
)

func TestRead(t *testing.T){
	fi, _ := os.Open("input.txt")
	r := bufio.NewReader(fi)
	scanner := NewScanner(r)

	nextRune := scanner.read()
	runeValue, _ := utf8.DecodeRuneInString("S")

	if (nextRune != runeValue){
		t.Fail()
	}

	nextRune = scanner.read()
	runeValue, _ = utf8.DecodeRuneInString("E")

	if (nextRune != runeValue){
		t.Fail()
	}	

}

func TestScan(t *testing.T) {
	file, _ := os.Open("input.txt")	
	reader := bufio.NewReader(file)

	scanner := NewScanner(reader)

	tok := scanner.Scan()
	if (tok.Type != SELECT){
		t.Fail();
	}

	t.Log(tok)

	tok = scanner.Scan()
	if (tok.Type != WS){
		t.Fail();
	}

	t.Log(tok)

	tok = scanner.Scan()
	if (tok.Type != ASTERISK){
		t.Fail();
	}

	t.Log(tok)
}