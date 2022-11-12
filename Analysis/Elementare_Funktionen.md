# Elementare Funktionen

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

![Wurzelfunktion](https://github.com/schbm/edu-notes/blob/main/Analysis/Wurzelfunktion.PNG?raw=true)

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
Da zum Beispiel $(a^2)^{\frac{1}{2}} = a^{2*\frac{1}{2}} = a^1 = a$ für $a \geq 0$.
Jedoch für $a < 0$ eribt diese Rechnung $\sqrt{a^2} = |a| = a \lor -a$.

Beispiel $x^{1/2}$ für $D=\mathbb{R}$:
![Nicht wohldefinierte Funktion](https://github.com/schbm/edu-notes/blob/main/Analysis/nicht_wohldefiniert.PNG?raw=true)

## Exponential Funktion

## Logarithmus
