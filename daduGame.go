package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("\n\n\n\n")
	fmt.Print("Selamat datang di permainan dadu \n")
	fmt.Print("Anda akan berperan sebagai wasit sekaligus penonton di game ini \n")
	fmt.Print("Tugas anda sebagai wasit adalah untuk memulai pertandingan di setiap sesinya \n")
	fmt.Print("Setiap akhir sesi, anda akan diperlihatkan kondisi pemain. \n")
	fmt.Print("permainan berakhir ketika hanya tesisa satu orang atau kurang yang memiliki dadu \n")
	fmt.Print("permainan dengan point terbanyak adalah juaranya \n \n \n")

	sumPlayer := ""
	sumPlayerInt := 0
	sumDice := ""
	sumDiceInt := 0
	step := 0
	var err error

	for time := 0; time < 100; time++ {
		if step == 0 {
			step = 1
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Tolong masukan jumlah pemain : ")
			sumPlayer, _ = reader.ReadString('\n')
			sumPlayer = strings.TrimSuffix(sumPlayer, "\r\n")
			sumPlayerInt, err = strconv.Atoi(sumPlayer)
			if err != nil {
				fmt.Print("Anda memasukan data bukan angka \n")
				step = 0
				continue
			}
			fmt.Println("menyimpan data jumlah pemain " + sumPlayer + " . . . .\n \n")
		}

		if step == 1 {
			step = 2
			fmt.Print("Tolong masukan jumlah dadu : ")
			reader := bufio.NewReader(os.Stdin)
			sumDice, _ = reader.ReadString('\n')
			sumDice = strings.TrimSuffix(sumDice, "\r\n")
			sumDiceInt, err = strconv.Atoi(sumDice)
			if err != nil {
				fmt.Print("Anda memasukan data bukan angka \n")
				step = 1
				continue
			}
			fmt.Println("menyimpan data jumlah dadu " + sumDice + " . . . .\n\n")
		}

		if step == 2 {
			step = 3
			fmt.Print("Apakah benar data jumlah pemain = ")
			fmt.Print(sumPlayer)
			fmt.Print(" dan jumlah dadu = " + sumDice + "\n")
			fmt.Print("Klik Y untuk benar dan memulai permainan atau N untuk salah dan mengisi data kembali atau X untuk keluar dan mengakhiri game\n")
			fmt.Print("Jawaban ")
			reader := bufio.NewReader(os.Stdin)
			answered, _ := reader.ReadString('\n')
			answered = strings.TrimSuffix(answered, "\r\n")
			if answered == "Y" {
				fmt.Print("Permainan Dimulai\n")
				permainan_dadu(sumDiceInt, sumPlayerInt)
			}
			if answered == "N" {
				fmt.Print("Mengulang proses pengisian data\n")
				step = 0
				continue
			}
			if answered == "X" {
				fmt.Print("Anda Menutup Permainan, Terimakasih telah bermain\n")
				break
			}
		}
	}
}

