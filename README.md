# Mini-CRM en Go (CLI)

Un petit CRM (Customer Relationship Management) en ligne de commande développé en Go.
Il permet de gérer des contacts avec des fonctionnalités simples : ajout, mise à jour, suppression, listing, et ajout via flags.

## Membres du groupe

Elise LABARRERE
Valentin CIRCOSTA

## Fonctionnalités

Afficher un menu interactif en boucle

Ajouter un contact (ID, Nom, Email)

Lister tous les contacts

Supprimer un contact par son ID

Mettre à jour un contact

Quitter l’application


## Concepts Go utilisés

Boucle infinie for {} pour le menu

Switch pour le choix utilisateur

Map pour stocker les contacts (map[int]Contact)

Gestion des erreurs avec if err != nil

Conversion de string en int avec strconv.Atoi

Lecture de l’entrée utilisateur avec bufio.Reader et os.Stdin

Flags avec le package flag pour l’ajout rapide


## Installation et lancement

Cloner le projet
 ```
git clone https://github.com/vcircosta/GO-TP1.git
cd GO-TP1
```

Lancer le projet
```
go run .
```

Lancer les tests

```
go test
```
