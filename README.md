# Dungeon Server

This is a golang server for a Dungeons and Dragons character/DM dashboard. All DnD content is pulled from the [open5e](https://api-beta.open5e.com/) API and character progression is saved in a Cloud Firestore.

Requires [Golang](https://golang.org/doc/install) 1.12 or higher.

## Firestore Setup

You must have the correct credentials set up to connect to a Firestore datastore. You can read about how to set up a new one [here](https://firebase.google.com/docs/firestore/quickstart). If you are trying to connect to an existing datastore then you only need the default credentials json file and the firestore config in your environment variables, e.g.

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/service_account.json"
export FIREBASE_CONFIG='/path/to/config.json' # This could also be the raw config json string 
```

## Getting Started

> `go run .`