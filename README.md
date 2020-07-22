# geoserver-cli [WIP]

## Done:

- server config 
- workspace  (CRUD)
- datastore (shapefile upload, RD)
- layer (RD)
- style (edit in vim,RU)


## Usage

### Server

use command server to manage server configuration

- create (alias: add ,new) : create a new server

  ```shell
  $ gctl server create server1 -u http://geoserver.org/geoserver -n admin -p geoserver
  ```

- List (alias: ls) :list current server saved in configuration

  ```shell
  $ gctl server list
  Servers:
  - 0 server1 http://localhost:8080/geoserver-2.16 
  - 1 server2 http://localhost:8080/geoserver
  ```

- Activate (alias: set) :set active server

  ```shell
  $ gctl server activate server1
  ```

- Delete (alias: rm remove): delete a server in configuration

  ```shell
  $ gctl server delete server1
  ```

  

### Workspace

use command workspace to manage workspace

- create (alias: add ,new) : create a workspace

  ```shell
  $ gctl workspace create testws
  ```

- List (alias: ls) :list workspaces

  ```shell
  $ gctl workspace list
  Workspaces:
    - 0 bonn 
    - 1 cite 
    - 2 tiger 
    - 3 nurc 
    - 4 sde 
  ```

- Rename (alias: mv) : rename a workspace

  ```shell
  $ gctl workspace rename bonn bonnn
  ```

- Delete (alias: rm remove): delete a workspace

  ```shell
  $ gctl workspace delete bonn
  ```

  

### Datastore

use command datastore to manage data store

- create (alias: add ,new) : create a new (Shapefile) datastore

  ```shell
  $ gctl datastore create bonn:roads -f Archive.zip --configure=all
  ```

- List (alias: ls) :list current datastore in a workspace 

  ```shell
  $ gctl datastore ls --workspace bonn                                                                                                                                                                           
  DataStores:
    - 0 road 
  $ gctl datastore ls bonn                                                                                                                            
  DataStores:
    - 0 road 
  ```

- Delete (alias: rm remove): delete a datastore

  ```shell
  $ gctl datastore delete bonn:roads -r
  ```

  

### Layer  

//TODO 

### Style

//TODO