# Exercice 3.2 - Questions de réflexion

## 1. Avantage de l'implémentation implicite des interfaces

Un type satisfait une interface du simple fait qu'il possède les bonnes
méthodes, sans déclaration ni dépendance explicite. On peut donc faire
satisfaire une interface à un type qu'on ne contrôle pas, et définir les
interfaces côté consommateur (là où le besoin existe) plutôt que côté
producteur. Résultat : un découplage fort et moins de couplage de compilation.

## 2. Quand utiliser `interface{}` (any), et ses inconvénients

Utile pour des conteneurs hétérogènes ou des API génériques (`encoding/json`,
`fmt`, etc.). Inconvénients : perte de la vérification de type à la compilation,
nécessité d'assertions / type switches au runtime (risque de `panic`), et code
moins lisible. Depuis Go 1.18, les génériques sont souvent préférables quand les
types manipulés sont homogènes.

## 3. Interfaces, modularité et testabilité

Les interfaces définissent des contrats de comportement : le code dépend
d'abstractions et non d'implémentations concrètes, ce qui le rend modulaire et
extensible. Pour les tests, on injecte des implémentations factices
(mocks/stubs) qui satisfont l'interface, sans toucher au code de production.
