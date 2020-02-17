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

### Usage
*prerequisites: Clone `geography` project:
https://github.com/pawelWritesCode/geo-vagrant-symfony3
#### 1. Change directory to `geography` root folder.
#### 2. Run command 
```bash
geo-generator geography:expand
```