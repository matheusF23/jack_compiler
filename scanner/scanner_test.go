package scanner

import (
	"encoding/binary"
	"os"
	"strings"
	"testing"
	"unicode/utf16"

	"jack_compiler/token"
)

func TestJackFilesAgainstReference(t *testing.T) {
	tests := []struct {
		name      string
		jack      string
		reference string
	}{
		{
			name:      "Main",
			jack:      "../tests/Square/Main.jack",
			reference: "../tests/Square/MainT.xml",
		},
		{
			name:      "Square",
			jack:      "../tests/Square/Square.jack",
			reference: "../tests/Square/SquareT.xml",
		},
		{
			name:      "SquareGame",
			jack:      "../tests/Square/SquareGame.jack",
			reference: "../tests/Square/SquareGameT.xml",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testJackFile(t, test.jack, test.reference)
		})
	}
}

func testJackFile(t *testing.T, jackPath string, xmlPath string) {
	t.Helper()

	// Lê o arquivo Jack.
	input, err := os.ReadFile(jackPath)
	if err != nil {
		t.Fatalf("erro ao ler %s: %v", jackPath, err)
	}

	// Lê o XML de referência.
	expectedXML, err := readReferenceXML(xmlPath)
	if err != nil {
		t.Fatalf("erro ao ler %s: %v", xmlPath, err)
	}

	// Cria o scanner.
	scan := NewScanner(input)

	// Gera o XML usando o scanner.
	actualXML := generateXML(scan)

	// Normaliza os dois conteúdos.
	actualXML = normalizeXML(actualXML)
	expectedXML = normalizeXML(expectedXML)

	// Compara.
	if actualXML != expectedXML {
		showDifference(t, actualXML, expectedXML)
		return
	}

	t.Log("XML corresponde à referência do Nand2Tetris")
}

func readReferenceXML(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// UTF-16 Little Endian.
	if len(data) >= 2 &&
		data[0] == 0xFF &&
		data[1] == 0xFE {

		data = data[2:]

		if len(data)%2 != 0 {
			return "", os.ErrInvalid
		}

		u16 := make([]uint16, len(data)/2)

		for i := range u16 {
			u16[i] = binary.LittleEndian.Uint16(data[i*2 : i*2+2])
		}

		result := string(utf16.Decode(u16))

		// Remove BOM Unicode, caso exista.
		result = strings.TrimPrefix(result, "\uFEFF")

		return result, nil
	}

	// UTF-8 com BOM.
	if len(data) >= 3 &&
		data[0] == 0xEF &&
		data[1] == 0xBB &&
		data[2] == 0xBF {

		data = data[3:]
	}

	return string(data), nil
}

func generateXML(scan *Scanner) string {
	var builder strings.Builder

	builder.WriteString("<tokens>\n")

	for tk := scan.NextToken(); tk.Type != token.EOF; tk = scan.NextToken() {
		builder.WriteString(tk.String())
		builder.WriteString("\n")
	}

	builder.WriteString("</tokens>\n")

	return builder.String()
}

func normalizeXML(xml string) string {
	// Remove BOM Unicode.
	xml = strings.TrimPrefix(xml, "\uFEFF")

	// Normaliza quebras de linha.
	xml = strings.ReplaceAll(xml, "\r\n", "\n")
	xml = strings.ReplaceAll(xml, "\r", "\n")

	// Remove espaços/quebras somente das extremidades.
	return strings.TrimSpace(xml)
}

func showDifference(t *testing.T, actual string, expected string) {
	t.Helper()

	actualLines := strings.Split(actual, "\n")
	expectedLines := strings.Split(expected, "\n")

	maxLines := len(actualLines)

	if len(expectedLines) > maxLines {
		maxLines = len(expectedLines)
	}

	for i := 0; i < maxLines; i++ {
		var actualLine string
		var expectedLine string

		if i < len(actualLines) {
			actualLine = actualLines[i]
		}

		if i < len(expectedLines) {
			expectedLine = expectedLines[i]
		}

		if actualLine != expectedLine {
			t.Fatalf(
				"XML diferente na linha %d\n\nEsperado:\n%q\n\nObtido:\n%q",
				i+1,
				expectedLine,
				actualLine,
			)
		}
	}

	t.Fatalf("os XMLs possuem diferenças")
}
