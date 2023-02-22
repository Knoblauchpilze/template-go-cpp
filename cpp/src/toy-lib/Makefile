
PROC_COUNT ?= 8

install: release
	cd build/Release && cmake --install . --config Release --prefix ../../../..

release:
	mkdir -p build/Release
	cd build/Release && cmake -DCMAKE_BUILD_TYPE=Release ../.. && make -j ${PROC_COUNT}

debug:
	mkdir -p build/Debug
	cd build/Debug && cmake -DCMAKE_BUILD_TYPE=Debug ../.. && make -j ${PROC_COUNT}

clean:
	rm -rf build
