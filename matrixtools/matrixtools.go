package matrixtools

import "matrices/arraytools"

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	// TODO
	result := m[i]

	return result
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	col := make([]float64, len(m))
	// TODO
	for i := range m {
		col[i] = m[i][j]
	}
	return col
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	// TODO
	arraytools.Add(m[i], m[j]) //arraytool eingefügt und dadurch enzelne spalten der Matritzen aufaddieren

}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	// TODO
	arraytools.ScalarMult(m[i], factor)
}

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {
	transposed := make([][]float64, len(m[0]))
	// TODO //Transposed = m[0]
	for i := range m[0] { //m fragt ab wie viele zeilen hast du und [0] fragt ab, wie viele Spalten wir haben
		transposed[i] = GetCol(m, i)
	}
	return transposed
}
