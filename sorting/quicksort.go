package sorting

type QuickSort struct {
	Data []int
}

func (qck *QuickSort) swap(source, destination *int) {
	*source, *destination = *destination, *source
}

func (qck *QuickSort) partition(lowIndex, highIndex int) int {

	pivot := qck.Data[highIndex]
	controllerIndex := lowIndex - 1

	for index := lowIndex; index <= highIndex-1; index++ {
		if qck.Data[index] <= pivot {
			controllerIndex++
			qck.swap(&qck.Data[controllerIndex], &qck.Data[index])
		}
	}

	qck.swap(&qck.Data[controllerIndex+1], &qck.Data[highIndex])

	return controllerIndex + 1
}

func (qck *QuickSort) Sort(left, right int) {

	if left < right {

		indexPivot := qck.partition(left, right)

		qck.Sort(left, indexPivot-1)

		qck.Sort(indexPivot+1, right)

	}
}
