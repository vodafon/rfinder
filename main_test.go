package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGoofleAPI(t *testing.T) {
	testFiles(t, "1.txt", "")
}

func testFiles(t *testing.T, suf, cfp string) {
	buf := &bytes.Buffer{}

	worker := Processor{
		w:   buf,
		cfg: MustLoadConfig(cfp),
	}

	for _, v := range lines(t, "testdata/t"+suf) {
		worker.Process(v)
	}

	exp := strings.Join(lines(t, "testdata/r"+suf), "\n") + "\n"
	res := buf.String()

	if exp != res {
		t.Errorf("Incorrect result. Expected\n%s\n, got\n%s\n", exp, res)
	}
}

func lines(t *testing.T, fp string) []string {
	_, err := os.Stat(fp)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(fp)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	res := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
}
