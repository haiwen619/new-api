package docs

import _ "embed"

//go:embed katuReadme.md
var katuReadme string

func GetKatuReadme() string {
	return katuReadme
}
