/**
 * Konversikan menit pada paramter menjadi satuan jam.
 * Misal: 
 * - 99 menjadi "1:39"
 * - 132 menjadi "2:12"
 * 
 * notes:
 * - type data pada parameter akan selalu INTEGER
 * - type data pada result akan selalu berupa STRING
 */



function konversiMenit(menit) {
  //beginanswer
  var jam = Math.floor(menit / 60)
  var sisa = menit - (jam * 60)
  if (sisa < 10) {
    return `${jam}:0${sisa}`
  }
  else {
    return `${jam}:${sisa}`
  }
  //endanswer
}

console.log(konversiMenit(61));
console.log(konversiMenit(94));
console.log(konversiMenit(51));
console.log(konversiMenit(187));

module.exports = konversiMenit;