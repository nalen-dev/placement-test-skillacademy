package deretangka

import "fmt"

func HitungDeretAngka() {
	angka := 3
	selisih := 2

	for i := 1; i < 100; i++ {
		angka += selisih
		selisih += 2
	}

	fmt.Println("Angka ke-100 dalam deret adalah:", angka)
}

func RecursiveRandomNumberGenerator(n int, seed int64) int64 {
	const (
		a int64 = 1103
		c int64 = 1234
		m int64 = 1 << 31
	)

	if n == 0 {
		return seed
	}

	prevRandom := RecursiveRandomNumberGenerator(n-1, seed)

	random := (a*prevRandom + c) % m

	return random
}
