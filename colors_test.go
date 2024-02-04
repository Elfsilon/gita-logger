package gita

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTextBuild(t *testing.T) {
	text := Text("Test")

	res := text.Build(&TextStyle{
		Style: Bold,
	})
	require.Equal(t, res, fmt.Sprintf("%vTest%v", Bold, Reset))

	res = text.Build(&TextStyle{
		Style:   Bold,
		Color:   Red,
		BGColor: GreenBG,
	})
	require.Equal(t, res, fmt.Sprintf("%v%v%vTest%v", Red, GreenBG, Bold, Reset))
}

func TestTextStyleApply(t *testing.T) {
	var text string = "Test"

	style := TextStyle{
		Style: Bold,
	}
	style.apply(&text)
	require.Equal(t, text, fmt.Sprintf("%vTest%v", Bold, Reset))

	text = "Test"
	style = TextStyle{
		Style:   Bold,
		Color:   Red,
		BGColor: GreenBG,
	}
	style.apply(&text)
	require.Equal(t, text, fmt.Sprintf("%v%v%vTest%v", Red, GreenBG, Bold, Reset))

	text = "Test"
	var emptyStyle *TextStyle = nil
	emptyStyle.apply(&text)
	require.Equal(t, text, "Test")
}
