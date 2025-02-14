package medium

type ProductOfNumbers struct {
	numbers []int
}

func Constructor() ProductOfNumbers {
	var emptyList []int
	return ProductOfNumbers{
		numbers: emptyList,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.numbers = append(this.numbers, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	slice := this.numbers[len(this.numbers)-k:]
	product := 1

	for _, num := range slice {
		product *= num
	}

	return product
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
