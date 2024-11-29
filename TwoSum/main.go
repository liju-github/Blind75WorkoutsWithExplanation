package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // Value to index map

	// Iterate through the nums array.
	// For each element 'num', we calculate its complement (target - num).
	// We then check if the complement already exists in the map.
	// If it does, we return the index of the complement and the current index 'i' as the solution.
	// If not, we store the current number and its index in the map for future reference.
	// This ensures that we are checking for each pair's complement in one pass.
	// If no pair sums up to the target, the function returns nil indicating no solution was found.
	for i, num := range nums {
		complement := target - num
		if idx, found := m[complement]; found {
			return []int{idx, i} // Return the indices of the two numbers whose sum equals the target
		}
		m[num] = i // Store the index of the current number
	}
	return nil // No solution found
}

func main() {
	nums := []int{2, 7, 11, 15} // Example input array
	target := 9                 // Example target sum
	result := twoSum(nums, target) // Call the twoSum function
	fmt.Println(result) // Output: [0, 1] because nums[0] + nums[1] = 2 + 7 = 9
}
