package IMEKeyboard

import (
	"fmt"
	"testing"
)

func TestKeyboardPos_ChooseWord(t *testing.T) {
	model, err := KBPSharedInstance(SogouBrand, "GooglePixel2XL", YamlFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	coordinate := model.KeyboardLastWord()
	fmt.Println(coordinate.X)
	fmt.Println(coordinate.Y)
}
