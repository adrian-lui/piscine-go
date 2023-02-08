package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return PrintNbrBase2(AtoiBase(nbr, baseFrom), baseTo)
}
