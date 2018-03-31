# Glowed
A Dockerized Golang/React web app that will display your Rocket League stats in a 'trophy room' way.

Data provied by [RocketLeagueStats](https://rocketleaguestats.com/).

# Usage
1. Obtain your own API Key
2. Set the following:
```
# Steam 64 ID / PSN Username / Xbox GamerTag or XUID
PROFILE=
# PC, PS, Xbox
PLATFORM=
# Rocket League Stats
RLS_API_KEY=
```
3. Run `docker-compose -d --build`
