package medium

func RearrangeArray(nums []int) []int {
    var positives []int
    var negatives []int

    for _, num := range nums {
        if num > 0 {
            positives = append(positives, num)
        } else {
            negatives = append(negatives, num)
        }
    }

    var result []int

    for i, pos := range positives {
        result = append(result, pos)
        result = append(result, negatives[i])
    }

    return result
}
