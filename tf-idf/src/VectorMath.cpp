#include "../include/VectorMath.hpp"
#include <iostream>
#include <cmath>

VectorMath::VectorMath() {
    printf("Hello to Vector Math");
}

// v1 and v2 must be the same size
double VectorMath::dot(std::vector<float> v1, std::vector<float> v2) {
    double result = 0.0;

    if (v1.size() != v2.size()) {
        return -1.0f;
    }

    for (int i = 0; i < v1.size(); i++) {
        result += v1[i] * v2[i];
    }
    
    return result;
}

double VectorMath::magnitude(std::vector<float> v) {
    double result = 0.0;
    for (const auto& val: v) {
        result += val * val;
    }

    return sqrt(result);
}