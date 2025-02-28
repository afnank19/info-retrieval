#ifndef VECTOR_MATH_H
#define VECTOR_MATH_H
#include <vector>

class VectorMath {
private:
public:
    VectorMath(/* args */);
    static double dot(std::vector<float> v1, std::vector<float> v2);
    static double magnitude(std::vector<float> v);
};


#endif
