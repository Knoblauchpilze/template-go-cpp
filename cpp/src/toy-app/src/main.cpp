
# include <iostream>
# include <Eigen/Eigen>

int
main(int /*argc*/, char** /*argv*/) {
  Eigen::Matrix2f rot;
  rot << 2.0f, 0.0f, 0.0f, 1.5f;

  Eigen::Vector2f v(1.0f, -1.0f);

  std::cout << "Hello world from cpp!" << std::endl;
  std::cout << "Result: " << std::endl << rot * v << std::endl;

  return EXIT_SUCCESS;
}
