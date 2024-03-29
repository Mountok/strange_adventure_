package readkeyboard

import (

	"github.com/eiannone/keyboard"
)


func ReadKeyBoard() (error,keyboard.Key) {
	if err := keyboard.Open(); err != nil {
        panic(err)
    }
    defer func() {
        _ = keyboard.Close()
    }()
	_, key, err := keyboard.GetKey()
    if err != nil {
        panic(err)
    }
	return nil,key
}