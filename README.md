 
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
    darktool --> cmd
    cmd[cmd\nUser Commands to CLI interface]
    cmd --> validator[validator\nlints freelancer configs to strict format]
    validator --> parser[parser\nProvides static typed access to parsedFreelancer configs]
    parser --> freelancer[freelancer\nFreelancer Data Parsers\npackage reflects\nFreelancer File Structure]
    validator --> denormalizer[denormalizer\nDenormalizes parsed data for more\nhuman readable view of freelancer configs]
    denormalizer --> parser
    parser --> filefind[filefind\nfinds freelancer files]
    freelancer --> inireader[inireader\nUniversal freelancer ini format reader\nLoads to changable structure\nThat can be rendered back]
    freelancer --> semantic[ORM mapper of inireader values for quick typed acess to values\nfor reading and writing, without need to map all file structure]
    semantic --> inireader
```

## Contributors

- [@dd84ai](https://github.com/dd84ai) // coding
- [@Groshyr](https://github.com/Groshyr) // spark of inspiration for project birth + beta tester + feature requester + domain expert
