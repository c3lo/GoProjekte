package main

import "fmt"
import "math"

/*Dieses Programm berechnet die Beleuchtungsstärke einer punktförmigen Lichtquelle.
Es wird angenommen, dass die Lichtstärke (engl. illuminance)
der Lichtquelle über den ganzen Abstrahlwinkel (in Sterandiant) konstant bleibt.
Daher ist in diesem Fall:
Lichtstärke = Lichtstrom(engl. luminousFlux)/Raumwinkel(engl. solidAngle)
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

func rad2srRect(angleX, angleY float64) float64 {
	/* falls der Lichtkegel kein Kegel, sondern eine Pyramide ist.
	Beispiele: Der Spiegel eines Overhead-Projektors,
	eine rechteckige Straßenlaterne*/
	return 4 * math.Asin(math.Sin(angleX)*math.Sin(angleY))
}

type IlluminanceParameters struct {
	OpeningAngle   float64
	LuminousFlux   float64
	Distance       float64
	AngleOfIncline float64
}

func (self *IlluminanceParameters) CalcIlluminance() float64 {
	solidAngle := rad2sr(self.OpeningAngle)
	return (self.LuminousFlux / (math.Pow(self.Distance, 2) * solidAngle)) * math.Cos(self.AngleOfIncline)
}

func main() {
	params := IlluminanceParameters{
		OpeningAngle:   sr2rad(1.67 * math.Pi),
		LuminousFlux:   3545,
		Distance:       1.67,
		AngleOfIncline: math.Atan(1.15 / 1.67),
	}

	fmt.Printf("Beleuchtungsstärke: %f Lux\n", params.CalcIlluminance())
}
