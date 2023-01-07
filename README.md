# Projet_loadbalancer_mandelbrot
Création d'une page web de génération de figure de mandelbrot à l'aide de loadbalancing(nginx) par Louis Demarcin 18090 et Logan Noel 18003
## Ce projet se compose de différentes parties:
### Une partie Backend : 
Pour notre backend, nous avaons décidé d'utilisé docker afin de représenter différents serveurs qui se partagerait les différents travaux. Nous retrouvons donc dans cette partie : 
- Les workers, qui sont les serveurs contenant la logique des mandelbrots
- Les configurations docker qui nous permettent de gérer l'infrastructure
### Une partie frontend : 
Celle-ci a été réalisée en python à l'aide de Flask. C'est dans cette partie que le client va pouvoir choisir la configuration de son mandelbrot et la générer. Afin de générer le résultat final le serveur de la partie frontend va séparer la requête initiales en une multitude de sous requêtes qu'il va envoyer de manière asynchrone à un serveur. Une fois qu'il aura obtenu toutes les réponses, le serveur réassemblera l'image. 
### La partie loadbalancing : 
C'est cette partie qui va recevoir la multitude de requêtes venant du frontend et va les loadbalancer au travers des différents serveurs disponibles. Nous avons chois de travailler avec un loadbalancer serveur :  Nginx. Nous avons ajouter sur ce github la configuration que l'on a faite pour lui. 
 

## Description de l'API du serveur :

Notre API est assez simpliste vu que la seule requêtes qu'on lui envoit est une demande de création d'image. 

Néanmoins afin de bien comprendre le processus voici un schéma le décrivant : 
<br />
![description](https://user-images.githubusercontent.com/75576766/210775222-1458784d-ffd4-4177-bbab-403f847743a2.png)

### Partie Backend:
Chacun des workers est un serveur go contenu sur un docker et contient l'API. Nous les avons générer via un docker-compose utilisant l'image de go.<br />
![image](https://user-images.githubusercontent.com/75576766/210776881-2a7781f4-8e6a-4a8c-a552-ba32d102ce05.png)

Chacun des workers n'utilise qu'un seul coeur de l'ordinateur, on leur a attribué les coeurs dans le docker-compose avec le cpuset:<br/>
![image](https://user-images.githubusercontent.com/75576766/210787225-2c45037d-b78b-4a61-91d1-13d5575038c2.png)

L'API ne contient qu'un seul appel : <br />

![description](https://user-images.githubusercontent.com/75576766/211143535-72d0dff4-8478-4f80-869b-5ea9549384dd.png)

Pour traiter la requête le serveur passe va utiliser plusieurs fonctions :
- png_generator(resolution_x, resolution_y, start_position_x, start_position_y, quantize_length, max_iteration float64, colormap [][3]int, wg sync.WaitGroup) : qui va nous permettre de créer l'image du mandelbrot sous format png.
- mandelbrot(max_iteration, c_real, c_imaginary float64): ( utilisée par png_generator ) Celle ci va nous donner le nombre d'itération à partir du quel on sait que l'élément ne fait pas partie de l'ensemble de mandelbrot. 
- colorize(iteration, max_iteration float64, colormap [][3]int) : ( utilisée par png generator ) Celle-ci va nous renvoyer la couleur que l'on doit donner à un pixel selon son nombre d'iteration. 

Le code ayant servi a générer les docker ainsi que le code serveur est retoruvable ici : https://github.com/Tateuf/projet_loadbalancer_mandelbrot/tree/main/backend%20docker-compose%20%2B%20worker

### Partie Frontend :
Nous avons décidé d'utilisé Flask pour créer notre serveur. 
Voici notre interface de base qui est en html, nous avaons rajouter un peu de js afin de rendre celle-ci plus esthétique : <br />
![image](https://user-images.githubusercontent.com/75576766/210783174-b3f90217-7240-49f1-b0bf-b6597562e11e.png)

Lorsque nous appuyons sur le boutton, cela va lancer un script js qui fait une une requête http à notre serveur "backend" du frontend. Celui-ci va subdiviser la requête en une centaine de sous requêtes qu'il va transmettre au loadbalancer de manière asyncrhone. Lorsqu'il aura récupérer chacune des réponses, il va reconstruire l'image et nous pourrons l'afficher dans l'interface utilisateur. 

Lien vers le code frontend : 
https://github.com/Tateuf/projet_loadbalancer_mandelbrot/tree/main/frontend

## Stratégie de répartion du loadbalancer :

Nous avons fait le choix d'utiliser nginx pour loadbalancer les requêtes. Nginx va recevoir toutes les requêtes et les répartir aux travers des différents serveurs qui lui sont attribués pour ce faire il suffit de configurer nginx comme montrer ci dessous : <br />
![image](https://user-images.githubusercontent.com/75576766/210776420-f684d79a-4f7d-4559-95bd-2eda77197d88.png)

Comme nous le voyons dans l'image, lorsque celui-ci va recevoir un appel, il va le redistribuer vers un des 4 serveurs disponibles ( créer avec docker ) 

Nous utilisons une stratégie de loadbalancing least connections, c'est à dire que le loadbalancer va rediriger la requêtes vers le serveur ayant le moins de connexion active. Il est très simple à mettre en place sur nginx, il suffit en effet de rajouter une seule ligne dans le fichier de configuration afin de terminer la méthode de loadbalancing.
<br />
![Sans titre](https://user-images.githubusercontent.com/75576766/210772008-9197432c-1f45-4ab5-bc14-8d36af537035.png)

D'autres méthodes sont disponibles : 
- Round Robin : dans laquelle les requêtes sont distribués de manière équivalentes entre les serveurs
- IP hash : on détermine le serveur traitant la requête à partir de l'addresse IP du client.

La méthode least connections nous semblait la plus approprié dans notre situation car les serveurs pourraient ne pas être exclusivement réserver aux traitements mandelbrot. 

Pour retrouver la configuration que l'on a faite :
https://github.com/Tateuf/projet_loadbalancer_mandelbrot/tree/main/conf

Pour plus d'information sur les différentes configurations : 
https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/

## Bibliothèques utilisées et description :

### Partie frontend :
- flask qui est un framework opensource de dévellopement web en python 
- aiohttp et asyncio qui nous permettent d'effectuer des requêtes aynschrones en python 
- pillow qui nous permet de faire du traitement d'image 
### Partie backend : 
- net/http qui nous permet de créer un serveur http en go 
- image qui nous permet de faire du traitement d'image en go
- math pour effectuer des opérations mathématique en go
