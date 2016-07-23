package vibrant

import (
	"image"
	_ "image/jpeg"
	"os"
	"sort"
	"strconv"
	"testing"
	"fmt"
	"golang.org/x/image/draw"
)

func ExamplePaletteBuilder() {
	file, _ := os.Open("test_files/1.jpg")
	decodedImage, _, _ := image.Decode(file)
	palette := NewPaletteBuilder(decodedImage).Generate()
	// Iterate over the swatches in the palette...
	for _, swatch := range palette.Swatches() {
		fmt.Printf("Swatch has color %v and population %d\n", swatch.RGBAInt(), swatch.Population())
	}
	for _, target := range palette.Targets() {//
		_ = palette.SwatchForTarget(target)
		// Do something with the swatch for a given target...
	}
}

func ExamplePaletteBuilder_maximumColorCount() {
	file, _ := os.Open("test_files/1.jpg")
	decodedImage, _, _ := image.Decode(file)
	// Use a custom color count.
	palette := NewPaletteBuilder(decodedImage).MaximumColorCount(32).Generate()
	// Iterate over the swatches in the palette...
	for _, swatch := range palette.Swatches() {
		fmt.Printf("Swatch has color %v and population %d\n", swatch.RGBAInt(), swatch.Population())
	}
	for _, target := range palette.Targets() {//
		_ = palette.SwatchForTarget(target)
		// Do something with the swatch for a given target...
	}
}

func ExamplePaletteBuilder_resizeImageArea() {
	file, _ := os.Open("test_files/1.jpg")
	decodedImage, _, _ := image.Decode(file)
	// Use a custom resize image area and scaler.
	palette := NewPaletteBuilder(decodedImage).ResizeImageArea(160 * 160).Scaler(draw.CatmullRom).Generate()
	// Iterate over the swatches in the palette...
	for _, swatch := range palette.Swatches() {
		fmt.Printf("Swatch has color %v and population %d\n", swatch.RGBAInt(), swatch.Population())
	}
	for _, target := range palette.Targets() {//
		_ = palette.SwatchForTarget(target)
		// Do something with the swatch for a given target...
	}
}

