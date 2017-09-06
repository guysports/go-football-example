# go-football-example
Sample program for consuming data from football-data.org in Go.

Using the REST-API provided by http://football-data.org, it gets the fixtures fora weekend and prints the results.

Usage
```
glide install
go install
go-football-example
```
Optional:
```
export FOOTBALL_API_KEY=<your_football_data_key>
```

Output
```
Crystal Palace FC 2 - 0 West Bromwich Albion FC
Sunderland AFC 2 - 2 West Ham United FC
Norwich City FC 1 - 2 Leicester City FC
Aston Villa FC 0 - 1 Stoke City FC
Manchester City FC 6 - 1 Newcastle United FC
AFC Bournemouth 1 - 1 Watford FC
Chelsea FC 1 - 3 Southampton FC
Everton FC 1 - 1 Liverpool FC
Swansea City FC 2 - 2 Tottenham Hotspur FC
Arsenal FC 3 - 0 Manchester United FC
```
