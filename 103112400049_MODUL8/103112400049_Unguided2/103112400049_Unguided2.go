package main

import "fmt"

func main() {
	var nomorPilihan int
	suaraMasuk := 0
	suaraSah := 0
	hitungSuaraCalon := make([]int, 21)

	for {
		_, err := fmt.Scan(&nomorPilihan)
		if err != nil || nomorPilihan == 0 {
			break
		}
		suaraMasuk++
		if nomorPilihan >= 1 && nomorPilihan <= 20 {
			suaraSah++
			hitungSuaraCalon[nomorPilihan]++
		}
	}
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	ketuaID := 0
	ketuaSuara := -1
	wakilID := 0
	wakilSuara := -1
	//103112400049 Hisyam Nurdiatmoko
	for id := 1; id <= 20; id++ {
		suaraCalonIni := hitungSuaraCalon[id]
		if suaraCalonIni == 0 {
			continue
		}
		if suaraCalonIni > ketuaSuara {
			wakilID = ketuaID
			wakilSuara = ketuaSuara

			ketuaID = id
			ketuaSuara = suaraCalonIni
		} else if suaraCalonIni == ketuaSuara {
			if id < ketuaID {
				wakilID = ketuaID
				wakilSuara = ketuaSuara

				ketuaID = id
			} else {
				if suaraCalonIni > wakilSuara {
					wakilID = id
					wakilSuara = suaraCalonIni
				} else if suaraCalonIni == wakilSuara {
					if id < wakilID {
						wakilID = id
					}
				}
			}
		} else {
			if suaraCalonIni > wakilSuara {
				wakilID = id
				wakilSuara = suaraCalonIni
			} else if suaraCalonIni == wakilSuara {
				if id < wakilID {
					wakilID = id
				}
			}
		}
	}
	fmt.Printf("Ketua RT: %d\n", ketuaID)
	fmt.Printf("Wakil ketua: %d\n", wakilID)
}
