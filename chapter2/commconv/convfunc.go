package commconv

func PoundToKg(lb Pound) Kilogram {
	return Kilogram(lb / 2.20462)
}

func KgToPound(kg Kilogram) Pound {
	return Pound(kg * 2.20462)
}

func InchToCenti(in Inches) Centimeter {
	return Centimeter(in * 2.54)
}

func CentiToInch(cm Centimeter) Inches {
	return Inches(cm / 2.54)
}
