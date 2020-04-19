# Ejercicio #5

Dado un nombre de usuario de [Lichess](http:lichess.org), obtener toda la
información pública de su perfil a través del servicio [REST API Get
`api/user/{username}`](https://lichess.org/api#operation/playerTopNbPerfType).

Es preciso tener en cuenta que la información devuelta por el servidor podrá
tener unos campos u otros en respuesta a información sobre usuarios diferentes
tal y como se muestra en los siguientes ejemplos.

## Ejemplo

Por ejemplo, con el usuario `clinares`:

``` sh
$ ./ex5 --user clinares

 clinares (Linares López, Carlos)

 * Madrid,ES

       Time spent
 total    7653528
    tv       2568

          # Games
      all   33689
     draw    1347
       ai       8
     winH   15054
    lossH   17280
      win   15056
       me       0
   import      67
    rated   33657
  playing       0
    drawH    1347
     loss   17286
 bookmark      33

        Variant Rating Incr.  Rd # Games
          blitz   1939   -24  45   32217
         bullet   1845   -43  62    1363
 correspondence   1489  -119 162      17
          horde   1321     0 251       5
         puzzle   1983   -49  80    1103
      classical   1883     0 252       2
          rapid   2182   -22 107      47
       chess960   1830     0 229       4

 * Completion rate: 90%

 * Following: 27
 * Followers: 51

 * Online: false
```

y, con otro usuario diferente la salida es:

``` sh
$ ./ex5 --user atorralba

 atorralba 

 * ES

       Time spent
 total    5678838
    tv     131301

          # Games
  playing       0
 bookmark       3
      win   12909
     winH   12906
       me       0
    drawH     894
    lossH    8509
   import       0
       ai      33
      all   22342
     draw     899
     loss    8534
    rated   22293

        Variant Rating Incr.  Rd # Games
  kingOfTheHill   1946   -13 108     273
         puzzle   2073    73  68    6268
          blitz   2199     5  45   10440
    racingKings   1746   -29 154      34
     threeCheck   1867    -2 129     127
     crazyhouse   1939    54 129     475
         atomic   1686  -115 118      44
          rapid   2108    32 105     185
          horde   1886    24 138     146
      classical   2043   -13 153      30
 correspondence   2184   251 139      26
         bullet   2010   -51  56    8956
       chess960   1881   -24  61    1294
    ultraBullet   1534   -29 110     115
      antichess   1757   -38 139     132

 * Completion rate: 97%

 * Following: 8
 * Followers: 10

 * Online: false
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
