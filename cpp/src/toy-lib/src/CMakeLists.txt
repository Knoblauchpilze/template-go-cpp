
target_sources (toy-lib PRIVATE
	${CMAKE_CURRENT_SOURCE_DIR}/Wrapper.cpp
	)

target_include_directories (toy-lib PUBLIC
	"${CMAKE_CURRENT_SOURCE_DIR}"
	)
