package word

import (
	"fmt"

	"github.com/tjarratt/babble"
)

func GenerateWord() {
	fmt.Println("GenerateWord called")
	fmt.Println(babble.NewBabbler().Babble())
}
