package main
import "fmt"

func waktu(n int, jam, menit, detik *int){
	*jam = n/3600
	*menit = (n%3600)/60
	*detik = n%60
	fmt.Print(*jam, *menit, *detik)
}

func main(){
	var masukan, j, m, d int
	fmt.Scan(&masukan)
	waktu(masukan, &j, &m, &d)
}
