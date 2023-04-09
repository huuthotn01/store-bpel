package schema

type CommonGetStatisticRequest struct {
	Start string
	End   string
}

type FilterGetStatisticRequest struct {
	BranchId []string
	Gender   []int
	Type     []string
	Start    string
	End      string
}
