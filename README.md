# Ebiten
Ebitengine (Eh-Bee-Ten-Gin) (anciennement connu sous le nom d'Ebiten) est un moteur de jeu open source pour le langage de programmation Go. L'API simple d'Ebitengine vous permet de développer rapidement et facilement des jeux 2D pouvant être déployés sur plusieurs plates-formes.
# Utilisation


Dans notre projet, cette libraire Ebiten nous sert à gérer un système de particules, parametré par des fichier .json.

```json
config.json
{
	"WindowTitle": "Project particles",
	"WindowSizeX": 800,
	"WindowSizeY": 600,
	"ParticleImage": "assets/particle.png",
	"Debug": true,
	"InitNumParticles": 1,
	"RandomSpawn": false,
	"ScaleX": 1,
	"ScaleY": 1,
	"ColorRed": 1,
	"ColorGreen": 1,
	"ColorBlue": 1,
	"SpawnX": 400,
	"SpawnY": 300,
	"SpawnRate": 0.5,
	"RandomSpeed": false,
	"SpeedFix" : false,
	"SpeedX": 5,
	"SpeedY": 5,	
	"SpeedXmin": -10,
	"SpeedXmax": 10,
	"SpeedYmin": -10,
	"SpeedYmax": 10,
	"Lifespan" : -1,
	"Opacity" : 1,
    "Gravity": false,
    "GravityVal" : 0.5,
	"Optimisation" : false,
	"Fade" : false,
    "SpawnAtMouse": true,
    "SpawnPerClick": 1,
    "RGBchange" : false,
	"Bounce" : false,
	"ColorBounce" : false,
	"Rotate" : false
}
```
# Vue détaillée

## Paramètrage de la fenêtre et setup

Cela donne le nom de la fenêtre.  
```json
"WindowTitle": "Project particles"
``` 
Cela sert à dimensionner la fenêtre de base (donc pour dimension 800x600).
```json
"WindowSizeX": 800
"WindowSizeY": 600
``` 
Cela sert à définir l'image pour la particule.
```json
"ParticleImage": "assets/particle.png"
``` 
Cela sert à afficher le debug sur la fenêtre, affichant les fps et le nombre de particules dans la liste.
```json
"Debug": true
``` 
Cela définit le nombre de particules au lancement du code.
```json   
"InitNumParticles": 1
``` 


## Paramètrage de la particule


### Spawn
Cela permet de générer les particules aléatoirement dans la fenêtre,  
```json
"RandomSpawn": true 
``` 
Autrement, ce sont ces coordonnées auxquelles sont générées les particules.  
```json
"SpawnX": 400
"SpawnY": 300
``` 
Cela permet d'ajuster les dimensions de la particule, soit 1 fois les dimensions de base ici.  
```json
"ScaleX": 1
"ScaleY": 1 
``` 
Cela permet de définir la couleur de la particule, 0 correspond au 0 et le 1 au 255. 
```json
"ColorRed": 1
"ColorGreen": 1 
"ColorBlue": 1
``` 

### Updates 
Cela définit le nombres de particules générées à chaque update, exemple 0.5 correspond à deux updates pour 1 particule. 
```json
"SpawnRate": 0.5 
``` 
Cela permet de définir une vitesse aléatoire 
```json
"RandomSpeed": true 
``` 
Si elle est à true les particules auront une vitesse aléatoire entre ces variables en X et en Y.
```json
"SpeedXmin": -3
"SpeedXmax": 3
"SpeedYmin": -3
"SpeedYmax": 3
```
Dans le cas ou RandomSpeed est false, les particules prenent une vitesse fixe en X et en Y.
```json
"SpeedX": 5
"SpeedY": 5 
``` 
Cela permet de figer l'ecran, ne plus faire bouger les particules.
```json
"SpeedFix": true
``` 
Cela correspond à la durée de vie de la particule, soit le nombre d'update avant sa disparition. La valeur de -1 désactive cette option.  
```json
"Lifespan": 100
``` 
Cela correspond à l'opacité de base la particule, 1 correspond à 100% (visible). Elle se couple généralement au Lifespan. 
```json
"Opacity" : 1
```  
Cela gère l'activation de la gravité.
```json
"Gravity": true
```
Si elle est true alors la valeur de la gravité est attribuée grâce à cela.
```json
"GravityVal" : 0.5
```  
Ceci est le mode d'optimisation, il permet de supprimer de la liste les particules qui ne doivent plus être affichées.
```json
"Optimisation" : true
```
Ceci est une manière de générer des particules, à l'endroit de la souris par un clic gauche.
```json
"SpawnAtMouse": true 
```
Ceci est le nombre de particules générées si l'option précédente est true.
```json
"SpawnPerClick": 1 
```
En complément, on peut activer ou non cette option qui diminue l'opacité des particules générées à mesure que le clic gauche reste enfoncé. Une fois relaché, l'opacité reviens à sa valeur de base.  
```json
"Fade" : true
```
C'est une option pour faire varier les couleurs avec les touches affichées à l'écran si actif. (couleurs disponibles : Rouge, Vert, Bleu, Cyan, Jaune, Violet, Magenta, Lime, Blanc, Orange).
```json
"RGBchange" : true
```
C'est une option, si activée, permettant aux particules de rebondir contre les bords de la fenêtre. 
```json 
"Bounce" : true
```
À cette option se couple la suivante; qui définit une couleur aléatoire pour la particule quand elle touche un bord de la fenêtre.
```json
"ColorBounce" : true
``` 
C'est une manière d'organiser la trajectoire des particules. Lorsqu'il est true, les particules générées sont au nombre de 3 pour 1 et entrent en rotation avec pour centre le milieu de l'ecran.
```json
"Rotate" : true
``` 
C'est un mode qui permet de déplacer le point de spawn si RandomSpawn est false avec les flèches directionnelles
```json
"Arrows" : true


# Pratique
Un commentaire en début de page mentionne les fonctions présente dans celle-ci.
Un interface sur la fenêtre des particules permet de modifier directement les paramètres de "RandomSpeed", "RandomSpawn", "SpawnAtMouse",  "RGBchange", "Gravity", "Bounce" (et "ColorBounce"), "Rotate" et "SpeedFix".
