function findMin(a, b, c) {
    let minValue = a;

    if (a < b) {
        minValue = a;
    } else {
        minValue = b;
    }
    if (c < minValue) {
        minValue = c;
    }

   return minValue;
   

   /* wrong algoritma */
   // let minValue = 0;

   //  if (a < b) {
   //      minValue = a;
   //  } else if (b < c) {
   //      minValue = b;
   //  } else {
   //      minValue = c;
   //  }

   //  return minValue;
}


function findMax(a, b, c) {
    return Math.max(a, b, c);
}

module.exports = { findMin, findMax };