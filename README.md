# Projet_loadbalancer_mandelbrot
Création d'une page web de génération de figure de mandelbrot à l'aide de loadbalancing(nginx)
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
Chacun des workers est un serveur go contenu sur un docker et contient l'API. Nous les avons générer via un docker-compose utilisant l'image de go. <br />
![image](https://user-images.githubusercontent.com/75576766/210776881-2a7781f4-8e6a-4a8c-a552-ba32d102ce05.png)

L'API ne contient qu'un seul appel : <br />

![image](https://user-images.githubusercontent.com/75576766/210778790-5a722d7f-6bde-436f-ac76-ee8a47ddd98c.png)

Pour traiter la requête le serveur passe va utiliser plusieurs fonctions :
- png_generator(resolution_x, resolution_y, start_position_x, start_position_y, quantize_length, max_iteration float64, colormap [][3]int, wg sync.WaitGroup) : qui va nous permettre de créer l'image du mandelbrot sous format png.
- mandelbrot(max_iteration, c_real, c_imaginary float64): ( utilisée par png_generator ) Celle ci va nous donner le nombre d'itération à partir du quel on sait que l'élément ne fait pas partie de l'ensemble de mandelbrot. 
- colorize(iteration, max_iteration float64, colormap [][3]int) : ( utilisée par png generator ) Celle-ci va nous renvoyer la couleur que l'on doit donner à un pixel selon son nombre d'iteration. 

### Partie Frontend :


## Stratégie de répartion du loadbalancer :

Nous avons fait le choix d'utiliser nginx pour loadbalancer les requêtes. Nginx va recevoir toutes les requêtes et les répartir aux travers des différents serveurs qui lui sont attribués pour ce faire il suffit de configurer nginx comme montrer ci dessous : <br />
![image](https://user-images.githubusercontent.com/75576766/210776420-f684d79a-4f7d-4559-95bd-2eda77197d88.png)

Comme nous le voyons dans l'image, lorsque celui-ci va recevoir un appel, il va le redistribuer vers un des 4 serveurs disponibles ( créer avec docker ) 

Nous utilisons une stratégie de loadbalancing least connections, c'est à dire que le loadbalancer va rediriger la requêtes vers le serveur ayant le moins de connexion active. Il est très simple à mettre en place sur nginx, il suffit en effet de rajouter une seule ligne dans le fichier de configuration afin de terminer la méthode de loadbalancing.
<br />
![Sans titre](https://user-images.githubusercontent.com/75576766/210772008-9197432c-1f45-4ab5-bc14-8d36af537035.png)

D'autres méthodes sont disponibles : 
- Round Robin : dans laquelle les requêtes sont distribués de manière équivalentes entre les serveurs
- IP hash : on détermine le serveur d'arriver de la requête via l'addresse IP du client.

La méthode least connections nous semblait la plus approprié dans notre situation car les serveurs pourraient ne pas être exclusivement réserver aux traitements mandelbrot. 

Pour plus d'information sur les différentes configurations : 
https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/


