package main

import (
	"fmt"
)

const Nmax int = 500

type latihan struct {
	namaLatihan   string
	durasiLatihan int
	kaloriLatihan int
	hari          int
	bulan         int
	tahun         int
}

type tabLatihan [Nmax]latihan

type user struct {
	ID            int
	nama          string
	umur          int
	beratBadan    int
	tinggiBadan   int
	jumlahLatihan int
	latihan       tabLatihan
}

type tabUser [Nmax]user

func main() {
	var dataUser tabUser
	var idCount int = 0
	var pilihan int

	fmt.Println("==============================================")
	fmt.Println("Selamat datang di aplikasi workout harian!")
	fmt.Println("1. Belum punya ID pengguna")
	fmt.Println("2. Sudah punya ID pengguna")
	fmt.Print("Masukkan nomor pilihan Anda: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		belumPunyaID(&dataUser, &idCount)
		cetakIDbaru(dataUser, idCount)
		sudahPunyaID(&dataUser[idCount-1])
	} else if pilihan == 2 {
		if idCount == 0 {
			fmt.Println("Belum ada pengguna yang terdaftar.")
			return
		}
		var inputID int
		fmt.Print("Masukkan ID pengguna Anda: ")
		fmt.Scan(&inputID)
		if inputID >= 0 && inputID < idCount {
			sudahPunyaID(&dataUser[inputID])
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func belumPunyaID(A *tabUser, id *int) {
	fmt.Println("\nHalo pengguna baru, silakan isi data berikut:")

	fmt.Print("Nama panggilan: ")
	fmt.Scan(&A[*id].nama)
	fmt.Print("Umur: ")
	fmt.Scan(&A[*id].umur)
	fmt.Print("Berat Badan: ")
	fmt.Scan(&A[*id].beratBadan)
	fmt.Print("Tinggi Badan: ")
	fmt.Scan(&A[*id].tinggiBadan)

	A[*id].ID = *id
	*id = *id + 1
}

func cetakIDbaru(A tabUser, id int) {
	uid := id - 1
	fmt.Println("\nSelamat, kamu sudah resmi menjadi pengguna!")
	fmt.Println("ID:", A[uid].ID)
	fmt.Println("Nama:", A[uid].nama)
	fmt.Println("Umur:", A[uid].umur)
	fmt.Println("Berat Badan:", A[uid].beratBadan)
	fmt.Println("Tinggi Badan:", A[uid].tinggiBadan)
}

func sudahPunyaID(userAktif *user) {
	fmt.Println("\nHalo", userAktif.nama, ", selamat datang!")

	for {
		var pilihan int
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Latihan")
		fmt.Println("2. Lihat Semua Latihan")
		fmt.Println("3. Lihat Latihan Berdasarkan Tanggal")
		fmt.Println("4. Ubah Data Latihan")
		fmt.Println("5. Hapus Data Latihan")
		fmt.Println("6. Rekomendasi latihan")
		fmt.Println("7. Urutkan Latihan Berdasarkan Durasi (Descending)")
		fmt.Println("8. Urutkan Latihan Berdasarkan Kalori (Descending)")
		fmt.Println("9. Keluar")

		fmt.Print("Masukkan pilihan: ") // <--- Tambahkan ini
		fmt.Scan(&pilihan)              // <--- Tambahkan ini

		switch pilihan {
		case 1:
			tambahLatihan(userAktif)
		case 2:
			tampilkanLatihan(userAktif)
		case 3:
			tampilkanLatihanTanggal(userAktif)
		case 4:
			ubahLatihan(userAktif)
		case 5:
			hapusLatihan(userAktif)
		case 6:
			rekomendasi(userAktif)
		case 7:
			urutkanLatihanDurasiDesc(userAktif)
		case 8:
			urutkanLatihanKaloriDesc(userAktif)
		case 9:
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahLatihan(userAktif *user) {
	if userAktif.jumlahLatihan >= Nmax {
		fmt.Println("Daftar latihan penuh.")
		return
	}

	fmt.Println("\nMasukkan tanggal latihan:")
	fmt.Print("Tanggal: ")
	fmt.Scan(&userAktif.latihan[userAktif.jumlahLatihan].hari)
	fmt.Print("Bulan: ")
	fmt.Scan(&userAktif.latihan[userAktif.jumlahLatihan].bulan)
	fmt.Print("Tahun: ")
	fmt.Scan(&userAktif.latihan[userAktif.jumlahLatihan].tahun)

	fmt.Println("\nMasukkan data latihan:")
	fmt.Print("Nama Latihan: ")
	fmt.Scan(&userAktif.latihan[userAktif.jumlahLatihan].namaLatihan)
	fmt.Print("Durasi (detik): ")
	fmt.Scan(&userAktif.latihan[userAktif.jumlahLatihan].durasiLatihan)
	fmt.Print("Kalori terbakar: ")
	fmt.Scan(&userAktif.latihan[userAktif.jumlahLatihan].kaloriLatihan)

	userAktif.jumlahLatihan++
	fmt.Println("Latihan berhasil ditambahkan.")
}

func tampilkanLatihan(userAktif *user) {
	if userAktif.jumlahLatihan == 0 {
		fmt.Println("Belum ada latihan.")
		return
	}

	fmt.Println("\nSemua Latihan:")
	for i := 0; i < userAktif.jumlahLatihan; i++ {
		lat := userAktif.latihan[i]
		fmt.Printf("[%d] %s | %d detik | %d kalori | Tanggal: %02d-%02d-%d\n",
			i+1, lat.namaLatihan, lat.durasiLatihan, lat.kaloriLatihan, lat.hari, lat.bulan, lat.tahun)
	}
}

func tampilkanLatihanTanggal(userAktif *user) {
	var h, b, t int
	fmt.Println("\nMasukkan tanggal yang ingin dilihat:")
	fmt.Print("Hari: ")
	fmt.Scan(&h)
	fmt.Print("Bulan: ")
	fmt.Scan(&b)
	fmt.Print("Tahun: ")
	fmt.Scan(&t)

	found := false
	fmt.Printf("\nLatihan pada %02d-%02d-%d:\n", h, b, t)
	for i := 0; i < userAktif.jumlahLatihan; i++ {
		lat := userAktif.latihan[i]
		if lat.hari == h && lat.bulan == b && lat.tahun == t {
			fmt.Printf("|%d |%s | %d detik | %d kalori\n",
				i+1, lat.namaLatihan, lat.durasiLatihan, lat.kaloriLatihan)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada data pada tanggal tersebut.")
	}
}

func ubahLatihan(userAktif *user) {
	var no int
	tampilkanLatihan(userAktif)
	fmt.Print("\nMasukkan nomor latihan yang ingin diubah: ")
	fmt.Scan(&no)
	if no < 1 || no > userAktif.jumlahLatihan {
		fmt.Println("Nomor tidak valid.")
		return
	}
	idx := no - 1
	fmt.Println("\nMasukkan data latihan baru:")
	fmt.Print("Nama Latihan: ")
	fmt.Scan(&userAktif.latihan[idx].namaLatihan)
	fmt.Print("Durasi (detik): ")
	fmt.Scan(&userAktif.latihan[idx].durasiLatihan)
	fmt.Print("Kalori terbakar: ")
	fmt.Scan(&userAktif.latihan[idx].kaloriLatihan)
	fmt.Print("Tanggal: ")
	fmt.Scan(&userAktif.latihan[idx].hari)
	fmt.Print("Bulan: ")
	fmt.Scan(&userAktif.latihan[idx].bulan)
	fmt.Print("Tahun: ")
	fmt.Scan(&userAktif.latihan[idx].tahun)
	fmt.Println("Latihan berhasil diubah.")
}

func hapusLatihan(userAktif *user) {
	var no int
	tampilkanLatihan(userAktif)
	fmt.Print("\nMasukkan nomor latihan yang ingin dihapus: ")
	fmt.Scan(&no)
	if no < 1 || no > userAktif.jumlahLatihan {
		fmt.Println("Nomor tidak valid.")
		return
	}
	idx := no - 1
	for i := idx; i < userAktif.jumlahLatihan-1; i++ {
		userAktif.latihan[i] = userAktif.latihan[i+1]
	}
	userAktif.jumlahLatihan--
	fmt.Println("Latihan berhasil dihapus.")
}

func rekomendasi(userAktif *user) {
	var beratBadan int

	fmt.Print("Masukkan berat badan anda saat ini: ")
	fmt.Scan(&beratBadan)

	switch {
	case beratBadan < 50:
		fmt.Println("\nRekomendasi latihan anda saat ini:")
		fmt.Println("1. Yoga ringan ")
		fmt.Println("2. Jalan kaki")
		fmt.Println("3. Stretching")
	case beratBadan >= 50 && beratBadan < 70:
		fmt.Println("\nRekomendasi latihan anda saat ini:")
		fmt.Println("1. Jogging")
		fmt.Println("2. Bodyweight workout")
		fmt.Println("3. Berenang")
	case beratBadan >= 70 && beratBadan < 90:
		fmt.Println("\nRekomendasi latihan anda saat ini:")
		fmt.Println("1. Bersepeda")
		fmt.Println("2. HIIT menengah")
		fmt.Println("3. Circuit training")
	case beratBadan >= 90:
		fmt.Println("\nRekomendasi latihan anda saat ini:")
		fmt.Println("1. Jalan cepat")
		fmt.Println("2. Elliptical")
		fmt.Println("3. Low-impact cardio")
	default:
		fmt.Println("Data berat badan tidak valid.")
	}
}

func urutkanLatihanDurasiDesc(userAktif *user) {
	var idxMax, pass, i int
	var temp latihan

	if userAktif.jumlahLatihan == 0 {
		fmt.Println("Belum ada latihan.")
		return
	}

	pass = 0
	for pass < userAktif.jumlahLatihan-1 {
		idxMax = pass
		i = pass + 1
		for i < userAktif.jumlahLatihan {
			if userAktif.latihan[idxMax].durasiLatihan < userAktif.latihan[i].durasiLatihan {
				idxMax = i
			}
			i++
		}
		temp = userAktif.latihan[pass]
		userAktif.latihan[pass] = userAktif.latihan[idxMax]
		userAktif.latihan[idxMax] = temp

		pass++
	}

	fmt.Println("\nLatihan berhasil diurutkan berdasarkan durasi (terlama ke tersingkat):")
	tampilkanLatihan(userAktif)
}

func urutkanLatihanKaloriDesc(userAktif *user) {
	var temp latihan
	var i, pass int

	if userAktif.jumlahLatihan == 0 {
		fmt.Println("Belum ada latihan.")
		return
	}

	pass = 1
	for pass < userAktif.jumlahLatihan {
		temp = userAktif.latihan[pass]
		i = pass
		for i > 0 && temp.kaloriLatihan > userAktif.latihan[i-1].kaloriLatihan {
			userAktif.latihan[i] = userAktif.latihan[i-1]
			i--
		}

		userAktif.latihan[i] = temp
		pass++
	}

	fmt.Println("\nLatihan berhasil diurutkan berdasarkan kalori (terbanyak ke tersedikit):")
	tampilkanLatihan(userAktif)
}
