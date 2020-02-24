# Ejercicios de Go

## Problema #1

Los múltiplos de 3 ó 5 inferiores a 10 son 3, 5, 6 y 9.

La suma de estos múltiplos es 23. Encontrar la suma de todos los mútliplos de 3
ó 5 menores que 1000.

NOTA: este es el primer problema de [Project Euler](http:projecteuler.net). No
hay ninguna intención de faltar a las normas de ese sitio, sino que se trata de
incentivar a otros a que jueguen allí.

### Ejemplo

``` sh
$ ./ex1 -bound 10
 The sum of all the multiples of 3 or 5 below 10 is 23
```

## Problema #2

La prueba de divisibilidad de un número *n* entre 7 consiste en los siguientes
pasos:

1. Dividir el número *n* en dos partes: *n1* con todos los dígitos menos las
unidades; *n2* que consiste sólo en las unidades
2. If *(n1-n2)* es divisible entre 7 **STOP**, *n* es divisible entre 7
3. En otro caso, *n* es divisible entre 7 si y sólo si lo es *(n1-n2)*

### Ejemplo

``` sh
$ ./ex2 --number 870458270989
 The number 870458270989 is not divisible by 7 and the remainder, indeed is 6
$ ./ex2 --number 870458270990
 The number 870458270990 is divisible by 7!
```

## Problema #3

Programar un paquete (*register*) que atienda peticiones de un pequeño servidor
HTTP en un puerto que debe indicarse como parámetro:

- Las peticiones a `"/"` deben simplemente dar la bienvenida al usuario
- Las peticiones a `"/timestamp"` deben mostrar la hora y fecha actual
- Las peticiones a `"/user"` deben mostrar el nombre de usuario actual

Para ello, el paquete debe ofrecer una única función exportable `Serve``

### Ejemplo

El servicio `Serve` podría usarse como sigue:

``` go
    register.Serve(port)
```

de modo que ejecutándolo:

``` sh
$ ./test --port 8080
```

cuando se abre el explorador en la dirección `localhost:8080` se ve:

``` text
Welcome to our tiny experimental server!
```

La dirección `localhost:8080/timestamp` podría devolver:

``` text
2020-02-24 17:49:31.623037828 +0100 CET m=+51.941346902
```

y, por último, la dirección `localhost:8080/user` devuelve:

``` text
clinares
```

# License #

goex is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your
option) any later version.

goex is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
for more details.

You should have received a copy of the GNU General Public License
along with goex.  If not, see <http://www.gnu.org/licenses/>.


# Author #

Carlos Linares Lopez <carlos.linares@uc3m.es>
