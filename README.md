 
# fldarkdata


## Description

Reusable library parser of Freelancer Discovery game data.
- As interesting features, it offers custom ORM like static typed access to data.
- With ability to read and write data back without requiring to write code for writing
    - because it reads and writes game data configs in a universal way

This library is used at least in projects:
- fldarklint (config formatter)
- fldarkstat (online flstat)


## How to use

TODO to add.


## Architecture

```mermaid
flowchart TD
    parser[parser\nProvides static typed access to parsedFreelancer configs]
    parser --> freelancer[freelancer\nFreelancer Data Parsers\npackage reflects\nFreelancer File Structure]
    parser --> filefind[filefind\nfinds freelancer files]
    freelancer --> inireader[inireader\nUniversal freelancer ini format reader\nLoads to changable structure\nThat can be rendered back]
    freelancer --> semantic[semantic\nORM mapper of inireader values for quick typed acess to values\nfor reading and writing, without need to map all file structure]
    semantic --> inireader
```
