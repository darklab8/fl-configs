# fldarkdata

## Description

Reusable library parser of Freelancer Discovery game data.

freelancer (resembling ini) config reader/writer with easily mapping variables to access in ORM - object relational mapping fashion.
This alone allows quickly accessing any config data with least amont of code effort for additional features.

# Features

- it offers custom ORM like static typed access to data.
- With ability to read and write data back without requiring to write code for writing
  - U can be just changing ORM mapped values

This library is used at least in projects:

- fldarklint (config formatter)
- fldarkstat (online flstat)

## Architecture

```mermaid
flowchart TD
    mapped[mapped\nProvides static typed access to parsedFreelancer configs]
    mapped --> freelancer[freelancer\nFreelancer Data Parsers\npackage reflects\nFreelancer File Structure]
    mapped --> filefind[filefind\nfinds freelancer files]
    freelancer --> inireader[inireader\nUniversal freelancer ini format reader\nLoads to changable structure\nThat can be rendered back]
    freelancer --> semantic[semantic\nORM mapper of inireader values for quick typed acess to values\nfor reading and writing, without need to map all file structure]
    semantic --> inireader
```
