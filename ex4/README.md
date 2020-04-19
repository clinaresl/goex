# Ejercicio #4

Emular la funcionalidad del comando Unix ``comm``: dados dos ficheros ``FILE1``
y ``FILE2``, dividir los contenidos de ambos ficheros en columnas: la primera
columna debe mostrar sólo las líneas que aparecen en ``FILE1``, la segunda
columna sólo aquellas que aparezcan en el segundo fichero, y la tercera columna
las que aparecen en ambos. Además, el comando debe aceptar los siguientes
argumentos:

- ``1`` deshabilita la presentación de líneas en la primera columna
- ``2`` desactiva la presentación de la segunda columna
- ``3`` para no mostrar las líneas comunes en la tercera columna

El programa no debe asumir, como en el caso del comando Unix ``comm`` que la
entrada está ordenada, y las líneas pueden mostrarse en cualquier orden.

## Ejemplo

Dados los contenidos de dos ficheros:

``` sh
$ more data/file1.txt
Gottfried Leibniz
Gottlob Frege
Charles Babbage
Kurt Goedel
Alan Turing
Martin Davis

$ more data/file2.txt
Charles Babbage
Richard Stallman
Alan Turing
Martin Davis
Ken Thompson

```

Entonces las columnas generadas por el programa serían:

``` sh
$ ./ex4 -file1 data/file1.txt -file2 data/file2.txt
Gottfried Leibniz
Gottlob Frege
		Charles Babbage
Kurt Goedel
		Alan Turing
		Martin Davis
	Richard Stallman
	Ken Thompson
```

Por ejemplo, para ver sólo los contenidos comunes de ambos ficheros se puede
hacer:

``` sh
$ ./ex4 -file1 data/file1.txt -file2 data/file2.txt -1 -2
		Martin Davis
		Charles Babbage
		Alan Turing
```

Nótese que los argumentos se pasan como ``-1 -2`` y no como ``-12`` como en el
caso del comando Unix ``comm``.

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
