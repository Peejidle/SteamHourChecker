# SteamHourChecker

A small command-line tool written in Go that looks up how many hours you've played a specific game on Steam.

Give it your Steam API key, your SteamID, and a game name, and it prints your total playtime and your playtime over the last two weeks.

## Requirements

- [Go](https://go.dev/dl/) installed
- A Steam Web API key
- Your SteamID64
- A Steam profile with **game details set to public** (private profiles return no data)

### Getting your API key

1. Log into Steam and go to https://steamcommunity.com/dev/apikey
2. Register a key (if it asks for a domain name, `localhost` works)

Keep this key private. Never commit it or share it.

### Finding your SteamID

The API needs your **SteamID64**, the long all-digits ID, not your vanity profile name. If your profile URL looks like `steamcommunity.com/profiles/76561198000000000`, the number at the end is your SteamID64. If you have a custom URL, a site like [steamid.io](https://steamid.io) can convert it.

## Usage

Clone the repo and run it:

```
git clone git@github.com:Peejidle/SteamHourChecker.git
cd SteamHourChecker
go run main.go
```

The program will prompt you for three things:

1. Your Steam API key
2. Your SteamID64 (just the numbers)
3. The name of the game you want to check

Example output:

```
Name: Some Game
Total Playtime: 142
Last 2 Weeks: 6
```

Hours are rounded down to the nearest whole hour.

## Notes

- The game name must match the name in your Steam library exactly, including capitalization and punctuation. If nothing matches, the program tells you to check for typos.
- Your API key and SteamID are only used for the request to Steam's API. Nothing is saved to disk.
- Data comes from Steam's `IPlayerService/GetOwnedGames` endpoint.

## Project status

A small side project. It does what it says and isn't under active development.