func TestPalette_Swatches(t *testing.T) {
	tests := map[string]map[uint32]Uint32Slice{
		"test_files/1.jpg": map[uint32]Uint32Slice{
			16: Uint32Slice([]uint32{
				0xffbbb893,
				0xff7a8199,
				0xff1e7f8f,
				0xff79585f,
				0xff6b403a,
				0xffc5b248,
				0xffa47a7b,
				0xff2b2129,
				0xff14bcdf,
				0xff80625b,
				0xffb26f44,
				0xff403233,
				0xff644e5a,
				0xff36496a,
				0xff3f394f,
				0xff795252,
			}),
			64: Uint32Slice([]uint32{
				0xffa27cd3,
				0xff403233,
				0xff23658d,
				0xffbdad75,
				0xff6dcce4,
				0xffa2d5cc,
				0xff70d3ad,
				0xff6e9a88,
				0xff644e5a,
				0xff15a0a9,
				0xffd43e34,
				0xffc77038,
				0xffc9b69a,
				0xff36496a,
				0xff923a2a,
				0xff3e228b,
				0xff777f40,
				0xff7d6a82,
				0xffaa8296,
				0xff3f394f,
				0xff795252,
				0xffb5a736,
				0xffd3ae42,
				0xffd6c950,
				0xff0ee4ea,
				0xff11afdd,
				0xffd4e9cf,
				0xff1d8173,
				0xffd5c2c6,
				0xffaaab97,
				0xff06a086,
				0xff38a280,
				0xff276dc6,
				0xff2a1c1c,
				0xff9a3168,
				0xffb33531,
				0xffe03e67,
				0xff9d6a43,
				0xffc97853,
				0xff774534,
				0xff5a3e3d,
				0xff1e8bd0,
				0xff66443c,
				0xff2c2451,
				0xff8255bf,
				0xffb6ac57,
				0xff7b55a2,
				0xff7d5e5b,
				0xffa77e78,
				0xff8b7376,
				0xffa2756d,
				0xff6d7ddd,
				0xff846564,
				0xff806c4a,
				0xff886247,
				0xff755768,
				0xff758756,
				0xff2f262e,
				0xff282637,
				0xff252b3e,
				0xffa78281,
				0xff785779,
				0xff88605c,
				0xff7c5858,
			}),
		},
		"test_files/2.jpg": map[uint32]Uint32Slice{
			16: Uint32Slice([]uint32{
				0xffd6d3b6,
				0xffb0b2b8,
				0xffaa9869,
				0xff415eb9,
				0xffa15e0d,
				0xff0d146e,
				0xff20201e,
				0xff2c2d41,
				0xff7c7b80,
				0xff826237,
				0xff977d3f,
				0xff52555f,
				0xff576177,
				0xff736a52,
				0xffe4e6e6,
				0xffdadbda,
			}),
			64: Uint32Slice([]uint32{
				0xffc0940c,
				0xff977d3f,
				0xff6282cf,
				0xff5a7499,
				0xff26201c,
				0xff52555f,
				0xff122cc1,
				0xff8194d2,
				0xfff4e126,
				0xffabb0b7,
				0xffdcd6b1,
				0xff576177,
				0xff736a52,
				0xffd2464c,
				0xffe4e6e6,
				0xff9a9793,
				0xff504b20,
				0xff817c7d,
				0xff2349d8,
				0xffbdb9b6,
				0xffcbad72,
				0xff0e1041,
				0xff83531a,
				0xffc8c6c5,
				0xffa76c0d,
				0xff4665d5,
				0xffd1ab39,
				0xff0e146f,
				0xff3647a2,
				0xff5b583f,
				0xff141c7d,
				0xff090e5f,
				0xff353b1a,
				0xffd4d3d3,
				0xff93561e,
				0xff0c0e1e,
				0xffdadbda,
				0xff393b43,
				0xff16218c,
				0xffa2a6ac,
				0xff969aa1,
				0xffb5a168,
				0xffeeb106,
				0xffbdc3cd,
				0xffcdced0,
				0xffd9c066,
				0xffd2cec8,
				0xffdb9e07,
				0xff9a4801,
				0xff868384,
				0xff817256,
				0xfff4e003,
				0xff1628aa,
				0xffb3351b,
				0xffa29a77,
				0xff15259b,
				0xff949dab,
				0xff74777f,
				0xff9a937c,
				0xff938e7c,
				0xff9c5816,
				0xff8c887c,
				0xffa75b09,
				0xffa06016,
			}),
		},
		"test_files/3.jpg": map[uint32]Uint32Slice{
			16: Uint32Slice([]uint32{
				0xff858e74,
				0xffc4a88a,
				0xff9c3c32,
				0xffa6a78a,
				0xff3e4c39,
				0xff224736,
				0xffcc504f,
				0xff653420,
				0xffede7d5,
				0xff677f69,
				0xff4c725d,
				0xffd3c6a9,
				0xffded4b9,
				0xff053125,
				0xffc6bc9b,
				0xffbeb598,
			}),
			64: Uint32Slice([]uint32{
				0xffc54343,
				0xff9eb091,
				0xffc88873,
				0xff66472e,
				0xffd25958,
				0xff78866e,
				0xffa37a5e,
				0xffa42b28,
				0xffeef0df,
				0xffe3dfc7,
				0xff657862,
				0xffc2a78b,
				0xffb0b396,
				0xff20200f,
				0xff602d19,
				0xff255543,
				0xff913f2d,
				0xff402612,
				0xff698570,
				0xffcbc5a7,
				0xffc6bc9b,
				0xff999779,
				0xffd9d2b6,
				0xff3c533e,
				0xff99443a,
				0xff45725e,
				0xffea6e6c,
				0xff482d19,
				0xffe5d6bc,
				0xff426854,
				0xff093e30,
				0xffdbc7ab,
				0xff959d84,
				0xffc3ad8c,
				0xff557762,
				0xff02261b,
				0xffc84e4e,
				0xff8f8f77,
				0xffa8a386,
				0xffbeb598,
				0xffefe1cc,
				0xff703321,
				0xff252d19,
				0xffb54640,
				0xff853525,
				0xff214635,
				0xff908a6e,
				0xffc4a27e,
				0xffa94037,
				0xff3c5f4e,
				0xff888263,
				0xff1b3c2d,
				0xffb2a988,
				0xff908363,
				0xff9ca790,
				0xff7c937b,
				0xfff8ede3,
				0xff9ba887,
				0xffa43833,
				0xff983828,
				0xffa0382f,
				0xff903830,
				0xff983830,
				0xffc0b090,
			}),
		},
		"test_files/4.jpg": map[uint32]Uint32Slice{
			16: Uint32Slice([]uint32{
				0xff4f2a29,
				0xffcb9f69,
				0xff6e7999,
				0xff13181b,
				0xffddbd72,
				0xffc49e46,
				0xffa9988a,
				0xff111015,
				0xff916f28,
				0xff9a8850,
				0xff9f803b,
				0xff917d59,
				0xffa89146,
				0xff131c2b,
				0xff191d2e,
				0xff101828,
			}),
			64: Uint32Slice([]uint32{
				0xff867650,
				0xff917d59,
				0xff970610,
				0xffdf8c9b,
				0xff3368d5,
				0xffedcdcf,
				0xffa89146,
				0xffbe893e,
				0xff893b16,
				0xff8a868e,
				0xff3b3336,
				0xffb32267,
				0xffdab943,
				0xff75302e,
				0xff6f4d2a,
				0xff5f2516,
				0xff7d623b,
				0xffcd9d71,
				0xff811d43,
				0xff996818,
				0xffd8b860,
				0xffa4acd6,
				0xffefe7ef,
				0xff6b2c23,
				0xff3e0409,
				0xffd751e1,
				0xffd3189b,
				0xffa5802a,
				0xff9e1955,
				0xffd355a6,
				0xff523125,
				0xffdcbc6f,
				0xff8c6f24,
				0xff471d16,
				0xffa87b73,
				0xffe6c47b,
				0xffb31ba7,
				0xffbea34a,
				0xffbd9642,
				0xffc19948,
				0xffe6c38f,
				0xff371914,
				0xfff065dd,
				0xffcea547,
				0xffeb95c3,
				0xffc7a55a,
				0xff281513,
				0xff3d2923,
				0xff131c2b,
				0xffa3853a,
				0xffc9a663,
				0xff372d2a,
				0xff9c1012,
				0xffaf8d3c,
				0xff656b75,
				0xff161013,
				0xff101015,
				0xff1a1513,
				0xff191d2e,
				0xffad9963,
				0xff342420,
				0xffb09c50,
				0xff10181d,
				0xff101828,
			}),
		},
	}
	for path, data := range tests {
		for maximumColorCount, expectedSwatches := range data {
			expectedSwatches.Sort()
			file, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			original, _, err := image.Decode(file)
			if err != nil {
				t.Fatal(err)
			}
			palette := NewPaletteBuilder(original).ResizeImageArea(0).MaximumColorCount(maximumColorCount).Generate()
			actualSwatches := palette.Swatches()
			if len(actualSwatches) != len(expectedSwatches) {
				t.Errorf("Expected %d swatches but generated %d for %s", len(expectedSwatches), len(actualSwatches), path)
			}
			for _, swatch := range actualSwatches {
				actual := swatch.RGBAInt().PackedRGBA()
				i := expectedSwatches.Search(actual)
				if !(i < len(expectedSwatches) && expectedSwatches[i] == actual) {
					t.Errorf("Swatch 0x%s was not expected in %s", strconv.FormatUint(uint64(actual), 16), path)
				}
			}
		}
	}
}

func TestPaletteBuilder_Region(t *testing.T) {
	file, err := os.Open("test_files/1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	original, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	tests := map[image.Rectangle]RGBAInt{
		original.Bounds():              RGBAInt(0xFF0FE1E8),
		image.Rect(205, 230, 260, 260): RGBAInt(0xFFF0054D),
		image.Rect(340, 380, 375, 410): RGBAInt(0xFF00F3F3),
		image.Rect(570, 380, 600, 400): RGBAInt(0xFF443EC0),
	}
	for region, expected := range tests {
		actual := NewPaletteBuilder(original).Region(region).Generate().VibrantSwatch().RGBAInt()
		if actual != expected {
			t.Errorf("Expected %v but generated %v as the vibrant color in %v\n", expected, actual, region)
		}
	}
}

func SearchUint32s(a []uint32, x uint32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

type Uint32Slice []uint32

func (p Uint32Slice) Len() int            { return len(p) }
func (p Uint32Slice) Less(i, j int) bool  { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p Uint32Slice) Sort()               { sort.Sort(p) }
func (p Uint32Slice) Search(x uint32) int { return SearchUint32s(p, x) }
