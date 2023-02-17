cmake_minimum_required (VERSION 3.7)

set (CMAKE_CXX_STANDARD 20)
set (CMAKE_CXX_STANDARD_REQUIRED ON)
set (CMAKE_CXX_EXTENSIONS OFF)

set (CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror")

set (CMAKE_INSTALL_PREFIX "../../../..")
set (CMAKE_INSTALL_INCLUDEDIR "include")

#set (CMAKE_VERBOSE_MAKEFILE ON)

project(toy-lib LANGUAGES CXX)

add_library (toy-lib SHARED)

add_subdirectory (
	${CMAKE_CURRENT_SOURCE_DIR}/src
	)

set_target_properties (toy-lib PROPERTIES PUBLIC_HEADER
	"src/lib_c_wrapper.h"
	)

message("inc: " ${CMAKE_INSTALL_INCLUDEDIR})

install (TARGETS toy-lib
	CONFIGURATIONS Release
	LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
	PUBLIC_HEADER DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/toy-lib
	)
