/*
 Diberikan sebuah string `line`. Isi dari string ini adalah sebuah gabungan dari sembarang
 karakter. Tugas kamu adalah mengubah karakter yang merupakan virus menjadi konsonan atau vokal sesuai dengan
 requirement berikut:

 1. Jika di dalam 'line' jumlah karakter konsonan lebih banyak dari jumlah karakter vokal, 
    maka ubahlah semua karakter yang merupakan virus menjadi 'a'.
 2. Jika di dalam 'line' jumlah karakter vokal lebih banyak dari jumlah karakter konsonan, 
    maka ubah semua karakter yang merupakan virus menjadi 'b'.
 3. Jika di dalam 'line' jumlah karakter vokal sama dengan jumlah karakter konsonan, 
    maka ubah semua karakter yang merupakan virus menjadi 'c'.
 4. Jika tidak ada virus sama sekali, maka line tidak perlu diubah.
 5. Program akan menampilkan 'line' yang telah diubah sesuai dengan requirement diatas

 Berikut ini adalah contoh karakter vokal dan konsonan.
 Vokal = a i u e o
 Virus = #
 Konsonan = selain Vokal & Virus

 contoh 1:
 line = 'abc#ab#ueo'

 maka output yang diharapkan adalah 'abcbabbueo'
 karena jumlah vokal adalah 5, jumlah konsonan adalah 3. Lalu, untuk kedua virus
 yang ditemukan, maka semua karakter virus diubah menjadi 'b'.


 contoh 2:
 line = 'abcabdueobbb'

 Karena tidak terdapat virus, maka output yang diharapkan adalah
 'abcabdueobbb'
 
 [Ketentuan]:
 - TIDAK BOLEH menggunakan built-in function!
*/

function fixData(line) {
  // beginanswer
  let vokal = 'aiueo';
  var vok = 0;
  var virus = 0;
  var hasil = "";

  for (var i = 0; i < line.length; i++) {
    if (line[i] === "#") {
      virus++;
    }
    for (var j = 0; j < vokal.length; j++) {
      if (line[i] === vokal[j]) {
        vok++;
      }
    }
  }

  var kons = line.length - (vok + virus);

  if (virus === 0) {
    console.log(line)
  }
  else {
    if (kons > vok) {
      for (var k = 0; k < line.length; k++) {
        if (line[k] === "#") {
          hasil += 'a'
        }
        else {
          hasil += line[k]
        }
      }
    }
    else if (vok > kons) {
      for (var k = 0; k < line.length; k++) {
        if (line[k] === "#") {
          hasil += 'b'
        }
        else {
          hasil += line[k]
        }
      }
    }
    else if (vok === kons) {
      for (var k = 0; k < line.length; k++) {
        if (line[k] === "#") {
          hasil += 'c'
        }
        else {
          hasil += line[k]
        }
      }
    }
    return hasil
  }
  //endanswer
}

console.log(fixData('aoi#fdg#ue'))
console.log(fixData('eh#xyz#oi#'))
console.log(fixData('#eui#bcl##'))

module.exports = fixData