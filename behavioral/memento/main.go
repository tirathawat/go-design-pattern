package main

import (
	"design-pattern/behavioral/memento/editor"
	"fmt"
)

func main() {
	myEditor := editor.Editor{
		Title:     "Hello World",
		Content:   "Hello World",
		FontSize:  12,
		FontColor: "red",
	}
	fmt.Println(myEditor)

	state := myEditor.Save()
	editor.StateManager.Push(state)

	myEditor.Title = "Hello Go"
	myEditor.Content = "Hello Go"
	state = myEditor.Save()
	editor.StateManager.Push(state)
	fmt.Println(myEditor)

	myEditor.Restore(editor.StateManager.Undo())
	fmt.Println(myEditor)

	myEditor.Restore(editor.StateManager.Redo())
	fmt.Println(myEditor)
}
