package parse_test

import (
	"testing"

	"github.com/soulteary/acm-fellows-api/model/define"
	"github.com/soulteary/acm-fellows-api/model/parse"
)

func TestGetListFromPage(t *testing.T) {
	pageData := []byte(`
	<table>
		<tbody>
			<tr role="row">
				<td><a href="http://example.com">John Doe</a></td>
				<td></td>
				<td>2019</td>
				<td>North America</td>
				<td></td>
			</tr>
			<tr role="row">
				<td><a href="http://example.com/jane-doe">Jane Doe</a></td>
				<td></td>
				<td>2018</td>
				<td>Europe</td>
				<td></td>
			</tr>
		</tbody>
	</table>`)

	expectedFellows := []define.Fellow{
		{Name: "John Doe", URL: "http://example.com", Year: "2019", Region: "North America"},
		{Name: "Jane Doe", URL: "http://example.com/jane-doe", Year: "2018", Region: "Europe"},
	}

	fellows, err := parse.GetListFromPage(pageData)
	if err != nil {
		t.Fatal(err)
	}

	if len(fellows) != len(expectedFellows) {
		t.Fatalf("GetListFromPage() = %d fellows, want %d fellows", len(fellows), len(expectedFellows))
	}

	for i := range fellows {
		if fellows[i].Name != expectedFellows[i].Name {
			t.Errorf("GetListFromPage() = %q, want %q", fellows[i].Name, expectedFellows[i].Name)
		}
		if fellows[i].URL != expectedFellows[i].URL {
			t.Errorf("GetListFromPage() = %q, want %q", fellows[i].URL, expectedFellows[i].URL)
		}
		if fellows[i].Year != expectedFellows[i].Year {
			t.Errorf("GetListFromPage() = %q, want %q", fellows[i].Year, expectedFellows[i].Year)
		}
		if fellows[i].Region != expectedFellows[i].Region {
			t.Errorf("GetListFromPage() = %q, want %q", fellows[i].Region, expectedFellows[i].Region)
		}
	}
}
