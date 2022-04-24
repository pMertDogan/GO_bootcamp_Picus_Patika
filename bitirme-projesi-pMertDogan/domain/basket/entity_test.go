package basket

import (
	"reflect"
	"testing"
)

func TestBaskets_GenerateProductIDTotalQuantityMap(t *testing.T) {
	tests := []struct {
		name string
		r    *Baskets
		want map[int]int
	}{
		{"same ID OK", &Baskets{
			{
				ProductID: 1,
				TotalQuantity: 2,
			},
			{
				ProductID: 1,
				TotalQuantity: 3,
			},
			{
				ProductID: 2,
				TotalQuantity: 4,
			},
		}, map[int]int{
			1: 3,
			2: 4,
		}},
		{"empty basket OK", &Baskets{}, map[int]int{}},
		{"one basket OK", &Baskets{
			{
				ProductID:     1,
				TotalQuantity: 1,
			},
		}, map[int]int{
			1: 1,
		}},
		{"two baskets OK", &Baskets{
			{
				ProductID:     1,
				TotalQuantity: 1,
			},
			{
				ProductID:     2,
				TotalQuantity: 2,
			},
		}, map[int]int{
			1: 1,
			2: 2,
		}},
		{"two baskets with different IDs OK", &Baskets{
			{
				ProductID:     763,
				TotalQuantity: 367,
			},
			{
				ProductID:     230,
				TotalQuantity: 729,
			},
		}, map[int]int{
			763: 367,
			230: 729,
		}},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GenerateProductIDTotalQuantityMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Baskets.GenerateProductIDTotalQuantityMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
