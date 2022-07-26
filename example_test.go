package uniseg_test

import (
	"fmt"

	"github.com/rivo/uniseg"
)

func ExampleGraphemes() {
	gr := uniseg.NewGraphemes("👍🏼!")
	for gr.Next() {
		fmt.Printf("%x ", gr.Runes())
	}
	// Output: [1f44d 1f3fc] [21]
}

func ExampleGraphemeClusterCount() {
	n := uniseg.GraphemeClusterCount("🇩🇪🏳️‍🌈")
	fmt.Println(n)
	// Output: 2
}

func ExampleFirstGraphemeCluster() {
	b := []byte("🇩🇪🏳️‍🌈")
	state := -1
	var c []byte
	for len(b) > 0 {
		c, b, state = uniseg.FirstGraphemeCluster(b, state)
		fmt.Println(string(c))
	}
	// Output: 🇩🇪
	//🏳️‍🌈
}

func ExampleFirstGraphemeClusterInString() {
	str := "🇩🇪🏳️‍🌈"
	state := -1
	var c string
	for len(str) > 0 {
		c, str, state = uniseg.FirstGraphemeClusterInString(str, state)
		fmt.Println(c)
	}
	// Output: 🇩🇪
	//🏳️‍🌈
}

func ExampleFirstWord() {
	b := []byte("Hello, world!")
	state := -1
	var c []byte
	for len(b) > 0 {
		c, b, state = uniseg.FirstWord(b, state)
		fmt.Printf("(%s)\n", string(c))
	}
	// Output: (Hello)
	//(,)
	//( )
	//(world)
	//(!)
}

func ExampleFirstWordInString() {
	str := "Hello, world!"
	state := -1
	var c string
	for len(str) > 0 {
		c, str, state = uniseg.FirstWordInString(str, state)
		fmt.Printf("(%s)\n", c)
	}
	// Output: (Hello)
	//(,)
	//( )
	//(world)
	//(!)
}

func ExampleFirstSentence() {
	b := []byte("This is sentence 1.0. And this is sentence two.")
	state := -1
	var c []byte
	for len(b) > 0 {
		c, b, state = uniseg.FirstSentence(b, state)
		fmt.Printf("(%s)\n", string(c))
	}
	// Output: (This is sentence 1.0. )
	//(And this is sentence two.)
}

func ExampleFirstSentenceInString() {
	str := "This is sentence 1.0. And this is sentence two."
	state := -1
	var c string
	for len(str) > 0 {
		c, str, state = uniseg.FirstSentenceInString(str, state)
		fmt.Printf("(%s)\n", c)
	}
	// Output: (This is sentence 1.0. )
	//(And this is sentence two.)
}

func ExampleFirstLineSegment() {
	b := []byte("First line.\nSecond line.")
	state := -1
	var (
		c         []byte
		mustBreak bool
	)
	for len(b) > 0 {
		c, b, mustBreak, state = uniseg.FirstLineSegment(b, state)
		fmt.Printf("(%s)", string(c))
		if mustBreak {
			fmt.Print("!")
		}
	}
	// Output: (First )(line.
	//)!(Second )(line.)!
}

func ExampleFirstLineSegmentInString() {
	str := "First line.\nSecond line."
	state := -1
	var (
		c         string
		mustBreak bool
	)
	for len(str) > 0 {
		c, str, mustBreak, state = uniseg.FirstLineSegmentInString(str, state)
		fmt.Printf("(%s)", c)
		if mustBreak {
			fmt.Println(" < must break")
		} else {
			fmt.Println(" < may break")
		}
	}
	// Output: (First ) < may break
	//(line.
	//) < must break
	//(Second ) < may break
	//(line.) < must break
}

func ExampleStep_graphemes() {
	b := []byte("🇩🇪🏳️‍🌈")
	state := -1
	var c []byte
	for len(b) > 0 {
		c, b, _, state = uniseg.Step(b, state)
		fmt.Println(string(c))
	}
	// Output: 🇩🇪
	//🏳️‍🌈
}

func ExampleStepString_graphemes() {
	str := "🇩🇪🏳️‍🌈"
	state := -1
	var c string
	for len(str) > 0 {
		c, str, _, state = uniseg.StepString(str, state)
		fmt.Println(c)
	}
	// Output: 🇩🇪
	//🏳️‍🌈
}

func ExampleStep_word() {
	b := []byte("Hello, world!")
	state := -1
	var (
		c          []byte
		boundaries int
	)
	for len(b) > 0 {
		c, b, boundaries, state = uniseg.Step(b, state)
		fmt.Print(string(c))
		if boundaries&uniseg.MaskWord != 0 {
			fmt.Print("|")
		}
	}
	// Output: Hello|,| |world|!|
}

func ExampleStepString_word() {
	str := "Hello, world!"
	state := -1
	var (
		c          string
		boundaries int
	)
	for len(str) > 0 {
		c, str, boundaries, state = uniseg.StepString(str, state)
		fmt.Print(c)
		if boundaries&uniseg.MaskWord != 0 {
			fmt.Print("|")
		}
	}
	// Output: Hello|,| |world|!|
}

func ExampleStep_sentence() {
	b := []byte("This is sentence 1.0. And this is sentence two.")
	state := -1
	var (
		c          []byte
		boundaries int
	)
	for len(b) > 0 {
		c, b, boundaries, state = uniseg.Step(b, state)
		fmt.Print(string(c))
		if boundaries&uniseg.MaskSentence != 0 {
			fmt.Print("|")
		}
	}
	// Output: This is sentence 1.0. |And this is sentence two.|
}

func ExampleStepString_sentence() {
	str := "This is sentence 1.0. And this is sentence two."
	state := -1
	var (
		c          string
		boundaries int
	)
	for len(str) > 0 {
		c, str, boundaries, state = uniseg.StepString(str, state)
		fmt.Print(c)
		if boundaries&uniseg.MaskSentence != 0 {
			fmt.Print("|")
		}
	}
	// Output: This is sentence 1.0. |And this is sentence two.|
}

func ExampleStep_lineBreaking() {
	b := []byte("First line.\nSecond line.")
	state := -1
	var (
		c          []byte
		boundaries int
	)
	for len(b) > 0 {
		c, b, boundaries, state = uniseg.Step(b, state)
		fmt.Print(string(c))
		if boundaries&uniseg.MaskLine == uniseg.LineCanBreak {
			fmt.Print("|")
		} else if boundaries&uniseg.MaskLine == uniseg.LineMustBreak {
			fmt.Print("‖")
		}
	}
	// Output: First |line.
	//‖Second |line.‖
}

func ExampleStepString_lineBreaking() {
	str := "First line.\nSecond line."
	state := -1
	var (
		c          string
		boundaries int
	)
	for len(str) > 0 {
		c, str, boundaries, state = uniseg.StepString(str, state)
		fmt.Print(c)
		if boundaries&uniseg.MaskLine == uniseg.LineCanBreak {
			fmt.Print("|")
		} else if boundaries&uniseg.MaskLine == uniseg.LineMustBreak {
			fmt.Print("‖")
		}
	}
	// Output: First |line.
	//‖Second |line.‖
}
