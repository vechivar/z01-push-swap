## INFORMATIONS SUR L'ALGORITHME

L'essentiel du projet a consisté à produire un algorithme capable de trier de grandes listes de nombres en un nombre d'opérations acceptable.
La version "naïve" de l'algorithme (aller chercher le minimum de A, le push dans B, et ainsi de suite) donnait un nombre d'opérations de l'ordre de 1500 pour 100 nombres aléatoires. Notre version est passée au-dessus de la barre des 700 opérations imposée dans l'audit dans 10 cas sur 10 000 tentatives aléatoires.
De nombreuses optimisations sont encore possibles (swap n'est jamais utilisé, par exemple), et l'algorithme est très mauvais pour pour trier de petites listes de nombres. Nous avons cependant considéré que notre travail était concluant.
Une procédure (très) spéciale est utilisée pour trier les petites piles (N <= 6).

## PRINCIPE GENERAL DE L'ALGORITHME (illustré pour N = 100)

Etape 0 : tous les nombres sont mélangés dans A

Etape 1 : trier les nombres en trois listes. Dans chaque liste, les nombres ne sont pas triés (symbolisé par ~)
```
                _________
                |       |
                | 33~65 |
_________       |_______|
|       |       |       |   
| 66~99 |       | 00~32 |
|_______|       |_______|
    A               B
```

Etape 2 : trier les nombres de A dans la liste de B (nombres entre x et y triés avec x au sommet symbolisés par [x-y])

``` 
                _________
                |       |   
                |[99-66]|
                |_______|
                |       |
                | 33~65 |
                |_______|
                |       |   
                | 00~32 |
________        |_______|
    A               B
```

Etape 2.5 : retransférer ces nombres dans A

```
                _________
                |       |
                | 33~65 |
_________       |_______|
|       |       |       |   
|[66-99]|       | 00~32 |
|_______|       |_______|
    A               B
```

Etape 3 : trier le prochain tiers de B dans A

```
_________                
|       |                
|[33-65]|                
|_______|       _________
|       |       |       |   
|[66-99]|       | 00~32 |
|_______|       |_______|
    A               B

```
Etape 4 : terminer le tri avec le dernier tiers

```
_________                
|       |                
|[00-32]|                
|_______|                
|       |                
|[33-65]|                
|_______|
|       |
|[66-99]|         (vide)
|_______|       _________
    A               B

```

Le coût de l'étape 1 est de l'ordre de (N + N/6) si effectué correctement. On trie ainsi trois listes de taille N/3 plutôt que une liste de taille N. Le coût de l'étape 1 est largement compensé, car il est très coûteux de parcourir des listes longues.

On peut aussi remarquer que les tris des étapes 2 et 4 sont plus efficaces que les tris de l'étape 3, car dans l'étape 3 la présence du premier tiers dans la liste à trier empêche d'optimiser les rotate pour le parcours de la liste.

## PRINCIPE DETAILLE DU TRI

On travaille avec des nombres successifs (transition facile à partir des nombres aléatoires en entrée) 

Pour trier les nombres de A vers la pile B, on effectue des rotations de la pile A jusqu'à trouver le maximum de A, puis on push ce nombre vers B.

ATTENTION : dans l'étape 2, les nombres sont triés du plus grand(en haut) vers le plus petit. Dans les étapes 3 et 4, ils sont triés du plus petit vers le plus grand. Dans les exemples suivants, on triera du plus petit vers le plus grand, mais le principe reste similaire.

Le tri est optimisé avec les méthodes suivantes.

(1) Lorsque l'on cherche le prochain nombre de A à transférer vers B, on peut effectuer des rotate ou des reverse-rotate en fonction de sa position dans la pile.

(2) Il est également possible d'envoyer le minimum de A vers B au lieu du maximum. Il faut alors faire un rotate sur la pile b pour placer le minimum en-dessous. A la fin, il suffit de faire des rotate sur B pour replacer le minimum au-dessus de la pile.

Exemple:

```
5                                                                                                                   3
1           1                                                                               2           4           4
3     (pb)  3       (pb)    3       (rb)    3       (ra)    4   (pb)        4   (pb)        4   (rb)    5   (pb)    5
4           4               4   1           4   5           2           2   5               5           1           1
2   X       2   5           2   5           2   1           3           3   1           3   1           2       X   2

```

Puis on rotate pour replacer 1 au-dessus.

(3) Lorsqu'on choisit de push un minimum, et qu'on rencontre le prochain minimum au sommet de la pile pendant les rotate, on peut push ce prochain minimum. Il faut alors effectuer un rotate de plus sur la pile triée pour passer les deux minimums en bas de la pile.

(4) Lorsqu'il faut rotate des minimums dans la pile triée, on peut éventuellement combiner ces rotate avec ceux utilisés pour parcourir la pile non-triée (utilisation de rr) si la pile non-triée est parcourue dans le bon sens.

(5) Pour décider si on push un (ou plusieurs) minimum ou un maximum, on choisir le plus petit nombre de rotate moyen pour accéder à ces valeurs. (ex : 7 rotate pour 2 minimums => 3.5 rotate par valeur. on choisit le minimum si les rotate pour atteindre le max sont >=4).