package piscine

func FindNextPrime(nb int) int {
	for counter := 0; counter < nb; counter++ {
		for modu := 2; modu <= (nb + counter); modu++ {
			if modu > (nb+counter)/2 {
				return nb + counter
			}
			if (nb+counter)%modu == 0 {
				break
			}
		}
	}
	return 2
}
