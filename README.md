# projet_loadbalancer_mandelbrot
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

## Stratégie de répartion du loadbalancer :

Nous utilisons un loadbalancing least connections, c'est à dire que le loadbalancer va rediriger la requêtes vers le serveur ayant le moins de connexion active. Il est très simple à mettre en place sur nginx, il suffit en effet de rajouter une seule ligne dans le fichier de configuration afin de terminer la méthode de loadbalancing. ![Sans titre](https://user-images.githubusercontent.com/75576766/210772008-9197432c-1f45-4ab5-bc14-8d36af537035.png)

D'autres méthodes sont disponibles : 
- Round Robin : dans laquelle les requêtes sont distribués de manière équivalentes entre les serveurs
- IP hash : on détermine le serveur d'arriver de la requête via l'addresse IP du client.

La méthode least connections nous semblait la plus approprié dans notre situation car les serveurs pourraient ne pas être exclusivement réserver aux traitements mandelbrot. 

Pour plus d'information sur les différentes configurations : 
https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/

