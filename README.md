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

```json
"WindowTitle": "Project particles"
``` 
Cela donne le nom de la fenêtre.  

```json
"WindowSizeX": 800,
"WindowSizeY": 600
``` 
servent à dimensionner la fenêtre de base (donc pour dimension 800x600).

```json
**"ParticleImage": "assets/particle.png"** définit l'image pour la particule.  
``` 
```json
**"Debug": bool** sert à afficher le debug sur la fenêtre, affichant les fps et le nombre de particules créées.
``` 
```json   
**"InitNumParticles": int** définit le nombre de particules au lancement du code.
``` 


## Paramètrage de la particule


### Spawn
```json
**"RandomSpawn": bool** si est true, les particules sont générées aléatoirement dans la fenêtre,  
``` 
```json
Autrement **"SpawnX": float64** et **"SpawnY": float64** sont les coordonnées auxquelles sont générées les particules.  
``` 
```json
**"ScaleX": float64** et **"ScaleY": float64** pour les dimensions de la particule, soit 1 fois les dimensions de base.  
``` 
```json
**"ColorRed": float64**, **"ColorGreen": float64** et **"ColorBlue": float64** définissent la couleur de la particule, 0 correspond au 0 et le 1 au 255. 
``` 

### Updates 
```json
**"SpawnRate": float64** définit le nombres de particules générées à chaque update, exemple 0.5 correspond à deux updates pour 1 particule. 
```  
```json
**"RandomSpeed": bool** si est false, les particules auront une vitesse de **"SpeedX": float64** et de **"SpeedY": float64**.  
Dans le cas contraire, elles auront une vitesse aléatoire entre **"SpeedXmin": float64** et **"SpeedXmax": float64** pour **X**,
Et entre **"SpeedYmin": float64** et **"SpeedYmax": float64** pour **Y**.  
``` 
```json
**"SpeedFix" : bool** est un option qui permet de figer la vitesse des particules ainsi que l'impact de la gravité sur elles.
``` 
```json
**"Lifespan": int** 
``` correspond à la durée de vie de la particule, soit le nombre d'update avant sa disparition. La valeur de -1 désactive cette option.  
```json
**"Opacity" : float64**
```  
gère l'opacité de la particule, 1 correspond à 100% (visible). Elle se couple généralement au Lifespan. 
```json
**"Gravity": bool**
```
gère l'activation de la gravité, si elle est true alors la valeur de la gravité est attribuée par **"GravityVal" : float64**.  
```json
**"Optimisation" : bool**
```
est le mode d'optimisation proposé par l'énoncé, il récupère les particules mortes pour faire de nouvelles particules plutôt que d'en générer de nouvelles.  
```json
**"SpawnAtMouse": bool** est une manière de générer des particules, à l'endroit de la souris par un clic gauche. Le nombre de particules générées est géré par l'option **"SpawnPerClick": int**. En complément, on peut activer ou non l'option **"Fade" : bool** qui diminue l'opacité des particules générées à mesure que le clic gauche reste enfoncé. Une fois relaché, l'opacité reviens à sa valeur de base.  
``` 
```json
**"RGBchange" : bool** est une option pour faire varier les couleurs rouge, vert et bleu avec les touches 'R','G','B' (couleurs disponibles : Rouge, Vert, Bleu, Cyan, Jaune, Violet).
``` 
```json 
**"Bounce" : bool** est une option, si activée, qui permet aux particules de rebondir contre les bords de la fenêtre.  
À cette option se couple la suivante; **"ColorBounce" : bool** qui une fois active définit une couleur aléatoire pour la particule quand elle touche un bord de la fenêtre.
``` 
```json
**"Rotate" : bool** est une manière d'organiser la trajectoire des particules. Lorsqu'il est true, les particules générées sont au nombre de 3 pour 1 et entrent en rotation avec pour centre le milieu de l'ecran.
``` 


# Pratique
Un commentaire en début de page mentionne les fonctions présente dans celle-ci.
Un interface sur la fenêtre des particules permet de modifier directement les paramètres de **"Gravity" : bool**, de **"Bounce" : bool** (et **"ColorBounce" : bool**), **"RandomSpeed": bool**, **"RGBchange" : bool**, **"SpawnAtMouse": bool**, **"SpeedFix" : bool**.
