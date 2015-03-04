package levenshtein

func Sort(items Sequence) {
	for i := 0; i < items.Length() - 2; i++ {
		bDistance := int(^uint(0) >> 1) // largest int
		bIndex := -1

		for j := i + 1; j < items.Length(); j++ {
			if i == j { continue }

			distance := Distance(items.Pair(i, j))
			if distance < bDistance {
				bDistance = distance
				bIndex = j
			}
		}

		items.Swap(i + 1, bIndex)
	}
}
