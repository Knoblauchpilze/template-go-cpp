cmake_minimum_required (VERSION 3.7)

set (CMAKE_CXX_STANDARD 20)
set (CMAKE_CXX_STANDARD_REQUIRED ON)
set (CMAKE_CXX_EXTENSIONS OFF)

set (CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror")

set (CMAKE_INSTALL_BINDIR "bin")

#set (CMAKE_VERBOSE_MAKEFILE ON)

project(toy-app LANGUAGES CXX)

find_package (Eigen3 3.4 REQUIRED NO_MODULE)

add_executable (toy-app)

add_subdirectory (
	${CMAKE_CURRENT_SOURCE_DIR}/src
	)

target_include_directories (toy-app PRIVATE
	"../../include"
	)

target_link_directories (toy-app PUBLIC
	"../../lib"
	)

target_link_libraries (toy-app PUBLIC
	Eigen3::Eigen
	toy-lib
	)

install (TARGETS toy-app
	CONFIGURATIONS Release
	RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
	)
