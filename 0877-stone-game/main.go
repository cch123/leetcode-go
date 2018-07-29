package main

type record struct {
	idx  int
	sum  int
	left int
}

func stoneGame(piles []int) bool {
	// init
	var f = make([][]record, len(piles))
	for i := 0; i < len(piles); i++ {
		f[i] = make([]record, len(piles))
		f[i][i] = record{
			idx:  i,
			sum:  piles[i],
			left: 0,
		}
	}

	for gap := 1; gap <= len(piles)-1; gap++ {
		// i left index
		// j right index
		for i := 0; i < len(piles); i++ {
			j := i + gap
			if j >= len(piles) {
				break
			}

			if f[i+1][j].left+piles[i] < f[i][j-1].left+piles[j] {
				// 说明应该取右边的
				f[i][j] = record{idx: j,
					sum:  f[i][j-1].left + piles[j],
					left: f[i][j-1].sum,
				}
			} else {
				f[i][j] = record{idx: i,
					sum:  f[i+1][j].left + piles[i],
					left: f[i+1][j].sum,
				}
			}
		}
	}

	resultRec := f[0][len(piles)-1]
	if resultRec.sum > resultRec.left {
		return true
	}

	return false
}
