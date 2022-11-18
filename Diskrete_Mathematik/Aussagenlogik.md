# Aussagenlogik

## Definitionen
Eine Aussage ist ein feststellender Satz, dem eindeutig
einer der beiden Wahrheitswerte wahr oder falsch zugeordnet werden kann.

Zeichen | Bedeutung | Beispiel
--------|-----------|---------
$\land$ | Konjunktion | $A \land B$
$\lor$ | Disjunktion | $A \land B$
$neg$ | Negation | $\neg A$
$\iff$ | Äquivalenz; genau dann, wenn | $A \iff \ B$
$\implies$ | Implikation; wenn, dann; Annahme, Folgerung; Hinreichend, Notwendig | $A \implies B$

- Tautologie = Aussage ist immer wahr
- Kontradiktion = Aussage ist immer falsch

### Äquivalenz
A | B | $A \iff B$
--------|-----------|---------
W | W | W
W | F | F
F | W | F
F | F | W

$$ (A \iff B) \iff (A \land B) \lor (\neg A \land \neg B)$$

Liegt eine Äquivalenz vor, dann ist die Spalte, in der die
Äquivalenz bestimmt wird, ist immer wahr.

### Implikation
A | B | $A \implies B$
--------|-----------|---------
W | W | W
W | F | F
F | W | W
F | F | W

$$ \neg (A \implies B) \iff \neg B \land A $$

$$ (A \implies B \implies C) \iff (A \implies B) \land (B \implies C)$$

- „ex falso quod libet sequitur“, aus dem Falschen folgt Beliebiges

Man sagt dem Mathematiker Bertrand Russel folgende Geschichte nach: Als er einem
Journalisten erklärte, dass man aus Falschem Beliebiges ableiten könnte, sagte der Journalist: 
„Gut: 1+1=3 ist falsch. Beweisen Sie mir, dass Sie der Papst sind!“ 

Nach kurzem Nachdenken kam die Antwort: „Wenn ich auf jeder Seite 1 abziehe, kann ich aus 1+1=3
folgern, dass 1=2 ist. Also ist 2=1. Der Papst und ich sind 2. Weil 2 aber gleich 1 ist, bin
ich somit der Papst.“

Das bedeutet: Wenn die Aussage A falsch ist, nützt uns die Implikation nichts, denn man
könnte damit etwas Richtiges, aber auch etwas Falsches herleiten.

### Aussagenlogische Formeln
Verknüpfte Aussagen sind aussagenlogische Formeln, können aus diesen, oder Aussagen, bestehen und bilden wiederum selbst wieder eine Aussage.

### Bindungsstärke
Es gilt $\neg \text{ vor } \land , \lor \text{ vor } \implies,\iff$

### Regeln
- Kommutativität

$$ A \land B \iff B \land A$$

- Assoziativität

$$ A \land (B \land C) \iff (A \land B) \land C $$

- Distributivität

$$ A \land (B \lor C) \iff (A \land B) \lor (A \land C) $$

$$ A \lor (B \land C) \iff (A \lor B) \land (A \lor C) $$

- Absorption

$$ A \lor (A \land B) \iff A $$

$$ A \land (A \lor B) \iff A

- Idempotenz

$$ A \lor A = A $$

$$ A \land A = A$$

- Doppelte Negation

$$ \neg (\neg A) \iff \neg \neg A \iff A $$

### De Morgan

$$ \neg (A \land B) \iff \neg A \lor \neg B$$

$$ \neg (A \lor B) \iff \neg A \land \neg B $$

### Negationsnormalform
Besteht nur wenn die Negation direkt vor Aussagen oder vor Konstanten steht e.g $\neg A \land \neg B$

### Verallgemeinerte Konjunktion
Eine verallg. Konjunktion ist:
- eine **einzelnze** Aussage oder dessen Negation
- einer der lolgischen Konstanten
- eine Konjunktion $A \land B$, falls A und B selbst verallgemeinerte Konjunktionen sind.
Diese liegt immer in Negationsnormalform vor.

### Disjunktive Normalform
Liegt vor, wenn die Aussage A eine verallgemeinerte Konjunktion ist, oder eine Disjunktion von verallgemeinerten Konjunktionen.
Analog dazu gibt es die Konjunktive normalform. (Dazu kann man die disjunktive Normalform negieren!)

$$ (A \land B) \lor (\neg A \land B) \lor (...)$$

### Karnaugh-Veitch

TODO