func permainan_dadu(sumDice int, sumPlayer int) {
	var pointPlayer []int
	var dicePlayer []int
	var addDice []int

	fmt.Print("Permainan akan segera dimulai!\n")
	fmt.Print("Total pemain adalah ")
	fmt.Print(sumPlayer)
	fmt.Print(", dengan masing masing dadu ")
	fmt.Print(sumDice)
	fmt.Print(" buah.\n")

	for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
		dicePlayer = append(dicePlayer, sumDice)
		pointPlayer = append(pointPlayer, 0)
		addDice = append(addDice, 0)
		fmt.Print("Pemain ke ")
		fmt.Print(idPlayer)
		fmt.Print(" memiliki ")
		fmt.Print(dicePlayer[idPlayer-1])
		fmt.Print(" dadu dan ")
		fmt.Print(pointPlayer[idPlayer-1])
		fmt.Println(" point ")
	}

	for time := 0; time < 100; time++ {
		playerWinner := 0
		velueWinner := 0
		endGame := 0
		fmt.Print("apakah anda ingin memulai sesi ")
		fmt.Print(time + 1)
		fmt.Print(" sekarang ? \n")
		fmt.Print("ketik Y untuk memulai, ketik N untuk membatalkan permainan dan melihat pemenang berdasarkan data terakhir : ")
		reader := bufio.NewReader(os.Stdin)
		answered, _ := reader.ReadString('\n')
		answered = strings.TrimSuffix(answered, "\r\n")

		if answered == "Y" {
			fmt.Print("Sesi ke ")
			fmt.Print(time + 1)
			fmt.Print(" Dimulai\n")
			for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
				if dicePlayer[idPlayer-1] > 0 {
					allVelueDice := ""
					fmt.Print("Pemain ke ")
					fmt.Print(idPlayer)
					fmt.Print(" melempar semua dadu dan memperoleh ")
					for rollDice := 0; rollDice < dicePlayer[idPlayer-1]; rollDice++ {
						velueDice := (rand.Intn(6) + 1)
						if allVelueDice == "" {
							allVelueDice = strconv.Itoa(velueDice)
						} else {
							allVelueDice = allVelueDice + ", " + strconv.Itoa(velueDice)
						}
						if velueDice == 6 {
							addDice[idPlayer-1] = addDice[idPlayer-1] - 1
							pointPlayer[idPlayer-1] = pointPlayer[idPlayer-1] + 1
						}
						if velueDice == 1 {
							plusDice := 1
							addDice[idPlayer-1] = addDice[idPlayer-1] - 1
							for a := idPlayer; a < sumPlayer; a++ {
								if dicePlayer[a] > 0 {
									addDice[a] = addDice[a] + plusDice
									plusDice = 0
								}
							}
							for b := 0; b < idPlayer-1; b++ {
								if dicePlayer[b] > 0 {
									addDice[b] = addDice[b] + plusDice
									plusDice = 0
								}
							}
							fmt.Print("Data ")
							fmt.Println(addDice)
						}
					}
					fmt.Println(allVelueDice)
					fmt.Println()
				}
				continue
			}

			endGame = 0
			for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
				dicePlayer[idPlayer-1] = dicePlayer[idPlayer-1] + addDice[idPlayer-1]
				addDice[idPlayer-1] = 0
				if velueWinner < pointPlayer[idPlayer-1] {
					playerWinner = idPlayer
					velueWinner = pointPlayer[idPlayer-1]
				}
				fmt.Print("Pemain ke ")
				fmt.Print(idPlayer)
				fmt.Print(" memiliki ")
				fmt.Print(dicePlayer[idPlayer-1])
				fmt.Print(" dadu dan ")
				fmt.Print(pointPlayer[idPlayer-1])
				fmt.Print(" point  \n")
				if dicePlayer[idPlayer-1] == 0 {
					endGame = endGame + 1
				}
			}

			if endGame == sumPlayer-1 {
				fmt.Print("Pemenang adalah Pemain ke ")
				fmt.Print(playerWinner)
				fmt.Print(" dengan ")
				fmt.Print(velueWinner)
				fmt.Print(" point  \n")
				return
			} else if playerWinner > 0 {
				fmt.Print("pemilik point tertinggi adalah Pemain ke ")
				fmt.Print(playerWinner)
				fmt.Print(" dengan ")
				fmt.Print(velueWinner)
				fmt.Print(" point  \n")
			}
			fmt.Print("------------------------------------------\n")
		}
		if answered == "N" {
			for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
				fmt.Print("Pemain ke ")
				fmt.Print(idPlayer)
				fmt.Print(" memiliki ")
				fmt.Print(dicePlayer)
				fmt.Print(" dadu dan ")
				fmt.Print(pointPlayer[idPlayer-1])
				fmt.Print(" point  \n")
			}

			if playerWinner > 0 {
				fmt.Print("pemilik point tertinggi adalah Pemain ke ")
				fmt.Print(playerWinner)
				fmt.Print(" dengan ")
				fmt.Print(velueWinner)
				fmt.Print(" point  \n")
			}
			fmt.Println("mengakhiri game")
			fmt.Println("mengakhiri game")
			break
		}
	}
}
