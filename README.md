# Mini-CRM en Go (CLI)

Un petit CRM (Customer Relationship Management) en ligne de commande développé en Go.
Il permet de gérer des contacts avec des fonctionnalités simples : ajout, listing, mise à jour et suppression.

## Membres du groupe

Elise LABARRERE
Valentin CIRCOSTA

## Fonctionnalités

Lors du lancement, 3 actions possibles :
1. Ajouter un contact
2. Lister tous les contacts
3. Mettre à jour un contact
4. Supprimer un contact

Par défaut, deux utilisateurs sont créés :
| ID   | Nom       | Email               |
|------|-----------|---------------------|
| 1    | Alice     | alice@example.com   |
| 2    | Bob       | bob@example.com     |

## Installation 

Cloner le projet :
 ```
git clone https://github.com/vcircosta/GO-TP1.git
cd GO-TP1
```

Installer les dépendances :
```
go mod tidy
```

## Utilisation du CRM

### Mode interactif (standard)

Afficher le menu principal :
```
go run ./cmd/crm
```

### Utilisation via les commandes

1. Ajouter un contact
    - `go run ./cmd/crm add "John Doe" "john@example.com"`
    - **Paramètre 1** : nom du contact
    - **Paramètre 2** : email du contact
2. Lister tous les contacts
    - `go run ./cmd/crm list`
    - Cette commande ne prend aucun paramètre.
3. Mettre à jour un contact
    - `go run ./cmd/crm update 1 "John Smith" "john.smith@example.com"`
    - **Paramètre 1** : ID du contact à modifier
    - **Paramètre 2** : nouveau nom
    - **Paramètre 3** : nouvel email
4. Supprimer un contact
    - Exemple avec paramètre : `go run ./cmd/crm delete 1`
    - **Paramètre 1** : ID du contact à supprimer

**Comportement en cas d’exécution :**
- Si une commande est saisie **sans paramètres**, le programme lance la fonction `handle` correspondante, ce qui affiche le menu interactif (fonctionnement normal).
- Si une commande comporte un **nombre incorrect de paramètres**, une erreur sur le nombre d'arguments est affiché suivi des utilisations corrects de la commande (avec et dans paramètres).

### Flags

| Option            | Description                                                                   |
| ----------------- | ----------------------------------------------------------------------------- |
| `--data-dir`      | Répertoire où stocker les données (`data` par défaut).                        |
| `--json`          | Active ou non le stockage en fichier JSON (`true` = JSON, `false` = mémoire). |
| `--verbose`, `-v` | Mode verbeux : affiche plus de détails lors de l’exécution.                   |

## Tests

Lancer les tests du mini-CRM (stockage en mémoire + stockage JSON) :

```
go test ./internal/storage -v
```

## Branches Git
| Branche     | Description                                                                                                                            |
| ----------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| **tp1**     | Programme simple sans persistance (utilise des `map`).                                                                                 |
| **tp2**     | Réorganisation des méthodes (`add`, `update`, `delete`) avec séparation des gestionnaires (`handle`) et des fonctions principales.     |
| **tp2-bis** | Introduction des **interfaces** pour rendre le programme plus modulaire, évolutif et maintenable.                                      |
| **tp3**     | Ajout de la **persistance des données en JSON** (la variable `memoryUseJson` dans `main` permet de choisir entre JSON et mémoire).     |
| **tp3-bis** | Transformation du programme en **CLI avec Cobra** (réorganisation des fichiers + ajout des sous-commandes pour les 4 fonctionnalités). |
