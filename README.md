# nba-mobler-engine
NBA Mobler Engine


Packages used
- SQLC
- Postgres DB Connector
- Goose(For setting up database schema)


Process Flow
- Teams need to be created in the system before adding/creating player
- After successfull creation of teams, when creating the player we pass the team's id to it when creating
- Now the game stats can be addes, and updated
- The postman endpoints for the application is below


NB: The postman collection is in the code base