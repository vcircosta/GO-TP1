# Mini-CRM en Go (CLI)

Un petit CRM (Customer Relationship Management) en ligne de commande développé en Go.
Il permet de gérer des contacts avec des fonctionnalités simples : ajout, listing, mise à jour et suppression.

## Membres du groupe

- Elise LABARRERE
- Valentin CIRCOSTA

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
    - `go run ./cmd/crm add`
        - Il sera demandé de renseigner un nom et un email pour ce nouvel utilisateur
    - `go run ./cmd/crm add "John Doe" "john@example.com"`
        - **Paramètre 1** : nom du contact
        - **Paramètre 2** : email du contact
2. Lister tous les contacts
    - `go run ./cmd/crm list`
        - Cette commande ne prend aucun paramètre.
3. Mettre à jour un contact
    - `go run ./cmd/crm update`
        - La liste des utilisateurs existants sera affiché il faudra choisir l'id de  l'utilisateur à modifier.
        - Puis il sera demandé de renseigner le nouveau nom et le nouvel email
    - `go run ./cmd/crm update 1 "John Smith" "john.smith@example.com"`
        - **Paramètre 1** : ID du contact à modifier
        - **Paramètre 2** : nouveau nom
        - **Paramètre 3** : nouvel email
4. Supprimer un contact
    - `go run ./cmd/crm delete`
        - La liste des utilisateurs existants sera affiché il faudra choisir l'id de  l'utilisateur à supprimer.
    - `go run ./cmd/crm delete 1`
        - **Paramètre 1** : ID du contact à supprimer
5. Changer le stockage par défaut
    - `go run ./cmd/crm set-storage`
        - Un choix sera demandé entre json, gorm, ou memory.
    - `go run ./cmd/crm set-storage gorm`
        - **Paramètre 1** : Change le type de stockage par défaut (json, gorm, ou memory).

Si une commande comporte un **nombre incorrect de paramètres**, une erreur sur le nombre d'arguments est affiché suivi des utilisations corrects de la commande (avec et dans paramètres).

### Flags

| Option            | Description                                                                    |
| ----------------- | ------------------------------------------------------------------------------ |
| `--data-dir`      | Répertoire où stocker les données (`data` par défaut).                         |
| `--file-name`     | Nom du fichier (sans extension) pour JSON ou SQLite (`contacts` par défaut).   |
| `--storage`       | Type de stockage par défaut : `json`, `gorm`, ou `memory` (`json` par défaut). |
| `--verbose`, `-v` | Mode verbeux : affiche plus de détails lors de l’exécution.                    |

## Tests

Lancer les tests du mini-CRM (stockage en mémoire + stockage JSON) :

```
go test ./internal/storage -v
```

## Branches Git
| Branche     | Description                                                                                                                            |
| ----------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| **tp1**     | Programme simple d'un menu de crm.                                                                                                     |
| **tp2**     | Réorganisation des méthodes (`add`, `update`, `delete`) avec séparation des gestionnaires (`handle`) et des fonctions principales.     |
| **tp2-bis** | Introduction des **interfaces** pour rendre le programme plus modulaire, évolutif et maintenable.                                      |
| **tp3**     | Ajout de la **persistance des données en JSON** (la variable `memoryUseJson` dans `main` permet de choisir entre JSON et mémoire).     |
| **tp3-bis** | Transformation du programme en **CLI avec Cobra** (réorganisation des fichiers + ajout des sous-commandes pour les 4 fonctionnalités). |
| **tp-4**    | Ajout du storer **GORM** et d'une commande set-storage pour changer le stockage par défaut.                                            |
