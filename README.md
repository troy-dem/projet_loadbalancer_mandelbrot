# projet_loadbalancer_mandelbrot
##Création d'une page web de génération de figure de mandelbrot à l'aide de loadbalancing(nginx)
##Ce projet se compoe de différentes parties:
### Une partie Backend : 
Pour notre backend, nous avaons décidé d'utilisé docker afin de représenter différents serveurs qui se partagerait les différents travaux. Nous retrouvons donc dans cette partie : 
- Les workers, qui sont les serveurs contenant la logique des mandelbrots
- Les configurations docker qui nous permettent de gérer l'infrastructure
### Une partie frontend : 
Celle-ci a été réalisée en python à l'aide de Flask. C'est dans cette partie que le client va pouvoir choisir la configuration de son mandelbrot et la générer. Afin de générer le résultat final le serveur de la partie frontend va séparer la requête initiales en une multitude de sous requêtes qu'il va envoyer de manière asynchrone à un serveur. Une fois qu'il aura obtenu toutes les réponses, le serveur réassemblera l'image. 
### La partie loadbalancing : 
C'est cette partie qui va recevoir la multitude de requêtes venant du frontend et va les loadbalancer au travers des différents serveurs disponibles. Nous avons chois de travailler avec un loadbalancer serveur Nginx. Nous avons ajouter sur ce github la configuration que l'on avait faite pour celui-ci afin. 
