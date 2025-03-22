# MyVCS - Système de Gestion de Versions

**MyVCS** est un système de gestion de versions léger, permettant de suivre les modifications, de gérer les commits, et de travailler avec des branches. Ce projet est conçu pour être intuitif tout en offrant des fonctionnalités essentielles pour le contrôle de versions.

---

## Fonctionnalités Principales

- **Initialisation de Dépôt** : Configure un dépôt pour le suivi de versions.
- **Ajout de Fichiers** : Sélectionne les fichiers à inclure dans le prochain commit.
- **Commit des Changements** : Sauvegarde des modifications avec un message descriptif.
- **Historique des Commits** : Visualise la chronologie des commits. 
- **Gestion des Branches** : Liste et crée des branches.
- **Clonage de Dépôt** : Clone un dépôt local existant.
- **Statut** : Montre l'état actuel du dépôt.
- **Aide** : Affiche les commandes disponibles.

---

## Installation

1. **Clonez le dépôt** :
   ```
   git clone https://github.com/votreutilisateur/myvcs.git
   cd myvcs
   ```

2. **Compilez le projet** :
   Assurez-vous que Go est installé sur votre système.
   ```
   go build -o vcs main.go
   ```

3. **Exécutez l'application** :
   ```
   ./vcs
   ```

---

## Commandes Disponibles

### 1. Initialisation de Dépôt
```
./vcs start
```
Crée un répertoire `.vcs` pour initialiser un nouveau dépôt.

### 2. Ajouter des Fichiers
```
./vcs keep <file>
```
Ajoute un fichier spécifique au suivi pour le prochain commit.

### 3. Sauvegarder un Commit
```
./vcs save "<message>"
```
Sauvegarde les changements en y associant un message descriptif.

### 4. Afficher l'Historique des Commits
```
./vcs timeline
```
Affiche une liste de tous les commits.

### 5. Afficher le Statut
```
./vcs status
```
Montre les fichiers modifiés, ajoutés ou suivis.

### 6. Gestion des Branches
- **Lister les branches** :
  ```
  ./vcs branch --list
  ```
- **Créer une nouvelle branche** :
  ```
  ./vcs branch --create <name>
  ```

### 7. Cloner un Dépôt Local
```
./vcs clone <local_path>
```
Clone un dépôt local existant dans un nouveau répertoire.

### 8. Afficher l'Aide
```
./vcs help
```
Affiche la liste des commandes disponibles et leur usage.

---


## Exemple d'Utilisation

1. **Initialiser un dépôt** :
   ```
   ./vcs start
   ```

2. **Ajouter un fichier** :
   ```
   ./vcs keep README.md
   ```

3. **Sauvegarder un commit** :
   ```
   ./vcs save "Initial commit : Ajout du README"
   ```

4. **Afficher l'historique** :
   ```
   ./vcs timeline
   ```

5. **Créer une branche** :
   ```
   ./vcs branch --create feature-docs
   ```

---

## Contribution

Les contributions sont encouragées ! Voici comment contribuer :

1. **Forkez le projet**.
2. **Créez une branche** pour vos modifications :
   ```
   git checkout -b feature/nom-de-la-fonctionnalité
   ```
3. **Commitez vos modifications** :
   ```
   git commit -m "Ajout de la fonctionnalité X"
   ```
4. **Poussez votre branche** :
   ```
   git push origin feature/nom-de-la-fonctionnalité
   ```
5. **Créez une Pull Request** sur le dépôt principal.



## Auteur

**Votre Nom**  
- GitHub : ((https://github.com/HafizBkr))  
- Email : hafizinovus@gmail.com

N'hésitez pas à poser des questions ou à proposer des améliorations. 🚀

