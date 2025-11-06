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
5. Quitter

Par défaut, deux utilisateurs sont créés :
| ID   | Nom       | Email               |
|------|-----------|---------------------|
| 1    | Alice     | alice@example.com   |
| 2    | Bob       | bob@example.com     |

## Installation et lancement

Cloner le projet :
 ```
git clone https://github.com/vcircosta/GO-TP1.git
cd GO-TP1
```

Lancer le projet :
```
go run ./cmd/crm
```

Lancer les tests du mini-crm :

```
go test ./internal/storage -v
```

## Initalisation de cobra

Installer les changements
```
go mod tidy
```

Exemples d'utilisation
```
go run ./cmd/crm add "John Doe" "john@example.com"
go run ./cmd/crm list
go run ./cmd/crm update 1 "John Smith" "john.smith@example.com"
go run ./cmd/crm delete 1
```
