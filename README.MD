GEOGRAPHY EXPANDER
===================
Simple CLI for expanding Geography API. 

### Building binary 
* Make sure you have added `$GOPATH/bin/`  to your `$PATH`

Run script `build.sh` from `backend-go/` dir 
```bash
    ./build.sh
```
CLI binary will be placed in `$GOPATH/bin/` dir, under name `geo-generator`.
It will be also available in `./bin/` folder.

### Usage
*prerequisites: Clone `geography` project:
https://github.com/pawelWritesCode/geo-vagrant-symfony3
#### 1. Change directory to `geography` root folder.
#### 2. Run one of available commands 
```bash
geo-generator geography:expand
geo-generator geography:shrink
geo-generator geography:rename
```