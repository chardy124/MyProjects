package particles

//import ("math/rand";"time";"project-particles/config";"log";"math")
import "fmt"
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	fmt.Println(len(s.Content))
}