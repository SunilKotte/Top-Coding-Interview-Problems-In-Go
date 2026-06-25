func pascalTriangleI(r, c int) int {
    // Pascal's triangle is 1‑indexed: row 1 = [1], col 1 = first element.
    // The value is C(r-1, c-1).
    n := r - 1
    k := c - 1

    // If k is out of range, return 0 (or handle as needed).
    if k < 0 || k > n {
        return 0
    }

    // Use the smaller of k and n-k for efficiency.
    if k > n-k {
        k = n - k
    }

    result := 1
    // Compute C(n, k) iteratively: result = product_{i=1}^{k} (n - k + i) / i
    for i := 1; i <= k; i++ {
        result = result * (n - k + i) / i
    }
    return result
}