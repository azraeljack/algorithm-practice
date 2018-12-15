package sort

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common"
	"github.com/azraeljack/algorithm-practice/common/testdata"
)

func TestBubbleSortNilSlice(t *testing.T) {
	// just make sure we don't crash no nil
	BubbleSort(nil)
}

func TestBubbleSortEmptySlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("emptySlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("emptySlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortSingleSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("singleSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("singleSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortDoubleSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("doubleSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("doubleSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortTripleSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("tripleSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("tripleSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortEqualSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("equalSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("equalSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortSortedSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("sortedSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("sortedSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortReserveSortedSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("reverseSortedSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("reverseSortedSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortRandomizedSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("randomizedSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("randomizedSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}

func TestBubbleSortLongRandomizedSlice(t *testing.T) {
	testData, err := testdata.GetSliceCopy("longRandomizedSlice")
	if err != nil {
		t.Errorf("failed to get test data, err: %v", err)
	}
	BubbleSort(testData)
	if !common.IsIntArraySorted(testData) {
		expected, _ := testdata.GetSliceCopy("longRandomizedSlice")
		t.Errorf("got %v, expected %v", testData, common.SortIntArray(expected))
	}
}
