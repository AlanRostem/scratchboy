cmake -S . -B build_windows_debug -DSCRATCH_DEBUG:bool=true
cd build_windows_debug
ninja.exe
cd ..