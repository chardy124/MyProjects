package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	InitSizeParticles		 float64
	InitVitesseParticles	 float64
	RandomSpawn              bool
	SpawnX, SpawnY           int
	SpawnRate                float64
	Gravite					 float64
	Gravitation				 float64
	Merge					 bool
	ExterieurDeLecranXmin	 float64
	ExterieurDeLecranXmax	 float64
	ExterieurDeLecranYmin	 float64
	ExterieurDeLecranYmax	 float64
	Rebond_bords			 bool
	Rebond_particules		 bool
	Acceleration			 float64
	Duree_Vie				 float64
	Nombre_Particules_Max	 float64
	Attraction			 	 float64
	AttractionX			 	 float64
	AttractionY			 	 float64
	Grossissement		 	 float64
}

var General Config