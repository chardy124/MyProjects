package particles
import ("testing";"project-particles/config";"fmt";"math")


func Test_Vide_Creation_Particules(t *testing.T){
	var systeme System = NewSystem()//création d'un nouveau système (avec aucune particule)
	if len(systeme.Content) != 0{//s'il y a au moins une particule...
		t.Error("NewSystem est vide.")
	}
}

func Test_Init_Particule(t *testing.T){
	config.General.InitNumParticles = 5//initialisation du nombre de particule à afficher au début à 5
	var systeme System = NewSystem()//création d'un nouveau système prenant les informations de NewSystem à son execution
	if len(systeme.Content) != 5{//si le nombre de particule du système est différent de 5...
		t.Error("Le nombre de particules est incorrect")
	}
}

func Test_Random_Spawn(t *testing.T){
	config.General.RandomSpawn = false//établir l'apparition aléatoire en position des particule comme désactivé
	config.General.SpawnX = 400//initialiser la position X à apparaitre des particules à 400 pixels de la gauche de la fenêtre 
	config.General.SpawnY = 300//initialiser la position Y à apparaitre des particules à 300 pixels du haut de la fenêtre
	var systeme System = NewSystem()//création d'un nouveau système prenant les informations de NewSystem à son execution
	for i := 0; i<len(systeme.Content); i++{//pour toutes les particules générées par la fonction NewSystem...
		if systeme.Content[i].PositionX != 400 || systeme.Content[i].PositionY != 300{//...si sa position est incorecte par rapport aux attente X et Y...
			t.Error("Ce n'est pas le bon endroit d'apparition")
		}
	}
}

func Test_deplacement(t *testing.T){
	config.General.InitVitesseParticles = 2
	var particules_temoin []Particle = NewSystem().Content//création d'un tableau de particules témoin qui servira de comparaison à celui auquel on appliquera la fonction
	var particules_test []Particle = deplacement(NewSystem().Content,0)//création d'un tableua de particules de test qui sert à se faire comparer à celui temoin afin de vérifier les changement effectué après application de la fonction
	for i := 0; i < len(particules_test); i++ {//pour toutes les particules du tableau test...
		if particules_test[i].PositionX == particules_temoin[i].PositionX &&//si la position X de la particule est égale à celle de correspondante du tableau temoin...
		particules_test[i].PositionY == particules_temoin[i].PositionY{//..et si la position Y de la particule est égale à celle de correspondante du tableau temoin...
			t.Error("La position de la particule",i," n'a pas été changée.")
		}
	} 
}

func Test_deplacement_v4(t *testing.T){
	var PositionX float64 = 400//initialiser la position X à apparaitre des particules à 400 pixels de la gauche de la fenêtre 
	var PositionY float64 = 300//initialiser la position Y à apparaitre des particules à 300 pixels du haut de la fenêtre
	var particules []Particle = ajout([]Particle{},PositionX,PositionY,1,1)//création d'un tableau de particules auquel on en ajoute une avec une vitesse non-nulle
	particules = deplacement(particules,0)//déplacement de la particule grâce à sa vitesse
	if particules[0].PositionX == 400 && particules[0].PositionY == 300{//si les coordonnées de la particule sont les mêmes qu'à sa création...
		t.Error("Les particules n'ont bougées.")
	}
}

