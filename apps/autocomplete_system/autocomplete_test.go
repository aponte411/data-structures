package autocomplete_system

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAutoCompleteSystem(t *testing.T) {
	fmt.Println("Testing")
	// History of usage
	sentences := []string{
		"machine learning",
		"metadata",
		"model serving",
		"metaflow",
		"mlflow",
		"features",
		"predictions",
	}
	times := []int{20, 15, 14, 13, 13, 10, 5}
	user_input := "m"
	// Top 3 words for user input
	expected := []string{
		"machine learning",
		"metadata",
		"model serving",
	}
	auto := InitAutoCompleteSystem(sentences, times)
	result := auto.Input(user_input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
