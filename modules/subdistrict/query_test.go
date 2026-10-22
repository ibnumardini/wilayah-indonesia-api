package subdistrict

import "testing"

func TestQueryOptionsValidate(t *testing.T) {
	tests := []struct {
		name    string
		options QueryOptions
		valid   bool
	}{
		{
			name:    "valid options",
			options: QueryOptions{Page: 1, Limit: 10, Sort: "asc"},
			valid:   true,
		},
		{
			name:    "invalid page",
			options: QueryOptions{Page: 0, Limit: 10, Sort: "asc"},
		},
		{
			name:    "invalid limit",
			options: QueryOptions{Page: 1, Limit: 101, Sort: "asc"},
		},
		{
			name:    "invalid sort",
			options: QueryOptions{Page: 1, Limit: 10, Sort: "random"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.options.Validate()
			if (err == nil) != test.valid {
				t.Fatalf("Validate() error = %v, valid = %v", err, test.valid)
			}
		})
	}
}

func TestQueryOptionsCheckSortDefaultsToAscending(t *testing.T) {
	options := QueryOptions{Sort: "random"}

	options.CheckSort()

	if options.Sort != "asc" {
		t.Fatalf("CheckSort() set sort to %q, want asc", options.Sort)
	}
}
