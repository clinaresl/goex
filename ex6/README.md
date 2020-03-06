## Ejercicio #6

Implementar una versión mejorada del comando Unix ``cal`` que muestra un
calendario en diferentes intervalos temporales. Los argumentos básicos que debe
reconocer son:

- ``-1`` muestra únicamente un mes
- ``-3`` muestra el mes anterior y el siguiente al actual
- ``--months NUMBER`` muestra el número de meses indicado a partir de la fecha
  de inicio
- ``sunday`` fuerza el domingo como primer día de la semana que, por
  defecto, empiezan los lunes.
- ``week-numbering`` numera todas las semanas de cada año

Además, el comando Unix ``cal`` muestra los calendarios en bloques de tres
meses. Se sugiere implementar un argumento adicional para poder mostrarlos en
bloques de cualquier tamaño:

- ``blocks NUMBER`` para mostrar bloques de un tamaño arbitrario

Por último, la salida debe mostrar en diferentes colores la fecha actual o una
fecha proporcionada por el usuario. También pueden distinguirse los domingos del
resto de la semana y cualesquiera otras leyendas como el nombre del mes o los
días de la semana. Por lo tanto, debe proporcionarse también el argumento:

- ``disable-highlighting`` para mostrar el calendario en el mismo color. Las
  fechas que se resaltan de otras (esto es, la fecha actual y la que proporcione
  el usuario si entrega alguna), se deben mostrar en vídeo inverso.

### Ejemplo

Ejecutando el programa sin argumentos, se debe mostrar únicamente el mes con la
fecha actual:

``` sh
$ ./ex6
```

![Example 1](pics/pic1.png)

El mismo calendario puede mostrarse pero en bloques de 4 meses como se indica a
continuación

``` sh
$ ./ex6 --blocks 4 2020
```

![Example 2](pics/pic2.png)

En el siguiente ejemplo se muestra el mismo calendario pero en bloques de 5
meses y numerando todas las semanas:

``` sh
$ ./ex6 --blocks 5 --week-numbering 2020
```

![Example 3](pics/pic3.png)


Por último, si se entrega una fecha en particular entonces debe resaltarse. El
siguiente ejemplo muestra tres meses consecutivos con la fecha elegida en el mes
mostrado en el centro:

``` sh
$ ./ex6 -3 16 8 2010
```

![Example 4](pics/pic4.png)


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
