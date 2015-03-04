package levenshtein

const MAX_DISTANCE = int(^uint(0) >> 1) // Largest possible int

func Sort(items Sequence) {
	nextBest := MAX_DISTANCE
	upperBound := MAX_DISTANCE
	length := items.Length()
	dists := make([]int, length)

	for i := 0; i < length - 2; i++ {
		best := MAX_DISTANCE
		index := -1

		for j := i + 1; j < length; j++ {
			if i == j { continue }
			if dists[j] > upperBound { continue }

			distance := Distance(items.Pair(i, j))
			dists[j] = distance
			if distance < best {
				best, nextBest = distance, best
				index = j
			}
		}

		top := i + 1
		items.Swap(top, index)
		dists[top], dists[index] = dists[index], dists[top]
		if nextBest == MAX_DISTANCE {
			upperBound = MAX_DISTANCE
		} else {
			upperBound = best + nextBest
		}
		nextBest = best
	}
}
