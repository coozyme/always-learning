// function expression
const sayHello = (greet) => {
    console.log(`${greet}!`)
}
 
const sayName = (name) => {
    console.log(`Nama saya ${name}`)
}

/*
   Best Practice

   const sayName = name => console.log(`Nama saya ${name}`);
*/


sayName("Leia");

/* output
Nama saya Leia
 */

const multiply = (a, b) => a * b;
console.log(multiply(3, 4));

/* output
12
 */