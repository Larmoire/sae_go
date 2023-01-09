# Ebiten
Ebitengine (Eh-Bee-Ten-Gin) (anciennement connu sous le nom d'Ebiten) est un moteur de jeu open source pour le langage de programmation Go. L'API simple d'Ebitengine vous permet de développer rapidement et facilement des jeux 2D pouvant être déployés sur plusieurs plates-formes.
# Utilisation
Tout d'abord, on importe la libraire.
```go
import (
	"github.com/hajimehoshi/ebiten/v2"
)
```
Dans notre projet, cette libraire nous sert à gérer un système de particules, parametré par un fichier config.json.

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
	"SpawnY": 200,
	"SpawnRate": 0,
	"RandomSpeed": false,
	"SpeedX": 5,
	"SpeedY": 5,	
	"SpeedXmin": -10,
	"SpeedXmax": 10,
	"SpeedYmin": -10,
	"SpeedYmax": 10,
	"Lifespan" : -1,
	"Opacity" : 1,
	"Gravity": false,
	"GravityVal" : 0.1,
	"Optimisation" : false,
	"SpawnAtMouse": false,
	"SpawnPerClick": 1,
	"Fade" : false,
	"RVBchange" : false,
	"Bounce" : false,
	"ColorBounce" : false
}
```
# Vue détaillée
## Paramètrage de la fenêtre et setup
**"WindowTitle": "Project particles"** donne le nom de la fenêtre.  

**"WindowSizeX": 800** et **"WindowSizeY": 600** servent à dimensionner la fenêtre de base (donc pour dimension 800x600). 
 
**"ParticleImage": "assets/particle.png"** définit l'image pour la particule.  

**"Debug": bool** sert à afficher le debug sur la fenêtre, affichant les fps et le nombre de particules créées.
   
**"InitNumParticles": int** définit le nombre de particules au lancement du code.

## Paramètrage de la particule
### Spawn

**"RandomSpawn": bool** si est true, les particules sont générées aléatoirement dans la fenêtre,  
Autrement **"SpawnX": float64** et **"SpawnY": float64** sont les coordonnées auxquelles sont générées les particules.  

**"ScaleX": float64** et **"ScaleY": float64** pour les dimensions de la particule, soit 1 fois les dimensions de base.  

**"ColorRed": float64**, **"ColorGreen": float64** et **"ColorBlue": float64** définissent la couleur de la particule, 0 correspond au 0 et le 1 au 255.  

### Updates 
**"SpawnRate": float64** définit le nombres de particules générées à chaque update, exemple 0.5 correspond à deux updates pour 1 particule.  

**"RandomSpeed": bool** si est false, les particules auront une vitesse de **"SpeedX": float64** et de **"SpeedY": float64**.  
Dans le cas contraire, elles auront une vitesse aléatoire entre **"SpeedXmin": float64** et **"SpeedXmax": float64** pour **X**,  
Et entre **"SpeedYmin": float64** et **"SpeedYmax": float64** pour **Y**.  

**"Lifespan": int** correspond à la durée de vie de la particule, soit le nombre d'update avant sa disparition. La valeur de -1 désactive cette option.  
**"Opacity" : float64** gère l'opacité de la particule, 1 correspond à 100% (visible). Elle se couple généralement au Lifespan.  

**"Gravity": bool** gère l'activation de la gravité, si elle est true alors la valeur de la gravité est attribuée par **"GravityVal" : float64**.  

**"Optimisation" : bool** est le mode d'optimisation proposé par l'énoncé, il récupère les particules mortes pour faire de nouvelles particules plutôt que d'en générer de nouvelles.  

**"SpawnAtMouse": bool** est une manière de générer des particules, à l'endroit de la souris par un clic gauche. Le nombre de particules générées est géré par l'option **"SpawnPerClick": int**. En complément, on peut activer ou non l'option **"Fade" : bool** qui diminue l'opacité des particules générées à mesure que le clic gauche reste enfoncé. Une fois relaché, l'opacité reviens à sa valeur de base.  

**"RVBchange" : bool** est une option pour faire varier les couleurs rouge, vert et bleu avec les touches 'R','V','B'. 
 
**"Bounce" : bool** est une option, si activée, qui permet aux particules de rebondir contre les bords de la fenêtre.  
À cette option se couple la suivante; **"ColorBounce" : bool** qui une fois active définit une couleur aléatoire pour la particule quand elle touche un bord de la fenêtre.
