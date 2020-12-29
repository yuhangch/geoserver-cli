# geoserver-cli [WIP]

## Done:

- server config 
- workspace  (CRUD)
- datastore (shapefile upload,  R D)
- layer ( R D)
- style (edit in vim (test on fish shell), R U)

## Install:

```shell
$ go get github.com/yuhangch/geoserver-cli
$ gctl  #alias geoserver-cli to gctl
```


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

use command layer to manage layers 

- List (alias: ls) :list all layers in a workspace 

  ```shell
  $ gctl layer ls                                                                                                                                                                           
  layers:
    - 0 bonn:roads 
    - 1 tiger:giant_polygon 
    - 2 tiger:poi 
    - 3 tiger:poly_landmarks 
    - 4 tiger:tiger_roads 
    - 5 nurc:Arc_Sample 
    - 6 nurc:Img_Sample 
  $ gctl layer ls tiger                                                                                                                            
  [tiger] layers:
    - 0 giant_polygon 
    - 1 poi 
    - 2 poly_landmarks 
    - 3 tiger_roads 
  ```

- Delete (alias: rm remove): delete a layer, -r to recurse delete

  ```shell
  $ gctl layer delete bonn:roads -r
  ```

### Style

//TODO
