# Réponses et remarques - TP 5a API REST net/http

## Gestion des routes avec net/http

Le `ServeMux` de la stdlib ne gère pas les paramètres d'URL nativement. Pour
distinguer `/items` (collection) de `/items/{id}` (élément précis), j'ai
enregistré deux patterns : `/items` (exact) et `/items/` (avec slash final,
qui matche tout ce qui commence par ce préfixe). L'ID est ensuite extrait
manuellement avec `strings.TrimPrefix`. C'est un peu verbeux mais ça évite
d'importer un routeur externe pour un truc aussi simple.

## Choix du mutex

J'ai utilisé un `sync.Mutex` simple. J'aurais pu mettre un `sync.RWMutex`
pour laisser passer plusieurs lectures en parallèle, mais sur cette appli
en mémoire la différence est vraiment négligeable.

Le truc qui m'a pris du temps : au début j'appelais `sendJSON` encore sous
le lock dans le PUT, et j'avais une deadlock sur certaines requêtes parce que
la goroutine HTTP restait bloquée à écrire la réponse pendant que le mutex
était tenu. J'ai corrigé en libérant le mutex avant d'appeler `sendJSON`.

## Génération des IDs

J'utilise `github.com/google/uuid` pour les IDs. J'ai d'abord essayé avec
un simple compteur global mais ça posait des problèmes de concurrence donc
autant prendre une lib propre.

## Tests manuels

```bash
# Lister tous les items
curl http://localhost:8080/items

# Récupérer un item par ID
curl http://localhost:8080/items/1

# Créer un item
curl -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Écran 4K","description":"27 pouces, 144Hz"}'

# Mettre à jour
curl -X PUT http://localhost:8080/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Clavier mécanique v2"}'

# Supprimer
curl -X DELETE http://localhost:8080/items/1
```

## Limites

- Tout est en mémoire, redémarrer le serveur efface tout
- Pas de pagination sur le GET /items
- Le PUT accepte un body vide sans renvoyer d'erreur
