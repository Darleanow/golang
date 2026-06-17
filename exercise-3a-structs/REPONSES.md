# Exercice 3 — Questions de réflexion

## 3.2 — Empêcher les dimensions invalides

Pour interdire un rectangle ou un cercle aux dimensions négatives, on n'expose
pas la construction directe par littéral : on passe par une fonction
« constructeur » qui valide les entrées et retourne `(valeur, error)`.

C'est ce que font `NewRectangle` et `NewCircle` dans le code : elles vérifient
les invariants (`Max >= Min`, `Radius >= 0`) et renvoient une erreur explicite
si la donnée est invalide, plutôt que de produire une forme incohérente.

## 3.3 — Receiver de valeur vs receiver de pointeur

**Receiver de valeur** (`func (r Rectangle)`) : la méthode reçoit une *copie*
de la structure. Toute modification reste locale à la copie et ne touche pas
l'original. Adapté aux méthodes de lecture seule.

**Receiver de pointeur** (`func (r *Rectangle)`) : la méthode reçoit l'adresse
de l'instance. Elle peut donc *muter* la valeur d'origine ; c'est aussi plus
efficace pour les grosses structures, qui ne sont pas copiées.

**Justification des choix du TP :**

- `Move` et `Scale` utilisent un **receiver de pointeur** : leur but est de
  modifier l'instance (déplacer le rectangle, redimensionner le cercle). Avec un
  receiver de valeur, la mutation porterait sur une copie et serait perdue.
- `Area`, `Perimeter`, `Width`, `Height`, `DistanceTo` et `Circumference`
  utilisent un **receiver de valeur** : elles ne font que lire les champs et
  calculer un résultat, sans rien modifier.
