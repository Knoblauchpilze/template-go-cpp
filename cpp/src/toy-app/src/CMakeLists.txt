
target_sources (toy-app PRIVATE
	${CMAKE_CURRENT_SOURCE_DIR}/main.cpp
	)

target_include_directories (toy-app PUBLIC
	"${CMAKE_CURRENT_SOURCE_DIR}"
	)
