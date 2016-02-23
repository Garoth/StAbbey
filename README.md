# StAbbey Overview

StAbbey is a work in progress game project. It aims to be something
like a digital boardgame. Right now, it's in very rough shape, but
has the following:

- Webserver with websockets to connect players and spectators
- Multiplayer interaction
- Basic monster pathing AI (A* implementation)
- Some work in process room generators
- Some player skills
- Basic web client (somewhat broken,
        but look at the original-client-prototype tag)
- Some game flow, like switching levels by using the stairs

Maybe there is something of interest for you.

## How To Run The StAbbey Server

0) ./stabbey.sh deps # Installs necessary dependencies locally
1) ./stabbey.sh run # Builds & runs system locally
2) Point your browser at localhost:8080
3) Be happy that you can now get to work and not fuck around with the build
   system for 30 minutes.

(View advanced options by typing ./stabbey.sh)
