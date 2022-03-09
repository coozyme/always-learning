#!/bin/sh

val=`expr 2 + 2`
val1='4'
val2='2'

Addition=`expr $val1 + $val2`
Subtraction=`expr $val1 - $val2`
Multiplication=`expr $val1 \* $val2`
Division=`expr $val1 / $val2`
Modulus=`expr $val1 % $val2`

echo "Addition : $Addition"
echo "Subtraction : $Subtraction"
echo "Multiplication : $Multiplication"
echo "Division : $Modulus"
echo "Modulus : $Modulus"