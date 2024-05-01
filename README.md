# Lockbin server
Part of the lockbin project

!WIP!

HereWillBeALinkToTheMainRepo

## Description
Lockbin aims to be a pastebin-like service but the data is available only after a given time and will be deleted after another given time.
The server don't know the data due the client side encryption.

## What does the server know?
- UUID (Unique identifier for every record)
- MasterKey  (The creator can delete a record with it)
- UnlockTime (The time after the record is available)
- DeleteTime (The time after the record is unavailable and deleted)
- Encrypted Message (The encrypted data what can be unlocked with the public password)

## Roadmap
- [ ] Server able to handle requests and store data properly
- [ ] Creating private servers - Add optional password protection, not letting anyone use the server
