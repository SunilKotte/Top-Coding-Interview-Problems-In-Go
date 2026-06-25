// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0. You must do it in place.

// Example 1

// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]

// Output: [[1,0,1],[0,0,0],[1,0,1]]

// Explanation:

// Element at position (1,1) is 0, so set entire row 1 and column 1 to 0.

// Example 2

// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]

// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Explanation:

// There are two zeroes: (0,0) and (0,3).

// Row 0 → all elements become 0
// Column 0 and column 3 → all elements become 0




func setZeroes(matrix [][]int) {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    n, m := len(matrix), len(matrix[0])

    // Step 1: Check if first row and first column contain zero
    firstRowZero := false
    for j := 0; j < m; j++ {
        if matrix[0][j] == 0 {
            firstRowZero = true
            break
        }
    }

    firstColZero := false
    for i := 0; i < n; i++ {
        if matrix[i][0] == 0 {
            firstColZero = true
            break
        }
    }

    // Step 2: Use first row/col as markers for the rest of the matrix
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    // Step 3: Zero out interior cells based on markers
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    // Step 4: Zero out the first row if needed
    if firstRowZero {
        for j := 0; j < m; j++ {
            matrix[0][j] = 0
        }
    }

    // Step 5: Zero out the first column if needed
    if firstColZero {
        for i := 0; i < n; i++ {
            matrix[i][0] = 0
        }
    }
}