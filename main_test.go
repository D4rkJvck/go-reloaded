package main

import (
	"fmt"
	"testing"
)

func Test_OverAll(t *testing.T) {

	input := "'if (    cap    ) I make you a BREAKFAST IN BED(  low  ,11(bin)           ),  just say thank you instead of :' hOW(  cap  ) did you get in my,house (  cap  ,  2  ) ? ' I have to pack 101 (  bin  ) outfits   . Packed 1a (  hex  ) , , , , , just to be sure .Don't be sad ,because sad backwards is das . And das not good . harold wilson (  cap,2  )(up, 2) :'I am a optimist , but a optimist who carries a raincoat  .  .  . '"
	input = FixSpaces(input)
	input = FixPunctuation(input)
	input = FixParentheses(input)
	input = ApplyInstances(FixAn(input))
	input = FixPunctuation(input)
	input = FixApostrophic(input)
	input = FixAn(input)
	expect := "'If I make you a BREAKFAST IN BED (low, 11 (bin)), just say thank you instead of:' How did you get in My, House? 'I have to pack 5 outfits. Packed 26,,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. HAROLD WILSON:' I am an optimist, but an optimist who carries a raincoat... '"
	
	if input != expect {
		t.Errorf("💥FAILED💥\nResult -> %v\nExpected -> %v\n\n", input, expect)
	} else {
		fmt.Print("\nOverAll: ✅\n\n\n")
	}
}

func Test_Functions(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		fix    func(string) string
		expect string
	}{

		{
			"Spaces",
			"'if (    cap    ) I make you a BREAKFAST IN BED(  low  ,11(bin)           ),  just say thank you instead of :' hOW(  cap  ) did you get in my,house (  cap  ,  2  ) ? ' I have to pack 101 (  bin  ) outfits   . Packed 1a (  hex  ) , , , , , just to be sure .Don't be sad ,because sad backwards is das . And das not good . harold wilson (  cap,2  )(up, 2) :'I am a optimist , but a optimist who carries a raincoat  .  .  . '",
			FixSpaces,
			"'if ( cap ) I make you a BREAKFAST IN BED( low ,11(bin) ), just say thank you instead of :' hOW( cap ) did you get in my,house ( cap , 2 ) ? ' I have to pack 101 ( bin ) outfits . Packed 1a ( hex ) , , , , , just to be sure .Don't be sad ,because sad backwards is das . And das not good . harold wilson ( cap,2 )(up, 2) :'I am a optimist , but a optimist who carries a raincoat . . . '",
		},
		{
			"Punctuation",
			"",
			FixPunctuation,
			"'if ( cap ) I make you a BREAKFAST IN BED( low, 11(bin) ), just say thank you instead of: ' hOW( cap ) did you get in my, house ( cap, 2 )? ' I have to pack 101 ( bin ) outfits. Packed 1a ( hex ),,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. harold wilson ( cap, 2 )(up, 2): 'I am a optimist, but a optimist who carries a raincoat... '",
		},
		{
			"Parentheses",
			"",
			FixParentheses,
			"'if (cap) I make you a BREAKFAST IN BED (low, 11 (bin)) , just say thank you instead of: ' hOW (cap) did you get in my, house (cap, 2) ? ' I have to pack 101 (bin) outfits. Packed 1a (hex) ,,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. harold wilson (cap, 2) (up, 2) : 'I am a optimist, but a optimist who carries a raincoat... '",
		},
		{
			"Instances",
			"",
			ApplyInstances,
			"'If I make you a BREAKFAST IN BED (low, 11 (bin)) , just say thank you instead of: ' How did you get in My, House ? ' I have to pack 5 outfits. Packed 26 ,,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. HAROLD WILSON : 'I am a optimist, but a optimist who carries a raincoat... '",
		},
		{
			"Punctuation",
			"",
			FixPunctuation,
			"'If I make you a BREAKFAST IN BED (low, 11 (bin)), just say thank you instead of: ' How did you get in My, House? ' I have to pack 5 outfits. Packed 26,,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. HAROLD WILSON: 'I am a optimist, but a optimist who carries a raincoat... '",
		},
		{
			"Apostrophic",
			"",
			FixApostrophic,
			"'If I make you a BREAKFAST IN BED (low, 11 (bin)), just say thank you instead of:' How did you get in My, House? 'I have to pack 5 outfits. Packed 26,,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. HAROLD WILSON:' I am a optimist, but a optimist who carries a raincoat... '",
		},
		{
			"AToAn",
			"",
			FixAn,
			"'If I make you a BREAKFAST IN BED (low, 11 (bin)), just say thank you instead of:' How did you get in My, House? 'I have to pack 5 outfits. Packed 26,,,,, just to be sure. Don't be sad, because sad backwards is das. And das not good. HAROLD WILSON:' I am an optimist, but an optimist who carries a raincoat... '",
		},
	}

	for i, test := range tests {
		test.input = test.fix(test.input)
		if test.input != test.expect {
			t.Errorf(test.name+"💥\nResult -> %v\nExpected -> %v\n\n", test.input, test.expect)
			break
		} else {
			fmt.Println(test.name + ": ✅\n")
			if i < len(tests)-1 {
				tests[i+1].input = test.fix(test.input)
			}
		}

	}
}