func Test_suppression_Position_X_min(t *testing.T){
	var PositionX float64 = -50//initialisation d'une position de coordonnées X = -50, donc en dehors, à gauche, de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle
	var systeme System
	systeme.Content = particules_test
	var libres int
	particules_test,libres = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100, &systeme, 0)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if libres != 1{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_X_max(t *testing.T){
	var PositionX float64 = float64(800) + 50//initialisation d'une position de coordonnées X = (taille de la fenêtre dans ce contexte) + 50, donc en dehors, à droite, de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle
	var systeme System
	systeme.Content = particules_test
	var libres int
	particules_test,libres = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100, &systeme, 0)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if libres != 1{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_Y_min(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = -50//et d'une coordonnée Y = -50 donc en dehors, au dessus, de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle
	var systeme System
	systeme.Content = particules_test
	var libres int
	particules_test,libres = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100, &systeme, 0)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if libres != 1{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_Y_max(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = float64(600) + 50//et d'une coordonnée Y = 50 donc en dehors, en dessous, de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle 
	var systeme System
	systeme.Content = particules_test
	var libres int
	particules_test,libres = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100, &systeme, 0)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if libres != 1{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Taille_min(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 2, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations avec comme particularité une taille inférieure à l'exigence qui suit
	var particules_test []Particle
	var systeme System
	systeme.Content = particules_test
	var libres int
	particules_test,libres = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 5, 100, &systeme, 0)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être trop petite
	if libres != 1{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Taille_max(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 150, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations avec comme particularité une taille supérieure à l'exigence qui suit
	var particules_test []Particle
	var systeme System
	systeme.Content = particules_test
	var libres int
	particules_test,libres = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100, &systeme, 0)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être trop grande
	if libres != 1{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_grossissement_Additif_Positif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle
	particules_test = grossissement(particules_temoin, false, 10,0)//application de la fonction de grossissement tel que la taille doit être incrémentée de 10 pixels
	if particules_test[0].ScaleX != float64(100)+10{//si la taille n'est pas celle attendue (110) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_grossissement_Additif_Negatif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, false, -10,0)//application de la fonction de grossissement tel que la taille doit être décrémentée de 10 pixels
	if particules_test[0].ScaleX != float64(100)-10{//si la taille n'est pas celle attendue (90) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas maigri.")
	}
}

func Test_grossissement_Multiplicatif_Positif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, true, 1.1,0)//application de la fonction de grossissement tel que la taille doit être multipliée par 1,1
	if particules_test[0].ScaleX != float64(100)*1.1{//si la taille n'est pas celle attendue (110) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_grossissement_Multiplicatif_Negatif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, true, 0.9,0)//application de la fonction de grossissement tel que la taille doit être multipliée par 0,9
	if particules_test[0].ScaleX != float64(100)*0.9{//si la taille n'est pas celle attendue (90) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_Ajout_Creation_1(t *testing.T){
	if len(ajout([]Particle{},0,0,0,0)) != 1{//si aucune particule ne se crée avec la fonction...
		t.Error("Aucune particule créée.")
	}
}

func Test_Ajout_Creation_2(t *testing.T){
	var test []Particle = ajout([]Particle{},0,0,0,0)//création d'une particule avec la fonction ajout avec les paramètres d'entrée autres que le tableau nuls
	if test[0].PositionX != 0{//si sa position X est non-nulle...
		t.Error("X différent de 0.")}
	if test[0].PositionY != 0{//si sa position Y est non-nulle...
		t.Error("Y différent de 0.")}
	if test[0].Rotation != 0{//si sa rotation est non-nulle...
		t.Error("Rotation non nulle")}
	if test[0].ScaleX != 0{//si sa taille X est non-nulle...
		t.Error("Taille X non nulle.")}
	if test[0].ScaleY != 0{//si sa taille Y est non-nulle...
		t.Error("Taille Y non nulle.")}
}

func Test_Ajout_Creation_3(t *testing.T){
	var test []Particle = ajout([]Particle{},10,3,4,7)//création d'une particule avec la fonction ajout avec les paramètres d'entrée autres que le tableau non-nuls
	if test[0].PositionX != 10{//si sa position X n'est pas égale à celle demandée...
		t.Error("X différent de 10.")}
	if test[0].PositionY != 3{//si sa position Y n'est pas égale à celle demandée...
		t.Error("Y différent de 3.")}
	if test[0].Rotation != 0{//si sa Rotation n'est pas nulle...
		t.Error("Rotation non nulle")}
	if test[0].ScaleX != 4{//si sa taille X n'est pas égale à celle demandée...
		t.Error("Taille X différente de 7.")}
	if test[0].ScaleY != 4{//si sa taille Y n'est pas égale à celle demandée...
		t.Error("Taille Y différente de 7.")}
}

func Test_Ajout_1(t *testing.T){
	var test []Particle = ajout(NewSystem().Content,10,3,4,7)//création d'un tableau de particules test ayant une particule de plus que le tableau contenu dans la sortie de NewSystem (cela a pour but de tester la fonction avec un tableau non-vide en entrée)
	if len(test) != len(NewSystem().Content)+1{//si la particule n'a pas été ajoutée
		t.Error("La particule n'a été ajoutée.")
	}
	if test[len(test)-1].PositionX != 10{//si sa position X n'est pas égale à celle demandée...
		t.Error("X différent de 10.")}
	if test[len(test)-1].PositionY != 3{//si sa position Y n'est pas égale à celle demandée...
		t.Error("Y différent de 3.")}
	if test[len(test)-1].Rotation != 0{//si sa Rotation n'est pas nulle...
		t.Error("Rotation non nulle")}
	if test[len(test)-1].ScaleX != 4{//si sa taille X n'est pas égale à celle demandée...
		t.Error("Taille X différente de 4.")}
	if test[len(test)-1].ScaleY != 4{//si sa taille Y n'est pas égale à celle demandée...
		t.Error("Taille Y différente de 4.")}
	if (test[len(test)-1].SpeedX)*(test[len(test)-1].SpeedX)+(test[len(test)-1].SpeedY)*(test[len(test)-1].SpeedY) >= float64(7*7)+1e-12 || 
		(test[len(test)-1].SpeedX)*(test[len(test)-1].SpeedX)+(test[len(test)-1].SpeedY)*(test[len(test)-1].SpeedY) <= float64(7*7)-1e-12{//si sa vitesse (théorème de Pythagore) n'est pas égale à celle demandée (donnée tel le module d'un nombre complexe SpeedX + i*SpeedY)...
		t.Error("Speed X différente de 7.",(test[len(test)-1].SpeedX)*(test[len(test)-1].SpeedX)+(test[len(test)-1].SpeedY)*(test[len(test)-1].SpeedY))}
}

func Test_Ajout_2(t *testing.T){
	var systeme System//création d'un nouveau système de particules pour tester la fonction ajout dans un système de particules
	systeme.Content = ajout(systeme.Content,10,3,4,7)//ajout de la particule dans le tableau de particules du système
	if len(systeme.Content) != 1{//si la particule n'a pas été ajoutée...
		t.Error("Il n'y a pas le bon nombre de particules crées")
	}
}

func Test_Creation_Particule_New_System(t *testing.T){
	var systeme System = NewSystem()//création d'un nouveau système de particules prenant celui de sortie de NewSystem
	if len(systeme.Content) == 0{//si aucune particule n'a été ajoutée...
		t.Error("La particule n'a été ajoutée")
	}
}

func Test_Spawn_Rate1(t *testing.T){
	config.General.Nombre_Particules_Max = -1
	config.General.SpawnRate = 1//initialisation du taux d'apparition
	var systeme System//création d'un nouveau système de particules
	systeme.Update()//actialisation du système de particules
	if len(systeme.Content) != 1{//si la particule n'a pas été ajoutée ou que trop en ont été ajoutées...
		t.Error("Le taux d(apparition est mauvais")
	}
}

func Test_Rebond(t *testing.T){
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	var systeme System
	var particule Particle
	systeme.Content = append(systeme.Content, particule)
	systeme.Content[0].SpeedX = -1000
	systeme.Content[0].PositionX = 400
	systeme.Content[0].PositionY = 300
	systeme.Update()	
	for i := 0; i < len(systeme.Content); i++ {
		if systeme.Content[i].PositionX > 800{
			t.Error("La particule sort de la fenêtre donc elle ne rebondie pas")
		}
		if systeme.Content[i].PositionX < 0{
			t.Error("La particule sort de la fenêtre donc elle ne rebondie pas")
		}
		if systeme.Content[i].PositionY > 600{
			t.Error("La particule sort de la fenêtre donc elle ne rebondie pas")
		}
		if systeme.Content[i].PositionY < 0{
			t.Error("La particule sort de la fenêtre donc elle ne rebondie pas")
		}
	}
}

func Test_Gravite(t *testing.T){
	config.General.Gravite = 6
	var systeme System
	var particule Particle
	systeme.Content = append(systeme.Content,particule)
	systeme.Content[0].SpeedY = -5
	systeme.Content[0].PositionX = 400
	systeme.Content[0].PositionY = 300
	systeme.Content = gravite(systeme.Content,0)
	for i := 0; i < len(systeme.Content); i++ {
		if systeme.Content[i].PositionY > 300{
			t.Error("La particule n'est pas soumise à la gravitée")
		}
		if systeme.Content[i].SpeedY < 0{
			t.Error("La particule monte")
		}
	}
}

func Test_abs(t *testing.T){
	if abs(-5) != 5{
		t.Error("Ce n'est pas la valeur absolue")
	}
}

func Test_ProcheX1(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule2.PositionX = 310
	particule1.ScaleX = 2
	particule2.ScaleX = 2
	if procheX(particule1,particule2) != true{
		t.Error("La particule est proche de la deuxième mais ça n'est pas indiqué")
	}
}

func Test_ProcheX2(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule2.PositionX = 350
	particule1.ScaleX = 2
	particule2.ScaleX = 2
	if procheX(particule1,particule2) != false{
		t.Error("La particule est proche de la deuxième mais ça n'est pas indiqué")
	}
}

func Test_ProcheY1(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionY = 300
	particule2.PositionY = 310
	particule1.ScaleY = 2
	particule2.ScaleY = 2
	if procheY(particule1,particule2) != true{
		t.Error("La particule est loin de la deuxième mais ça n'est pas indiqué")
	}
}

func Test_ProcheY2(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionY = 300
	particule2.PositionY = 350
	particule1.ScaleY = 2
	particule2.ScaleY = 2
	if procheY(particule1,particule2) != false{
		t.Error("La particule est proche de la deuxième mais ça n'est pas indiqué")
	}
}

func Test_collision1(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule2.PositionX = 310
	particule1.ScaleX = 2
	particule2.ScaleX = 2
	particule1.PositionY = 300
	particule2.PositionY = 310
	particule1.ScaleY = 2
	particule2.ScaleY = 2
	if !procheX(particule1,particule2) || !procheY(particule1,particule2){
		t.Error("Il y a collision mais ça n'est pas indiqué")
	}
}

func Test_collision2(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule2.PositionX = 350
	particule1.ScaleX = 2
	particule2.ScaleX = 2
	particule1.PositionY = 300
	particule2.PositionY = 350
	particule1.ScaleY = 2
	particule2.ScaleY = 2
	if procheX(particule1,particule2) && procheY(particule1,particule2){
		t.Error("Il n'y a pas de collision mais ça n'est pas indiqué")
	}
}


func Test_collision3(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 290
	particule1.PositionY = 190
	particule2.PositionX = 300
	particule2.PositionY = 200
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	if !collision(particule1,particule2){
		t.Error("La particule1 est censée être en collision avec un x inférieur et un y supérieur à la particule2")
	}
}

func Test_collision4(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 280
	particule1.PositionY = 180
	particule2.PositionX = 300
	particule2.PositionY = 200
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	if !collision(particule1,particule2){
		t.Error("La particule1 est censée être en collision avec le coin inférieur droit de la particule2")
	}
}

func Test_collision5(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule1.PositionY = 180
	particule2.PositionX = 300
	particule2.PositionY = 200
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	if !collision(particule1,particule2){
		t.Error("La particule1 est censée être en collision avec le segment inférieur de la particule2")
	}
}

func Test_collision6(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 290
	particule1.PositionY = 190
	particule2.PositionX = 290
	particule2.PositionY = 190
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	if !collision(particule1,particule2){
		t.Error("La particule1 est censée être en collision avec la aprticule2 car même coordonnées")
	}
}

func Test_rebond_particules1(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule1.PositionY = 180
	particule1.SpeedX = 2
	particule1.SpeedY = 2
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.PositionX = 300
	particule2.PositionY = 200
	particule2.SpeedX = 3
	particule2.SpeedY = 3
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	if !collision(particule1,particule2) && particule1.SpeedX != 3 && particule1.SpeedY != 3 &&
	particule2.SpeedX != 2 && particule2.SpeedY != 2 {
		t.Error("Les particules n'ont pas rebondis")
	}
}

func Test_rebond_particules2(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 300
	particule1.PositionY = 310
	particule1.SpeedX = 2
	particule1.SpeedY = 2
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.PositionX = 300
	particule2.PositionY = 311
	particule2.SpeedX = 3
	particule2.SpeedY = 3
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	var systeme System
	systeme.Content = append(systeme.Content,particule1)
	systeme.Content = append(systeme.Content,particule2)
	systeme.Content = rebond_particules(systeme.Content,0)
	if systeme.Content[0].SpeedX != 3 && systeme.Content[0].SpeedY !=3 &&
	systeme.Content[1].SpeedX != 2 && systeme.Content[1].SpeedY != 2{
		t.Error("Les particules ne rebondissent pas")
	}
}

func Test_acceleration(t *testing.T){
	config.General.Acceleration = 2
	var particule Particle
	particule.SpeedX = 5
	particule.SpeedY = 5
	particule = acceleration(particule)
	fmt.Println()
	if particule.SpeedX != 10 && particule.SpeedY != 10{
		t.Error("La particule n'a pas acceléré")
	}
}

func Test_attraction1(t *testing.T){
	var particule Particle
	particule.PositionX = 200
	particule.PositionY = 200
	particule.SpeedX = 2
	particule.SpeedY = 2
	particule = attraction(particule,1, 0, true, 400, 400)
	if particule.SpeedX <= 2 && particule.SpeedY <= 2{
		t.Error("Il n'y a pas d'accelération",particule.SpeedX, particule.SpeedY)
	}
}

func Test_attraction2(t *testing.T){
	var particule Particle
	particule.PositionX = 200
	particule.PositionY = 200
	particule.SpeedX = 3
	particule.SpeedY = 3
	particule = attraction(particule, 2, 0, false, 400, 400)
	if particule.SpeedX <= 2.99999999 && particule.SpeedY <=2.99999999{
		t.Error("Il n'y a pas d'accelération",particule.SpeedX, particule.SpeedY)
	}
}

func Test_attraction3(t *testing.T){
	var particule Particle
	particule.PositionX = 200
	particule.PositionY = 200
	particule.SpeedX = 3
	particule.SpeedY = 3
	particule = attraction(particule, 2, 90, false, 400, 400)
	if particule.SpeedX <= 2.99999999 && particule.SpeedY <=2.99999999{
		t.Error("Il n'y a pas d'accelération",particule.SpeedX, particule.SpeedY)
	}
}

func Test_attraction4(t *testing.T){
	var particule Particle
	particule.PositionX = 200
	particule.PositionY = 200
	particule.SpeedX = 3
	particule.SpeedY = 3
	particule = attraction(particule, 2, 180, false, 400, 400)
	if particule.SpeedX <= 2.99999999 && particule.SpeedY <=2.99999999{
		t.Error("Il n'y a pas d'accelération",particule.SpeedX, particule.SpeedY)
	}
}

func Test_attraction5(t *testing.T){
	var particule Particle
	particule.PositionX = 200
	particule.PositionY = 200
	particule.SpeedX = 3
	particule.SpeedY = 3
	particule = attraction(particule, 2, 270, false, 400, 400)
	if particule.SpeedX <= 2.99999999 && particule.SpeedY <=2.99999999{
		t.Error("Il n'y a pas d'accelération",particule.SpeedX, particule.SpeedY)
	}
}


func Test_gravitation_positive(t *testing.T){
	config.General.Gravitation = 1
	config.General.InitNumParticles = 100
	var systeme System = NewSystem()
	for i := 0; i < len(systeme.Content); i++ {
		systeme.Content[i].SpeedX = 0
		systeme.Content[i].SpeedY = 0 
	}
	var centreX1 float64
	var centreY1 float64
	for i := 0; i < len(systeme.Content); i++ {
		centreX1 += systeme.Content[i].PositionX
		centreY1 += systeme.Content[i].PositionY
	}
	centreX1 /= float64(len(systeme.Content))
	centreY1 /= float64(len(systeme.Content))
	var distance_moyenne1 float64
	for i := 0; i < len(systeme.Content); i++ {
		distance_moyenne1 += math.Sqrt((systeme.Content[i].PositionX - centreX1)*(systeme.Content[i].PositionX - centreX1) +(systeme.Content[i].PositionY - centreY1)*(systeme.Content[i].PositionY - centreY1))
	}
	distance_moyenne1 /= float64(len(systeme.Content))
	systeme.Content = gravitation(systeme.Content, 0)
	systeme.Content = deplacement(systeme.Content, 0)

	for i := 0; i < len(systeme.Content); i++ {
		systeme.Content[i].SpeedX = 0
		systeme.Content[i].SpeedY = 0 
	}
	var centreX2 float64
	var centreY2 float64
	for i := 0; i < len(systeme.Content); i++ {
		centreX2 += systeme.Content[i].PositionX
		centreY2 += systeme.Content[i].PositionY
	}
	centreX2 /= float64(len(systeme.Content))
	centreY2 /= float64(len(systeme.Content))
	var distance_moyenne2 float64
	for i := 0; i < len(systeme.Content); i++ {
		distance_moyenne2 += math.Sqrt((systeme.Content[i].PositionX - centreX2)*(systeme.Content[i].PositionX - centreX2) +(systeme.Content[i].PositionY - centreY2)*(systeme.Content[i].PositionY - centreY2))
	}
	distance_moyenne2 /= float64(len(systeme.Content))
	if distance_moyenne1 <= distance_moyenne2 {
		t.Error("Il n'y a pas de gravitation")
	}
}

func Test_gravitation_nulle(t *testing.T){
	config.General.Gravitation = 0
	config.General.InitNumParticles = 100
	var systeme System = NewSystem()
	for i := 0; i < len(systeme.Content); i++ {
		systeme.Content[i].SpeedX = 0
		systeme.Content[i].SpeedY = 0 
	}
	var centreX1 float64
	var centreY1 float64
	for i := 0; i < len(systeme.Content); i++ {
		centreX1 += systeme.Content[i].PositionX
		centreY1 += systeme.Content[i].PositionY
	}
	centreX1 /= float64(len(systeme.Content))
	centreY1 /= float64(len(systeme.Content))
	var distance_moyenne1 float64
	for i := 0; i < len(systeme.Content); i++ {
		distance_moyenne1 += math.Sqrt((systeme.Content[i].PositionX - centreX1)*(systeme.Content[i].PositionX - centreX1) +(systeme.Content[i].PositionY - centreY1)*(systeme.Content[i].PositionY - centreY1))
	}
	distance_moyenne1 /= float64(len(systeme.Content))
	systeme.Content = gravitation(systeme.Content, 0)
	systeme.Content = deplacement(systeme.Content, 0)

	for i := 0; i < len(systeme.Content); i++ {
		systeme.Content[i].SpeedX = 0
		systeme.Content[i].SpeedY = 0 
	}
	var centreX2 float64
	var centreY2 float64
	for i := 0; i < len(systeme.Content); i++ {
		centreX2 += systeme.Content[i].PositionX
		centreY2 += systeme.Content[i].PositionY
	}
	centreX2 /= float64(len(systeme.Content))
	centreY2 /= float64(len(systeme.Content))
	var distance_moyenne2 float64
	for i := 0; i < len(systeme.Content); i++ {
		distance_moyenne2 += math.Sqrt((systeme.Content[i].PositionX - centreX2)*(systeme.Content[i].PositionX - centreX2) +(systeme.Content[i].PositionY - centreY2)*(systeme.Content[i].PositionY - centreY2))
	}
	distance_moyenne2 /= float64(len(systeme.Content))
	if distance_moyenne1 <= distance_moyenne2 {
		t.Error("Il n'y a pas de gravitation")
	}
}

func Test_gravitation_négative(t *testing.T){
	config.General.Gravitation = -1
	config.General.InitNumParticles = 100
	var systeme System = NewSystem()
	for i := 0; i < len(systeme.Content); i++ {
		systeme.Content[i].SpeedX = 0
		systeme.Content[i].SpeedY = 0 
	}
	var centreX1 float64
	var centreY1 float64
	for i := 0; i < len(systeme.Content); i++ {
		centreX1 += systeme.Content[i].PositionX
		centreY1 += systeme.Content[i].PositionY
	}
	centreX1 /= float64(len(systeme.Content))
	centreY1 /= float64(len(systeme.Content))
	var distance_moyenne1 float64
	for i := 0; i < len(systeme.Content); i++ {
		distance_moyenne1 += math.Sqrt((systeme.Content[i].PositionX - centreX1)*(systeme.Content[i].PositionX - centreX1) +(systeme.Content[i].PositionY - centreY1)*(systeme.Content[i].PositionY - centreY1))
	}
	distance_moyenne1 /= float64(len(systeme.Content))
	systeme.Content = gravitation(systeme.Content, 0)
	systeme.Content = deplacement(systeme.Content, 0)

	for i := 0; i < len(systeme.Content); i++ {
		systeme.Content[i].SpeedX = 0
		systeme.Content[i].SpeedY = 0 
	}
	var centreX2 float64
	var centreY2 float64
	for i := 0; i < len(systeme.Content); i++ {
		centreX2 += systeme.Content[i].PositionX
		centreY2 += systeme.Content[i].PositionY
	}
	centreX2 /= float64(len(systeme.Content))
	centreY2 /= float64(len(systeme.Content))
	var distance_moyenne2 float64
	for i := 0; i < len(systeme.Content); i++ {
		distance_moyenne2 += math.Sqrt((systeme.Content[i].PositionX - centreX2)*(systeme.Content[i].PositionX - centreX2) +(systeme.Content[i].PositionY - centreY2)*(systeme.Content[i].PositionY - centreY2))
	}
	distance_moyenne2 /= float64(len(systeme.Content))
	if distance_moyenne1 <= distance_moyenne2 {
		t.Error("Il n'y a pas de gravitation")
	}
}

func Test_fusion1(t *testing.T){
	config.General.Merge = true
	config.General.InitNumParticles = 100
	config.General.Rebond_bords = true
	var particule1 Particle
	var particule2 Particle
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	var libres int
	var systeme System
	systeme.Content = append(systeme.Content, particule1)
	systeme.Content = append(systeme.Content, particule2)
	systeme.Content, libres = fusion(systeme.Content, &systeme, libres)
	if libres != 1{
		t.Error("Il n'y a pas eu de fusion")
	}
}

func Test_fusion2(t *testing.T){
	config.General.Merge = true
	config.General.InitNumParticles = 100
	config.General.Rebond_bords = true
	var particule1 Particle
	var particule2 Particle
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 3
	particule2.ScaleY = 3
	var libres int
	var systeme System
	systeme.Content = append(systeme.Content, particule1)
	systeme.Content = append(systeme.Content, particule2)
	systeme.Content, libres = fusion(systeme.Content, &systeme, libres)
	if masse(systeme.Content[0]) < 13-10e-12 || masse(systeme.Content[0]) > 13+10e+12{
		t.Error("Il n'y a pas eu de fusion")
	}
}

func Test_fusion3(t *testing.T){
	config.General.Merge = true
	config.General.InitNumParticles = 100
	config.General.Rebond_bords = true
	var particule1 Particle
	var particule2 Particle
	particule1.ScaleX = 1
	particule1.ScaleY = 1
	particule2.ScaleX = 1
	particule2.ScaleY = 1
	var libres int
	var systeme System
	systeme.Content = append(systeme.Content, particule1)
	systeme.Content = append(systeme.Content, particule2)
	systeme.Content, libres = fusion(systeme.Content, &systeme, libres)
	if masse(systeme.Content[0]) < 2-10e12 || masse(systeme.Content[0]) > 2+10e12{
		t.Error("Il n'y a pas eu de fusion")
	}
}

func Test_masse(t *testing.T){
	var particule Particle
	particule.ScaleX = 2
	particule.ScaleY = 3
	if masse(particule) != 6{
		t.Error("La particule n'a pas la bonne masse")
	}
}

func Test_growth(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 3
	particule2.ScaleY = 3
	var masse_particule = masse(particule2)
	var particule3 Particle = growth(particule1,masse_particule)
	if masse(particule3) < 13-10e12 || masse(particule3) > 13+10e12{
		t.Error("Le grossissement n'est pas le bon")
	}
}

func Test_distance(t *testing.T){
	var particule1 Particle
	var particule2 Particle
	particule1.PositionX = 500
	particule1.PositionY = 100
	particule2.PositionX = 300
	particule2.PositionY = 700
	particule1.ScaleX = 2
	particule1.ScaleY = 2
	particule2.ScaleX = 2
	particule2.ScaleY = 2
	if distance(particule1,particule2) < 200-10e12 || distance(particule1,particule2) > 200+10e12{
		t.Error("Il n'y a pas la bonne distance entre les deux particules")
	}
}