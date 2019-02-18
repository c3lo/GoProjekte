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

func deg2rad(opening_angle float64) float64 {
  return (opening_angle*math.Pi)/180
}

func rad2sr(opening_angle float64) float64 {
  return 2*math.Pi*(1 - math.Cos(opening_angle/2))
}

func sr2rad(solid_angle float64) float64 {
  return 2*math.Acos(-solid_angle/(2*math.Pi) + 1)
}

// falls der Lichtkegel kein Kegel, sondern eine Pyramide ist.
// Beispiele: Der Spiegel eines Overhead-Projektors,
// eine rechteckige Straßenlaterne*/
func rad2srRect(angle_x, angle_y float64) float64{
  return 4*math.Asin(math.Sin(angle_x)*math.Sin(angle_y))
}

func calc_illuminance(opening_angle, luminous_flux, distance, angle_of_incline float64) float64{
  solid_angle := rad2sr(opening_angle)
  return (luminous_flux/(math.Pow(distance, 2)*solid_angle))*math.Cos(angle_of_incline)
}

func main()  {
  fmt.Println("Beleuchtungsstärke:", calc_illuminance(sr2rad(1.67*math.Pi), 3545, 1.67, math.Atan(1.15/1.67)), "Lux")
}
