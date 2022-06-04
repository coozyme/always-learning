let profile = {
    firstName: "John",
    lastName: "Doe",
    age: 18
}
 
let {firstName, lastName, age} = profile;
 
console.log(firstName, lastName, age);
 
/* output:
John Doe 18
*/


/* Destructuring Assignment */

firstName = "Dimas";
age = 20;
 
// menginisialisasi nilai baru melalui destructuring object
({firstName, age} = profile);
 
console.log(firstName);
console.log(age);
 
/* output:
John
18
*/


/* Default Values */


console.log(firstName)
console.log(age)
console.log(isMale)
 
/* output:
John
18
undefined
*/

/* Assigning to Different Local Variable Names */

const {firstName: localFirstName, lastName: localLastName, age: localAge} = profile;
 
console.log(localFirstName);
console.log(localLastName);
console.log(localAge);
 
 
/* output:
John
Doe
18
*/