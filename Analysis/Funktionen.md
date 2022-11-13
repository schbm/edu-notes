
# Funktionen
## Grundlagen

### Schreibweise
Da Element $x$ aus $D$ heisst "Argument". Der Wert $f(x) \in Z$ welche durch die Zuordnung entsteht, heisst "Funktionswert".
Der Funktionswert $f(x)$ ensteht durch Anwender der Funktion $f$ auf das Argument $x$.:

$$ D = \text{Definitionsmenge} $$

$$ Z = \text{Zielmenge}$$

```math
f =
  \begin{cases}
    D \to Z \\
    x \mapsto \text{Zordnungsvorschrift}
  \end{cases}
```


### Bild
Menge aller Funktionswerte:
$$\text{Bild}(f) = f(D)$$
Das Bild der Funktion ist immer eine Teilmenge der Zielmenge.

### Graph
Menge aller Wertepaare $(x,y)$.

$$ Graph(f) = \set{(x,y) \in D \times Z \mid f(x) = y } $$

![Kartesisches Koordinatensystem](https://www.grund-wissen.de/mathematik/_images/koordinatensystem-kartesisch.png)

### Implizite Funktionen
Bei implizit definierten Funktionen müssen Funktionswerte durch lösen einer Gleichung bestimmt werden.

### Wohldefiniertheit
Die Funktionsdefinition ist wohldefiniert, wenn es zu jedem Wert auf der 1. Koordinatenachse, 
nur einen einzigen Wert auf der 2. Koordinatenachse gibt.
Oder: Eine Funktionsdefinition ist wohldefiniert,
wenn jede zur 2. Koordinatenachse parallele Gerade den Funktionsgraphen genau einmal schneidet.

## Umkehrfunktionen
![Begriffe Umkehrfunktion](https://images.gutefrage.net/media/fragen-antworten/bilder/295658895/0_big.webp?v=1541879960000)

$f$ ist injektiv, wenn für alle $x,y \in D$ mit $x \neq y$

$$ f(x) \neq f(y) $$ 

$f$ ist surjektiv, wenn alle Werte aus Z 'erreicht' werden:

$$ Bild(f) = Z $$

$f$ ist bijektiv (und somit umkehrbar), wenn sie surjektiv und injektiv ist:
```math
f^{-1} =
  \begin{cases}
    Z \to D \\
    y \mapsto f(x) = y
  \end{cases}
```

$$ f^{-1}(f(x)) = x \text{ für alle } x \in D $$

$$ f(f^{-1}(y)) = y \text{ für alle } y \in Z $$

$$ (f^{-1})^{-1} = f $$

### Exponentialfunktion

$$ log_a = exp_a^{-1} $$

### Potenzfunktion
Für $a > 0$ ist $\sqrt[a]{}$ die Umkehrfunktion:

```math
pot_a^+ =
  \begin{cases}
    \mathbb{R}^+ \to \mathbb{R}^+ \\
    x \mapsto x^a
  \end{cases}
```

$$ \sqrt[a]{*} = (pot_a^+)^{-1} $$

## Funktionseigenschaften
### TODO

## Funktionen mit einem Computer
### TODO
