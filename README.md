# Glowed
A Dockerized Golang/React web app that will display your Rocket League stats in a 'trophy room' way.

Data provied by [RocketLeagueStats](https://rocketleaguestats.com/).

## Local Development
### Configure
1. Obtain your own [API Key](https://developers.rocketleaguestats.com/).
2. Set the following in your `.env`:
```
# Steam 64 ID / PSN Username / Xbox GamerTag or XUID
PROFILE=GamerTag
# 1 - PC, 2- PS, 3 - Xbox
PLATFORM=3
# Rocket League Stats
RLS_API_KEY=
```
### Run
1. `$ docker-compose -d --build`.
