// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de_DE": &dictionary{index: de_DEIndex, data: de_DEData},
		"en_GB": &dictionary{index: en_GBIndex, data: en_GBData},
		"fr_CH": &dictionary{index: fr_CHIndex, data: fr_CHData},
	}
	fallback := language.MustParse("en-GB")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%d books available": 1,
	"Launching soon!":    2,
	"Welcome!":           0,
}

var de_DEIndex = []uint32{ // 4 elements
	0x00000000, 0x0000000c, 0x00000045, 0x00000056,
} // Size: 40 bytes

const de_DEData string = "" + // Size: 86 bytes
	"\x02Willkommen!\x14\x01\x81\x01\x00=\x01\x15\x02Ein Buch erhältlich\x00" +
	"\x1a\x02%[1]d Bücher erhältlich\x02Bald verfügbar!"

var en_GBIndex = []uint32{ // 4 elements
	0x00000000, 0x00000009, 0x0000003c, 0x0000004c,
} // Size: 40 bytes

const en_GBData string = "" + // Size: 76 bytes
	"\x02Welcome!\x14\x01\x81\x01\x00=\x01\x13\x02One book available\x00\x16" +
	"\x02%[1]d books available\x02Launching soon!"

var fr_CHIndex = []uint32{ // 4 elements
	0x00000000, 0x0000000c, 0x00000043, 0x00000059,
} // Size: 40 bytes

const fr_CHData string = "" + // Size: 89 bytes
	"\x02Bienvenue !\x14\x01\x81\x01\x00=\x01\x14\x02Un livre disponible\x00" +
	"\x19\x02%[1]d livres disponibles\x02Bientôt disponible !"

	// Total table size 371 bytes (0KiB); checksum: 47347436
