package populate

func convertSetToArr[T any](m map[string]T) []interface{} {
	arr := make([]interface{}, 0)

	for k := range m {
		arr = append(arr, k)
	}

	return arr
}

func convertToInterfaceSlice[T any](names []T) []interface{} {
	interfaces := make([]interface{}, len(names))

	for i, v := range names {
		interfaces[i] = v
	}

	return interfaces
}

// batchProducts splits the array into chunks of 32767
// (half the PostgreSQL limit for the number of parameters in
// a query since we are passing in two parameters for each product)
func batchProducts(batch []Product) [][]Product {
	var batches [][]Product

	for i := 0; i < len(batch); i += 32767 {
		end := i + 32767

		if end > len(batch) {
			end = len(batch)
		}

		batches = append(batches, batch[i:end])
	}

	return batches
}
