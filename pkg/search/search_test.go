package search

import (
	"context"
	"reflect"
	"sort"
	"testing"
)

func BenchmarkAll(b *testing.B) {
	want := [][]Result{
		[]Result{
			{
				Phrase:  "mama ela kashy",
				Line:    "tgvrtgvwmama ela kashy",
				LineNum: 5,
				ColNum:  9,
			},
			{
				Phrase:  "mama ela kashy",
				Line:    "qwetgvrtgvwmama ela kashy",
				LineNum: 6,
				ColNum:  12,
			},
			{
				Phrase:  "mama ela kashy",
				Line:    "mama ela kashy",
				LineNum: 7,
				ColNum:  1,
			},
		},
		[]Result{
			{
				Phrase:  "mama ela kashy",
				Line:    "qwemama ela kashy",
				LineNum: 8,
				ColNum:  4,
			},
			{
				Phrase:  "mama ela kashy",
				Line:    "aerg werg werg werg wergmama ela kashy",
				LineNum: 8,
				ColNum:  25,
			},
			{
				Phrase:  "mama ela kashy",
				Line:    "aerg qwewerg werg werg wergmama ela kashy",
				LineNum: 9,
				ColNum:  28,
			},
		}}

	ctx := context.Background()
	files := []string{}
	files = append(files, "data/data", "data/data1")
	result := [][]Result{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := All(ctx, "mama ela kashy", files)
		for va := range ch {
			sort.Slice(va[:], func(i, j int) bool {
				return va[i].LineNum < va[j].LineNum
			})
			result = append(result, va)
		}

		b.StopTimer()
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result \ngot: %v; \nwant: %v", result, want)
		}

		b.StartTimer()
	}
}
