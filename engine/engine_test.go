package engine

import (
	"testing"
)

func TestEngine_ShowThrowIfInputFileNotExists(t *testing.T) {
	nonExistingInputFile := "nonexisting.csv"
	res, err := Process(nonExistingInputFile)
	if res != "" {
		t.Error("Expected empty string")
	}
	if err == nil {
		t.Error("Expected error got nil")
	}
}

func TestEngine_ShowThrowIfUnallowedFileExtension(t *testing.T) {
	inputFile := "Workbook2.json"
	res, err := Process(inputFile)
	if res != "" {
		t.Error("Expected empty string")
	}
	if err == nil {
		t.Error("Expected error got nil")
	}
}

func TestEngine_ShowReturnHtmlTableForCsvInput(t *testing.T) {
	inputFile := "test.csv"
	expectedhtml := "<table><thead><tr><th>Name</th><th>Address</th><th>Postcode</th><th>Phone</th><th>Credit Limit</th><th>Birthday</th></thead><tbody><tr><td>Johnson, John</td><td>Voorstraat 32</td><td>3122gg</td><td>020 3849381</td><td>10000</td><td>01/01/1987</td></tr><tr><td>Anderson, Paul</td><td>Dorpsplein 3A</td><td>4532 AA</td><td>030 3458986</td><td>109093</td><td>03/12/1965</td></tr><tr><td>Wicket, Steve</td><td>Mendelssohnstraat 54d</td><td>3423 ba</td><td>0313-398475</td><td>934</td><td>03/06/1964</td></tr><tr><td>Benetar, Pat</td><td>Driehoog 3zwart</td><td>2340 CC</td><td>06-28938945</td><td>54</td><td>04/09/1964</td></tr><tr><td>Gibson, Mal</td><td>Vredenburg 21</td><td>3209 DD</td><td>06-48958986</td><td>54.5</td><td>09/11/1978</td></tr><tr><td>Friendly, User</td><td>Sint Jansstraat 32</td><td>4220 EE</td><td>0885-291029</td><td>63.6</td><td>10/08/1980</td></tr><tr><td>Smith, John</td><td>Børkestraße 32</td><td>87823</td><td>+44 728 889838</td><td>9898.3</td><td>20/09/1999</td></tr></tbody></table>"
	res, err := Process(inputFile)

	if res != expectedhtml {
		t.Error("Results don't match")
	}
	if err != nil {
		t.Error("Expected nil got error")
	}
}

func TestEngine_ShowReturnHtmlTableForPrnInput(t *testing.T) {
	inputFile := "test.prn"
	expectedhtml := "<table><thead><tr><th>Name</th><th>Address</th><th>Postcode</th><th>Phone</th><th>Credit Limit</th><th>Birthday</th></thead><tbody><tr><td>Johnson, John</td><td>Voorstraat 32</td><td>3122gg</td><td>020 3849381</td><td>1000000</td><td>19870101</td></tr><tr><td>Anderson, Paul</td><td>Dorpsplein 3A</td><td>4532 AA</td><td>030 3458986</td><td>10909300</td><td>19651203</td></tr><tr><td>Wicket, Steve</td><td>Mendelssohnstraat 54d</td><td> 3423 ba</td><td>0313-398475</td><td>93400</td><td>19640603</td></tr><tr><td>Benetar, Pat</td><td>Driehoog 3zwart</td><td>2340 CC</td><td>06-28938945</td><td>54</td><td>19640904</td></tr><tr><td>Gibson, Mal</td><td>Vredenburg 21</td><td>3209 DD</td><td>06-48958986</td><td>5450</td><td>19781109</td></tr><tr><td>Friendly, User</td><td>Sint Jansstraat 32</td><td>4220 EE</td><td>0885-291029</td><td>6360</td><td>19800810</td></tr><tr><td>Smith, John</td><td>Børkestraße 32</td><td>87823</td><td>+44 728 889838</td><td>989830</td><td>19990920</td></tr></tbody></table>"
	res, err := Process(inputFile)

	if res != expectedhtml {
		t.Error("Results don't match")
	}
	if err != nil {
		t.Error("Expected nil got error")
	}
}
