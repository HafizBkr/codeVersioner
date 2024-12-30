// cmd/help.go
package cmd

import "fmt"

func ShowHelp() {
    help := `
Usage: vcs <command> [options]

Commands:
    start               Initialise un nouveau repository
    keep <file>        Ajoute un fichier à la staging area
    save <message>     Crée un nouveau commit avec les fichiers stagés
    timeline           Affiche l'historique des commits
    status             Affiche l'état actuel du repository
    branch <name>      Crée une nouvelle branche
    help               Affiche ce message d'aide

Examples:
    vcs start
    vcs keep file.txt
    vcs status
    vcs save "Initial commit"
    vcs branch develop
    vcs timeline
`
    fmt.Println(help)
}
