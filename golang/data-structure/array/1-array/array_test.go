package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("Should return index of element", func(t *testing.T) {
		//arrange
		arr := []string{"cat", "dog", "bird", "fish"}
		var x string = "fish"

		want := 3

		// act
		got, err := Find(arr, x)

		if err != nil {
			fmt.Println("err nil", err)
		}

		assert.Equal(t, want, got)
	})

	t.Run("Should not found inde of element", func(t *testing.T) {
		arr := []string{"cat", "dog", "bird", "fish"}

		x := "elephent"

		want := 8
		wantErr := "not found"

		got, err := Find(arr, x)

		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, err)
	})
}
