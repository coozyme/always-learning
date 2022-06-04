/**
 * TODO:
 * 1. Buatlah class bernama Animal dengan ketentuan:
 *    - Memiliki properti:
 *      - name: string
 *      - age: int
 *      - isMammal: boolean
 *    - Memiliki constructor untuk menginisialisasi properti:
 *      - name
 *      - age
 *      - isMammal
 * 2. Buatlah class bernama Rabbit dengan ketentuan:
 *    - Merupakan turunan dari class Animal
 *    - Memiliki method:
 *      - eat yang mengembalikan nilai string `${this.name} sedang makan!`
 *    - Ketika diinstansiasi, properti isMammal harus bernilai true
 * 3. Buatlah class bernama Eagle dengan ketentuan:
 *    - Merupakan turunan dari class Animal
 *    - Memiliki method:
 *      - fly yang mengembalikan nilai string `${this.name} sedang terbang!`
 *    - Ketika diinstansiasi, properti isMammal harus bernilai false
 * 4. Buatlah instance dari class Rabbit bernama "myRabbit" dengan ketentuan:
 *    - properti name bernilai: "Labi"
 *    - properti age bernilai: 2
 * 5. Buatlah instance dari class Eagle bernama "myEagle" dengan ketentuan:
 *    - properti name bernilai: "Elo"
 *    - properti age bernilai: 4
 */


// TODO

class Animal {
   constructor(name, age, isMammal) {
      this.name = name 
      this.age = age
      this.isMammal = isMammal
   }

}

class Rabbit extends Animal {
   constructor(name, age, isMammal) {
      super(name, age, isMammal)
      this.isMammal = true
  }
   eat() {
      return `${this.name} sedang makan!`
   }
}

class Eagle extends Animal {
    constructor(name, age, isMammal) {
      super(name, age, isMammal)
      this.isMammal = false
  }

   fly() {
      return `${this.name} sedang terbang!`
   }
}


const myRabbit = new Rabbit

myRabbit.name = "Labi"
myRabbit.age = 2

const myEagle = new Eagle

myEagle.name = "Elo"
myEagle.age = 4


console.log(
   myRabbit,
   myRabbit.eat(),
   myEagle,
   myEagle.fly()
)

/**
 * Jangan hapus kode di bawah ini
 */

module.exports = {
  Animal, Rabbit, Eagle, myRabbit, myEagle,
};