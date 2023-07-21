Pada kode yang telah saya buat, terdapat fungsi isBalanced yang memiliki fungsi untuk
memeriksa apakah tanda kurung dalam sebuah string s seimbang (balanced) atau tidak. 
Saya melakukan analisis kompleksitas untuk fungsi ini dengan mengidentifikasi 
beberapa operasi utama yang ada di dalamnya. 
Pertama, terdapat operasi inisialisasi stack (slice) dengan panjang 0.
Operasi ini memiliki kompleksitas 0(1) karena inisialisasi stack hanya dilakukan sekali pada awal fungsi,
dan tidak tergantung pada ukuran string input s. 
Selanjutnya, dilakukan iterasi pada setiap karakter dalam string s. 
Operasi ini memiliki kompleksitas 0(n), di mana n adalah panjang dari string s. 
Karena operasi ini bergantung pada ukuran input, kompleksitas keseluruhan fungsi isBalanced adalah 0(n). 
Operasi push dan pop pada stack (slice) memiliki kompleksitas 0(1), 
karena slice diimplementasikan dengan array dinamis dan operasi ini dapat dilakukan dalam waktu konstan, 
tanpa ketergantungan pada ukuran input. Selain itu, terdapat operasi pencarian pada map bracketPairs,
yang memiliki kompleksitas 0(1) pada rata-rata karena Go map menggunakan hashing untuk mencari kunci.
Namun, dalam kasus terburuk, pencarian pada map bisa menjadi 0(log n), di mana n adalah jumlah elemen di map. 
Namun, dalam konteks ini, karena bracketPairs hanya memiliki 3 elemen yang tetap, pencarian pada map dapat dianggap sebagai 0(1). 
Kompleksitas dari fungsi isBalanced pada kode main yang saya buat dipengaruhi oleh ukuran dari masing-masing string
input (input1, input2, dan input3). 
Jika ukuran string-input ini berbeda-beda, kompleksitas dari fungsi isBalanced akan berbeda pula, 
sesuai dengan panjang dari masing-masing string tersebut. 
Dengan menggunakan pendekatan stack untuk melacak tanda kurung dan melakukan iterasi sekali pada setiap karakter dalam string, 
saya berhasil membuat solusi yang efisien dalam mengecek keseimbangan tanda kurung pada sebuah string.
