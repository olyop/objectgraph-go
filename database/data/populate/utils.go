package populate

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strings"
	"time"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func convertSetToArr[T any](m map[string]T) []any {
	arr := make([]any, 0)

	for k := range m {
		arr = append(arr, k)
	}

	return arr
}

func convertToInterfaceSlice[T any](names []T) []any {
	interfaces := make([]any, len(names))

	for i, v := range names {
		interfaces[i] = v
	}

	return interfaces
}

// batchProducts splits the array into chunks of 32767
// (half the PostgreSQL limit for the number of parameters in
// a query since we are passing in two parameters for each product)
func batchProducts(batch []product) [][]product {
	var batches [][]product

	for i := 0; i < len(batch); i += 32767 {
		end := i + 32767

		if end > len(batch) {
			end = len(batch)
		}

		batches = append(batches, batch[i:end])
	}

	return batches
}

func generateRandomMobileNumber() string {
	return fmt.Sprintf("04%d", rand.IntN(1000000000))
}

func generateRandomDOB() int64 {
	year := rand.IntN(40) + 1970
	month := rand.IntN(11) + 1
	day := rand.IntN(27) + 1

	dob := fmt.Sprintf("%d-%02d-%02d", year, month, day)

	parsedDOB, err := time.Parse(time.DateOnly, dob)
	if err != nil {
		log.Default().Fatal(err)
	}

	if parsedDOB.UnixMilli() < 0 {
		return generateRandomDOB()
	}

	return parsedDOB.UnixMilli()
}

func generateUserNameFromEmailAddress(emailAddress string) string {
	indexOfAt := strings.Index(emailAddress, "@")

	return emailAddress[:indexOfAt]
}
