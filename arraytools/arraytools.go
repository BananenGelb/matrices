package arraytools

// Mult erwartet ein Array und einen skalaren Faktor.
// Multipliziert jedes Element des Arrays mit dem Faktor.
func ScalarMult(a []float64, factor float64) {
	// TODO
	for i, wert := range a { //zeihen den wert aus der Postition "pos"
		a[i] = wert * factor //multiplizieren mit dem wert der position und setzt es an Position wieder ein
	}

}

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
func Add(a, b []float64) {
	// TODO
	for i := range a {
		a[i] += b[i]
	}
}
