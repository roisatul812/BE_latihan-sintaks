package main

import "fmt"

func main() {

	// ========================================
	// 1. Variabel dengan tipe berbeda
	// ========================================

	// String
	nama := "Isa"

	// Integer
	umur := 20

	// Float64
	ipk := 3.75

	// Boolean
	isActive := true

	// Slice
	hobi := []string{"Coding", "Desain", "Membaca"}

	fmt.Println("=== VARIABEL ===")
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("IPK:", ipk)
	fmt.Println("Status Aktif:", isActive)
	fmt.Println("Hobi:", hobi)

	// ========================================
	// 2. Map Data Mahasiswa
	// ========================================

	dataMahasiswa := make(map[string]int)

	// Menambah data
	dataMahasiswa["Budi"] = 85
	dataMahasiswa["Sari"] = 90
	dataMahasiswa["Andi"] = 88

	fmt.Println("\n=== MAP ===")
	fmt.Println("Data mahasiswa:", dataMahasiswa)

	// Membaca data dengan pengecekan keberadaan
	nilai, ada := dataMahasiswa["Sari"]

	if ada {
		fmt.Println("Nilai Sari:", nilai)
	} else {
		fmt.Println("Data Sari tidak ditemukan")
	}

	// Menghapus data
	delete(dataMahasiswa, "Andi")

	fmt.Println("Setelah Andi dihapus:", dataMahasiswa)

	// Menelusuri seluruh isi map
	fmt.Println("Seluruh data mahasiswa:")

	for nama, nilai := range dataMahasiswa {
		fmt.Println(nama, ":", nilai)
	}

	// ========================================
	// 3. Pointer
	// ========================================

	fmt.Println("\n=== POINTER ===")

	// ----------------------------------------
	// Perbandingan Pass by Value dan Pointer
	// ----------------------------------------

	nilaiAwal := 42

	fmt.Println("\n=== PASS BY VALUE VS POINTER ===")
	fmt.Println("Nilai awal:", nilaiAwal)

	// Pass by value
	ubahNilai(nilaiAwal)
	fmt.Println("Setelah pass by value:", nilaiAwal)

	// Pass by pointer
	ubahLewatPointer(&nilaiAwal)
	fmt.Println("Setelah pass by pointer:", nilaiAwal)

	// ----------------------------------------
	// Function swap menggunakan pointer
	// ----------------------------------------

	a := 10
	b := 20

	fmt.Println("\nSebelum swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	swap(&a, &b)

	fmt.Println("Setelah swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// ----------------------------------------
	// Function updateSlice menggunakan pointer
	// ----------------------------------------

	hobiBaru := []string{"Coding", "Desain"}

	fmt.Println("\nSlice sebelum update:", hobiBaru)

	updateSlice(&hobiBaru, "Membaca")

	fmt.Println("Slice setelah update:", hobiBaru)

	// ========================================
	// 4. Struct Student
	// ========================================

	fmt.Println("\n=== STRUCT STUDENT ===")

	student := Student{
		ID:       1,
		Name:     "Isa",
		Grade:    85,
		IsActive: false,
	}

	fmt.Println("Info awal:", student.GetInfo())

	// Update grade
	student.UpdateGrade(90)

	fmt.Println("Setelah UpdateGrade:", student.GetInfo())

	// Activate
	student.Activate()

	fmt.Println("Setelah Activate:", student.GetInfo())

	// Deactivate
	student.Deactivate()

	fmt.Println("Setelah Deactivate:", student.GetInfo())
}

// ========================================
// Function Pass by Value dan Pointer
// ========================================

// Pass by value
func ubahNilai(x int) {
	x = 100
}

// Pass by pointer
func ubahLewatPointer(x *int) {
	*x = 100
}

// ========================================
// Function Pointer
// ========================================

// Menukar dua integer menggunakan pointer
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Menambahkan item baru ke slice menggunakan pointer
func updateSlice(s *[]string, newItem string) {
	*s = append(*s, newItem)
}

// ========================================
// Struct Student
// ========================================

type Student struct {
	ID       int
	Name     string
	Grade    float64
	IsActive bool
}

// GetInfo menggunakan value receiver
func (s Student) GetInfo() string {
	return fmt.Sprintf(
		"ID: %d, Name: %s, Grade: %.2f, Active: %t",
		s.ID,
		s.Name,
		s.Grade,
		s.IsActive,
	)
}

// UpdateGrade menggunakan pointer receiver
func (s *Student) UpdateGrade(grade float64) {
	s.Grade = grade
}

// Activate menggunakan pointer receiver
func (s *Student) Activate() {
	s.IsActive = true
}

// Deactivate menggunakan pointer receiver
func (s *Student) Deactivate() {
	s.IsActive = false
}