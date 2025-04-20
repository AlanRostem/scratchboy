DIR="build_linux_debug"
cmake -S . -B $DIR -DSCRATCH_DEBUG:bool=true
make --directory=$DIR