package main

import "fmt"
import "math"

/*
Dieses Programm berechnet die Beleuchtungsstärke einer punktförmigen
Lichtquelle.  Es wird angenommen, dass die Lichtstärke (engl. illuminance) der
Lichtquelle über den ganzen Abstrahlwinkel (in Sterandiant) konstant bleibt.
Daher ist in diesem Fall: Lichtstärke = Lichtstrom(engl.
luminous_flux)/Raumwinkel(engl. solid_angle)
*/

func deg2rad(openingAngle float64) float64 {
	return (openingAngle * math.Pi) / 180
}

func rad2sr(openingAngle float64) float64 {
	return 2 * math.Pi * (1 - math.Cos(openingAngle/2))
}

func sr2rad(solidAngle float64) float64 {
	return 2 * math.Acos(-solidAngle/(2*math.Pi)+1)
}

// falls der Lichtkegel kein Kegel, sondern eine Pyramide ist.
// Beispiele: Der Spiegel eines Overhead-Projektors,
// eine rechteckige Straßenlaterne*/
func rad2srRect(angleX, angleY float64) float64 {
	return 4 * math.Asin(math.Sin(angleX)*math.Sin(angleY))
}

func calcIlluminance(openingAngle, luminousFlux, distance, angleOfIncline float64) float64 {
	solid_angle := rad2sr(openingAngle)
	return (luminousFlux / (math.Pow(distance, 2) * solid_angle)) * math.Cos(angleOfIncline)
}

func main() {
	fmt.Println("Beleuchtungsstärke:", calcIlluminance(sr2rad(1.67*math.Pi), 3545, 1.67, math.Atan(1.15/1.67)), "Lux")
}
