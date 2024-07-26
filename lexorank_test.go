package lexorank

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRank(t *testing.T) {
	tests := []struct {
		prev, next string
		want       string
		wantErr    bool
	}{
		{".", "a", "", true},
		{"a", ",", "", true},
		{"|", "a", "", true},
		{"a", "]", "", true},
		{"b", "a", "", true},
		{"", "", "", false},
		{"001", "001", "001", false},
		{"1", "10", "10", false},
		{"", "0", "0", false},
		{"1", "3", "2", false},
		{"1", "9", "5", false},
		{"10", "20", "1U", false},
		{"RaNDOm1", "TeSTCaSE001", "RaNDOm10000U", false},
		{"491q0VpP5", "Z0sJKz8gs", "G4x8>Bz6T", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("case: %s, %s", tt.prev, tt.next), func(t *testing.T) {
			got, err := Rank(tt.prev, tt.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankN(t *testing.T) {
	tests := []struct {
		prev    string
		next    string
		n       int
		want    []string
		wantErr bool
	}{
		{".", "a", 10, nil, true},
		{"a", ",", 10, nil, true},
		{"|", "a", 10, nil, true},
		{"a", "]", 10, nil, true},
		{"b", "a", 10, nil, true},
		{"0", "z", 5, []string{"U0", "U1", "U2", "U3", "U4"}, false},
		{"0", "z", 20, []string{"U00", "U01", "U02", "U03", "U04", "U05", "U06", "U07", "U08", "U09", "U10", "U11", "U12", "U13", "U14", "U15", "U16", "U17", "U18", "U19"}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("case: %s, %s, %d", tt.prev, tt.want, tt.n), func(t *testing.T) {
			got, err := RankN(tt.prev, tt.next, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("RankN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RankN() got = %v, want %v", got, tt.want)
			}
			for _, item := range got {
				if item < tt.prev {
					t.Errorf("RankN() got item = %v, less than %v", item, tt.prev)
				}
				if item < tt.prev || item > tt.next {
					t.Errorf("RankN() got item = %v, greater than %v", item, tt.next)
				}
			}
		})
	}
}
