package getdoc

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMethod(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join("_testdata", "method.html"))
	if err != nil {
		t.Fatal(err)
	}

	v, err := ParseMethod(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	expected := &Method{
		Name:        "langpack.getDifference",
		Description: []string{"Get new strings in languagepack"},
		Parameters: map[string]string{
			"from_version": "Previous localization pack version",
			"lang_code":    "Language code",
			"lang_pack":    "Language pack",
		},
		Errors: []Error{
			{Code: 400, Type: "LANG_PACK_INVALID", Description: "The provided language pack is invalid"},
		},
	}
	require.Equal(t, expected, v)
}
