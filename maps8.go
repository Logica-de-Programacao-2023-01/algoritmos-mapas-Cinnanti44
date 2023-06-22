package main

func calculateExpenseSettlement(expenses map[string]float64) map[string]float64 {
	totalExpenses := 0.0
	for _, expense := range expenses {
		totalExpenses += expense
	}

	averageExpense := totalExpenses / float64(len(expenses))

	settlement := make(map[string]float64)
	for person, expense := range expenses {
		settlement[person] = expense - averageExpense
	}

	return settlement
}
