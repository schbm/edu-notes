
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

### Implizite Funktionen
Bei implizit definierten Funktionen müssen Funktionswerte durch lösen einer Gleichung bestimmt werden.

### Wohldefiniertheit
Die Funktionsdefinition ist wohldefiniert, wenn es zu jedem Wert auf der 1. Koordinatenachse, 
nur einen einzigen Wert auf der 2. Koordinatenachse gibt.
Oder: Eine Funktionsdefinition ist wohldefiniert,
wenn jede zur 2. Koordinatenachse parallele Gerade den Funktionsgraphen genau einmal schneidet.

## Wurzelfunktion
```math
\sqrt[n]{} =
  \begin{cases}
     \mathbb{R}^+ \to  \mathbb{R}^+ \\
    a \mapsto \text{positive Lösung } x^n = a
  \end{cases}
```

Für $a \geq 0$, dann gilt: $\sqrt{a^2} = a$ oder $\sqrt{a}^2 = a$.
Für $a < 0$ und gerade Exponenten, dann gilt $\sqrt{a^2} = |a|$

## Potenzfunktion
Für $n \geq 0$ und ganzzahlig:
```math
\cdot^n =
  \begin{cases}
     \mathbb{R} \to  \mathbb{R} \\
    x \mapsto x^n
  \end{cases}
```

Für $n < 0$ und ganzzahlig:
```math
\cdot^n =
  \begin{cases}
     \mathbb{R} \backslash \set{0} \to  \mathbb{R} \backslash \set{0} \\
    x \mapsto  x^n
  \end{cases}
```
da $a^{-n} = \frac{1}{a^n}$ und $0^{-n} = \frac{1}{0^n}$ und dies ist nur für $n=0$ definiert da $0^n=1$

Für $n \geq 0$ und $n \in \mathbb{R}$:
```math
\cdot^n =
  \begin{cases}
     \mathbb{R}^+ \to  \mathbb{R}^+ \\
    x \mapsto x^n
  \end{cases}
```

Für $n < 0$ und $n \in \mathbb{R}$:
```math
\cdot^n =
  \begin{cases}
     \mathbb{R}^+ \backslash \set{0} \to  \mathbb{R}^+ \backslash \set{0} \\
    x \mapsto  x^n
  \end{cases}
```
