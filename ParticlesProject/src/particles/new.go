package particles

import ("project-particles/config";"math/rand";"time")

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	var WindowSizeX int = config.General.WindowSizeX
	var WindowSizeY int = config.General.WindowSizeY
	var particules []Particle
	rand.Seed(time.Now().UnixNano())
	var nbrScreens int = 11
	for i := 0 ; i < config.General.InitNumParticles ; i++ {
		particules = append(particules, Particle{
			PositionX: rand.Float64() * float64(WindowSizeX*nbrScreens + WindowSizeX*nbrScreens/2),
			PositionY: rand.Float64() * float64(WindowSizeY*nbrScreens + WindowSizeY*nbrScreens/2),
			Rotation: 0,
			ScaleX: 0.1,
			ScaleY: 0.1,
			ColorRed: 30.0,
			ColorGreen: 30.0,
			ColorBlue: 30.0,
			Opacity: 1.0,
			Duree_Vie: -1.0,
		})
	}


	return System{Content: particules, reste: 0}
	/*rand.Seed(time.Now().UnixNano())
	var nombre_particules int = config.General.InitNumParticles
	if config.General.RandomSpawn{
		for i := 0; i < nombre_particules; i++ {
			var x float64 = rand.Float64()* float64(config.General.WindowSizeX)
			var y float64 = rand.Float64()* float64(config.General.WindowSizeY)
			var taille float64 = config.General.InitSizeParticles
			var vitesse float64 = config.General.InitVitesseParticles
			particules = ajout(particules,x,y,taille,vitesse)
		}
	}else{
		var x float64 = float64(config.General.SpawnX)
		var y float64 = float64(config.General.SpawnY)
		var taille float64 = rand.Float64()*2config.General.InitSizeParticles
		var vitesse float64 = config.General.InitVitesseParticles
		for i := 0; i < nombre_particules; i++ {
			particules = ajout(particules,x,y,taille,vitesse)
		}
	}*/
}