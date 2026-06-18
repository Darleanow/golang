# Réponses aux questions - TP 4a Goroutines

## Exercice 1 - Sortie prématurée

**Que constate-t-on ? Toutes les goroutines terminent-elles avant l'arrêt du programme ?**

Sans synchronisation, `main` lance les 5 goroutines puis continue son exécution
immédiatement et affiche « Toutes les goroutines lancées. ». Or, quand `main`
se termine, **tout le programme s'arrête**, y compris les goroutines encore en
cours : celles qui dorment encore dans `time.Sleep` sont tuées avant d'afficher
« tâche terminée ». On voit donc souvent moins de messages de fin que de début,
et l'ordre varie d'une exécution à l'autre.

Dans le code, un `time.Sleep(600 ms)` final masque le problème en laissant le
temps aux goroutines de finir - mais ce n'est qu'un bricolage : la durée est
devinée, pas garantie. La vraie solution est le `WaitGroup` (Exercice 2).

## Exercice 2 - sync.WaitGroup

**Le comportement a-t-il changé ? Toutes les goroutines terminent-elles ?**

Oui. `wg.Add(1)` avant chaque lancement et `defer wg.Done()` dans la goroutine
permettent à `wg.Wait()` de **bloquer `main`** tant que le compteur n'est pas
revenu à zéro. Toutes les goroutines vont donc jusqu'au bout avant l'affichage
final, de façon déterministe (plus de course avec la fin du programme).

## Exercice 3 - Canaux

**Quel est l'ordre des résultats ? Suit-il l'ordre des IDs ?**

Non. Les goroutines s'exécutent en parallèle et dorment une durée **aléatoire** ;
elles finissent donc dans un ordre imprévisible. Le canal délivre les messages
dans leur **ordre d'arrivée** (ordre d'achèvement), qui ne correspond pas à
l'ordre des IDs 1→5.

Détail important : le canal est **bufferisé** (capacité 5). Comme on appelle
`wg.Wait()` *avant* de lire le canal, un canal non-bufferisé provoquerait un
interblocage - les goroutines resteraient bloquées sur l'envoi (personne ne lit
encore), donc `wg.Done()` ne serait jamais atteint et `wg.Wait()` attendrait
indéfiniment. Avec un buffer suffisant, les envois passent sans lecteur immédiat.

## Exercice 4 - Pool de travailleurs

**Ordre de traitement et effet du nombre de travailleurs sur le temps total ?**

Les 10 tâches sont réparties entre les 3 travailleurs qui les consomment au fur
et à mesure depuis le canal `taches`. L'ordre de traitement et d'affichage est
non déterministe : il dépend de quel travailleur est libre et de la durée
aléatoire de chaque tâche.

Plus il y a de travailleurs, plus de tâches s'exécutent **en parallèle**, donc
le temps total diminue - jusqu'à une limite : au-delà du nombre de tâches (ou du
nombre de cœurs `GOMAXPROCS`), ajouter des travailleurs n'accélère plus rien.
Le pool borne aussi la concurrence, ce qui évite de lancer 10 000 goroutines
d'un coup pour 10 000 tâches.
